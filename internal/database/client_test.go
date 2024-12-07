package database

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cufee/aftermath/tests/env"
)

func MustTestClient(t *testing.T) *client {
	env.LoadTestEnv(t)
	client, err := NewSQLiteClient(filepath.Join(os.Getenv("DATABASE_PATH"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		panic(err)
	}
	return client
}
