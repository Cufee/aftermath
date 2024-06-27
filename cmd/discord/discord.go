package discord

import (
	"net/http"

	"github.com/cufee/aftermath/cmd/core"

	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/router"
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
