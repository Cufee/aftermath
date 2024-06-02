package discord

import (
	"net/http"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/discord/commands"
	"github.com/cufee/aftermath/cmds/discord/common"
	"github.com/cufee/aftermath/cmds/discord/middleware"
	"github.com/cufee/aftermath/cmds/discord/router"
)

func NewRouterHandler(coreClient core.Client, token string, publicKey string) (http.HandlerFunc, error) {
	rt, err := router.NewRouter(coreClient, token, publicKey)
	if err != nil {
		return nil, err
	}

	rt.LoadCommands(commands.Ping)

	// should always be loaded last
	rt.LoadMiddleware(middleware.FetchUser(common.ContextKeyUser, coreClient.DB))

	// rt.UpdateCommands()
	// if err != nil {
	// 	return nil, err
	// }

	return rt.HTTPHandler()
}
