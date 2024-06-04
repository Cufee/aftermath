package router

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
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
	commands   []command
}

type command struct {
	data    discordgo.ApplicationCommand
	handler builder.CommandHandler
	match   func(string) bool
}

/*
Initializes the client using the router settings or returns an existing client
*/
func (r *Router) client() (*rest.Client, error) {
	// if r.botClient != nil {
	// 	return r.botClient, nil
	// }

	// client, err := disgo.New(r.token,
	// 	bot.WithEventListenerFunc(r.handleEvent),
	// )
	// r.botClient = client

	return nil, nil
}

/*
Loads commands into the router, does not update bot commands through Discord API
*/
func (r *Router) LoadCommands(commands ...func() (discordgo.ApplicationCommand, func(s string) bool, builder.CommandHandler)) {
	for _, build := range commands {
		d, m, h := build()
		r.commands = append(r.commands, command{d, h, m})
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

func (r *Router) handleInteraction(ctx context.Context, interaction discordgo.Interaction, reply chan<- discordgo.InteractionResponseData, done chan<- struct{}) {
	defer func() {
		done <- struct{}{}
	}()

	reply <- discordgo.InteractionResponseData{
		Content: "pong",
	}
}
