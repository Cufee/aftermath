// Code generated by ent, DO NOT EDIT.

package achievementssnapshot

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cufee/aftermath/internal/database/models"
)

const (
	// Label holds the string label denoting the achievementssnapshot type in the database.
	Label = "achievements_snapshot"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldReferenceID holds the string denoting the reference_id field in the database.
	FieldReferenceID = "reference_id"
	// FieldBattles holds the string denoting the battles field in the database.
	FieldBattles = "battles"
	// FieldLastBattleTime holds the string denoting the last_battle_time field in the database.
	FieldLastBattleTime = "last_battle_time"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// EdgeAccount holds the string denoting the account edge name in mutations.
	EdgeAccount = "account"
	// Table holds the table name of the achievementssnapshot in the database.
	Table = "achievements_snapshots"
	// AccountTable is the table that holds the account relation/edge.
	AccountTable = "achievements_snapshots"
	// AccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountInverseTable = "accounts"
	// AccountColumn is the table column denoting the account relation/edge.
	AccountColumn = "account_id"
)

// Columns holds all SQL columns for achievementssnapshot fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldType,
	FieldAccountID,
	FieldReferenceID,
	FieldBattles,
	FieldLastBattleTime,
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
	DefaultCreatedAt func() int
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() int
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int
	// AccountIDValidator is a validator for the "account_id" field. It is called by the builders before save.
	AccountIDValidator func(string) error
	// ReferenceIDValidator is a validator for the "reference_id" field. It is called by the builders before save.
	ReferenceIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type models.SnapshotType) error {
	switch _type {
	case "live", "daily":
		return nil
	default:
		return fmt.Errorf("achievementssnapshot: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the AchievementsSnapshot queries.
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

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByReferenceID orders the results by the reference_id field.
func ByReferenceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldReferenceID, opts...).ToFunc()
}

// ByBattles orders the results by the battles field.
func ByBattles(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBattles, opts...).ToFunc()
}

// ByLastBattleTime orders the results by the last_battle_time field.
func ByLastBattleTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastBattleTime, opts...).ToFunc()
}

// ByAccountField orders the results by account field.
func ByAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountStep(), sql.OrderByField(field, opts...))
	}
}
func newAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
	)
}
