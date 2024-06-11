package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
)

func (c *client) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
	_, err := c.prisma.CronTask.FindMany(db.CronTask.CreatedAt.Before(expiration)).Delete().Exec(ctx)
	if err != nil && !db.IsErrNotFound(err) {
		return err
	}
	return nil
}

func (c *client) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
	_, err := c.prisma.AccountSnapshot.FindMany(db.AccountSnapshot.CreatedAt.Before(expiration)).Delete().Exec(ctx)
	if err != nil && !db.IsErrNotFound(err) {
		return err
	}
	_, err = c.prisma.VehicleSnapshot.FindMany(db.VehicleSnapshot.CreatedAt.Before(expiration)).Delete().Exec(ctx)
	if err != nil && !db.IsErrNotFound(err) {
		return err
	}
	return nil
}
