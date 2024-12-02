package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
)

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

func ToAccount(a *Account) model.Account {
	model := model.Account{
		ID:               a.ID,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
		LastBattleTime:   a.LastBattleTime,
		AccountCreatedAt: a.CreatedAt,
		Realm:            a.Realm,
		Nickname:         a.Nickname,
		Private:          a.Private,
	}
	if a.ClanID != "" {
		model.ClanID = &a.ClanID
	}
	return model
}

func FromAccount(r *model.Account, clan *model.Clan) Account {
	account := Account{
		ID:             r.ID,
		Realm:          r.Realm,
		Nickname:       r.Nickname,
		Private:        r.Private,
		CreatedAt:      r.AccountCreatedAt,
		LastBattleTime: r.LastBattleTime,
	}
	if clan != nil {
		account.ClanID = clan.ID
		account.ClanTag = clan.Tag
	}
	return account
}
