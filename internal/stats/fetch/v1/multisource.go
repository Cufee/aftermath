package fetch

import (
	"context"
	"fmt"
	"io"
	"slices"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"

	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/external/wargaming"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/retry"
	"github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
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

func (c *multiSourceClient) Search(ctx context.Context, nickname string, realm types.Realm, limit int) (types.Account, error) {
	if !wargaming.ValidatePlayerNickname(nickname) {
		return types.Account{}, ErrInvalidSearch
	}

	accounts, err := c.wargaming.SearchAccounts(ctx, realm, nickname, types.WithLimit(limit))
	if err != nil {
		return types.Account{}, parseWargamingError(err)
	}
	if len(accounts) < 1 {
		return types.Account{}, ErrAccountNotFound
	}

	return accounts[0], nil
}

func (c *multiSourceClient) BroadSearch(ctx context.Context, nickname string, limitPerRealm int) ([]AccountWithRealm, error) {
	data := make(chan AccountWithRealm, 9)
	errors := make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		accounts, err := c.wargaming.SearchAccounts(ctx, types.RealmNorthAmerica, nickname, types.WithLimit(limitPerRealm))
		if err != nil {
			errors <- err
			return
		}
		for _, a := range accounts {
			data <- AccountWithRealm{a, types.RealmNorthAmerica}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		accounts, err := c.wargaming.SearchAccounts(ctx, types.RealmEurope, nickname, types.WithLimit(limitPerRealm))
		if err != nil {
			errors <- err
			return
		}
		for _, a := range accounts {
			data <- AccountWithRealm{a, types.RealmEurope}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		accounts, err := c.wargaming.SearchAccounts(ctx, types.RealmAsia, nickname, types.WithLimit(limitPerRealm))
		if err != nil {
			errors <- err
			return
		}
		for _, a := range accounts {
			data <- AccountWithRealm{a, types.RealmAsia}
		}
	}()
	wg.Wait()
	close(data)
	close(errors)

	// return first error
	if len(data) == 0 && len(errors) > 0 {
		for err := range errors {
			if err != nil {
				return nil, parseWargamingError(err)
			}
		}
	}

	var accounts []AccountWithRealm
	for a := range data {
		accounts = append(accounts, a)
	}

	return accounts, nil
}

/*
Gets account info from wg and updates cache
*/
func (c *multiSourceClient) Account(ctx context.Context, id string) (models.Account, error) {
	realm, ok := c.wargaming.RealmFromID(id)
	if !ok {
		return models.Account{}, wargaming.ErrRealmNotSupported
	}

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
		return models.Account{}, parseWargamingError(wgAccount.Err)
	}

	account := WargamingToAccount(realm, wgAccount.Data, clan, false)
	go func(account models.Account) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		aErr, err := c.database.UpsertAccounts(ctx, &account)
		if err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
		if err := aErr[account.ID]; err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
	}(account)

	return account, nil
}

/*
Get live account stats from wargaming
  - Each request will be retried c.retriesPerRequest times
  - This function will assume the player ID is valid and optimistically run all request concurrently, returning the first error once all request finish
*/
func (c *multiSourceClient) CurrentStats(ctx context.Context, id string, opts ...StatsOption) (AccountStatsOverPeriod, error) {
	var options statsOptions
	for _, apply := range opts {
		apply(&options)
	}

	realm, ok := c.wargaming.RealmFromID(id)
	if !ok {
		return AccountStatsOverPeriod{}, wargaming.ErrRealmNotSupported
	}

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

		var vehicleIDs []string // nil array will return all vehicles
		if options.VehicleIDs != nil && len(options.VehicleIDs) < 100 {
			// wg only supports 100 vehicles IDs per request
			vehicleIDs = options.VehicleIDs
		}

		vehicles = retry.Retry(func() ([]types.VehicleStatsFrame, error) {
			return c.wargaming.AccountVehicles(ctx, realm, id, vehicleIDs)
		}, c.retriesPerRequest, c.retrySleepInterval)

		if vehicles.Err != nil || len(vehicles.Data) < 1 || !options.WithWN8 {
			return
		}

		// manually filter vehicles for cases where the slice of ids was 100+
		if options.VehicleIDs != nil && len(options.VehicleIDs) >= 100 {
			var filtered []types.VehicleStatsFrame
			for _, v := range vehicles.Data {
				if !slices.Contains(options.VehicleIDs, fmt.Sprint(v.TankID)) {
					continue
				}
				filtered = append(filtered, v)
			}
			vehicles.Data = filtered
		}

		var ids []string
		for _, v := range vehicles.Data {
			tid := fmt.Sprint(v.TankID)
			ids = append(ids, tid)
		}
		a, err := c.database.GetVehicleAverages(ctx, ids)
		averages = retry.DataWithErr[map[string]frame.StatsFrame]{Data: a, Err: err}
	}()

	group.Wait()

	if account.Err != nil {
		return AccountStatsOverPeriod{}, parseWargamingError(account.Err)
	}
	if vehicles.Err != nil {
		return AccountStatsOverPeriod{}, parseWargamingError(vehicles.Err)
	}
	if averages.Err != nil {
		// not critical, this will only affect WN8
		log.Err(averages.Err).Msg("failed to get tank averages")
	}

	stats := WargamingToStats(realm, account.Data, clan, vehicles.Data)
	if options.WithWN8 {
		stats.AddWN8(averages.Data)
	}

	go func(account models.Account) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		aErr, err := c.database.UpsertAccounts(ctx, &account)
		if err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
		if err := aErr[account.ID]; err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
	}(stats.Account)

	return stats, nil
}

