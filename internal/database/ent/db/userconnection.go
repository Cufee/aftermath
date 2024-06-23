// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/ent/db/userconnection"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserConnection is the model entity for the UserConnection schema.
type UserConnection struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Type holds the value of the "type" field.
	Type models.ConnectionType `json:"type,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// ReferenceID holds the value of the "reference_id" field.
	ReferenceID string `json:"reference_id,omitempty"`
	// Permissions holds the value of the "permissions" field.
	Permissions string `json:"permissions,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserConnectionQuery when eager-loading is set.
	Edges        UserConnectionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserConnectionEdges holds the relations/edges for other nodes in the graph.
type UserConnectionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserConnectionEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserConnection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userconnection.FieldMetadata:
			values[i] = new([]byte)
		case userconnection.FieldCreatedAt, userconnection.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case userconnection.FieldID, userconnection.FieldType, userconnection.FieldUserID, userconnection.FieldReferenceID, userconnection.FieldPermissions:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserConnection fields.
func (uc *UserConnection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userconnection.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				uc.ID = value.String
			}
		case userconnection.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				uc.CreatedAt = value.Int64
			}
		case userconnection.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				uc.UpdatedAt = value.Int64
			}
		case userconnection.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				uc.Type = models.ConnectionType(value.String)
			}
		case userconnection.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uc.UserID = value.String
			}
		case userconnection.FieldReferenceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference_id", values[i])
			} else if value.Valid {
				uc.ReferenceID = value.String
			}
		case userconnection.FieldPermissions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permissions", values[i])
			} else if value.Valid {
				uc.Permissions = value.String
			}
		case userconnection.FieldMetadata:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field metadata", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &uc.Metadata); err != nil {
					return fmt.Errorf("unmarshal field metadata: %w", err)
				}
			}
		default:
			uc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserConnection.
// This includes values selected through modifiers, order, etc.
func (uc *UserConnection) Value(name string) (ent.Value, error) {
	return uc.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserConnection entity.
func (uc *UserConnection) QueryUser() *UserQuery {
	return NewUserConnectionClient(uc.config).QueryUser(uc)
}

// Update returns a builder for updating this UserConnection.
// Note that you need to call UserConnection.Unwrap() before calling this method if this UserConnection
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserConnection) Update() *UserConnectionUpdateOne {
	return NewUserConnectionClient(uc.config).UpdateOne(uc)
}

// Unwrap unwraps the UserConnection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserConnection) Unwrap() *UserConnection {
	_tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("db: UserConnection is not a transactional entity")
	}
	uc.config.driver = _tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserConnection) String() string {
	var builder strings.Builder
	builder.WriteString("UserConnection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", uc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", uc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", uc.Type))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(uc.UserID)
	builder.WriteString(", ")
	builder.WriteString("reference_id=")
	builder.WriteString(uc.ReferenceID)
	builder.WriteString(", ")
	builder.WriteString("permissions=")
	builder.WriteString(uc.Permissions)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", uc.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// UserConnections is a parsable slice of UserConnection.
type UserConnections []*UserConnection
