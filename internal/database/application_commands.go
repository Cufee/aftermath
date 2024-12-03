package database

import (
	"context"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/sqlite"
)

func (c *client) GetCommandsByID(ctx context.Context, commandIDs ...string) ([]models.ApplicationCommand, error) {
	if len(commandIDs) < 1 {
		return nil, nil
	}

	stmt := t.ApplicationCommand.
		SELECT(t.ApplicationCommand.AllColumns).
		WHERE(t.ApplicationCommand.ID.IN(toStringSlice(commandIDs...)...))

	var result []m.ApplicationCommand
	err := c.query(ctx, stmt, &result)
	if err != nil {
		return nil, err
	}

	var commands []models.ApplicationCommand
	for _, r := range result {
		commands = append(commands, models.ToApplicationCommand(&r))
	}
	return commands, nil
}

func (c *client) GetCommandsByHash(ctx context.Context, commandHashes ...string) ([]models.ApplicationCommand, error) {
	if len(commandHashes) < 1 {
		return nil, nil
	}

	stmt := t.ApplicationCommand.
		SELECT(t.ApplicationCommand.AllColumns).
		WHERE(t.ApplicationCommand.OptionsHash.IN(toStringSlice(commandHashes...)...))

	var result []m.ApplicationCommand
	err := c.query(ctx, stmt, &result)
	if err != nil {
		return nil, err
	}

	var commands []models.ApplicationCommand
	for _, r := range result {
		commands = append(commands, models.ToApplicationCommand(&r))
	}
	return commands, nil
}

func (c *client) UpsertCommands(ctx context.Context, commands ...models.ApplicationCommand) error {
	if len(commands) < 1 {
		return nil
	}

	return c.withTx(ctx, func(tx *transaction) error {
		for _, command := range commands {
			stmt := t.ApplicationCommand.
				INSERT(t.ApplicationCommand.AllColumns).
				MODEL(command.Model()).
				ON_CONFLICT(t.ApplicationCommand.ID).
				DO_UPDATE(
					s.SET(
						t.ApplicationCommand.OptionsHash.SET(t.ApplicationCommand.EXCLUDED.OptionsHash),
						t.ApplicationCommand.UpdatedAt.SET(t.ApplicationCommand.EXCLUDED.UpdatedAt),
						t.ApplicationCommand.Version.SET(t.ApplicationCommand.EXCLUDED.Version),
						t.ApplicationCommand.Name.SET(t.ApplicationCommand.EXCLUDED.Name),
					),
				)
			_, err := tx.exec(ctx, stmt)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
