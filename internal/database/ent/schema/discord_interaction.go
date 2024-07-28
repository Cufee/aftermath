package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

type DiscordInteraction struct {
	ent.Schema
}

func (DiscordInteraction) Fields() []ent.Field {
	return append(defaultFields,
		field.String("result").NotEmpty().Immutable(),
		field.String("user_id").NotEmpty().Immutable(),
		field.String("event_id").NotEmpty().Immutable(),
		field.String("guild_id").Immutable(),
		field.String("snowflake").Default(""),
		field.String("channel_id").Immutable(),
		field.String("message_id").Immutable(),
		field.Enum("type").
			GoType(models.DiscordInteractionType("")),
		field.String("locale"),
		field.JSON("metadata", map[string]any{}),
	)
}

func (DiscordInteraction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("discord_interactions").Field("user_id").Required().Immutable().Unique(),
	}
}

func (DiscordInteraction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("user_id"),
		index.Fields("snowflake"),
		index.Fields("created_at"),
		index.Fields("user_id", "type", "created_at"),
		index.Fields("channel_id", "type", "created_at"),
	}
}
