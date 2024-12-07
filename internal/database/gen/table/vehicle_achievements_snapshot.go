//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/sqlite"
)

var VehicleAchievementsSnapshot = newVehicleAchievementsSnapshotTable("", "vehicle_achievements_snapshot", "")

type vehicleAchievementsSnapshotTable struct {
	sqlite.Table

	// Columns
	ID             sqlite.ColumnString
	CreatedAt      sqlite.ColumnString
	Type           sqlite.ColumnString
	VehicleID      sqlite.ColumnString
	ReferenceID    sqlite.ColumnString
	Battles        sqlite.ColumnInteger
	LastBattleTime sqlite.ColumnString
	Frame          sqlite.ColumnString
	AccountID      sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type VehicleAchievementsSnapshotTable struct {
	vehicleAchievementsSnapshotTable

	EXCLUDED vehicleAchievementsSnapshotTable
}

// AS creates new VehicleAchievementsSnapshotTable with assigned alias
func (a VehicleAchievementsSnapshotTable) AS(alias string) *VehicleAchievementsSnapshotTable {
	return newVehicleAchievementsSnapshotTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new VehicleAchievementsSnapshotTable with assigned schema name
func (a VehicleAchievementsSnapshotTable) FromSchema(schemaName string) *VehicleAchievementsSnapshotTable {
	return newVehicleAchievementsSnapshotTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new VehicleAchievementsSnapshotTable with assigned table prefix
func (a VehicleAchievementsSnapshotTable) WithPrefix(prefix string) *VehicleAchievementsSnapshotTable {
	return newVehicleAchievementsSnapshotTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new VehicleAchievementsSnapshotTable with assigned table suffix
func (a VehicleAchievementsSnapshotTable) WithSuffix(suffix string) *VehicleAchievementsSnapshotTable {
	return newVehicleAchievementsSnapshotTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newVehicleAchievementsSnapshotTable(schemaName, tableName, alias string) *VehicleAchievementsSnapshotTable {
	return &VehicleAchievementsSnapshotTable{
		vehicleAchievementsSnapshotTable: newVehicleAchievementsSnapshotTableImpl(schemaName, tableName, alias),
		EXCLUDED:                         newVehicleAchievementsSnapshotTableImpl("", "excluded", ""),
	}
}

func newVehicleAchievementsSnapshotTableImpl(schemaName, tableName, alias string) vehicleAchievementsSnapshotTable {
	var (
		IDColumn             = sqlite.StringColumn("id")
		CreatedAtColumn      = sqlite.StringColumn("created_at")
		TypeColumn           = sqlite.StringColumn("type")
		VehicleIDColumn      = sqlite.StringColumn("vehicle_id")
		ReferenceIDColumn    = sqlite.StringColumn("reference_id")
		BattlesColumn        = sqlite.IntegerColumn("battles")
		LastBattleTimeColumn = sqlite.StringColumn("last_battle_time")
		FrameColumn          = sqlite.StringColumn("frame")
		AccountIDColumn      = sqlite.StringColumn("account_id")
		allColumns           = sqlite.ColumnList{IDColumn, CreatedAtColumn, TypeColumn, VehicleIDColumn, ReferenceIDColumn, BattlesColumn, LastBattleTimeColumn, FrameColumn, AccountIDColumn}
		mutableColumns       = sqlite.ColumnList{CreatedAtColumn, TypeColumn, VehicleIDColumn, ReferenceIDColumn, BattlesColumn, LastBattleTimeColumn, FrameColumn, AccountIDColumn}
	)

	return vehicleAchievementsSnapshotTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		Type:           TypeColumn,
		VehicleID:      VehicleIDColumn,
		ReferenceID:    ReferenceIDColumn,
		Battles:        BattlesColumn,
		LastBattleTime: LastBattleTimeColumn,
		Frame:          FrameColumn,
		AccountID:      AccountIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}