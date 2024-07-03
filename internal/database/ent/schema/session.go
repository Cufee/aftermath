package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Session struct {
	ent.Schema
}

func (Session) Fields() []ent.Field {
	return append(defaultFields,
		field.Time("expires_at"),
		field.String("user_id").Immutable(),
		field.String("public_id").NotEmpty().Immutable().Unique(),
		field.JSON("metadata", map[string]string{}),
	)
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("sessions").Field("user_id").Required().Immutable().Unique(),
	}
}

func (Session) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("public_id", "expires_at"),
	}
}
