package private

import (
	"encoding/json"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/permissions"
)

func init() {
	commands.LoadedInternal.Add(
		builder.NewCommand("interaction").
			Middleware(middleware.RequirePermissions(permissions.GlobalAdmin)).
			Options(
				builder.NewOption("id", discordgo.ApplicationCommandOptionString),
				builder.NewOption("user_id", discordgo.ApplicationCommandOptionString),
				builder.NewOption("snowflake", discordgo.ApplicationCommandOptionString),
			).
			Handler(func(ctx common.Context) error {
				snowflake, _ := ctx.Options().Value("snowflake").(string)
				interactionID, _ := ctx.Options().Value("id").(string)
				userID, _ := ctx.Options().Value("user_id").(string)

				var data any
				var err error
				if snowflake != "" {
					data, err = ctx.Core().Database().FindDiscordInteractions(ctx.Ctx(), database.WithSnowflake(snowflake))
				}
				if interactionID != "" {
					data, err = ctx.Core().Database().GetDiscordInteraction(ctx.Ctx(), interactionID)
				}
				if userID != "" {
					data, err = ctx.Core().Database().FindDiscordInteractions(ctx.Ctx(), database.WithUserID(userID), database.WithSentAfter(time.Now().Add(-time.Hour*24)), database.WithLimit(10))
				}
				if err != nil {
					return ctx.Reply().Send(err.Error())
				}

				d, _ := json.MarshalIndent(data, "", "  ")
				return ctx.Reply().File(d, "requests.json").Send()
			}),
	)
}
