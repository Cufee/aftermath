package fetch

import "github.com/pkg/errors"

var (
	AccountNotFound        = errors.New("no results found")
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
	}
}
