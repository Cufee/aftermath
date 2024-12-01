// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/vehicle"
	"golang.org/x/text/language"
)

// Vehicle is the model entity for the Vehicle schema.
type Vehicle struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Tier holds the value of the "tier" field.
	Tier int `json:"tier,omitempty"`
	// LocalizedNames holds the value of the "localized_names" field.
	LocalizedNames map[language.Tag]string `json:"localized_names,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Vehicle) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehicle.FieldLocalizedNames:
			values[i] = new([]byte)
		case vehicle.FieldTier:
			values[i] = new(sql.NullInt64)
		case vehicle.FieldID:
			values[i] = new(sql.NullString)
		case vehicle.FieldCreatedAt, vehicle.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Vehicle fields.
func (v *Vehicle) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehicle.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				v.ID = value.String
			}
		case vehicle.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case vehicle.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case vehicle.FieldTier:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tier", values[i])
			} else if value.Valid {
				v.Tier = int(value.Int64)
			}
		case vehicle.FieldLocalizedNames:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field localized_names", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &v.LocalizedNames); err != nil {
					return fmt.Errorf("unmarshal field localized_names: %w", err)
				}
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Vehicle.
// This includes values selected through modifiers, order, etc.
func (v *Vehicle) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// Update returns a builder for updating this Vehicle.
// Note that you need to call Vehicle.Unwrap() before calling this method if this Vehicle
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Vehicle) Update() *VehicleUpdateOne {
	return NewVehicleClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Vehicle entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Vehicle) Unwrap() *Vehicle {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("db: Vehicle is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Vehicle) String() string {
	var builder strings.Builder
	builder.WriteString("Vehicle(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("tier=")
	builder.WriteString(fmt.Sprintf("%v", v.Tier))
	builder.WriteString(", ")
	builder.WriteString("localized_names=")
	builder.WriteString(fmt.Sprintf("%v", v.LocalizedNames))
	builder.WriteByte(')')
	return builder.String()
}

// Vehicles is a parsable slice of Vehicle.
type Vehicles []*Vehicle
