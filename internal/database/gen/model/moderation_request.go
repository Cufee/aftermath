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

type ModerationRequest struct {
	ID               string    `sql:"primary_key" db:"id"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
	ModeratorComment *string   `db:"moderator_comment"`
	Context          *string   `db:"context"`
	ReferenceID      string    `db:"reference_id"`
	ActionReason     *string   `db:"action_reason"`
	ActionStatus     string    `db:"action_status"`
	Data             []byte    `db:"data"`
	RequestorID      string    `db:"requestor_id"`
	ModeratorID      *string   `db:"moderator_id"`
}
