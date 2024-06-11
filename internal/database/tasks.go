package database

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/encoding"
)

type TaskType string

const (
	TaskTypeUpdateClans              TaskType = "UPDATE_CLANS"
	TaskTypeRecordSessions           TaskType = "RECORD_ACCOUNT_SESSIONS"
	TaskTypeUpdateAccountWN8         TaskType = "UPDATE_ACCOUNT_WN8"
	TaskTypeRecordPlayerAchievements TaskType = "UPDATE_ACCOUNT_ACHIEVEMENTS"

	TaskTypeDatabaseCleanup TaskType = "CLEANUP_DATABASE"
)

// Task statuses
type taskStatus string

const (
	TaskStatusScheduled  taskStatus = "TASK_SCHEDULED"
	TaskStatusInProgress taskStatus = "TASK_IN_PROGRESS"
	TaskStatusComplete   taskStatus = "TASK_COMPLETE"
	TaskStatusFailed     taskStatus = "TASK_FAILED"
)

type Task struct {
	ID        string    `json:"id"`
	Type      TaskType  `json:"kind"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	ReferenceID string   `json:"referenceId"`
	Targets     []string `json:"targets"`

	Logs []TaskLog `json:"logs"`

	Status         taskStatus `json:"status"`
	ScheduledAfter time.Time  `json:"scheduledAfter"`
	LastRun        time.Time  `json:"lastRun"`

	Data map[string]any `json:"data"`
}

func (t Task) FromModel(model db.CronTaskModel) Task {
	t.ID = model.ID
	t.Type = TaskType(model.Type)

	t.Status = taskStatus(model.Status)
	t.ReferenceID = model.ReferenceID

	t.LastRun = model.LastRun
	t.CreatedAt = model.CreatedAt
	t.UpdatedAt = model.UpdatedAt
	t.ScheduledAfter = model.ScheduledAfter

	t.decodeData(model.DataEncoded)
	t.decodeLogs(model.LogsEncoded)
	t.decodeTargets(model.TargetsEncoded)
	return t
}

func (t *Task) LogAttempt(log TaskLog) {
	t.Logs = append(t.Logs, log)
}

func (t *Task) OnCreated() {
	t.LastRun = time.Now()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}
func (t *Task) OnUpdated() {
	t.UpdatedAt = time.Now()
}

func (t *Task) encodeTargets() []byte {
	return []byte(strings.Join(t.Targets, ";"))
}
func (t *Task) decodeTargets(targets []byte) {
	if string(targets) != "" {
		t.Targets = strings.Split(string(targets), ";")
	}
}

func (t *Task) encodeLogs() []byte {
	if t.Logs == nil {
		return []byte{}
	}
	data, _ := json.Marshal(t.Logs)
	return data
}
func (t *Task) decodeLogs(logs []byte) {
	_ = json.Unmarshal(logs, &t.Logs)
}

func (t *Task) encodeData() []byte {
	if t.Data == nil {
		return []byte{}
	}
	data, _ := encoding.EncodeGob(t.Data)
	return data
}
func (t *Task) decodeData(data []byte) {
	t.Data = make(map[string]any)
	_ = encoding.DecodeGob(data, &t.Data)
}

type TaskLog struct {
	Targets   []string  `json:"targets" bson:"targets"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Comment   string    `json:"result" bson:"result"`
	Error     string    `json:"error" bson:"error"`
}

func NewAttemptLog(task Task, comment, err string) TaskLog {
	return TaskLog{
		Targets:   task.Targets,
		Timestamp: time.Now(),
		Comment:   comment,
		Error:     err,
	}
}

/*
Returns up limit tasks that have TaskStatusInProgress and were last updates 1+ hours ago
*/
func (c *client) GetStaleTasks(ctx context.Context, limit int) ([]Task, error) {
	models, err := c.prisma.CronTask.FindMany(db.CronTask.And(
		db.CronTask.Status.Equals(string(TaskStatusInProgress)),
		db.CronTask.UpdatedAt.Before(time.Now().Add(time.Hour*-1)),
	)).OrderBy(db.CronTask.ScheduledAfter.Order(db.ASC)).Take(limit).Exec(ctx)
	if err != nil && !db.IsErrNotFound(err) {
		return nil, err
	}

	var tasks []Task
	for _, model := range models {
		tasks = append(tasks, Task{}.FromModel(model))
	}

	return tasks, nil
}

