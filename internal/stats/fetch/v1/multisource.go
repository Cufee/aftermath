package fetch

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/ent/db"
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

func (c *multiSourceClient) Search(ctx context.Context, nickname, realm string, limit int) (*types.Account, error) {
	accounts, err := c.wargaming.SearchAccounts(ctx, realm, nickname, limit)
	if err != nil {
		return nil, err
	}
	if len(accounts) < 1 {
		return nil, AccountNotFound
	}

	return &accounts[0], nil
}

func (c *multiSourceClient) BroadSearch(ctx context.Context, nickname string, limitPerRealm int) ([]*AccountWithRealm, error) {
	data := make(chan AccountWithRealm, 9)
	errors := make(chan error, 3)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		accounts, err := c.wargaming.SearchAccounts(ctx, "NA", nickname, limitPerRealm)
		if err != nil {
			errors <- err
			return
		}
		for _, a := range accounts {
			data <- AccountWithRealm{a, "NA"}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		accounts, err := c.wargaming.SearchAccounts(ctx, "EU", nickname, limitPerRealm)
		if err != nil {
			errors <- err
			return
		}
		for _, a := range accounts {
			data <- AccountWithRealm{a, "EU"}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		accounts, err := c.wargaming.SearchAccounts(ctx, "AS", nickname, limitPerRealm)
		if err != nil {
			errors <- err
			return
		}
		for _, a := range accounts {
			data <- AccountWithRealm{a, "AS"}
		}
	}()
	wg.Wait()
	close(data)
	close(errors)

	// return first error
	if len(errors) == 3 || (len(data) == 0 && len(errors) > 0) {
		for err := range errors {
			return nil, err
		}
	}

	var accounts []*AccountWithRealm
	for a := range data {
		accounts = append(accounts, &a)
	}

	return accounts, nil
}

/*
Gets account info from wg and updates cache
*/
func (c *multiSourceClient) Account(ctx context.Context, id string) (*models.Account, error) {
	realm := c.wargaming.RealmFromAccountID(id)

	var group sync.WaitGroup
	group.Add(2)

	var clan types.ClanMember
	var wgAccount retry.DataWithErr[types.ExtendedAccount]

	go func(account *retry.DataWithErr[types.ExtendedAccount]) {
		defer group.Done()

		*account = retry.Retry(func() (types.ExtendedAccount, error) {
			return c.wargaming.AccountByID(ctx, realm, id)
		}, c.retriesPerRequest, c.retrySleepInterval)
	}(&wgAccount)

	go func(clan *types.ClanMember) {
		defer group.Done()
		// we should not retry this request since it is not critical and will fail on accounts without a clan
		*clan, _ = c.wargaming.AccountClan(ctx, realm, id)
	}(&clan)

	group.Wait()

	if wgAccount.Err != nil {
		return nil, wgAccount.Err
	}

	account := WargamingToAccount(realm, &wgAccount.Data, &clan, false)
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		aErr, err := c.database.UpsertAccounts(ctx, account)
		if err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
		if err := aErr[account.ID]; err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
	}

	return account, nil
}

