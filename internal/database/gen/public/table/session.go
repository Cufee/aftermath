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

var Session = newSessionTable("public", "session", "")

type sessionTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	CreatedAt postgres.ColumnString
	UpdatedAt postgres.ColumnString
	ExpiresAt postgres.ColumnString
	PublicID  postgres.ColumnString
	Metadata  postgres.ColumnBytea
	UserID    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
	DefaultColumns postgres.ColumnList
}

type SessionTable struct {
	sessionTable

	EXCLUDED sessionTable
}

// AS creates new SessionTable with assigned alias
func (a SessionTable) AS(alias string) *SessionTable {
	return newSessionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SessionTable with assigned schema name
func (a SessionTable) FromSchema(schemaName string) *SessionTable {
	return newSessionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SessionTable with assigned table prefix
func (a SessionTable) WithPrefix(prefix string) *SessionTable {
	return newSessionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SessionTable with assigned table suffix
func (a SessionTable) WithSuffix(suffix string) *SessionTable {
	return newSessionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSessionTable(schemaName, tableName, alias string) *SessionTable {
	return &SessionTable{
		sessionTable: newSessionTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newSessionTableImpl("", "excluded", ""),
	}
}

func newSessionTableImpl(schemaName, tableName, alias string) sessionTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		CreatedAtColumn = postgres.StringColumn("created_at")
		UpdatedAtColumn = postgres.StringColumn("updated_at")
		ExpiresAtColumn = postgres.StringColumn("expires_at")
		PublicIDColumn  = postgres.StringColumn("public_id")
		MetadataColumn  = postgres.ByteaColumn("metadata")
		UserIDColumn    = postgres.StringColumn("user_id")
		allColumns      = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, ExpiresAtColumn, PublicIDColumn, MetadataColumn, UserIDColumn}
		mutableColumns  = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, ExpiresAtColumn, PublicIDColumn, MetadataColumn, UserIDColumn}
		defaultColumns  = postgres.ColumnList{MetadataColumn}
	)

	return sessionTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		ExpiresAt: ExpiresAtColumn,
		PublicID:  PublicIDColumn,
		Metadata:  MetadataColumn,
		UserID:    UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
