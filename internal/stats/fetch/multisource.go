package fetch

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
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

	stats := wargamingToStats(realm, account.Data, clan, vehicles.Data)
	if options.withWN8 {
		stats.AddWN8(averages.Data)
	}

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

	current.Data.PeriodStart = periodStart
	current.Data.RatingBattles = StatsWithVehicles{} // blitzstars do not provide rating battles stats
	current.Data.RegularBattles = blitzstarsToStats(current.Data.RegularBattles.Vehicles, histories.Data, periodStart)

	stats := current.Data
	if options.withWN8 {
		stats.AddWN8(averages.Data)
	}
	return stats, nil
}

func (c *multiSourceClient) CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error) {
	return c.blitzstars.CurrentTankAverages(ctx)
}
