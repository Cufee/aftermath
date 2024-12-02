package database

import (
	"context"
	"time"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/json"
	s "github.com/go-jet/jet/v2/sqlite"
)

func (c *client) CreateAuthNonce(ctx context.Context, publicID, identifier string, expiresAt time.Time, meta map[string]string) (models.AuthNonce, error) {
	insert := m.AuthNonce{
		ID:         c.newID(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Active:     true,
		ExpiresAt:  expiresAt,
		Identifier: identifier,
		PublicID:   publicID,
	}
	if meta != nil {
		metaEncoded, err := json.Marshal(meta)
		if err != nil {
			return models.AuthNonce{}, err
		}
		insert.Metadata = string(metaEncoded)
	}

	var result m.AuthNonce
	stmt := t.AuthNonce.
		INSERT(t.AuthNonce.AllColumns).
		MODEL(insert).
		RETURNING(t.AuthNonce.AllColumns)
	err := c.query(ctx, stmt, &result)
	if err != nil {
		return models.AuthNonce{}, err
	}

	nonce := models.ToAuthNonce(&result)
	return nonce, nonce.Valid()
}

func (c *client) FindAuthNonce(ctx context.Context, publicID string) (models.AuthNonce, error) {
	var record m.AuthNonce
	stmt := t.AuthNonce.
		SELECT(t.AuthNonce.AllColumns).
		WHERE(t.AuthNonce.PublicID.EQ(s.String(publicID)))
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.AuthNonce{}, err
	}

	nonce := models.ToAuthNonce(&record)
	return nonce, nonce.Valid()
}

func (c *client) SetAuthNonceActive(ctx context.Context, nonceID string, active bool) error {
	stmt := t.AuthNonce.
		UPDATE().
		SET(t.AuthNonce.Active.SET(s.Bool(active))).
		WHERE(t.AuthNonce.ID.EQ(s.String(nonceID)))
	_, err := c.exec(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}

// func (c *client) CreateSession(ctx context.Context, publicID, userID string, expiresAt time.Time, meta map[string]string) (models.Session, error) {
// 	user, err := c.GetOrCreateUserByID(ctx, userID)
// 	if err != nil {
// 		return models.Session{}, err
// 	}

// 	record, err := c.db.Session.Create().SetPublicID(publicID).SetUser(c.db.User.GetX(ctx, user.ID)).SetExpiresAt(expiresAt).SetMetadata(meta).Save(ctx)
// 	if err != nil {
// 		return models.Session{}, err
// 	}

// 	nonce := toSession(record)
// 	return nonce, nonce.Valid()
// }

// func (c *client) FindSession(ctx context.Context, publicID string) (models.Session, error) {
// 	record, err := c.db.Session.Query().Where(session.PublicID(publicID), session.ExpiresAtGT(time.Now())).First(ctx)
// 	if err != nil {
// 		return models.Session{}, err
// 	}

// 	session := toSession(record)
// 	return session, session.Valid()
// }

// func (c *client) SetSessionExpiresAt(ctx context.Context, sessionID string, expiresAt time.Time) error {
// 	err := c.db.Session.UpdateOneID(sessionID).SetExpiresAt(expiresAt).Exec(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
