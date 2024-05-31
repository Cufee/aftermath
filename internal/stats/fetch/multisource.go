package fetch

import (
	"context"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type multiSourceClient struct {
	retriesPerRequest  int
	retrySleepInterval time.Duration

	wargaming  wargaming.Client
	blitzstars blitzstars.Client
}

func NewMultiSourceClient(wargaming wargaming.Client, blitzstars blitzstars.Client) (Client, error) {
	mc := multiSourceClient{
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
	current, err := c.CurrentStats(ctx, id)
	if err != nil {
		return AccountStatsOverPeriod{}, err
	}

	// TODO: Get tank averages for WN8

	// TODO: Lookup a session from the database first

	// Return career stats if stats are requested for 90+ days, blitzstars and aftermath do not track that far
	if time.Since(periodStart).Hours()/24 >= 90 {
		// TODO: Set wn8 for all vehicles
		return current, nil
	}

	// blitzstars does not provide rating battle data
	periodStats := AccountStatsOverPeriod{
		Account:        current.Account,
		Clan:           current.Clan,
		PeriodStart:    periodStart,
		PeriodEnd:      current.LastBattleTime,
		LastBattleTime: current.LastBattleTime,
		RegularBattles: StatsWithVehicles{Vehicles: make(map[string]VehicleStatsFrame)},
		RatingBattles:  StatsWithVehicles{Vehicles: make(map[string]VehicleStatsFrame)},
	}

	histories := withRetry(func() (map[int][]blitzstars.TankHistoryEntry, error) {
		return c.blitzstars.AccountTankHistories(ctx, id)
	}, c.retriesPerRequest, c.retrySleepInterval)
	if histories.Err != nil {
		return AccountStatsOverPeriod{}, histories.Err
	}

	for _, vehicle := range current.RegularBattles.Vehicles {
		if vehicle.LastBattleTime.Before(periodStart) {
			continue
		}

		id, err := strconv.Atoi(vehicle.VehicleID)
		if err != nil || id == 0 {
			continue
		}

		entries := histories.Data[id]
		// Sort entries by number of battles in descending order
		slices.SortFunc(entries, func(i, j blitzstars.TankHistoryEntry) int {
			return j.Stats.Battles - i.Stats.Battles
		})

		var selectedEntry blitzstars.TankHistoryEntry
		for _, entry := range entries {
			if entry.LastBattleTime < int(periodStats.PeriodStart.Unix()) {
				selectedEntry = entry
				break
			}
		}

		if selectedEntry.Stats.Battles < int(vehicle.Battles) {
			selectedFrame := wargamingToFrame(selectedEntry.Stats)
			vehicle.StatsFrame.Subtract(selectedFrame)

			// TODO: add WN8

			periodStats.RegularBattles.Vehicles[vehicle.VehicleID] = vehicle
			periodStats.RegularBattles.Add(*vehicle.StatsFrame)
		}
	}

	return periodStats, nil
}
