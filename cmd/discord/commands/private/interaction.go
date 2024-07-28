package private

import (
	"encoding/json"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/common"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/internal/permissions"
)

func init() {
	commands.LoadedInternal.Add(
		builder.NewCommand("interaction").
			Middleware(middleware.RequirePermissions(permissions.GlobalAdmin)).
			Options(
				builder.NewOption("id", discordgo.ApplicationCommandOptionString),
			).
			Handler(func(ctx common.Context) error {
				interactionID, ok := ctx.Options().Value("id").(string)
				if !ok {
					ctx.Reply().Send("invalid invalid interaction id")
				}

				data, err := ctx.Core().Database().GetDiscordInteraction(ctx.Ctx(), interactionID)
				if err != nil {
					return ctx.Reply().Send(err.Error())
				}

				d, _ := json.MarshalIndent(data, "", "  ")
				return ctx.Reply().File(d, "requests.json").Send()
			}),
	)
}
