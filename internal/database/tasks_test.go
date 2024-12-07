package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
)

func TestCronTasks(t *testing.T) {
	is := is.New(t)

	client := MustTestClient(t)
	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.CronTask.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.CronTask.TableName()))

	task1 := models.Task{
		ID:             "task-1",
		Type:           models.TaskTypeRecordSnapshots,
		ScheduledAfter: time.Now(),
		ReferenceID:    "ref-1",
		Targets:        []string{"t1"},
		TriesLeft:      1,
	}
	task2 := models.Task{
		ID:             "task-2",
		Type:           models.TaskTypeRecordSnapshots,
		ReferenceID:    "ref-2",
		ScheduledAfter: time.Now(),
		Data: map[string]string{
			"key-1": "value-1",
		},
		TriesLeft: 3,
	}

	err := client.CreateTasks(context.Background(), task1, task2)
	is.NoErr(err)

	// GetRecentTasks
	{
		tasks, err := client.GetRecentTasks(context.Background(), time.Now().Add(time.Hour*-1))
		is.NoErr(err)
		is.True(len(tasks) == 2)
	}

	// UpdateTasks
	{
		task1.Status = models.TaskStatusComplete
		err := client.UpdateTasks(context.Background(), task1)
		is.NoErr(err)

		updated, err := client.GetTaskByID(context.Background(), task1.ID)
		is.NoErr(err)
		is.True(updated.ID == task1.ID)
		is.True(updated.Status == task1.Status)
	}

	{
		// Start/Abandon tasks
		started, err := client.GetAndStartTasks(context.Background(), 1)
		is.NoErr(err)
		is.True(len(started) == 1)
		is.True(started[0].Status == models.TaskStatusInProgress)

		err = client.AbandonTasks(context.Background(), started[0].ID)
		is.NoErr(err)

		abandoned, err := client.GetTaskByID(context.Background(), started[0].ID)
		is.NoErr(err)
		is.True(abandoned.ID == started[0].ID)
		is.True(abandoned.Status == models.TaskStatusScheduled)
	}
}
