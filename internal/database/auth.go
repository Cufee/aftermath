package database

import (
	"context"
	"time"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/json"
	s "github.com/go-jet/jet/v2/sqlite"
	"github.com/lucsky/cuid"
)

func (c *client) CreateAuthNonce(ctx context.Context, publicID, identifier string, expiresAt time.Time, meta map[string]string) (models.AuthNonce, error) {
	insert := m.AuthNonce{
		ID:         cuid.New(),
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
		insert.Metadata = metaEncoded
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

func (c *client) CreateSession(ctx context.Context, publicID, userID string, expiresAt time.Time, meta map[string]string) (models.Session, error) {
	model := m.Session{
		ID:        cuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: expiresAt,
		PublicID:  publicID,
		UserID:    userID,
	}
	if meta != nil {
		data, err := json.Marshal(meta)
		if err != nil {
			return models.Session{}, err
		}
		model.Metadata = data
	}

	stmt := t.Session.
		INSERT().
		MODEL(model).
		RETURNING()

	var record m.Session
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.Session{}, err
	}

	return models.ToSession(&record), nil
}

func (c *client) FindSession(ctx context.Context, publicID string) (models.Session, error) {
	stmt := t.Session.
		SELECT(t.Session.AllColumns).
		WHERE(t.Session.PublicID.EQ(s.String(publicID)))

	var record m.Session
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.Session{}, err
	}

	session := models.ToSession(&record)
	return session, session.Valid()
}

func (c *client) SetSessionExpiresAt(ctx context.Context, sessionID string, expiresAt time.Time) error {
	stmt := t.Session.
		UPDATE(t.Session.ExpiresAt).
		WHERE(t.Session.ID.EQ(s.String(sessionID))).
		SET(t.Session.ExpiresAt.SET(s.DATETIME(expiresAt)))

	_, err := c.exec(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}
