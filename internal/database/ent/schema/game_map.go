package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"golang.org/x/text/language"
)

// Vehicle holds the schema definition for the Vehicle entity.
type GameMap struct {
	ent.Schema
}

// Fields of the Vehicle.
func (GameMap) Fields() []ent.Field {
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
		field.Ints("game_modes"),
		field.Int("supremacy_points"),
		field.JSON("localized_names", map[language.Tag]string{}),
	}

}

// Edges of the Vehicle.
func (GameMap) Edges() []ent.Edge {
	return nil
}

func (GameMap) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
