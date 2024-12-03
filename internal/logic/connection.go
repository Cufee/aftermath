package logic

import (
	"context"
	"errors"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
)

var ErrConnectionNotFound = errors.New("connection not found")

/*
Updates all user connections, marking them as not default unless the reference id matches the provided reference id
  - As a side effect and if the new connection is verified, updates a user content record of type models.UserContentTypePersonalBackground to match the new reference id
*/
func UpdateDefaultUserConnection(ctx context.Context, db database.Client, userID string, referenceID string) (models.UserConnection, error) {
	if referenceID == "" {
		return models.UserConnection{}, errors.New("invalid reference id")
	}
	if userID == "" {
		return models.UserConnection{}, errors.New("invalid user id")
	}

	user, err := db.GetUserByID(ctx, userID, database.WithConnections(), database.WithContent())
	if err != nil {
		return models.UserConnection{}, err
	}

	var connection models.UserConnection
	for _, conn := range user.Connections {
		if conn.Type != models.ConnectionTypeWargaming {
			continue
		}
		if conn.ReferenceID == referenceID {
			connection = conn
		}

		conn.Selected = conn.ReferenceID == referenceID
		_, err := db.UpdateUserConnection(ctx, conn)
		if err != nil {
			if database.IsNotFound(err) {
				return models.UserConnection{}, ErrConnectionNotFound
			}
			return models.UserConnection{}, err
		}
	}

	if connection.ID == "" {
		return models.UserConnection{}, ErrConnectionNotFound
	}

	// as a side effect, update the user content in the background
	go func(db database.Client, u models.User, c models.UserConnection, r string) {
		content, ok := u.Content(models.UserContentTypePersonalBackground)
		if !ok {
			return
		}

		content.ReferenceID = userID
		if c.Verified == true {
			content.ReferenceID = r

			// check if there is an existing record for the same account and update it to not reference the account anymore
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			existing, err := db.GetUserContentFromRef(ctx, r, models.UserContentTypePersonalBackground)
			if err != nil && !database.IsNotFound(err) {
				log.Err(err).Msg("failed to get an existing content for a reference id")
			}

			if existing.ID != "" && existing.UserID != content.UserID {
				// if the a background with a reference to this account exists, update it to reference the user instead
				// - the current user has their account verified, so we assume that they are the current owner and should take over for showing the image on account
				existing.ReferenceID = existing.UserID // set the reference id to the userID

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				_, err = db.UpdateUserContent(ctx, existing)
				log.Err(err).Msg("failed to update existing content for a user")

			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err = db.UpdateUserContent(ctx, content)
		if err != nil {
			log.Err(err).Msg("failed to update user content as a side effect of default connection update")
		}
	}(db, user, connection, referenceID)

	return connection, nil
}
