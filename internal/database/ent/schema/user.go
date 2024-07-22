package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Time("created_at").
			Immutable().
			Default(timeNow),
		field.Time("updated_at").
			Default(timeNow).
			UpdateDefault(timeNow),
		//
		field.String("username").Default(""),
		field.String("permissions").Default(""),
		field.Strings("feature_flags").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("discord_interactions", DiscordInteraction.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("subscriptions", UserSubscription.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("connections", UserConnection.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("widgets", WidgetSettings.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("content", UserContent.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("sessions", Session.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("moderation_requests", ModerationRequest.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("moderation_actions", ModerationRequest.Type).Annotations(entsql.OnDelete(entsql.NoAction)),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("username"),
	}
}
