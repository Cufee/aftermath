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

func (a *Account) Model() model.Account {
	model := model.Account{
		ID:               a.ID,
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        time.Now().Unix(),
		LastBattleTime:   a.LastBattleTime.Unix(),
		AccountCreatedAt: a.CreatedAt.Unix(),
		Realm:            a.Realm,
		Nickname:         a.Nickname,
		Private:          a.Private,
	}
	if a.ClanID != "" {
		model.ClanID = &a.ClanID
	}
	return model
}

func ToAccount(r *model.Account, clan *model.Clan) Account {
	account := Account{
		ID:             r.ID,
		Realm:          r.Realm,
		Nickname:       r.Nickname,
		Private:        r.Private,
		CreatedAt:      time.Unix(r.AccountCreatedAt, 0),
		LastBattleTime: time.Unix(r.LastBattleTime, 0),
	}
	if clan != nil {
		account.ClanID = clan.ID
		account.ClanTag = clan.Tag
	}
	return account
}
