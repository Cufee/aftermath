package fetch

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type AccountStatsOverPeriod struct {
	Account types.Account `json:"account"`
	Clan    types.Clan    `json:"clan"`

	PeriodStart time.Time `json:"start"`
	PeriodEnd   time.Time `json:"end"`

	RegularBattles StatsWithVehicles `json:"regular"`
	RatingBattles  StatsWithVehicles `json:"rating"`

	LastBattleTime time.Time `json:"lastBattleTime"`
}

type StatsWithVehicles struct {
	frame.StatsFrame
	Vehicles map[string]frame.VehicleStatsFrame
}

type Client interface {
	// Search(id string) (AccountStatsOverPeriod, error)
	CurrentStats(ctx context.Context, id string) (AccountStatsOverPeriod, error)
	PeriodStats(ctx context.Context, id string, from time.Time) (AccountStatsOverPeriod, error)
}