/*
Get live account stats from wargaming
  - Each request will be retried c.retriesPerRequest times
  - This function will assume the player ID is valid and optimistically run all request concurrently, returning the first error once all request finish
*/
func (c *multiSourceClient) CurrentStats(ctx context.Context, id string, opts ...StatsOption) (*AccountStatsOverPeriod, error) {
	var options statsOptions
	for _, apply := range opts {
		apply(&options)
	}

	realm := c.wargaming.RealmFromAccountID(id)

	var group sync.WaitGroup
	group.Add(3)

	var clan types.ClanMember
	var account retry.DataWithErr[types.ExtendedAccount]
	var vehicles = retry.DataWithErr[[]types.VehicleStatsFrame]{Data: make([]types.VehicleStatsFrame, 0, 300)}
	var averages retry.DataWithErr[map[string]frame.StatsFrame]

	go func(ctx context.Context, realm, id string, account *retry.DataWithErr[types.ExtendedAccount]) {
		defer group.Done()

		*account = retry.Retry(func() (types.ExtendedAccount, error) {
			return c.wargaming.AccountByID(ctx, realm, id)
		}, c.retriesPerRequest, c.retrySleepInterval)
	}(ctx, realm, id, &account)

	go func(ctx context.Context, realm, id string, clan *types.ClanMember) {
		defer group.Done()
		// we should not retry this request since it is not critical and will fail on accounts without a clan
		*clan, _ = c.wargaming.AccountClan(ctx, realm, id)
	}(ctx, realm, id, &clan)

	go func(ctx context.Context, realm, id string, vehicles *retry.DataWithErr[[]types.VehicleStatsFrame], averages *retry.DataWithErr[map[string]frame.StatsFrame]) {
		defer group.Done()

		var vehicleIDs []string // nil array will return all vehicles
		if options.vehicleID != "" {
			vehicleIDs = append(vehicleIDs, options.vehicleID)
		}

		*vehicles = retry.Retry(func() ([]types.VehicleStatsFrame, error) {
			return c.wargaming.AccountVehicles(ctx, realm, id, vehicleIDs)
		}, c.retriesPerRequest, c.retrySleepInterval)

		if vehicles.Err != nil || len(vehicles.Data) < 1 || !options.withWN8 {
			return
		}

		var ids []string
		for _, v := range vehicles.Data {
			tid := fmt.Sprint(v.TankID)
			ids = append(ids, tid)
		}
		a, err := c.database.GetVehicleAverages(ctx, ids)
		*averages = retry.DataWithErr[map[string]frame.StatsFrame]{Data: a, Err: err}
	}(ctx, realm, id, &vehicles, &averages)

	group.Wait()

	if account.Err != nil {
		return nil, account.Err
	}
	if vehicles.Err != nil {
		return nil, vehicles.Err
	}
	if averages.Err != nil {
		// not critical, this will only affect WN8
		log.Err(averages.Err).Msg("failed to get tank averages")
	}

	stats := WargamingToStats(realm, &account.Data, &clan, vehicles.Data)
	if options.withWN8 {
		stats.AddWN8(averages.Data)
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		aErr, err := c.database.UpsertAccounts(ctx, stats.Account)
		if err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
		if err := aErr[stats.Account.ID]; err != nil {
			log.Err(err).Msg("failed to update account cache")
		}
	}

	return stats, nil
}

