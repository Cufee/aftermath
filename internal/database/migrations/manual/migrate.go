package manual

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/gen/public/model"
	"github.com/cufee/aftermath/internal/database/gen/public/table"
	"github.com/cufee/aftermath/internal/database/models"
	pg "github.com/go-jet/jet/v2/postgres"
	"github.com/lucsky/cuid"
	"github.com/pkg/errors"
)

func Migrate(ctx context.Context, client database.Client) error {
	err := migration_08042025(ctx, client)
	if err != nil {
		return err
	}

	// err = migration_08042025_4(ctx, client)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func startMigration(ctx context.Context, client database.Client, key string) (bool, func(context.Context), error) {
	var data model.ManualMigration
	err := table.ManualMigration.SELECT(table.ManualMigration.AllColumns).WHERE(table.ManualMigration.Key.EQ(pg.String(key))).QueryContext(ctx, client.Unsafe(), &data)
	if err != nil && !database.IsNotFound(err) {
		return false, func(ctx context.Context) {}, err
	}
	if database.IsNotFound(err) {
		data = model.ManualMigration{
			ID:        cuid.New(),
			CreatedAt: models.TimeToString(time.Now()),
			UpdatedAt: models.TimeToString(time.Now()),
			Key:       key,
			Finished:  false,
			Metadata:  make([]byte, 0),
		}

		stmt := table.ManualMigration.
			INSERT(table.ManualMigration.AllColumns).
			MODEL(data).
			RETURNING(table.ManualMigration.AllColumns)
		err := stmt.QueryContext(ctx, client.Unsafe(), &data)
		if err != nil {
			return false, func(ctx context.Context) {}, err
		}
	}
	if data.Finished {
		return false, func(context.Context) {}, nil
	}
	return true, func(ctx context.Context) {
		err := finishMigration(ctx, client, data.ID)
		if err != nil {
			panic(err)
		}
	}, nil
}

func finishMigration(ctx context.Context, client database.Client, id string) error {
	stmt := table.ManualMigration.
		UPDATE(
			table.ManualMigration.Finished,
			table.ManualMigration.UpdatedAt,
		).
		WHERE(table.ManualMigration.ID.EQ(pg.String(id))).
		SET(
			table.ManualMigration.Finished.SET(pg.Bool(true)),
			table.ManualMigration.UpdatedAt.SET(pg.String(models.TimeToString(time.Now()))),
		)
	result, err := stmt.ExecContext(ctx, client.Unsafe())
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("bad number of rows affected")
	}
	return nil
}
