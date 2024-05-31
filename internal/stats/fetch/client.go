package fetch

import (
	"time"

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
	StatsFrame
	Vehicles map[string]VehicleStatsFrame
}

type Client interface {
	// Search(id string) (AccountStatsOverPeriod, error)
	CurrentStats(id string) (AccountStatsOverPeriod, error)
	PeriodStats(id string, from time.Time) (AccountStatsOverPeriod, error)
}
