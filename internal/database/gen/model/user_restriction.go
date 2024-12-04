//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type UserRestriction struct {
	ID               string    `sql:"primary_key" db:"id"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	ExpiresAt        time.Time `db:"expires_at"`
	Type             string    `db:"type"`
	Restriction      string    `db:"restriction"`
	PublicReason     string    `db:"public_reason"`
	ModeratorComment string    `db:"moderator_comment"`
	Events           string    `db:"events"`
	UserID           string    `db:"user_id"`
}
