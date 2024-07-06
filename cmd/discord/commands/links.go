package commands

import (
	"context"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
)

func init() {
	LoadedPublic.add(
		builder.NewCommand("links").
			Ephemeral().
			Options(
				builder.NewOption("add", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_links_add_name"), builder.SetDescKey("command_option_links_add_desc")).
					Options(
						builder.NewOption("nickname", discordgo.ApplicationCommandOptionString).
							Required().
							Min(5).
							Max(30).
							Params(
								builder.SetNameKey("common_option_stats_nickname_name"),
								builder.SetDescKey("common_option_stats_nickname_description"),
							),
						builder.NewOption("server", discordgo.ApplicationCommandOptionString).
							Required().
							Params(
								builder.SetNameKey("common_option_stats_realm_name"),
								builder.SetDescKey("common_option_stats_realm_description"),
							).
							Choices(
								builder.NewChoice("realm_na", "NA").Params(builder.SetNameKey("common_label_realm_na")),
								builder.NewChoice("realm_eu", "EU").Params(builder.SetNameKey("common_label_realm_eu")),
								builder.NewChoice("realm_as", "AS").Params(builder.SetNameKey("common_label_realm_as")),
							),
					),
				builder.NewOption("favorite", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_links_fav_name"), builder.SetDescKey("command_option_links_fav_desc")).
					Options(
						builder.NewOption("selected", discordgo.ApplicationCommandOptionString).
							Autocomplete().
							Required(),
					),
				builder.NewOption("verify", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_links_verify_name"), builder.SetDescKey("command_option_links_verify_desc")).
					Options(
						builder.NewOption("server", discordgo.ApplicationCommandOptionString).
							Required().
							Params(
								builder.SetNameKey("common_option_stats_realm_name"),
								builder.SetDescKey("common_option_stats_realm_description"),
							).
							Choices(
								builder.NewChoice("realm_na", "NA").Params(builder.SetNameKey("common_label_realm_na")),
								builder.NewChoice("realm_eu", "EU").Params(builder.SetNameKey("common_label_realm_eu")),
								builder.NewChoice("realm_as", "AS").Params(builder.SetNameKey("common_label_realm_as")),
							),
					),

				builder.NewOption("view", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_links_view_name"), builder.SetDescKey("command_option_links_view_desc")),
				builder.NewOption("remove", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_links_remove_name"), builder.SetDescKey("command_option_links_remove_desc")).
					Options(
						builder.NewOption("selected", discordgo.ApplicationCommandOptionString).
							Autocomplete().
							Required(),
					),
			).
			Handler(func(ctx *common.Context) error {
				// handle subcommand
				subcommand, subOptions, _ := ctx.Options().Subcommand()
				switch subcommand {
				default:
					return ctx.Error("received an unexpected subcommand: " + subcommand)

				case "verify":
					options := getDefaultStatsOptions(subOptions)
					realm := strings.ToLower(options.Server)
					loginURL := fmt.Sprintf("%s/r/verify/%s", constants.FrontendURL, realm)
					return ctx.Reply().Format("command_links_verify_response_fmt", ctx.Localize("common_label_realm_"+realm), loginURL).Send()

				case "favorite":
					value, _ := subOptions.Value("selected").(string)
					parts := strings.Split(value, "#")
					if len(parts) != 4 || parts[0] != "valid" {
						return ctx.Reply().Send("links_error_connection_not_found")
					}
					accountID := parts[1]
					nickname := parts[2]
					realm := parts[3]

					var found bool
					var newConnectionVerified bool
					for _, conn := range ctx.User.Connections {
						if conn.Type != models.ConnectionTypeWargaming {
							continue
						}
						if conn.ReferenceID == accountID {
							newConnectionVerified, _ = conn.Metadata["verified"].(bool)
							found = true
						}
						conn.Metadata["default"] = conn.ReferenceID == accountID

						_, err := ctx.Core.Database().UpdateConnection(ctx.Context, conn)
						if err != nil {
							if database.IsNotFound(err) {
								return ctx.Reply().Send("links_error_connection_not_found")
							}
							return ctx.Err(err)
						}
					}
					if !found {
						return ctx.Reply().Send("links_error_connection_not_found")
					}

					var content = []string{fmt.Sprintf(ctx.Localize("command_links_set_default_successfully_fmt"), nickname, realm)}
					if !newConnectionVerified {
						content = append(content, ctx.Localize("command_links_verify_cta"))
					}
					return ctx.Reply().Text(content...).Send()

				case "remove":
					value, _ := subOptions.Value("selected").(string)
					parts := strings.Split(value, "#")
					if len(parts) != 4 || parts[0] != "valid" {
						return ctx.Reply().Send("links_error_connection_not_found_selected")
					}
					accountID := parts[1]

					for _, conn := range ctx.User.Connections {
						if conn.Type != models.ConnectionTypeWargaming {
							continue
						}
						if conn.ReferenceID == accountID {
							err := ctx.Core.Database().DeleteConnection(ctx.Context, conn.ID)
							if err != nil {
								return ctx.Err(err)
							}
							return ctx.Reply().Send("command_links_unlinked_successfully")
						}
					}
					return ctx.Reply().Send("links_error_connection_not_found_selected")

				case "view":
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

					accounts, err := ctx.Core.Database().GetAccounts(ctx.Context, linkedAccounts)
					if err != nil && !database.IsNotFound(err) {
						return ctx.Err(err)
					}
					if len(accounts) == 0 {
						return ctx.Reply().Send("stats_error_connection_not_found_personal")
					}

					var longestName int
					for _, a := range accounts {
						if l := utf8.RuneCountInString(a.Nickname); l > longestName {
							longestName = l
						}
					}

					var nicknames []string
					for _, a := range accounts {
						nicknames = append(nicknames, accountToRow(a, longestName, currentDefault == a.ID))
					}
					return ctx.Reply().Send(strings.Join(nicknames, "\n"))

				case "add":
					var wgConnections []models.UserConnection
					for _, conn := range ctx.User.Connections {
						if conn.Type != models.ConnectionTypeWargaming {
							continue
						}
						wgConnections = append(wgConnections, conn)
					}
					if len(wgConnections) >= 3 {
						return ctx.Reply().Send("links_error_too_many_connections")
					}

					options := getDefaultStatsOptions(subOptions)
					message, valid := options.Validate(ctx)
					if !valid {
						return ctx.Reply().Send(message)
					}

					account, err := ctx.Core.Fetch().Search(ctx.Context, options.Nickname, options.Server)
					if err != nil {
						if err.Error() == "no results found" {
							return ctx.Reply().Format("stats_error_nickname_not_fount_fmt", options.Nickname, strings.ToUpper(options.Server)).Send()
						}
						return ctx.Err(err)
					}

					var found bool
					for _, conn := range wgConnections {
						if conn.ReferenceID == fmt.Sprint(account.ID) {
							conn.Metadata["verified"] = false
							found = true
						}
						conn.Metadata["default"] = conn.ReferenceID == fmt.Sprint(account.ID)

						_, err = ctx.Core.Database().UpsertConnection(ctx.Context, conn)
						if err != nil {
							return ctx.Err(err)
						}
					}
					if !found {
						meta := make(map[string]any)
						meta["verified"] = false
						meta["default"] = true
						_, err = ctx.Core.Database().UpsertConnection(ctx.Context, models.UserConnection{
							Type:        models.ConnectionTypeWargaming,
							ReferenceID: fmt.Sprint(account.ID),
							UserID:      ctx.User.ID,
							Metadata:    meta,
						})
						if err != nil {
							return ctx.Err(err)
						}
					}

					go func(id string) {
						c, cancel := context.WithTimeout(context.Background(), time.Second)
						defer cancel()
						_, _ = ctx.Core.Fetch().Account(c, id) // Make sure the account is cached
					}(fmt.Sprint(account.ID))

					return ctx.Reply().Format("command_links_linked_successfully_fmt", account.Nickname, strings.ToUpper(options.Server)).Send()
				}
			}),
	)

	LoadedPublic.add(
		builder.NewCommand("autocomplete_links").
			ComponentType(func(s string) bool {
				return s == "autocomplete_links"
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
