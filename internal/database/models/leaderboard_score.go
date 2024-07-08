package models

import (
	"time"
)

type LeaderboardID string

// Values provides list valid values for Enum.
func (LeaderboardID) Values() []string {
	var kinds []string
	for _, s := range []LeaderboardID{
		LeaderboardsMasteryWeightedPlayers,
		LeaderboardsMasteryWeightedClans,
	} {
		kinds = append(kinds, string(s))
	}
	return kinds
}

type ScoreType string

// Values provides list valid values for Enum.
func (ScoreType) Values() []string {
	var kinds []string
	for _, s := range []ScoreType{
		LeaderboardScoreCustom,
		LeaderboardScoreHourly,
		LeaderboardScoreDaily,
	} {
		kinds = append(kinds, string(s))
	}
	return kinds
}

const (
	LeaderboardsMasteryWeightedPlayers LeaderboardID = "mastery-weighted-players"
	LeaderboardsMasteryWeightedClans   LeaderboardID = "mastery-weighted-clans"

	LeaderboardScoreCustom ScoreType = "custom"
	LeaderboardScoreHourly ScoreType = "hourly"
	LeaderboardScoreDaily  ScoreType = "daily"
)

type LeaderboardScore struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	Type          ScoreType
	ReferenceID   string
	LeaderboardID LeaderboardID

	Score float32
	Meta  map[string]any
}
