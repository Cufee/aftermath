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

var GameMode = newGameModeTable("", "game_mode", "")

type gameModeTable struct {
	sqlite.Table

	// Columns
	ID             sqlite.ColumnString
	CreatedAt      sqlite.ColumnTimestamp
	UpdatedAt      sqlite.ColumnTimestamp
	LocalizedNames sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type GameModeTable struct {
	gameModeTable

	EXCLUDED gameModeTable
}

// AS creates new GameModeTable with assigned alias
func (a GameModeTable) AS(alias string) *GameModeTable {
	return newGameModeTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GameModeTable with assigned schema name
func (a GameModeTable) FromSchema(schemaName string) *GameModeTable {
	return newGameModeTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GameModeTable with assigned table prefix
func (a GameModeTable) WithPrefix(prefix string) *GameModeTable {
	return newGameModeTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GameModeTable with assigned table suffix
func (a GameModeTable) WithSuffix(suffix string) *GameModeTable {
	return newGameModeTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGameModeTable(schemaName, tableName, alias string) *GameModeTable {
	return &GameModeTable{
		gameModeTable: newGameModeTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newGameModeTableImpl("", "excluded", ""),
	}
}

func newGameModeTableImpl(schemaName, tableName, alias string) gameModeTable {
	var (
		IDColumn             = sqlite.StringColumn("id")
		CreatedAtColumn      = sqlite.TimestampColumn("created_at")
		UpdatedAtColumn      = sqlite.TimestampColumn("updated_at")
		LocalizedNamesColumn = sqlite.StringColumn("localized_names")
		allColumns           = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, LocalizedNamesColumn}
		mutableColumns       = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, LocalizedNamesColumn}
	)

	return gameModeTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:             IDColumn,
		CreatedAt:      CreatedAtColumn,
		UpdatedAt:      UpdatedAtColumn,
		LocalizedNames: LocalizedNamesColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
