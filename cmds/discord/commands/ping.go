package commands

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
)

var Ping = Command{
	discord.MessageCommandCreate{
		Name: "ping",
	},
	func(ctx *context) error {
		return ctx.event.CreateMessage(discord.MessageCreate{Content: fmt.Sprintf("<@%s>", ctx.User.ID)})
	},
}
