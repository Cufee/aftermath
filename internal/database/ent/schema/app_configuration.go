package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// AppConfiguration holds the schema definition for the AppConfiguration entity.
type AppConfiguration struct {
	ent.Schema
}

// Fields of the AppConfiguration.
func (AppConfiguration) Fields() []ent.Field {
	return append(defaultFields,
		field.String("key").Unique().NotEmpty(),
		//
		field.Any("value"),
		field.JSON("metadata", map[string]any{}).Optional(),
	)
}

// Edges of the AppConfiguration.
func (AppConfiguration) Edges() []ent.Edge {
	return nil
}

func (AppConfiguration) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("key"),
	}
}
