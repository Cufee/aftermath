package database

import (
	"context"
	"database/sql"
	"fmt"
	"runtime/debug"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath-assets/types"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/frame"

	_ "github.com/mattn/go-sqlite3"
)

var _ Client = &client{}

type AuthClient interface {
	CreateAuthNonce(ctx context.Context, publicID, identifier string, expiresAt time.Time, meta map[string]string) (models.AuthNonce, error)
	FindAuthNonce(ctx context.Context, publicID string) (models.AuthNonce, error)
	SetAuthNonceActive(ctx context.Context, nonceID string, active bool) error

	CreateSession(ctx context.Context, publicID, userID string, expiresAt time.Time, meta map[string]string) (models.Session, error)
	SetSessionExpiresAt(ctx context.Context, sessionID string, expiresAt time.Time) error
	FindSession(ctx context.Context, publicID string) (models.Session, error)
}

type AccountsClient interface {
	GetAccounts(ctx context.Context, ids []string) ([]models.Account, error)
	GetAccountByID(ctx context.Context, id string) (models.Account, error)
	GetRealmAccountIDs(ctx context.Context, realm string) ([]string, error)
	AccountSetPrivate(ctx context.Context, id string, value bool) error
	UpsertAccounts(ctx context.Context, accounts []models.Account) (map[string]error, error)
}

type GlossaryClient interface {
	GetAllVehicles(ctx context.Context) (map[string]models.Vehicle, error)
	GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error)
	GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error)

	UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) (map[string]error, error)
	UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) (map[string]error, error)

	GetMap(ctx context.Context, id string) (types.Map, error)
	UpsertMaps(ctx context.Context, maps map[string]types.Map) error
}

type UsersClient interface {
	GetUserByID(ctx context.Context, id string, opts ...UserGetOption) (models.User, error)
	GetOrCreateUserByID(ctx context.Context, id string, opts ...UserGetOption) (models.User, error)
	UpsertUserWithPermissions(ctx context.Context, userID string, perms permissions.Permissions) (models.User, error)

	GetConnection(ctx context.Context, connection string) (models.UserConnection, error)
	UpdateConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error)
	UpsertConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error)
	DeleteUserConnection(ctx context.Context, userID, connectionID string) error

	GetWidgetSettings(ctx context.Context, settingsID string) (models.WidgetOptions, error)
	GetUserWidgetSettings(ctx context.Context, userID string, referenceID []string) ([]models.WidgetOptions, error)
	UpdateWidgetSettings(ctx context.Context, id string, settings models.WidgetOptions) (models.WidgetOptions, error)
	CreateWidgetSettings(ctx context.Context, userID string, settings models.WidgetOptions) (models.WidgetOptions, error)
}

type SnapshotsClient interface {
	GetAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...Query) ([]models.AccountSnapshot, error)
	GetVehicleSnapshots(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...Query) ([]models.VehicleSnapshot, error)

	GetAccountLastBattleTimes(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...Query) (map[string]time.Time, error)
	GetVehicleLastBattleTimes(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...Query) (map[string]time.Time, error)

	CreateAccountSnapshots(ctx context.Context, snapshots ...models.AccountSnapshot) error
	CreateAccountVehicleSnapshots(ctx context.Context, accountID string, snapshots ...models.VehicleSnapshot) error

	DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error
}

type LeaderboardsClient interface {
	CreateLeaderboardScores(ctx context.Context, scores ...models.LeaderboardScore) error
	GetLeaderboardScores(ctx context.Context, leaderboardID string, scoreType models.ScoreType, options ...Query) ([]models.LeaderboardScore, error)

	DeleteExpiredLeaderboardScores(ctx context.Context, expiration time.Time, scoreType models.ScoreType) error
}

type TasksClient interface {
	CreateTasks(ctx context.Context, tasks ...models.Task) error
	GetTasks(ctx context.Context, ids ...string) ([]models.Task, error)
	UpdateTasks(ctx context.Context, tasks ...models.Task) error
	DeleteTasks(ctx context.Context, ids ...string) error
	AbandonTasks(ctx context.Context, ids ...string) error

	GetStaleTasks(ctx context.Context, limit int) ([]models.Task, error)
	GetRecentTasks(ctx context.Context, createdAfter time.Time, status ...models.TaskStatus) ([]models.Task, error)
	GetAndStartTasks(ctx context.Context, limit int) ([]models.Task, error)

	DeleteExpiredTasks(ctx context.Context, expiration time.Time) error
}

type DiscordDataClient interface {
	UpsertCommands(ctx context.Context, commands ...models.ApplicationCommand) error
	GetCommandsByID(ctx context.Context, commandIDs ...string) ([]models.ApplicationCommand, error)

	CreateDiscordInteraction(ctx context.Context, data models.DiscordInteraction) error
	GetDiscordInteraction(ctx context.Context, referenceID string) (models.DiscordInteraction, error)
	DeleteExpiredInteractions(ctx context.Context, expiration time.Time) error
}

type Client interface {
	AuthClient
	UsersClient

	GlossaryClient
	AccountsClient
	SnapshotsClient
	LeaderboardsClient

	TasksClient

	DiscordDataClient

	Disconnect() error
}

type client struct {
	db *db.Client
}

func (c *client) withTx(ctx context.Context, fn func(tx *db.Tx) error) error {
	var err error
	tx, err := c.db.Tx(ctx)
	if err != nil {
		return err
	}

	if err = fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return err
}

func (c *client) Disconnect() error {
	return c.db.Close()
}

type clientOptions struct {
	debug bool
}

type ClientOption func(*clientOptions)

func WithDebug() func(*clientOptions) {
	return func(opts *clientOptions) {
		opts.debug = true
	}
}

func NewSQLiteClient(filePath string, options ...ClientOption) (*client, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal().Interface("error", r).Str("stack", string(debug.Stack())).Msg("NewSQLiteClient panic")
		}
	}()

	opts := clientOptions{}
	for _, apply := range options {
		apply(&opts)
	}

	var dbOptions []db.Option
	if opts.debug {
		dbOptions = append(dbOptions, db.Debug())
	}

	sqldb, err := sql.Open("sqlite3", fmt.Sprintf("file://%s?_fk=1&_auto_vacuum=2&_synchronous=1&_journal_mode=WAL", filePath)) // _mutex
	if err != nil {
		return nil, err
	}
	sqldb.SetMaxOpenConns(1)
	dbOptions = append(dbOptions, db.Driver(entsql.OpenDB("sqlite3", sqldb)))

	return &client{
		db: db.NewClient(dbOptions...),
	}, nil
}

func toAnySlice[T any](s ...T) []any {
	var a []any
	for _, i := range s {
		a = append(a, i)
	}
	return a
}

func wrap(query string) *entsql.Builder {
	wrapper := &entsql.Builder{}
	return wrapper.Wrap(func(b *entsql.Builder) { b.WriteString(query) })
}

func batch[T any](ops []T, size int) [][]T {
	var batched [][]T
	for i := 0; i < len(ops); i += size {
		end := i + size
		if end > len(ops) {
			end = len(ops)
		}
		batched = append(batched, ops[i:end])
	}

	return batched
}
