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

var Clans = newClansTable("", "clans", "")

type clansTable struct {
	sqlite.Table

	// Columns
	ID        sqlite.ColumnString
	CreatedAt sqlite.ColumnTimestamp
	UpdatedAt sqlite.ColumnTimestamp
	Tag       sqlite.ColumnString
	Name      sqlite.ColumnString
	EmblemID  sqlite.ColumnString
	Members   sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
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
		IDColumn        = sqlite.StringColumn("id")
		CreatedAtColumn = sqlite.TimestampColumn("created_at")
		UpdatedAtColumn = sqlite.TimestampColumn("updated_at")
		TagColumn       = sqlite.StringColumn("tag")
		NameColumn      = sqlite.StringColumn("name")
		EmblemIDColumn  = sqlite.StringColumn("emblem_id")
		MembersColumn   = sqlite.StringColumn("members")
		allColumns      = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, TagColumn, NameColumn, EmblemIDColumn, MembersColumn}
		mutableColumns  = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, TagColumn, NameColumn, EmblemIDColumn, MembersColumn}
	)

	return clansTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

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
	}
}
