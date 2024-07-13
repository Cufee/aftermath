package commands

import (
	"fmt"
	"slices"
	"strings"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/search"
	prepare "github.com/cufee/aftermath/internal/stats/prepare/common/v1"
)

func init() {
	LoadedPublic.add(
		builder.NewCommand("autocomplete_linked_accounts").
			ComponentType(func(s string) bool {
				var keys []string
				keys = append(keys, "autocomplete_links_favorite_selected", "autocomplete_links_remove_selected") // links
				keys = append(keys, "autocomplete_my_session_account", "autocomplete_my_stats_account")           // my
				keys = append(keys, "autocomplete_widget_account")                                                // widget
				return slices.Contains(keys, s)
			}).
			Handler(func(ctx *common.Context) error {
				var currentDefault string
				var linkedAccounts []string
				for _, conn := range ctx.User.Connections {
					if conn.Type != models.ConnectionTypeWargaming {
						continue
					}
					linkedAccounts = append(linkedAccounts, conn.ReferenceID)
					if def, _ := conn.Metadata["default"].(bool); def {
						currentDefault = conn.ReferenceID
					}

				}
				if len(linkedAccounts) < 1 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("links_error_no_accounts_linked"), Value: "error#links_error_no_accounts_linked"})
				}

				accounts, err := ctx.Core.Database().GetAccounts(ctx.Context, linkedAccounts)
				if err != nil {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("links_error_no_accounts_linked"), Value: "error#links_error_no_accounts_linked"})
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
				return ctx.Reply().Choices(opts...)
			}),
	)

	LoadedPublic.add(
		builder.NewCommand("autocomplete_tank_search").
			ComponentType(func(s string) bool {
				var keys []string
				keys = append(keys, "autocomplete_stats_tank", "autocomplete_session_tank")       // stats/session
				keys = append(keys, "autocomplete_my_session_tank", "autocomplete_my_stats_tank") // my
				return slices.Contains(keys, s)
			}).
			Handler(func(ctx *common.Context) error {
				options := getDefaultStatsOptions(ctx.Options())
				// if the tank was already found, return the tank
				if options.TankID != "" {
					vehicle, ok := search.GetVehicleFromCache(ctx.Locale, options.TankID)
					if !ok {
						return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_found"), Value: "error#stats_autocomplete_not_found"})
					}
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: fmt.Sprintf("%s %s", prepare.IntToRoman(vehicle.Tier), vehicle.Name(ctx.Locale)), Value: fmt.Sprintf("valid#vehicle#%s", vehicle.ID)})
				}

				if len(options.TankSearch) < 3 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_enough_length"), Value: "error#stats_autocomplete_not_enough_length"})
				}

				vehicles, ok := search.SearchVehicles(ctx.Locale, options.TankSearch, 5)
				if !ok || len(vehicles) < 1 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_autocomplete_not_found"), Value: "error#stats_autocomplete_not_found"})
				}

				var opts []*discordgo.ApplicationCommandOptionChoice
				for _, v := range vehicles {
					opts = append(opts, &discordgo.ApplicationCommandOptionChoice{Name: fmt.Sprintf("%s %s", prepare.IntToRoman(v.Tier), v.Name(ctx.Locale)), Value: fmt.Sprintf("valid#vehicle#%s", v.ID)})
				}
				return ctx.Reply().Choices(opts...)
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
