package database

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
)

func TestAuthNonce(t *testing.T) {
	client := MustTestClient(t)
	is := is.New(t)

	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.AuthNonce.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.AuthNonce.TableName()))

	t.Run("all auth fields are set correctly", func(t *testing.T) {
		is := is.New(t)

		nonce, err := client.CreateAuthNonce(context.Background(), "pid-1", "id-1", time.Now().Add(time.Hour), nil)
		is.NoErr(err)
		is.True(nonce.ID != "")
		is.True(nonce.Active == true)
		is.True(nonce.PublicID == "pid-1")
		is.True(nonce.Identifier == "id-1")
		is.True(nonce.ExpiresAt.After(time.Now().Add(time.Minute)))
	})
	t.Run("find auth by public id", func(t *testing.T) {
		is := is.New(t)

		nonce, err := client.CreateAuthNonce(context.Background(), "pid-2", "id-2", time.Now().Add(time.Hour), nil)
		is.NoErr(err)

		found, err := client.FindAuthNonce(context.Background(), nonce.PublicID)
		is.NoErr(err)
		is.True(found.PublicID == nonce.PublicID)

		_, err = client.FindAuthNonce(context.Background(), "invalid")
		is.True(IsNotFound(err))
	})
	t.Run("set nonce as not active", func(t *testing.T) {
		is := is.New(t)

		nonce, err := client.CreateAuthNonce(context.Background(), "pid-3", "id-3", time.Now().Add(time.Hour), nil)
		is.NoErr(err)

		err = client.SetAuthNonceActive(context.Background(), nonce.ID, false)
		is.NoErr(err)

		nonActive, err := client.FindAuthNonce(context.Background(), nonce.PublicID)
		is.True(errors.Is(err, models.ErrInvalidNonce))
		is.True(!nonActive.Active)

		err = client.SetAuthNonceActive(context.Background(), nonce.ID, true)
		is.NoErr(err)

		active, err := client.FindAuthNonce(context.Background(), nonce.PublicID)
		is.NoErr(err)
		is.True(active.Active)
	})
}

func TestSession(t *testing.T) {
	client := MustTestClient(t)
	is := is.New(t)

	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.Session.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.Session.TableName()))

	user, err := client.GetOrCreateUserByID(context.Background(), "user-TestSession")
	is.NoErr(err)

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", table.User.TableName(), user.ID))

	t.Run("all session fields are set correctly", func(t *testing.T) {
		is := is.New(t)

		session, err := client.CreateSession(context.Background(), "pid-10", user.ID, time.Now().Add(time.Hour), nil)
		is.NoErr(err)
		is.True(session.ID != "")
		is.True(session.UserID == user.ID)
		is.True(session.PublicID == "pid-10")
		is.True(session.ExpiresAt.After(time.Now().Add(time.Minute)))
	})
	t.Run("create and find a session", func(t *testing.T) {
		is := is.New(t)

		session, err := client.CreateSession(context.Background(), "pid-11", user.ID, time.Now().Add(time.Hour), nil)
		is.NoErr(err)

		found, err := client.FindSession(context.Background(), session.PublicID)
		is.NoErr(err)
		is.True(found.ID == session.ID)
	})
	t.Run("set session expiration", func(t *testing.T) {
		is := is.New(t)

		session, err := client.CreateSession(context.Background(), "pid-12", user.ID, time.Now().Add(time.Hour), nil)
		is.NoErr(err)

		{
			expired := time.Now().UTC()
			err = client.SetSessionExpiresAt(context.Background(), session.ID, expired)
			is.NoErr(err)

			updated, err := client.FindSession(context.Background(), session.PublicID)
			is.True(errors.Is(err, models.ErrSessionExpired))
			is.True(updated.ExpiresAt.Unix() == expired.Unix())
		}

		{
			nonExpired := time.Now().Add(time.Hour * 24).UTC()
			err = client.SetSessionExpiresAt(context.Background(), session.ID, nonExpired)
			is.NoErr(err)

			updated2, err := client.FindSession(context.Background(), session.PublicID)
			is.NoErr(err)
			is.True(updated2.ExpiresAt.Unix() == nonExpired.Unix())
		}
	})
}
