package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.Int("created_at").
			Immutable().
			DefaultFunc(timeNow),
		field.Int("updated_at").
			DefaultFunc(timeNow).
			UpdateDefault(timeNow),
		//
		field.Int("tier").
			Min(0). // vehicle that does not exist in official glossary has tier set to 0
			Max(10),
		field.JSON("localized_names", map[string]string{}),
	}

}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return nil
}

func (Vehicle) Indexes() []ent.Index {
	return nil
}
