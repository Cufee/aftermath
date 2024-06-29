package models

import (
	"time"

	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type AchievementsSnapshot struct {
	ID          string
	Type        SnapshotType
	CreatedAt   time.Time
	AccountID   string
	ReferenceID string // accountID or vehicleID

	Battles        int
	LastBattleTime time.Time
	Data           types.AchievementsFrame
}
