// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/applicationcommand"
)

// ApplicationCommand is the model entity for the ApplicationCommand schema.
type ApplicationCommand struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// OptionsHash holds the value of the "options_hash" field.
	OptionsHash  string `json:"options_hash,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ApplicationCommand) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case applicationcommand.FieldCreatedAt, applicationcommand.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case applicationcommand.FieldID, applicationcommand.FieldName, applicationcommand.FieldVersion, applicationcommand.FieldOptionsHash:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ApplicationCommand fields.
func (ac *ApplicationCommand) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case applicationcommand.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ac.ID = value.String
			}
		case applicationcommand.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = value.Int64
			}
		case applicationcommand.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = value.Int64
			}
		case applicationcommand.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ac.Name = value.String
			}
		case applicationcommand.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				ac.Version = value.String
			}
		case applicationcommand.FieldOptionsHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field options_hash", values[i])
			} else if value.Valid {
				ac.OptionsHash = value.String
			}
		default:
			ac.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ApplicationCommand.
// This includes values selected through modifiers, order, etc.
func (ac *ApplicationCommand) Value(name string) (ent.Value, error) {
	return ac.selectValues.Get(name)
}

// Update returns a builder for updating this ApplicationCommand.
// Note that you need to call ApplicationCommand.Unwrap() before calling this method if this ApplicationCommand
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *ApplicationCommand) Update() *ApplicationCommandUpdateOne {
	return NewApplicationCommandClient(ac.config).UpdateOne(ac)
}

// Unwrap unwraps the ApplicationCommand entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *ApplicationCommand) Unwrap() *ApplicationCommand {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("db: ApplicationCommand is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *ApplicationCommand) String() string {
	var builder strings.Builder
	builder.WriteString("ApplicationCommand(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(ac.Name)
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(ac.Version)
	builder.WriteString(", ")
	builder.WriteString("options_hash=")
	builder.WriteString(ac.OptionsHash)
	builder.WriteByte(')')
	return builder.String()
}

// ApplicationCommands is a parsable slice of ApplicationCommand.
type ApplicationCommands []*ApplicationCommand
