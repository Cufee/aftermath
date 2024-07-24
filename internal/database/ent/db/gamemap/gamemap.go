// Code generated by ent, DO NOT EDIT.

package gamemap

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the gamemap type in the database.
	Label = "game_map"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldGameModes holds the string denoting the game_modes field in the database.
	FieldGameModes = "game_modes"
	// FieldSupremacyPoints holds the string denoting the supremacy_points field in the database.
	FieldSupremacyPoints = "supremacy_points"
	// FieldLocalizedNames holds the string denoting the localized_names field in the database.
	FieldLocalizedNames = "localized_names"
	// Table holds the table name of the gamemap in the database.
	Table = "game_maps"
)

// Columns holds all SQL columns for gamemap fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldGameModes,
	FieldSupremacyPoints,
	FieldLocalizedNames,
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
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the GameMap queries.
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

// BySupremacyPoints orders the results by the supremacy_points field.
func BySupremacyPoints(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupremacyPoints, opts...).ToFunc()
}
