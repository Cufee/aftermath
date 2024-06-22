package models

import (
	"time"

	"github.com/cufee/aftermath/internal/stats/frame"
)

type AccountSnapshot struct {
	ID             string
	Type           SnapshotType
	CreatedAt      time.Time
	AccountID      string
	ReferenceID    string
	LastBattleTime time.Time
	RatingBattles  frame.StatsFrame
	RegularBattles frame.StatsFrame
}
