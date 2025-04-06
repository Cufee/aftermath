package database

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	t "github.com/cufee/aftermath/internal/database/gen/public/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"golang.org/x/sync/errgroup"
)

func (c *client) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
	_, err := c.exec(ctx, t.CronTask.DELETE().WHERE(t.CronTask.CreatedAt.LT(timeToField(expiration))))
	return err
}

func (c *client) DeleteExpiredInteractions(ctx context.Context, expiration time.Time) error {
	_, err := c.exec(ctx, t.DiscordInteraction.DELETE().WHERE(t.DiscordInteraction.CreatedAt.LT(timeToField(expiration))))
	return err
}

func (c *client) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) (int, error) {
	var rowsAffected atomic.Int64

	var group errgroup.Group
	group.Go(func() error {
		sql, args := snapshotCleanup(t.AccountSnapshot.TableName(), t.AccountSnapshot.AccountID.Name(), expiration)
		result, err := c.db.ExecContext(ctx, sql, args...)
		if err != nil {
			return err
		}
		affected, _ := result.RowsAffected()
		rowsAffected.Add(affected)
		log.Debug().Int64("deleted", affected).Msg("account snapshot cleanup complete")
		return nil
	})
	group.Go(func() error {
		sql, args := snapshotCleanup(t.AccountAchievementsSnapshot.TableName(), t.AccountAchievementsSnapshot.AccountID.Name(), expiration)
		result, err := c.db.ExecContext(ctx, sql, args...)
		if err != nil {
			return err
		}
		affected, _ := result.RowsAffected()
		rowsAffected.Add(affected)
		log.Debug().Int64("deleted", affected).Msg("account achievements snapshot cleanup complete")
		return nil

	})
	group.Go(func() error {
		sql, args := snapshotCleanup(t.VehicleSnapshot.TableName(), t.VehicleSnapshot.VehicleID.Name(), expiration)
		result, err := c.db.ExecContext(ctx, sql, args...)
		if err != nil {
			return err
		}
		affected, _ := result.RowsAffected()
		rowsAffected.Add(affected)
		log.Debug().Int64("deleted", affected).Msg("vehicle snapshot cleanup complete")
		return nil
	})

	group.Go(func() error {
		sql, args := snapshotCleanup(t.VehicleAchievementsSnapshot.TableName(), t.VehicleAchievementsSnapshot.VehicleID.Name(), expiration)
		result, err := c.db.ExecContext(ctx, sql, args...)
		if err != nil {
			return err
		}
		affected, _ := result.RowsAffected()
		rowsAffected.Add(affected)
		log.Debug().Int64("deleted", affected).Msg("vehicle achievements snapshot cleanup complete")
		return nil
	})

	if err := group.Wait(); err != nil {
		return int(rowsAffected.Load()), err
	}
	return int(rowsAffected.Load()), nil
}

func snapshotCleanup(table string, idField string, expiration time.Time) (string, []any) {
	timeStr := models.TimeToString(expiration)
	return fmt.Sprintf(`
DELETE FROM %[1]s
WHERE id IN (
    SELECT s1.id
    FROM %[1]s AS s1
    WHERE s1.created_at < $1
    AND EXISTS (
        SELECT 1
        FROM %[1]s AS s2
        WHERE s1.%[2]s = s2.%[2]s
        AND s1.reference_id = s2.reference_id
        AND s1.type = s2.type
        AND s2.type = $3
        AND s2.created_at >= $2
    )
)`, table, idField), []any{timeStr, timeStr, models.SnapshotTypeDaily}
}
