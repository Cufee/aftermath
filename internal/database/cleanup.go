package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/achievementssnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/crontask"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
)

func (c *libsqlClient) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
	tx, err := c.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.CronTask.Delete().Where(crontask.CreatedAtLT(expiration.Unix())).Exec(ctx)
	if err != nil && !IsNotFound(err) {
		return rollback(tx, err)
	}
	return tx.Commit()
}

func (c *libsqlClient) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
	_, err := c.db.AccountSnapshot.Delete().Where(accountsnapshot.CreatedAtLT(expiration.Unix())).Exec(ctx)
	if err != nil && !IsNotFound(err) {
		return err
	}

	_, err = c.db.VehicleSnapshot.Delete().Where(vehiclesnapshot.CreatedAtLT(expiration.Unix())).Exec(ctx)
	if err != nil && !IsNotFound(err) {
		return err
	}

	_, err = c.db.AchievementsSnapshot.Delete().Where(achievementssnapshot.CreatedAtLT(expiration.Unix())).Exec(ctx)
	if err != nil && !IsNotFound(err) {
		return err
	}

	return nil
}
