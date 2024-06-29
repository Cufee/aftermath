package logic

import (
	"context"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"

	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type vehicleBattleData struct {
	LastBattle time.Time
	Battles    int
}

/*
Gets all accounts from WH and their respective snapshots
  - an account is considered valid is it has played a battle since the last snapshot, or has no snapshots
  - as side effect, invalid accounts will be set as private
*/
func getAndValidateAccounts(ctx context.Context, wgClient wargaming.Client, dbClient database.Client, force bool, realm string, accountIDs ...string) (map[string]types.ExtendedAccount, []string, map[string]*models.AccountSnapshot, error) {
	accounts, err := wgClient.BatchAccountByID(ctx, realm, accountIDs)
	if err != nil {
		return nil, nil, nil, errors.Wrap(err, "failed to fetch accounts")
	}

	// existing snapshots for accounts
	existingSnapshots, err := dbClient.GetManyAccountSnapshots(ctx, accountIDs, models.SnapshotTypeDaily)
	if err != nil && !database.IsNotFound(err) {
		return nil, nil, nil, errors.Wrap(err, "failed to get existing snapshots")
	}
	existingSnapshotsMap := make(map[string]*models.AccountSnapshot)
	for _, snapshot := range existingSnapshots {
		existingSnapshotsMap[snapshot.AccountID] = &snapshot
	}

	// make a new slice just in case some accounts were not returned/are private
	var needAnUpdate []string
	for _, id := range accountIDs {
		data, ok := accounts[id]
		if !ok {
			go func(id string) {
				log.Debug().Str("accountId", id).Msg("account is private")
				// update account cache (if it exists) to set account as private
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				err := dbClient.AccountSetPrivate(ctx, id, true)
				if err != nil {
					log.Err(err).Str("accountId", id).Msg("failed to set account status as private")
				}
			}(id)
			continue
		}
		if data.LastBattleTime < 1 {
			continue
		}
		if s, ok := existingSnapshotsMap[id]; !force && (ok && data.LastBattleTime == int(s.LastBattleTime.Unix())) {
			// last snapshot is the same, we can skip it
			continue
		}
		needAnUpdate = append(needAnUpdate, id)
	}
	return accounts, needAnUpdate, existingSnapshotsMap, nil
}

func RecordAccountSnapshots(ctx context.Context, wgClient wargaming.Client, dbClient database.Client, realm string, force bool, accountIDs ...string) (map[string]error, error) {
	if len(accountIDs) < 1 {
		return nil, nil
	}
	createdAt := time.Now()

	accounts, accountsNeedAnUpdate, _, err := getAndValidateAccounts(ctx, wgClient, dbClient, force, realm, accountIDs...)
	if err != nil {
		return nil, err
	}
	if len(accountsNeedAnUpdate) < 1 {
		return nil, nil
	}

	var group sync.WaitGroup
	var clans = make(map[string]types.ClanMember)
	var accountAchievements retry.DataWithErr[map[string]types.AchievementsFrame]

	var accountVehiclesMx sync.Mutex
	var accountVehicleAchievementsMx sync.Mutex
	accountVehicles := make(map[string]retry.DataWithErr[[]types.VehicleStatsFrame], len(accountsNeedAnUpdate))
	accountVehicleAchievements := make(map[string]retry.DataWithErr[map[string]types.AchievementsFrame], len(accountsNeedAnUpdate))

	for _, id := range accountsNeedAnUpdate {
		group.Add(1)
		// get account vehicle stats
		go func(id string) {
			defer group.Done()
			data, err := wgClient.AccountVehicles(ctx, realm, id)

			accountVehiclesMx.Lock()
			accountVehicles[id] = retry.DataWithErr[[]types.VehicleStatsFrame]{Data: data, Err: err}
			accountVehiclesMx.Unlock()
		}(id)

		group.Add(1)
		// get account vehicle achievements
		go func(id string) {
			defer group.Done()
			data, err := wgClient.AccountVehicleAchievements(ctx, realm, id)

			accountVehicleAchievementsMx.Lock()
			accountVehicleAchievements[id] = retry.DataWithErr[map[string]types.AchievementsFrame]{Data: data, Err: err}
			accountVehicleAchievementsMx.Unlock()
		}(id)
	}

	group.Add(1)
	// get account clans, not critical
	go func() {
		defer group.Done()
		// clans are optional-ish
		data, err := wgClient.BatchAccountClan(ctx, realm, accountsNeedAnUpdate)
		if err != nil {
			log.Err(err).Msg("failed to get batch account clans")
		}
		clans = data
	}()

	group.Add(1)
	// get account achievements, not critical
	go func() {
		defer group.Done()
		// clans are optional-ish
		data, err := wgClient.BatchAccountAchievements(ctx, realm, accountsNeedAnUpdate)
		if err != nil {
			log.Err(err).Msg("failed to get batch account achievements")
		}
		accountAchievements.Data = data
		accountAchievements.Err = err
	}()

	group.Wait()

	var accountErrors = make(map[string]error)
	var accountUpdates = make(map[string]models.Account)

	var accountSnapshots []models.AccountSnapshot
	var vehicleSnapshots = make(map[string][]models.VehicleSnapshot)
	var achievementsSnapshots = make(map[string][]models.AchievementsSnapshot)

	for _, accountID := range accountsNeedAnUpdate {
		vehicles := accountVehicles[accountID]
		if vehicles.Err != nil {
			accountErrors[accountID] = vehicles.Err
			continue
		}

		// get existing vehicle snapshots from db
		existingSnapshots, err := dbClient.GetVehicleSnapshots(ctx, accountID, accountID, models.SnapshotTypeDaily)
		if err != nil && !database.IsNotFound(err) {
			accountErrors[accountID] = err
			continue
		}
		existingSnapshotsMap := make(map[string]*models.VehicleSnapshot)
		for _, snapshot := range existingSnapshots {
			existingSnapshotsMap[snapshot.VehicleID] = &snapshot
		}

		snapshotStats := fetch.WargamingToStats(realm, accounts[accountID], clans[accountID], vehicles.Data)
		{ // account snapshot
			accountSnapshots = append(accountSnapshots, models.AccountSnapshot{
				Type:           models.SnapshotTypeDaily,
				CreatedAt:      createdAt,
				AccountID:      snapshotStats.Account.ID,
				ReferenceID:    snapshotStats.Account.ID,
				LastBattleTime: snapshotStats.LastBattleTime,
				RatingBattles:  snapshotStats.RatingBattles.StatsFrame,
				RegularBattles: snapshotStats.RegularBattles.StatsFrame,
			})
			accountUpdates[accountID] = snapshotStats.Account
		}

		// vehicle snapshots
		var vehicleLastBattleTimes = make(map[string]vehicleBattleData)
		vehicleStats := fetch.WargamingVehiclesToFrame(vehicles.Data)
		if len(vehicles.Data) > 0 {
			for id, vehicle := range vehicleStats {
				if s, ok := existingSnapshotsMap[id]; !force && ok && s.LastBattleTime.Equal(vehicle.LastBattleTime) {
					// last snapshot is the same, we can skip it
					continue
				}
				vehicleLastBattleTimes[id] = vehicleBattleData{vehicle.LastBattleTime, int(vehicle.Battles.Float())}
				vehicleSnapshots[accountID] = append(vehicleSnapshots[accountID], models.VehicleSnapshot{
					Type:           models.SnapshotTypeDaily,
					LastBattleTime: vehicle.LastBattleTime,
					Stats:          *vehicle.StatsFrame,
					VehicleID:      vehicle.VehicleID,
					AccountID:      accountID,
					ReferenceID:    accountID,
					CreatedAt:      createdAt,
				})
			}
		}

		// account achievement snapshot
		if accountAchievements.Err != nil {
			accountErrors[accountID] = errors.Wrap(accountAchievements.Err, "failed to get account achievements")
		}
		if achievements, ok := accountAchievements.Data[accountID]; ok {
			achievementsSnapshots[accountID] = append(achievementsSnapshots[accountID], models.AchievementsSnapshot{
				Data:           achievements,
				CreatedAt:      createdAt,
				AccountID:      accountID,
				ReferenceID:    accountID,
				Type:           models.SnapshotTypeDaily,
				LastBattleTime: snapshotStats.LastBattleTime,
				Battles:        int(snapshotStats.RegularBattles.Battles.Float()),
			})
		}

		// account vehicle achievement snapshots
		if achievements, ok := accountVehicleAchievements[accountID]; ok {
			if achievements.Err != nil {
				accountErrors[accountID] = errors.Wrap(achievements.Err, "failed to get vehicle achievements")
			}
			for vehicleID, a := range achievements.Data {
				battleData, ok := vehicleLastBattleTimes[vehicleID]
				if !ok {
					// vehicle was not played, no need to save achievements
					continue
				}

				achievementsSnapshots[accountID] = append(achievementsSnapshots[accountID], models.AchievementsSnapshot{
					Data:           a,
					CreatedAt:      createdAt,
					AccountID:      accountID,
					ReferenceID:    vehicleID,
					Battles:        battleData.Battles,
					LastBattleTime: battleData.LastBattle,
					Type:           models.SnapshotTypeDaily,
				})
			}
		}
	}

	for _, accountSnapshot := range accountSnapshots {
		vehicles := vehicleSnapshots[accountSnapshot.AccountID]
		if len(vehicles) > 0 {
			// save all vehicle snapshots)
			vErr, err := dbClient.CreateAccountVehicleSnapshots(ctx, accountSnapshot.AccountID, vehicles...)
			if err != nil {
				accountErrors[accountSnapshot.AccountID] = err
				continue
			}
			if len(vErr) > 0 {
				accountErrors[accountSnapshot.AccountID] = errors.Errorf("failed to insert %d vehicle snapshots", len(vErr))
				continue
			}
		}

		achievements := achievementsSnapshots[accountSnapshot.AccountID]
		if len(achievements) > 0 {
			achErr, err := dbClient.CreateAccountAchievementSnapshots(ctx, accountSnapshot.AccountID, achievements...)
			if err != nil {
				accountErrors[accountSnapshot.AccountID] = errors.Wrap(err, "failed to save account achievements to database")
				continue
			}

			for id, e := range achErr {
				if e != nil {
					err = errors.Wrapf(e, "failed to save some achievement snapshots, id %s", id)
					break
				}
			}
			if err != nil {
				accountErrors[accountSnapshot.AccountID] = err
				continue
			}

		}

		// save account snapshot
		aErr, err := dbClient.CreateAccountSnapshots(ctx, accountSnapshot)
		if err != nil {
			accountErrors[accountSnapshot.AccountID] = errors.Wrap(err, "failed to save account snapshots to database")
			continue
		}
		if err := aErr[accountSnapshot.AccountID]; err != nil {
			accountErrors[accountSnapshot.AccountID] = errors.Wrap(err, "failed to save account snapshots to database")
			continue
		}

		// update account cache, non critical and should not fail the flow
		_, err = dbClient.UpsertAccounts(ctx, []models.Account{accountUpdates[accountSnapshot.AccountID]})
		if err != nil {
			log.Err(err).Str("accountId", accountSnapshot.AccountID).Msg("failed to upsert account")
		}
	}

	return accountErrors, nil
}
