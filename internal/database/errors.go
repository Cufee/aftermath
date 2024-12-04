package database

import (
	"database/sql"
	"errors"
)

var ErrNotFound = sql.ErrNoRows

func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, sql.ErrNoRows) || errors.Is(err, ErrNotFound)
}