/*
GetAndStartTasks retrieves up to limit number of tasks matching the referenceId and updates their status to in progress
  - this func will block until all other calls to task update funcs are done
*/
func (c *client) GetAndStartTasks(ctx context.Context, limit int) ([]Task, error) {
	if limit < 1 {
		return nil, nil
	}
	if err := c.tasksUpdateSem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer c.tasksUpdateSem.Release(1)

	models, err := c.prisma.CronTask.
		FindMany(
			db.CronTask.Status.Equals(string(TaskStatusScheduled)),
			db.CronTask.ScheduledAfter.Lt(time.Now()),
		).
		OrderBy(db.CronTask.ScheduledAfter.Order(db.ASC)).
		Take(limit).
		Exec(ctx)
	if err != nil {
		if db.IsErrNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	if len(models) < 1 {
		return nil, nil
	}

	var ids []string
	var tasks []Task
	for _, model := range models {
		tasks = append(tasks, Task{}.FromModel(model))
		ids = append(ids, model.ID)
	}

	_, err = c.prisma.CronTask.
		FindMany(db.CronTask.ID.In(ids)).
		Update(db.CronTask.Status.Set(string(TaskStatusInProgress))).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (c *client) CreateTasks(ctx context.Context, tasks ...Task) error {
	if len(tasks) < 1 {
		return nil
	}
	// we do not block using c.tasksUpdateSem here because the order of ops in GetAndStartTasks makes this safe

	var txns []db.PrismaTransaction
	for _, task := range tasks {
		task.OnCreated()
		txns = append(txns, c.prisma.CronTask.CreateOne(
			db.CronTask.Type.Set(string(task.Type)),
			db.CronTask.ReferenceID.Set(task.ReferenceID),
			db.CronTask.TargetsEncoded.Set(task.encodeTargets()),
			db.CronTask.Status.Set(string(TaskStatusScheduled)),
			db.CronTask.LastRun.Set(time.Now()),
			db.CronTask.ScheduledAfter.Set(task.ScheduledAfter),
			db.CronTask.LogsEncoded.Set(task.encodeLogs()),
			db.CronTask.DataEncoded.Set(task.encodeData()),
		).Tx())
	}

	err := c.prisma.Prisma.Transaction(txns...).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

/*
UpdateTasks will update all tasks passed in
  - the following fields will be replaced: targets, status, leastRun, scheduleAfterm logs, data
  - this func will block until all other calls to task update funcs are done
*/
func (c *client) UpdateTasks(ctx context.Context, tasks ...Task) error {
	if len(tasks) < 1 {
		return nil
	}

	if err := c.tasksUpdateSem.Acquire(ctx, 1); err != nil {
		return err
	}
	defer c.tasksUpdateSem.Release(1)

	var txns []db.PrismaTransaction
	for _, task := range tasks {
		task.OnUpdated()
		txns = append(txns, c.prisma.CronTask.
			FindUnique(db.CronTask.ID.Equals(task.ID)).
			Update(
				db.CronTask.TargetsEncoded.Set(task.encodeTargets()),
				db.CronTask.Status.Set(string(task.Status)),
				db.CronTask.LastRun.Set(task.LastRun),
				db.CronTask.ScheduledAfter.Set(task.ScheduledAfter),
				db.CronTask.LogsEncoded.Set(task.encodeLogs()),
				db.CronTask.DataEncoded.Set(task.encodeData()),
			).Tx())
	}

	err := c.prisma.Prisma.Transaction(txns...).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

/*
DeleteTasks will delete all tasks matching by ids
  - this func will block until all other calls to task update funcs are done
*/
func (c *client) DeleteTasks(ctx context.Context, ids ...string) error {
	if len(ids) < 1 {
		return nil
	}

	if err := c.tasksUpdateSem.Acquire(ctx, 1); err != nil {
		return nil
	}
	defer c.tasksUpdateSem.Release(1)

	_, err := c.prisma.CronTask.FindMany(db.CronTask.ID.In(ids)).Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
