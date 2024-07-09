package database

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func TestGetLeaderboardScores(t *testing.T) {
	client, err := NewSQLiteClient(filepath.Join(os.Getenv("DATABASE_PATH"), os.Getenv("DATABASE_NAME")), WithDebug())
	assert.NoError(t, err, "new client should not error")
	defer client.Disconnect()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	client.db.LeaderboardScore.Delete().Where().Exec(ctx)
	defer client.db.LeaderboardScore.Delete().Exec(ctx)

	score1CreatedAt := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	score1 := models.LeaderboardScore{
		CreatedAt:     score1CreatedAt,
		UpdatedAt:     score1CreatedAt,
		Type:          models.LeaderboardScoreHourly,
		ReferenceID:   "r1",
		LeaderboardID: models.LeaderboardsMasteryWeightedPlayers,
		Score:         1,
		Meta:          map[string]any{},
	}
	score2CreatedAt := time.Date(2023, 6, 2, 0, 0, 0, 0, time.UTC)
	score2 := models.LeaderboardScore{
		CreatedAt:     score2CreatedAt,
		UpdatedAt:     score2CreatedAt,
		Type:          models.LeaderboardScoreHourly,
		ReferenceID:   "r1",
		LeaderboardID: models.LeaderboardsMasteryWeightedPlayers,
		Score:         2,
		Meta:          map[string]any{},
	}
	score3CreatedAt := time.Date(2023, 6, 3, 0, 0, 0, 0, time.UTC)
	score3 := models.LeaderboardScore{
		CreatedAt:     score3CreatedAt,
		UpdatedAt:     score3CreatedAt,
		Type:          models.LeaderboardScoreHourly,
		ReferenceID:   "r1",
		LeaderboardID: models.LeaderboardsMasteryWeightedPlayers,
		Score:         3,
		Meta:          map[string]any{},
	}

	sErr, err := client.CreateLeaderboardScores(context.Background(), score1, score2, score3)
	assert.NoError(t, err, "failed to create scores")
	assert.Nil(t, sErr, "failed to create scores")

	{ // default order is created_at desc, this should return score3
		scores, err := client.GetLeaderboardScores(context.Background(), models.LeaderboardsMasteryWeightedPlayers, models.LeaderboardScoreHourly)
		assert.NoError(t, err, "failed to get leaderboard scores")
		assert.Len(t, scores, 1, "scores are grouped by referenceID, this should return exactly 1 result")
		assert.Equal(t, score3.Score, scores[0].Score, "default sort is desc, should return score3")
	}
}
