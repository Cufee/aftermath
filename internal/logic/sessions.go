package logic

import (
	"context"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/aftermath/internal/stats/fetch"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type vehicleResponseData map[string][]types.VehicleStatsFrame

func RecordAccountSnapshots(ctx context.Context, wgClient wargaming.Client, dbClient database.Client, realm string, force bool, accountIDs ...string) (map[string]error, error) {
	if len(accountIDs) < 1 {
		return nil, nil
	}

	createdAt := time.Now()
	accounts, err := wgClient.BatchAccountByID(ctx, realm, accountIDs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch accounts")
	}

	// existing snaphsots for accounts
	existingSnapshots, err := dbClient.GetManyAccountSnapshots(ctx, accountIDs, database.SnapshotTypeDaily)
	if err != nil && !db.IsErrNotFound(err) {
		return nil, errors.Wrap(err, "failed to get existing snapshots")
	}
	existingSnapshotsMap := make(map[string]*database.AccountSnapshot)
	for _, snapshot := range existingSnapshots {
		existingSnapshotsMap[snapshot.AccountID] = &snapshot
	}

	// make a new slice just in case some accounts were not returned/are private
	var validAccouts []string
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
		if data.LastBattleTime == 0 {
			continue
		}
		if s, ok := existingSnapshotsMap[id]; !force && ok && data.LastBattleTime == int(s.LastBattleTime.Unix()) {
			// last snapshot is the same, we can skip it
			continue
		}
		validAccouts = append(validAccouts, id)
	}

	if len(validAccouts) == 0 {
		return nil, nil
	}

	// clans are options-ish
	clans, err := wgClient.BatchAccountClan(ctx, realm, validAccouts)
	if err != nil {
		log.Err(err).Msg("failed to get batch account clans")
		clans = make(map[string]types.ClanMember)
	}

	vehicleCh := make(chan retry.DataWithErr[vehicleResponseData], len(validAccouts))
	var group sync.WaitGroup
	group.Add(len(validAccouts))
	for _, id := range validAccouts {
		go func(id string) {
			defer group.Done()
			data, err := wgClient.AccountVehicles(ctx, realm, id)
			vehicleCh <- retry.DataWithErr[vehicleResponseData]{Data: vehicleResponseData{id: data}, Err: err}
		}(id)
	}
	group.Wait()
	close(vehicleCh)

	var accountErrors = make(map[string]error)
	var accountUpdates []database.Account
	var snapshots []database.AccountSnapshot
	var vehicleSnapshots []database.VehicleSnapshot
	for result := range vehicleCh {
		// there is only 1 key in this map
		for id, vehicles := range result.Data {
			if result.Err != nil {
				accountErrors[id] = result.Err
				continue
			}
			existingSnapshots, err := dbClient.GetVehicleSnapshots(ctx, id, id, database.SnapshotTypeDaily)
			if err != nil && !db.IsErrNotFound(err) {
				accountErrors[id] = err
				continue
			}
			existingSnapshotsMap := make(map[string]*database.VehicleSnapshot)
			for _, snapshot := range existingSnapshots {
				existingSnapshotsMap[snapshot.VehicleID] = &snapshot
			}

			stats := fetch.WargamingToStats(realm, accounts[id], clans[id], vehicles)
			accountUpdates = append(accountUpdates, database.Account{
				Realm:    stats.Realm,
				ID:       stats.Account.ID,
				Nickname: stats.Account.Nickname,

				Private:        false,
				CreatedAt:      stats.Account.CreatedAt,
				LastBattleTime: stats.LastBattleTime,

				ClanID:  stats.Account.ClanID,
				ClanTag: stats.Account.ClanTag,
			})
			snapshots = append(snapshots, database.AccountSnapshot{
				Type:           database.SnapshotTypeDaily,
				CreatedAt:      createdAt,
				AccountID:      stats.Account.ID,
				ReferenceID:    stats.Account.ID,
				LastBattleTime: stats.LastBattleTime,
				RatingBattles:  stats.RatingBattles.StatsFrame,
				RegularBattles: stats.RegularBattles.StatsFrame,
			})

			if len(vehicles) < 1 {
				continue
			}
			vehicleStats := fetch.WargamingVehiclesToFrame(vehicles)
			for id, vehicle := range vehicleStats {
				if vehicle.LastBattleTime.Unix() < 1 {
					// vehicle was never played
					continue
				}
				if s, ok := existingSnapshotsMap[id]; !force && ok && s.Stats.Battles == vehicle.Battles {
					// last snapshot is the same, we can skip it
					continue
				}

				vehicleSnapshots = append(vehicleSnapshots, database.VehicleSnapshot{
					CreatedAt:      createdAt,
					Type:           database.SnapshotTypeDaily,
					LastBattleTime: vehicle.LastBattleTime,
					AccountID:      stats.Account.ID,
					VehicleID:      vehicle.VehicleID,
					ReferenceID:    stats.Account.ID,
					Stats:          *vehicle.StatsFrame,
				})
			}
		}
	}

	err = dbClient.CreateAccountSnapshots(ctx, snapshots...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save account snapshots to database")
	}

	err = dbClient.CreateVehicleSnapshots(ctx, vehicleSnapshots...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to save vehicle snapshots to database")
	}

	return accountErrors, nil
}
