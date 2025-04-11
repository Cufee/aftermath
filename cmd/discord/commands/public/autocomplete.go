package public

import (
	"context"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/glossary"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("autocomplete_linked_accounts").
			ComponentType(func(s string) bool {
				var keys []string
				keys = append(keys, "autocomplete_links_favorite_selected", "autocomplete_links_remove_selected") // links
				keys = append(keys, "autocomplete_my_session_account", "autocomplete_my_career_account")          // my
				keys = append(keys, "autocomplete_widget_account")                                                // widget
				return slices.Contains(keys, s)
			}).
			Handler(func(ctx common.Context) error {
				var currentDefault string
				var linkedAccounts []string
				for _, conn := range ctx.User().Connections {
					if conn.Type != models.ConnectionTypeWargaming {
						continue
					}
					linkedAccounts = append(linkedAccounts, conn.ReferenceID)
					if conn.Selected {
						currentDefault = conn.ReferenceID
					}

				}
				if len(linkedAccounts) < 1 {
					return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("links_error_no_accounts_linked"), Value: "error#links_error_no_accounts_linked"}).Send()
				}

				accounts, err := ctx.Core().Database().GetAccounts(ctx.Ctx(), linkedAccounts)
				if err != nil {
					return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("links_error_no_accounts_linked"), Value: "error#links_error_no_accounts_linked"}).Send()
				}

				var opts []*discordgo.ApplicationCommandOptionChoice
				for _, a := range accounts {
					opts = append(opts, &discordgo.ApplicationCommandOptionChoice{Name: accountToRow(a, currentDefault == a.ID), Value: fmt.Sprintf("valid#%s#%s#%s", a.ID, a.Nickname, a.Realm)})
				}
				return ctx.Reply().Choices(opts...).Send()
			}),
	)

	commands.LoadedPublic.Add(
		builder.NewCommand("autocomplete_tank_search").
			ComponentType(func(s string) bool {
				var keys []string
				keys = append(keys, "autocomplete_career_tank", "autocomplete_stats_tank", "autocomplete_session_tank") // stats/session
				keys = append(keys, "autocomplete_my_session_tank", "autocomplete_my_career_tank")                      // my
				return slices.Contains(keys, s)
			}).
			Handler(func(ctx common.Context) error {
				options := commands.GetDefaultStatsOptions(ctx.Options())
				// if the tank was already found, return the tank
				if options.TankID != "" {
					vehicle, ok := glossary.GetVehicleFromCache(ctx.Locale(), options.TankID)
					if !ok {
						return ctx.Reply().IsError(common.ApplicationError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_found"), Value: "error#stats_autocomplete_not_found"}).Send()
					}
					return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: fmt.Sprintf("%s %s", logic.IntToRoman(vehicle.Tier), vehicle.Name(ctx.Locale())), Value: fmt.Sprintf("valid#vehicle#%s", vehicle.ID)}).Send()
				}

				if len(options.TankSearch) < 3 {
					return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_enough_length"), Value: "error#stats_autocomplete_not_enough_length"}).Send()
				}

				vehicles, ok := glossary.SearchVehicles(ctx.Locale(), options.TankSearch, 5)
				if !ok || len(vehicles) < 1 {
					return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_found"), Value: "error#stats_autocomplete_not_found"}).Send()
				}

				var opts []*discordgo.ApplicationCommandOptionChoice
				for _, v := range vehicles {
					content := fmt.Sprintf("%s %s", logic.IntToRoman(v.Tier), v.Name(ctx.Locale()))
					opts = append(opts, &discordgo.ApplicationCommandOptionChoice{Name: content, Value: fmt.Sprintf("valid#vehicle#%s", v.ID)})
				}
				return ctx.Reply().Choices(opts...).Send()
			}),
	)

	commands.LoadedPublic.Add(
		builder.NewCommand("autocomplete_account_search").
			ComponentType(func(s string) bool {
				var keys []string
				keys = append(keys, "autocomplete_manage_accounts_search_nickname")                                                 // manage
				keys = append(keys, "autocomplete_career_nickname", "autocomplete_stats_nickname", "autocomplete_session_nickname") // stats/session
				keys = append(keys, "autocomplete_links_add_nickname")                                                              // links
				return slices.Contains(keys, s)
			}).
			Handler(func(ctx common.Context) error {
				options := commands.GetDefaultStatsOptions(ctx.Options())
				// if the account was already found, return the account
				if options.AccountID != "" {
					fCtx, cancel := context.WithTimeout(ctx.Ctx(), time.Millisecond*2500)
					defer cancel()

					account, err := ctx.Core().Fetch().Account(fCtx, options.AccountID)
					if err != nil {
						if os.IsTimeout(err) || fCtx.Err() != nil {
							log.Err(fCtx.Err()).Msg("account fetch for autocompletion timed out")
							return ctx.Reply().IsError(common.ApplicationError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("wargaming_error_outage_short"), Value: "error#wargaming_error_outage_short"}).Send()
						}
						return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_not_found"), Value: "error#nickname_autocomplete_not_found"}).Send()
					}
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: fmt.Sprintf("[%s] %s", account.Realm, account.Nickname), Value: fmt.Sprintf("valid#account#%s#%s", account.ID, account.Realm)}).Send()
				}

				if len(options.NicknameSearch) < 3 {
					return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_not_enough_length"), Value: "error#nickname_autocomplete_not_enough_length"}).Send()
				}
				if !wargaming.ValidatePlayerNickname(options.NicknameSearch) {
					return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_invalid_input"), Value: "error#nickname_autocomplete_invalid_input"}).Send()
				}

				sCtx, cancel := context.WithTimeout(ctx.Ctx(), time.Millisecond*2500)
				defer cancel()

				accounts, err := ctx.Core().Fetch().BroadSearch(sCtx, options.NicknameSearch, 2)
				if err != nil {
					if os.IsTimeout(err) || sCtx.Err() != nil {
						log.Err(sCtx.Err()).Msg("broad search accounts for autocompletion timed out")
						return ctx.Reply().IsError(common.ApplicationError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("wargaming_error_outage_short"), Value: "error#wargaming_error_outage_short"}).Send()
					}

					log.Err(err).Msg("failed to broad search accounts")
					return ctx.Reply().IsError(common.ApplicationError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_not_found"), Value: "error#nickname_autocomplete_not_found"}).Send()
				}
				if len(accounts) < 1 {
					return ctx.Reply().IsError(common.UserError).Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_not_found"), Value: "error#nickname_autocomplete_not_found"}).Send()
				}

				slices.SortFunc(accounts, func(a, b fetch.AccountWithRealm) int {
					return strings.Compare(b.Realm.String()+b.Nickname, a.Realm.String()+a.Nickname)
				})

				var opts []*discordgo.ApplicationCommandOptionChoice
				for _, a := range accounts {
					opts = append(opts, &discordgo.ApplicationCommandOptionChoice{Name: fmt.Sprintf("[%s] %s", a.Realm, a.Nickname), Value: fmt.Sprintf("valid#account#%d#%s", a.ID, a.Realm)})
				}
				return ctx.Reply().Choices(opts...).Send()
			}),
	)
}

func accountToRow(account models.Account, isDefault bool) string {
	var row string
	if isDefault {
		row += "⬢ "
	} else {
		row += "⬡ "
	}
	row += "[" + account.Realm.String() + "] "
	row += account.Nickname
	return row
}
