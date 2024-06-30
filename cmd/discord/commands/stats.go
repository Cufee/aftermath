package commands

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
	render "github.com/cufee/aftermath/internal/stats/render/common/v1"
	"github.com/rs/zerolog/log"
)

func init() {
	LoadedPublic.add(
		builder.NewCommand("stats").
			Options(defaultStatsOptions...).
			Handler(func(ctx *common.Context) error {
				options := getDefaultStatsOptions(ctx)
				message, valid := options.Validate(ctx)
				if !valid {
					return ctx.Reply().Send(message)
				}

				var accountID string
				var backgroundURL = "static://bg-default"

				switch {
				case options.UserID != "":
					// mentioned another user, check if the user has an account linked
					mentionedUser, _ := ctx.Core.Database().GetUserByID(ctx.Context, options.UserID, database.WithConnections(), database.WithSubscriptions(), database.WithContent())
					defaultAccount, hasDefaultAccount := mentionedUser.Connection(models.ConnectionTypeWargaming)
					if !hasDefaultAccount {
						return ctx.Reply().Send("stats_error_connection_not_found_vague")
					}
					accountID = defaultAccount.ReferenceID
					// TODO: Get user background

				case options.Nickname != "" && options.Server != "":
					// nickname provided and server selected - lookup the account
					account, err := ctx.Core.Fetch().Search(ctx.Context, options.Nickname, options.Server)
					if err != nil {
						if err.Error() == "no results found" {
							return ctx.Reply().Format("stats_error_nickname_not_fount_fmt", options.Nickname, strings.ToUpper(options.Server)).Send()
						}
						return ctx.Err(err)
					}
					accountID = fmt.Sprint(account.ID)

				default:
					defaultAccount, hasDefaultAccount := ctx.User.Connection(models.ConnectionTypeWargaming)
					if !hasDefaultAccount {
						return ctx.Reply().Send("stats_error_nickname_or_server_missing")
					}
					// command used without options, but user has a default connection
					accountID = defaultAccount.ReferenceID
					// TODO: Get user background
				}

				image, meta, err := ctx.Core.Render(ctx.Locale).Period(context.Background(), accountID, options.PeriodStart, render.WithBackground(backgroundURL))
				if err != nil {
					return ctx.Err(err)
				}

				button, saveErr := saveInteractionData(ctx, "stats", models.DiscordInteractionOptions{
					BackgroundImageURL: backgroundURL,
					PeriodStart:        options.PeriodStart,
					AccountID:          accountID,
				})
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
				if ctx.User.HasPermission(permissions.UseDebugFeatures) {
					timings = append(timings, "```")
					for name, duration := range meta.Timings {
						timings = append(timings, fmt.Sprintf("%s: %v", name, duration.Milliseconds()))
					}
					timings = append(timings, "```")
				}

				return ctx.Reply().File(buf.Bytes(), "stats_command_by_aftermath.png").Component(button).Text(timings...).Send()
			}),
	)
}
