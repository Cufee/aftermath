package logic

import (
	"context"
	"os"

	"github.com/cufee/aftermath/cmd/discord/commands/builder"
	"github.com/cufee/aftermath/cmd/discord/rest"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/discord"
	"github.com/cufee/aftermath/internal/log"
)

type command struct {
	requested *builder.Command
	current   *discord.ApplicationCommand
	cached    *models.ApplicationCommand
}

func UpdateCommands(ctx context.Context, db database.Client, rest *rest.Client, commands []builder.Command) error {
	var commandByName = make(map[string]command)
	for _, cmd := range commands {
		if cmd.Type != builder.CommandTypeChat {
			continue
		}
		commandByName[cmd.Name] = command{requested: &cmd}
	}

	current, err := rest.GetGlobalApplicationCommands(ctx)
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

	cachedCommands, err := db.GetCommandsByID(context.Background(), currentIDs...)
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
			hash, err := hashAny(data.requested.ApplicationCommand)
			if err != nil {
				return err
			}
			command, err := rest.CreateGlobalApplicationCommand(ctx, data.requested.ApplicationCommand)
			if err != nil {
				return err
			}
			err = db.UpsertCommands(context.Background(), models.ApplicationCommand{ID: command.ID, Name: command.Name, Hash: hash, Version: command.Version})
			if err != nil {
				return err
			}
			log.Debug().Str("name", data.requested.Name).Str("id", command.ID).Msg("created and cached a global command")

		case data.requested == nil && data.current != nil:
			log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("deleting a global command")
			err := rest.DeleteGlobalApplicationCommand(ctx, data.current.ID)
			if err != nil {
				return err
			}
			log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("deleted a global command")

		case data.current != nil && data.requested != nil:
			hash, err := hashAny(data.requested.ApplicationCommand)
			if err != nil {
				return err
			}
			if data.cached == nil || data.cached.Hash != hash || os.Getenv("FORCE_UPDATE_DISCORD_COMMANDS") == "true" {
				log.Debug().Str("name", data.current.Name).Str("id", data.current.ID).Msg("updating a global command")
				hash, err := hashAny(data.requested.ApplicationCommand)
				if err != nil {
					return err
				}
				command, err := rest.UpdateGlobalApplicationCommand(ctx, data.current.ID, data.requested.ApplicationCommand)
				if err != nil {
					return err
				}
				err = db.UpsertCommands(context.Background(), models.ApplicationCommand{ID: command.ID, Name: command.Name, Hash: hash, Version: command.Version})
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
