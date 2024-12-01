package private

import (
	"time"

	"github.com/cufee/aftermath/internal/json"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
)

func init() {
	commands.LoadedInternal.Add(
		builder.NewCommand("requests").
			Middleware(middleware.RequirePermissions(permissions.ContentModerator)).
			Options(
				builder.NewOption("view", discordgo.ApplicationCommandOptionSubCommand).
					Options(
						builder.NewOption("user_id", discordgo.ApplicationCommandOptionString),
						builder.NewOption("request_id", discordgo.ApplicationCommandOptionString),
					),
				builder.NewOption("set_status", discordgo.ApplicationCommandOptionSubCommand).
					Options(
						builder.NewOption("status", discordgo.ApplicationCommandOptionString).Required(),
						builder.NewOption("request_id", discordgo.ApplicationCommandOptionString).Required(),
					),
			).
			Handler(func(ctx common.Context) error {
				subcommand, subOptions, _ := ctx.Options().Subcommand()

				switch subcommand {
				case "view":
					requestID, _ := subOptions.Value("request_id").(string)
					userID, _ := subOptions.Value("user_id").(string)
					if requestID == "" && userID == "" {
						return ctx.Reply().Send("request id or user id needs to be provided")
					}

					var data any
					var err error
					if requestID != "" {
						data, err = ctx.Core().Database().GetModerationRequest(ctx.Ctx(), requestID)
					}
					if userID != "" {
						data, err = ctx.Core().Database().FindUserModerationRequests(ctx.Ctx(), userID, nil, nil, time.Now().Add(-time.Hour*72))
					}
					if err != nil {
						return ctx.Reply().Send("failed to get requests", err.Error())
					}

					d, _ := json.MarshalIndent(data, "", "  ")
					return ctx.Reply().File(d, "requests.json").Send()

				case "set_status":
					requestID, _ := subOptions.Value("request_id").(string)
					if requestID == "" {
						return ctx.Reply().Send("review id needs to be provided")
					}

					status, ok := subOptions.Value("status").(string)
					if !ok {
						return ctx.Reply().Send("status cannot be left blank")
					}

					data, err := ctx.Core().Database().GetModerationRequest(ctx.Ctx(), requestID)
					if err != nil {
						return ctx.Reply().Send("failed to get request", err.Error())
					}

					data.ActionStatus = models.ModerationStatus(status)
					data, err = ctx.Core().Database().UpdateModerationRequest(ctx.Ctx(), data)
					if err != nil {
						return ctx.Reply().Send("failed to update request", err.Error())
					}
					d, _ := json.MarshalIndent(data, "", "  ")
					return ctx.Reply().File(d, "request.json").Send()

				default:
					return ctx.Error("invalid subcommand: " + subcommand)
				}
			}),
	)
}
