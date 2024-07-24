package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

type WidgetSettings struct {
	ent.Schema
}

// Fields of the Vehicle.
func (WidgetSettings) Fields() []ent.Field {
	return append(defaultFields,
		field.String("reference_id"),
		field.String("title").Optional(),
		field.String("user_id").Immutable(),
		field.Time("session_from").Optional(),
		field.String("session_reference_id").Optional(),
		field.JSON("metadata", map[string]any{}),
		field.JSON("styles", models.WidgetStyling{}),
	)
}

// Edges of the Vehicle.
func (WidgetSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("widgets").Field("user_id").Required().Immutable().Unique(),
	}
}

func (WidgetSettings) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
