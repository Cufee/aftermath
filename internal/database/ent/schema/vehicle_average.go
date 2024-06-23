package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.Int64("created_at").
			Immutable().
			DefaultFunc(timeNow),
		field.Int64("updated_at").
			DefaultFunc(timeNow).
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
	return nil
}
