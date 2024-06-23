// Code generated by ent, DO NOT EDIT.

package crontask

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/models"
)

const (
	// Label holds the string label denoting the crontask type in the database.
	Label = "cron_task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldReferenceID holds the string denoting the reference_id field in the database.
	FieldReferenceID = "reference_id"
	// FieldTargets holds the string denoting the targets field in the database.
	FieldTargets = "targets"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldScheduledAfter holds the string denoting the scheduled_after field in the database.
	FieldScheduledAfter = "scheduled_after"
	// FieldLastRun holds the string denoting the last_run field in the database.
	FieldLastRun = "last_run"
	// FieldLogs holds the string denoting the logs field in the database.
	FieldLogs = "logs"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// Table holds the table name of the crontask in the database.
	Table = "cron_tasks"
)

// Columns holds all SQL columns for crontask fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldType,
	FieldReferenceID,
	FieldTargets,
	FieldStatus,
	FieldScheduledAfter,
	FieldLastRun,
	FieldLogs,
	FieldData,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() int64
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() int64
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int64
	// ReferenceIDValidator is a validator for the "reference_id" field. It is called by the builders before save.
	ReferenceIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type models.TaskType) error {
	switch _type {
	case "UPDATE_CLANS", "RECORD_ACCOUNT_SESSIONS", "UPDATE_ACCOUNT_WN8", "UPDATE_ACCOUNT_ACHIEVEMENTS", "CLEANUP_DATABASE":
		return nil
	default:
		return fmt.Errorf("crontask: invalid enum value for type field: %q", _type)
	}
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s models.TaskStatus) error {
	switch s {
	case "TASK_SCHEDULED", "TASK_IN_PROGRESS", "TASK_COMPLETE", "TASK_FAILED":
		return nil
	default:
		return fmt.Errorf("crontask: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the CronTask queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByReferenceID orders the results by the reference_id field.
func ByReferenceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReferenceID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByScheduledAfter orders the results by the scheduled_after field.
func ByScheduledAfter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScheduledAfter, opts...).ToFunc()
}

// ByLastRun orders the results by the last_run field.
func ByLastRun(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastRun, opts...).ToFunc()
}
