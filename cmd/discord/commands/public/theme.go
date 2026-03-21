package public

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/render/themes"
)

func buildThemeChoices() []builder.OptionChoice {
	var choices []builder.OptionChoice
	for _, id := range themes.AvailableThemes() {
		choices = append(choices, builder.NewChoice(id, id))
	}
	return choices
}

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("theme").
			Middleware(middleware.RequirePermissions(permissions.UseTextCommands)).
			Ephemeral().
			Params(builder.SetNameKey("command_theme_name"), builder.SetDescKey("command_theme_description")).
			Options(
				builder.NewOption("select", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_theme_option_select_name"), builder.SetDescKey("command_theme_option_select_description")).
					Options(
						builder.NewOption("theme", discordgo.ApplicationCommandOptionString).
							Choices(buildThemeChoices()...),
					),
				builder.NewOption("clear", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_theme_option_clear_name"), builder.SetDescKey("command_theme_option_clear_description")),
			).
			Handler(func(ctx common.Context) error {
				subcommand, subOptions, _ := ctx.Options().Subcommand()
				switch subcommand {
				default:
					return ctx.Reply().Send("command_theme_current_fmt")

				case "select":
					selected, hasSelection := common.GetOption[string](subOptions, "theme")
					if !hasSelection {
						currentTheme, hasTheme := ctx.User().Content(models.UserContentTypeThemePreference)
						var currentID string
						if hasTheme {
							currentID = string(currentTheme.Value)
						} else {
							currentID = "default"
						}

						ids := themes.AvailableThemes()
						sort.Strings(ids)

						var lines []string
						for _, id := range ids {
							name := ctx.Localize(fmt.Sprintf("command_theme_select_option_theme_choice_%s_name", id))
							if name == "" {
								name = id
							}
							if id == currentID {
								lines = append(lines, "⬢ "+name)
							} else {
								lines = append(lines, "⬡ "+name)
							}
						}

						return ctx.Reply().Format("command_theme_current_fmt", strings.Join(lines, "\n")).Send()
					}

					if selected == "default" {
						existing, err := ctx.Core().Database().GetUserContentFromRef(ctx.Ctx(), ctx.User().ID, models.UserContentTypeThemePreference)
						if err != nil && !database.IsNotFound(err) {
							return ctx.Err(err, common.ApplicationError)
						}
						if !database.IsNotFound(err) {
							err = ctx.Core().Database().DeleteUserContent(ctx.Ctx(), existing.ID)
							if err != nil {
								return ctx.Err(err, common.ApplicationError)
							}
						}
						return ctx.Reply().Send("command_theme_reset_success")
					}

					if _, ok := themes.GetTheme(selected); !ok {
						return ctx.Reply().IsError(common.UserError).Send("command_theme_not_found")
					}

					existing, err := ctx.Core().Database().GetUserContentFromRef(ctx.Ctx(), ctx.User().ID, models.UserContentTypeThemePreference)
					if err != nil && !database.IsNotFound(err) {
						return ctx.Err(err, common.ApplicationError)
					}
					if database.IsNotFound(err) {
						existing = models.UserContent{
							UserID:      ctx.User().ID,
							ReferenceID: ctx.User().ID,
						}
					}

					existing.Type = models.UserContentTypeThemePreference
					existing.Value = []byte(selected)
					_, err = ctx.Core().Database().UpsertUserContent(ctx.Ctx(), existing)
					if err != nil {
						return ctx.Err(err, common.ApplicationError)
					}

					return ctx.Reply().Format("command_theme_set_success_fmt", selected).Send()

				case "clear":
					existing, err := ctx.Core().Database().GetUserContentFromRef(ctx.Ctx(), ctx.User().ID, models.UserContentTypeThemePreference)
					if err != nil && !database.IsNotFound(err) {
						return ctx.Err(err, common.ApplicationError)
					}
					if !database.IsNotFound(err) {
						err = ctx.Core().Database().DeleteUserContent(ctx.Ctx(), existing.ID)
						if err != nil {
							return ctx.Err(err, common.ApplicationError)
						}
					}
					return ctx.Reply().Send("command_theme_reset_success")
				}
			}),
	)
}
