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

type AuthNonces struct {
	ID         string     `sql:"primary_key" db:"id"`
	CreatedAt  *time.Time `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	Active     *bool      `db:"active"`
	ExpiresAt  *time.Time `db:"expires_at"`
	Identifier *string    `db:"identifier"`
	PublicID   *string    `db:"public_id"`
	Metadata   *string    `db:"metadata"`
}
