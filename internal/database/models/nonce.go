package models

import (
	"time"

	"github.com/pkg/errors"
)

var ErrInvalidNonce = errors.New("invalid nonce")
var ErrNonceExpired = errors.New("nonce expired")

type AuthNonce struct {
	ID string

	Active     bool
	PublicID   string
	Identifier string

	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time

	Meta map[string]string
}

func (n AuthNonce) Valid() error {
	if n.ExpiresAt.Before(time.Now()) {
		return ErrNonceExpired
	}
	if n.Active && n.PublicID != "" && n.Identifier != "" {
		return nil
	}
	return ErrInvalidNonce
}
