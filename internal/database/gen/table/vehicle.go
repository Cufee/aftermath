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

var Vehicle = newVehicleTable("", "vehicle", "")

type vehicleTable struct {
	sqlite.Table

	// Columns
	ID             sqlite.ColumnString
	CreatedAt      sqlite.ColumnTimestamp
	UpdatedAt      sqlite.ColumnTimestamp
	Tier           sqlite.ColumnInteger
	LocalizedNames sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type VehicleTable struct {
	vehicleTable

	EXCLUDED vehicleTable
}

// AS creates new VehicleTable with assigned alias
func (a VehicleTable) AS(alias string) *VehicleTable {
	return newVehicleTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new VehicleTable with assigned schema name
func (a VehicleTable) FromSchema(schemaName string) *VehicleTable {
	return newVehicleTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new VehicleTable with assigned table prefix
func (a VehicleTable) WithPrefix(prefix string) *VehicleTable {
	return newVehicleTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new VehicleTable with assigned table suffix
func (a VehicleTable) WithSuffix(suffix string) *VehicleTable {
	return newVehicleTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newVehicleTable(schemaName, tableName, alias string) *VehicleTable {
	return &VehicleTable{
		vehicleTable: newVehicleTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newVehicleTableImpl("", "excluded", ""),
	}
}

func newVehicleTableImpl(schemaName, tableName, alias string) vehicleTable {
	var (
		IDColumn             = sqlite.StringColumn("id")
		CreatedAtColumn      = sqlite.TimestampColumn("created_at")
		UpdatedAtColumn      = sqlite.TimestampColumn("updated_at")
		TierColumn           = sqlite.IntegerColumn("tier")
		LocalizedNamesColumn = sqlite.StringColumn("localized_names")
		allColumns           = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, TierColumn, LocalizedNamesColumn}
		mutableColumns       = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, TierColumn, LocalizedNamesColumn}
	)

	return vehicleTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		Tier:           TierColumn,
		LocalizedNames: LocalizedNamesColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}