package public

import (
	"bytes"
	"context"
	"fmt"

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
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/pkg/errors"
)

func init() {
	commands.LoadedPublic.Add(
		builder.NewCommand("session").
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

				ioptions := statsOptions{StatsOptions: options}

				switch {
				case options.UserID != "":
					// mentioned another user, check if the user has an account linked
					mentionedUser, _ := ctx.Core().Database().GetUserByID(ctx.Ctx(), options.UserID, database.WithConnections(), database.WithSubscriptions(), database.WithContent())
					defaultAccount, hasDefaultAccount := mentionedUser.Connection(models.ConnectionTypeWargaming, nil, utils.Pointer(true))
					if !hasDefaultAccount {
						return ctx.Reply().Send("stats_error_connection_not_found_vague")
					}
					accountID = defaultAccount.ReferenceID

					if img, content, err := logic.GetBackgroundImageFromRef(ctx.Ctx(), ctx.Core().Database(), accountID); err == nil {
						opts = append(opts, stats.WithBackground(img, true))
						ioptions.BackgroundID = content.ID
					}

				case options.AccountID != "":
					// account selected from autocomplete
					accountID = options.AccountID

					if img, content, err := logic.GetBackgroundImageFromRef(ctx.Ctx(), ctx.Core().Database(), accountID); err == nil {
						opts = append(opts, stats.WithBackground(img, true))
						ioptions.BackgroundID = content.ID
					}

				case options.NicknameSearch != "" && options.AccountID == "":
					// nickname provided, but user did not select an option from autocomplete
					accounts, err := accountsFromBadInput(ctx.Ctx(), ctx.Core().Fetch(), options.NicknameSearch)
					if err != nil {
						return ctx.Err(err)
					}
					if len(accounts) == 0 {
						return ctx.Reply().Send("stats_account_not_found")
					}

					realms := make(map[string]struct{})
					for _, a := range accounts {
						realms[a.Realm.String()] = struct{}{}
					}
					if len(realms) > 1 {
						reply, err := realmSelectButtons(ctx, ctx.ID(), accounts)
						if err != nil {
							return ctx.Err(err)
						}
						return reply.Send()
					}

					// one or more options on the same server - just pick the first one
					message = "stats_bad_nickname_input_hint"
					accountID = fmt.Sprint(accounts[0].ID)

				default:
					defaultAccount, hasDefaultAccount := ctx.User().Connection(models.ConnectionTypeWargaming, nil, utils.Pointer(true))
					if !hasDefaultAccount {
						return ctx.Reply().Send("command_session_help_message")
					}
					// command used without options, but user has a default connection
					accountID = defaultAccount.ReferenceID

					background, _ := ctx.User().Content(models.UserContentTypePersonalBackground)
					if img, err := logic.UserContentToImage(background); err == nil {
						opts = append(opts, stats.WithBackground(img, true))
						ioptions.BackgroundID = background.ID
					}
				}

				image, meta, err := ctx.Core().Stats(ctx.Locale()).SessionImage(context.Background(), accountID, options.PeriodStart, opts...)
				if err != nil {
					if errors.Is(err, stats.ErrAccountNotTracked) || (errors.Is(err, fetch.ErrSessionNotFound) && options.Days < 1) {
						return ctx.Reply().Send("session_error_account_was_not_tracked")
					}
					if errors.Is(err, fetch.ErrSessionNotFound) {
						return ctx.Reply().Send("session_error_no_session_for_period")
					}
					return ctx.Err(err)
				}

				ioptions.AccountID = accountID
				button, saveErr := ioptions.refreshButton(ctx, ctx.ID())
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

				return ctx.Reply().WithAds().Hint(message).File(buf.Bytes(), "session_command_by_aftermath.png").Component(button).Text(timings...).Send()
			}),
	)
}
