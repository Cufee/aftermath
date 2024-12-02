package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/sqlite"
)

type accountWithClan struct {
	model.Accounts
	Clan *model.Clans
}

func toAccount(model *accountWithClan) models.Account {
	account := models.Account{
		ID:             model.ID,
		Realm:          model.Realm,
		Nickname:       model.Nickname,
		Private:        model.Private,
		CreatedAt:      model.AccountCreatedAt,
		LastBattleTime: model.LastBattleTime,
	}
	if model.Clan != nil {
		account.ClanID = model.Clan.ID
		account.ClanTag = model.Clan.Tag
	}
	return account
}

func (c *client) GetRealmAccountIDs(ctx context.Context, realm string) ([]string, error) {
	stmt := s.SELECT(table.Accounts.ID).
		FROM(table.Accounts).
		WHERE(table.Accounts.Realm.EQ(s.UPPER(s.String(realm))))

	var result []model.Accounts
	err := stmt.Query(c.db, &result)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, a := range result {
		ids = append(ids, a.ID)
	}

	return ids, nil
}

func (c *client) GetAccountByID(ctx context.Context, id string) (models.Account, error) {
	stmt := s.SELECT(table.Accounts.AllColumns, table.Clans.ID, table.Clans.Tag).
		FROM(table.Accounts.INNER_JOIN(table.Clans, table.Clans.ID.EQ(table.Accounts.ClanID))).
		WHERE(table.Accounts.ID.EQ(s.String(id)))

	var result accountWithClan
	err := stmt.Query(c.db, &result)
	if err != nil {
		return models.Account{}, err
	}
	return toAccount(&result), nil
}

func (c *client) GetAccounts(ctx context.Context, ids []string) ([]models.Account, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	stmt := s.SELECT(table.Accounts.AllColumns, table.Clans.ID, table.Clans.Tag).
		FROM(table.Accounts.INNER_JOIN(table.Clans, table.Clans.ID.EQ(table.Accounts.ClanID))).
		WHERE(table.Accounts.ID.IN(toStringSlice(ids...)...))

	var result []accountWithClan
	err := stmt.Query(c.db, &result)
	if err != nil {
		return nil, err
	}

	var accounts []models.Account
	for _, a := range result {
		accounts = append(accounts, toAccount(&a))
	}

	return accounts, nil
}

// func (c *client) UpsertAccounts(ctx context.Context, accounts ...*models.Account) (map[string]error, error) {
// 	if len(accounts) < 1 {
// 		return nil, nil
// 	}

// 	var ids []string
// 	accountsMap := make(map[string]*models.Account)
// 	for _, a := range accounts {
// 		ids = append(ids, a.ID)
// 		accountsMap[a.ID] = a
// 	}

// 	records, err := c.db.Account.Query().Where(account.IDIn(ids...)).All(ctx)
// 	if err != nil && !IsNotFound(err) {
// 		return nil, err
// 	}

// 	errors := make(map[string]error)
// 	return errors, c.withTx(ctx, func(tx *db.Tx) error {
// 		for _, r := range records {
// 			update, ok := accountsMap[r.ID]
// 			if !ok {
// 				continue // this should never happen tho
// 			}

// 			err = tx.Account.UpdateOneID(r.ID).
// 				SetRealm(strings.ToUpper(update.Realm)).
// 				SetNickname(update.Nickname).
// 				SetPrivate(update.Private).
// 				SetLastBattleTime(update.LastBattleTime).
// 				Exec(ctx)
// 			if err != nil {
// 				errors[r.ID] = err
// 			}

// 			delete(accountsMap, r.ID)
// 		}

// 		var writes []*db.AccountCreate
// 		for _, a := range accountsMap {
// 			writes = append(writes, tx.Account.Create().
// 				SetID(a.ID).
// 				SetRealm(strings.ToUpper(a.Realm)).
// 				SetNickname(a.Nickname).
// 				SetPrivate(a.Private).
// 				SetAccountCreatedAt(a.CreatedAt).
// 				SetLastBattleTime(a.LastBattleTime),
// 			)
// 		}

// 		return tx.Account.CreateBulk(writes...).Exec(ctx)
// 	})
// }

// func (c *client) AccountSetPrivate(ctx context.Context, id string, value bool) error {
// 	err := c.db.Account.UpdateOneID(id).SetPrivate(value).Exec(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
