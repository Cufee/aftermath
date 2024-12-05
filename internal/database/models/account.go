package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

type Account struct {
	ID       string      `json:"id"`
	Realm    types.Realm `json:"realm"`
	Nickname string      `json:"nickname"`

	Private        bool      `json:"private"`
	CreatedAt      time.Time `json:"createdAt"`
	LastBattleTime time.Time `json:"lastBattleTime"`

	ClanID  string `json:"clanId"`
	ClanTag string `json:"clanTag"`
}

func (a *Account) Model() model.Account {
	model := model.Account{
		ID:               a.ID,
		CreatedAt:        TimeToString(time.Now()),
		UpdatedAt:        TimeToString(time.Now()),
		LastBattleTime:   TimeToString(a.LastBattleTime),
		AccountCreatedAt: TimeToString(a.CreatedAt),
		Realm:            a.Realm.String(),
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
		Realm:          types.Realm(r.Realm),
		Nickname:       r.Nickname,
		Private:        r.Private,
		CreatedAt:      StringToTime(r.AccountCreatedAt),
		LastBattleTime: StringToTime(r.LastBattleTime),
	}
	if clan != nil {
		account.ClanID = clan.ID
		account.ClanTag = clan.Tag
	}
	return account
}
