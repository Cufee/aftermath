// Code generated by ent, DO NOT EDIT.

package discordinteraction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldUpdatedAt, v))
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldResult, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldUserID, v))
}

// EventID applies equality check predicate on the "event_id" field. It's identical to EventIDEQ.
func EventID(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldEventID, v))
}

// GuildID applies equality check predicate on the "guild_id" field. It's identical to GuildIDEQ.
func GuildID(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldGuildID, v))
}

// Snowflake applies equality check predicate on the "snowflake" field. It's identical to SnowflakeEQ.
func Snowflake(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldSnowflake, v))
}

// ChannelID applies equality check predicate on the "channel_id" field. It's identical to ChannelIDEQ.
func ChannelID(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldChannelID, v))
}

// MessageID applies equality check predicate on the "message_id" field. It's identical to MessageIDEQ.
func MessageID(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldMessageID, v))
}

// Locale applies equality check predicate on the "locale" field. It's identical to LocaleEQ.
func Locale(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldLocale, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldUpdatedAt, v))
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldResult, v))
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldResult, v))
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldResult, vs...))
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldResult, vs...))
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldResult, v))
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldResult, v))
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldResult, v))
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldResult, v))
}

// ResultContains applies the Contains predicate on the "result" field.
func ResultContains(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContains(FieldResult, v))
}

// ResultHasPrefix applies the HasPrefix predicate on the "result" field.
func ResultHasPrefix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasPrefix(FieldResult, v))
}

// ResultHasSuffix applies the HasSuffix predicate on the "result" field.
func ResultHasSuffix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasSuffix(FieldResult, v))
}

// ResultEqualFold applies the EqualFold predicate on the "result" field.
func ResultEqualFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldResult, v))
}

// ResultContainsFold applies the ContainsFold predicate on the "result" field.
func ResultContainsFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldResult, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldUserID, v))
}

// EventIDEQ applies the EQ predicate on the "event_id" field.
func EventIDEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldEventID, v))
}

// EventIDNEQ applies the NEQ predicate on the "event_id" field.
func EventIDNEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldEventID, v))
}

// EventIDIn applies the In predicate on the "event_id" field.
func EventIDIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldEventID, vs...))
}

// EventIDNotIn applies the NotIn predicate on the "event_id" field.
func EventIDNotIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldEventID, vs...))
}

// EventIDGT applies the GT predicate on the "event_id" field.
func EventIDGT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldEventID, v))
}

// EventIDGTE applies the GTE predicate on the "event_id" field.
func EventIDGTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldEventID, v))
}

// EventIDLT applies the LT predicate on the "event_id" field.
func EventIDLT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldEventID, v))
}

// EventIDLTE applies the LTE predicate on the "event_id" field.
func EventIDLTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldEventID, v))
}

// EventIDContains applies the Contains predicate on the "event_id" field.
func EventIDContains(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContains(FieldEventID, v))
}

// EventIDHasPrefix applies the HasPrefix predicate on the "event_id" field.
func EventIDHasPrefix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasPrefix(FieldEventID, v))
}

// EventIDHasSuffix applies the HasSuffix predicate on the "event_id" field.
func EventIDHasSuffix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasSuffix(FieldEventID, v))
}

// EventIDEqualFold applies the EqualFold predicate on the "event_id" field.
func EventIDEqualFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldEventID, v))
}

// EventIDContainsFold applies the ContainsFold predicate on the "event_id" field.
func EventIDContainsFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldEventID, v))
}

// GuildIDEQ applies the EQ predicate on the "guild_id" field.
func GuildIDEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldGuildID, v))
}

// GuildIDNEQ applies the NEQ predicate on the "guild_id" field.
func GuildIDNEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldGuildID, v))
}

// GuildIDIn applies the In predicate on the "guild_id" field.
func GuildIDIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldGuildID, vs...))
}

// GuildIDNotIn applies the NotIn predicate on the "guild_id" field.
func GuildIDNotIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldGuildID, vs...))
}

// GuildIDGT applies the GT predicate on the "guild_id" field.
func GuildIDGT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldGuildID, v))
}

// GuildIDGTE applies the GTE predicate on the "guild_id" field.
func GuildIDGTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldGuildID, v))
}

// GuildIDLT applies the LT predicate on the "guild_id" field.
func GuildIDLT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldGuildID, v))
}

