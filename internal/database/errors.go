package database

import "github.com/cufee/aftermath/internal/database/ent/db"

func IsNotFound(err error) bool {
	return db.IsNotFound(err)
}
