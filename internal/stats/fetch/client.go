package fetch

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type AccountStatsOverPeriod struct {
	Realm string `json:"realm"`

	Account database.Account `json:"account"`

	PeriodStart time.Time `json:"start"`
	PeriodEnd   time.Time `json:"end"`

	RegularBattles StatsWithVehicles `json:"regular"`
	RatingBattles  StatsWithVehicles `json:"rating"`

	LastBattleTime time.Time `json:"lastBattleTime"`
}

func (stats *AccountStatsOverPeriod) AddWN8(averages map[string]frame.StatsFrame) {
	var weightedTotal, battlesTotal float32
	for id, data := range stats.RegularBattles.Vehicles {
		tankAverages, ok := averages[id]
		if !ok || data.Battles < 1 {
			continue
		}
		weightedTotal += data.Battles.Float() * data.WN8(tankAverages).Float()
		battlesTotal += data.Battles.Float()
	}
	if battlesTotal < 1 {
		return
	}

	wn8 := int(weightedTotal) / int(battlesTotal)
	stats.RegularBattles.SetWN8(wn8)
}

type StatsWithVehicles struct {
	frame.StatsFrame
	Vehicles map[string]frame.VehicleStatsFrame
}

type Client interface {
	Account(ctx context.Context, id string) (database.Account, error)
	Search(ctx context.Context, nickname, realm string) (types.Account, error)
	CurrentStats(ctx context.Context, id string, opts ...statsOption) (AccountStatsOverPeriod, error)

	PeriodStats(ctx context.Context, id string, from time.Time, opts ...statsOption) (AccountStatsOverPeriod, error)
	SessionStats(ctx context.Context, id string, sessionStart time.Time, opts ...statsOption) (AccountStatsOverPeriod, AccountStatsOverPeriod, error)

	CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error)
}

type statsOptions struct {
	withWN8 bool
}

type statsOption func(*statsOptions)

func WithWN8() statsOption {
	return func(so *statsOptions) {
		so.withWN8 = true
	}
}
