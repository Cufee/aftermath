package database

import (
	"context"
	"strings"
	"time"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/models"
)

func toAccount(model db.Account) models.Account {
	return models.Account{
		ID:             model.ID,
		Realm:          model.Realm,
		Nickname:       model.Nickname,
		Private:        model.Private,
		CreatedAt:      time.Unix(int64(model.AccountCreatedAt), 0),
		LastBattleTime: time.Unix(int64(model.LastBattleTime), 0),
	}
}

func (c *libsqlClient) GetRealmAccountIDs(ctx context.Context, realm string) ([]string, error) {
	result, err := c.db.Account.Query().Where(account.Realm(strings.ToLower(realm))).Select(account.FieldID).All(ctx)
	if err != nil {
		return nil, err
	}

	var accounts []string
	for _, model := range result {
		accounts = append(accounts, model.ID)
	}

	return accounts, nil
}

func (c *libsqlClient) GetAccountByID(ctx context.Context, id string) (models.Account, error) {
	// model, err := c.prisma.Account.FindUnique(db.Account.ID.Equals(id)).With(db.Account.Clan.Fetch()).Exec(ctx)
	// if err != nil {
	// 	return Account{}, err
	// }

	// account := Account{}.FromModel(*model)
	// if clan, ok := model.Clan(); ok {
	// 	account.AddClan(clan)
	// }

	return models.Account{}, nil
}

func (c *libsqlClient) GetAccounts(ctx context.Context, ids []string) ([]models.Account, error) {
	// if len(ids) < 1 {
	// 	return nil, nil
	// }

	// models, err := c.prisma.Account.FindMany(db.Account.ID.In(ids)).With(db.Account.Clan.Fetch()).Exec(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	var accounts []models.Account
	// for _, model := range models {
	// 	account := Account{}.FromModel(model)
	// 	if clan, ok := model.Clan(); ok {
	// 		account.AddClan(clan)
	// 	}
	// 	accounts = append(accounts, account)
	// }

	return accounts, nil
}

func (c *libsqlClient) UpsertAccounts(ctx context.Context, accounts []models.Account) map[string]error {
	// if len(accounts) < 1 {
	// 	return nil
	// }

	// var mx sync.Mutex
	// var wg sync.WaitGroup
	errors := make(map[string]error, len(accounts))

	// // we don't really want to exit if one fails
	// for _, account := range accounts {
	// 	wg.Add(1)
	// 	go func(account Account) {
	// 		defer wg.Done()
	// 		optional := []db.AccountSetParam{db.Account.Private.Set(account.Private)}
	// 		if account.ClanID != "" {
	// 			clan, err := c.prisma.Clan.FindUnique(db.Clan.ID.Equals(account.ClanID)).Exec(ctx)
	// 			if err == nil {
	// 				optional = append(optional, db.Account.Clan.Link(db.Clan.ID.Equals(clan.ID)))
	// 			}
	// 		}

	// 		_, err := c.prisma.Account.
	// 			UpsertOne(db.Account.ID.Equals(account.ID)).
	// 			Create(db.Account.ID.Set(account.ID),
	// 				db.Account.LastBattleTime.Set(account.LastBattleTime),
	// 				db.Account.AccountCreatedAt.Set(account.CreatedAt),
	// 				db.Account.Realm.Set(account.Realm),
	// 				db.Account.Nickname.Set(account.Nickname),
	// 				optional...,
	// 			).
	// 			Update(
	// 				append(optional,
	// 					db.Account.Nickname.Set(account.Nickname),
	// 					db.Account.LastBattleTime.Set(account.LastBattleTime),
	// 				)...,
	// 			).
	// 			Exec(ctx)
	// 		if err != nil {
	// 			mx.Lock()
	// 			errors[account.ID] = err
	// 			mx.Unlock()
	// 			return
	// 		}
	// 	}(account)
	// }
	// wg.Wait()

	return errors
}

func (c *libsqlClient) AccountSetPrivate(ctx context.Context, id string, value bool) error {
	// _, err := c.prisma.Account.FindUnique(db.Account.ID.Equals(id)).Update(db.Account.Private.Set(value)).Exec(ctx)
	// if err != nil && !database.IsNotFound(err) {
	// 	return err
	// }
	return nil
}
