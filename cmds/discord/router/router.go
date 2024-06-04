package router

import (
	"fmt"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/discord/commands/builder"
	"github.com/cufee/aftermath/cmds/discord/middleware"
	"github.com/cufee/aftermath/cmds/discord/rest"
)

func NewRouter(coreClient core.Client, token, publicKey string) (*Router, error) {
	restClient, err := rest.NewClient(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new rest client :%w", err)
	}

	return &Router{
		core:       coreClient,
		token:      token,
		publicKey:  publicKey,
		restClient: restClient,
	}, nil
}

type Router struct {
	core core.Client

	token      string
	publicKey  string
	restClient *rest.Client

	middleware []middleware.MiddlewareFunc
	commands   []builder.Command
}

/*
Loads commands into the router, does not update bot commands through Discord API
*/
func (r *Router) LoadCommands(commands ...builder.Command) {
	r.commands = append(r.commands, commands...)
}

/*
Loads interactions into the router
*/
func (r *Router) LoadMiddleware(middleware ...middleware.MiddlewareFunc) {
	r.middleware = append(r.middleware, middleware...)
}

/*
Updates all loaded commands using the Discord REST API
*/
func (r *Router) UpdateLoadedCommands() error {
	// if len(r.commands) < 1 {
	// 	return errors.New("no commands to update")
	// }

	// client, err := r.client()
	// if err != nil {
	// 	return err
	// }

	// var options []discord.ApplicationCommandCreate
	// for _, cmd := range r.commands {
	// 	options = append(options, cmd.ApplicationCommandCreate)

	// 	d, _ := json.MarshalIndent(cmd.ApplicationCommandCreate, "", "  ")
	// 	println(string(d))

	// }

	// if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), options); err != nil {
	// 	return err
	// }
	return nil
}
