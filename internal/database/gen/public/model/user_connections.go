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

type UserConnections struct {
	ID          string     `sql:"primary_key" db:"id"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	Type        *string    `db:"type"`
	ReferenceID *string    `db:"reference_id"`
	Permissions *string    `db:"permissions"`
	Metadata    *string    `db:"metadata"`
	UserID      *string    `db:"user_id"`
}
