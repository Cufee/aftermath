package tasks

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/cmds/core"
	"github.com/cufee/aftermath/internal/database"
)

func init() {
	defaultHandlers[database.TaskTypeDatabaseCleanup] = TaskHandler{
		Process: func(client core.Client, task database.Task) (string, error) {
			if task.Data == nil {
				return "no data provided", errors.New("no data provided")
			}
			snapshotExpiration, ok := task.Data["expiration_snapshots"].(int64)
			if !ok {
				return "invalid expiration_snapshots", errors.New("failed to cast expiration_snapshots to time")
			}
			taskExpiration, ok := task.Data["expiration_tasks"].(int64)
			if !ok {
				task.Data["triesLeft"] = int(0) // do not retry
				return "invalid expiration_tasks", errors.New("failed to cast expiration_tasks to time")
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			err := client.Database().DeleteExpiredTasks(ctx, time.Unix(taskExpiration, 0))
			if err != nil {
				return "failed to delete expired tasks", err
			}

			err = client.Database().DeleteExpiredSnapshots(ctx, time.Unix(snapshotExpiration, 0))
			if err != nil {
				return "failed to delete expired snapshots", err
			}

			return "cleanup complete", nil
		},
		ShouldRetry: func(task *database.Task) bool {
			return false
		},
	}
}

func CreateCleanupTasks(client core.Client) error {
	now := time.Now()

	task := database.Task{
		Type:           database.TaskTypeDatabaseCleanup,
		ReferenceID:    "database_cleanup",
		ScheduledAfter: now,
		Data: map[string]any{
			"expiration_snapshots": now.Add(-1 * time.Hour * 24 * 90).Unix(), // 90 days
			"expiration_tasks":     now.Add(-1 * time.Hour * 24 * 7).Unix(),  // 7 days
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return client.Database().CreateTasks(ctx, task)
}
