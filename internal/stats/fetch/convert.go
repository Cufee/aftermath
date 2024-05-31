package fetch

import (
	"strconv"
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func timestampToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

func wargamingToStats(accountData types.ExtendedAccount, clanMember types.ClanMember, vehicleData []types.VehicleStatsFrame) AccountStatsOverPeriod {
	stats := AccountStatsOverPeriod{
		Account: accountData.Account,
		Clan:    clanMember.Clan,
		RegularBattles: StatsWithVehicles{
			StatsFrame: wargamingToFrame(accountData.Statistics.All),
			Vehicles:   wargamingVehiclesToFrame(vehicleData),
		},
		RatingBattles: StatsWithVehicles{
			StatsFrame: wargamingToFrame(accountData.Statistics.Rating),
			Vehicles:   make(map[string]VehicleStatsFrame),
		},
		LastBattleTime: timestampToTime(accountData.LastBattleTime),
		PeriodStart:    timestampToTime(accountData.CreatedAt),
		PeriodEnd:      timestampToTime(accountData.LastBattleTime),
	}
	// An account can be blank with no last battle played
	if stats.PeriodEnd.Before(stats.PeriodStart) {
		stats.PeriodEnd = stats.PeriodStart
	}

	return stats
}

func wargamingToFrame(wg types.StatsFrame) StatsFrame {
	return StatsFrame{
		Battles:              ValueInt(wg.Battles),
		BattlesWon:           ValueInt(wg.Wins),
		BattlesSurvived:      ValueInt(wg.SurvivedBattles),
		DamageDealt:          ValueInt(wg.DamageDealt),
		DamageReceived:       ValueInt(wg.DamageReceived),
		ShotsFired:           ValueInt(wg.Shots),
		ShotsHit:             ValueInt(wg.Hits),
		Frags:                ValueInt(wg.Frags),
		MaxFrags:             ValueInt(wg.MaxFrags),
		EnemiesSpotted:       ValueInt(wg.Spotted),
		CapturePoints:        ValueInt(wg.CapturePoints),
		DroppedCapturePoints: ValueInt(wg.DroppedCapturePoints),
		Rating:               ValueSpecialRating(wg.Rating),
	}
}

func wargamingVehiclesToFrame(wg []types.VehicleStatsFrame) map[string]VehicleStatsFrame {
	frame := make(map[string]VehicleStatsFrame)

	for _, stats := range wg {
		id := strconv.Itoa(stats.TankID)
		inner := wargamingToFrame(stats.Stats)
		frame[id] = VehicleStatsFrame{
			VehicleID:      id,
			StatsFrame:     &inner,
			MarkOfMastery:  stats.MarkOfMastery,
			LastBattleTime: timestampToTime(stats.LastBattleTime),
		}
	}

	return frame
}
