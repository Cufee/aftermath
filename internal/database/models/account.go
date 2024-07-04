package models

import "time"

type Account struct {
	ID       string `json:"id"`
	Realm    string `json:"realm"`
	Nickname string `json:"nickname"`

	Private        bool      `json:"private"`
	CreatedAt      time.Time `json:"createdAt"`
	LastBattleTime time.Time `json:"lastBattleTime"`

	ClanID  string `json:"clanId"`
	ClanTag string `json:"clanTag"`
}
