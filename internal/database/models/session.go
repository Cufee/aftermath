package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
	"github.com/pkg/errors"
)

var (
	ErrSessionExpired = errors.New("session expired")
	ErrSessionInvalid = errors.New("session invalid")
)

type Session struct {
	ID string

	UserID   string
	PublicID string

	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	Meta map[string]string
}

func (n Session) Valid() error {
	if n.ExpiresAt.Before(time.Now()) {
		return ErrSessionExpired
	}
	if n.UserID != "" && n.PublicID != "" {
		return nil
	}
	return ErrSessionInvalid
}

func ToSession(record *model.Session) Session {
	session := Session{
		ID:       record.ID,
		UserID:   record.UserID,
		PublicID: record.PublicID,

		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		ExpiresAt: record.ExpiresAt,
	}
	json.Unmarshal(record.Metadata, &session.Meta)

	if session.Meta == nil {
		session.Meta = make(map[string]string, 0)
	}
	return session
}

func (record *Session) Model() model.Session {
	session := model.Session{
		ID:       utils.StringOr(record.ID, cuid.New()),
		UserID:   record.UserID,
		PublicID: record.PublicID,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: record.ExpiresAt,
	}
	if record.Meta != nil {
		data, _ := json.Marshal(record.Meta)
		session.Metadata = data
	}
	return session
}
