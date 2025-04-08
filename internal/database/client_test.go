package database

import (
	"fmt"
	"os"
	"testing"

	"github.com/cufee/aftermath/tests/env"
)

func MustTestClient(t *testing.T) *client {
	env.LoadTestEnv(t)

	connString := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_NAME"))
	client, err := NewPostgresClient(connString)
	if err != nil {
		panic(err)
	}
	return client
}
