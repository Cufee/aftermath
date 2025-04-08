package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/model"
	"github.com/cufee/aftermath/internal/json"
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

func ToAuthNonce(record *model.AuthNonce) AuthNonce {
	nonce := AuthNonce{
		ID:         record.ID,
		Active:     record.Active,
		PublicID:   record.PublicID,
		Identifier: record.Identifier,

		CreatedAt: StringToTime(record.CreatedAt),
		UpdatedAt: StringToTime(record.UpdatedAt),
		ExpiresAt: StringToTime(record.ExpiresAt),
	}
	json.Unmarshal(record.Metadata, &nonce.Meta)

	if nonce.Meta == nil {
		nonce.Meta = make(map[string]string, 0)
	}
	return nonce
}
