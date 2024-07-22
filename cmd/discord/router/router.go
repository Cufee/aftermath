package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/logic"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/cmd/discord/rest"
)

type Router interface {
	discord.Commander
	HTTPHandler() (http.HandlerFunc, error)
}

func NewRouter(coreClient core.Client, token, publicKey string) (*router, error) {
	restClient, err := rest.NewClient(token)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new rest client :%w", err)
	}

	return &router{
		core:       coreClient,
		token:      token,
		publicKey:  publicKey,
		restClient: restClient,
	}, nil
}

type router struct {
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
func (r *router) LoadCommands(commands ...builder.Command) {
	r.commands = append(r.commands, commands...)
}

/*
Loads interactions into the router
*/
func (r *router) LoadMiddleware(middleware ...middleware.MiddlewareFunc) {
	r.middleware = append(r.middleware, middleware...)
}

/*
Updates all loaded commands using the Discord REST API
  - any commands that are loaded will be created/updated
  - any commands that were not loaded will be deleted
*/
func (r *router) UpdateLoadedCommands(ctx context.Context) error {
	return logic.UpdateCommands(ctx, r.core.Database(), r.restClient, r.commands)
}
