package fetch

import "github.com/pkg/errors"

var (
	ErrAccountNotFound     = errors.New("no results found")
	ErrInvalidSearch       = errors.New("invalid search")
	ErrSessionNotFound     = errors.New("account sessions not found")
	ErrInvalidSessionStart = errors.New("invalid session start provided")
	ErrSourceNotAvailable  = errors.New("source not available")
)

func parseWargamingError(err error) error {
	if err == nil {
		return nil
	}
	switch err.Error() {
	default:
		return err

	case "SOURCE_NOT_AVAILABLE":
		return ErrSourceNotAvailable
	case "INVALID_SEARCH":
		return ErrInvalidSearch
	}
}
