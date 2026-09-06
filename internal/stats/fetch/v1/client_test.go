package fetch

import (
	"testing"

	"github.com/cufee/aftermath/internal/stats/frame"
)

func TestAddWN8OnlyWeightsVehiclesWithAvailableWN8(t *testing.T) {
	rated := frame.StatsFrame{Battles: 10}
	rated.SetWN8(1000)

	zeroRated := frame.StatsFrame{Battles: 10}
	unrated := frame.StatsFrame{Battles: 80}

	stats := AccountStatsOverPeriod{
		RegularBattles: StatsWithVehicles{
			Vehicles: map[string]frame.VehicleStatsFrame{
				"rated":     {StatsFrame: &rated},
				"zeroRated": {StatsFrame: &zeroRated},
				"unrated":   {StatsFrame: &unrated},
			},
		},
	}
	stats.AddWN8(map[string]frame.StatsFrame{
		"rated": {},
		"zeroRated": {
			Battles:              10,
			BattlesWon:           5,
			DamageDealt:          10_000,
			Frags:                10,
			EnemiesSpotted:       10,
			DroppedCapturePoints: 10,
		},
		"unrated": {},
	})

	if got := stats.RegularBattles.WN8().Float(); got != 500 {
		t.Fatalf("expected WN8 500 from rated vehicles, got %v", got)
	}
}
