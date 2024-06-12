package fetch

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/rs/zerolog/log"
)

var _ Client = &multiSourceClient{}

type multiSourceClient struct {
	retriesPerRequest  int
	retrySleepInterval time.Duration

	database   database.Client
	wargaming  wargaming.Client
	blitzstars blitzstars.Client
}

func NewMultiSourceClient(wargaming wargaming.Client, blitzstars blitzstars.Client, database database.Client) (*multiSourceClient, error) {
	return &multiSourceClient{
		database:   database,
		wargaming:  wargaming,
		blitzstars: blitzstars,

		retriesPerRequest:  2,
		retrySleepInterval: time.Millisecond * 100,
	}, nil
}

func (c *multiSourceClient) Search(ctx context.Context, nickname, realm string) (types.Account, error) {
	accounts, err := c.wargaming.SearchAccounts(ctx, realm, nickname)
	if err != nil {
		return types.Account{}, err
	}
	if len(accounts) < 1 {
		return types.Account{}, errors.New("no results found")
	}

	return accounts[0], nil
}

/*
Gets account info from wg and updates cache
*/
func (c *multiSourceClient) Account(ctx context.Context, id string) (database.Account, error) {
	realm := c.wargaming.RealmFromAccountID(id)

	var group sync.WaitGroup
	group.Add(2)

	var clan types.ClanMember
	var wgAccount retry.DataWithErr[types.ExtendedAccount]

	go func() {
		defer group.Done()

		wgAccount = retry.Retry(func() (types.ExtendedAccount, error) {
			return c.wargaming.AccountByID(ctx, realm, id)
		}, c.retriesPerRequest, c.retrySleepInterval)
	}()
	go func() {
		defer group.Done()
		// we should not retry this request since it is not critical and will fail on accounts without a clan
		clan, _ = c.wargaming.AccountClan(ctx, realm, id)
	}()

	group.Wait()

	if wgAccount.Err != nil {
		return database.Account{}, wgAccount.Err
	}

	account := wargamingToAccount(realm, wgAccount.Data, clan, false)
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		errMap := c.database.UpsertAccounts(ctx, []database.Account{account})
		if err := errMap[id]; err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
	}()

	return account, nil
}

/*
Get live account stats from wargaming
  - Each request will be retried c.retriesPerRequest times
  - This function will assume the player ID is valid and optimistically run all request concurrently, returning the first error once all request finish
*/
func (c *multiSourceClient) CurrentStats(ctx context.Context, id string, opts ...statsOption) (AccountStatsOverPeriod, error) {
	var options statsOptions
	for _, apply := range opts {
		apply(&options)
	}

	realm := c.wargaming.RealmFromAccountID(id)

	var group sync.WaitGroup
	group.Add(3)

	var clan types.ClanMember
	var account retry.DataWithErr[types.ExtendedAccount]
	var vehicles retry.DataWithErr[[]types.VehicleStatsFrame]
	var averages retry.DataWithErr[map[string]frame.StatsFrame]

	go func() {
		defer group.Done()

		account = retry.Retry(func() (types.ExtendedAccount, error) {
			return c.wargaming.AccountByID(ctx, realm, id)
		}, c.retriesPerRequest, c.retrySleepInterval)
	}()
	go func() {
		defer group.Done()
		// we should not retry this request since it is not critical and will fail on accounts without a clan
		clan, _ = c.wargaming.AccountClan(ctx, realm, id)
	}()
	go func() {
		defer group.Done()

		vehicles = retry.Retry(func() ([]types.VehicleStatsFrame, error) {
			return c.wargaming.AccountVehicles(ctx, realm, id)
		}, c.retriesPerRequest, c.retrySleepInterval)

		if vehicles.Err != nil || len(vehicles.Data) < 1 || !options.withWN8 {
			return
		}

		var ids []string
		for _, v := range vehicles.Data {
			ids = append(ids, fmt.Sprint(v.TankID))
		}
		a, err := c.database.GetVehicleAverages(ctx, ids)
		averages = retry.DataWithErr[map[string]frame.StatsFrame]{Data: a, Err: err}
	}()

	group.Wait()

	if account.Err != nil {
		return AccountStatsOverPeriod{}, account.Err
	}
	if vehicles.Err != nil {
		return AccountStatsOverPeriod{}, vehicles.Err
	}
	if averages.Err != nil {
		// not critical, this will only affect WN8
		log.Err(averages.Err).Msg("failed to get tank averages")
	}

	stats := WargamingToStats(realm, account.Data, clan, vehicles.Data)
	if options.withWN8 {
		stats.AddWN8(averages.Data)
	}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		errMap := c.database.UpsertAccounts(ctx, []database.Account{stats.Account})
		if err := errMap[id]; err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
	}()

	return stats, nil
}

