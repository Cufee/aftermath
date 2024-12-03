package database

import (
	"context"
	"time"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/sqlite"
)

/*
Returns up limit tasks that have TaskStatusInProgress and were last updates 1+ hours ago
*/
func (c *client) GetStaleTasks(ctx context.Context, limit int) ([]models.Task, error) {
	var records []m.CronTask
	stmt := t.CronTask.
		SELECT(t.CronTask.AllColumns).
		WHERE(s.AND(
			t.CronTask.Status.EQ(s.String(string(models.TaskStatusInProgress))),
			t.CronTask.LastRun.LT(s.DATETIME(time.Now().Add(time.Hour*-1))),
		)).
		LIMIT(int64(limit))

	err := c.query(ctx, stmt, &records)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, r := range records {
		tasks = append(tasks, models.ToCronTask(&r))
	}

	return tasks, nil
}

/*
GetRecentTasks retrieves tasks sorted by DESC(task.last_run)
  - createdAfter - tasks.created_at gt
  - status (optional) - status of the tasks
*/
func (c *client) GetRecentTasks(ctx context.Context, createdAfter time.Time, status ...models.TaskStatus) ([]models.Task, error) {
	where := []s.BoolExpression{t.CronTask.CreatedAt.GT(s.DATETIME(createdAfter))}
	if len(status) > 0 {
		var s []string
		for _, st := range status {
			s = append(s, string(st))
		}
		where = append(where, t.CronTask.Status.IN(stringsToExp(s)...))
	}

	var records []m.CronTask
	stmt := t.CronTask.
		SELECT(t.CronTask.AllColumns).
		WHERE(s.AND(where...)).
		ORDER_BY(t.CronTask.LastRun.DESC())

	err := c.query(ctx, stmt, &records)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, r := range records {
		tasks = append(tasks, models.ToCronTask(&r))
	}

	return tasks, nil
}

/*
GetAndStartTasks retrieves up to limit number of tasks and updates their status to in progress
*/
func (c *client) GetAndStartTasks(ctx context.Context, limit int) ([]models.Task, error) {
	if limit < 1 {
		return nil, nil
	}

	var tasks []models.Task
	return tasks, c.withTx(ctx, func(tx *transaction) error {
		stmt := t.CronTask.
			SELECT(t.CronTask.AllColumns).
			WHERE(s.AND(
				t.CronTask.TriesLeft.GT(s.Int32(0)),
				t.CronTask.ScheduledAfter.LT(s.DATETIME(time.Now())),
				t.CronTask.Status.EQ(s.String(string(models.TaskStatusScheduled))),
			)).
			ORDER_BY(t.CronTask.ScheduledAfter.ASC()).
			LIMIT(int64(limit))

		var selected []m.CronTask
		err := tx.query(ctx, stmt, &selected)
		if err != nil {
			return err
		}

		for _, r := range selected {
			task := models.ToCronTask(&r)
			task.OnUpdated()
			task.Status = models.TaskStatusInProgress

			stmt := t.CronTask.
				UPDATE(
					t.CronTask.UpdatedAt,
					t.CronTask.LastRun,
					t.CronTask.Status,
				).
				SET(s.SET(
					t.CronTask.Status.SET(s.String(string(task.Status))),
					t.CronTask.UpdatedAt.SET(s.DATETIME(task.UpdatedAt)),
					t.CronTask.LastRun.SET(s.DATETIME(task.LastRun)),
				)).
				WHERE(t.CronTask.ID.EQ(s.String(task.ID)))

			_, err := tx.exec(ctx, stmt)
			if err != nil {
				return err
			}

			tasks = append(tasks, task)
		}
		return nil
	})
}

/*
AbandonTasks retrieves tasks by ids and marks them as scheduled, adding a log entry
*/
func (c *client) AbandonTasks(ctx context.Context, ids ...string) error {
	if len(ids) < 1 {
		return nil
	}

	return c.withTx(ctx, func(tx *transaction) error {
		stmt := t.CronTask.SELECT(t.CronTask.AllColumns).WHERE(t.CronTask.ID.IN(stringsToExp(ids)...))

		var records []m.CronTask
		err := tx.query(ctx, stmt, &records)
		if err != nil {
			return err
		}

		now := time.Now()
		for _, r := range records {
			task := models.ToCronTask(&r)
			task.OnUpdated()

			task.Status = models.TaskStatusScheduled
			task.LogAttempt(models.TaskLog{
				Comment:   "task was abandoned",
				Timestamp: now,
			})

			stmt := t.CronTask.
				UPDATE(
					t.CronTask.Logs,
					t.CronTask.Status,
					t.CronTask.LastRun,
					t.CronTask.UpdatedAt,
				).
				MODEL(task.Model()).
				WHERE(t.CronTask.ID.EQ(s.String(task.ID)))

			_, err := tx.exec(ctx, stmt)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (c *client) CreateTasks(ctx context.Context, tasks ...models.Task) error {
	if len(tasks) < 1 {
		return nil
	}
	return c.withTx(ctx, func(tx *transaction) error {
		var inserts []m.CronTask
		for _, task := range tasks {
			task.OnCreated()
			task.OnUpdated()
			inserts = append(inserts, task.Model())
		}

		stmt := t.CronTask.INSERT().MODELS(inserts)
		_, err := tx.exec(ctx, stmt)
		return err
	})
}

func (c *client) GetTasks(ctx context.Context, ids ...string) ([]models.Task, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	var records []m.CronTask
	stmt := t.CronTask.SELECT(t.CronTask.AllColumns).WHERE(t.CronTask.ID.IN(stringsToExp(ids)...))

	err := c.query(ctx, stmt, &records)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, r := range records {
		tasks = append(tasks, models.ToCronTask(&r))
	}

	return tasks, nil
}

/*
UpdateTasks will update all tasks passed in
*/
func (c *client) UpdateTasks(ctx context.Context, tasks ...models.Task) error {
	if len(tasks) < 1 {
		return nil
	}

	return c.withTx(ctx, func(tx *transaction) error {
		for _, task := range tasks {
			task.OnUpdated()
			stmt := t.CronTask.
				UPDATE(
					t.CronTask.UpdatedAt,
					t.CronTask.Targets,
					t.CronTask.Logs,
					t.CronTask.Status,
					t.CronTask.ScheduledAfter,
					t.CronTask.TriesLeft,
					t.CronTask.LastRun,
					t.CronTask.Data,
				).
				MODEL(task.Model()).
				WHERE(t.CronTask.ID.EQ(s.String(task.ID)))

			_, err := tx.exec(ctx, stmt)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

/*
DeleteTasks will delete all tasks matching by ids
*/
func (c *client) DeleteTasks(ctx context.Context, ids ...string) error {
	if len(ids) < 1 {
		return nil
	}
	return c.withTx(ctx, func(tx *transaction) error {
		_, err := tx.exec(ctx, t.CronTask.DELETE().WHERE(t.CronTask.ID.IN(stringsToExp(ids)...)))
		return err
	})
}