func (c *multiSourceClient) SessionStats(ctx context.Context, id string, sessionStart time.Time, opts ...StatsOption) (AccountStatsOverPeriod, AccountStatsOverPeriod, error) {
	var options = statsOptions{SnapshotType: models.SnapshotTypeDaily, ReferenceID: id}
	for _, apply := range opts {
		apply(&options)
	}

	if days := time.Since(sessionStart).Hours() / 24; sessionStart.IsZero() || days >= 91 {
		return AccountStatsOverPeriod{}, AccountStatsOverPeriod{}, ErrInvalidSessionStart
	}

	sessionBefore := sessionStart

	var accountSnapshot retry.DataWithErr[models.AccountSnapshot]
	var vehiclesSnapshots retry.DataWithErr[[]models.VehicleSnapshot]
	var averages retry.DataWithErr[map[string]frame.StatsFrame]
	var current retry.DataWithErr[AccountStatsOverPeriod]

	var group sync.WaitGroup
	group.Add(1)
	go func() {
		defer group.Done()

		stats, err := c.CurrentStats(ctx, id, opts...)
		current = retry.DataWithErr[AccountStatsOverPeriod]{Data: stats, Err: err}

		if err != nil || stats.RegularBattles.Battles < 1 || !options.WithWN8 {
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
		var opts = []database.Query{database.WithCreatedBefore(sessionBefore)}
		if options.ReferenceID != "" {
			opts = append(opts, database.WithReferenceIDIn(options.ReferenceID))
		}

		defer group.Done()
		s, err := c.database.GetAccountSnapshots(ctx, []string{id}, options.SnapshotType, opts...)
		if err != nil {
			accountSnapshot = retry.DataWithErr[models.AccountSnapshot]{Err: err}
			return
		}
		if len(s) < 1 {
			accountSnapshot = retry.DataWithErr[models.AccountSnapshot]{Err: database.ErrNotFound}
			return
		}
		accountSnapshot = retry.DataWithErr[models.AccountSnapshot]{Data: s[0]}

		v, err := c.database.GetVehicleSnapshots(ctx, id, options.VehicleIDs, options.SnapshotType, opts...)
		vehiclesSnapshots = retry.DataWithErr[[]models.VehicleSnapshot]{Data: v, Err: err}
	}()

	// wait for all requests to finish and check errors
	group.Wait()
	if current.Err != nil {
		return AccountStatsOverPeriod{}, AccountStatsOverPeriod{}, current.Err
	}
	if accountSnapshot.Err != nil {
		if database.IsNotFound(accountSnapshot.Err) {
			return AccountStatsOverPeriod{}, AccountStatsOverPeriod{}, ErrSessionNotFound
		}
		return AccountStatsOverPeriod{}, AccountStatsOverPeriod{}, accountSnapshot.Err
	}
	if averages.Err != nil {
		// not critical, this will only affect WN8
		log.Err(averages.Err).Msg("failed to get tank averages")
	}

	if options.VehicleIDs != nil {
		// if vehicle ids are set, there is not enough data to show rating stats
		current.Data.RatingBattles = StatsWithVehicles{}
		accountSnapshot.Data.RatingBattles = frame.StatsFrame{}
	}

	session := current.Data
	session.PeriodEnd = time.Now()
	session.PeriodStart = accountSnapshot.Data.LastBattleTime
	session.RatingBattles.StatsFrame.Subtract(accountSnapshot.Data.RatingBattles)
	session.RegularBattles.Vehicles = make(map[string]frame.VehicleStatsFrame, len(current.Data.RegularBattles.Vehicles))

	career := AccountStatsOverPeriod{
		Realm:          current.Data.Realm,
		Account:        current.Data.Account,
		LastBattleTime: accountSnapshot.Data.LastBattleTime,
		PeriodStart:    current.Data.Account.CreatedAt,
		PeriodEnd:      accountSnapshot.Data.LastBattleTime,
	}
	career.Account.LastBattleTime = accountSnapshot.Data.LastBattleTime
	career.RatingBattles.StatsFrame = accountSnapshot.Data.RatingBattles
	career.RatingBattles.Vehicles = make(map[string]frame.VehicleStatsFrame, 0)
	career.RegularBattles.Vehicles = make(map[string]frame.VehicleStatsFrame, len(vehiclesSnapshots.Data))

	// save career vehicle stats and make a snapshots map
	snapshotsMap := make(map[string]*models.VehicleSnapshot, len(vehiclesSnapshots.Data))
	for _, data := range vehiclesSnapshots.Data {
		career.RegularBattles.Vehicles[data.VehicleID] = frame.VehicleStatsFrame{
			LastBattleTime: data.LastBattleTime,
			VehicleID:      data.VehicleID,
			StatsFrame:     &data.Stats,
		}
		snapshotsMap[data.VehicleID] = &data
	}

	// compute per vehicle sessions
	for id, current := range current.Data.RegularBattles.Vehicles {
		snapshot, exists := snapshotsMap[id]
		if !exists {
			session.RegularBattles.Vehicles[id] = current
			continue
		}
		if current.Battles == 0 || current.Battles == snapshot.Stats.Battles {
			continue
		}

		sessionFrame := current
		stats := *current.StatsFrame
		stats.Subtract(snapshot.Stats)
		sessionFrame.StatsFrame = &stats
		session.RegularBattles.Vehicles[id] = sessionFrame
	}

	// compute total from selected/all vehicles
	if options.VehicleIDs != nil {
		session.RegularBattles.StatsFrame = VehiclesToFrame(session.RegularBattles.Vehicles)
	} else {
		session.RegularBattles.StatsFrame.Subtract(accountSnapshot.Data.RegularBattles)
	}

	if options.WithWN8 {
		career.AddWN8(averages.Data)
		session.AddWN8(averages.Data)
	}

	return session, career, nil

}

func (c *multiSourceClient) CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error) {
	return c.blitzstars.CurrentTankAverages(ctx)
}

