package router

import (
	"encoding/hex"
	"encoding/json"
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
		core:         coreClient,
		token:        token,
		publicKey:    publicKey,
		commands:     make(map[string]commands.Command),
		interactions: make(map[string]interactions.Interaction),
	}, nil
}

type Router struct {
	core      core.Client
	token     string
	publicKey string
	botClient bot.Client

	middleware []middleware.MiddlewareFunc

	commands     map[string]commands.Command
	interactions map[string]interactions.Interaction
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
	for _, cmd := range commands {
		r.commands[cmd.CommandName()] = cmd
	}
}

/*
Loads interactions into the router
*/
func (r *Router) LoadInteractions(interactions ...interactions.Interaction) {
	for _, intc := range interactions {
		r.interactions[intc.Name] = intc
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
		return errors.New("no commands to update")
	}

	client, err := r.client()
	if err != nil {
		return err
	}

	var options []discord.ApplicationCommandCreate
	for _, cmd := range r.commands {
		options = append(options, cmd.ApplicationCommandCreate)

		d, _ := json.MarshalIndent(cmd.ApplicationCommandCreate, "", "  ")
		println(string(d))

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
