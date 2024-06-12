package fetch

import (
	"context"
	"testing"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/stretchr/testify/assert"
)

func TestSessionStats(t *testing.T) {
	db, err := database.NewClient()
	assert.NoError(t, err, "failed to create a database client")
	defer db.Prisma().Disconnect()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	sessionStartTime := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)

	client := &multiSourceClient{database: db}
	stats, err := client.SessionStats(ctx, "", sessionStartTime.Add(time.Hour*-1))
	assert.NoError(t, err, "failed to get session stats")

	_ = stats
}
