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
		CreatedAt:      time.Unix(record.CreatedAt, 0),
		UpdatedAt:      time.Unix(record.UpdatedAt, 0),
		ReferenceID:    record.ReferenceID,
		Targets:        record.Targets,
		Logs:           record.Logs,
		Status:         record.Status,
		ScheduledAfter: time.Unix(record.ScheduledAfter, 0),
		LastRun:        time.Unix(record.LastRun, 0),
		Data:           record.Data,
	}
}

/*
Returns up limit tasks that have TaskStatusInProgress and were last updates 1+ hours ago
*/
func (c *libsqlClient) GetStaleTasks(ctx context.Context, limit int) ([]models.Task, error) {
	records, err := c.db.CronTask.Query().
		Where(
			crontask.StatusEQ(models.TaskStatusInProgress),
			crontask.LastRunLT(time.Now().Add(time.Hour*-1).Unix()),
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
func (c *libsqlClient) GetRecentTasks(ctx context.Context, createdAfter time.Time, status ...models.TaskStatus) ([]models.Task, error) {
	var where []predicate.CronTask
	where = append(where, crontask.CreatedAtGT(createdAfter.Unix()))
	if len(status) > 0 {
		where = append(where, crontask.StatusIn(status...))
	}

	records, err := c.db.CronTask.Query().Where(where...).All(ctx)
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
func (c *libsqlClient) GetAndStartTasks(ctx context.Context, limit int) ([]models.Task, error) {
	if limit < 1 {
		return nil, nil
	}

	tx, err := c.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	records, err := tx.CronTask.Query().Where(crontask.StatusEQ(models.TaskStatusScheduled), crontask.ScheduledAfterLT(time.Now().Unix())).Order(crontask.ByScheduledAfter(sql.OrderAsc())).Limit(limit).All(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	var ids []string
	var tasks []models.Task
	for _, r := range records {
		t := toCronTask(r)
		t.OnUpdated()

		t.Status = models.TaskStatusInProgress
		tasks = append(tasks, t)
		ids = append(ids, r.ID)
	}

	err = tx.CronTask.Update().Where(crontask.IDIn(ids...)).SetStatus(models.TaskStatusInProgress).Exec(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return tasks, tx.Commit()
}

func (c *libsqlClient) CreateTasks(ctx context.Context, tasks ...models.Task) error {
	if len(tasks) < 1 {
		return nil
	}

	var inserts []*db.CronTaskCreate
	for _, t := range tasks {
		t.OnCreated()
		t.OnUpdated()

		inserts = append(inserts,
			c.db.CronTask.Create().
				SetType(t.Type).
				SetData(t.Data).
				SetLogs(t.Logs).
				SetStatus(t.Status).
				SetTargets(t.Targets).
				SetLastRun(t.LastRun.Unix()).
				SetReferenceID(t.ReferenceID).
				SetScheduledAfter(t.ScheduledAfter.Unix()),
		)
	}

	return c.db.CronTask.CreateBulk(inserts...).Exec(ctx)
}

/*
UpdateTasks will update all tasks passed in
  - the following fields will be replaced: targets, status, leastRun, scheduleAfterm logs, data
  - this func will block until all other calls to task update funcs are done
*/
func (c *libsqlClient) UpdateTasks(ctx context.Context, tasks ...models.Task) error {
	if len(tasks) < 1 {
		return nil
	}

	tx, err := c.db.Tx(ctx)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		t.OnUpdated()

		err := tx.CronTask.UpdateOneID(t.ID).
			SetData(t.Data).
			SetLogs(t.Logs).
			SetStatus(t.Status).
			SetTargets(t.Targets).
			SetLastRun(t.LastRun.Unix()).
			SetScheduledAfter(t.ScheduledAfter.Unix()).
			Exec(ctx)

		if err != nil {
			return rollback(tx, err)
		}
	}

	return tx.Commit()
}

/*
DeleteTasks will delete all tasks matching by ids
  - this func will block until all other calls to task update funcs are done
*/
func (c *libsqlClient) DeleteTasks(ctx context.Context, ids ...string) error {
	if len(ids) < 1 {
		return nil
	}

	tx, err := c.db.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.CronTask.Delete().Where(crontask.IDIn(ids...)).Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}
