// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/gamemode"
	"golang.org/x/text/language"
)

// GameMode is the model entity for the GameMode schema.
type GameMode struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LocalizedNames holds the value of the "localized_names" field.
	LocalizedNames map[language.Tag]string `json:"localized_names,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GameMode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gamemode.FieldLocalizedNames:
			values[i] = new([]byte)
		case gamemode.FieldID:
			values[i] = new(sql.NullString)
		case gamemode.FieldCreatedAt, gamemode.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GameMode fields.
func (gm *GameMode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gamemode.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				gm.ID = value.String
			}
		case gamemode.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gm.CreatedAt = value.Time
			}
		case gamemode.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gm.UpdatedAt = value.Time
			}
		case gamemode.FieldLocalizedNames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field localized_names", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gm.LocalizedNames); err != nil {
					return fmt.Errorf("unmarshal field localized_names: %w", err)
				}
			}
		default:
			gm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GameMode.
// This includes values selected through modifiers, order, etc.
func (gm *GameMode) Value(name string) (ent.Value, error) {
	return gm.selectValues.Get(name)
}

// Update returns a builder for updating this GameMode.
// Note that you need to call GameMode.Unwrap() before calling this method if this GameMode
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GameMode) Update() *GameModeUpdateOne {
	return NewGameModeClient(gm.config).UpdateOne(gm)
}

// Unwrap unwraps the GameMode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GameMode) Unwrap() *GameMode {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("db: GameMode is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GameMode) String() string {
	var builder strings.Builder
	builder.WriteString("GameMode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("localized_names=")
	builder.WriteString(fmt.Sprintf("%v", gm.LocalizedNames))
	builder.WriteByte(')')
	return builder.String()
}

// GameModes is a parsable slice of GameMode.
type GameModes []*GameMode
