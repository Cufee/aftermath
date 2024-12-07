package database

import (
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/qrm"
)

var ErrNotFound = qrm.ErrNoRows

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, sql.ErrNoRows) || errors.Is(err, qrm.ErrNoRows) || errors.Is(err, ErrNotFound)
}
