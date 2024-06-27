package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

// AccountSnapshot holds the schema definition for the AccountSnapshot entity.
type AccountSnapshot struct {
	ent.Schema
}

// Fields of the AccountSnapshot.
func (AccountSnapshot) Fields() []ent.Field {
	return append(defaultFields,
		field.Enum("type").
			GoType(models.SnapshotType("")),
		field.Time("last_battle_time"),
		//
		field.String("account_id").NotEmpty().Immutable(),
		field.String("reference_id").NotEmpty(),
		//
		field.Int("rating_battles"),
		field.JSON("rating_frame", frame.StatsFrame{}),
		//
		field.Int("regular_battles"),
		field.JSON("regular_frame", frame.StatsFrame{}),
	)
}

// Edges of the AccountSnapshot.
func (AccountSnapshot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("snapshots").Field("account_id").Required().Immutable().Unique(),
	}
}

func (AccountSnapshot) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("created_at"),
		index.Fields("type", "account_id", "created_at"),
		index.Fields("type", "account_id", "reference_id"),
		index.Fields("type", "account_id", "reference_id", "created_at"),
	}
}
