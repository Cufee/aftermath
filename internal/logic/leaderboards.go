package logic

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

func UpdateAccountAchievementsLeaderboardScores(ctx context.Context, wgClient wargaming.Client, dbClient database.Client, realm string, force bool, accountIDs ...string) (map[string]error, error) {
	if len(accountIDs) < 1 {
		return nil, nil
	}
	createdAt := time.Now()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	accounts, err := wgClient.BatchAccountByID(ctx, realm, accountIDs, "account_id", "last_battle_time")
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch accounts")
	}

	accountsNeedAnUpdate, err := filterActiveAccounts(ctx, dbClient, realm, accounts, false)
	if err != nil {
		return nil, err
	}
	if len(accountsNeedAnUpdate) < 1 {
		return nil, nil
	}

	// get current account achievements
	achievements, err := wgClient.BatchAccountAchievements(ctx, realm, accountsNeedAnUpdate)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account achievements")
	}

	var scores []models.LeaderboardScore
	// calculate the current score for each player
	for id, achievements := range achievements {
		scores = append(scores, models.LeaderboardScore{
			ReferenceID:   id,
			CreatedAt:     createdAt,
			UpdatedAt:     createdAt,
			Type:          models.LeaderboardScoreHourly,
			Score:         masteryAchievementsScore(achievements),
			LeaderboardID: models.LeaderboardsMasteryWeightedPlayers,
			Meta: map[string]any{
				"achievements": achievements,
			},
		})
	}

	sErrors, err := dbClient.CreateLeaderboardScores(ctx, scores...)
	if err != nil {
		return nil, err
	}
	if len(sErrors) > 0 {
		return sErrors, nil
	}

	return nil, nil
}

func masteryAchievementsScore(achievements types.AchievementsFrame) float32 {
	return float32(achievements.MarkOfMastery)*25 + float32(achievements.MarkOfMasteryI)*5 + float32(achievements.MarkOfMasteryII)*1 + float32(achievements.MarkOfMasteryIII)*1
}
