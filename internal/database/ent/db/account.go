// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/ent/db/clan"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// LastBattleTime holds the value of the "last_battle_time" field.
	LastBattleTime int64 `json:"last_battle_time,omitempty"`
	// AccountCreatedAt holds the value of the "account_created_at" field.
	AccountCreatedAt int64 `json:"account_created_at,omitempty"`
	// Realm holds the value of the "realm" field.
	Realm string `json:"realm,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Private holds the value of the "private" field.
	Private bool `json:"private,omitempty"`
	// ClanID holds the value of the "clan_id" field.
	ClanID string `json:"clan_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges        AccountEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Clan holds the value of the clan edge.
	Clan *Clan `json:"clan,omitempty"`
	// Snapshots holds the value of the snapshots edge.
	Snapshots []*AccountSnapshot `json:"snapshots,omitempty"`
	// VehicleSnapshots holds the value of the vehicle_snapshots edge.
	VehicleSnapshots []*VehicleSnapshot `json:"vehicle_snapshots,omitempty"`
	// AchievementSnapshots holds the value of the achievement_snapshots edge.
	AchievementSnapshots []*AchievementsSnapshot `json:"achievement_snapshots,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// ClanOrErr returns the Clan value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) ClanOrErr() (*Clan, error) {
	if e.Clan != nil {
		return e.Clan, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: clan.Label}
	}
	return nil, &NotLoadedError{edge: "clan"}
}

// SnapshotsOrErr returns the Snapshots value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) SnapshotsOrErr() ([]*AccountSnapshot, error) {
	if e.loadedTypes[1] {
		return e.Snapshots, nil
	}
	return nil, &NotLoadedError{edge: "snapshots"}
}

// VehicleSnapshotsOrErr returns the VehicleSnapshots value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) VehicleSnapshotsOrErr() ([]*VehicleSnapshot, error) {
	if e.loadedTypes[2] {
		return e.VehicleSnapshots, nil
	}
	return nil, &NotLoadedError{edge: "vehicle_snapshots"}
}

// AchievementSnapshotsOrErr returns the AchievementSnapshots value or an error if the edge
// was not loaded in eager-loading.
func (e AccountEdges) AchievementSnapshotsOrErr() ([]*AchievementsSnapshot, error) {
	if e.loadedTypes[3] {
		return e.AchievementSnapshots, nil
	}
	return nil, &NotLoadedError{edge: "achievement_snapshots"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldPrivate:
			values[i] = new(sql.NullBool)
		case account.FieldCreatedAt, account.FieldUpdatedAt, account.FieldLastBattleTime, account.FieldAccountCreatedAt:
			values[i] = new(sql.NullInt64)
		case account.FieldID, account.FieldRealm, account.FieldNickname, account.FieldClanID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				a.ID = value.String
			}
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Int64
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Int64
			}
		case account.FieldLastBattleTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_battle_time", values[i])
			} else if value.Valid {
				a.LastBattleTime = value.Int64
			}
		case account.FieldAccountCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_created_at", values[i])
			} else if value.Valid {
				a.AccountCreatedAt = value.Int64
			}
		case account.FieldRealm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field realm", values[i])
			} else if value.Valid {
				a.Realm = value.String
			}
		case account.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				a.Nickname = value.String
			}
		case account.FieldPrivate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field private", values[i])
			} else if value.Valid {
				a.Private = value.Bool
			}
		case account.FieldClanID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field clan_id", values[i])
			} else if value.Valid {
				a.ClanID = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryClan queries the "clan" edge of the Account entity.
func (a *Account) QueryClan() *ClanQuery {
	return NewAccountClient(a.config).QueryClan(a)
}

// QuerySnapshots queries the "snapshots" edge of the Account entity.
func (a *Account) QuerySnapshots() *AccountSnapshotQuery {
	return NewAccountClient(a.config).QuerySnapshots(a)
}

// QueryVehicleSnapshots queries the "vehicle_snapshots" edge of the Account entity.
func (a *Account) QueryVehicleSnapshots() *VehicleSnapshotQuery {
	return NewAccountClient(a.config).QueryVehicleSnapshots(a)
}

// QueryAchievementSnapshots queries the "achievement_snapshots" edge of the Account entity.
func (a *Account) QueryAchievementSnapshots() *AchievementsSnapshotQuery {
	return NewAccountClient(a.config).QueryAchievementSnapshots(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("db: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", a.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("last_battle_time=")
	builder.WriteString(fmt.Sprintf("%v", a.LastBattleTime))
	builder.WriteString(", ")
	builder.WriteString("account_created_at=")
	builder.WriteString(fmt.Sprintf("%v", a.AccountCreatedAt))
	builder.WriteString(", ")
	builder.WriteString("realm=")
	builder.WriteString(a.Realm)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(a.Nickname)
	builder.WriteString(", ")
	builder.WriteString("private=")
	builder.WriteString(fmt.Sprintf("%v", a.Private))
	builder.WriteString(", ")
	builder.WriteString("clan_id=")
	builder.WriteString(a.ClanID)
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account