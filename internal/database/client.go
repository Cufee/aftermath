package database

import (
	"context"
	"database/sql"
	"fmt"
	"runtime/debug"
	"time"

	"github.com/cufee/aftermath-assets/types"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/log"
	"github.com/cufee/aftermath/internal/stats/frame"
	pg "github.com/go-jet/jet/v2/postgres"
	"golang.org/x/text/language"

	_ "github.com/lib/pq"
)

//go:generate go run ./generate.go

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
	UpsertAccounts(ctx context.Context, accounts ...*models.Account) (map[string]error, error)
}

type GlossaryClient interface {
	GetAllVehicles(ctx context.Context) (map[string]models.Vehicle, error)
	GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error)
	GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error)

	UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) (map[string]error, error)
	UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) (map[string]error, error)

	GetMap(ctx context.Context, id string) (types.Map, error)
	UpsertMaps(ctx context.Context, maps map[string]types.Map) error

	GetGameModeNames(ctx context.Context, id string) (map[language.Tag]string, error)
	UpsertGameModes(ctx context.Context, modes map[string]map[language.Tag]string) (map[string]error, error)
}

type UsersClient interface {
	GetUserByID(ctx context.Context, id string, opts ...UserQueryOption) (models.User, error)
	GetOrCreateUserByID(ctx context.Context, id string, opts ...UserQueryOption) (models.User, error)

	FindUserConnections(ctx context.Context, opts ...ConnectionQueryOption) ([]models.UserConnection, error)
	GetUserConnection(ctx context.Context, connection string) (models.UserConnection, error)
	UpdateUserConnection(ctx context.Context, id string, connection models.UserConnection) (models.UserConnection, error)
	UpsertUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error)
	DeleteUserConnection(ctx context.Context, userID, connectionID string) error
	SetReferenceConnectionsUnverified(ctx context.Context, referenceID string) error

	GetWidgetSettings(ctx context.Context, settingsID string) (models.WidgetOptions, error)
	GetUserWidgetSettings(ctx context.Context, userID string, referenceID []string) ([]models.WidgetOptions, error)
	UpdateWidgetSettings(ctx context.Context, id string, settings models.WidgetOptions) (models.WidgetOptions, error)
	CreateWidgetSettings(ctx context.Context, userID string, settings models.WidgetOptions) (models.WidgetOptions, error)

	GetUserContent(ctx context.Context, id string) (models.UserContent, error)
	GetUserContentFromRef(ctx context.Context, referenceID string, kind models.UserContentType) (models.UserContent, error)
	CreateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error)
	UpdateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error)
	UpsertUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error)
	DeleteUserContent(ctx context.Context, id string) error

	CreateUserSubscription(ctx context.Context, subscription models.UserSubscription) (models.UserSubscription, error)
	UpdateUserSubscription(ctx context.Context, id string, subscription models.UserSubscription) (models.UserSubscription, error)
}

type SnapshotsClient interface {
	CreateVehicleSnapshots(ctx context.Context, snapshots ...*models.VehicleSnapshot) error
	GetAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...Query) ([]models.AccountSnapshot, error)
	GetVehicleSnapshots(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...Query) ([]models.VehicleSnapshot, error)

	CreateAccountSnapshots(ctx context.Context, snapshots ...*models.AccountSnapshot) error
	GetAccountLastBattleTimes(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...Query) (map[string]time.Time, error)
	GetVehiclesLastBattleTimes(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...Query) (map[string]time.Time, error)

	DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) (int, error)
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

	CreateDiscordInteraction(ctx context.Context, data models.DiscordInteraction) (models.DiscordInteraction, error)
	GetDiscordInteraction(ctx context.Context, id string) (models.DiscordInteraction, error)
	FindDiscordInteractions(ctx context.Context, opts ...InteractionQuery) ([]models.DiscordInteraction, error)
	DeleteExpiredInteractions(ctx context.Context, expiration time.Time) error
}

type ModerationClient interface {
	GetModerationRequest(ctx context.Context, id string) (models.ModerationRequest, error)
	FindUserModerationRequests(ctx context.Context, userID string, referenceIDs []string, status []models.ModerationStatus, since time.Time) ([]models.ModerationRequest, error)
	CreateModerationRequest(ctx context.Context, request models.ModerationRequest) (models.ModerationRequest, error)
	UpdateModerationRequest(ctx context.Context, request models.ModerationRequest) (models.ModerationRequest, error)

	GetUserRestriction(ctx context.Context, id string) (models.UserRestriction, error)
	GetUserRestrictions(ctx context.Context, userID string) ([]models.UserRestriction, error)
	CreateUserRestriction(ctx context.Context, data models.UserRestriction) (models.UserRestriction, error)
	UpdateUserRestriction(ctx context.Context, data models.UserRestriction) (models.UserRestriction, error)
}

type Client interface {
	AuthClient
	UsersClient

	GlossaryClient
	AccountsClient
	SnapshotsClient

	TasksClient

	DiscordDataClient

	ModerationClient
	Disconnect() error

	Unsafe() *sql.DB
}

type ClientOption func(*clientOptions)

func WithDebug() func(*clientOptions) {
	return func(opts *clientOptions) {
		opts.debug = true
	}
}

func NewPostgresClient(connString string, options ...ClientOption) (*client, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal().Interface("error", r).Str("stack", string(debug.Stack())).Msg("NewClient panic")
		}
	}()

	opts := clientOptions{}
	for _, apply := range options {
		apply(&opts)
	}

	sqldb, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	sqldb.SetMaxOpenConns(25)
	sqldb.SetMaxIdleConns(10)

	if err := sqldb.Ping(); err != nil {
		return nil, err
	}

	return &client{
		options: opts,
		db:      sqldb,
	}, nil
}

type client struct {
	options clientOptions
	db      *sql.DB
}

func (c *client) Unsafe() *sql.DB {
	return c.db
}

func (c *client) query(ctx context.Context, stmt pg.Statement, dst interface{}) error {
	if c.options.debug {
		println("SQL Query:", stmt.DebugSql())
	}
	return stmt.QueryContext(ctx, c.db, dst)
}

func (c *client) rows(ctx context.Context, stmt pg.Statement) (*pg.Rows, error) {
	if c.options.debug {
		println("SQL Rows:", stmt.DebugSql())
	}
	return stmt.Rows(ctx, c.db)
}

func (c *client) exec(ctx context.Context, stmt pg.Statement) (sql.Result, error) {
	if c.options.debug {
		println("SQL Exec:", stmt.DebugSql())
	}
	return stmt.ExecContext(ctx, c.db)
}

type transaction struct {
	tx      *sql.Tx
	options clientOptions
}

func (t *transaction) query(ctx context.Context, stmt pg.Statement, dst interface{}) error {
	if t.options.debug {
		println("SQL Query:", stmt.DebugSql())
	}
	return stmt.QueryContext(ctx, t.tx, dst)
}

func (t *transaction) exec(ctx context.Context, stmt pg.Statement) (sql.Result, error) {
	if t.options.debug {
		println("SQL Exec:", stmt.DebugSql())
	}
	return stmt.ExecContext(ctx, t.tx)
}

func (c *client) withTx(ctx context.Context, fn func(tx *transaction) error) error {
	var err error
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err = fn(&transaction{tx: tx, options: c.options}); err != nil {
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

func stringsToExp(s []string) []pg.Expression {
	var a []pg.Expression
	for _, i := range s {
		a = append(a, pg.String(i))
	}
	return a
}

func timeToField(t time.Time) pg.StringExpression {
	return pg.String(models.TimeToString(t))
}