func (c *multiSourceClient) PeriodStats(ctx context.Context, id string, periodStart time.Time, opts ...StatsOption) (*AccountStatsOverPeriod, error) {
	var options statsOptions
	for _, apply := range opts {
		apply(&options)
	}

	var histories retry.DataWithErr[map[int][]blitzstars.TankHistoryEntry]
	var averages retry.DataWithErr[map[string]frame.StatsFrame]
	var current retry.DataWithErr[*AccountStatsOverPeriod]

	var group sync.WaitGroup
	group.Add(1)
	go func(current *retry.DataWithErr[*AccountStatsOverPeriod], averages *retry.DataWithErr[map[string]frame.StatsFrame]) {
		defer group.Done()

		stats, err := c.CurrentStats(ctx, id, opts...)
		*current = retry.DataWithErr[*AccountStatsOverPeriod]{Data: stats, Err: err}

		if err != nil || stats.RegularBattles.Battles < 1 || !options.withWN8 {
			return
		}

		var ids []string
		for id := range stats.RegularBattles.Vehicles {
			ids = append(ids, id)
		}
		a, err := c.database.GetVehicleAverages(ctx, ids)
		*averages = retry.DataWithErr[map[string]frame.StatsFrame]{Data: a, Err: err}
	}(&current, &averages)

	// TODO: lookup a session from the database first
	// if a session exists in the database, we don't need BlitzStars and have better data

	// return career stats if stats are requested for 0 or 90+ days, we do not track that far
	if days := time.Since(periodStart).Hours() / 24; days >= 91 || days < 1 {
		group.Wait()
		if current.Err != nil {
			return nil, current.Err
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
	go func(histories *retry.DataWithErr[map[int][]blitzstars.TankHistoryEntry]) {
		defer group.Done()
		*histories = retry.Retry(func() (map[int][]blitzstars.TankHistoryEntry, error) {
			return c.blitzstars.AccountTankHistories(ctx, id)
		}, c.retriesPerRequest, c.retrySleepInterval)
	}(&histories)

	// wait for all requests to finish and check errors
	group.Wait()
	if current.Err != nil {
		return nil, current.Err
	}
	if histories.Err != nil {
		return nil, histories.Err
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

func (c *multiSourceClient) SessionStats(ctx context.Context, id string, sessionStart time.Time, opts ...StatsOption) (*AccountStatsOverPeriod, *AccountStatsOverPeriod, error) {
	var options = statsOptions{snapshotType: models.SnapshotTypeDaily, referenceID: id}
	for _, apply := range opts {
		apply(&options)
	}

	if days := time.Since(sessionStart).Hours() / 24; sessionStart.IsZero() || days >= 91 {
		return nil, nil, ErrInvalidSessionStart
	}

	sessionBefore := sessionStart

	var accountSnapshot retry.DataWithErr[models.AccountSnapshot]
	var vehiclesSnapshots retry.DataWithErr[[]models.VehicleSnapshot]
	var averages retry.DataWithErr[map[string]frame.StatsFrame]
	var current retry.DataWithErr[*AccountStatsOverPeriod]

	var group sync.WaitGroup
	group.Add(1)
	go func(current *retry.DataWithErr[*AccountStatsOverPeriod], averages *retry.DataWithErr[map[string]frame.StatsFrame]) {
		defer group.Done()

		stats, err := c.CurrentStats(ctx, id, opts...)
		*current = retry.DataWithErr[*AccountStatsOverPeriod]{Data: stats, Err: err}

		if err != nil || stats.RegularBattles.Battles < 1 || !options.withWN8 {
			return
		}

		var ids []string
		for id := range stats.RegularBattles.Vehicles {
			ids = append(ids, id)
		}
		a, err := c.database.GetVehicleAverages(ctx, ids)
		*averages = retry.DataWithErr[map[string]frame.StatsFrame]{Data: a, Err: err}
	}(&current, &averages)

	group.Add(1)
	go func(account *retry.DataWithErr[models.AccountSnapshot], vehicles *retry.DataWithErr[[]models.VehicleSnapshot]) {
		var opts = []database.Query{database.WithCreatedBefore(sessionBefore)}
		if options.referenceID != "" {
			opts = append(opts, database.WithReferenceIDIn(options.referenceID))
		}

		defer group.Done()
		s, err := c.database.GetAccountSnapshots(ctx, []string{id}, options.snapshotType, opts...)
		if err != nil {
			*account = retry.DataWithErr[models.AccountSnapshot]{Err: err}
			return
		}
		if len(s) < 1 {
			*account = retry.DataWithErr[models.AccountSnapshot]{Err: new(db.NotFoundError)}
			return
		}
		*account = retry.DataWithErr[models.AccountSnapshot]{Data: s[0]}

		var vIDs []string
		if options.vehicleID != "" {
			vIDs = append(vIDs, options.vehicleID)
		}

		v, err := c.database.GetVehicleSnapshots(ctx, id, vIDs, options.snapshotType, opts...)
		*vehicles = retry.DataWithErr[[]models.VehicleSnapshot]{Data: v, Err: err}
	}(&accountSnapshot, &vehiclesSnapshots)

	// wait for all requests to finish and check errors
	group.Wait()
	if current.Err != nil {
		return nil, nil, current.Err
	}
	if accountSnapshot.Err != nil {
		if database.IsNotFound(accountSnapshot.Err) {
			return nil, nil, ErrSessionNotFound
		}
		return nil, nil, accountSnapshot.Err
	}
	if averages.Err != nil {
		// not critical, this will only affect WN8
		log.Err(averages.Err).Msg("failed to get tank averages")
	}

	session := current.Data
	session.PeriodEnd = time.Now()
	session.PeriodStart = accountSnapshot.Data.LastBattleTime
	session.RatingBattles.StatsFrame.Subtract(&accountSnapshot.Data.RatingBattles)
	session.RegularBattles.StatsFrame.Subtract(&accountSnapshot.Data.RegularBattles)
	session.RegularBattles.Vehicles = make(map[string]frame.VehicleStatsFrame, len(current.Data.RegularBattles.Vehicles))

	career := &AccountStatsOverPeriod{
		Realm:          current.Data.Realm,
		Account:        current.Data.Account,
		LastBattleTime: accountSnapshot.Data.LastBattleTime,
		PeriodStart:    current.Data.Account.CreatedAt,
		PeriodEnd:      accountSnapshot.Data.LastBattleTime,
	}
	career.Account.LastBattleTime = accountSnapshot.Data.LastBattleTime
	career.RatingBattles.StatsFrame = accountSnapshot.Data.RatingBattles
	career.RegularBattles.StatsFrame = accountSnapshot.Data.RegularBattles
	career.RatingBattles.Vehicles = make(map[string]frame.VehicleStatsFrame, 0)
	career.RegularBattles.Vehicles = make(map[string]frame.VehicleStatsFrame, len(vehiclesSnapshots.Data))

	snapshotsMap := make(map[string]*models.VehicleSnapshot, len(vehiclesSnapshots.Data))
	for _, data := range vehiclesSnapshots.Data {
		snapshotsMap[data.VehicleID] = &data

		career.RegularBattles.Vehicles[data.VehicleID] = frame.VehicleStatsFrame{
			LastBattleTime: data.LastBattleTime,
			VehicleID:      data.VehicleID,
			StatsFrame:     &data.Stats,
		}
	}

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
		stats.Subtract(&snapshot.Stats)
		sessionFrame.StatsFrame = &stats
		session.RegularBattles.Vehicles[id] = sessionFrame
	}

	if options.withWN8 {
		career.AddWN8(averages.Data)
		session.AddWN8(averages.Data)
	}

	return session, career, nil

}

func (c *multiSourceClient) CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error) {
	return c.blitzstars.CurrentTankAverages(ctx)
}

func (c *multiSourceClient) ReplayRemote(ctx context.Context, fileURL string) (*Replay, error) {
	unpacked, err := replay.UnpackRemote(ctx, fileURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack a remote replay")
	}
	return c.replay(ctx, unpacked)
}

func (c *multiSourceClient) Replay(ctx context.Context, file io.ReaderAt, size int64) (*Replay, error) {
	unpacked, err := replay.Unpack(file, size)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unpack a remote replay")
	}
	return c.replay(ctx, unpacked)
}

func (c *multiSourceClient) replay(ctx context.Context, unpacked *replay.UnpackedReplay) (*Replay, error) {
	replay := replay.Prettify(unpacked.BattleResult, unpacked.Meta)

	var vehicles []string
	for _, player := range append(replay.Teams.Allies, replay.Teams.Enemies...) {
		vehicles = append(vehicles, player.VehicleID)
	}

	averages, err := c.database.GetVehicleAverages(ctx, vehicles)
	if err != nil {
		log.Err(err).Msg("failed to get tank averages for a replay")
	}

	// calculate and cache WN8
	_ = replay.Protagonist.Performance.WN8(averages[replay.Protagonist.VehicleID])
	for i, player := range append(replay.Teams.Allies, replay.Teams.Enemies...) {
		avg := averages[player.VehicleID]
		_ = player.Performance.WN8(avg)
		if i < len(replay.Teams.Allies) {
			replay.Teams.Allies[i] = player
		} else {
			replay.Teams.Enemies[i-len(replay.Teams.Allies)] = player
		}
	}

	mapData, err := c.database.GetMap(ctx, replay.MapID)
	if err != nil && !database.IsNotFound(err) {
		return nil, errors.Wrap(err, "failed to get map glossary")
	}

	return &Replay{mapData, replay, c.wargaming.RealmFromAccountID(replay.Protagonist.ID)}, nil
}
