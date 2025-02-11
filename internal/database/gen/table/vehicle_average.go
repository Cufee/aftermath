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

var VehicleAverage = newVehicleAverageTable("", "vehicle_average", "")

type vehicleAverageTable struct {
	sqlite.Table

	// Columns
	ID        sqlite.ColumnString
	CreatedAt sqlite.ColumnString
	UpdatedAt sqlite.ColumnString
	Data      sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type VehicleAverageTable struct {
	vehicleAverageTable

	EXCLUDED vehicleAverageTable
}

// AS creates new VehicleAverageTable with assigned alias
func (a VehicleAverageTable) AS(alias string) *VehicleAverageTable {
	return newVehicleAverageTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new VehicleAverageTable with assigned schema name
func (a VehicleAverageTable) FromSchema(schemaName string) *VehicleAverageTable {
	return newVehicleAverageTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new VehicleAverageTable with assigned table prefix
func (a VehicleAverageTable) WithPrefix(prefix string) *VehicleAverageTable {
	return newVehicleAverageTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new VehicleAverageTable with assigned table suffix
func (a VehicleAverageTable) WithSuffix(suffix string) *VehicleAverageTable {
	return newVehicleAverageTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newVehicleAverageTable(schemaName, tableName, alias string) *VehicleAverageTable {
	return &VehicleAverageTable{
		vehicleAverageTable: newVehicleAverageTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newVehicleAverageTableImpl("", "excluded", ""),
	}
}

func newVehicleAverageTableImpl(schemaName, tableName, alias string) vehicleAverageTable {
	var (
		IDColumn        = sqlite.StringColumn("id")
		CreatedAtColumn = sqlite.StringColumn("created_at")
		UpdatedAtColumn = sqlite.StringColumn("updated_at")
		DataColumn      = sqlite.StringColumn("data")
		allColumns      = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DataColumn}
		mutableColumns  = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, DataColumn}
	)

	return vehicleAverageTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		Data:      DataColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
