package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/prisma/db"
)

type ApplicationCommand struct {
	ID      string
	Hash    string
	Name    string
	Version string
}

func (c ApplicationCommand) fromModel(model db.ApplicationCommandModel) ApplicationCommand {
	return ApplicationCommand{
		ID:      model.ID,
		Name:    model.Name,
		Hash:    model.OptionsHash,
		Version: model.Version,
	}
}

func (c *client) GetCommandsByID(ctx context.Context, commandIDs ...string) ([]ApplicationCommand, error) {
	if len(commandIDs) < 1 {
		return nil, nil
	}

	models, err := c.prisma.ApplicationCommand.FindMany(db.ApplicationCommand.ID.In(commandIDs)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var commands []ApplicationCommand
	for _, model := range models {
		commands = append(commands, ApplicationCommand{}.fromModel(model))
	}
	return commands, nil
}

func (c *client) GetCommandsByHash(ctx context.Context, commandHashes ...string) ([]ApplicationCommand, error) {
	if len(commandHashes) < 1 {
		return nil, nil
	}

	models, err := c.prisma.ApplicationCommand.FindMany(db.ApplicationCommand.OptionsHash.In(commandHashes)).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var commands []ApplicationCommand
	for _, model := range models {
		commands = append(commands, ApplicationCommand{}.fromModel(model))
	}
	return commands, nil
}

func (c *client) UpsertCommands(ctx context.Context, commands ...ApplicationCommand) error {
	if len(commands) < 1 {
		return nil
	}

	var tx []db.PrismaTransaction
	for _, cmd := range commands {
		tx = append(tx, c.prisma.ApplicationCommand.UpsertOne(db.ApplicationCommand.ID.Equals(cmd.ID)).
			Create(
				db.ApplicationCommand.ID.Set(cmd.ID),
				db.ApplicationCommand.Name.Set(cmd.Name),
				db.ApplicationCommand.Version.Set(cmd.Version),
				db.ApplicationCommand.OptionsHash.Set(cmd.Hash),
			).
			Update(
				db.ApplicationCommand.Name.Set(cmd.Name),
				db.ApplicationCommand.Version.Set(cmd.Version),
				db.ApplicationCommand.OptionsHash.Set(cmd.Hash),
			).Tx(),
		)
	}
	return c.prisma.Prisma.Transaction(tx...).Exec(ctx)
}
