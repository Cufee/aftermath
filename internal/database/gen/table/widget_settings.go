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

var WidgetSettings = newWidgetSettingsTable("", "widget_settings", "")

type widgetSettingsTable struct {
	sqlite.Table

	// Columns
	ID                 sqlite.ColumnString
	CreatedAt          sqlite.ColumnTimestamp
	UpdatedAt          sqlite.ColumnTimestamp
	ReferenceID        sqlite.ColumnString
	Title              sqlite.ColumnString
	SessionFrom        sqlite.ColumnTimestamp
	Metadata           sqlite.ColumnString
	Styles             sqlite.ColumnString
	UserID             sqlite.ColumnString
	SessionReferenceID sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type WidgetSettingsTable struct {
	widgetSettingsTable

	EXCLUDED widgetSettingsTable
}

// AS creates new WidgetSettingsTable with assigned alias
func (a WidgetSettingsTable) AS(alias string) *WidgetSettingsTable {
	return newWidgetSettingsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new WidgetSettingsTable with assigned schema name
func (a WidgetSettingsTable) FromSchema(schemaName string) *WidgetSettingsTable {
	return newWidgetSettingsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new WidgetSettingsTable with assigned table prefix
func (a WidgetSettingsTable) WithPrefix(prefix string) *WidgetSettingsTable {
	return newWidgetSettingsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new WidgetSettingsTable with assigned table suffix
func (a WidgetSettingsTable) WithSuffix(suffix string) *WidgetSettingsTable {
	return newWidgetSettingsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newWidgetSettingsTable(schemaName, tableName, alias string) *WidgetSettingsTable {
	return &WidgetSettingsTable{
		widgetSettingsTable: newWidgetSettingsTableImpl(schemaName, tableName, alias),
		EXCLUDED:            newWidgetSettingsTableImpl("", "excluded", ""),
	}
}

func newWidgetSettingsTableImpl(schemaName, tableName, alias string) widgetSettingsTable {
	var (
		IDColumn                 = sqlite.StringColumn("id")
		CreatedAtColumn          = sqlite.TimestampColumn("created_at")
		UpdatedAtColumn          = sqlite.TimestampColumn("updated_at")
		ReferenceIDColumn        = sqlite.StringColumn("reference_id")
		TitleColumn              = sqlite.StringColumn("title")
		SessionFromColumn        = sqlite.TimestampColumn("session_from")
		MetadataColumn           = sqlite.StringColumn("metadata")
		StylesColumn             = sqlite.StringColumn("styles")
		UserIDColumn             = sqlite.StringColumn("user_id")
		SessionReferenceIDColumn = sqlite.StringColumn("session_reference_id")
		allColumns               = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, ReferenceIDColumn, TitleColumn, SessionFromColumn, MetadataColumn, StylesColumn, UserIDColumn, SessionReferenceIDColumn}
		mutableColumns           = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, ReferenceIDColumn, TitleColumn, SessionFromColumn, MetadataColumn, StylesColumn, UserIDColumn, SessionReferenceIDColumn}
	)

	return widgetSettingsTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                 IDColumn,
		CreatedAt:          CreatedAtColumn,
		UpdatedAt:          UpdatedAtColumn,
		ReferenceID:        ReferenceIDColumn,
		Title:              TitleColumn,
		SessionFrom:        SessionFromColumn,
		Metadata:           MetadataColumn,
		Styles:             StylesColumn,
		UserID:             UserIDColumn,
		SessionReferenceID: SessionReferenceIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}