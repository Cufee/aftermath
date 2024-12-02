package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/permissions"
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

		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		ExpiresAt: record.ExpiresAt,

		ModeratorComment: record.ModeratorComment,
		PublicReason:     record.PublicReason,
		Restriction:      permissions.Parse(record.Restriction, permissions.Blank),

		Events: make([]RestrictionUpdate, 0),
	}
	json.Unmarshal([]byte(record.Events), &r.Events)
	return r
}
