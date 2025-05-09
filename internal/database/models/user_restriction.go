package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
)

type UserRestrictionType string

const (
	RestrictionTypePartial  UserRestrictionType = "partial"  // restrict a specific permissions.Permissions value
	RestrictionTypeComplete UserRestrictionType = "complete" // restricts user from using any and all features
)

func (UserRestrictionType) Values() []string {
	return []string{string(RestrictionTypePartial), string(RestrictionTypeComplete)}
}

type UserRestriction struct {
	ID     string
	Type   UserRestrictionType
	UserID string

	Restriction      permissions.Permissions
	PublicReason     string
	ModeratorComment string

	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time

	Events []RestrictionUpdate
}

func (r *UserRestriction) AddEvent(modID string, summary string, context string) {
	r.Events = append(r.Events, RestrictionUpdate{
		ModeratorID: modID,
		Summary:     summary,
		Context:     context,
	})
}

type RestrictionUpdate struct {
	ModeratorID string
	Summary     string
	Context     string
}

func ToUserRestriction(record *model.UserRestriction) UserRestriction {
	r := UserRestriction{
		ID:     record.ID,
		Type:   UserRestrictionType(record.Type),
		UserID: record.UserID,

		CreatedAt: StringToTime(record.CreatedAt),
		UpdatedAt: StringToTime(record.UpdatedAt),
		ExpiresAt: StringToTime(record.ExpiresAt),

		ModeratorComment: record.ModeratorComment,
		PublicReason:     record.PublicReason,
		Restriction:      permissions.Parse(record.Restriction, permissions.Blank),

		Events: make([]RestrictionUpdate, 0),
	}
	json.Unmarshal(record.Events, &r.Events)
	return r
}

func (record UserRestriction) Model() model.UserRestriction {
	r := model.UserRestriction{
		ID:     utils.StringOr(record.ID, cuid.New()),
		Type:   string(record.Type),
		UserID: record.UserID,

		CreatedAt: TimeToString(record.CreatedAt),
		UpdatedAt: TimeToString(record.UpdatedAt),
		ExpiresAt: TimeToString(record.ExpiresAt),

		ModeratorComment: record.ModeratorComment,
		PublicReason:     record.PublicReason,
		Restriction:      record.Restriction.Encode(),
	}
	r.Events, _ = json.Marshal(record.Events)
	return r
}
