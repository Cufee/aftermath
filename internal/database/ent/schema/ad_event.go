package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AppConfiguration holds the schema definition for the AppConfiguration entity.
type AdEvent struct {
	ent.Schema
}

// Fields of the AppConfiguration.
func (AdEvent) Fields() []ent.Field {
	return append(defaultFields,
		field.String("user_id").NotEmpty(),
		field.String("guild_id"),
		field.String("channel_id").NotEmpty(),
		field.String("locale"),
		field.String("message_id"),
		field.JSON("metadata", map[string]any{}).Optional(),
	)
}

// Edges of the AppConfiguration.
func (AdEvent) Edges() []ent.Edge {
	return nil
}

func (AdEvent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("user_id", "guild_id", "channel_id", "created_at"),
	}
}
