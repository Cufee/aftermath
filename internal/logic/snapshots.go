package logic

import (
	"context"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/utils"

	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

var (
	accountsPool         = utils.NewPool[models.Account]()
	vehicleSnapshotsPool = utils.NewPool[models.VehicleSnapshot]()
	accountSnapshotsPool = utils.NewPool[models.AccountSnapshot]()
)

type snapshotRecordOptions struct {
	force bool
	kind  models.SnapshotType
}

type SnapshotRecordOption func(*snapshotRecordOptions)

func WithForceUpdate() SnapshotRecordOption {
	return func(sro *snapshotRecordOptions) {
		sro.force = true
	}
}

func WithType(t models.SnapshotType) SnapshotRecordOption {
	return func(sro *snapshotRecordOptions) {
		sro.kind = t
	}
}

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

/*
Filter passed in accounts and return active account ids
  - an account is considered active if it has played a battle since the last snapshot, or has no snapshots
*/
func filterActiveAccounts(ctx context.Context, dbClient database.Client, input AccountsWithReference, accounts map[string]types.ExtendedAccount, options snapshotRecordOptions) ([]string, error) {
	var ids []string
	for id := range accounts {
		ids = append(ids, id)
	}

	// existing snapshots for accounts
	var dbOpts []database.Query
	dbOpts = append(dbOpts, database.WithReferenceIDIn(input.ReferenceIDs()...))

	existingLastBattleTimes, err := dbClient.GetAccountLastBattleTimes(ctx, ids, options.kind, dbOpts...)
	if err != nil && !database.IsNotFound(err) {
		return nil, errors.Wrap(err, "failed to get existing snapshots")
	}

	// make a new slice just in case some accounts were not returned/are private
	var needAnUpdate []string
	for id, data := range accounts {
		if data.LastBattleTime < 1 {
			continue
		}
		if s, ok := existingLastBattleTimes[id]; !options.force && (ok && !s.IsZero() && data.LastBattleTime == int(s.Unix())) {
			// last snapshot is the same, we can skip it
			continue
		}
		needAnUpdate = append(needAnUpdate, id)
	}
	return needAnUpdate, nil
}

func RecordAccountSnapshots(ctx context.Context, wgClient wargaming.Client, dbClient database.Client, realm types.Realm, input AccountsWithReference, opts ...SnapshotRecordOption) (map[string]error, error) {
	if len(input) < 1 {
		return nil, nil
	}

	options := snapshotRecordOptions{kind: models.SnapshotTypeDaily, force: false}
	for _, apply := range opts {
		apply(&options)
	}

	createdAt := time.Now()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	accounts, err := wgClient.BatchAccountByID(ctx, realm, input.AccountIDs())
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch accounts")
	}

	accountsNeedAnUpdate, err := filterActiveAccounts(ctx, dbClient, input, accounts, options)
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

	var accountVehiclesBTMx sync.Mutex
	accountVehiclesBT := make(map[string]retry.DataWithErr[map[string]time.Time], len(accountsNeedAnUpdate))

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
		data, _ := wgClient.BatchAccountClan(ctx, realm, accountsNeedAnUpdate)
		clans = data
	}()

	// get account vehicle last battle times
	for _, accountID := range accountsNeedAnUpdate {
		group.Add(1)
		go func(id string) {
			defer group.Done()
			lastBattles, err := dbClient.GetVehiclesLastBattleTimes(ctx, id, nil, options.kind)
			accountVehiclesBTMx.Lock()
			accountVehiclesBT[id] = retry.DataWithErr[map[string]time.Time]{Data: lastBattles, Err: err}
			defer accountVehiclesBTMx.Unlock()
		}(accountID)
	}

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

		existingLastBattleTimes := accountVehiclesBT[accountID]
		if existingLastBattleTimes.Err != nil {
			accountErrors[accountID] = existingLastBattleTimes.Err
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
				Type:           options.kind,
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
		vehicleStats := fetch.WargamingVehiclesToFrame(vehicles.Data)
		if len(vehicles.Data) > 0 {
			for id, vehicle := range vehicleStats {
				if !options.force && existingLastBattleTimes.Data[id].Unix() >= vehicle.LastBattleTime.Unix() {
					// skip vehicles that were not played
					continue
				}

				sht := vehicleSnapshotsPool.Get()
				defer vehicleSnapshotsPool.Put(sht)

				*sht = models.VehicleSnapshot{
					Type:           options.kind,
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

outer:
	for _, accountSnapshot := range accountSnapshots {
		// update account cache
		aErr, err := dbClient.UpsertAccounts(ctx, accountUpdates[accountSnapshot.AccountID])
		if err != nil {
			log.Err(err).Str("accountId", accountSnapshot.AccountID).Msg("failed to upsert account")
		}
		for _, err := range aErr {
			log.Err(err).Str("accountId", accountSnapshot.AccountID).Msg("failed to upsert account")
			accountErrors[accountSnapshot.AccountID] = err
			continue outer
		}

		// save account snapshot
		err = dbClient.CreateAccountSnapshots(ctx, accountSnapshot)
		if err != nil {
			accountErrors[accountSnapshot.AccountID] = errors.Wrap(err, "failed to save account snapshots to database")
			continue
		}

		// save all vehicle snapshots)
		err = dbClient.CreateVehicleSnapshots(ctx, vehicleSnapshots[accountSnapshot.AccountID]...)
		if err != nil {
			accountErrors[accountSnapshot.AccountID] = err
			continue
		}
	}

	return accountErrors, nil
}
