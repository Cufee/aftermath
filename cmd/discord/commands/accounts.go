package commands

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/rs/zerolog/log"
)

func init() {
	LoadedPublic.add(
		builder.NewCommand("accounts").
			Ephemeral().
			Options(
				builder.NewOption("default", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_accounts_default_name"), builder.SetDescKey("command_option_accounts_default_description")).
					Options(
						builder.NewOption("selected", discordgo.ApplicationCommandOptionString).
							Autocomplete().
							Required(),
					),
				builder.NewOption("unlink", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_accounts_unlink_name"), builder.SetDescKey("command_option_accounts_unlink_description")).
					Options(
						builder.NewOption("selected", discordgo.ApplicationCommandOptionString).
							Autocomplete().
							Required(),
					),
				builder.NewOption("linked", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_accounts_linked_name"), builder.SetDescKey("command_option_accounts_linked_description")),
			).
			Handler(func(ctx *common.Context) error {
				// handle subcommand
				subcommand, subOptions, _ := ctx.Options().Subcommand()
				switch subcommand {
				default:
					return ctx.Error("received an unexpected subcommand: " + subcommand)

				case "default":
					value, _ := subOptions.Value("selected").(string)
					parts := strings.Split(value, "#")
					if len(parts) != 4 || parts[0] != "valid" {
						return ctx.Reply().Send("accounts_error_connection_not_found")
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
								return ctx.Reply().Send("accounts_error_connection_not_found")
							}
							return ctx.Err(err)
						}
					}
					if !found {
						return ctx.Reply().Send("accounts_error_connection_not_found")
					}

					var content = []string{fmt.Sprintf(ctx.Localize("command_accounts_set_default_successfully_fmt"), nickname, realm)}
					if !newConnectionVerified {
						content = append(content, ctx.Localize("command_accounts_verify_cta"))
					}
					return ctx.Reply().Text(content...).Send()

				case "unlink":
					value, _ := subOptions.Value("selected").(string)
					parts := strings.Split(value, "#")
					if len(parts) != 4 || parts[0] != "valid" {
						return ctx.Reply().Send("accounts_error_connection_not_found_selected")
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
							return ctx.Reply().Send("command_accounts_unlinked_successfully")
						}
					}
					return ctx.Reply().Send("accounts_error_connection_not_found_selected")

				case "linked":
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
				}
			}),
	)

	LoadedPublic.add(
		builder.NewCommand("autocomplete_accounts").
			ComponentType(func(s string) bool {
				return s == "autocomplete_accounts"
			}).
			Handler(func(ctx *common.Context) error {
				command, _, ok := ctx.Options().Subcommand()
				if !ok {
					log.Error().Str("command", "autocomplete_accounts").Msg("interaction is not a subcommand")
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("accounts_error_connection_not_found"), Value: "error#accounts_error_connection_not_found"})
				}
				if command != "default" && command != "unlink" {
					log.Error().Str("command", command).Msg("invalid subcommand received in autocomplete_accounts")
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("accounts_error_connection_not_found"), Value: "error#accounts_error_connection_not_found"})
				}

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
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("accounts_error_connection_not_found"), Value: "error#accounts_error_connection_not_found"})
				}

				accounts, err := ctx.Core.Database().GetAccounts(ctx.Context, linkedAccounts)
				if err != nil {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("accounts_error_connection_not_found"), Value: "error#accounts_error_connection_not_found"})
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
