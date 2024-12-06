package tests

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/pkg/errors"
)

var _ fetch.Client = &staticTestingFetch{}

type staticTestingFetch struct{}

func StaticTestingFetch() *staticTestingFetch {
	return &staticTestingFetch{}
}

func (c *staticTestingFetch) VerifyAccountToken(ctx context.Context, id string, token string) (bool, error) {
	return true, nil
}

func (c *staticTestingFetch) Account(ctx context.Context, id string) (models.Account, error) {
	if account, ok := staticAccounts[id]; ok {
		return account, nil
	}
	return models.Account{}, errors.New("account not found")
}
func (c *staticTestingFetch) Search(ctx context.Context, nickname string, realm types.Realm, limit int) (types.Account, error) {
	return types.Account{}, nil
}
func (c *staticTestingFetch) BroadSearch(ctx context.Context, nickname string, limit int) ([]fetch.AccountWithRealm, error) {
	return nil, nil
}
func (c *staticTestingFetch) CurrentStats(ctx context.Context, id string, opts ...fetch.StatsOption) (fetch.AccountStatsOverPeriod, error) {
	account, err := c.Account(ctx, id)
	if err != nil {
		return fetch.AccountStatsOverPeriod{}, err
	}

	var vehicles = make(map[string]frame.VehicleStatsFrame)
	for id := range 10 {
		f := DefaultVehicleStatsFrameBig1(fmt.Sprint(id))
		f.SetWN8(9999)
		vehicles[fmt.Sprint(id)] = f
	}

	stats := fetch.AccountStatsOverPeriod{
		Account: account,
		Realm:   account.Realm,

		PeriodEnd:      time.Now(),
		PeriodStart:    time.Now().Add(time.Hour * 25 * 1),
		LastBattleTime: time.Now(),

		RegularBattles: fetch.StatsWithVehicles{
			Vehicles:   vehicles,
			StatsFrame: DefaultStatsFrameBig1,
		},
		RatingBattles: fetch.StatsWithVehicles{
			StatsFrame: DefaultStatsFrameBig2,
		},
	}
	stats.RegularBattles.SetWN8(9999)
	return stats, nil
}

func (c *staticTestingFetch) PeriodStats(ctx context.Context, id string, from time.Time, opts ...fetch.StatsOption) (fetch.AccountStatsOverPeriod, error) {
	current, err := c.CurrentStats(ctx, id, opts...)
	if err != nil {
		return fetch.AccountStatsOverPeriod{}, err
	}

	current.PeriodStart = from
	current.RegularBattles.SetWN8(9999)
	current.RegularBattles.StatsFrame.Subtract(DefaultStatsFrameSmall1)
	current.RatingBattles.StatsFrame.Subtract(DefaultStatsFrameSmall2)

	for id, stats := range current.RegularBattles.Vehicles {
		stats.SetWN8(9999)
		stats.StatsFrame.Subtract(DefaultStatsFrameSmall1)
		current.RegularBattles.Vehicles[id] = stats
	}
	return current, nil
}
func (c *staticTestingFetch) SessionStats(ctx context.Context, id string, sessionStart time.Time, opts ...fetch.StatsOption) (fetch.AccountStatsOverPeriod, fetch.AccountStatsOverPeriod, error) {
	session, err := c.PeriodStats(ctx, id, sessionStart, opts...)
	if err != nil {
		return fetch.AccountStatsOverPeriod{}, fetch.AccountStatsOverPeriod{}, err
	}
	career, err := c.CurrentStats(ctx, id, opts...)
	if err != nil {
		return fetch.AccountStatsOverPeriod{}, fetch.AccountStatsOverPeriod{}, err
	}

	session.RegularBattles.SetWN8(3495)
	career.RegularBattles.SetWN8(3495)

	for id, stats := range session.RegularBattles.Vehicles {
		stats.SetWN8(3495)
		session.RegularBattles.Vehicles[id] = stats
	}
	for id, stats := range career.RegularBattles.Vehicles {
		stats.SetWN8(3295)
		career.RegularBattles.Vehicles[id] = stats
	}

	return session, career, nil
}

func (c *staticTestingFetch) CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error) {
	// TODO: add some data
	return map[string]frame.StatsFrame{}, nil
}

func (c *staticTestingFetch) ReplayRemote(ctx context.Context, url string) (fetch.Replay, error) {
	// TODO: add some data
	return fetch.Replay{}, nil
}

func (c *staticTestingFetch) Replay(ctx context.Context, file io.ReaderAt, size int64) (fetch.Replay, error) {
	// TODO: add some data
	return fetch.Replay{}, nil
}
