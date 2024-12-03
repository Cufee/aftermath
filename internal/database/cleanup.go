package database

import (
	"context"
	"time"

	t "github.com/cufee/aftermath/internal/database/gen/table"
	s "github.com/go-jet/jet/v2/sqlite"
)

func (c *client) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
	_, err := c.exec(ctx, t.CronTasks.DELETE().WHERE(t.CronTasks.CreatedAt.LT(s.DATETIME(expiration))))
	return err
}

func (c *client) DeleteExpiredInteractions(ctx context.Context, expiration time.Time) error {
	_, err := c.exec(ctx, t.DiscordInteraction.DELETE().WHERE(t.DiscordInteraction.CreatedAt.LT(s.DATETIME(expiration))))
	return err
}

func (c *client) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
	_, err := c.exec(ctx, t.AccountSnapshot.DELETE().WHERE(t.AccountSnapshot.CreatedAt.LT(s.DATETIME(expiration))))
	if err != nil {
		return err
	}

	_, err = c.exec(ctx, t.VehicleSnapshot.DELETE().WHERE(t.VehicleSnapshot.CreatedAt.LT(s.DATETIME(expiration))))
	return err
}