func (c *multiSourceClient) ReplayRemote(ctx context.Context, fileURL string) (Replay, error) {
	unpacked, err := replay.UnpackRemote(ctx, fileURL)
	if err != nil {
		return Replay{}, errors.Wrap(err, "failed to unpack a remote replay")
	}
	return c.replay(ctx, unpacked)
}

func (c *multiSourceClient) Replay(ctx context.Context, file io.ReaderAt, size int64) (Replay, error) {
	unpacked, err := replay.Unpack(file, size)
	if err != nil {
		return Replay{}, errors.Wrap(err, "failed to unpack a remote replay")
	}
	return c.replay(ctx, unpacked)
}

func (c *multiSourceClient) replay(ctx context.Context, unpacked *replay.UnpackedReplay) (Replay, error) {
	replay := replay.Prettify(unpacked.BattleResult, unpacked.Meta)

	var vehicles []string
	var players = make(map[types.Realm][]string)
	for _, player := range append(replay.Teams.Allies, replay.Teams.Enemies...) {
		vehicles = append(vehicles, player.VehicleID)

		realm, ok := c.wargaming.RealmFromID(replay.Protagonist.ID)
		if !ok {
			continue
		}
		players[realm] = append(players[realm], player.ID)
	}

	var group errgroup.Group
	var playerDataMx sync.Mutex
	var playerData = make(map[string]types.ExtendedAccount)
	for realm, ids := range players {
		group.Go(func() error {
			data, err := c.wargaming.BatchAccountByID(ctx, realm, ids)
			if err != nil {
				return parseWargamingError(err)
			}

			playerDataMx.Lock()
			defer playerDataMx.Unlock()
			for id, player := range data {
				playerData[id] = player
			}
			return nil
		})
	}
	_ = group.Wait() // this error is not critical and will result in some missing data

	averages, err := c.database.GetVehicleAverages(ctx, vehicles)
	if err != nil {
		log.Err(err).Msg("failed to get tank averages for a replay")
	}

	// calculate and cache WN8
	_ = replay.Protagonist.Performance.WN8(averages[replay.Protagonist.VehicleID])
	for i, player := range append(replay.Teams.Allies, replay.Teams.Enemies...) {
		// calculate and cache WN8
		avg := averages[player.VehicleID]
		_ = player.Performance.WN8(avg)

		// set winrate cache
		if stats, ok := playerData[player.ID]; ok {
			frame := WargamingToFrame(stats.Statistics.All)
			player.Performance.SetWinRate(float32(frame.BattlesWon) / float32(frame.Battles) * 100)
		}

		if i < len(replay.Teams.Allies) {
			replay.Teams.Allies[i] = player
		} else {
			replay.Teams.Enemies[i-len(replay.Teams.Allies)] = player
		}
	}

	mapData, err := c.database.GetMap(ctx, replay.MapID)
	if err != nil && !database.IsNotFound(err) {
		return Replay{}, errors.Wrap(err, "failed to get map glossary")
	}

	return Replay{mapData, replay}, nil
}
