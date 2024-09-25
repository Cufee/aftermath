package replay

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cufee/aftermath/internal/stats/frame"
)

type Outcome string

const (
	OutcomeVictory Outcome = "victory"
	OutcomeDefeat  Outcome = "defeat"
	OutcomeDraw    Outcome = "draw"
)

type battleType struct {
	ID  int    `json:"id" protobuf:"10"`
	Tag string `json:"tag" protobuf:"11"`
}

func (bt battleType) String() string {
	return bt.Tag
}

var (
	BattleTypeUnknown   = battleType{-1, "battle_type_unknown"}
	BattleTypeRandom    = battleType{0, "battle_type_regular"}
	BattleTypeSupremacy = battleType{1, "battle_type_supremacy"}
)

var battleTypes = map[int]battleType{
	0: BattleTypeRandom,
	1: BattleTypeSupremacy,
}

type gameMode struct {
	ID      int      `json:"id" protobuf:"15"`
	Name    string   `json:"name" protobuf:"16"`
	Special bool     `json:"special" protobuf:"17"` // Signifies if WN8 should be calculated
	Tags    []string `json:"tags"`
}

func (gm gameMode) String() string {
	return gm.Name
}

var gameModes = map[int]gameMode{
	0:  {-1, "game_mode_unknown", false, nil},
	2:  {2, "game_mode_training", false, nil},
	6:  {6, "game_mode_tutorial", true, nil},
	18: {18, "game_mode_tutorial", true, nil},
	// 3:  GameModeCompany,
	1:  {1, "game_mode_regular", false, nil},
	4:  {4, "game_mode_regular", true, []string{"tournament"}},
	5:  {5, "game_mode_regular", true, []string{"quick_tournament"}},
	7:  {7, "game_mode_rating", false, nil},
	8:  {8, "game_mode_arcade", true, nil},
	29: {29, "game_mode_arcade", true, []string{"tournament"}},
	36: {36, "game_mode_arcade", true, []string{"quick_tournament"}},
	// 21: GameModeVirtual,
	22: {22, "game_mode_scuffle", false, nil},
	30: {30, "game_mode_scuffle", false, []string{"tournament"}},
	37: {37, "game_mode_scuffle", false, []string{"quick_tournament"}},
	23: {23, "game_mode_hellgames", true, nil},
	28: {23, "game_mode_hellgames", true, []string{"tournament"}},
	35: {35, "game_mode_hellgames", true, []string{"quick_tournament"}},
	24: {24, "game_mode_lunar", true, nil},
	31: {31, "game_mode_lunar", true, []string{"tournament"}},
	38: {38, "game_mode_lunar", true, []string{"quick_tournament"}},
	25: {25, "game_mode_nanomaps", false, nil},
	32: {32, "game_mode_nanomaps", false, []string{"tournament"}},
	39: {39, "game_mode_nanomaps", false, []string{"quick_tournament"}},
	26: {26, "game_mode_vampiric", true, nil},
	33: {33, "game_mode_vampiric", true, []string{"tournament"}},
	40: {40, "game_mode_vampiric", true, []string{"quick_tournament"}},
	27: {27, "game_mode_bossmode", true, nil},
	34: {34, "game_mode_bossmode", true, []string{"tournament"}},
	41: {41, "game_mode_bossmode", true, []string{"quick_tournament"}},
	42: {42, "game_mode_gravitizing", true, nil},
	43: {43, "game_mode_gravitizing", true, []string{"tournament"}},
	44: {42, "game_mode_gravitizing", true, []string{"quick_tournament"}},
	45: {45, "game_mode_tenvsten", false, nil},
	46: {45, "game_mode_tenvsten", false, []string{"tournament"}},
	47: {45, "game_mode_tenvsten", false, []string{"quick_tournament"}},
}

type Replay struct {
	MapID      string     `json:"mapId" protobuf:"10"`
	GameMode   gameMode   `json:"gameMode" protobuf:"11"`
	BattleType battleType `json:"battleType" protobuf:"12"`

	Outcome        Outcome   `json:"victory" protobuf:"15"`
	BattleTime     time.Time `json:"battleTime" protobuf:"16"`
	BattleDuration int       `json:"battleDuration" protobuf:"17"`

	Spoils      Spoils `json:"spoils" protobuf:"20"`
	Protagonist Player `json:"protagonist" protobuf:"21"`

	Teams Teams `json:"teams" protobuf:"22"`
}

