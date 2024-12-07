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

var DiscordInteraction = newDiscordInteractionTable("", "discord_interaction", "")

type discordInteractionTable struct {
	sqlite.Table

	// Columns
	ID        sqlite.ColumnString
	CreatedAt sqlite.ColumnString
	UpdatedAt sqlite.ColumnString
	Result    sqlite.ColumnString
	EventID   sqlite.ColumnString
	GuildID   sqlite.ColumnString
	Snowflake sqlite.ColumnString
	ChannelID sqlite.ColumnString
	MessageID sqlite.ColumnString
	Type      sqlite.ColumnString
	Locale    sqlite.ColumnString
	Metadata  sqlite.ColumnString
	UserID    sqlite.ColumnString

	AllColumns     sqlite.ColumnList
	MutableColumns sqlite.ColumnList
}

type DiscordInteractionTable struct {
	discordInteractionTable

	EXCLUDED discordInteractionTable
}

// AS creates new DiscordInteractionTable with assigned alias
func (a DiscordInteractionTable) AS(alias string) *DiscordInteractionTable {
	return newDiscordInteractionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DiscordInteractionTable with assigned schema name
func (a DiscordInteractionTable) FromSchema(schemaName string) *DiscordInteractionTable {
	return newDiscordInteractionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DiscordInteractionTable with assigned table prefix
func (a DiscordInteractionTable) WithPrefix(prefix string) *DiscordInteractionTable {
	return newDiscordInteractionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DiscordInteractionTable with assigned table suffix
func (a DiscordInteractionTable) WithSuffix(suffix string) *DiscordInteractionTable {
	return newDiscordInteractionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDiscordInteractionTable(schemaName, tableName, alias string) *DiscordInteractionTable {
	return &DiscordInteractionTable{
		discordInteractionTable: newDiscordInteractionTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newDiscordInteractionTableImpl("", "excluded", ""),
	}
}

func newDiscordInteractionTableImpl(schemaName, tableName, alias string) discordInteractionTable {
	var (
		IDColumn        = sqlite.StringColumn("id")
		CreatedAtColumn = sqlite.StringColumn("created_at")
		UpdatedAtColumn = sqlite.StringColumn("updated_at")
		ResultColumn    = sqlite.StringColumn("result")
		EventIDColumn   = sqlite.StringColumn("event_id")
		GuildIDColumn   = sqlite.StringColumn("guild_id")
		SnowflakeColumn = sqlite.StringColumn("snowflake")
		ChannelIDColumn = sqlite.StringColumn("channel_id")
		MessageIDColumn = sqlite.StringColumn("message_id")
		TypeColumn      = sqlite.StringColumn("type")
		LocaleColumn    = sqlite.StringColumn("locale")
		MetadataColumn  = sqlite.StringColumn("metadata")
		UserIDColumn    = sqlite.StringColumn("user_id")
		allColumns      = sqlite.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, ResultColumn, EventIDColumn, GuildIDColumn, SnowflakeColumn, ChannelIDColumn, MessageIDColumn, TypeColumn, LocaleColumn, MetadataColumn, UserIDColumn}
		mutableColumns  = sqlite.ColumnList{CreatedAtColumn, UpdatedAtColumn, ResultColumn, EventIDColumn, GuildIDColumn, SnowflakeColumn, ChannelIDColumn, MessageIDColumn, TypeColumn, LocaleColumn, MetadataColumn, UserIDColumn}
	)

	return discordInteractionTable{
		Table: sqlite.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		Result:    ResultColumn,
		EventID:   EventIDColumn,
		GuildID:   GuildIDColumn,
		Snowflake: SnowflakeColumn,
		ChannelID: ChannelIDColumn,
		MessageID: MessageIDColumn,
		Type:      TypeColumn,
		Locale:    LocaleColumn,
		Metadata:  MetadataColumn,
		UserID:    UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
