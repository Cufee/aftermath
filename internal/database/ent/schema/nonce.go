package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type AuthNonce struct {
	ent.Schema
}

func (AuthNonce) Fields() []ent.Field {
	return append(defaultFields,
		field.Bool("active"),
		field.Time("expires_at").Immutable(),
		field.String("identifier").NotEmpty().Immutable(),
		field.String("public_id").NotEmpty().Immutable().Unique(),
		field.JSON("metadata", map[string]string{}),
	)
}

func (AuthNonce) Edges() []ent.Edge {
	return nil
}

func (AuthNonce) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("public_id", "active", "expires_at"),
	}
}
