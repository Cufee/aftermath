// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/ent/db/usercontent"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserContent is the model entity for the UserContent schema.
type UserContent struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Type holds the value of the "type" field.
	Type models.UserContentType `json:"type,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// ReferenceID holds the value of the "reference_id" field.
	ReferenceID string `json:"reference_id,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// Metadata holds the value of the "metadata" field.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserContentQuery when eager-loading is set.
	Edges        UserContentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserContentEdges holds the relations/edges for other nodes in the graph.
type UserContentEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserContentEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserContent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usercontent.FieldMetadata:
			values[i] = new([]byte)
		case usercontent.FieldID, usercontent.FieldType, usercontent.FieldUserID, usercontent.FieldReferenceID, usercontent.FieldValue:
			values[i] = new(sql.NullString)
		case usercontent.FieldCreatedAt, usercontent.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserContent fields.
func (uc *UserContent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usercontent.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				uc.ID = value.String
			}
		case usercontent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				uc.CreatedAt = value.Time
			}
		case usercontent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				uc.UpdatedAt = value.Time
			}
		case usercontent.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				uc.Type = models.UserContentType(value.String)
			}
		case usercontent.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uc.UserID = value.String
			}
		case usercontent.FieldReferenceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference_id", values[i])
			} else if value.Valid {
				uc.ReferenceID = value.String
			}
		case usercontent.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				uc.Value = value.String
			}
		case usercontent.FieldMetadata:
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

// GetValue returns the ent.Value that was dynamically selected and assigned to the UserContent.
// This includes values selected through modifiers, order, etc.
func (uc *UserContent) GetValue(name string) (ent.Value, error) {
	return uc.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserContent entity.
func (uc *UserContent) QueryUser() *UserQuery {
	return NewUserContentClient(uc.config).QueryUser(uc)
}

// Update returns a builder for updating this UserContent.
// Note that you need to call UserContent.Unwrap() before calling this method if this UserContent
// was returned from a transaction, and the transaction was committed or rolled back.
func (uc *UserContent) Update() *UserContentUpdateOne {
	return NewUserContentClient(uc.config).UpdateOne(uc)
}

// Unwrap unwraps the UserContent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uc *UserContent) Unwrap() *UserContent {
	_tx, ok := uc.config.driver.(*txDriver)
	if !ok {
		panic("db: UserContent is not a transactional entity")
	}
	uc.config.driver = _tx.drv
	return uc
}

// String implements the fmt.Stringer.
func (uc *UserContent) String() string {
	var builder strings.Builder
	builder.WriteString("UserContent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", uc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(uc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(uc.UpdatedAt.Format(time.ANSIC))
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
	builder.WriteString("value=")
	builder.WriteString(uc.Value)
	builder.WriteString(", ")
	builder.WriteString("metadata=")
	builder.WriteString(fmt.Sprintf("%v", uc.Metadata))
	builder.WriteByte(')')
	return builder.String()
}

// UserContents is a parsable slice of UserContent.
type UserContents []*UserContent
