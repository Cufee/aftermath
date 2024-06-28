package tests

import (
	"context"
	"errors"
	"time"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/frame"
)

var ErrNotFound = &db.NotFoundError{}

var _ database.Client = &staticTestingDatabase{}

type staticTestingDatabase struct{}

func StaticTestingDatabase() *staticTestingDatabase {
	return &staticTestingDatabase{}
}

func (c *staticTestingDatabase) Disconnect() error {
	return nil
}

func (c *staticTestingDatabase) GetAccounts(ctx context.Context, ids []string) ([]models.Account, error) {
	var accounts []models.Account
	for _, id := range ids {
		if a, ok := staticAccounts[id]; ok {
			accounts = append(accounts, a)
		}
	}
	return accounts, nil
}
func (c *staticTestingDatabase) GetAccountByID(ctx context.Context, id string) (models.Account, error) {
	if account, ok := staticAccounts[id]; ok {
		return account, nil
	}
	return models.Account{}, ErrNotFound
}
func (c *staticTestingDatabase) GetRealmAccountIDs(ctx context.Context, realm string) ([]string, error) {
	return nil, errors.New("GetRealmAccountIDs not implemented")
}
func (c *staticTestingDatabase) AccountSetPrivate(ctx context.Context, id string, value bool) error {
	return errors.New("AccountSetPrivate not implemented")
}
func (c *staticTestingDatabase) UpsertAccounts(ctx context.Context, accounts []models.Account) (map[string]error, error) {
	for _, acc := range accounts {
		if account, ok := staticAccounts[acc.ID]; ok {
			staticAccounts[acc.ID] = account
		}
	}
	return nil, nil
}

func (c *staticTestingDatabase) GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error) {
	vehicles := make(map[string]models.Vehicle)
	for _, id := range ids {
		vehicles[id] = models.Vehicle{ID: id, Tier: 10, LocalizedNames: map[string]string{"en": "Test Vehicle " + id}}
	}
	return vehicles, nil
}
func (c *staticTestingDatabase) GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error) {
	// TODO: get some kind of data in
	return map[string]frame.StatsFrame{}, nil
}
func (c *staticTestingDatabase) UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) (map[string]error, error) {
	return nil, errors.New("UpsertVehicles not implemented")
}
func (c *staticTestingDatabase) UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) (map[string]error, error) {
	return nil, errors.New("UpsertVehicleAverages not implemented")
}

func (c *staticTestingDatabase) GetUserByID(ctx context.Context, id string, opts ...database.UserGetOption) (models.User, error) {
	return models.User{}, errors.New("GetUserByID not implemented")
}
func (c *staticTestingDatabase) GetOrCreateUserByID(ctx context.Context, id string, opts ...database.UserGetOption) (models.User, error) {
	return models.User{}, errors.New("GetOrCreateUserByID not implemented")
}
func (c *staticTestingDatabase) UpsertUserWithPermissions(ctx context.Context, userID string, perms permissions.Permissions) (models.User, error) {
	return models.User{}, errors.New("UpsertUserWithPermissions not implemented")
}
func (c *staticTestingDatabase) UpdateConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	return models.UserConnection{}, errors.New("UpdateConnection not implemented")
}
func (c *staticTestingDatabase) UpsertConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	return models.UserConnection{}, errors.New("UpsertConnection not implemented")
}

func (c *staticTestingDatabase) GetAccountSnapshot(ctx context.Context, accountID, referenceID string, kind models.SnapshotType, options ...database.SnapshotQuery) (models.AccountSnapshot, error) {
	return models.AccountSnapshot{}, errors.New("GetAccountSnapshot not implemented")
}
func (c *staticTestingDatabase) CreateAccountSnapshots(ctx context.Context, snapshots ...models.AccountSnapshot) (map[string]error, error) {
	return nil, errors.New("CreateAccountSnapshots not implemented")
}
func (c *staticTestingDatabase) GetLastAccountSnapshots(ctx context.Context, accountID string, limit int) ([]models.AccountSnapshot, error) {
	return nil, errors.New("GetLastAccountSnapshots not implemented")
}
func (c *staticTestingDatabase) GetManyAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...database.SnapshotQuery) ([]models.AccountSnapshot, error) {
	return nil, errors.New("GetManyAccountSnapshots not implemented")
}
func (c *staticTestingDatabase) GetVehicleSnapshots(ctx context.Context, accountID, referenceID string, kind models.SnapshotType, options ...database.SnapshotQuery) ([]models.VehicleSnapshot, error) {
	return nil, errors.New("GetVehicleSnapshots not implemented")
}
func (c *staticTestingDatabase) CreateAccountVehicleSnapshots(ctx context.Context, accountID string, snapshots ...models.VehicleSnapshot) (map[string]error, error) {
	return nil, errors.New("CreateAccountVehicleSnapshots not implemented")
}
func (c *staticTestingDatabase) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
	return errors.New("DeleteExpiredSnapshots not implemented")
}

func (c *staticTestingDatabase) CreateTasks(ctx context.Context, tasks ...models.Task) error {
	return errors.New("not implemented")
}
func (c *staticTestingDatabase) GetTasks(ctx context.Context, ids ...string) ([]models.Task, error) {
	return nil, errors.New("CreateTasks not implemented")
}
func (c *staticTestingDatabase) UpdateTasks(ctx context.Context, tasks ...models.Task) error {
	return errors.New("UpdateTasks not implemented")
}
func (c *staticTestingDatabase) DeleteTasks(ctx context.Context, ids ...string) error {
	return errors.New("DeleteTasks not implemented")
}
func (c *staticTestingDatabase) AbandonTasks(ctx context.Context, ids ...string) error {
	return errors.New("AbandonTasks not implemented")
}

func (c *staticTestingDatabase) GetStaleTasks(ctx context.Context, limit int) ([]models.Task, error) {
	return nil, errors.New("GetStaleTasks not implemented")
}
func (c *staticTestingDatabase) GetRecentTasks(ctx context.Context, createdAfter time.Time, status ...models.TaskStatus) ([]models.Task, error) {
	return nil, errors.New("GetRecentTasks not implemented")
}
func (c *staticTestingDatabase) GetAndStartTasks(ctx context.Context, limit int) ([]models.Task, error) {
	return nil, errors.New("GetAndStartTasks not implemented")
}
func (c *staticTestingDatabase) DeleteExpiredTasks(ctx context.Context, expiration time.Time) error {
	return errors.New("DeleteExpiredTasks not implemented")
}

func (c *staticTestingDatabase) UpsertCommands(ctx context.Context, commands ...models.ApplicationCommand) error {
	return errors.New("UpsertCommands not implemented")
}
func (c *staticTestingDatabase) GetCommandsByID(ctx context.Context, commandIDs ...string) ([]models.ApplicationCommand, error) {
	return nil, errors.New("GetCommandsByID not implemented")
}

func (c *staticTestingDatabase) CreateDiscordInteraction(ctx context.Context, data models.DiscordInteraction) error {
	return errors.New("CreateDiscordInteraction not implemented")
}
func (c *staticTestingDatabase) GetDiscordInteraction(ctx context.Context, referenceID string) (models.DiscordInteraction, error) {
	return models.DiscordInteraction{}, errors.New("GetDiscordInteraction not implemented")
}
func (c *staticTestingDatabase) DeleteExpiredInteractions(ctx context.Context, expiration time.Time) error {
	return errors.New("DeleteExpiredInteractions not implemented")
}
