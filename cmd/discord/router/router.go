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

type applyRouterOption func(*routerOptions)

type routerOptions struct {
	restClientOptions []rest.ClientOption
}

func WithRestClientOptions(opts ...rest.ClientOption) applyRouterOption {
	return func(options *routerOptions) {
		options.restClientOptions = append(options.restClientOptions, opts...)
	}
}

func NewRouter(coreClient core.Client, token string, requestValidator RequestValidator, opts ...applyRouterOption) (*router, error) {
	var options routerOptions
	for _, apply := range opts {
		apply(&options)
	}

	restClient, err := rest.NewClient(token, options.restClientOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new rest client :%w", err)
	}

	r := router{
		core:       coreClient,
		token:      token,
		restClient: restClient,
		validator:  requestValidator,
	}

	return &r, nil
}

type router struct {
	core core.Client

	token      string
	restClient *rest.Client

	middleware []middleware.MiddlewareFunc
	commands   []builder.Command

	validator RequestValidator
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
