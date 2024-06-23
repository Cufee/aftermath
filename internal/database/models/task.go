package models

import (
	"time"
)

type TaskType string

const (
	TaskTypeUpdateClans              TaskType = "UPDATE_CLANS"
	TaskTypeRecordSessions           TaskType = "RECORD_ACCOUNT_SESSIONS"
	TaskTypeUpdateAccountWN8         TaskType = "UPDATE_ACCOUNT_WN8"
	TaskTypeRecordPlayerAchievements TaskType = "UPDATE_ACCOUNT_ACHIEVEMENTS"

	TaskTypeDatabaseCleanup TaskType = "CLEANUP_DATABASE"
)

// Values provides list valid values for Enum.
func (TaskType) Values() []string {
	var kinds []string
	for _, s := range []TaskType{
		TaskTypeUpdateClans,
		TaskTypeRecordSessions,
		TaskTypeUpdateAccountWN8,
		TaskTypeRecordPlayerAchievements,
		//
		TaskTypeDatabaseCleanup,
	} {
		kinds = append(kinds, string(s))
	}
	return kinds
}

// Task statuses
type TaskStatus string

const (
	TaskStatusScheduled  TaskStatus = "TASK_SCHEDULED"
	TaskStatusInProgress TaskStatus = "TASK_IN_PROGRESS"
	TaskStatusComplete   TaskStatus = "TASK_COMPLETE"
	TaskStatusFailed     TaskStatus = "TASK_FAILED"
)

// Values provides list valid values for Enum.
func (TaskStatus) Values() []string {
	var kinds []string
	for _, s := range []TaskStatus{
		TaskStatusScheduled,
		TaskStatusInProgress,
		TaskStatusComplete,
		TaskStatusFailed,
	} {
		kinds = append(kinds, string(s))
	}
	return kinds
}

type Task struct {
	ID        string    `json:"id"`
	Type      TaskType  `json:"kind"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	ReferenceID string   `json:"referenceId"`
	Targets     []string `json:"targets"`

	Logs []TaskLog `json:"logs"`

	Status         TaskStatus `json:"status"`
	ScheduledAfter time.Time  `json:"scheduledAfter"`
	LastRun        time.Time  `json:"lastRun"`

	Data map[string]any `json:"data"`
}

func (t *Task) LogAttempt(log TaskLog) {
	t.Logs = append(t.Logs, log)
}

func (t *Task) OnCreated() {
	t.Status = TaskStatusScheduled
	t.LastRun = time.Now()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.Logs = append(t.Logs, TaskLog{
		Comment:   "task created",
		Timestamp: time.Now(),
	})
}
func (t *Task) OnUpdated() {
	t.LastRun = time.Now()
	t.UpdatedAt = time.Now()
	t.Logs = append(t.Logs, TaskLog{
		Comment:   "task updated",
		Timestamp: time.Now(),
	})
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
