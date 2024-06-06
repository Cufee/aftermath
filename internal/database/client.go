package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/stats/frame"
)

type Client interface {
	SetVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) error
	GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error)

	GetUserByID(ctx context.Context, id string, opts ...userGetOption) (User, error)
	GetOrCreateUserByID(ctx context.Context, id string, opts ...userGetOption) (User, error)
	UpdateConnection(ctx context.Context, connection UserConnection) (UserConnection, error)
	UpsertConnection(ctx context.Context, connection UserConnection) (UserConnection, error)
}

// var _ Client = &client{} // just a marker to see if it is implemented correctly

type client struct {
	prisma *db.PrismaClient
}

func NewClient() (*client, error) {
	prisma := db.NewClient()
	err := prisma.Connect()
	if err != nil {
		return nil, err
	}

	return &client{prisma: prisma}, nil
}