// GuildIDLTE applies the LTE predicate on the "guild_id" field.
func GuildIDLTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldGuildID, v))
}

// GuildIDContains applies the Contains predicate on the "guild_id" field.
func GuildIDContains(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContains(FieldGuildID, v))
}

// GuildIDHasPrefix applies the HasPrefix predicate on the "guild_id" field.
func GuildIDHasPrefix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasPrefix(FieldGuildID, v))
}

// GuildIDHasSuffix applies the HasSuffix predicate on the "guild_id" field.
func GuildIDHasSuffix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasSuffix(FieldGuildID, v))
}

// GuildIDEqualFold applies the EqualFold predicate on the "guild_id" field.
func GuildIDEqualFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldGuildID, v))
}

// GuildIDContainsFold applies the ContainsFold predicate on the "guild_id" field.
func GuildIDContainsFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldGuildID, v))
}

// SnowflakeEQ applies the EQ predicate on the "snowflake" field.
func SnowflakeEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldSnowflake, v))
}

// SnowflakeNEQ applies the NEQ predicate on the "snowflake" field.
func SnowflakeNEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldSnowflake, v))
}

// SnowflakeIn applies the In predicate on the "snowflake" field.
func SnowflakeIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldSnowflake, vs...))
}

// SnowflakeNotIn applies the NotIn predicate on the "snowflake" field.
func SnowflakeNotIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldSnowflake, vs...))
}

// SnowflakeGT applies the GT predicate on the "snowflake" field.
func SnowflakeGT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldSnowflake, v))
}

// SnowflakeGTE applies the GTE predicate on the "snowflake" field.
func SnowflakeGTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldSnowflake, v))
}

// SnowflakeLT applies the LT predicate on the "snowflake" field.
func SnowflakeLT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldSnowflake, v))
}

// SnowflakeLTE applies the LTE predicate on the "snowflake" field.
func SnowflakeLTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldSnowflake, v))
}

// SnowflakeContains applies the Contains predicate on the "snowflake" field.
func SnowflakeContains(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContains(FieldSnowflake, v))
}

// SnowflakeHasPrefix applies the HasPrefix predicate on the "snowflake" field.
func SnowflakeHasPrefix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasPrefix(FieldSnowflake, v))
}

// SnowflakeHasSuffix applies the HasSuffix predicate on the "snowflake" field.
func SnowflakeHasSuffix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasSuffix(FieldSnowflake, v))
}

// SnowflakeEqualFold applies the EqualFold predicate on the "snowflake" field.
func SnowflakeEqualFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldSnowflake, v))
}

// SnowflakeContainsFold applies the ContainsFold predicate on the "snowflake" field.
func SnowflakeContainsFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldSnowflake, v))
}

// ChannelIDEQ applies the EQ predicate on the "channel_id" field.
func ChannelIDEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldChannelID, v))
}

// ChannelIDNEQ applies the NEQ predicate on the "channel_id" field.
func ChannelIDNEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldChannelID, v))
}

// ChannelIDIn applies the In predicate on the "channel_id" field.
func ChannelIDIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldChannelID, vs...))
}

// ChannelIDNotIn applies the NotIn predicate on the "channel_id" field.
func ChannelIDNotIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldChannelID, vs...))
}

// ChannelIDGT applies the GT predicate on the "channel_id" field.
func ChannelIDGT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldChannelID, v))
}

// ChannelIDGTE applies the GTE predicate on the "channel_id" field.
func ChannelIDGTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldChannelID, v))
}

// ChannelIDLT applies the LT predicate on the "channel_id" field.
func ChannelIDLT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldChannelID, v))
}

// ChannelIDLTE applies the LTE predicate on the "channel_id" field.
func ChannelIDLTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldChannelID, v))
}

// ChannelIDContains applies the Contains predicate on the "channel_id" field.
func ChannelIDContains(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContains(FieldChannelID, v))
}

// ChannelIDHasPrefix applies the HasPrefix predicate on the "channel_id" field.
func ChannelIDHasPrefix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasPrefix(FieldChannelID, v))
}

// ChannelIDHasSuffix applies the HasSuffix predicate on the "channel_id" field.
func ChannelIDHasSuffix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasSuffix(FieldChannelID, v))
}

// ChannelIDEqualFold applies the EqualFold predicate on the "channel_id" field.
func ChannelIDEqualFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldChannelID, v))
}

