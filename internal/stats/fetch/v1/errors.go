package fetch

import "github.com/pkg/errors"

var (
	AccountNotFound        = errors.New("no results found")
	ErrSessionNotFound     = errors.New("account sessions not found")
	ErrInvalidSessionStart = errors.New("invalid session start provided")
)
