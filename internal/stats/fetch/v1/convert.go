package fetch

import (
	"strconv"
	"time"

	assets "github.com/cufee/aftermath-assets/types"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/fetch/v1/replay"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

func timestampToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

func WargamingToAccount(realm types.Realm, account types.ExtendedAccount, clan types.ClanMember, private bool) models.Account {
	a := models.Account{
		ID:       strconv.Itoa(account.ID),
		Realm:    realm,
		Nickname: account.Nickname,

		Private:        private,
		CreatedAt:      timestampToTime(account.CreatedAt),
		LastBattleTime: timestampToTime(account.LastBattleTime),
	}
	if clan.ClanID > 0 {
		a.ClanTag = clan.Clan.Tag
		a.ClanID = strconv.Itoa(clan.ClanID)
	}
	return a
}

func WargamingToStats(realm types.Realm, accountData types.ExtendedAccount, clanMember types.ClanMember, vehicleData []types.VehicleStatsFrame) AccountStatsOverPeriod {
	unratedVehicles := WargamingVehiclesToFrame(vehicleData)

	stats := AccountStatsOverPeriod{
		Realm: realm,
		// we got the stats, so the account is obv not private at this point
		Account: WargamingToAccount(realm, accountData, clanMember, false),
		RegularBattles: StatsWithVehicles{
			StatsFrame: vehiclesToOverview(unratedVehicles),
			Vehicles:   unratedVehicles,
		},
		RatingBattles: StatsWithVehicles{
			StatsFrame: WargamingToFrame(accountData.Statistics.Rating),
			Vehicles:   make(map[string]frame.VehicleStatsFrame),
		},
		LastBattleTime: timestampToTime(accountData.LastBattleTime),
		PeriodEnd:      timestampToTime(accountData.LastBattleTime),
		PeriodStart:    timestampToTime(accountData.CreatedAt),
	}
	// An account can be blank with no last battle played
	if stats.LastBattleTime.IsZero() {
		stats.PeriodEnd = stats.PeriodStart
	}

	return stats
}

func vehiclesToOverview(vehicles map[string]frame.VehicleStatsFrame) frame.StatsFrame {
	var frame frame.StatsFrame
	for _, v := range vehicles {
		frame.Add(*v.StatsFrame)
	}
	return frame
}

func WargamingToFrame(wg types.StatsFrame) frame.StatsFrame {
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
		RawRating:            frame.ValueSpecialRating(wg.Rating),
	}
}

func WargamingVehiclesToFrame(wg []types.VehicleStatsFrame) map[string]frame.VehicleStatsFrame {
	stats := make(map[string]frame.VehicleStatsFrame)

	for _, record := range wg {
		id := strconv.Itoa(record.TankID)
		inner := WargamingToFrame(record.Stats)
		stats[id] = frame.VehicleStatsFrame{
			VehicleID:      id,
			StatsFrame:     &inner,
			MarkOfMastery:  record.MarkOfMastery,
			LastBattleTime: timestampToTime(record.LastBattleTime),
		}
	}

	return stats
}

type Replay struct {
	Map assets.Map
	replay.Replay
}
