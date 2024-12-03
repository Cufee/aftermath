package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
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

func ToVehicleSnapshot(record *model.VehicleSnapshot) VehicleSnapshot {
	s := VehicleSnapshot{
		ID:             record.ID,
		Type:           SnapshotType(record.Type),
		CreatedAt:      record.CreatedAt,
		LastBattleTime: record.LastBattleTime,
		ReferenceID:    record.ReferenceID,
		AccountID:      record.AccountID,
		VehicleID:      record.VehicleID,
	}
	json.Unmarshal([]byte(record.Frame), &s.Stats)
	return s
}

func (record *VehicleSnapshot) Model() model.VehicleSnapshot {
	s := model.VehicleSnapshot{
		ID:             utils.StringOr(record.ID, cuid.New()),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Type:           string(record.Type),
		VehicleID:      record.VehicleID,
		ReferenceID:    record.ReferenceID,
		Battles:        int32(record.Stats.Battles),
		LastBattleTime: record.LastBattleTime,
		AccountID:      record.AccountID,
	}
	frame, _ := json.Marshal(record.Stats)
	s.Frame = string(frame)
	return s
}