func (c *multiSourceClient) PeriodStats(ctx context.Context, id string, periodStart time.Time, opts ...statsOption) (AccountStatsOverPeriod, error) {
	var options statsOptions
	for _, apply := range opts {
		apply(&options)
	}

	var histories retry.DataWithErr[map[int][]blitzstars.TankHistoryEntry]
	var averages retry.DataWithErr[map[string]frame.StatsFrame]
	var current retry.DataWithErr[AccountStatsOverPeriod]

	var group sync.WaitGroup
	group.Add(1)
	go func() {
		defer group.Done()

		stats, err := c.CurrentStats(ctx, id)
		current = retry.DataWithErr[AccountStatsOverPeriod]{Data: stats, Err: err}

		if err != nil || stats.RegularBattles.Battles < 1 || !options.withWN8 {
			return
		}

		var ids []string
		for id := range stats.RegularBattles.Vehicles {
			ids = append(ids, id)
		}
		a, err := c.database.GetVehicleAverages(ctx, ids)
		averages = retry.DataWithErr[map[string]frame.StatsFrame]{Data: a, Err: err}
	}()

	// TODO: lookup a session from the database first
	// if a session exists in the database, we don't need BlitzStars and have better data

	// return career stats if stats are requested for 0 or 90+ days, we do not track that far
	if days := time.Since(periodStart).Hours() / 24; days > 90 || days < 1 {
		group.Wait()
		if current.Err != nil {
			return AccountStatsOverPeriod{}, current.Err
		}
		if averages.Err != nil {
			// not critical, this will only affect WN8
			log.Err(averages.Err).Msg("failed to get tank averages")
		}

		stats := current.Data
		if options.withWN8 {
			stats.AddWN8(averages.Data)
		}
		return stats, nil
	}

	group.Add(1)
	go func() {
		defer group.Done()
		histories = retry.Retry(func() (map[int][]blitzstars.TankHistoryEntry, error) {
			return c.blitzstars.AccountTankHistories(ctx, id)
		}, c.retriesPerRequest, c.retrySleepInterval)
	}()

	// wait for all requests to finish and check errors
	group.Wait()
	if current.Err != nil {
		return AccountStatsOverPeriod{}, current.Err
	}
	if histories.Err != nil {
		return AccountStatsOverPeriod{}, histories.Err
	}
	if averages.Err != nil {
		// not critical, this will only affect WN8
		log.Err(averages.Err).Msg("failed to get tank averages")
	}

	current.Data.PeriodEnd = time.Now()
	current.Data.PeriodStart = periodStart
	current.Data.RatingBattles = StatsWithVehicles{} // blitzstars do not provide rating battles stats
	current.Data.RegularBattles = blitzstarsToStats(current.Data.RegularBattles.Vehicles, histories.Data, periodStart)

	stats := current.Data
	if options.withWN8 {
		stats.AddWN8(averages.Data)
	}
	return stats, nil
}

