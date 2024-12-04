package database

import (
	"context"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
)

func TestAccounts(t *testing.T) {
	client := MustTestClient(t)
	is := is.New(t)

	t.Run("upsert and check a new account", func(t *testing.T) {
		errors, err := client.UpsertAccounts(context.Background(), &models.Account{
			ID:       "id-1",
			Realm:    "realm",
			Nickname: "nickname-1",

			Private:        false,
			CreatedAt:      time.Now(),
			LastBattleTime: time.Now(),
		})
		is.NoErr(err)
		for _, err := range errors {
			is.NoErr(err)
		}

		account, err := client.GetAccountByID(context.Background(), "id-1")
		is.NoErr(err)
		is.True(account.Nickname == "nickname-1")

		errors, err = client.UpsertAccounts(context.Background(), &models.Account{
			ID:       "id-1",
			Realm:    "realm",
			Nickname: "nickname-2",

			Private:        false,
			CreatedAt:      time.Now(),
			LastBattleTime: time.Now(),
		})
		is.NoErr(err)
		for _, err := range errors {
			is.NoErr(err)
		}

		account, err = client.GetAccountByID(context.Background(), "id-1")
		is.NoErr(err)
		is.True(account.Nickname == "nickname-2")
	})

	t.Run("get multiple accounts", func(t *testing.T) {
		errors, err := client.UpsertAccounts(context.Background(),
			&models.Account{
				ID:       "id-21",
				Realm:    "realm",
				Nickname: "nickname-21",

				Private:        false,
				CreatedAt:      time.Now(),
				LastBattleTime: time.Now(),
			},
			&models.Account{
				ID:       "id-22",
				Realm:    "realm",
				Nickname: "nickname-22",

				Private:        false,
				CreatedAt:      time.Now(),
				LastBattleTime: time.Now(),
			})
		is.NoErr(err)
		for _, err := range errors {
			is.NoErr(err)
		}

		accounts, err := client.GetAccounts(context.Background(), []string{"id-21", "id-22"})
		is.NoErr(err)
		is.True(len(accounts) == 2)
		is.True(accounts[0].ID != accounts[1].ID)
		is.True(accounts[0].Realm == accounts[1].Realm)
		is.True(accounts[0].Nickname != accounts[1].Nickname)
	})

	t.Run("set account to private", func(t *testing.T) {
		errors, err := client.UpsertAccounts(context.Background(), &models.Account{
			ID:       "id-10",
			Realm:    "realm",
			Nickname: "nickname-10",

			Private:        false,
			CreatedAt:      time.Now(),
			LastBattleTime: time.Now(),
		})
		is.NoErr(err)
		for _, err := range errors {
			is.NoErr(err)
		}

		account, err := client.GetAccountByID(context.Background(), "id-10")
		is.NoErr(err)
		is.True(!account.Private)

		err = client.AccountSetPrivate(context.Background(), "id-10", true)
		is.NoErr(err)
		for _, err := range errors {
			is.NoErr(err)
		}

		account, err = client.GetAccountByID(context.Background(), "id-10")
		is.NoErr(err)
		is.True(account.Private)
	})

}
