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

var ApplicationCommand = newApplicationCommandTable("", "application_command", "")

type applicationCommandTable struct {
	sqlite.Table

	// Columns
	ID          sqlite.ColumnString
	CreatedAt   sqlite.ColumnInteger
	UpdatedAt   sqlite.ColumnInteger
	Name        sqlite.ColumnString
	Version     sqlite.ColumnString
	OptionsHash sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type ApplicationCommandTable struct {
	applicationCommandTable

	EXCLUDED applicationCommandTable
}

// AS creates new ApplicationCommandTable with assigned alias
func (a ApplicationCommandTable) AS(alias string) *ApplicationCommandTable {
	return newApplicationCommandTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ApplicationCommandTable with assigned schema name
func (a ApplicationCommandTable) FromSchema(schemaName string) *ApplicationCommandTable {
	return newApplicationCommandTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ApplicationCommandTable with assigned table prefix
func (a ApplicationCommandTable) WithPrefix(prefix string) *ApplicationCommandTable {
	return newApplicationCommandTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ApplicationCommandTable with assigned table suffix
func (a ApplicationCommandTable) WithSuffix(suffix string) *ApplicationCommandTable {
	return newApplicationCommandTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newApplicationCommandTable(schemaName, tableName, alias string) *ApplicationCommandTable {
	return &ApplicationCommandTable{
		applicationCommandTable: newApplicationCommandTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newApplicationCommandTableImpl("", "excluded", ""),
	}
}

func newApplicationCommandTableImpl(schemaName, tableName, alias string) applicationCommandTable {
	var (
		IDColumn          = sqlite.StringColumn("id")
		CreatedAtColumn   = sqlite.IntegerColumn("created_at")
		UpdatedAtColumn   = sqlite.IntegerColumn("updated_at")
		NameColumn        = sqlite.StringColumn("name")
		VersionColumn     = sqlite.StringColumn("version")
		OptionsHashColumn = sqlite.StringColumn("options_hash")
		allColumns        = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, NameColumn, VersionColumn, OptionsHashColumn}
		mutableColumns    = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, NameColumn, VersionColumn, OptionsHashColumn}
	)

	return applicationCommandTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		Name:        NameColumn,
		Version:     VersionColumn,
		OptionsHash: OptionsHashColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
