package database

import (
	"context"
	"strings"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/models"
)

func toAccount(model *db.Account) models.Account {
	account := models.Account{
		ID:             model.ID,
		Realm:          model.Realm,
		Nickname:       model.Nickname,
		Private:        model.Private,
		CreatedAt:      model.AccountCreatedAt,
		LastBattleTime: model.LastBattleTime,
	}
	if model.Edges.Clan != nil {
		account.ClanID = model.Edges.Clan.ID
		account.ClanTag = model.Edges.Clan.Tag
	}
	return account
}

func (c *client) GetRealmAccountIDs(ctx context.Context, realm string) ([]string, error) {
	result, err := c.db.Account.Query().Where(account.Realm(strings.ToUpper(realm))).IDs(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *client) GetAccountByID(ctx context.Context, id string) (models.Account, error) {
	result, err := c.db.Account.Query().Where(account.ID(id)).WithClan().First(ctx)
	if err != nil {
		return models.Account{}, err
	}
	return toAccount(result), nil
}

func (c *client) GetAccounts(ctx context.Context, ids []string) ([]models.Account, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	result, err := c.db.Account.Query().Where(account.IDIn(ids...)).WithClan().All(ctx)
	if err != nil {
		return nil, err
	}

	var accounts []models.Account
	for _, a := range result {
		accounts = append(accounts, toAccount(a))
	}

	return accounts, nil
}

func (c *client) UpsertAccounts(ctx context.Context, accounts ...*models.Account) (map[string]error, error) {
	if len(accounts) < 1 {
		return nil, nil
	}

	var ids []string
	accountsMap := make(map[string]*models.Account)
	for _, a := range accounts {
		ids = append(ids, a.ID)
		accountsMap[a.ID] = a
	}

	records, err := c.db.Account.Query().Where(account.IDIn(ids...)).All(ctx)
	if err != nil && !IsNotFound(err) {
		return nil, err
	}

	errors := make(map[string]error)
	return errors, c.withTx(ctx, func(tx *db.Tx) error {
		for _, r := range records {
			update, ok := accountsMap[r.ID]
			if !ok {
				continue // this should never happen tho
			}

			err = tx.Account.UpdateOneID(r.ID).
				SetRealm(strings.ToUpper(update.Realm)).
				SetNickname(update.Nickname).
				SetPrivate(update.Private).
				SetLastBattleTime(update.LastBattleTime).
				Exec(ctx)
			if err != nil {
				errors[r.ID] = err
			}

			delete(accountsMap, r.ID)
		}

		var writes []*db.AccountCreate
		for _, a := range accountsMap {
			writes = append(writes, tx.Account.Create().
				SetID(a.ID).
				SetRealm(strings.ToUpper(a.Realm)).
				SetNickname(a.Nickname).
				SetPrivate(a.Private).
				SetAccountCreatedAt(a.CreatedAt).
				SetLastBattleTime(a.LastBattleTime),
			)
		}

		return tx.Account.CreateBulk(writes...).Exec(ctx)
	})
}

func (c *client) AccountSetPrivate(ctx context.Context, id string, value bool) error {
	err := c.db.Account.UpdateOneID(id).SetPrivate(value).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
