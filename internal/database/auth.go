package database

// import (
// 	"context"
// 	"time"

// 	"github.com/cufee/aftermath/internal/database/ent/db"
// 	"github.com/cufee/aftermath/internal/database/ent/db/authnonce"
// 	"github.com/cufee/aftermath/internal/database/ent/db/session"
// 	"github.com/cufee/aftermath/internal/database/models"
// )

// func toAuthNonce(record *db.AuthNonce) models.AuthNonce {
// 	return models.AuthNonce{
// 		ID:         record.ID,
// 		Active:     record.Active,
// 		PublicID:   record.PublicID,
// 		Identifier: record.Identifier,
// 		Meta:       record.Metadata,

// 		CreatedAt: record.CreatedAt,
// 		UpdatedAt: record.UpdatedAt,
// 		ExpiresAt: record.ExpiresAt,
// 	}
// }

// func (c *client) CreateAuthNonce(ctx context.Context, publicID, identifier string, expiresAt time.Time, meta map[string]string) (models.AuthNonce, error) {
// 	record, err := c.db.AuthNonce.Create().SetActive(true).SetExpiresAt(expiresAt).SetIdentifier(identifier).SetPublicID(publicID).SetMetadata(meta).Save(ctx)
// 	if err != nil {
// 		return models.AuthNonce{}, err
// 	}

// 	nonce := toAuthNonce(record)
// 	return nonce, nonce.Valid()
// }

// func (c *client) FindAuthNonce(ctx context.Context, publicID string) (models.AuthNonce, error) {
// 	record, err := c.db.AuthNonce.Query().Where(authnonce.PublicID(publicID), authnonce.Active(true), authnonce.ExpiresAtGT(time.Now())).First(ctx)
// 	if err != nil {
// 		return models.AuthNonce{}, err
// 	}

// 	nonce := toAuthNonce(record)
// 	return nonce, nonce.Valid()
// }

// func (c *client) SetAuthNonceActive(ctx context.Context, nonceID string, active bool) error {
// 	err := c.db.AuthNonce.UpdateOneID(nonceID).SetActive(active).Exec(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func toSession(record *db.Session) models.Session {
// 	return models.Session{
// 		ID:       record.ID,
// 		UserID:   record.UserID,
// 		PublicID: record.PublicID,
// 		Meta:     record.Metadata,

// 		CreatedAt: record.CreatedAt,
// 		UpdatedAt: record.UpdatedAt,
// 		ExpiresAt: record.ExpiresAt,
// 	}
// }

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
