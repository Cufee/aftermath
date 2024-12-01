package logic

import (
	"context"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"

	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

var (
	accountsPool         = newPool[models.Account]()
	vehicleSnapshotsPool = newPool[models.VehicleSnapshot]()
	accountSnapshotsPool = newPool[models.AccountSnapshot]()
)

type AccountsWithReference map[string]string

func WithReference(accountID string, referenceID string) AccountsWithReference {
	var d AccountsWithReference = make(AccountsWithReference)
	d[accountID] = referenceID
	return d
}

func (d AccountsWithReference) ReferenceIDs() []string {
	var slice []string
	for _, reference := range d {
		slice = append(slice, reference)
	}
	return slice
}

func (d AccountsWithReference) AccountIDs() []string {
	var slice []string
	for id := range d {
		slice = append(slice, id)
	}
	return slice
}

func WithDefaultReference(ids []string) AccountsWithReference {
	var d AccountsWithReference = make(AccountsWithReference)
	for _, id := range ids {
		d[id] = id
	}
	return d
}

type vehicleBattleData struct {
	LastBattle time.Time
	Battles    int
}

/*
Filter passed in accounts and return active account ids
  - an account is considered active if it has played a battle since the last snapshot, or has no snapshots
*/
func filterActiveAccounts(ctx context.Context, dbClient database.Client, input AccountsWithReference, accounts map[string]types.ExtendedAccount, force bool) ([]string, error) {
	var ids []string
	for id := range accounts {
		ids = append(ids, id)
	}

	// existing snapshots for accounts
	var opts []database.Query
	opts = append(opts, database.WithReferenceIDIn(input.ReferenceIDs()...))

	existingLastBattleTimes, err := dbClient.GetAccountLastBattleTimes(ctx, ids, models.SnapshotTypeDaily, opts...)
	if err != nil && !database.IsNotFound(err) {
		return nil, errors.Wrap(err, "failed to get existing snapshots")
	}

	// make a new slice just in case some accounts were not returned/are private
	var needAnUpdate []string
	for id, data := range accounts {
		if data.LastBattleTime < 1 {
			continue
		}
		if s, ok := existingLastBattleTimes[id]; !force && (ok && !s.IsZero() && data.LastBattleTime == int(s.Unix())) {
			// last snapshot is the same, we can skip it
			continue
		}
		needAnUpdate = append(needAnUpdate, id)
	}
	return needAnUpdate, nil
}

func RecordAccountSnapshots(ctx context.Context, wgClient wargaming.Client, dbClient database.Client, realm string, force bool, input AccountsWithReference) (map[string]error, error) {
	if len(input) < 1 {
		return nil, nil
	}
	createdAt := time.Now()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	accounts, err := wgClient.BatchAccountByID(ctx, realm, input.AccountIDs())
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch accounts")
	}

	accountsNeedAnUpdate, err := filterActiveAccounts(ctx, dbClient, input, accounts, force)
	if err != nil {
		return nil, err
	}
	if len(accountsNeedAnUpdate) < 1 {
		return nil, nil
	}

	var group sync.WaitGroup
	var clans = make(map[string]types.ClanMember)

	var accountVehiclesMx sync.Mutex
	accountVehicles := make(map[string]retry.DataWithErr[[]types.VehicleStatsFrame], len(accountsNeedAnUpdate))

	for _, id := range accountsNeedAnUpdate {
		group.Add(1)
		// get account vehicle stats
		go func(id string) {
			defer group.Done()
			data, err := wgClient.AccountVehicles(ctx, realm, id, nil) // nil will return all vehicles

			accountVehiclesMx.Lock()
			accountVehicles[id] = retry.DataWithErr[[]types.VehicleStatsFrame]{Data: data, Err: err}
			accountVehiclesMx.Unlock()
		}(id)
	}

	group.Add(1)
	// get account clans, not critical
	go func() {
		defer group.Done()
		// clans are optional-ish
		data, err := wgClient.BatchAccountClan(ctx, realm, accountsNeedAnUpdate)
		if err != nil && err.Error() != "SOURCE_NOT_AVAILABLE" {
			log.Err(err).Msg("failed to get batch account clans")
		}
		clans = data
	}()

	group.Wait()

	var accountErrors = make(map[string]error)
	var accountUpdates = make(map[string]*models.Account)

	var accountSnapshots []*models.AccountSnapshot
	var vehicleSnapshots = make(map[string][]*models.VehicleSnapshot)

	for _, accountID := range accountsNeedAnUpdate {
		vehicles := accountVehicles[accountID]
		if vehicles.Err != nil {
			accountErrors[accountID] = vehicles.Err
			continue
		}

		accountRefID := input[accountID]
		if accountRefID == "" {
			accountRefID = accountID
		}

		snapshotStats := fetch.WargamingToStats(realm, accounts[accountID], clans[accountID], vehicles.Data)
		{ // account snapshot
			sht := accountSnapshotsPool.Get()
			defer accountSnapshotsPool.Put(sht)

			*sht = models.AccountSnapshot{
				Type:           models.SnapshotTypeDaily,
				CreatedAt:      createdAt,
				ReferenceID:    accountRefID,
				AccountID:      snapshotStats.Account.ID,
				LastBattleTime: snapshotStats.LastBattleTime,
				RatingBattles:  snapshotStats.RatingBattles.StatsFrame,
				RegularBattles: snapshotStats.RegularBattles.StatsFrame,
			}
			accountSnapshots = append(accountSnapshots, sht)

			asht := accountsPool.Get()
			defer accountsPool.Put(asht)

			*asht = snapshotStats.Account
			accountUpdates[accountID] = asht
		}

		// vehicle snapshots
		var vehicleLastBattleTimes = make(map[string]vehicleBattleData)
		vehicleStats := fetch.WargamingVehiclesToFrame(vehicles.Data)
		if len(vehicles.Data) > 0 {
			for id, vehicle := range vehicleStats {
				vehicleLastBattleTimes[id] = vehicleBattleData{vehicle.LastBattleTime, int(vehicle.Battles.Float())}

				sht := vehicleSnapshotsPool.Get()
				defer vehicleSnapshotsPool.Put(sht)

				*sht = models.VehicleSnapshot{
					Type:           models.SnapshotTypeDaily,
					LastBattleTime: vehicle.LastBattleTime,
					Stats:          *vehicle.StatsFrame,
					VehicleID:      vehicle.VehicleID,
					ReferenceID:    accountRefID,
					AccountID:      accountID,
					CreatedAt:      createdAt,
				}
				vehicleSnapshots[accountID] = append(vehicleSnapshots[accountID], sht)
			}
		}
	}

	for _, accountSnapshot := range accountSnapshots {
		vehicles := vehicleSnapshots[accountSnapshot.AccountID]
		if len(vehicles) > 0 {
			// save all vehicle snapshots)
			err := dbClient.CreateAccountVehicleSnapshots(ctx, accountSnapshot.AccountID, vehicles...)
			if err != nil {
				accountErrors[accountSnapshot.AccountID] = err
				continue
			}
		}

		// save account snapshot
		err := dbClient.CreateAccountSnapshots(ctx, accountSnapshot)
		if err != nil {
			accountErrors[accountSnapshot.AccountID] = errors.Wrap(err, "failed to save account snapshots to database")
			continue
		}

		// update account cache, non critical and should not fail the flow
		_, err = dbClient.UpsertAccounts(ctx, accountUpdates[accountSnapshot.AccountID])
		if err != nil {
			log.Err(err).Str("accountId", accountSnapshot.AccountID).Msg("failed to upsert account")
		}
	}

	return accountErrors, nil
}
