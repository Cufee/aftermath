package public

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/constants"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/permissions"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("links").
			Middleware(middleware.RequirePermissions(permissions.UseTextCommands, permissions.CreatePersonalConnection, permissions.RemovePersonalConnection, permissions.UpdatePersonalConnection)).
			Ephemeral().
			Options(
				builder.NewOption("add", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_option_links_add_name"), builder.SetDescKey("command_option_links_add_desc")).
					Options(commands.NicknameOption.Required()),
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
			Handler(func(ctx common.Context) error {
				// handle subcommand
				subcommand, subOptions, _ := ctx.Options().Subcommand()
				switch subcommand {
				default:
					return ctx.Error("received an unexpected subcommand: " + subcommand)

				case "verify":
					server, _ := ctx.Options().Value("server").(string)
					realm := strings.ToLower(server)
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

					updated, err := logic.UpdateDefaultUserConnection(ctx.Ctx(), ctx.Core().Database(), ctx.User().ID, accountID)
					if err != nil {
						if errors.Is(err, logic.ErrConnectionNotFound) {
							return ctx.Reply().Send("links_error_connection_not_found")
						}
						return ctx.Err(err)
					}

					var content = []string{fmt.Sprintf(ctx.Localize("command_links_set_default_successfully_fmt"), nickname, realm)}
					if updated.Verified != true {
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

					for _, conn := range ctx.User().Connections {
						if conn.Type != models.ConnectionTypeWargaming {
							continue
						}
						if conn.ReferenceID == accountID {
							err := ctx.Core().Database().DeleteUserConnection(ctx.Ctx(), ctx.User().ID, conn.ID)
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
					for _, conn := range ctx.User().Connections {
						if conn.Type != models.ConnectionTypeWargaming {
							continue
						}
						linkedAccounts = append(linkedAccounts, conn.ReferenceID)
						if conn.Selected {
							currentDefault = conn.ReferenceID
						}
					}

					accounts, err := ctx.Core().Database().GetAccounts(ctx.Ctx(), linkedAccounts)
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
					for _, conn := range ctx.User().Connections {
						if conn.Type != models.ConnectionTypeWargaming {
							continue
						}
						wgConnections = append(wgConnections, conn)
					}
					if len(wgConnections) >= 3 {
						return ctx.Reply().Send("links_error_too_many_connections")
					}

					options := commands.GetDefaultStatsOptions(subOptions)
					message, valid := options.Validate(ctx)
					if !valid {
						return ctx.Reply().Send(message)
					}
					if options.AccountID == "" {
						return ctx.Reply().Send("links_error_no_account_selected")
					}

					account, err := ctx.Core().Fetch().Account(ctx.Ctx(), options.AccountID)
					if err != nil {
						return ctx.Reply().Send("nickname_autocomplete_not_found")
					}

					var found bool
					for _, conn := range wgConnections {
						if conn.ReferenceID == account.ID {
							conn.Verified = false
							found = true
						}
						conn.Selected = conn.ReferenceID == account.ID

						_, err := ctx.Core().Database().UpsertUserConnection(ctx.Ctx(), conn)
						if err != nil {
							return ctx.Err(err)
						}
					}
					if !found {
						_, err := ctx.Core().Database().UpsertUserConnection(ctx.Ctx(), models.UserConnection{
							Type:        models.ConnectionTypeWargaming,
							Verified:    false,
							Selected:    true,
							ReferenceID: fmt.Sprint(account.ID),
							UserID:      ctx.User().ID,
						})
						if err != nil {
							return ctx.Err(err)
						}
					}

					go func(id string) {
						c, cancel := context.WithTimeout(context.Background(), time.Second)
						defer cancel()
						_, _ = ctx.Core().Fetch().Account(c, id) // Make sure the account is cached
					}(fmt.Sprint(account.ID))

					return ctx.Reply().Format("command_links_linked_successfully_fmt", account.Nickname, options.Realm.String()).Send()
				}
			}),
	)
}
