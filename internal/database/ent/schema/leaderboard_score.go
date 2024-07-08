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
		field.String("reference_id"),
		field.Enum("leaderboard_id").
			GoType(models.LeaderboardID("")),
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
		index.Fields("created_at", "type"),
		index.Fields("score", "type"),
		index.Fields("leaderboard_id", "type"),
		index.Fields("leaderboard_id", "score", "type"),
		index.Fields("leaderboard_id", "reference_id", "type"),
		index.Fields("leaderboard_id", "reference_id", "score", "type"),
	}
}
