package database

// import (
// 	"context"
// 	"time"

// 	m "github.com/cufee/aftermath/internal/database/gen/model"
// 	t "github.com/cufee/aftermath/internal/database/gen/table"
// 	"github.com/cufee/aftermath/internal/database/models"
// 	s "github.com/go-jet/jet/v2/sqlite"
// )

// func (c *client) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
// 	err := c.withTx(ctx, func(tx *db.Tx) error {
// 		_, err := tx.CronTask.Delete().Where(crontask.CreatedAtLT(expiration)).Exec(ctx)
// 		return err
// 	})
// 	return err
// }

// func (c *client) DeleteExpiredInteractions(ctx context.Context, expiration time.Time) error {
// 	err := c.withTx(ctx, func(tx *db.Tx) error {
// 		_, err := tx.DiscordInteraction.Delete().Where(discordinteraction.CreatedAtLT(expiration)).Exec(ctx)
// 		return err
// 	})
// 	return err
// }

// func (c *client) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
// 	_, err := c.db.AccountSnapshot.Delete().Where(accountsnapshot.CreatedAtLT(expiration)).Exec(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = c.db.VehicleSnapshot.Delete().Where(vehiclesnapshot.CreatedAtLT(expiration)).Exec(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
