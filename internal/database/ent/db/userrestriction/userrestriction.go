// Code generated by ent, DO NOT EDIT.

package userrestriction

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cufee/aftermath/internal/database/models"
)

const (
	// Label holds the string label denoting the userrestriction type in the database.
	Label = "user_restriction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRestriction holds the string denoting the restriction field in the database.
	FieldRestriction = "restriction"
	// FieldPublicReason holds the string denoting the public_reason field in the database.
	FieldPublicReason = "public_reason"
	// FieldModeratorComment holds the string denoting the moderator_comment field in the database.
	FieldModeratorComment = "moderator_comment"
	// FieldEvents holds the string denoting the events field in the database.
	FieldEvents = "events"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the userrestriction in the database.
	Table = "user_restrictions"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "user_restrictions"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for userrestriction fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldExpiresAt,
	FieldType,
	FieldUserID,
	FieldRestriction,
	FieldPublicReason,
	FieldModeratorComment,
	FieldEvents,
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
	// UserIDValidator is a validator for the "user_id" field. It is called by the builders before save.
	UserIDValidator func(string) error
	// RestrictionValidator is a validator for the "restriction" field. It is called by the builders before save.
	RestrictionValidator func(string) error
	// PublicReasonValidator is a validator for the "public_reason" field. It is called by the builders before save.
	PublicReasonValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type models.UserRestrictionType) error {
	switch _type {
	case "partial", "complete":
		return nil
	default:
		return fmt.Errorf("userrestriction: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the UserRestriction queries.
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

// ByExpiresAt orders the results by the expires_at field.
func ByExpiresAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiresAt, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByRestriction orders the results by the restriction field.
func ByRestriction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRestriction, opts...).ToFunc()
}

// ByPublicReason orders the results by the public_reason field.
func ByPublicReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicReason, opts...).ToFunc()
}

// ByModeratorComment orders the results by the moderator_comment field.
func ByModeratorComment(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModeratorComment, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}