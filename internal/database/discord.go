package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/applicationcommand"
	"github.com/cufee/aftermath/internal/database/models"
)

func toApplicationCommand(record *db.ApplicationCommand) models.ApplicationCommand {
	return models.ApplicationCommand{
		ID:      record.ID,
		Name:    record.Name,
		Hash:    record.OptionsHash,
		Version: record.Version,
	}

}

func (c *client) GetCommandsByID(ctx context.Context, commandIDs ...string) ([]models.ApplicationCommand, error) {
	if len(commandIDs) < 1 {
		return nil, nil
	}

	records, err := c.db.ApplicationCommand.Query().Where(applicationcommand.IDIn(commandIDs...)).All(ctx)
	if err != nil {
		return nil, err
	}

	var commands []models.ApplicationCommand
	for _, c := range records {
		commands = append(commands, toApplicationCommand(c))
	}
	return commands, nil
}

func (c *client) GetCommandsByHash(ctx context.Context, commandHashes ...string) ([]models.ApplicationCommand, error) {
	if len(commandHashes) < 1 {
		return nil, nil
	}

	records, err := c.db.ApplicationCommand.Query().Where(applicationcommand.OptionsHashIn(commandHashes...)).All(ctx)
	if err != nil {
		return nil, err
	}
	var commands []models.ApplicationCommand
	for _, c := range records {
		commands = append(commands, toApplicationCommand(c))
	}
	return commands, nil
}

func (c *client) UpsertCommands(ctx context.Context, commands ...models.ApplicationCommand) error {
	if len(commands) < 1 {
		return nil
	}

	var ids []string
	commandsMap := make(map[string]*models.ApplicationCommand)
	for _, c := range commands {
		ids = append(ids, c.ID)
		commandsMap[c.ID] = &c
	}

	tx, cancel, err := c.txWithLock(ctx)
	if err != nil {
		return err
	}
	defer cancel()

	existing, err := tx.ApplicationCommand.Query().Where(applicationcommand.IDIn(ids...)).All(ctx)
	if err != nil && !IsNotFound(err) {
		return rollback(tx, err)
	}

	for _, c := range existing {
		update, ok := commandsMap[c.ID]
		if !ok {
			continue
		}

		err := tx.ApplicationCommand.UpdateOneID(c.ID).
			SetName(update.Name).
			SetVersion(update.Version).
			SetOptionsHash(update.Hash).
			Exec(ctx)
		if err != nil {
			return rollback(tx, err)
		}

		delete(commandsMap, c.ID)
	}

	var inserts []*db.ApplicationCommandCreate
	for _, cmd := range commandsMap {
		inserts = append(inserts,
			c.db.ApplicationCommand.Create().
				SetID(cmd.ID).
				SetName(cmd.Name).
				SetVersion(cmd.Version).
				SetOptionsHash(cmd.Hash),
		)
	}

	err = tx.ApplicationCommand.CreateBulk(inserts...).Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}
