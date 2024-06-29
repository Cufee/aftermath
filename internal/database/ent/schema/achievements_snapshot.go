package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type AchievementsSnapshot struct {
	ent.Schema
}

func (AchievementsSnapshot) Fields() []ent.Field {
	return append(defaultFields,
		field.Enum("type").
			GoType(models.SnapshotType("")),
		field.String("account_id").NotEmpty().Immutable(),
		field.String("reference_id").NotEmpty(),
		//
		field.Int("battles"),
		field.Time("last_battle_time"),
		field.JSON("data", types.AchievementsFrame{}),
	)
}

func (AchievementsSnapshot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("achievement_snapshots").Field("account_id").Required().Immutable().Unique(),
	}
}

func (AchievementsSnapshot) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("created_at"),
		index.Fields("account_id", "reference_id"),
		index.Fields("account_id", "reference_id", "created_at"),
	}
}
