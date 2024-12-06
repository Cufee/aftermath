package database

import (
	"context"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestUsers(t *testing.T) {
	client := MustTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	//
	_ = ctx
	_ = client

	t.Run("", func(t *testing.T) {
		is := is.New(t)
		_ = is

	})
}
