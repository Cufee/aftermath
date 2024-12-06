package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/matryer/is"
)

func TestUsers(t *testing.T) {
	is := is.New(t)

	client := MustTestClient(t)

	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.UserContent.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.UserContent.TableName()))
	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.UserConnection.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.UserConnection.TableName()))
	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.UserSubscription.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.UserSubscription.TableName()))

	user, err := client.GetOrCreateUserByID(context.Background(), "user-TestUsers")
	is.NoErr(err)

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", table.User.TableName(), user.ID))

	t.Run("create and get user connections", func(t *testing.T) {
		is := is.New(t)

		verified := models.UserConnection{
			ID:          "id-1",
			Type:        models.ConnectionTypeWargaming,
			Verified:    true,
			Selected:    false,
			UserID:      user.ID,
			ReferenceID: "r-1",
		}
		verifiedInserted, err := client.CreateUserConnection(context.Background(), verified)
		is.NoErr(err)
		is.True(verifiedInserted.Verified)
		is.True(!verifiedInserted.Selected)
		is.True(verified.UserID == user.ID)

		selected := models.UserConnection{
			ID:          "id-2",
			Type:        models.ConnectionTypeWargaming,
			Verified:    false,
			Selected:    true,
			UserID:      user.ID,
			ReferenceID: "r-2",
		}
		selectedInserted, err := client.CreateUserConnection(context.Background(), selected)
		is.NoErr(err)
		is.True(selectedInserted.Selected)
		is.True(!selectedInserted.Verified)
		is.True(selectedInserted.UserID == user.ID)

		regular := models.UserConnection{
			ID:          "id-3",
			Type:        models.ConnectionTypeWargaming,
			Verified:    false,
			Selected:    false,
			UserID:      user.ID,
			ReferenceID: "r-3",
		}
		regularInserted, err := client.CreateUserConnection(context.Background(), regular)
		is.NoErr(err)
		is.True(!regularInserted.Verified)
		is.True(!regularInserted.Selected)
		is.True(regularInserted.UserID == user.ID)

		localUser, err := client.GetOrCreateUserByID(context.Background(), user.ID, WithConnections())
		is.NoErr(err)
		{
			found, ok := localUser.Connection(models.ConnectionTypeWargaming, utils.Pointer(true), nil)
			is.True(ok)
			is.True(found.Verified)
			is.True(found.ID == verified.ID)
		}
		{
			found, ok := localUser.Connection(models.ConnectionTypeWargaming, nil, utils.Pointer(true))
			is.True(ok)
			is.True(found.Selected)
			is.True(found.ID == selected.ID)
		}
		{
			found, ok := localUser.Connection(models.ConnectionTypeWargaming, utils.Pointer(false), utils.Pointer(false))
			is.True(ok)
			is.True(found.ID == regular.ID)
			is.True(!found.Selected && !found.Verified)
		}
	})

	t.Run("create and get user subscription", func(t *testing.T) {
		is := is.New(t)

		data := models.UserSubscription{
			ID:          "id-1",
			Type:        models.SubscriptionTypePlus,
			UserID:      user.ID,
			ExpiresAt:   time.Now().Add(time.Hour),
			ReferenceID: "r-1",
			Permissions: permissions.SubscriptionAftermathPlus,
		}
		inserted, err := client.CreateUserSubscription(context.Background(), data)
		is.NoErr(err)
		is.True(inserted.ID == data.ID)
		is.True(inserted.Permissions.Has(data.Permissions))

		data.ExpiresAt = time.Now().Add(time.Hour * -1)
		updated, err := client.UpdateUserSubscription(context.Background(), data.ID, data)
		is.NoErr(err)
		is.True(updated.ExpiresAt.Before(time.Now()))

		localUser, err := client.GetOrCreateUserByID(context.Background(), user.ID, WithSubscriptions())
		is.NoErr(err)

		_, ok := localUser.Subscription(models.SubscriptionTypePlus)
		is.True(!ok)
	})

	t.Run("create and get user content", func(t *testing.T) {
		is := is.New(t)

		data := models.UserContent{
			ID:          "id-1",
			Type:        models.UserContentTypePersonalBackground,
			UserID:      user.ID,
			ReferenceID: "ref-1",
			Value:       "value-1",
		}
		inserted, err := client.CreateUserContent(context.Background(), data)
		is.NoErr(err)
		is.True(inserted.Value == data.Value)
		is.True(inserted.ReferenceID == data.ReferenceID)
	})

}
