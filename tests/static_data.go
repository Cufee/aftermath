package tests

import (
	"fmt"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

const DefaultAccountNA = "10003333004"
const DefaultAccountNAShort = "2000444005"
const DefaultAccountEU = "20000555006"
const DefaultAccountAS = "30000666007"

var staticAccounts = map[string]models.Account{
	DefaultAccountNA:      {ID: DefaultAccountNA, Realm: "NA", Nickname: "Stylish_Player_One", CreatedAt: time.Now(), LastBattleTime: time.Now(), ClanID: "clan1", ClanTag: "AMTH"},
	DefaultAccountNAShort: {ID: DefaultAccountNA, Realm: "NA", Nickname: "test_acc", CreatedAt: time.Now(), LastBattleTime: time.Now(), ClanID: "clan1", ClanTag: "AMTH"},
	DefaultAccountEU:      {ID: DefaultAccountEU, Realm: "EU", Nickname: "test_account_eu_1", CreatedAt: time.Now(), LastBattleTime: time.Now(), ClanID: "clan2", ClanTag: "AMTH"},
	DefaultAccountAS:      {ID: DefaultAccountAS, Realm: "AS", Nickname: "test_account_as_1", CreatedAt: time.Now(), LastBattleTime: time.Now(), ClanID: "clan3", ClanTag: "AMTH"},
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
		Rating:               228,
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
		Rating:               228,
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
		Rating:               228,
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
		Rating:               228,
	}
)

func DefaultVehicleStatsFrameBig1(id string) frame.VehicleStatsFrame {
	f := DefaultStatsFrameBig1
	return frame.VehicleStatsFrame{
		VehicleID:  id,
		StatsFrame: &f,
	}
}

func DefaultVehicleStatsFrameBig2(id string) frame.VehicleStatsFrame {
	f := DefaultStatsFrameBig2
	return frame.VehicleStatsFrame{
		VehicleID:  id,
		StatsFrame: &f,
	}
}

func DefaultVehicleStatsFrameSmall1(id string) frame.VehicleStatsFrame {
	f := DefaultStatsFrameSmall1
	return frame.VehicleStatsFrame{
		VehicleID:  id,
		StatsFrame: &f,
	}
}

func DefaultVehicleStatsFrameSmall2(id string) frame.VehicleStatsFrame {
	f := DefaultStatsFrameSmall2
	return frame.VehicleStatsFrame{
		VehicleID:  id,
		StatsFrame: &f,
	}
}

var (
	DefaultUserWithEdges = models.User{ID: "user1", Connections: []models.UserConnection{
		Connection(DefaultAccountNA, models.ConnectionTypeWargaming, false, true),
		Connection(DefaultAccountEU, models.ConnectionTypeWargaming, true, false),
		Connection(DefaultAccountAS, models.ConnectionTypeWargaming, false, false),
	}}
)

func Connection(id string, kind models.ConnectionType, verified, selected bool) models.UserConnection {
	return models.UserConnection{
		ID:          fmt.Sprint(time.Now().Nanosecond()),
		ReferenceID: id,
		Type:        kind,
		Selected:    selected,
		Verified:    verified,
	}
}
