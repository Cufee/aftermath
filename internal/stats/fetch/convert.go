package fetch

import (
	"slices"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/external/blitzstars"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func timestampToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

func wargamingToAccount(realm string, account types.ExtendedAccount, clan types.ClanMember, private bool) database.Account {
	return database.Account{
		ID:       strconv.Itoa(account.ID),
		Realm:    realm,
		Nickname: account.Nickname,
		ClanTag:  clan.Clan.Tag,
		ClanID:   strconv.Itoa(clan.ClanID),

		Private:        private,
		CreatedAt:      timestampToTime(account.CreatedAt),
		LastBattleTime: timestampToTime(account.LastBattleTime),
	}
}

func wargamingToStats(realm string, accountData types.ExtendedAccount, clanMember types.ClanMember, vehicleData []types.VehicleStatsFrame) AccountStatsOverPeriod {
	stats := AccountStatsOverPeriod{
		Realm: realm,
		// we got the stats, so the account is obv not private at this point
		Account: wargamingToAccount(realm, accountData, clanMember, false),
		RegularBattles: StatsWithVehicles{
			StatsFrame: wargamingToFrame(accountData.Statistics.All),
			Vehicles:   wargamingVehiclesToFrame(vehicleData),
		},
		RatingBattles: StatsWithVehicles{
			StatsFrame: wargamingToFrame(accountData.Statistics.Rating),
			Vehicles:   make(map[string]frame.VehicleStatsFrame),
		},
		LastBattleTime: timestampToTime(accountData.LastBattleTime),
		PeriodStart:    time.Now(),
		PeriodEnd:      timestampToTime(accountData.LastBattleTime),
	}
	// An account can be blank with no last battle played
	if stats.PeriodEnd.Before(stats.PeriodStart) {
		stats.PeriodEnd = stats.PeriodStart
	}

	return stats
}

func wargamingToFrame(wg types.StatsFrame) frame.StatsFrame {
	return frame.StatsFrame{
		Battles:              frame.ValueInt(wg.Battles),
		BattlesWon:           frame.ValueInt(wg.Wins),
		BattlesSurvived:      frame.ValueInt(wg.SurvivedBattles),
		DamageDealt:          frame.ValueInt(wg.DamageDealt),
		DamageReceived:       frame.ValueInt(wg.DamageReceived),
		ShotsFired:           frame.ValueInt(wg.Shots),
		ShotsHit:             frame.ValueInt(wg.Hits),
		Frags:                frame.ValueInt(wg.Frags),
		MaxFrags:             frame.ValueInt(wg.MaxFrags),
		EnemiesSpotted:       frame.ValueInt(wg.Spotted),
		CapturePoints:        frame.ValueInt(wg.CapturePoints),
		DroppedCapturePoints: frame.ValueInt(wg.DroppedCapturePoints),
		Rating:               frame.ValueSpecialRating(wg.Rating),
	}
}

func wargamingVehiclesToFrame(wg []types.VehicleStatsFrame) map[string]frame.VehicleStatsFrame {
	stats := make(map[string]frame.VehicleStatsFrame)

	for _, record := range wg {
		id := strconv.Itoa(record.TankID)
		inner := wargamingToFrame(record.Stats)
		stats[id] = frame.VehicleStatsFrame{
			VehicleID:      id,
			StatsFrame:     &inner,
			MarkOfMastery:  record.MarkOfMastery,
			LastBattleTime: timestampToTime(record.LastBattleTime),
		}
	}

	return stats
}

func blitzstarsToStats(vehicles map[string]frame.VehicleStatsFrame, histories map[int][]blitzstars.TankHistoryEntry, from time.Time) StatsWithVehicles {
	stats := StatsWithVehicles{
		Vehicles: make(map[string]frame.VehicleStatsFrame),
	}

	for _, vehicle := range vehicles {
		if vehicle.LastBattleTime.Before(from) {
			continue
		}

		id, err := strconv.Atoi(vehicle.VehicleID)
		if err != nil || id == 0 {
			continue
		}

		entries := histories[id]
		// Sort entries by number of battles in descending order
		slices.SortFunc(entries, func(i, j blitzstars.TankHistoryEntry) int {
			return j.Stats.Battles - i.Stats.Battles
		})

		var selectedEntry blitzstars.TankHistoryEntry
		for _, entry := range entries {
			if entry.LastBattleTime < int(from.Unix()) {
				selectedEntry = entry
				break
			}
		}

		if selectedEntry.Stats.Battles < int(vehicle.Battles) {
			selectedFrame := wargamingToFrame(selectedEntry.Stats)
			vehicle.StatsFrame.Subtract(selectedFrame)

			stats.Vehicles[vehicle.VehicleID] = vehicle
			stats.Add(*vehicle.StatsFrame)
		}
	}

	return stats
}
