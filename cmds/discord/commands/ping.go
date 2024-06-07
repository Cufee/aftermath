package commands

import (
	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/common"
)

func init() {
	Loaded.add(
		builder.NewCommand("ping").
			Params(builder.SetDescKey("Pong!")).
			Handler(func(ctx *common.Context) error {
				return ctx.Reply("Pong!")
			}),
	)
}
