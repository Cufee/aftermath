package logic

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

func RecordCurrentAchievementsLeaderboards(ctx context.Context, wgClient wargaming.Client, dbClient database.Client, scoreType models.ScoreType, realm string, force bool, accountIDs ...string) (map[string]error, error) {
	if len(accountIDs) < 1 {
		return nil, nil
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	var playerLeaderboard retry.DataWithErr[map[string]models.LeaderboardScore]
	wg.Add(1) // get last leaderboard entry for each account
	go func() {
		defer wg.Done()
		leaderboard, err := dbClient.GetLeaderboardScores(ctx, models.LeaderboardsMasteryWeightedPlayers, scoreType, database.WithReferenceIDIn(accountIDs...))
		if err != nil && !database.IsNotFound(err) {
			playerLeaderboard = retry.DataWithErr[map[string]models.LeaderboardScore]{Err: err}
			return
		}
		leaderboardMap := make(map[string]models.LeaderboardScore)
		for _, score := range leaderboard {
			leaderboardMap[score.ReferenceID] = score
		}
		playerLeaderboard = retry.DataWithErr[map[string]models.LeaderboardScore]{Data: leaderboardMap}
	}()

	var accounts retry.DataWithErr[map[string]types.ExtendedAccount]
	wg.Add(1) // get account last battle time for all accounts
	go func() {
		defer wg.Done()
		data, err := wgClient.BatchAccountByID(ctx, realm, accountIDs, "account_id", "last_battle_time")
		accounts = retry.DataWithErr[map[string]types.ExtendedAccount]{Data: data, Err: err}
	}()

	wg.Wait()
	if playerLeaderboard.Err != nil {
		return nil, errors.Wrap(playerLeaderboard.Err, "failed to get player leaderboard scores")
	}
	if accounts.Err != nil {
		return nil, errors.Wrap(accounts.Err, "failed to get accounts from wargaming")
	}

	var accountsToUpdate []string
	for _, id := range accountIDs {
		account, ok := accounts.Data[id]
		if !ok {
			continue // wargaming returned no data for account
		}
		score, sok := playerLeaderboard.Data[id]
		lastBattle, lok := score.Meta["lastBattleTime"].(time.Time)
		if !sok || !lok || lastBattle.Unix() < int64(account.LastBattleTime) { // player has no previous score on the leaderboard, the metadata is missing, or they played
			accountsToUpdate = append(accountsToUpdate, id)
		}
	}
	if len(accountsToUpdate) < 1 {
		return nil, nil
	}

	var accountClans retry.DataWithErr[map[string]types.ClanMember]
	wg.Add(1) // get account clans for all accounts
	go func() {
		defer wg.Done()
		data, err := wgClient.BatchAccountClan(ctx, realm, accountIDs, "account_id", "clan.clan_id")
		accountClans = retry.DataWithErr[map[string]types.ClanMember]{Data: data, Err: err}
	}()

	var currentAchievements retry.DataWithErr[map[string]types.AchievementsFrame]
	wg.Add(1) // get current achievements for all players
	go func() {
		defer wg.Done()
		data, err := wgClient.BatchAccountAchievements(ctx, realm, accountsToUpdate)
		currentAchievements = retry.DataWithErr[map[string]types.AchievementsFrame]{Data: data, Err: err}
	}()

	wg.Wait()
	if currentAchievements.Err != nil {
		return nil, currentAchievements.Err
	}
	if accountClans.Err != nil { // this error is not critical, but it will delay the leaderboard update for clans
		log.Err(accountClans.Err).Msg("failed to get account clans")
	}

	var accountErrors = make(map[string]error)
	var leaderboardScores []models.LeaderboardScore
	var clanScores = make(map[string]models.LeaderboardScore)
	for _, id := range accountsToUpdate {
		account, ok := accounts.Data[id]
		if !ok {
			continue // this should never happen due to the check above, but just in case
		}
		achievements, ok := currentAchievements.Data[id]
		if !ok {
			accountErrors[id] = errors.New("failed to get account achievements")
			continue
		}

		playerScore := models.LeaderboardScore{
			Type:          scoreType,
			ReferenceID:   id,
			LeaderboardID: models.LeaderboardsMasteryWeightedPlayers,
			Score:         masteryAchievementsScore(achievements),
			Meta: map[string]any{
				"lastBattleTime": account.LastBattleTime,
				"frame":          achievements,
			},
		}
		leaderboardScores = append(leaderboardScores, playerScore)

		clan, ok := accountClans.Data[id]
		if !ok || clan.ClanID == 0 { // account has no clan, or the request failed
			continue
		}

		clanID := fmt.Sprint(clan.ClanID)
		clanScore := clanScores[clanID]
		clanScore.Type = scoreType
		clanScore.ReferenceID = clanID
		clanScore.LeaderboardID = models.LeaderboardsMasteryWeightedClans

		clanScore.Score += playerScore.Score
		clanFrame, _ := clanScore.Meta["frame"].(types.AchievementsFrame)
		clanFrame.Add(&achievements)
		clanScore.Meta["frame"] = clanFrame

		clanScores[clanID] = clanScore
	}

	for _, score := range clanScores {
		leaderboardScores = append(leaderboardScores, score)
	}

	insertErrors, err := dbClient.CreateLeaderboardScores(ctx, leaderboardScores...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save leaderboard scores")
	}
	for id, err := range insertErrors {
		if _, ok := clanScores[id]; !ok {
			accountErrors[id] = err // if this is not a clan-related error
		}
	}
	if len(accountErrors) > 0 {
		return accountErrors, nil
	}
	return nil, nil
}

func masteryAchievementsScore(achievements types.AchievementsFrame) float32 {
	return float32(achievements.MarkOfMastery)*25 + float32(achievements.MarkOfMasteryI)*5 + float32(achievements.MarkOfMasteryII)*1 + float32(achievements.MarkOfMasteryIII)*1
}
