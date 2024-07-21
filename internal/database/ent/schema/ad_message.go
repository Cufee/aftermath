package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"golang.org/x/text/language"
)

// AppConfiguration holds the schema definition for the AppConfiguration entity.
type AdMessage struct {
	ent.Schema
}

// Fields of the AppConfiguration.
func (AdMessage) Fields() []ent.Field {
	return append(defaultFields,
		field.Bool("enabled"),
		field.Int("weight"),
		field.Float32("chance"),
		field.JSON("message", map[language.Tag]string{}),
		field.JSON("metadata", map[string]any{}).Optional(),
	)
}

// Edges of the AppConfiguration.
func (AdMessage) Edges() []ent.Edge {
	return nil
}

func (AdMessage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("weight", "enabled"),
	}
}
