package tasks

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/cmd/core"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
)

func init() {
	defaultHandlers[models.TaskTypeDatabaseCleanup] = TaskHandler{
		Process: func(ctx context.Context, client core.Client, task *models.Task) error {
			if task.Data == nil {
				task.Data["triesLeft"] = "0" // do not retry
				return errors.New("no data provided")
			}
			{
				taskExpiration, err := time.Parse(time.RFC3339, task.Data["expiration_tasks"])
				if err != nil {
					task.Data["triesLeft"] = "0" // do not retry
					return errors.Wrap(err, "failed to parse expiration_tasks to time")
				}
				err = client.Database().DeleteExpiredTasks(ctx, taskExpiration)
				if err != nil && !database.IsNotFound(err) {
					return errors.Wrap(err, "failed to delete expired tasks")
				}
			}
			{
				interactionExpiration, err := time.Parse(time.RFC3339, task.Data["expiration_interactions"])
				if err != nil {
					task.Data["triesLeft"] = "0" // do not retry
					return errors.Wrap(err, "failed to parse interactionExpiration to time")
				}
				err = client.Database().DeleteExpiredInteractions(ctx, interactionExpiration)
				if err != nil && !database.IsNotFound(err) {
					return errors.Wrap(err, "failed to delete expired interactions")
				}
			}
			{
				snapshotExpiration, err := time.Parse(time.RFC3339, task.Data["expiration_snapshots"])
				if err != nil {
					task.Data["triesLeft"] = "0" // do not retry
					return errors.Wrap(err, "failed to parse expiration_snapshots to time")
				}
				err = client.Database().DeleteExpiredSnapshots(ctx, snapshotExpiration)
				if err != nil && !database.IsNotFound(err) {
					return errors.Wrap(err, "failed to delete expired snapshots")
				}
			}

			return nil
		},
		ShouldRetry: func(task *models.Task) bool {
			return false
		},
	}
}

func CreateCleanupTasks(client core.Client) error {
	now := time.Now()

	task := models.Task{
		Type:           models.TaskTypeDatabaseCleanup,
		ReferenceID:    "database_cleanup",
		ScheduledAfter: now,
		Data: map[string]string{
			"expiration_snapshots":    now.Add(-1 * time.Hour * 24 * 90).Format(time.RFC3339), // 90 days
			"expiration_interactions": now.Add(-1 * time.Hour * 24 * 7).Format(time.RFC3339),  // 7 days
			"expiration_tasks":        now.Add(-1 * time.Hour * 24 * 7).Format(time.RFC3339),  // 7 days
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return client.Database().CreateTasks(ctx, task)
}
