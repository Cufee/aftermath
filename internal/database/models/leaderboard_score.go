package models

import (
	"time"
)

type ScoreType string

// Values provides list valid values for Enum.
func (ScoreType) Values() []string {
	var kinds []string
	for _, s := range []ScoreType{
		LeaderboardScoreHourly,
		LeaderboardScoreDaily,
	} {
		kinds = append(kinds, string(s))
	}
	return kinds
}

const (
	LeaderboardsMasteryWeightedPlayers string = "mastery-weighted-players"
	LeaderboardsMasteryWeightedClans   string = "mastery-weighted-clans"

	LeaderboardScoreHourly ScoreType = "hourly"
	LeaderboardScoreDaily  ScoreType = "daily"
)

type LeaderboardScore struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	Type          ScoreType
	AccountID     string
	ReferenceID   string
	LeaderboardID string

	Score float32
	Meta  map[string]any
}
