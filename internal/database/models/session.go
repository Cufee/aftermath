package models

import (
	"time"

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
