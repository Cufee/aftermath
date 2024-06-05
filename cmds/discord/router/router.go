package router

import (
	"fmt"
	"slices"

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
  - any commands that are loaded will be created/updated
  - any commands that were not loaded will be deleted
*/
func (r *Router) UpdateLoadedCommands() error {
	toOverwrite := make(map[string]discordgo.ApplicationCommand)

	for _, cmd := range r.commands {
		if cmd.Type != builder.CommandTypeChat {
			continue
		}
		toOverwrite[cmd.Name] = cmd.ApplicationCommand
	}

	current, err := r.restClient.GetGlobalApplicationCommands()
	if err != nil {
		return err
	}

	var toDelete []string
	var currentCommands []string
	for _, command := range current {
		currentCommands = append(currentCommands, command.Name)
		if _, ok := toOverwrite[command.Name]; !ok {
			toDelete = append(toDelete, command.ID)
			continue
		}

		newCmd := toOverwrite[command.Name]
		newCmd.ID = command.ID

		toOverwrite[command.Name] = newCmd
	}

	for _, cmd := range r.commands {
		if cmd.Type != builder.CommandTypeChat {
			continue
		}
		if !slices.Contains(currentCommands, cmd.Name) {
			// it will be created during the bulk overwrite call
			continue
		}
		// Check if the command needs to be updated
		// if true {
		delete(toOverwrite, cmd.Name)
		// 	continue
		// }
	}

	for _, id := range toDelete {
		err := r.restClient.DeleteGlobalApplicationCommand(id)
		if err != nil {
			return fmt.Errorf("failed to delete a command: %w", err)
		}
	}

	if len(toOverwrite) < 1 {
		return nil
	}

	var commandUpdates []discordgo.ApplicationCommand
	for _, cmd := range toOverwrite {
		commandUpdates = append(commandUpdates, cmd)
	}

	err = r.restClient.OverwriteGlobalApplicationCommands(commandUpdates)
	if err != nil {
		return fmt.Errorf("failed to bulk update application commands: %w", err)
	}

	return nil
}
