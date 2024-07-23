package private

import (
	"encoding/json"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/permissions"
)

func init() {
	commands.LoadedInternal.Add(
		builder.NewCommand("restriction").
			Middleware(middleware.RequirePermissions(permissions.CreateSoftUserRestriction, permissions.RemoveSoftRestriction, permissions.CreateHardUserRestriction, permissions.RemoveHardRestriction)).
			Options(
				builder.NewOption("view", discordgo.ApplicationCommandOptionSubCommand).
					Options(
						builder.NewOption("user_id", discordgo.ApplicationCommandOptionString),
						builder.NewOption("restriction_id", discordgo.ApplicationCommandOptionString),
					),
				builder.NewOption("update", discordgo.ApplicationCommandOptionSubCommand).
					Options(
						builder.NewOption("restriction_id", discordgo.ApplicationCommandOptionString).Required(),
						builder.NewOption("expires_at_unix", discordgo.ApplicationCommandOptionNumber).Required(),
					),
			).
			Handler(func(ctx common.Context) error {
				subcommand, subOptions, _ := ctx.Options().Subcommand()

				userID, _ := subOptions.Value("user_id").(string)
				restrictionID, _ := subOptions.Value("restriction_id").(string)

				if userID == "" && restrictionID == "" {
					return ctx.Reply().Send("user id or restriction id needs to be provided")
				}

				switch subcommand {
				case "view":
					var data any
					var err error
					if restrictionID != "" {
						data, err = ctx.Core().Database().GetUserRestriction(ctx.Ctx(), restrictionID)
					}
					if userID != "" {
						data, err = ctx.Core().Database().GetUserRestrictions(ctx.Ctx(), userID)
					}
					if err != nil {
						return ctx.Reply().Send("failed to get restriction", err.Error())
					}
					d, _ := json.MarshalIndent(data, "", "  ")
					return ctx.Reply().Send(string(d))

				case "update":
					data, err := ctx.Core().Database().GetUserRestriction(ctx.Ctx(), restrictionID)
					if err != nil {
						return ctx.Reply().Send("failed to get restriction", err.Error())
					}
					expiresAtUnix, ok := subOptions.Value("expires_at_unix").(float64)
					if !ok {
						return ctx.Reply().Send("failed to parse expires_at_unix")
					}
					expiresAt := time.Unix(int64(expiresAtUnix), 0)
					data.ExpiresAt = expiresAt
					data.AddEvent(ctx.User().ID, "updated expires_at", "expired_at updated from a /restriction update command")

					data, err = ctx.Core().Database().UpdateUserRestriction(ctx.Ctx(), data)
					if err != nil {
						return ctx.Reply().Send("failed to update restriction", err.Error())
					}
					d, _ := json.MarshalIndent(data, "", "  ")
					return ctx.Reply().Send(string(d))

				default:
					return ctx.Error("invalid subcommand: " + subcommand)
				}
			}),
	)
}
