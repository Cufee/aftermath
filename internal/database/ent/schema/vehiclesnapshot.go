package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

// VehicleSnapshot holds the schema definition for the VehicleSnapshot entity.
type VehicleSnapshot struct {
	ent.Schema
}

// Fields of the VehicleSnapshot.
func (VehicleSnapshot) Fields() []ent.Field {
	return append(defaultFields,
		field.Enum("type").
			GoType(models.SnapshotType("")),
		//
		field.String("account_id").NotEmpty().Immutable(),
		field.String("vehicle_id").NotEmpty().Immutable(),
		field.String("reference_id").NotEmpty(),
		//
		field.Int("battles"),
		field.Int("last_battle_time"),
		field.JSON("frame", frame.StatsFrame{}),
	)
}

// Edges of the VehicleSnapshot.
func (VehicleSnapshot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("vehicle_snapshots").Field("account_id").Required().Immutable().Unique(),
	}
}

func (VehicleSnapshot) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("vehicle_id", "created_at"),
		index.Fields("account_id", "created_at"),
		index.Fields("account_id", "type", "created_at"),
	}
}
