package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/cufee/am-wg-proxy-next/v2/types"
	"github.com/lucsky/cuid"
)

type AccountAchievementsSnapshot struct {
	ID        string
	CreatedAt time.Time

	Type           SnapshotType
	Battles        int
	LastBattleTime time.Time

	AccountID   string
	ReferenceID string

	Stats types.AchievementsFrame
}

func ToAccountAchievementsSnapshot(record *model.AccountAchievementsSnapshot) AccountAchievementsSnapshot {
	s := AccountAchievementsSnapshot{
		ID:             record.ID,
		Type:           SnapshotType(record.Type),
		CreatedAt:      StringToTime(record.CreatedAt),
		LastBattleTime: StringToTime(record.LastBattleTime),
		ReferenceID:    record.ReferenceID,
		AccountID:      record.AccountID,
	}
	json.Unmarshal(record.Frame, &s.Stats)
	return s
}

func (record *AccountAchievementsSnapshot) Model() model.AccountAchievementsSnapshot {
	s := model.AccountAchievementsSnapshot{
		ID:             utils.StringOr(record.ID, cuid.New()),
		CreatedAt:      TimeToString(record.CreatedAt),
		Type:           string(record.Type),
		Battles:        int32(record.Battles),
		LastBattleTime: TimeToString(record.LastBattleTime),
		AccountID:      record.AccountID,
		ReferenceID:    record.ReferenceID,
	}
	s.Frame, _ = json.Marshal(record.Stats)
	return s
}
