package public

import (
	"bytes"
	"context"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/permissions"
	stats "github.com/cufee/aftermath/internal/stats/client/v1"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/pkg/errors"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("my").
			Middleware(middleware.RequirePermissions(permissions.UseTextCommands, permissions.UseImageCommands)).
			Params(builder.SetNameKey("command_my_name"), builder.SetDescKey("command_my_description")).
			Options(
				builder.NewOption("stats", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_my_stats_name"), builder.SetDescKey("command_my_stats_description")).
					Options(
						commands.DaysOption,
						commands.VehicleOption,
						builder.NewOption("account", discordgo.ApplicationCommandOptionString).
							Params(builder.SetNameKey("command_option_my_account_name"), builder.SetDescKey("command_option_my_account_description")).
							Autocomplete(),
					),
				builder.NewOption("session", discordgo.ApplicationCommandOptionSubCommand).
					Params(builder.SetNameKey("command_my_session_name"), builder.SetDescKey("command_my_session_description")).
					Options(
						commands.DaysOption,
						commands.VehicleOption,
						builder.NewOption("account", discordgo.ApplicationCommandOptionString).
							Params(builder.SetNameKey("command_option_my_account_name"), builder.SetDescKey("command_option_my_account_description")).
							Autocomplete(),
					),
			).
			Handler(func(ctx common.Context) error {
				subcommand, subOptions, _ := ctx.Options().Subcommand()
				options := commands.GetDefaultStatsOptions(ctx.Options())
				message, valid := options.Validate(ctx)
				if !valid {
					return ctx.Reply().Send(message)
				}

				var accountID string
				var opts = []stats.RequestOption{stats.WithWN8(), stats.WithVehicleID(options.TankID)}

				ioptions := statsOptions{StatsOptions: options}

				if img, content, err := logic.GetAccountBackgroundImage(ctx.Ctx(), ctx.Core().Database(), accountID); err == nil {
					opts = append(opts, stats.WithBackground(img, true))
					ioptions.BackgroundID = content.ID
				} else {
					background, _ := ctx.User().Content(models.UserContentTypePersonalBackground)
					if img, err := logic.UserContentToImage(background); err == nil {
						opts = append(opts, stats.WithBackground(img, true))
						ioptions.BackgroundID = background.ID
					}
				}

				value, _ := subOptions.Value("account").(string)
				parts := strings.Split(value, "#")
				if len(parts) == 4 && parts[0] == "valid" {
					accountID = parts[1]
				} else {
					defaultAccount, hasDefaultAccount := ctx.User().Connection(models.ConnectionTypeWargaming, nil, utils.Pointer(true))
					if !hasDefaultAccount {
						return ctx.Reply().Send("my_error_no_account_linked")
					}
					accountID = defaultAccount.ReferenceID
				}

				var err error
				var image stats.Image
				switch subcommand {
				case "stats":
					image, _, err = ctx.Core().Stats(ctx.Locale()).PeriodImage(context.Background(), accountID, options.PeriodStart, opts...)
				case "session":
					image, _, err = ctx.Core().Stats(ctx.Locale()).SessionImage(context.Background(), accountID, options.PeriodStart, opts...)
				default:
					return ctx.Error("invalid subcommand in /my - " + subcommand)
				}
				if err != nil {
					if errors.Is(err, blitzstars.ErrServiceUnavailable) {
						return ctx.Reply().
							Hint(ctx.InteractionID()).
							Component(discordgo.ActionsRow{Components: []discordgo.MessageComponent{common.ButtonJoinPrimaryGuild(ctx.Localize("buttons_have_a_question_question"))}}).
							Send("blitz_stars_error_service_down")
					}
					if errors.Is(err, stats.ErrAccountNotTracked) || (errors.Is(err, fetch.ErrSessionNotFound) && options.Days < 1) {
						return ctx.Reply().Send("session_error_account_was_not_tracked")
					}
					if errors.Is(err, fetch.ErrSessionNotFound) {
						return ctx.Reply().Send("session_error_no_session_for_period")
					}
					return ctx.Err(err)
				}

				ioptions.AccountID = accountID
				button, saveErr := ioptions.refreshButton(ctx, subcommand)
				if saveErr != nil {
					// nil button will not cause an error and will be ignored
					log.Err(err).Str("interactionId", ctx.ID()).Str("command", "session").Msg("failed to save discord interaction")
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return ctx.Err(err)
				}
				return ctx.Reply().File(buf.Bytes(), "session_command_by_aftermath.png").Component(button).Send()
			}),
	)
}
