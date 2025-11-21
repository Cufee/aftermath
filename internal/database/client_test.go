package database

import (
	"context"
	"os"
	"testing"

	"github.com/cufee/aftermath/tests/env"
)

func MustTestClient(t *testing.T) *client {
	env.LoadTestEnv(t)

	connString := os.Getenv("DATABASE_URL")
	client, err := NewPostgresClient(context.Background(), connString)
	if err != nil {
		panic(err)
	}
	return client
}
