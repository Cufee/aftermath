// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/clan"
)

// Clan is the model entity for the Clan schema.
type Clan struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// EmblemID holds the value of the "emblem_id" field.
	EmblemID string `json:"emblem_id,omitempty"`
	// Members holds the value of the "members" field.
	Members []string `json:"members,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ClanQuery when eager-loading is set.
	Edges        ClanEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ClanEdges holds the relations/edges for other nodes in the graph.
type ClanEdges struct {
	// Accounts holds the value of the accounts edge.
	Accounts []*Account `json:"accounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading.
func (e ClanEdges) AccountsOrErr() ([]*Account, error) {
	if e.loadedTypes[0] {
		return e.Accounts, nil
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Clan) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case clan.FieldMembers:
			values[i] = new([]byte)
		case clan.FieldCreatedAt, clan.FieldUpdatedAt:
			values[i] = new(sql.NullInt64)
		case clan.FieldID, clan.FieldTag, clan.FieldName, clan.FieldEmblemID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Clan fields.
func (c *Clan) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case clan.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case clan.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Int64
			}
		case clan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Int64
			}
		case clan.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				c.Tag = value.String
			}
		case clan.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case clan.FieldEmblemID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field emblem_id", values[i])
			} else if value.Valid {
				c.EmblemID = value.String
			}
		case clan.FieldMembers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field members", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Members); err != nil {
					return fmt.Errorf("unmarshal field members: %w", err)
				}
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Clan.
// This includes values selected through modifiers, order, etc.
func (c *Clan) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryAccounts queries the "accounts" edge of the Clan entity.
func (c *Clan) QueryAccounts() *AccountQuery {
	return NewClanClient(c.config).QueryAccounts(c)
}

// Update returns a builder for updating this Clan.
// Note that you need to call Clan.Unwrap() before calling this method if this Clan
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Clan) Update() *ClanUpdateOne {
	return NewClanClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Clan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Clan) Unwrap() *Clan {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("db: Clan is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Clan) String() string {
	var builder strings.Builder
	builder.WriteString("Clan(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", c.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", c.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("tag=")
	builder.WriteString(c.Tag)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("emblem_id=")
	builder.WriteString(c.EmblemID)
	builder.WriteString(", ")
	builder.WriteString("members=")
	builder.WriteString(fmt.Sprintf("%v", c.Members))
	builder.WriteByte(')')
	return builder.String()
}

// Clans is a parsable slice of Clan.
type Clans []*Clan