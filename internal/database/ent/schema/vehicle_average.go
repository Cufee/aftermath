package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/stats/frame"
)

// VehicleAverage holds the schema definition for the VehicleAverage entity.
type VehicleAverage struct {
	ent.Schema
}

// Fields of the VehicleAverage.
func (VehicleAverage) Fields() []ent.Field {
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
		field.JSON("data", frame.StatsFrame{}),
	}
}

// Edges of the VehicleAverage.
func (VehicleAverage) Edges() []ent.Edge {
	return nil
}

func (VehicleAverage) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
	}
}
