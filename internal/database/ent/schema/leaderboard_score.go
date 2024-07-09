package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

type LeaderboardScore struct {
	ent.Schema
}

func (LeaderboardScore) Fields() []ent.Field {
	return append(
		defaultFields,
		field.Enum("type").
			GoType(models.ScoreType("")),
		field.Float32("score"),
		field.String("account_id"),
		field.String("reference_id"),
		field.String("leaderboard_id"),
		field.JSON("meta", map[string]any{}),
	)
}

func (LeaderboardScore) Edges() []ent.Edge {
	return nil
}

func (LeaderboardScore) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("created_at"),
		index.Fields("account_id"),
		index.Fields("reference_id"),
		index.Fields("leaderboard_id", "type", "account_id"),
		index.Fields("leaderboard_id", "type", "reference_id"),
	}
}
