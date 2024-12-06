package database

import (
	"context"
	"fmt"
	"time"

	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
)

func (c *client) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
	_, err := c.exec(ctx, t.CronTask.DELETE().WHERE(t.CronTask.CreatedAt.LT(timeToField(expiration))))
	return err
}

func (c *client) DeleteExpiredInteractions(ctx context.Context, expiration time.Time) error {
	_, err := c.exec(ctx, t.DiscordInteraction.DELETE().WHERE(t.DiscordInteraction.CreatedAt.LT(timeToField(expiration))))
	return err
}

func (c *client) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
	{
		sql, args := snapshotCleanup(t.AccountSnapshot.TableName(), t.AccountSnapshot.AccountID.Name(), expiration)
		result, err := c.db.ExecContext(ctx, sql, args...)
		if err != nil {
			return err
		}
		affected, _ := result.RowsAffected()
		log.Debug().Int64("deleted", affected).Msg("account snapshot cleanup complete")
	}
	{
		sql, args := snapshotCleanup(t.VehicleSnapshot.TableName(), t.VehicleSnapshot.VehicleID.Name(), expiration)
		result, err := c.db.ExecContext(ctx, sql, args...)
		if err != nil {
			return err
		}
		affected, _ := result.RowsAffected()
		log.Debug().Int64("deleted", affected).Msg("vehicle snapshot cleanup complete")
	}
	return nil
}

func snapshotCleanup(table string, idField string, expiration time.Time) (string, []any) {
	time := models.TimeToString(expiration)
	return fmt.Sprintf(`
DELETE from %[1]s
WHERE id IN (
	SELECT id
	FROM %[1]s AS snapshots_1
	WHERE created_at < ?
	AND (
		type != '%[3]s'
		OR EXISTS (
			SELECT 1
			FROM %[1]s AS snapshots_2
			WHERE snapshots_1.%[2]s = snapshots_2.%[2]s
			AND snapshots_1.reference_id = snapshots_2.reference_id
			AND snapshots_2.created_at >= ?
		)
	)
);`, table, idField, models.SnapshotTypeDaily), []any{time, time}
}
