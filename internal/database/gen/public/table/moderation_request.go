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

var ModerationRequest = newModerationRequestTable("public", "moderation_request", "")

type moderationRequestTable struct {
	postgres.Table

	// Columns
	ID               postgres.ColumnString
	CreatedAt        postgres.ColumnString
	UpdatedAt        postgres.ColumnString
	ModeratorComment postgres.ColumnString
	Context          postgres.ColumnString
	ReferenceID      postgres.ColumnString
	ActionReason     postgres.ColumnString
	ActionStatus     postgres.ColumnString
	Data             postgres.ColumnBytea
	RequestorID      postgres.ColumnString
	ModeratorID      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
	DefaultColumns postgres.ColumnList
}

type ModerationRequestTable struct {
	moderationRequestTable

	EXCLUDED moderationRequestTable
}

// AS creates new ModerationRequestTable with assigned alias
func (a ModerationRequestTable) AS(alias string) *ModerationRequestTable {
	return newModerationRequestTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ModerationRequestTable with assigned schema name
func (a ModerationRequestTable) FromSchema(schemaName string) *ModerationRequestTable {
	return newModerationRequestTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ModerationRequestTable with assigned table prefix
func (a ModerationRequestTable) WithPrefix(prefix string) *ModerationRequestTable {
	return newModerationRequestTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ModerationRequestTable with assigned table suffix
func (a ModerationRequestTable) WithSuffix(suffix string) *ModerationRequestTable {
	return newModerationRequestTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newModerationRequestTable(schemaName, tableName, alias string) *ModerationRequestTable {
	return &ModerationRequestTable{
		moderationRequestTable: newModerationRequestTableImpl(schemaName, tableName, alias),
		EXCLUDED:               newModerationRequestTableImpl("", "excluded", ""),
	}
}

func newModerationRequestTableImpl(schemaName, tableName, alias string) moderationRequestTable {
	var (
		IDColumn               = postgres.StringColumn("id")
		CreatedAtColumn        = postgres.StringColumn("created_at")
		UpdatedAtColumn        = postgres.StringColumn("updated_at")
		ModeratorCommentColumn = postgres.StringColumn("moderator_comment")
		ContextColumn          = postgres.StringColumn("context")
		ReferenceIDColumn      = postgres.StringColumn("reference_id")
		ActionReasonColumn     = postgres.StringColumn("action_reason")
		ActionStatusColumn     = postgres.StringColumn("action_status")
		DataColumn             = postgres.ByteaColumn("data")
		RequestorIDColumn      = postgres.StringColumn("requestor_id")
		ModeratorIDColumn      = postgres.StringColumn("moderator_id")
		allColumns             = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, ModeratorCommentColumn, ContextColumn, ReferenceIDColumn, ActionReasonColumn, ActionStatusColumn, DataColumn, RequestorIDColumn, ModeratorIDColumn}
		mutableColumns         = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, ModeratorCommentColumn, ContextColumn, ReferenceIDColumn, ActionReasonColumn, ActionStatusColumn, DataColumn, RequestorIDColumn, ModeratorIDColumn}
		defaultColumns         = postgres.ColumnList{DataColumn}
	)

	return moderationRequestTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		CreatedAt:        CreatedAtColumn,
		UpdatedAt:        UpdatedAtColumn,
		ModeratorComment: ModeratorCommentColumn,
		Context:          ContextColumn,
		ReferenceID:      ReferenceIDColumn,
		ActionReason:     ActionReasonColumn,
		ActionStatus:     ActionStatusColumn,
		Data:             DataColumn,
		RequestorID:      RequestorIDColumn,
		ModeratorID:      ModeratorIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
