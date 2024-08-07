// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLastBattleTime holds the string denoting the last_battle_time field in the database.
	FieldLastBattleTime = "last_battle_time"
	// FieldAccountCreatedAt holds the string denoting the account_created_at field in the database.
	FieldAccountCreatedAt = "account_created_at"
	// FieldRealm holds the string denoting the realm field in the database.
	FieldRealm = "realm"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldPrivate holds the string denoting the private field in the database.
	FieldPrivate = "private"
	// FieldClanID holds the string denoting the clan_id field in the database.
	FieldClanID = "clan_id"
	// EdgeClan holds the string denoting the clan edge name in mutations.
	EdgeClan = "clan"
	// EdgeVehicleSnapshots holds the string denoting the vehicle_snapshots edge name in mutations.
	EdgeVehicleSnapshots = "vehicle_snapshots"
	// EdgeAccountSnapshots holds the string denoting the account_snapshots edge name in mutations.
	EdgeAccountSnapshots = "account_snapshots"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// ClanTable is the table that holds the clan relation/edge.
	ClanTable = "accounts"
	// ClanInverseTable is the table name for the Clan entity.
	// It exists in this package in order to avoid circular dependency with the "clan" package.
	ClanInverseTable = "clans"
	// ClanColumn is the table column denoting the clan relation/edge.
	ClanColumn = "clan_id"
	// VehicleSnapshotsTable is the table that holds the vehicle_snapshots relation/edge.
	VehicleSnapshotsTable = "vehicle_snapshots"
	// VehicleSnapshotsInverseTable is the table name for the VehicleSnapshot entity.
	// It exists in this package in order to avoid circular dependency with the "vehiclesnapshot" package.
	VehicleSnapshotsInverseTable = "vehicle_snapshots"
	// VehicleSnapshotsColumn is the table column denoting the vehicle_snapshots relation/edge.
	VehicleSnapshotsColumn = "account_id"
	// AccountSnapshotsTable is the table that holds the account_snapshots relation/edge.
	AccountSnapshotsTable = "account_snapshots"
	// AccountSnapshotsInverseTable is the table name for the AccountSnapshot entity.
	// It exists in this package in order to avoid circular dependency with the "accountsnapshot" package.
	AccountSnapshotsInverseTable = "account_snapshots"
	// AccountSnapshotsColumn is the table column denoting the account_snapshots relation/edge.
	AccountSnapshotsColumn = "account_id"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLastBattleTime,
	FieldAccountCreatedAt,
	FieldRealm,
	FieldNickname,
	FieldPrivate,
	FieldClanID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// RealmValidator is a validator for the "realm" field. It is called by the builders before save.
	RealmValidator func(string) error
	// NicknameValidator is a validator for the "nickname" field. It is called by the builders before save.
	NicknameValidator func(string) error
	// DefaultPrivate holds the default value on creation for the "private" field.
	DefaultPrivate bool
)

// OrderOption defines the ordering options for the Account queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByLastBattleTime orders the results by the last_battle_time field.
func ByLastBattleTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastBattleTime, opts...).ToFunc()
}

// ByAccountCreatedAt orders the results by the account_created_at field.
func ByAccountCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountCreatedAt, opts...).ToFunc()
}

// ByRealm orders the results by the realm field.
func ByRealm(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRealm, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// ByPrivate orders the results by the private field.
func ByPrivate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivate, opts...).ToFunc()
}

// ByClanID orders the results by the clan_id field.
func ByClanID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClanID, opts...).ToFunc()
}

// ByClanField orders the results by clan field.
func ByClanField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClanStep(), sql.OrderByField(field, opts...))
	}
}

// ByVehicleSnapshotsCount orders the results by vehicle_snapshots count.
func ByVehicleSnapshotsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVehicleSnapshotsStep(), opts...)
	}
}

// ByVehicleSnapshots orders the results by vehicle_snapshots terms.
func ByVehicleSnapshots(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVehicleSnapshotsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAccountSnapshotsCount orders the results by account_snapshots count.
func ByAccountSnapshotsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccountSnapshotsStep(), opts...)
	}
}

// ByAccountSnapshots orders the results by account_snapshots terms.
func ByAccountSnapshots(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountSnapshotsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newClanStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClanInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ClanTable, ClanColumn),
	)
}
func newVehicleSnapshotsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VehicleSnapshotsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VehicleSnapshotsTable, VehicleSnapshotsColumn),
	)
}
func newAccountSnapshotsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountSnapshotsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AccountSnapshotsTable, AccountSnapshotsColumn),
	)
}
