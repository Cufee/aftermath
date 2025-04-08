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

var UserContents = newUserContentsTable("public", "user_contents", "")

type userContentsTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnString
	CreatedAt   postgres.ColumnTimestampz
	UpdatedAt   postgres.ColumnTimestampz
	Type        postgres.ColumnString
	ReferenceID postgres.ColumnString
	Value       postgres.ColumnString
	Metadata    postgres.ColumnString
	UserID      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
	DefaultColumns postgres.ColumnList
}

type UserContentsTable struct {
	userContentsTable

	EXCLUDED userContentsTable
}

// AS creates new UserContentsTable with assigned alias
func (a UserContentsTable) AS(alias string) *UserContentsTable {
	return newUserContentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserContentsTable with assigned schema name
func (a UserContentsTable) FromSchema(schemaName string) *UserContentsTable {
	return newUserContentsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserContentsTable with assigned table prefix
func (a UserContentsTable) WithPrefix(prefix string) *UserContentsTable {
	return newUserContentsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserContentsTable with assigned table suffix
func (a UserContentsTable) WithSuffix(suffix string) *UserContentsTable {
	return newUserContentsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserContentsTable(schemaName, tableName, alias string) *UserContentsTable {
	return &UserContentsTable{
		userContentsTable: newUserContentsTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newUserContentsTableImpl("", "excluded", ""),
	}
}

func newUserContentsTableImpl(schemaName, tableName, alias string) userContentsTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		CreatedAtColumn   = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn   = postgres.TimestampzColumn("updated_at")
		TypeColumn        = postgres.StringColumn("type")
		ReferenceIDColumn = postgres.StringColumn("reference_id")
		ValueColumn       = postgres.StringColumn("value")
		MetadataColumn    = postgres.StringColumn("metadata")
		UserIDColumn      = postgres.StringColumn("user_id")
		allColumns        = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, TypeColumn, ReferenceIDColumn, ValueColumn, MetadataColumn, UserIDColumn}
		mutableColumns    = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, TypeColumn, ReferenceIDColumn, ValueColumn, MetadataColumn, UserIDColumn}
		defaultColumns    = postgres.ColumnList{}
	)

	return userContentsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		Type:        TypeColumn,
		ReferenceID: ReferenceIDColumn,
		Value:       ValueColumn,
		Metadata:    MetadataColumn,
		UserID:      UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
