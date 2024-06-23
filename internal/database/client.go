package database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/frame"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var _ Client = &libsqlClient{}

type AccountsClient interface {
	GetAccounts(ctx context.Context, ids []string) ([]models.Account, error)
	GetAccountByID(ctx context.Context, id string) (models.Account, error)
	GetRealmAccountIDs(ctx context.Context, realm string) ([]string, error)
	AccountSetPrivate(ctx context.Context, id string, value bool) error
	UpsertAccounts(ctx context.Context, accounts []models.Account) (map[string]error, error)
}

type GlossaryClient interface {
	GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error)
	GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error)

	UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) (map[string]error, error)
	UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) (map[string]error, error)
}

type UsersClient interface {
	GetUserByID(ctx context.Context, id string, opts ...userGetOption) (models.User, error)
	GetOrCreateUserByID(ctx context.Context, id string, opts ...userGetOption) (models.User, error)
	UpsertUserWithPermissions(ctx context.Context, userID string, perms permissions.Permissions) (models.User, error)

	UpdateConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error)
	UpsertConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error)
}

type SnapshotsClient interface {
	GetAccountSnapshot(ctx context.Context, accountID, referenceID string, kind models.SnapshotType, options ...SnapshotQuery) (models.AccountSnapshot, error)
	CreateAccountSnapshots(ctx context.Context, snapshots ...models.AccountSnapshot) error
	GetLastAccountSnapshots(ctx context.Context, accountID string, limit int) ([]models.AccountSnapshot, error)
	GetManyAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...SnapshotQuery) ([]models.AccountSnapshot, error)

	GetVehicleSnapshots(ctx context.Context, accountID, referenceID string, kind models.SnapshotType, options ...SnapshotQuery) ([]models.VehicleSnapshot, error)
	CreateVehicleSnapshots(ctx context.Context, snapshots ...models.VehicleSnapshot) error

	DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error
}

type TasksClient interface {
	CreateTasks(ctx context.Context, tasks ...models.Task) error
	UpdateTasks(ctx context.Context, tasks ...models.Task) error
	DeleteTasks(ctx context.Context, ids ...string) error

	GetStaleTasks(ctx context.Context, limit int) ([]models.Task, error)
	GetRecentTasks(ctx context.Context, createdAfter time.Time, status ...models.TaskStatus) ([]models.Task, error)
	GetAndStartTasks(ctx context.Context, limit int) ([]models.Task, error)

	DeleteExpiredTasks(ctx context.Context, expiration time.Time) error
}

type DiscordDataClient interface {
	UpsertCommands(ctx context.Context, commands ...models.ApplicationCommand) error
	GetCommandsByID(ctx context.Context, commandIDs ...string) ([]models.ApplicationCommand, error)
}

type Client interface {
	UsersClient

	GlossaryClient
	AccountsClient
	SnapshotsClient

	TasksClient

	DiscordDataClient

	Disconnect() error
}

type libsqlClient struct {
	db              *db.Client
	transactionLock *sync.Mutex
}

func (c *libsqlClient) Disconnect() error {
	return c.db.Close()
}

func (c *libsqlClient) txWithLock(ctx context.Context) (*db.Tx, func(), error) {
	c.transactionLock.Lock()
	tx, err := c.db.Tx(ctx)
	if err != nil {
		c.transactionLock.Unlock()
		return tx, func() {}, nil
	}
	return tx, c.transactionLock.Unlock, nil
}

func NewLibSQLClient(primaryUrl string) (*libsqlClient, error) {
	driver, err := sql.Open("libsql", primaryUrl)
	if err != nil {
		return nil, err
	}

	dbClient := db.NewClient(db.Driver(entsql.OpenDB(dialect.SQLite, driver)))
	return &libsqlClient{
		transactionLock: &sync.Mutex{},
		db:              dbClient,
	}, nil
}

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *db.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
