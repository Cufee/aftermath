package private

import (
	"time"

	"github.com/cufee/aftermath/internal/json"

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
				builder.NewOption("timeframe_hours", discordgo.ApplicationCommandOptionNumber),
				builder.NewOption("all_time", discordgo.ApplicationCommandOptionBoolean),
			).
			Handler(func(ctx common.Context) error {
				snowflake, _ := ctx.Options().Value("snowflake").(string)
				interactionID, _ := ctx.Options().Value("id").(string)
				userID, _ := ctx.Options().Value("user_id").(string)
				timeframeHours, _ := ctx.Options().Value("timeframe_hours").(float64)
				allTime, _ := ctx.Options().Value("all_time").(bool)

				if timeframeHours < 0 {
					return ctx.Reply().Send("timeframe_hours cannot be a negative value")
				}

				var data any
				var err error
				if snowflake != "" {
					data, err = ctx.Core().Database().FindDiscordInteractions(ctx.Ctx(), database.WithSnowflake(snowflake))
				}
				if interactionID != "" {
					data, err = ctx.Core().Database().GetDiscordInteraction(ctx.Ctx(), interactionID)
				}
				if userID != "" {
					query := []database.InteractionQuery{database.WithUserID(userID)}
					if !allTime {
						if timeframeHours <= 0 {
							timeframeHours = 24
						}
						query = append(query, database.WithSentAfter(time.Now().Add(-time.Duration(timeframeHours*float64(time.Hour)))))
					}
					data, err = ctx.Core().Database().FindDiscordInteractions(ctx.Ctx(), query...)
				}
				if err != nil {
					return ctx.Reply().Send(err.Error())
				}

				d, _ := json.MarshalIndent(data, "", "  ")
				return ctx.Reply().File(d, "requests.json").Send()
			}),
	)
}
