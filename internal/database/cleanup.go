package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/achievementssnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/crontask"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
)

func (c *client) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
	tx, cancel, err := c.txWithLock(ctx)
	if err != nil {
		return err
	}
	defer cancel()

	_, err = tx.CronTask.Delete().Where(crontask.CreatedAtLT(expiration.Unix())).Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

func (c *client) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
	_, err := c.db.AccountSnapshot.Delete().Where(accountsnapshot.CreatedAtLT(expiration.Unix())).Exec(ctx)
	if err != nil {
		return err
	}

	_, err = c.db.VehicleSnapshot.Delete().Where(vehiclesnapshot.CreatedAtLT(expiration.Unix())).Exec(ctx)
	if err != nil {
		return err
	}

	_, err = c.db.AchievementsSnapshot.Delete().Where(achievementssnapshot.CreatedAtLT(expiration.Unix())).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
