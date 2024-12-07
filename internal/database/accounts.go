package database

import (
	"context"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/sqlite"
)

type accountWithClan struct {
	m.Account
	Clan *m.Clan
}

func (c *client) GetRealmAccountIDs(ctx context.Context, realm string) ([]string, error) {
	stmt := t.Account.
		SELECT(t.Account.ID).
		WHERE(t.Account.Realm.EQ(s.UPPER(s.String(realm))))

	var result []m.Account
	err := c.query(ctx, stmt, &result)
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
	stmt := s.
		SELECT(t.Account.AllColumns, t.Clan.ID, t.Clan.Tag).
		FROM(t.Account.LEFT_JOIN(t.Clan, t.Clan.ID.EQ(t.Account.ClanID))).
		WHERE(t.Account.ID.EQ(s.String(id))).
		LIMIT(1)

	var result accountWithClan
	err := c.query(ctx, stmt, &result)
	if err != nil {
		return models.Account{}, err
	}

	return models.ToAccount(&result.Account, result.Clan), nil
}

func (c *client) GetAccounts(ctx context.Context, ids []string) ([]models.Account, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	stmt := s.
		SELECT(t.Account.AllColumns, t.Clan.ID, t.Clan.Tag).
		FROM(t.Account.LEFT_JOIN(t.Clan, t.Clan.ID.EQ(t.Account.ClanID))).
		WHERE(t.Account.ID.IN(stringsToExp(ids)...))

	var result []accountWithClan
	err := c.query(ctx, stmt, &result)
	if err != nil {
		return nil, err
	}

	var accounts []models.Account
	for _, a := range result {
		accounts = append(accounts, models.ToAccount(&a.Account, a.Clan))
	}

	return accounts, nil
}

func (c *client) UpsertAccounts(ctx context.Context, accounts ...*models.Account) (map[string]error, error) {
	if len(accounts) < 1 {
		return nil, nil
	}

	errors := make(map[string]error, len(accounts))
	return errors, c.withTx(ctx, func(tx *transaction) error {
		for _, a := range accounts {
			stmt := t.Account.
				INSERT(t.Account.AllColumns).
				MODEL(a.Model()).
				ON_CONFLICT(t.Account.ID).
				DO_UPDATE(
					s.SET(
						t.Account.UpdatedAt.SET(t.Account.EXCLUDED.UpdatedAt),
						t.Account.ClanID.SET(t.Account.EXCLUDED.ClanID),
						t.Account.Private.SET(t.Account.EXCLUDED.Private),
						t.Account.Nickname.SET(t.Account.EXCLUDED.Nickname),
						t.Account.LastBattleTime.SET(t.Account.EXCLUDED.LastBattleTime),
					),
				)
			_, errors[a.ID] = tx.exec(ctx, stmt)
		}
		return nil
	})

}

func (c *client) AccountSetPrivate(ctx context.Context, id string, value bool) error {
	stmt := t.Account.
		UPDATE(t.Account.Private).
		SET(
			t.Account.Private.SET(s.Bool(value)),
		).
		WHERE(t.Account.ID.EQ(s.String(id)))
	_, err := c.exec(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}
