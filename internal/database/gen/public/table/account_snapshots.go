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

var AccountSnapshots = newAccountSnapshotsTable("public", "account_snapshots", "")

type accountSnapshotsTable struct {
	postgres.Table

	// Columns
	ID             postgres.ColumnString
	CreatedAt      postgres.ColumnTimestampz
	UpdatedAt      postgres.ColumnTimestampz
	Type           postgres.ColumnString
	LastBattleTime postgres.ColumnTimestampz
	ReferenceID    postgres.ColumnString
	RatingBattles  postgres.ColumnInteger
	RatingFrame    postgres.ColumnString
	RegularBattles postgres.ColumnInteger
	RegularFrame   postgres.ColumnString
	AccountID      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
	DefaultColumns postgres.ColumnList
}

type AccountSnapshotsTable struct {
	accountSnapshotsTable

	EXCLUDED accountSnapshotsTable
}

// AS creates new AccountSnapshotsTable with assigned alias
func (a AccountSnapshotsTable) AS(alias string) *AccountSnapshotsTable {
	return newAccountSnapshotsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AccountSnapshotsTable with assigned schema name
func (a AccountSnapshotsTable) FromSchema(schemaName string) *AccountSnapshotsTable {
	return newAccountSnapshotsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AccountSnapshotsTable with assigned table prefix
func (a AccountSnapshotsTable) WithPrefix(prefix string) *AccountSnapshotsTable {
	return newAccountSnapshotsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AccountSnapshotsTable with assigned table suffix
func (a AccountSnapshotsTable) WithSuffix(suffix string) *AccountSnapshotsTable {
	return newAccountSnapshotsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAccountSnapshotsTable(schemaName, tableName, alias string) *AccountSnapshotsTable {
	return &AccountSnapshotsTable{
		accountSnapshotsTable: newAccountSnapshotsTableImpl(schemaName, tableName, alias),
		EXCLUDED:              newAccountSnapshotsTableImpl("", "excluded", ""),
	}
}

func newAccountSnapshotsTableImpl(schemaName, tableName, alias string) accountSnapshotsTable {
	var (
		IDColumn             = postgres.StringColumn("id")
		CreatedAtColumn      = postgres.TimestampzColumn("created_at")
		UpdatedAtColumn      = postgres.TimestampzColumn("updated_at")
		TypeColumn           = postgres.StringColumn("type")
		LastBattleTimeColumn = postgres.TimestampzColumn("last_battle_time")
		ReferenceIDColumn    = postgres.StringColumn("reference_id")
		RatingBattlesColumn  = postgres.IntegerColumn("rating_battles")
		RatingFrameColumn    = postgres.StringColumn("rating_frame")
		RegularBattlesColumn = postgres.IntegerColumn("regular_battles")
		RegularFrameColumn   = postgres.StringColumn("regular_frame")
		AccountIDColumn      = postgres.StringColumn("account_id")
		allColumns           = postgres.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, TypeColumn, LastBattleTimeColumn, ReferenceIDColumn, RatingBattlesColumn, RatingFrameColumn, RegularBattlesColumn, RegularFrameColumn, AccountIDColumn}
		mutableColumns       = postgres.ColumnList{CreatedAtColumn, UpdatedAtColumn, TypeColumn, LastBattleTimeColumn, ReferenceIDColumn, RatingBattlesColumn, RatingFrameColumn, RegularBattlesColumn, RegularFrameColumn, AccountIDColumn}
		defaultColumns       = postgres.ColumnList{}
	)

	return accountSnapshotsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		Type:           TypeColumn,
		LastBattleTime: LastBattleTimeColumn,
		ReferenceID:    ReferenceIDColumn,
		RatingBattles:  RatingBattlesColumn,
		RatingFrame:    RatingFrameColumn,
		RegularBattles: RegularBattlesColumn,
		RegularFrame:   RegularFrameColumn,
		AccountID:      AccountIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
		DefaultColumns: defaultColumns,
	}
}