// ChannelIDContainsFold applies the ContainsFold predicate on the "channel_id" field.
func ChannelIDContainsFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldChannelID, v))
}

// MessageIDEQ applies the EQ predicate on the "message_id" field.
func MessageIDEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldMessageID, v))
}

// MessageIDNEQ applies the NEQ predicate on the "message_id" field.
func MessageIDNEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldMessageID, v))
}

// MessageIDIn applies the In predicate on the "message_id" field.
func MessageIDIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldMessageID, vs...))
}

// MessageIDNotIn applies the NotIn predicate on the "message_id" field.
func MessageIDNotIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldMessageID, vs...))
}

// MessageIDGT applies the GT predicate on the "message_id" field.
func MessageIDGT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldMessageID, v))
}

// MessageIDGTE applies the GTE predicate on the "message_id" field.
func MessageIDGTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldMessageID, v))
}

// MessageIDLT applies the LT predicate on the "message_id" field.
func MessageIDLT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldMessageID, v))
}

// MessageIDLTE applies the LTE predicate on the "message_id" field.
func MessageIDLTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldMessageID, v))
}

// MessageIDContains applies the Contains predicate on the "message_id" field.
func MessageIDContains(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContains(FieldMessageID, v))
}

// MessageIDHasPrefix applies the HasPrefix predicate on the "message_id" field.
func MessageIDHasPrefix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasPrefix(FieldMessageID, v))
}

// MessageIDHasSuffix applies the HasSuffix predicate on the "message_id" field.
func MessageIDHasSuffix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasSuffix(FieldMessageID, v))
}

// MessageIDEqualFold applies the EqualFold predicate on the "message_id" field.
func MessageIDEqualFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldMessageID, v))
}

// MessageIDContainsFold applies the ContainsFold predicate on the "message_id" field.
func MessageIDContainsFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldMessageID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v models.DiscordInteractionType) predicate.DiscordInteraction {
	vc := v
	return predicate.DiscordInteraction(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v models.DiscordInteractionType) predicate.DiscordInteraction {
	vc := v
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...models.DiscordInteractionType) predicate.DiscordInteraction {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscordInteraction(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...models.DiscordInteractionType) predicate.DiscordInteraction {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldType, v...))
}

// LocaleEQ applies the EQ predicate on the "locale" field.
func LocaleEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEQ(FieldLocale, v))
}

// LocaleNEQ applies the NEQ predicate on the "locale" field.
func LocaleNEQ(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNEQ(FieldLocale, v))
}

// LocaleIn applies the In predicate on the "locale" field.
func LocaleIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldIn(FieldLocale, vs...))
}

// LocaleNotIn applies the NotIn predicate on the "locale" field.
func LocaleNotIn(vs ...string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldNotIn(FieldLocale, vs...))
}

// LocaleGT applies the GT predicate on the "locale" field.
func LocaleGT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGT(FieldLocale, v))
}

// LocaleGTE applies the GTE predicate on the "locale" field.
func LocaleGTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldGTE(FieldLocale, v))
}

// LocaleLT applies the LT predicate on the "locale" field.
func LocaleLT(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLT(FieldLocale, v))
}

// LocaleLTE applies the LTE predicate on the "locale" field.
func LocaleLTE(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldLTE(FieldLocale, v))
}

// LocaleContains applies the Contains predicate on the "locale" field.
func LocaleContains(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContains(FieldLocale, v))
}

// LocaleHasPrefix applies the HasPrefix predicate on the "locale" field.
func LocaleHasPrefix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasPrefix(FieldLocale, v))
}

// LocaleHasSuffix applies the HasSuffix predicate on the "locale" field.
func LocaleHasSuffix(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldHasSuffix(FieldLocale, v))
}

// LocaleEqualFold applies the EqualFold predicate on the "locale" field.
func LocaleEqualFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldEqualFold(FieldLocale, v))
}

// LocaleContainsFold applies the ContainsFold predicate on the "locale" field.
func LocaleContainsFold(v string) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.FieldContainsFold(FieldLocale, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.DiscordInteraction {
	return predicate.DiscordInteraction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DiscordInteraction) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DiscordInteraction) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DiscordInteraction) predicate.DiscordInteraction {
	return predicate.DiscordInteraction(sql.NotPredicates(p))
}
