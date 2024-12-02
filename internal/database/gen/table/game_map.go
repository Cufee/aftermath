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

var GameMap = newGameMapTable("", "game_map", "")

type gameMapTable struct {
	sqlite.Table

	// Columns
	ID              sqlite.ColumnString
	CreatedAt       sqlite.ColumnTimestamp
	UpdatedAt       sqlite.ColumnTimestamp
	GameModes       sqlite.ColumnString
	SupremacyPoints sqlite.ColumnInteger
	LocalizedNames  sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type GameMapTable struct {
	gameMapTable

	EXCLUDED gameMapTable
}

// AS creates new GameMapTable with assigned alias
func (a GameMapTable) AS(alias string) *GameMapTable {
	return newGameMapTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new GameMapTable with assigned schema name
func (a GameMapTable) FromSchema(schemaName string) *GameMapTable {
	return newGameMapTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new GameMapTable with assigned table prefix
func (a GameMapTable) WithPrefix(prefix string) *GameMapTable {
	return newGameMapTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new GameMapTable with assigned table suffix
func (a GameMapTable) WithSuffix(suffix string) *GameMapTable {
	return newGameMapTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newGameMapTable(schemaName, tableName, alias string) *GameMapTable {
	return &GameMapTable{
		gameMapTable: newGameMapTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newGameMapTableImpl("", "excluded", ""),
	}
}

func newGameMapTableImpl(schemaName, tableName, alias string) gameMapTable {
	var (
		IDColumn              = sqlite.StringColumn("id")
		CreatedAtColumn       = sqlite.TimestampColumn("created_at")
		UpdatedAtColumn       = sqlite.TimestampColumn("updated_at")
		GameModesColumn       = sqlite.StringColumn("game_modes")
		SupremacyPointsColumn = sqlite.IntegerColumn("supremacy_points")
		LocalizedNamesColumn  = sqlite.StringColumn("localized_names")
		allColumns            = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, GameModesColumn, SupremacyPointsColumn, LocalizedNamesColumn}
		mutableColumns        = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, GameModesColumn, SupremacyPointsColumn, LocalizedNamesColumn}
	)

	return gameMapTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedAt:       UpdatedAtColumn,
		GameModes:       GameModesColumn,
		SupremacyPoints: SupremacyPointsColumn,
		LocalizedNames:  LocalizedNamesColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
