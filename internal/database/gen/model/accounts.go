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
	ID               string `sql:"primary_key"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	LastBattleTime   time.Time
	AccountCreatedAt time.Time
	Realm            string
	Nickname         string
	Private          bool
	ClanID           *string
}
