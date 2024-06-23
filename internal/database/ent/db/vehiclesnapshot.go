// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

// VehicleSnapshot is the model entity for the VehicleSnapshot schema.
type VehicleSnapshot struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Type holds the value of the "type" field.
	Type models.SnapshotType `json:"type,omitempty"`
	// AccountID holds the value of the "account_id" field.
	AccountID string `json:"account_id,omitempty"`
	// VehicleID holds the value of the "vehicle_id" field.
	VehicleID string `json:"vehicle_id,omitempty"`
	// ReferenceID holds the value of the "reference_id" field.
	ReferenceID string `json:"reference_id,omitempty"`
	// Battles holds the value of the "battles" field.
	Battles int `json:"battles,omitempty"`
	// LastBattleTime holds the value of the "last_battle_time" field.
	LastBattleTime int64 `json:"last_battle_time,omitempty"`
	// Frame holds the value of the "frame" field.
	Frame frame.StatsFrame `json:"frame,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VehicleSnapshotQuery when eager-loading is set.
	Edges        VehicleSnapshotEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VehicleSnapshotEdges holds the relations/edges for other nodes in the graph.
type VehicleSnapshotEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VehicleSnapshotEdges) AccountOrErr() (*Account, error) {
	if e.Account != nil {
		return e.Account, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: account.Label}
	}
	return nil, &NotLoadedError{edge: "account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VehicleSnapshot) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehiclesnapshot.FieldFrame:
			values[i] = new([]byte)
		case vehiclesnapshot.FieldCreatedAt, vehiclesnapshot.FieldUpdatedAt, vehiclesnapshot.FieldBattles, vehiclesnapshot.FieldLastBattleTime:
			values[i] = new(sql.NullInt64)
		case vehiclesnapshot.FieldID, vehiclesnapshot.FieldType, vehiclesnapshot.FieldAccountID, vehiclesnapshot.FieldVehicleID, vehiclesnapshot.FieldReferenceID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VehicleSnapshot fields.
func (vs *VehicleSnapshot) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehiclesnapshot.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				vs.ID = value.String
			}
		case vehiclesnapshot.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vs.CreatedAt = value.Int64
			}
		case vehiclesnapshot.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vs.UpdatedAt = value.Int64
			}
		case vehiclesnapshot.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				vs.Type = models.SnapshotType(value.String)
			}
		case vehiclesnapshot.FieldAccountID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				vs.AccountID = value.String
			}
		case vehiclesnapshot.FieldVehicleID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field vehicle_id", values[i])
			} else if value.Valid {
				vs.VehicleID = value.String
			}
		case vehiclesnapshot.FieldReferenceID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reference_id", values[i])
			} else if value.Valid {
				vs.ReferenceID = value.String
			}
		case vehiclesnapshot.FieldBattles:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field battles", values[i])
			} else if value.Valid {
				vs.Battles = int(value.Int64)
			}
		case vehiclesnapshot.FieldLastBattleTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_battle_time", values[i])
			} else if value.Valid {
				vs.LastBattleTime = value.Int64
			}
		case vehiclesnapshot.FieldFrame:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field frame", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &vs.Frame); err != nil {
					return fmt.Errorf("unmarshal field frame: %w", err)
				}
			}
		default:
			vs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VehicleSnapshot.
// This includes values selected through modifiers, order, etc.
func (vs *VehicleSnapshot) Value(name string) (ent.Value, error) {
	return vs.selectValues.Get(name)
}

// QueryAccount queries the "account" edge of the VehicleSnapshot entity.
func (vs *VehicleSnapshot) QueryAccount() *AccountQuery {
	return NewVehicleSnapshotClient(vs.config).QueryAccount(vs)
}

// Update returns a builder for updating this VehicleSnapshot.
// Note that you need to call VehicleSnapshot.Unwrap() before calling this method if this VehicleSnapshot
// was returned from a transaction, and the transaction was committed or rolled back.
func (vs *VehicleSnapshot) Update() *VehicleSnapshotUpdateOne {
	return NewVehicleSnapshotClient(vs.config).UpdateOne(vs)
}

// Unwrap unwraps the VehicleSnapshot entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vs *VehicleSnapshot) Unwrap() *VehicleSnapshot {
	_tx, ok := vs.config.driver.(*txDriver)
	if !ok {
		panic("db: VehicleSnapshot is not a transactional entity")
	}
	vs.config.driver = _tx.drv
	return vs
}

// String implements the fmt.Stringer.
func (vs *VehicleSnapshot) String() string {
	var builder strings.Builder
	builder.WriteString("VehicleSnapshot(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", vs.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", vs.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", vs.Type))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(vs.AccountID)
	builder.WriteString(", ")
	builder.WriteString("vehicle_id=")
	builder.WriteString(vs.VehicleID)
	builder.WriteString(", ")
	builder.WriteString("reference_id=")
	builder.WriteString(vs.ReferenceID)
	builder.WriteString(", ")
	builder.WriteString("battles=")
	builder.WriteString(fmt.Sprintf("%v", vs.Battles))
	builder.WriteString(", ")
	builder.WriteString("last_battle_time=")
	builder.WriteString(fmt.Sprintf("%v", vs.LastBattleTime))
	builder.WriteString(", ")
	builder.WriteString("frame=")
	builder.WriteString(fmt.Sprintf("%v", vs.Frame))
	builder.WriteByte(')')
	return builder.String()
}

// VehicleSnapshots is a parsable slice of VehicleSnapshot.
type VehicleSnapshots []*VehicleSnapshot