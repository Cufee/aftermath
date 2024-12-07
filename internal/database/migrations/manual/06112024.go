package manual

import (
	"context"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/go-jet/jet/v2/sqlite"
)

func ManualMigration_06112024(ctx context.Context, client database.Client) error {
	shouldRun, cleanup, err := startMigration(ctx, client, "ManualMigration_06112024")
	if err != nil {
		return err
	}
	if !shouldRun {
		log.Debug().Str("key", "ManualMigration_06112024").Msg("migration already complete")
		return nil
	}
	log.Debug().Str("key", "ManualMigration_06112024").Msg("running migration")
	defer cleanup(ctx)

	// Update user connections
	{
		var allConnections []model.UserConnection
		err := table.UserConnection.SELECT(table.UserConnection.AllColumns).QueryContext(ctx, client.Unsafe(), &allConnections)
		if database.IsNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}

		var setAsDefault []sqlite.Expression
		for _, record := range allConnections {
			conn := models.ToUserConnection(&record)
			if conn.Metadata["default"] == true {
				setAsDefault = append(setAsDefault, sqlite.String(conn.ID))
			}
		}

		_, err = table.UserConnection.
			UPDATE(table.UserConnection.Selected).
			WHERE(table.UserConnection.ID.IN(setAsDefault...)).
			SET(table.UserConnection.Selected.SET(sqlite.Bool(true))).
			ExecContext(ctx, client.Unsafe())

		if database.IsNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}
	}
	// Remove widget session references
	{
		m := model.WidgetSettings{
			SessionFrom:        nil,
			SessionReferenceID: nil,
		}
		_, err := table.WidgetSettings.
			UPDATE(
				table.WidgetSettings.SessionFrom,
				table.WidgetSettings.SessionReferenceID,
			).
			MODEL(m).
			WHERE(sqlite.Bool(true).EQ(sqlite.Bool(true))).
			ExecContext(ctx, client.Unsafe())

		if database.IsNotFound(err) {
			return nil
		}
		if err != nil {
			return err
		}
	}
	return nil
}
