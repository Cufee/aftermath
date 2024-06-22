package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/models"
)

func (c *libsqlClient) GetCommandsByID(ctx context.Context, commandIDs ...string) ([]models.ApplicationCommand, error) {
	// if len(commandIDs) < 1 {
	// 	return nil, nil
	// }

	// models, err := c.prisma.models.ApplicationCommand.FindMany(db.models.ApplicationCommand.ID.In(commandIDs)).Exec(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	var commands []models.ApplicationCommand
	// for _, model := range models {
	// 	commands = append(commands, models.ApplicationCommand{}.fromModel(model))
	// }
	return commands, nil
}

func (c *libsqlClient) GetCommandsByHash(ctx context.Context, commandHashes ...string) ([]models.ApplicationCommand, error) {
	// if len(commandHashes) < 1 {
	// 	return nil, nil
	// }

	// models, err := c.prisma.models.ApplicationCommand.FindMany(db.models.ApplicationCommand.OptionsHash.In(commandHashes)).Exec(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	var commands []models.ApplicationCommand
	// for _, model := range models {
	// 	commands = append(commands, models.ApplicationCommand{}.fromModel(model))
	// }
	return commands, nil
}

func (c *libsqlClient) UpsertCommands(ctx context.Context, commands ...models.ApplicationCommand) error {
	// if len(commands) < 1 {
	// 	return nil
	// }

	// var tx []db.PrismaTransaction
	// for _, cmd := range commands {
	// 	tx = append(tx, c.prisma.models.ApplicationCommand.UpsertOne(db.models.ApplicationCommand.ID.Equals(cmd.ID)).
	// 		Create(
	// 			db.models.ApplicationCommand.ID.Set(cmd.ID),
	// 			db.models.ApplicationCommand.Name.Set(cmd.Name),
	// 			db.models.ApplicationCommand.Version.Set(cmd.Version),
	// 			db.models.ApplicationCommand.OptionsHash.Set(cmd.Hash),
	// 		).
	// 		Update(
	// 			db.models.ApplicationCommand.Name.Set(cmd.Name),
	// 			db.models.ApplicationCommand.Version.Set(cmd.Version),
	// 			db.models.ApplicationCommand.OptionsHash.Set(cmd.Hash),
	// 		).Tx(),
	// 	)
	// }
	// return c.prisma.Prisma.Transaction(tx...).Exec(ctx)
	return nil
}
