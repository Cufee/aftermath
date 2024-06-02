package router

import (
	"context"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/cufee/aftermath/cmds/discord/commands"
	"github.com/cufee/aftermath/cmds/discord/interactions"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handlers"
	"github.com/disgoorg/disgo/httpserver"
)

func NewRouter(token, publicKey string) (*Router, error) {
	return &Router{
		token:               token,
		publicKey:           publicKey,
		commandHandlers:     make(map[string]commands.Handler),
		interactionHandlers: make(map[string]interactions.Handler),
	}, nil
}

type Router struct {
	token     string
	publicKey string
	botClient bot.Client

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
		bot.WithEventListenerFunc(addContext(r.commandHandler)),
		bot.WithEventListenerFunc(addContext(r.interactionHandler)),
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
	r.commands = commands
	for _, cmd := range commands {
		r.commandHandlers[cmd.CommandName()] = cmd.Handler
	}
}

/*
Loads interactions into the router
*/
func (r *Router) LoadInteractions(interactions ...interactions.Interaction) {
	r.interactions = interactions
	for _, intc := range interactions {
		r.interactionHandlers[intc.Name] = intc.Handler
	}
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

func addContext[E bot.Event](next func(ctx context.Context, event E)) (f func(e E)) {
	return func(e E) {
		ctx := context.Background()
		next(ctx, e)
	}
}
