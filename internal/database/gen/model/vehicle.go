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

type Vehicle struct {
	ID             string    `sql:"primary_key" db:"id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
	Tier           int32     `db:"tier"`
	LocalizedNames []byte    `db:"localized_names"`
}