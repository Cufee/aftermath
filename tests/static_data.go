package tests

import (
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

const DefaultAccountNA = "account1"
const DefaultAccountEU = "account2"
const DefaultAccountAS = "account3"

var staticAccounts = map[string]models.Account{
	DefaultAccountNA: {ID: DefaultAccountNA, Realm: "NA", Nickname: "@test_account_na_1", CreatedAt: time.Now(), LastBattleTime: time.Now(), ClanID: "clan1", ClanTag: "TEST1"},
	DefaultAccountEU: {ID: DefaultAccountEU, Realm: "EU", Nickname: "@test_account_eu_1", CreatedAt: time.Now(), LastBattleTime: time.Now(), ClanID: "clan2", ClanTag: "TEST2"},
	DefaultAccountAS: {ID: DefaultAccountAS, Realm: "AS", Nickname: "@test_account_as_1", CreatedAt: time.Now(), LastBattleTime: time.Now(), ClanID: "clan3", ClanTag: "TEST3"},
}

var (
	DefaultStatsFrameBig1 = frame.StatsFrame{
		Battles:              100,
		BattlesWon:           80,
		BattlesSurvived:      70,
		DamageDealt:          234567,
		DamageReceived:       123567,
		ShotsHit:             690,
		ShotsFired:           1000,
		Frags:                420,
		MaxFrags:             50,
		EnemiesSpotted:       240,
		CapturePoints:        0,
		DroppedCapturePoints: 0,
		Rating:               100,
	}
	DefaultStatsFrameBig2 = frame.StatsFrame{
		Battles:              45,
		BattlesWon:           34,
		BattlesSurvived:      60,
		DamageDealt:          234567 / 2,
		DamageReceived:       123567 / 2,
		ShotsHit:             690 / 2,
		ShotsFired:           1000 / 2,
		Frags:                420 / 2,
		MaxFrags:             50 / 2,
		EnemiesSpotted:       240 / 2,
		CapturePoints:        0,
		DroppedCapturePoints: 0,
		Rating:               100,
	}
	DefaultStatsFrameSmall1 = frame.StatsFrame{
		Battles:              20,
		BattlesWon:           16,
		BattlesSurvived:      14,
		DamageDealt:          23456 * 2,
		DamageReceived:       12356 * 2,
		ShotsHit:             69 * 2,
		ShotsFired:           100 * 2,
		Frags:                42 * 2,
		MaxFrags:             5 * 2,
		EnemiesSpotted:       24 * 2,
		CapturePoints:        0,
		DroppedCapturePoints: 0,
		Rating:               100,
	}
	DefaultStatsFrameSmall2 = frame.StatsFrame{
		Battles:              10,
		BattlesWon:           8,
		BattlesSurvived:      7,
		DamageDealt:          23456,
		DamageReceived:       12356,
		ShotsHit:             69,
		ShotsFired:           100,
		Frags:                42,
		MaxFrags:             5,
		EnemiesSpotted:       24,
		CapturePoints:        0,
		DroppedCapturePoints: 0,
		Rating:               100,
	}
)

func DefaultVehicleStatsFrameBig1(id string) frame.VehicleStatsFrame {
	return frame.VehicleStatsFrame{
		VehicleID:  id,
		StatsFrame: &DefaultStatsFrameBig1,
	}
}

func DefaultVehicleStatsFrameBig2(id string) frame.VehicleStatsFrame {
	return frame.VehicleStatsFrame{
		VehicleID:  id,
		StatsFrame: &DefaultStatsFrameBig2,
	}
}

func DefaultVehicleStatsFrameSmall1(id string) frame.VehicleStatsFrame {
	return frame.VehicleStatsFrame{
		VehicleID:  id,
		StatsFrame: &DefaultStatsFrameSmall1,
	}
}

func DefaultVehicleStatsFrameSmall2(id string) frame.VehicleStatsFrame {
	return frame.VehicleStatsFrame{
		VehicleID:  id,
		StatsFrame: &DefaultStatsFrameSmall2,
	}
}
