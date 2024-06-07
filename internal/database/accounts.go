package database

import (
	"context"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
)

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

func (a Account) FromModel(model db.AccountModel) Account {
	a.ID = model.ID
	a.Realm = model.Realm
	a.Nickname = model.Nickname
	a.Private = model.Private
	a.CreatedAt = model.CreatedAt
	a.LastBattleTime = model.LastBattleTime
	return a
}

func (a *Account) AddClan(model *db.ClanModel) {
	a.ClanID = model.ID
	a.ClanTag = model.Tag
}

func (c *client) GetRealmAccounts(ctx context.Context, realm string) ([]Account, error) {
	models, err := c.prisma.Account.FindMany(db.Account.Realm.Equals(realm)).With(db.Account.Clan.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var accounts []Account
	for _, model := range models {
		account := Account{}.FromModel(model)
		if clan, ok := model.Clan(); ok {
			account.AddClan(clan)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (c *client) GetAccountByID(ctx context.Context, id string) (Account, error) {
	model, err := c.prisma.Account.FindUnique(db.Account.ID.Equals(id)).With(db.Account.Clan.Fetch()).Exec(ctx)
	if err != nil {
		return Account{}, err
	}

	account := Account{}.FromModel(*model)
	if clan, ok := model.Clan(); ok {
		account.AddClan(clan)
	}

	return account, nil
}

func (c *client) GetAccounts(ctx context.Context, ids []string) ([]Account, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	models, err := c.prisma.Account.FindMany(db.Account.ID.In(ids)).With(db.Account.Clan.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}

	var accounts []Account
	for _, model := range models {
		account := Account{}.FromModel(model)
		if clan, ok := model.Clan(); ok {
			account.AddClan(clan)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (c *client) UpsertAccounts(ctx context.Context, accounts []Account) map[string]error {
	if len(accounts) < 1 {
		return nil
	}

	var mx sync.Mutex
	var wg sync.WaitGroup
	errors := make(map[string]error, len(accounts))

	// we don't really want to exit if one fails
	for _, account := range accounts {
		wg.Add(1)
		go func(account Account) {
			defer wg.Done()
			optional := []db.AccountSetParam{db.Account.Private.Set(account.Private)}
			if account.ClanID != "" {
				optional = append(optional, db.Account.Clan.Link(db.Clan.ID.Equals(account.ClanID)))
			}

			_, err := c.prisma.Account.
				UpsertOne(db.Account.ID.Equals(account.ID)).
				Create(db.Account.ID.Set(account.ID),
					db.Account.LastBattleTime.Set(account.LastBattleTime),
					db.Account.AccountCreatedAt.Set(account.CreatedAt),
					db.Account.Realm.Set(account.Realm),
					db.Account.Nickname.Set(account.Nickname),
					optional...,
				).
				Exec(ctx)
			if err != nil {
				mx.Lock()
				errors[account.ID] = err
				mx.Unlock()
			}
		}(account)
	}
	wg.Wait()

	return errors
}