func Prettify(battle battleResults, meta replayMeta) Replay {
	var replay Replay

	replay.GameMode = gameModes[-1]
	if gm, ok := gameModes[int(battle.RoomType)]; ok {
		replay.GameMode = gm
	}

	// ModeAndMap
	replay.BattleType = BattleTypeUnknown
	if bt, ok := battleTypes[battle.GameMode()]; ok {
		replay.BattleType = bt
	}

	replay.MapID = fmt.Sprint(battle.MapID())
	ts, _ := strconv.ParseInt(meta.BattleStartTime, 10, 64)
	replay.BattleTime = time.Unix(ts, 0)
	replay.BattleDuration = int(meta.BattleDuration)

	replay.Spoils = Spoils{
		Exp:     frame.ValueInt(battle.Author.TotalXP),
		Credits: frame.ValueInt(battle.Author.TotalCredits),
		// TODO: Find where mastery is set
		// MasteryBadge: data.MasteryBadge,
	}

	var allyTeam, enemyTeam uint32
	players := make(map[int]playerInfo)
	for _, p := range battle.Players {
		players[int(p.AccountID)] = p.Info
		if p.AccountID == battle.Author.AccountID {
			allyTeam = p.Info.Team
		} else {
			enemyTeam = p.Info.Team
		}
	}
	for _, result := range battle.PlayerResults {
		info, ok := players[int(result.Info.AccountID)]
		if !ok {
			continue
		}
		player := playerFromData(battle, info, result.Info)
		if player.ID == fmt.Sprint(battle.Author.AccountID) {
			replay.Protagonist = player
		}
		if info.Team == allyTeam {
			replay.Teams.Allies = append(replay.Teams.Allies, player)
		} else {
			replay.Teams.Enemies = append(replay.Teams.Enemies, player)
		}
	}

	replay.Outcome = OutcomeDraw
	if battle.WinnerTeam == allyTeam {
		replay.Outcome = OutcomeVictory
	}
	if battle.WinnerTeam == enemyTeam {
		replay.Outcome = OutcomeDefeat
	}

	return replay
}

type Teams struct {
	Allies  []Player `json:"allies" protobuf:"10,repeated"`
	Enemies []Player `json:"enemies" protobuf:"11,repeated"`
}

type Player struct {
	ID       string `json:"id" protobuf:"10"`
	ClanID   string `json:"clanId" protobuf:"11"`
	ClanTag  string `json:"clanTag" protobuf:"12"`
	Nickname string `json:"nickname" protobuf:"13"`

	VehicleID string         `json:"vehicleId" protobuf:"15"`
	PlatoonID *int           `json:"platoonId" protobuf:"16"`
	TimeAlive frame.ValueInt `json:"timeAlive" protobuf:"17"`
	HPLeft    frame.ValueInt `json:"hpLeft" protobuf:"18"`

	Performance  Performance    `json:"performance" protobuf:"20"`
	Achievements map[string]int `json:"achievements" protobuf:"21"`
}

func playerFromData(battle battleResults, info playerInfo, result playerResultsInfo) Player {
	var player Player
	player.ID = fmt.Sprint(result.AccountID)
	player.Nickname = info.Nickname
	player.VehicleID = fmt.Sprint(result.TankID)
	if info.ClanTag != nil && info.ClanID != nil {
		player.ClanTag = *info.ClanTag
		player.ClanID = fmt.Sprint(*info.ClanID)
	}

	if info.PlatoonID != nil {
		id := int(*info.PlatoonID)
		player.PlatoonID = &id
	}

	var stats frame.StatsFrame
	stats.Battles = 1
	if info.Team == battle.WinnerTeam {
		stats.BattlesWon = 1
	}

	if result.HitpointsLeft != nil {
		player.HPLeft = frame.ValueInt(*result.HitpointsLeft)
	}
	if player.HPLeft > 0 {
		stats.BattlesSurvived = 1
	}

	stats.DamageDealt = frame.ValueInt(result.DamageDealt)
	stats.DamageReceived = frame.ValueInt(result.DamageReceived)
	stats.ShotsHit = frame.ValueInt(result.ShotsHit)
	stats.ShotsFired = frame.ValueInt(result.ShotsFired)
	stats.Frags = frame.ValueInt(result.EnemiesDestroyed)
	stats.MaxFrags = frame.ValueInt(stats.Frags)
	stats.EnemiesSpotted = frame.ValueInt(result.DamageAssisted)
	// TODO: Parse this from replays, it seems that those fields are only present when a battle was won by cap
	// frame.CapturePoints =
	// frame.DroppedCapturePoints =
	player.Performance = Performance{
		DamageBlocked:    frame.ValueInt(result.DamageBlocked),
		DamageReceived:   frame.ValueInt(result.DamageReceived),
		DamageAssisted:   frame.ValueInt(result.DamageAssisted + result.DamageAssistedTrack),
		DistanceTraveled: frame.ValueInt(result.DistanceTraveled),
		StatsFrame:       stats,
	}

	// player.Achievements = make(map[int]int)
	// for _, a := range append(result.Achievements, result.AchievementsOther...) {
	// 	player.Achievements[int(a.Tag)] = int(a.Value)
	// }

	return player
}

type Performance struct {
	DamageBlocked         frame.ValueInt `json:"damageBlocked" protobuf:"10"`
	DamageReceived        frame.ValueInt `json:"damageReceived" protobuf:"11"`
	DamageAssisted        frame.ValueInt `json:"damageAssisted" protobuf:"12"`
	DistanceTraveled      frame.ValueInt `json:"distanceTraveled" protobuf:"13"`
	SupremacyPointsEarned frame.ValueInt `json:"supremacyPointsEarned" protobuf:"14"`
	SupremacyPointsStolen frame.ValueInt `json:"supremacyPointsStolen" protobuf:"15"`

	frame.StatsFrame `json:",inline" protobuf:"20"`
}

type Spoils struct {
	Exp          frame.ValueInt `json:"exp" protobuf:"10"`
	Credits      frame.ValueInt `json:"credits" protobuf:"11"`
	MasteryBadge int            `json:"masteryBadge" protobuf:"12"`
}
