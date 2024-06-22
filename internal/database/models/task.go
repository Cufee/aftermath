package models

import (
	"encoding/json"
	"strings"
	"time"

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
