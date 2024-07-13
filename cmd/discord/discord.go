package discord

import (
	"context"
	"net/http"
	"time"

	"github.com/cufee/aftermath/cmd/core"

	"github.com/cufee/aftermath/cmd/discord/commands"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/router"
)

func NewRouterHandler(coreClient core.Client, token string, publicKey string, cmds ...builder.Command) (http.HandlerFunc, error) {
	rt, err := router.NewRouter(coreClient, token, publicKey)
	if err != nil {
		return nil, err
	}

	rt.LoadCommands(cmds...)
	rt.LoadCommands(commands.Help().Build())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	err = rt.UpdateLoadedCommands(ctx)
	if err != nil {
		return nil, err
	}

	return rt.HTTPHandler()
}
