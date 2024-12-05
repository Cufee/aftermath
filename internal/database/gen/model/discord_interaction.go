//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type DiscordInteraction struct {
	ID        string    `sql:"primary_key" db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Result    string    `db:"result"`
	EventID   string    `db:"event_id"`
	GuildID   string    `db:"guild_id"`
	Snowflake string    `db:"snowflake"`
	ChannelID string    `db:"channel_id"`
	MessageID string    `db:"message_id"`
	Type      string    `db:"type"`
	Locale    string    `db:"locale"`
	Metadata  []byte    `db:"metadata"`
	UserID    string    `db:"user_id"`
}
