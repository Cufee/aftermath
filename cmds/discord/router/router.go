package router

import (
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/cmds/discord/commands"
	"github.com/cufee/aftermath/cmds/discord/interactions"
	"github.com/cufee/aftermath/cmds/discord/middleware"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handlers"
	"github.com/disgoorg/disgo/httpserver"
)

func NewRouter(coreClient core.Client, token, publicKey string) (*Router, error) {
	return &Router{
		core:                coreClient,
		token:               token,
		publicKey:           publicKey,
		commandHandlers:     make(map[string]commands.Handler),
		interactionHandlers: make(map[string]interactions.Handler),
	}, nil
}

type Router struct {
	core      core.Client
	token     string
	publicKey string
	botClient bot.Client

	middleware []middleware.MiddlewareFunc

	commands            []commands.Command
	interactions        []interactions.Interaction
	commandHandlers     map[string]commands.Handler
	interactionHandlers map[string]interactions.Handler
}

/*
Initializes the client using the router settings or returns an existing client
*/
func (r *Router) client() (bot.Client, error) {
	if r.botClient != nil {
		return r.botClient, nil
	}

	client, err := disgo.New(r.token,
		bot.WithEventListenerFunc(r.handleEvent),
	)
	r.botClient = client

	return r.botClient, err
}

/*
Returns a handler for the current router
*/
func (r *Router) HTTPHandler() (http.HandlerFunc, error) {
	client, err := r.client()
	if err != nil {
		return nil, err
	}

	hexDecodedKey, err := hex.DecodeString(r.publicKey)
	if err != nil {
		return nil, errors.New("invalid public key")
	}

	handler := httpserver.HandleInteraction(hexDecodedKey, client.Logger(), handlers.DefaultHTTPServerEventHandlerFunc(client))
	return handler, err
}

/*
Loads commands into the router, does not update bot commands through Discord API
*/
func (r *Router) LoadCommands(commands ...commands.Command) {
	r.commands = append(r.commands, commands...)
	for _, cmd := range commands {
		r.commandHandlers[cmd.CommandName()] = cmd.Handler
	}
}

/*
Loads interactions into the router
*/
func (r *Router) LoadInteractions(interactions ...interactions.Interaction) {
	r.interactions = append(r.interactions, interactions...)
	for _, intc := range interactions {
		r.interactionHandlers[intc.Name] = intc.Handler
	}
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
func (r *Router) UpdateCommands() error {
	if len(r.commands) < 1 {
		return nil
	}

	client, err := r.client()
	if err != nil {
		return err
	}

	var options []discord.ApplicationCommandCreate
	for _, cmd := range r.commands {
		options = append(options, cmd)
	}

	if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), options); err != nil {
		return err
	}
	return nil
}

// func addContext[E bot.Event](next func(ctx context.Context, event E)) (f func(e E)) {
// 	return func(e E) {
// 		ctx := context.Background()
// 		next(ctx, e)
// 	}
// }

// func fetchUser[E bot.Event](next func(ctx context.Context, event E)) func(ctx context.Context, event E) {
// 	return func(ctx context.Context, event E) {
// 		next(ctx, event)
// 	}
// }
