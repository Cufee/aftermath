package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/stats/frame"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
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

func ToAccountSnapshot(record *model.AccountSnapshot) AccountSnapshot {
	s := AccountSnapshot{
		ID:             record.ID,
		Type:           SnapshotType(record.Type),
		AccountID:      record.AccountID,
		ReferenceID:    record.ReferenceID,
		CreatedAt:      StringToTime(record.CreatedAt),
		LastBattleTime: StringToTime(record.LastBattleTime),
	}
	json.Unmarshal(record.RatingFrame, &s.RatingBattles)
	json.Unmarshal(record.RegularFrame, &s.RegularBattles)
	return s
}

func (record *AccountSnapshot) Model() model.AccountSnapshot {
	s := model.AccountSnapshot{
		ID:             utils.StringOr(record.ID, cuid.New()),
		CreatedAt:      TimeToString(record.CreatedAt),
		Type:           string(record.Type),
		LastBattleTime: TimeToString(record.LastBattleTime),
		ReferenceID:    record.ReferenceID,
		RatingBattles:  int32(record.RatingBattles.Battles),
		RegularBattles: int32(record.RegularBattles.Battles),
		AccountID:      record.AccountID,
	}
	s.RatingFrame, _ = json.Marshal(record.RatingBattles)
	s.RegularFrame, _ = json.Marshal(record.RegularBattles)

	return s
}
