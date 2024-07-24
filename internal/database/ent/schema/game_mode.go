package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"golang.org/x/text/language"
)

type GameMode struct {
	ent.Schema
}

func (GameMode) Fields() []ent.Field {
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
		field.JSON("localized_names", map[language.Tag]string{}),
	}

}

func (GameMode) Edges() []ent.Edge {
	return nil
}

func (GameMode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
