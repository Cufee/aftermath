package fetch

import (
	"context"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type multiSourceClient struct {
	retriesPerRequest  int
	retrySleepInterval time.Duration

	database   *database.Client
	wargaming  wargaming.Client
	blitzstars blitzstars.Client
}

func NewMultiSourceClient(wargaming wargaming.Client, blitzstars blitzstars.Client, database *database.Client) (Client, error) {
	mc := multiSourceClient{
		database:   database,
		wargaming:  wargaming,
		blitzstars: blitzstars,

		retriesPerRequest:  2,
		retrySleepInterval: time.Millisecond * 100,
	}
	return &mc, nil
}

/*
Get live account stats from wargaming
  - Each request will be retried c.retriesPerRequest times
  - This function will assume the player ID is valid and optimistically run all request concurrently, returning the first error once all request finish
*/
func (c *multiSourceClient) CurrentStats(ctx context.Context, id string) (AccountStatsOverPeriod, error) {
	realm := c.wargaming.RealmFromAccountID(id)

	var group sync.WaitGroup
	group.Add(3)

	var clan types.ClanMember
	var account DataWithErr[types.ExtendedAccount]
	var vehicles DataWithErr[[]types.VehicleStatsFrame]

	go func() {
		defer group.Done()
		account = withRetry(func() (types.ExtendedAccount, error) { return c.wargaming.AccountByID(ctx, realm, id) }, c.retriesPerRequest, c.retrySleepInterval)
	}()
	go func() {
		defer group.Done()
		// we should not retry this request since it is not critical and will fail on accounts without a clan
		clan, _ = c.wargaming.AccountClan(ctx, realm, id)
	}()
	go func() {
		defer group.Done()
		vehicles = withRetry(func() ([]types.VehicleStatsFrame, error) { return c.wargaming.AccountVehicles(ctx, realm, id) }, c.retriesPerRequest, c.retrySleepInterval)
	}()

	group.Wait()

	if account.Err != nil {
		return AccountStatsOverPeriod{}, account.Err
	}
	if vehicles.Err != nil {
		return AccountStatsOverPeriod{}, vehicles.Err
	}

	// TODO: Get tank averages for WN8

	return wargamingToStats(account.Data, clan, vehicles.Data), nil
}

func (c *multiSourceClient) PeriodStats(ctx context.Context, id string, periodStart time.Time) (AccountStatsOverPeriod, error) {
	var histories DataWithErr[map[int][]blitzstars.TankHistoryEntry]
	var averages DataWithErr[map[string]frame.StatsFrame]
	var current DataWithErr[AccountStatsOverPeriod]

	var group sync.WaitGroup
	group.Add(1)
	go func() {
		defer group.Done()
		// Get current stats
		stats, err := c.CurrentStats(ctx, id)
		current = DataWithErr[AccountStatsOverPeriod]{stats, err}
		if err != nil {
			return
		}

		// Get vehicle averages from db, a list of tanks is required for a more optimal query here
		var vehicles []string
		for _, vehicle := range stats.RegularBattles.Vehicles {
			vehicles = append(vehicles, vehicle.VehicleID)
		}
		data, err := c.database.GetVehicleAverages(ctx, vehicles)
		averages = DataWithErr[map[string]frame.StatsFrame]{data, err}
	}()

	// TODO: Lookup a session from the database first

	// Return career stats if stats are requested for 90+ days, blitzstars and aftermath do not track that far
	if time.Since(periodStart).Hours()/24 >= 90 {
		group.Wait()
		if current.Err != nil {
			return AccountStatsOverPeriod{}, current.Err
		}

		// TODO: Set wn8 for all vehicles

		return current.Data, nil
	}

	group.Add(1)
	go func() {
		defer group.Done()
		histories = withRetry(func() (map[int][]blitzstars.TankHistoryEntry, error) {
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
	current.Data.RegularBattles = blitzstarsToStats(current.Data.RegularBattles.Vehicles, histories.Data, averages.Data, periodStart)
	return current.Data, nil
}
