package models

import "time"

type Account struct {
	ID       string
	Realm    string
	Nickname string

	Private        bool
	CreatedAt      time.Time
	LastBattleTime time.Time

	ClanID  string
	ClanTag string
}
