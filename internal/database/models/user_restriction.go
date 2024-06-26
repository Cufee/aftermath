package models

import (
	"time"

	"github.com/cufee/aftermath/internal/permissions"
)

type UserRestrictionType string

const (
	RestrictionTypePartial  UserRestrictionType = "partial"  // restrict a specific permissions.Permissions value
	RestrictionTypeComplete UserRestrictionType = "complete" // restricts user from using any and all features
)

func (r *UserRestrictionType) Values() []string {
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
}
