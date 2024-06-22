package database

import (
	"context"
	"time"
)

func (c *libsqlClient) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
	// _, err := c.prisma.CronTask.FindMany(db.CronTask.CreatedAt.Before(expiration)).Delete().Exec(ctx)
	// if err != nil && !database.IsNotFound(err) {
	// 	return err
	// }
	return nil
}

func (c *libsqlClient) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
	// _, err := c.prisma.AccountSnapshot.FindMany(db.AccountSnapshot.CreatedAt.Before(expiration)).Delete().Exec(ctx)
	// if err != nil && !database.IsNotFound(err) {
	// 	return err
	// }
	// _, err = c.prisma.VehicleSnapshot.FindMany(db.VehicleSnapshot.CreatedAt.Before(expiration)).Delete().Exec(ctx)
	// if err != nil && !database.IsNotFound(err) {
	// 	return err
	// }
	return nil
}
