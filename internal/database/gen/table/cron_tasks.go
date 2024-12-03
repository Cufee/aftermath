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

var CronTasks = newCronTasksTable("", "cron_tasks", "")

type cronTasksTable struct {
	sqlite.Table

	// Columns
	ID             sqlite.ColumnString
	CreatedAt      sqlite.ColumnTimestamp
	UpdatedAt      sqlite.ColumnTimestamp
	Type           sqlite.ColumnString
	ReferenceID    sqlite.ColumnString
	Targets        sqlite.ColumnString
	Status         sqlite.ColumnString
	ScheduledAfter sqlite.ColumnTimestamp
	LastRun        sqlite.ColumnTimestamp
	TriesLeft      sqlite.ColumnInteger
	Logs           sqlite.ColumnString
	Data           sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type CronTasksTable struct {
	cronTasksTable

	EXCLUDED cronTasksTable
}

// AS creates new CronTasksTable with assigned alias
func (a CronTasksTable) AS(alias string) *CronTasksTable {
	return newCronTasksTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CronTasksTable with assigned schema name
func (a CronTasksTable) FromSchema(schemaName string) *CronTasksTable {
	return newCronTasksTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CronTasksTable with assigned table prefix
func (a CronTasksTable) WithPrefix(prefix string) *CronTasksTable {
	return newCronTasksTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CronTasksTable with assigned table suffix
func (a CronTasksTable) WithSuffix(suffix string) *CronTasksTable {
	return newCronTasksTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCronTasksTable(schemaName, tableName, alias string) *CronTasksTable {
	return &CronTasksTable{
		cronTasksTable: newCronTasksTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newCronTasksTableImpl("", "excluded", ""),
	}
}

func newCronTasksTableImpl(schemaName, tableName, alias string) cronTasksTable {
	var (
		IDColumn             = sqlite.StringColumn("id")
		CreatedAtColumn      = sqlite.TimestampColumn("created_at")
		UpdatedAtColumn      = sqlite.TimestampColumn("updated_at")
		TypeColumn           = sqlite.StringColumn("type")
		ReferenceIDColumn    = sqlite.StringColumn("reference_id")
		TargetsColumn        = sqlite.StringColumn("targets")
		StatusColumn         = sqlite.StringColumn("status")
		ScheduledAfterColumn = sqlite.TimestampColumn("scheduled_after")
		LastRunColumn        = sqlite.TimestampColumn("last_run")
		TriesLeftColumn      = sqlite.IntegerColumn("tries_left")
		LogsColumn           = sqlite.StringColumn("logs")
		DataColumn           = sqlite.StringColumn("data")
		allColumns           = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, TypeColumn, ReferenceIDColumn, TargetsColumn, StatusColumn, ScheduledAfterColumn, LastRunColumn, TriesLeftColumn, LogsColumn, DataColumn}
		mutableColumns       = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, TypeColumn, ReferenceIDColumn, TargetsColumn, StatusColumn, ScheduledAfterColumn, LastRunColumn, TriesLeftColumn, LogsColumn, DataColumn}
	)

	return cronTasksTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		Type:           TypeColumn,
		ReferenceID:    ReferenceIDColumn,
		Targets:        TargetsColumn,
		Status:         StatusColumn,
		ScheduledAfter: ScheduledAfterColumn,
		LastRun:        LastRunColumn,
		TriesLeft:      TriesLeftColumn,
		Logs:           LogsColumn,
		Data:           DataColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
