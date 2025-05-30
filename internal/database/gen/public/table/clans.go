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

var Clans = newClansTable("public", "clans", "")

type clansTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	CreatedAt postgres.ColumnTimestampz
	UpdatedAt postgres.ColumnTimestampz
	Tag       postgres.ColumnString
	Name      postgres.ColumnString
	EmblemID  postgres.ColumnString
	Members   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
	DefaultColumns postgres.ColumnList
}

type ClansTable struct {
	clansTable

	EXCLUDED clansTable
}

// AS creates new ClansTable with assigned alias
func (a ClansTable) AS(alias string) *ClansTable {
	return newClansTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ClansTable with assigned schema name
func (a ClansTable) FromSchema(schemaName string) *ClansTable {
	return newClansTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ClansTable with assigned table prefix
func (a ClansTable) WithPrefix(prefix string) *ClansTable {
	return newClansTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ClansTable with assigned table suffix
func (a ClansTable) WithSuffix(suffix string) *ClansTable {
	return newClansTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newClansTable(schemaName, tableName, alias string) *ClansTable {
	return &ClansTable{
		clansTable: newClansTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newClansTableImpl("", "excluded", ""),
	}
}

func newClansTableImpl(schemaName, tableName, alias string) clansTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn = postgres.TimestampzColumn("updated_at")
		TagColumn       = postgres.StringColumn("tag")
		NameColumn      = postgres.StringColumn("name")
		EmblemIDColumn  = postgres.StringColumn("emblem_id")
		MembersColumn   = postgres.StringColumn("members")
		allColumns      = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, TagColumn, NameColumn, EmblemIDColumn, MembersColumn}
		mutableColumns  = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, TagColumn, NameColumn, EmblemIDColumn, MembersColumn}
		defaultColumns  = postgres.ColumnList{EmblemIDColumn}
	)

	return clansTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		Tag:       TagColumn,
		Name:      NameColumn,
		EmblemID:  EmblemIDColumn,
		Members:   MembersColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
