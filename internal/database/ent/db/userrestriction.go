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
	"github.com/cufee/aftermath/internal/database/ent/db/userrestriction"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserRestriction is the model entity for the UserRestriction schema.
type UserRestriction struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// Type holds the value of the "type" field.
	Type models.UserRestrictionType `json:"type,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Restriction holds the value of the "restriction" field.
	Restriction string `json:"restriction,omitempty"`
	// PublicReason holds the value of the "public_reason" field.
	PublicReason string `json:"public_reason,omitempty"`
	// ModeratorComment holds the value of the "moderator_comment" field.
	ModeratorComment string `json:"moderator_comment,omitempty"`
	// Events holds the value of the "events" field.
	Events []models.RestrictionUpdate `json:"events,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserRestrictionQuery when eager-loading is set.
	Edges        UserRestrictionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserRestrictionEdges holds the relations/edges for other nodes in the graph.
type UserRestrictionEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserRestrictionEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserRestriction) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userrestriction.FieldEvents:
			values[i] = new([]byte)
		case userrestriction.FieldID, userrestriction.FieldType, userrestriction.FieldUserID, userrestriction.FieldRestriction, userrestriction.FieldPublicReason, userrestriction.FieldModeratorComment:
			values[i] = new(sql.NullString)
		case userrestriction.FieldCreatedAt, userrestriction.FieldUpdatedAt, userrestriction.FieldExpiresAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserRestriction fields.
func (ur *UserRestriction) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userrestriction.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ur.ID = value.String
			}
		case userrestriction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ur.CreatedAt = value.Time
			}
		case userrestriction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ur.UpdatedAt = value.Time
			}
		case userrestriction.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				ur.ExpiresAt = value.Time
			}
		case userrestriction.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				ur.Type = models.UserRestrictionType(value.String)
			}
		case userrestriction.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ur.UserID = value.String
			}
		case userrestriction.FieldRestriction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field restriction", values[i])
			} else if value.Valid {
				ur.Restriction = value.String
			}
		case userrestriction.FieldPublicReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_reason", values[i])
			} else if value.Valid {
				ur.PublicReason = value.String
			}
		case userrestriction.FieldModeratorComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field moderator_comment", values[i])
			} else if value.Valid {
				ur.ModeratorComment = value.String
			}
		case userrestriction.FieldEvents:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field events", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ur.Events); err != nil {
					return fmt.Errorf("unmarshal field events: %w", err)
				}
			}
		default:
			ur.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserRestriction.
// This includes values selected through modifiers, order, etc.
func (ur *UserRestriction) Value(name string) (ent.Value, error) {
	return ur.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserRestriction entity.
func (ur *UserRestriction) QueryUser() *UserQuery {
	return NewUserRestrictionClient(ur.config).QueryUser(ur)
}

// Update returns a builder for updating this UserRestriction.
// Note that you need to call UserRestriction.Unwrap() before calling this method if this UserRestriction
// was returned from a transaction, and the transaction was committed or rolled back.
func (ur *UserRestriction) Update() *UserRestrictionUpdateOne {
	return NewUserRestrictionClient(ur.config).UpdateOne(ur)
}

// Unwrap unwraps the UserRestriction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ur *UserRestriction) Unwrap() *UserRestriction {
	_tx, ok := ur.config.driver.(*txDriver)
	if !ok {
		panic("db: UserRestriction is not a transactional entity")
	}
	ur.config.driver = _tx.drv
	return ur
}

// String implements the fmt.Stringer.
func (ur *UserRestriction) String() string {
	var builder strings.Builder
	builder.WriteString("UserRestriction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ur.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ur.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ur.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(ur.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", ur.Type))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(ur.UserID)
	builder.WriteString(", ")
	builder.WriteString("restriction=")
	builder.WriteString(ur.Restriction)
	builder.WriteString(", ")
	builder.WriteString("public_reason=")
	builder.WriteString(ur.PublicReason)
	builder.WriteString(", ")
	builder.WriteString("moderator_comment=")
	builder.WriteString(ur.ModeratorComment)
	builder.WriteString(", ")
	builder.WriteString("events=")
	builder.WriteString(fmt.Sprintf("%v", ur.Events))
	builder.WriteByte(')')
	return builder.String()
}

// UserRestrictions is a parsable slice of UserRestriction.
type UserRestrictions []*UserRestriction
