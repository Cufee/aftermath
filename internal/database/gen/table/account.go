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

var Account = newAccountTable("", "account", "")

type accountTable struct {
	sqlite.Table

	// Columns
	ID               sqlite.ColumnString
	CreatedAt        sqlite.ColumnTimestamp
	UpdatedAt        sqlite.ColumnTimestamp
	LastBattleTime   sqlite.ColumnTimestamp
	AccountCreatedAt sqlite.ColumnTimestamp
	Realm            sqlite.ColumnString
	Nickname         sqlite.ColumnString
	Private          sqlite.ColumnBool
	ClanID           sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type AccountTable struct {
	accountTable

	EXCLUDED accountTable
}

// AS creates new AccountTable with assigned alias
func (a AccountTable) AS(alias string) *AccountTable {
	return newAccountTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AccountTable with assigned schema name
func (a AccountTable) FromSchema(schemaName string) *AccountTable {
	return newAccountTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AccountTable with assigned table prefix
func (a AccountTable) WithPrefix(prefix string) *AccountTable {
	return newAccountTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AccountTable with assigned table suffix
func (a AccountTable) WithSuffix(suffix string) *AccountTable {
	return newAccountTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAccountTable(schemaName, tableName, alias string) *AccountTable {
	return &AccountTable{
		accountTable: newAccountTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newAccountTableImpl("", "excluded", ""),
	}
}

func newAccountTableImpl(schemaName, tableName, alias string) accountTable {
	var (
		IDColumn               = sqlite.StringColumn("id")
		CreatedAtColumn        = sqlite.TimestampColumn("created_at")
		UpdatedAtColumn        = sqlite.TimestampColumn("updated_at")
		LastBattleTimeColumn   = sqlite.TimestampColumn("last_battle_time")
		AccountCreatedAtColumn = sqlite.TimestampColumn("account_created_at")
		RealmColumn            = sqlite.StringColumn("realm")
		NicknameColumn         = sqlite.StringColumn("nickname")
		PrivateColumn          = sqlite.BoolColumn("private")
		ClanIDColumn           = sqlite.StringColumn("clan_id")
		allColumns             = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, LastBattleTimeColumn, AccountCreatedAtColumn, RealmColumn, NicknameColumn, PrivateColumn, ClanIDColumn}
		mutableColumns         = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, LastBattleTimeColumn, AccountCreatedAtColumn, RealmColumn, NicknameColumn, PrivateColumn, ClanIDColumn}
	)

	return accountTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		CreatedAt:        CreatedAtColumn,
		UpdatedAt:        UpdatedAtColumn,
		LastBattleTime:   LastBattleTimeColumn,
		AccountCreatedAt: AccountCreatedAtColumn,
		Realm:            RealmColumn,
		Nickname:         NicknameColumn,
		Private:          PrivateColumn,
		ClanID:           ClanIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
