package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"golang.org/x/text/language"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
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
		field.Int("tier").
			Min(0). // vehicle that does not exist in official glossary has tier set to 0
			Max(10),
		field.JSON("localized_names", map[language.Tag]string{}),
	}

}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return nil
}

func (Vehicle) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
