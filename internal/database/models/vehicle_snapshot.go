package models

import (
	"time"

	"github.com/cufee/aftermath/internal/stats/frame"
)

type SnapshotType string

const (
	SnapshotTypeLive   SnapshotType = "live"
	SnapshotTypeDaily  SnapshotType = "daily"
	SnapshotTypeWidget SnapshotType = "widget"
)

// Values provides list valid values for Enum.
func (SnapshotType) Values() []string {
	var kinds []string
	for _, s := range []SnapshotType{
		SnapshotTypeLive,
		SnapshotTypeDaily,
	} {
		kinds = append(kinds, string(s))
	}
	return kinds
}

type VehicleSnapshot struct {
	ID        string
	CreatedAt time.Time

	Type           SnapshotType
	LastBattleTime time.Time

	AccountID   string
	VehicleID   string
	ReferenceID string

	Stats frame.StatsFrame
}
