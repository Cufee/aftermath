package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable(),
		field.Int64("created_at").
			Immutable().
			DefaultFunc(timeNow),
		field.Int64("updated_at").
			DefaultFunc(timeNow).
			UpdateDefault(timeNow),
		//
		field.String("permissions").Default(""),
		field.Strings("feature_flags").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("discord_interactions", DiscordInteraction.Type),
		edge.To("subscriptions", UserSubscription.Type),
		edge.To("connections", UserConnection.Type),
		edge.To("content", UserContent.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
