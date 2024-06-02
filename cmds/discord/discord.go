package discord

import (
	"net/http"

	"github.com/cufee/aftermath/cmds/discord/commands"
	"github.com/cufee/aftermath/cmds/discord/router"
)

func NewRouterHandler(token string, publicKey string) (http.HandlerFunc, error) {
	rt, err := router.NewRouter(token, publicKey)
	if err != nil {
		return nil, err
	}

	rt.LoadCommands(commands.Ping)

	// rt.UpdateCommands()
	// if err != nil {
	// 	return nil, err
	// }

	return rt.HTTPHandler()
}
