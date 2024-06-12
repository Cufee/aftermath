package fetch

import "github.com/pkg/errors"

var (
	ErrSessionNotFound     = errors.New("account sessions not found")
	ErrInvalidSessionStart = errors.New("invalid session start provided")
)
