package router

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/middleware"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/logic"
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

type command struct {
	requested *builder.Command
	current   *discordgo.ApplicationCommand
	cached    *models.ApplicationCommand
}

/*
Updates all loaded commands using the Discord REST API
  - any commands that are loaded will be created/updated
  - any commands that were not loaded will be deleted
*/
func (r *Router) UpdateLoadedCommands(ctx context.Context) error {
	var commandByName = make(map[string]command)
	for _, cmd := range r.commands {
		if cmd.Type != builder.CommandTypeChat {
			continue
		}
		commandByName[cmd.Name] = command{requested: &cmd}
	}

	current, err := r.restClient.GetGlobalApplicationCommands(ctx)
	if err != nil {
		return err
	}

	var currentIDs []string
	for _, cmd := range current {
		command := commandByName[cmd.Name]
		command.current = &cmd
		commandByName[cmd.Name] = command

		currentIDs = append(currentIDs, cmd.ID)
	}

	cachedCommands, err := r.core.Database().GetCommandsByID(context.Background(), currentIDs...)
	if err != nil {
		return err
	}

	for _, cmd := range cachedCommands {
		command := commandByName[cmd.Name]
		command.cached = &cmd
		commandByName[cmd.Name] = command
	}

	for _, data := range commandByName {
		switch {
		case data.requested != nil && data.current == nil:
			log.Debug().Str("name", data.requested.Name).Msg("creating a global command")
			hash, err := logic.HashAny(data.requested.ApplicationCommand)
			if err != nil {
				return err
			}
			command, err := r.restClient.CreateGlobalApplicationCommand(ctx, data.requested.ApplicationCommand)
			if err != nil {
				return err
			}
			err = r.core.Database().UpsertCommands(context.Background(), models.ApplicationCommand{ID: command.ID, Name: command.Name, Hash: hash, Version: command.Version})
			if err != nil {
				return err
			}
			log.Debug().Str("name", data.requested.Name).Str("id", command.ID).Msg("created and cached a global command")

		case data.requested == nil && data.current != nil:
			log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("deleting a global command")
			err := r.restClient.DeleteGlobalApplicationCommand(ctx, data.current.ID)
			if err != nil {
				return err
			}
			log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("deleted a global command")

		case data.current != nil && data.requested != nil:
			hash, err := logic.HashAny(data.requested.ApplicationCommand)
			if err != nil {
				return err
			}
			if data.cached == nil || data.cached.Hash != hash {
				log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("updating a global command")
				hash, err := logic.HashAny(data.requested.ApplicationCommand)
				if err != nil {
					return err
				}
				command, err := r.restClient.UpdateGlobalApplicationCommand(ctx, data.current.ID, data.requested.ApplicationCommand)
				if err != nil {
					return err
				}
				err = r.core.Database().UpsertCommands(context.Background(), models.ApplicationCommand{ID: command.ID, Name: command.Name, Hash: hash, Version: command.Version})
				if err != nil {
					return err
				}
				log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("updated global command")
			} else {
				log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("global command is up to date")
			}

		default:
			log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("global command does not require any action")
		}
	}

	return nil
}
