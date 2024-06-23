package database

import (
	"context"
	"strings"
	"time"

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
		CreatedAt:      time.Unix(int64(model.AccountCreatedAt), 0),
		LastBattleTime: time.Unix(int64(model.LastBattleTime), 0),
	}
	if model.Edges.Clan != nil {
		account.ClanID = model.Edges.Clan.ID
		account.ClanTag = model.Edges.Clan.Tag
	}
	return account
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
	result, err := c.db.Account.Query().Where(account.ID(id)).WithClan().Only(ctx)
	if err != nil {
		return models.Account{}, err
	}
	return toAccount(result), nil
}

func (c *libsqlClient) GetAccounts(ctx context.Context, ids []string) ([]models.Account, error) {
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

func (c *libsqlClient) UpsertAccounts(ctx context.Context, accounts []models.Account) error {
	if len(accounts) < 1 {
		return nil
	}

	var ids []string
	accountsMap := make(map[string]*models.Account)
	for _, a := range accounts {
		ids = append(ids, a.ID)
		accountsMap[a.ID] = &a
	}

	tx, err := c.db.Tx(ctx)
	if err != nil {
		return err
	}

	records, err := tx.Account.Query().Where(account.IDIn(ids...)).All(ctx)
	if err != nil && !IsNotFound(err) {
		return rollback(tx, err)
	}

	for _, r := range records {
		update, ok := accountsMap[r.ID]
		if !ok {
			continue // this should never happen tho
		}

		err = tx.Account.UpdateOneID(r.ID).
			SetRealm(update.Realm).
			SetNickname(update.Nickname).
			SetPrivate(update.Private).
			SetLastBattleTime(update.LastBattleTime.Unix()).
			SetClanID(update.ClanID).
			Exec(ctx)
		if err != nil {
			return rollback(tx, err)
		}

		delete(accountsMap, r.ID)
	}

	var inserts []*db.AccountCreate
	for _, a := range accountsMap {
		inserts = append(inserts,
			c.db.Account.Create().
				SetID(a.ID).
				SetRealm(a.Realm).
				SetNickname(a.Nickname).
				SetPrivate(a.Private).
				SetAccountCreatedAt(a.CreatedAt.Unix()).
				SetLastBattleTime(a.LastBattleTime.Unix()).
				SetClanID(a.ClanID),
		)
	}

	err = c.db.Account.CreateBulk(inserts...).Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	return tx.Commit()
}

func (c *libsqlClient) AccountSetPrivate(ctx context.Context, id string, value bool) error {
	err := c.db.Account.UpdateOneID(id).SetPrivate(value).Exec(ctx)
	if err != nil && !IsNotFound(err) {
		return err
	}
	return nil
}
