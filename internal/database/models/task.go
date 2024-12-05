package models

import (
	"time"

	"github.com/cufee/aftermath/internal/json"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
)

type TaskType string

const (
	TaskTypeUpdateClans                   TaskType = "UPDATE_CLANS"
	TaskTypeRecordSnapshots               TaskType = "RECORD_SNAPSHOTS"
	TaskTypeAchievementsLeaderboardUpdate TaskType = "ACHIEVEMENT_LEADERBOARDS"

	TaskTypeDatabaseCleanup TaskType = "CLEANUP_DATABASE"
)

func (TaskType) Values() []string {
	var kinds []string
	for _, s := range []TaskType{
		TaskTypeUpdateClans,
		TaskTypeRecordSnapshots,
		TaskTypeAchievementsLeaderboardUpdate,
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
	TriesLeft      int        `json:"triesLeft"`
	LastRun        time.Time  `json:"lastRun"`

	Data map[string]string `json:"data"`
}

func (t *Task) LogAttempt(log TaskLog) {
	t.Logs = append(t.Logs, log)
}

func (t *Task) OnCreated() {
	t.Status = TaskStatusScheduled
	t.LastRun = time.Now()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
}
func (t *Task) OnUpdated() {
	t.LastRun = time.Now()
	t.UpdatedAt = time.Now()
}

type TaskLog struct {
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Comment   string    `json:"result" bson:"result"`
	Error     string    `json:"error" bson:"error"`
}

func NewAttemptLog(task Task, comment string, err error) TaskLog {
	return TaskLog{
		Timestamp: time.Now(),
		Comment:   comment,
		Error:     err.Error(),
	}
}

func ToCronTask(record *model.CronTask) Task {
	t := Task{
		ID:             record.ID,
		Type:           TaskType(record.Type),
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
		ReferenceID:    record.ReferenceID,
		Status:         TaskStatus(record.Status),
		ScheduledAfter: record.ScheduledAfter,
		TriesLeft:      int(record.TriesLeft),
		LastRun:        record.LastRun,
	}
	json.Unmarshal(record.Targets, &t.Targets)
	json.Unmarshal(record.Data, &t.Data)
	json.Unmarshal(record.Logs, &t.Logs)
	return t
}

func (record Task) Model() model.CronTask {
	t := model.CronTask{
		ID:             utils.StringOr(record.ID, cuid.New()),
		Type:           string(record.Type),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		ReferenceID:    record.ReferenceID,
		Status:         string(record.Status),
		ScheduledAfter: record.ScheduledAfter,
		TriesLeft:      int32(record.TriesLeft),
		LastRun:        record.LastRun,
	}
	t.Targets, _ = json.Marshal(record.Targets)
	t.Data, _ = json.Marshal(record.Data)
	t.Logs, _ = json.Marshal(record.Logs)
	return t
}
