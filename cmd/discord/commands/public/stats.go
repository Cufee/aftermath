package public

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
	"github.com/cufee/aftermath/internal/permissions"
	stats "github.com/cufee/aftermath/internal/stats/client/v1"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("stats").
			Middleware(middleware.RequirePermissions(permissions.UseTextCommands, permissions.UseImageCommands)).
			Options(commands.DefaultStatsOptions...).
			Handler(func(ctx common.Context) error {
				options := commands.GetDefaultStatsOptions(ctx.Options())
				message, valid := options.Validate(ctx)
				if !valid {
					return ctx.Reply().Send(message)
				}

				var accountID string
				var opts = []stats.RequestOption{stats.WithWN8(), stats.WithVehicleID(options.TankID)}

				ioptions := models.DiscordInteractionOptions{
					PeriodStart: options.PeriodStart,
					VehicleID:   options.TankID,
				}

				switch {
				case options.UserID != "":
					// mentioned another user, check if the user has an account linked
					mentionedUser, _ := ctx.Core().Database().GetUserByID(ctx.Ctx(), options.UserID, database.WithConnections(), database.WithSubscriptions(), database.WithContent())
					defaultAccount, hasDefaultAccount := mentionedUser.Connection(models.ConnectionTypeWargaming, map[string]any{"default": true})
					if !hasDefaultAccount {
						return ctx.Reply().Send("stats_error_connection_not_found_vague")
					}
					accountID = defaultAccount.ReferenceID

					if img, content, err := logic.GetBackgroundImageFromRef(ctx.Ctx(), ctx.Core().Database(), accountID); err == nil {
						opts = append(opts, stats.WithBackground(img))
						ioptions.BackgroundContentID = content.ID
					}

				case options.Nickname != "" && options.Server != "":
					// nickname provided and server selected - lookup the account
					account, err := ctx.Core().Fetch().Search(ctx.Ctx(), options.Nickname, options.Server)
					if err != nil {
						if err.Error() == "no results found" {
							return ctx.Reply().Format("stats_error_nickname_not_fount_fmt", options.Nickname, strings.ToUpper(options.Server)).Send()
						}
						return ctx.Err(err)
					}
					accountID = fmt.Sprint(account.ID)

					if img, content, err := logic.GetBackgroundImageFromRef(ctx.Ctx(), ctx.Core().Database(), accountID); err == nil {
						opts = append(opts, stats.WithBackground(img))
						ioptions.BackgroundContentID = content.ID
					}

				default:
					defaultAccount, hasDefaultAccount := ctx.User().Connection(models.ConnectionTypeWargaming, map[string]any{"default": true})
					if !hasDefaultAccount {
						return ctx.Reply().Send("command_stats_help_message")
					}
					// command used without options, but user has a default connection
					accountID = defaultAccount.ReferenceID

					background, _ := ctx.User().Content(models.UserContentTypePersonalBackground)
					if img, err := logic.UserContentToImage(background); err == nil {
						opts = append(opts, stats.WithBackground(img))
						ioptions.BackgroundContentID = background.ID
					}
				}

				image, meta, err := ctx.Core().Stats(ctx.Locale()).PeriodImage(context.Background(), accountID, options.PeriodStart, opts...)
				if err != nil {
					return ctx.Err(err)
				}

				ioptions.AccountID = accountID
				button, saveErr := saveInteractionData(ctx, "stats", ioptions)
				if saveErr != nil {
					// nil button will not cause an error and will be ignored
					log.Err(err).Str("interactionId", ctx.ID()).Str("command", "session").Msg("failed to save discord interaction")
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return ctx.Err(err)
				}

				var timings []string
				if ctx.User().HasPermission(permissions.UseDebugFeatures) {
					timings = append(timings, "```")
					for name, duration := range meta.Timings {
						timings = append(timings, fmt.Sprintf("%s: %v", name, duration.Milliseconds()))
					}
					timings = append(timings, "```")
				}

				return ctx.Reply().WithAds().Hint(message).File(buf.Bytes(), "stats_command_by_aftermath.png").Component(button).Text(timings...).Send()
			}),
	)
}
