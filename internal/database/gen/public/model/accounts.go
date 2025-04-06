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

type Accounts struct {
	ID               string     `sql:"primary_key" db:"id"`
	CreatedAt        *time.Time `db:"created_at"`
	UpdatedAt        *time.Time `db:"updated_at"`
	LastBattleTime   *time.Time `db:"last_battle_time"`
	AccountCreatedAt *time.Time `db:"account_created_at"`
	Realm            *string    `db:"realm"`
	Nickname         *string    `db:"nickname"`
	Private          *bool      `db:"private"`
	ClanID           *string    `db:"clan_id"`
}
