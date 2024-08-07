package public

import (
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/search"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("autocomplete_linked_accounts").
			ComponentType(func(s string) bool {
				var keys []string
				keys = append(keys, "autocomplete_links_favorite_selected", "autocomplete_links_remove_selected") // links
				keys = append(keys, "autocomplete_my_session_account", "autocomplete_my_stats_account")           // my
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
					if def, _ := conn.Metadata["default"].(bool); def {
						currentDefault = conn.ReferenceID
					}

				}
				if len(linkedAccounts) < 1 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("links_error_no_accounts_linked"), Value: "error#links_error_no_accounts_linked"}).Send()
				}

				accounts, err := ctx.Core().Database().GetAccounts(ctx.Ctx(), linkedAccounts)
				if err != nil {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("links_error_no_accounts_linked"), Value: "error#links_error_no_accounts_linked"}).Send()
				}

				var longestName int
				for _, a := range accounts {
					if l := utf8.RuneCountInString(a.Nickname); l > longestName {
						longestName = l
					}
				}

				var opts []*discordgo.ApplicationCommandOptionChoice
				for _, a := range accounts {
					opts = append(opts, &discordgo.ApplicationCommandOptionChoice{Name: accountToRow(a, longestName, currentDefault == a.ID), Value: fmt.Sprintf("valid#%s#%s#%s", a.ID, a.Nickname, a.Realm)})
				}
				return ctx.Reply().Choices(opts...).Send()
			}),
	)

	commands.LoadedPublic.Add(
		builder.NewCommand("autocomplete_tank_search").
			ComponentType(func(s string) bool {
				var keys []string
				keys = append(keys, "autocomplete_stats_tank", "autocomplete_session_tank")       // stats/session
				keys = append(keys, "autocomplete_my_session_tank", "autocomplete_my_stats_tank") // my
				return slices.Contains(keys, s)
			}).
			Handler(func(ctx common.Context) error {
				options := commands.GetDefaultStatsOptions(ctx.Options())
				// if the tank was already found, return the tank
				if options.TankID != "" {
					vehicle, ok := search.GetVehicleFromCache(ctx.Locale(), options.TankID)
					if !ok {
						return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_found"), Value: "error#stats_autocomplete_not_found"}).Send()
					}
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: fmt.Sprintf("%s %s", prepare.IntToRoman(vehicle.Tier), vehicle.Name(ctx.Locale())), Value: fmt.Sprintf("valid#vehicle#%s", vehicle.ID)}).Send()
				}

				if len(options.TankSearch) < 3 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_enough_length"), Value: "error#stats_autocomplete_not_enough_length"}).Send()
				}

				vehicles, ok := search.SearchVehicles(ctx.Locale(), options.TankSearch, 5)
				if !ok || len(vehicles) < 1 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_found"), Value: "error#stats_autocomplete_not_found"}).Send()
				}

				var opts []*discordgo.ApplicationCommandOptionChoice
				for _, v := range vehicles {
					content := fmt.Sprintf("%s %s", prepare.IntToRoman(v.Tier), v.Name(ctx.Locale()))
					opts = append(opts, &discordgo.ApplicationCommandOptionChoice{Name: content, Value: fmt.Sprintf("valid#vehicle#%s", v.ID)})
				}
				return ctx.Reply().Choices(opts...).Send()
			}),
	)

	commands.LoadedPublic.Add(
		builder.NewCommand("autocomplete_account_search").
			ComponentType(func(s string) bool {
				var keys []string
				keys = append(keys, "autocomplete_manage_accounts_search_nickname")                 // manage
				keys = append(keys, "autocomplete_stats_nickname", "autocomplete_session_nickname") // stats/session
				keys = append(keys, "autocomplete_links_add_nickname")                              // links
				return slices.Contains(keys, s)
			}).
			Handler(func(ctx common.Context) error {
				options := commands.GetDefaultStatsOptions(ctx.Options())
				// if the account was already found, return the account
				if options.AccountID != "" {
					account, err := ctx.Core().Fetch().Account(ctx.Ctx(), options.AccountID)
					if err != nil {
						return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_not_found"), Value: "error#nickname_autocomplete_not_found"}).Send()
					}
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: fmt.Sprintf("[%s] %s", account.Realm, account.Nickname)}).Send()
				}

				if len(options.NicknameSearch) < 5 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_not_enough_length"), Value: "error#nickname_autocomplete_not_enough_length"}).Send()
				}
				if !commands.ValidatePlayerName(options.NicknameSearch) {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_invalid_input"), Value: "error#nickname_autocomplete_invalid_input"}).Send()
				}

				accounts, err := ctx.Core().Fetch().BroadSearch(ctx.Ctx(), options.NicknameSearch)
				if err != nil {
					log.Err(err).Msg("failed to broad search accounts")
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_not_found"), Value: "error#nickname_autocomplete_not_found"}).Send()
				}
				if len(accounts) < 1 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("nickname_autocomplete_not_found"), Value: "error#nickname_autocomplete_not_found"}).Send()
				}

				slices.SortFunc(accounts, func(a, b fetch.AccountWithRealm) int {
					return strings.Compare(b.Realm+b.Nickname, a.Realm+a.Nickname)
				})

				var opts []*discordgo.ApplicationCommandOptionChoice
				for _, a := range accounts {
					opts = append(opts, &discordgo.ApplicationCommandOptionChoice{Name: fmt.Sprintf("[%s] %s", a.Realm, a.Nickname), Value: fmt.Sprintf("valid#account#%d#%s", a.ID, a.Realm)})
				}
				return ctx.Reply().Choices(opts...).Send()
			}),
	)
}

func accountToRow(account models.Account, padding int, isDefault bool) string {
	var row string
	row += "[" + account.Realm + "] "
	row += account.Nickname + strings.Repeat(" ", padding-utf8.RuneCountInString(account.Nickname))
	if isDefault {
		row += " ⭐"
	}
	return row
}
