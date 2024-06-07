package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/stats/frame"
	"golang.org/x/sync/semaphore"
)

type Client interface {
	GetAccountByID(ctx context.Context, id string) (Account, error)
	GetAccounts(ctx context.Context, ids []string) ([]Account, error)
	GetRealmAccounts(ctx context.Context, realm string) ([]Account, error)
	UpsertAccounts(ctx context.Context, accounts []Account) map[string]error

	GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error)
	UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) error
	GetVehicles(ctx context.Context, ids []string) (map[string]Vehicle, error)
	UpsertVehicles(ctx context.Context, vehicles map[string]Vehicle) error

	GetUserByID(ctx context.Context, id string, opts ...userGetOption) (User, error)
	GetOrCreateUserByID(ctx context.Context, id string, opts ...userGetOption) (User, error)
	UpdateConnection(ctx context.Context, connection UserConnection) (UserConnection, error)
	UpsertConnection(ctx context.Context, connection UserConnection) (UserConnection, error)

	CreateTasks(ctx context.Context, tasks ...Task) error
	UpdateTasks(ctx context.Context, tasks ...Task) error
	DeleteTasks(ctx context.Context, ids ...string) error
}

var _ Client = &client{}

type client struct {
	prisma *db.PrismaClient
	// Prisma does not currently support updateManyAndReturn
	// in order to avoid a case where we
	tasksUpdateSem *semaphore.Weighted
}

func (c *client) Prisma() *db.PrismaClient {
	return c.prisma
}

func NewClient() (*client, error) {
	prisma := db.NewClient()
	err := prisma.Connect()
	if err != nil {
		return nil, err
	}

	return &client{prisma: prisma, tasksUpdateSem: semaphore.NewWeighted(1)}, nil
}
