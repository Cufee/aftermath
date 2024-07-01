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
	"github.com/lucsky/cuid"
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
					connectionID := parts[1]
					nickname := parts[2]
					realm := parts[3]

					if err := cuid.IsCuid(connectionID); err != nil {
						return ctx.Reply().Send("accounts_error_connection_not_found")
					}

					for _, conn := range ctx.User.Connections {
						if conn.ID != connectionID {
							continue
						}
						conn.Metadata["default"] = true
						_, err := ctx.Core.Database().UpdateConnection(ctx.Context, conn)
						if err != nil {
							if database.IsNotFound(err) {
								return ctx.Reply().Send("accounts_error_connection_not_found")
							}
							return ctx.Err(err)
						}

						var content = []string{fmt.Sprintf(ctx.Localize("command_accounts_set_default_successfully_fmt"), nickname, realm)}
						if verified, _ := conn.Metadata["verified"].(bool); !verified {
							content = append(content, ctx.Localize("command_accounts_verify_cta"))
						}
						return ctx.Reply().Text(content...).Send()
					}
					return ctx.Reply().Send("accounts_error_connection_not_found")

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
						nicknames = append(nicknames, accountToRow(a, longestName, currentDefault == a.ID || len(linkedAccounts) == 1))
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
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_error_connection_not_found_personal"), Value: "error#stats_error_connection_not_found_personal"})
				}
				if command != "default" {
					log.Error().Str("command", command).Msg("invalid subcommand received in autocomplete_accounts")
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_error_connection_not_found_personal"), Value: "error#stats_error_connection_not_found_personal"})
				}

				var currentDefault string
				var linkedAccounts []string
				var accountToConnection = make(map[string]string)
				for _, conn := range ctx.User.Connections {
					if conn.Type != models.ConnectionTypeWargaming {
						continue
					}
					linkedAccounts = append(linkedAccounts, conn.ReferenceID)
					accountToConnection[conn.ReferenceID] = conn.ID
					if def, _ := conn.Metadata["default"].(bool); def {
						currentDefault = conn.ReferenceID
					}

				}
				if len(linkedAccounts) < 1 {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_error_connection_not_found_personal"), Value: "error#stats_error_connection_not_found_personal"})
				}

				accounts, err := ctx.Core.Database().GetAccounts(ctx.Context, linkedAccounts)
				if err != nil {
					return ctx.Reply().Choices(&discordgo.ApplicationCommandOptionChoice{Name: ctx.Localize("stats_error_connection_not_found_personal"), Value: "error#stats_error_connection_not_found_personal"})
				}

				var longestName int
				for _, a := range accounts {
					if l := utf8.RuneCountInString(a.Nickname); l > longestName {
						longestName = l
					}
				}

				var opts []*discordgo.ApplicationCommandOptionChoice
				for _, a := range accounts {
					opts = append(opts, &discordgo.ApplicationCommandOptionChoice{Name: accountToRow(a, longestName, currentDefault == a.ID || len(linkedAccounts) == 1), Value: fmt.Sprintf("valid#%s#%s#%s", accountToConnection[a.ID], a.Nickname, a.Realm)})
				}
				return ctx.Reply().Choices(opts...)
			}),
	)
}

func accountToRow(account models.Account, padding int, isDefault bool) string {
	var row string
	row += account.Nickname + strings.Repeat(" ", padding-utf8.RuneCountInString(account.Nickname))
	row += " [" + account.Realm + "]"
	if isDefault {
		row += " ✅"
	}
	if account.Private {
		row += " 🔒"
	}
	return row
}