func (c *multiSourceClient) SessionStats(ctx context.Context, id string, sessionStart time.Time, opts ...statsOption) (AccountStatsOverPeriod, error) {
	var options statsOptions
	for _, apply := range opts {
		apply(&options)
	}

	// sessions and period stats are tracked differenttly
	// we look up the latest session _before_ sessionBefore, not after sessionStart
	if sessionStart.IsZero() {
		sessionStart = time.Now()
	}
	if days := time.Since(sessionStart).Hours() / 24; days > 90 {
		return AccountStatsOverPeriod{}, ErrInvalidSessionStart
	}

	sessionBefore := sessionStart
	if time.Since(sessionStart).Hours() >= 24 {
		// 3 days would mean today and yest, so before the reset 2 days ago and so on
		sessionBefore = sessionBefore.Add(time.Hour * 24 * -1)
	}

	var accountSnapshot retry.DataWithErr[database.AccountSnapshot]
	var vehiclesSnapshots retry.DataWithErr[[]database.VehicleSnapshot]
	var averages retry.DataWithErr[map[string]frame.StatsFrame]
	var current retry.DataWithErr[AccountStatsOverPeriod]

	var group sync.WaitGroup
	group.Add(1)
	go func() {
		defer group.Done()

		stats, err := c.CurrentStats(ctx, id)
		current = retry.DataWithErr[AccountStatsOverPeriod]{Data: stats, Err: err}

		if err != nil || stats.RegularBattles.Battles < 1 || !options.withWN8 {
			return
		}

		var ids []string
		for id := range stats.RegularBattles.Vehicles {
			ids = append(ids, id)
		}
		a, err := c.database.GetVehicleAverages(ctx, ids)
		averages = retry.DataWithErr[map[string]frame.StatsFrame]{Data: a, Err: err}
	}()

	group.Add(1)
	go func() {
		defer group.Done()
		s, err := c.database.GetAccountSnapshot(ctx, id, id, database.SnapshotTypeDaily, database.WithCreatedBefore(sessionBefore))
		accountSnapshot = retry.DataWithErr[database.AccountSnapshot]{Data: s, Err: err}
		v, err := c.database.GetVehicleSnapshots(ctx, id, id, database.SnapshotTypeDaily, database.WithCreatedBefore(sessionBefore))
		vehiclesSnapshots = retry.DataWithErr[[]database.VehicleSnapshot]{Data: v, Err: err}
	}()

	// wait for all requests to finish and check errors
	group.Wait()
	if current.Err != nil {
		return AccountStatsOverPeriod{}, current.Err
	}
	if accountSnapshot.Err != nil {
		if db.IsErrNotFound(accountSnapshot.Err) {
			return AccountStatsOverPeriod{}, ErrSessionNotFound
		}
		return AccountStatsOverPeriod{}, accountSnapshot.Err
	}
	if averages.Err != nil {
		// not critical, this will only affect WN8
		log.Err(averages.Err).Msg("failed to get tank averages")
	}

	stats := current.Data
	stats.PeriodEnd = time.Now()
	stats.PeriodStart = sessionStart
	stats.RatingBattles.StatsFrame.Subtract(accountSnapshot.Data.RatingBattles)
	stats.RegularBattles.StatsFrame.Subtract(accountSnapshot.Data.RegularBattles)

	snapshotsMap := make(map[string]int, len(vehiclesSnapshots.Data))
	for i, data := range vehiclesSnapshots.Data {
		snapshotsMap[data.VehicleID] = i
	}

	stats.RegularBattles.Vehicles = make(map[string]frame.VehicleStatsFrame, len(current.Data.RegularBattles.Vehicles))
	for id, current := range current.Data.RegularBattles.Vehicles {
		snapshotIndex, exists := snapshotsMap[id]
		if !exists {
			stats.RegularBattles.Vehicles[id] = current
			continue
		}

		snapshot := vehiclesSnapshots.Data[snapshotIndex]
		if current.Battles == 0 || current.Battles == snapshot.Stats.Battles {
			continue
		}
		current.StatsFrame.Subtract(snapshot.Stats)
		stats.RegularBattles.Vehicles[id] = current
	}

	if options.withWN8 {
		stats.AddWN8(averages.Data)
	}
	return stats, nil

}

func (c *multiSourceClient) CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error) {
	return c.blitzstars.CurrentTankAverages(ctx)
}
