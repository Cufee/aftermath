package database

import (
	"context"
	"testing"
	"time"
)

func TestUsers(t *testing.T) {
	client := MustTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	//
	_ = ctx
	_ = client
}
