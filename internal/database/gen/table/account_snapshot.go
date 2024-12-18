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

var AccountSnapshot = newAccountSnapshotTable("", "account_snapshot", "")

type accountSnapshotTable struct {
	sqlite.Table

	// Columns
	ID             sqlite.ColumnString
	CreatedAt      sqlite.ColumnString
	Type           sqlite.ColumnString
	LastBattleTime sqlite.ColumnString
	ReferenceID    sqlite.ColumnString
	RatingBattles  sqlite.ColumnInteger
	RatingFrame    sqlite.ColumnString
	RegularBattles sqlite.ColumnInteger
	RegularFrame   sqlite.ColumnString
	AccountID      sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type AccountSnapshotTable struct {
	accountSnapshotTable

	EXCLUDED accountSnapshotTable
}

// AS creates new AccountSnapshotTable with assigned alias
func (a AccountSnapshotTable) AS(alias string) *AccountSnapshotTable {
	return newAccountSnapshotTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AccountSnapshotTable with assigned schema name
func (a AccountSnapshotTable) FromSchema(schemaName string) *AccountSnapshotTable {
	return newAccountSnapshotTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AccountSnapshotTable with assigned table prefix
func (a AccountSnapshotTable) WithPrefix(prefix string) *AccountSnapshotTable {
	return newAccountSnapshotTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AccountSnapshotTable with assigned table suffix
func (a AccountSnapshotTable) WithSuffix(suffix string) *AccountSnapshotTable {
	return newAccountSnapshotTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAccountSnapshotTable(schemaName, tableName, alias string) *AccountSnapshotTable {
	return &AccountSnapshotTable{
		accountSnapshotTable: newAccountSnapshotTableImpl(schemaName, tableName, alias),
		EXCLUDED:             newAccountSnapshotTableImpl("", "excluded", ""),
	}
}

func newAccountSnapshotTableImpl(schemaName, tableName, alias string) accountSnapshotTable {
	var (
		IDColumn             = sqlite.StringColumn("id")
		CreatedAtColumn      = sqlite.StringColumn("created_at")
		TypeColumn           = sqlite.StringColumn("type")
		LastBattleTimeColumn = sqlite.StringColumn("last_battle_time")
		ReferenceIDColumn    = sqlite.StringColumn("reference_id")
		RatingBattlesColumn  = sqlite.IntegerColumn("rating_battles")
		RatingFrameColumn    = sqlite.StringColumn("rating_frame")
		RegularBattlesColumn = sqlite.IntegerColumn("regular_battles")
		RegularFrameColumn   = sqlite.StringColumn("regular_frame")
		AccountIDColumn      = sqlite.StringColumn("account_id")
		allColumns           = sqlite.ColumnList{IDColumn, CreatedAtColumn, TypeColumn, LastBattleTimeColumn, ReferenceIDColumn, RatingBattlesColumn, RatingFrameColumn, RegularBattlesColumn, RegularFrameColumn, AccountIDColumn}
		mutableColumns       = sqlite.ColumnList{CreatedAtColumn, TypeColumn, LastBattleTimeColumn, ReferenceIDColumn, RatingBattlesColumn, RatingFrameColumn, RegularBattlesColumn, RegularFrameColumn, AccountIDColumn}
	)

	return accountSnapshotTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
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
	}
}
