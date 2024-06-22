package commands

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/render"
	"github.com/cufee/aftermath/internal/stats/render/assets"
)

func init() {
	Loaded.add(
		builder.NewCommand("stats").
			Options(defaultStatsOptions...).
			Handler(func(ctx *common.Context) error {
				options := getDefaultStatsOptions(ctx)
				message, valid := options.Validate(ctx)
				if !valid {
					return ctx.Reply(message)
				}

				var accountID string
				background, _ := assets.GetLoadedImage("bg-default")

				switch {
				case options.UserID != "":
					// mentioned another user, check if the user has an account linked
					mentionedUser, _ := ctx.Core.Database().GetUserByID(ctx.Context, options.UserID, database.WithConnections(), database.WithSubscriptions(), database.WithContent())
					defaultAccount, hasDefaultAccount := mentionedUser.Connection(models.ConnectionTypeWargaming)
					if !hasDefaultAccount {
						return ctx.Reply("stats_error_connection_not_found_vague")
					}
					accountID = defaultAccount.ReferenceID
					// TODO: Get user background

				case options.Nickname != "" && options.Server != "":
					// nickname provided and server selected - lookup the account
					account, err := ctx.Core.Fetch().Search(ctx.Context, options.Nickname, options.Server)
					if err != nil {
						if err.Error() == "no results found" {
							return ctx.ReplyFmt("stats_error_nickname_not_fount_fmt", options.Nickname, strings.ToUpper(options.Server))
						}
						return ctx.Err(err)
					}
					accountID = fmt.Sprint(account.ID)

				default:
					defaultAccount, hasDefaultAccount := ctx.User.Connection(models.ConnectionTypeWargaming)
					if !hasDefaultAccount {
						return ctx.Reply("stats_error_nickname_or_server_missing")
					}
					// command used without options, but user has a default connection
					accountID = defaultAccount.ReferenceID
					// TODO: Get user background
				}

				image, _, err := ctx.Core.Render(ctx.Locale).Period(context.Background(), accountID, options.PeriodStart, render.WithBackground(background))
				if err != nil {
					return ctx.Err(err)
				}

				var buf bytes.Buffer
				err = image.PNG(&buf)
				if err != nil {
					return ctx.Err(err)
				}

				return ctx.File(&buf, "stats_command_by_aftermath.png")
			}),
	)
}
