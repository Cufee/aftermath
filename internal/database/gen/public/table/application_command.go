//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var ApplicationCommand = newApplicationCommandTable("public", "application_command", "")

type applicationCommandTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnString
	CreatedAt   postgres.ColumnString
	UpdatedAt   postgres.ColumnString
	Name        postgres.ColumnString
	Version     postgres.ColumnString
	OptionsHash postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
	DefaultColumns postgres.ColumnList
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
		IDColumn          = postgres.StringColumn("id")
		CreatedAtColumn   = postgres.StringColumn("created_at")
		UpdatedAtColumn   = postgres.StringColumn("updated_at")
		NameColumn        = postgres.StringColumn("name")
		VersionColumn     = postgres.StringColumn("version")
		OptionsHashColumn = postgres.StringColumn("options_hash")
		allColumns        = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, NameColumn, VersionColumn, OptionsHashColumn}
		mutableColumns    = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, NameColumn, VersionColumn, OptionsHashColumn}
		defaultColumns    = postgres.ColumnList{}
	)

	return applicationCommandTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		Name:        NameColumn,
		Version:     VersionColumn,
		OptionsHash: OptionsHashColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
