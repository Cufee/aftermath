package database

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/crontask"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
)

func toCronTask(record *db.CronTask) models.Task {
	return models.Task{
		ID:             record.ID,
		Type:           record.Type,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
		ReferenceID:    record.ReferenceID,
		Targets:        record.Targets,
		Logs:           record.Logs,
		Status:         record.Status,
		ScheduledAfter: record.ScheduledAfter,
		LastRun:        record.LastRun,
		Data:           record.Data,
	}
}

/*
Returns up limit tasks that have TaskStatusInProgress and were last updates 1+ hours ago
*/
func (c *client) GetStaleTasks(ctx context.Context, limit int) ([]models.Task, error) {
	records, err := c.db.CronTask.Query().
		Where(
			crontask.StatusEQ(models.TaskStatusInProgress),
			crontask.LastRunLT(time.Now().Add(time.Hour*-1)),
		).
		Order(crontask.ByScheduledAfter(sql.OrderAsc())).
		Limit(limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, r := range records {
		tasks = append(tasks, toCronTask(r))
	}

	return tasks, nil
}

/*
Returns all tasks that were created after createdAfter, sorted by ScheduledAfter (DESC)
*/
func (c *client) GetRecentTasks(ctx context.Context, createdAfter time.Time, status ...models.TaskStatus) ([]models.Task, error) {
	var where []predicate.CronTask
	where = append(where, crontask.CreatedAtGT(createdAfter))
	if len(status) > 0 {
		where = append(where, crontask.StatusIn(status...))
	}

	records, err := c.db.CronTask.Query().Where(where...).Order(crontask.ByLastRun(sql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, r := range records {
		tasks = append(tasks, toCronTask(r))
	}

	return tasks, nil
}

/*
GetAndStartTasks retrieves up to limit number of tasks matching the referenceId and updates their status to in progress
*/
func (c *client) GetAndStartTasks(ctx context.Context, limit int) ([]models.Task, error) {
	if limit < 1 {
		return nil, nil
	}

	var tasks []models.Task
	return tasks, c.withTx(ctx, func(tx *db.Tx) error {
		records, err := tx.CronTask.Query().Where(
			crontask.TriesLeftGT(0),
			crontask.ScheduledAfterLT(time.Now()),
			crontask.StatusEQ(models.TaskStatusScheduled),
		).Order(crontask.ByScheduledAfter(sql.OrderAsc())).Limit(limit).All(ctx)
		if err != nil {
			return err
		}

		var ids []string
		for _, r := range records {
			t := toCronTask(r)
			t.OnUpdated()

			t.Status = models.TaskStatusInProgress
			tasks = append(tasks, t)
			ids = append(ids, r.ID)
		}

		return tx.CronTask.Update().Where(crontask.IDIn(ids...)).SetStatus(models.TaskStatusInProgress).Exec(ctx)
	})
}

/*
AbandonTasks retrieves tasks by ids and marks them as scheduled, adding a log noting it was abandoned
*/
func (c *client) AbandonTasks(ctx context.Context, ids ...string) error {
	if len(ids) < 1 {
		return nil
	}

	var tasks []models.Task
	err := c.withTx(ctx, func(tx *db.Tx) error {
		records, err := tx.CronTask.Query().Where(crontask.IDIn(ids...)).All(ctx)
		if err != nil {
			return err
		}

		now := time.Now()
		for _, r := range records {
			t := toCronTask(r)
			t.OnUpdated()

			t.Status = models.TaskStatusScheduled
			t.LogAttempt(models.TaskLog{
				Comment:   "task was abandoned",
				Timestamp: now,
			})
			tasks = append(tasks, t)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return c.UpdateTasks(ctx, tasks...)
}

func (c *client) CreateTasks(ctx context.Context, tasks ...models.Task) error {
	if len(tasks) < 1 {
		return nil
	}
	return c.withTx(ctx, func(tx *db.Tx) error {
		var inserts []*db.CronTaskCreate
		for _, t := range tasks {
			t.OnCreated()
			t.OnUpdated()

			inserts = append(inserts,
				tx.CronTask.Create().
					SetType(t.Type).
					SetData(t.Data).
					SetLogs(t.Logs).
					SetStatus(t.Status).
					SetTargets(t.Targets).
					SetLastRun(t.LastRun).
					SetTriesLeft(t.TriesLeft).
					SetReferenceID(t.ReferenceID).
					SetScheduledAfter(t.ScheduledAfter),
			)
		}
		return tx.CronTask.CreateBulk(inserts...).Exec(ctx)
	})
}

func (c *client) GetTasks(ctx context.Context, ids ...string) ([]models.Task, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	records, err := c.db.CronTask.Query().Where(crontask.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for _, r := range records {
		tasks = append(tasks, toCronTask(r))
	}

	return tasks, nil
}

/*
UpdateTasks will update all tasks passed in
  - the following fields will be replaced: targets, status, leastRun, scheduleAfterm logs, data
  - this func will block until all other calls to task update funcs are done
*/
func (c *client) UpdateTasks(ctx context.Context, tasks ...models.Task) error {
	if len(tasks) < 1 {
		return nil
	}

	return c.withTx(ctx, func(tx *db.Tx) error {
		for _, t := range tasks {
			t.OnUpdated()

			err := tx.CronTask.UpdateOneID(t.ID).
				SetData(t.Data).
				SetLogs(t.Logs).
				SetStatus(t.Status).
				SetTargets(t.Targets).
				SetLastRun(t.LastRun).
				SetScheduledAfter(t.ScheduledAfter).
				Exec(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

/*
DeleteTasks will delete all tasks matching by ids
  - this func will block until all other calls to task update funcs are done
*/
func (c *client) DeleteTasks(ctx context.Context, ids ...string) error {
	if len(ids) < 1 {
		return nil
	}

	return c.withTx(ctx, func(tx *db.Tx) error {
		_, err := tx.CronTask.Delete().Where(crontask.IDIn(ids...)).Exec(ctx)
		return err
	})

}
