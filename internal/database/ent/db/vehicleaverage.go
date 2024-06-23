// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/vehicleaverage"
	"github.com/cufee/aftermath/internal/stats/frame"
)

// VehicleAverage is the model entity for the VehicleAverage schema.
type VehicleAverage struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Data holds the value of the "data" field.
	Data         frame.StatsFrame `json:"data,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VehicleAverage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehicleaverage.FieldData:
			values[i] = new([]byte)
		case vehicleaverage.FieldCreatedAt, vehicleaverage.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case vehicleaverage.FieldID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VehicleAverage fields.
func (va *VehicleAverage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehicleaverage.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				va.ID = value.String
			}
		case vehicleaverage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				va.CreatedAt = value.Int64
			}
		case vehicleaverage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				va.UpdatedAt = value.Int64
			}
		case vehicleaverage.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &va.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %w", err)
				}
			}
		default:
			va.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VehicleAverage.
// This includes values selected through modifiers, order, etc.
func (va *VehicleAverage) Value(name string) (ent.Value, error) {
	return va.selectValues.Get(name)
}

// Update returns a builder for updating this VehicleAverage.
// Note that you need to call VehicleAverage.Unwrap() before calling this method if this VehicleAverage
// was returned from a transaction, and the transaction was committed or rolled back.
func (va *VehicleAverage) Update() *VehicleAverageUpdateOne {
	return NewVehicleAverageClient(va.config).UpdateOne(va)
}

// Unwrap unwraps the VehicleAverage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (va *VehicleAverage) Unwrap() *VehicleAverage {
	_tx, ok := va.config.driver.(*txDriver)
	if !ok {
		panic("db: VehicleAverage is not a transactional entity")
	}
	va.config.driver = _tx.drv
	return va
}

// String implements the fmt.Stringer.
func (va *VehicleAverage) String() string {
	var builder strings.Builder
	builder.WriteString("VehicleAverage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", va.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", va.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", va.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("data=")
	builder.WriteString(fmt.Sprintf("%v", va.Data))
	builder.WriteByte(')')
	return builder.String()
}

// VehicleAverages is a parsable slice of VehicleAverage.
type VehicleAverages []*VehicleAverage
