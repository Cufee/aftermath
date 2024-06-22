package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
)

// func (t Task) FromModel(model db.CronTaskModel) Task {
// 	t.ID = model.ID
// 	t.Type = TaskType(model.Type)

// 	t.Status = TaskStatus(model.Status)
// 	t.ReferenceID = model.ReferenceID

// 	t.LastRun = model.LastRun
// 	t.CreatedAt = model.CreatedAt
// 	t.UpdatedAt = model.UpdatedAt
// 	t.ScheduledAfter = model.ScheduledAfter

// 	t.decodeData(model.DataEncoded)
// 	t.decodeLogs(model.LogsEncoded)
// 	t.decodeTargets(model.TargetsEncoded)
// 	return t
// }

/*
Returns up limit tasks that have TaskStatusInProgress and were last updates 1+ hours ago
*/
func (c *libsqlClient) GetStaleTasks(ctx context.Context, limit int) ([]models.Task, error) {
	// models, err := c.prisma.CronTask.FindMany(
	// 	db.CronTask.Status.Equals(string(TaskStatusInProgress)),
	// 	db.CronTask.LastRun.Before(time.Now().Add(time.Hour*-1)),
	// ).OrderBy(db.CronTask.ScheduledAfter.Order(db.ASC)).Take(limit).Exec(ctx)
	// if err != nil && !database.IsNotFound(err) {
	// 	return nil, err
	// }

	var tasks []models.Task
	// for _, model := range models {
	// 	tasks = append(tasks, Task{}.FromModel(model))
	// }

	return tasks, nil
}

/*
Returns all tasks that were created after createdAfter, sorted by ScheduledAfter (DESC)
*/
func (c *libsqlClient) GetRecentTasks(ctx context.Context, createdAfter time.Time, status ...models.TaskStatus) ([]models.Task, error) {
	// var statusStr []string
	// for _, s := range status {
	// 	statusStr = append(statusStr, string(s))
	// }

	// params := []db.CronTaskWhereParam{db.CronTask.CreatedAt.After(createdAfter)}
	// if len(statusStr) > 0 {
	// 	params = append(params, db.CronTask.Status.In(statusStr))
	// }

	// models, err := c.prisma.CronTask.FindMany(params...).OrderBy(db.CronTask.ScheduledAfter.Order(db.DESC)).Exec(ctx)
	// if err != nil && !database.IsNotFound(err) {
	// 	return nil, err
	// }

	var tasks []models.Task
	// for _, model := range models {
	// 	tasks = append(tasks, Task{}.FromModel(model))
	// }

	return tasks, nil
}

/*
GetAndStartTasks retrieves up to limit number of tasks matching the referenceId and updates their status to in progress
  - this func will block until all other calls to task update funcs are done
*/
func (c *libsqlClient) GetAndStartTasks(ctx context.Context, limit int) ([]models.Task, error) {
	// if limit < 1 {
	// 	return nil, nil
	// }
	// if err := c.tasksUpdateSem.Acquire(ctx, 1); err != nil {
	// 	return nil, err
	// }
	// defer c.tasksUpdateSem.Release(1)

	// models, err := c.prisma.CronTask.
	// 	FindMany(
	// 		db.CronTask.Status.Equals(string(TaskStatusScheduled)),
	// 		db.CronTask.ScheduledAfter.Lt(time.Now()),
	// 	).
	// 	OrderBy(db.CronTask.ScheduledAfter.Order(db.ASC)).
	// 	Take(limit).
	// 	Exec(ctx)
	// if err != nil {
	// 	if database.IsNotFound(err) {
	// 		return nil, nil
	// 	}
	// 	return nil, err
	// }
	// if len(models) < 1 {
	// 	return nil, nil
	// }

	// var ids []string
	var tasks []models.Task
	// for _, model := range models {
	// 	tasks = append(tasks, Task{}.FromModel(model))
	// 	ids = append(ids, model.ID)
	// }

	// _, err = c.prisma.CronTask.
	// 	FindMany(db.CronTask.ID.In(ids)).
	// 	Update(db.CronTask.Status.Set(string(TaskStatusInProgress)), db.CronTask.LastRun.Set(time.Now())).
	// 	Exec(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	return tasks, nil
}

func (c *libsqlClient) CreateTasks(ctx context.Context, tasks ...models.Task) error {
	// if len(tasks) < 1 {
	// 	return nil
	// }
	// // we do not block using c.tasksUpdateSem here because the order of ops in GetAndStartTasks makes this safe

	// var txns []db.PrismaTransaction
	// for _, task := range tasks {
	// 	task.OnCreated()
	// 	txns = append(txns, c.prisma.CronTask.CreateOne(
	// 		db.CronTask.Type.Set(string(task.Type)),
	// 		db.CronTask.ReferenceID.Set(task.ReferenceID),
	// 		db.CronTask.TargetsEncoded.Set(task.encodeTargets()),
	// 		db.CronTask.Status.Set(string(TaskStatusScheduled)),
	// 		db.CronTask.LastRun.Set(time.Now()),
	// 		db.CronTask.ScheduledAfter.Set(task.ScheduledAfter),
	// 		db.CronTask.LogsEncoded.Set(task.encodeLogs()),
	// 		db.CronTask.DataEncoded.Set(task.encodeData()),
	// 	).Tx())
	// }

	// err := c.prisma.Prisma.Transaction(txns...).Exec(ctx)
	// if err != nil {
	// 	return err
	// }
	return nil
}

/*
UpdateTasks will update all tasks passed in
  - the following fields will be replaced: targets, status, leastRun, scheduleAfterm logs, data
  - this func will block until all other calls to task update funcs are done
*/
func (c *libsqlClient) UpdateTasks(ctx context.Context, tasks ...models.Task) error {
	// if len(tasks) < 1 {
	// 	return nil
	// }

	// if err := c.tasksUpdateSem.Acquire(ctx, 1); err != nil {
	// 	return err
	// }
	// defer c.tasksUpdateSem.Release(1)

	// var txns []db.PrismaTransaction
	// for _, task := range tasks {
	// 	task.OnUpdated()
	// 	txns = append(txns, c.prisma.CronTask.
	// 		FindUnique(db.CronTask.ID.Equals(task.ID)).
	// 		Update(
	// 			db.CronTask.TargetsEncoded.Set(task.encodeTargets()),
	// 			db.CronTask.Status.Set(string(task.Status)),
	// 			db.CronTask.LastRun.Set(task.LastRun),
	// 			db.CronTask.ScheduledAfter.Set(task.ScheduledAfter),
	// 			db.CronTask.LogsEncoded.Set(task.encodeLogs()),
	// 			db.CronTask.DataEncoded.Set(task.encodeData()),
	// 		).Tx())
	// }

	// err := c.prisma.Prisma.Transaction(txns...).Exec(ctx)
	// if err != nil {
	// 	return err
	// }
	return nil
}

/*
DeleteTasks will delete all tasks matching by ids
  - this func will block until all other calls to task update funcs are done
*/
func (c *libsqlClient) DeleteTasks(ctx context.Context, ids ...string) error {
	// if len(ids) < 1 {
	// 	return nil
	// }

	// if err := c.tasksUpdateSem.Acquire(ctx, 1); err != nil {
	// 	return nil
	// }
	// defer c.tasksUpdateSem.Release(1)

	// _, err := c.prisma.CronTask.FindMany(db.CronTask.ID.In(ids)).Delete().Exec(ctx)
	// if err != nil {
	// 	return err
	// }
	return nil
}
