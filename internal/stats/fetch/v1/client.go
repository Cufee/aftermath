package fetch

import (
	"context"
	"io"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type AccountWithRealm struct {
	types.Account
	Realm types.Realm
}

type AccountStatsOverPeriod struct {
	Realm types.Realm `json:"realm"`

	Account models.Account `json:"account"`

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
	Account(ctx context.Context, id string) (models.Account, error)
	Search(ctx context.Context, nickname string, realm types.Realm, limit int) (types.Account, error)
	BroadSearch(ctx context.Context, nickname string, limit int) ([]AccountWithRealm, error)
	CurrentStats(ctx context.Context, id string, opts ...StatsOption) (AccountStatsOverPeriod, error)

	SessionStats(ctx context.Context, id string, sessionStart time.Time, opts ...StatsOption) (AccountStatsOverPeriod, AccountStatsOverPeriod, error)

	ReplayRemote(ctx context.Context, fileURL string) (Replay, error)
	Replay(ctx context.Context, file io.ReaderAt, size int64) (Replay, error)
	CurrentTankAverages(ctx context.Context) (map[string]frame.StatsFrame, error)
}

type statsOptions struct {
	WithWN8      bool
	VehicleIDs   []string
	ReferenceID  string
	SnapshotType models.SnapshotType
}

type StatsOption func(*statsOptions)

func ParseOptions(opts []StatsOption) statsOptions {
	var o statsOptions
	for _, apply := range opts {
		apply(&o)
	}
	return o
}

func WithWN8() StatsOption {
	return func(so *statsOptions) {
		so.WithWN8 = true
	}
}
func WithType(sType models.SnapshotType) StatsOption {
	return func(so *statsOptions) {
		so.SnapshotType = sType
	}
}
func WithReferenceID(reference string) StatsOption {
	return func(so *statsOptions) {
		so.ReferenceID = reference
	}
}
func WithVehicleIDs(ids ...string) StatsOption {
	return func(so *statsOptions) {
		so.VehicleIDs = ids
	}
}
