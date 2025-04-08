package database

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/matryer/is"
)

func TestModerationRequest(t *testing.T) {
	client := MustTestClient(t)
	is := is.New(t)

	client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.ModerationRequest.TableName()))
	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s;", table.ModerationRequest.TableName()))

	user, err := client.GetOrCreateUserByID(context.Background(), "user-TestModerationRequest")
	is.NoErr(err)

	defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", table.User.TableName(), user.ID))

	t.Run("fund user moderation request", func(t *testing.T) {
		is := is.New(t)

		localUser, err := client.GetOrCreateUserByID(context.Background(), "user-TestModerationRequest-1")
		is.NoErr(err)

		defer client.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = '%s';", table.User.TableName(), localUser.ID))

		data1 := models.ModerationRequest{
			ReferenceID:    "ref-2",
			RequestorID:    user.ID,
			RequestContext: "context-2",
			ActionStatus:   models.ModerationStatusSubmitted,
		}
		data2 := models.ModerationRequest{
			ReferenceID:    "ref-3",
			RequestorID:    localUser.ID,
			RequestContext: "context-3",
			ActionStatus:   models.ModerationStatusSubmitted,
		}

		_, err = client.CreateModerationRequest(context.Background(), data1)
		is.NoErr(err)
		_, err = client.CreateModerationRequest(context.Background(), data2)
		is.NoErr(err)
		{
			found, err := client.FindUserModerationRequests(context.Background(), localUser.ID, nil, nil, time.Now().Add(time.Hour*-24))
			is.NoErr(err)
			is.True(len(found) == 1)
			is.True(found[0].RequestorID == localUser.ID)
		}
	})

	t.Run("create and get a moderation request", func(t *testing.T) {
		is := is.New(t)

		data := models.ModerationRequest{
			ReferenceID:    "ref-1",
			RequestorID:    user.ID,
			RequestContext: "context-1",
			ActionStatus:   models.ModerationStatusSubmitted,
		}
		inserted, err := client.CreateModerationRequest(context.Background(), data)
		is.NoErr(err)
		is.True(inserted.RequestorID == user.ID)

		{
			found, err := client.GetModerationRequest(context.Background(), inserted.ID)
			is.NoErr(err)
			is.True(found.ID == inserted.ID)
		}
		{
			found, err := client.FindUserModerationRequests(context.Background(), user.ID, nil, nil, time.Now().Add(-time.Hour))
			is.NoErr(err)
			is.True(len(found) > 0)
		}
	})
	t.Run("update a moderation request", func(t *testing.T) {
		is := is.New(t)

		data := models.ModerationRequest{
			ReferenceID:    "ref-4",
			RequestorID:    user.ID,
			RequestContext: "context-4",
			ActionStatus:   models.ModerationStatusSubmitted,
		}
		inserted, err := client.CreateModerationRequest(context.Background(), data)
		is.NoErr(err)
		is.True(inserted.RequestorID == user.ID)
		{
			inserted.ActionStatus = models.ModerationStatusDeclined
			inserted.ModeratorComment = "comment-4"
			inserted.ModeratorID = &user.ID
			updated, err := client.UpdateModerationRequest(context.Background(), inserted)
			is.NoErr(err)
			is.True(updated.ActionStatus == inserted.ActionStatus)
			is.True(updated.ActionStatus != data.ActionStatus)
			is.True(updated.ModeratorComment == inserted.ModeratorComment)
			is.True(*updated.ModeratorID == *inserted.ModeratorID)
		}
	})

}
