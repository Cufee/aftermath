package discord

import (
	"net/http"

	"github.com/cufee/aftermath/cmds/core"

	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/router"
)

func NewRouterHandler(coreClient core.Client, token string, publicKey string, commands ...builder.Command) (http.HandlerFunc, error) {
	rt, err := router.NewRouter(coreClient, token, publicKey)
	if err != nil {
		return nil, err
	}

	rt.LoadCommands(commands...)

	err = rt.UpdateLoadedCommands()
	if err != nil {
		return nil, err
	}

	return rt.HTTPHandler()
}
