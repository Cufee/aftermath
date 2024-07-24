package tests

import (
	"context"
	"errors"
	"fmt"

	"time"

	"github.com/cufee/aftermath-assets/types"
	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/stats/frame"
	"golang.org/x/text/language"
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
		} else {
			accounts = append(accounts, models.Account{
				ID:       id,
				Nickname: "some_account_" + id,
				Realm:    "NA",
			})
		}
	}
	return accounts, nil
}
func (c *staticTestingDatabase) GetAccountByID(ctx context.Context, id string) (models.Account, error) {
	if account, ok := staticAccounts[id]; ok {
		return account, nil
	}
	return models.Account{
		ID:       id,
		Realm:    "NA",
		Nickname: "some_account" + id,
	}, nil
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

func (c *staticTestingDatabase) GetAllVehicles(ctx context.Context) (map[string]models.Vehicle, error) {
	vehicles := make(map[string]models.Vehicle)
	for i := range 10 {
		id := fmt.Sprint(i)
		vehicles[id] = models.Vehicle{ID: id, Tier: 10, LocalizedNames: map[language.Tag]string{language.English: "Test Vehicle " + id}}
	}
	return vehicles, nil
}
func (c *staticTestingDatabase) GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error) {
	vehicles := make(map[string]models.Vehicle)
	for _, id := range ids {
		vehicles[id] = models.Vehicle{ID: id, Tier: 10, LocalizedNames: map[language.Tag]string{language.English: "Test Vehicle " + id}}
	}
	return vehicles, nil
}
func (c *staticTestingDatabase) GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error) {
	averages := make(map[string]frame.StatsFrame)
	for _, id := range ids {
		averages[id] = DefaultStatsFrameSmall2
	}
	return averages, nil
}
func (c *staticTestingDatabase) UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) (map[string]error, error) {
	return nil, errors.New("UpsertVehicles not implemented")
}
func (c *staticTestingDatabase) UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) (map[string]error, error) {
	return nil, errors.New("UpsertVehicleAverages not implemented")
}

func (c *staticTestingDatabase) GetUserByID(ctx context.Context, id string, opts ...database.UserGetOption) (models.User, error) {
	return DefaultUserWithEdges, nil
}
func (c *staticTestingDatabase) GetOrCreateUserByID(ctx context.Context, id string, opts ...database.UserGetOption) (models.User, error) {
	return c.GetUserByID(ctx, id)
}
func (c *staticTestingDatabase) UpsertUserWithPermissions(ctx context.Context, userID string, perms permissions.Permissions) (models.User, error) {
	u, err := c.GetUserByID(ctx, userID)
	if err != nil {
		return u, err
	}
	u.Permissions = perms
	return u, nil
}
func (c *staticTestingDatabase) GetUserConnection(ctx context.Context, id string) (models.UserConnection, error) {
	return models.UserConnection{}, errors.New("GetConnection not implemented")
}
func (c *staticTestingDatabase) UpdateUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	return connection, nil
}
func (c *staticTestingDatabase) UpsertUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	return connection, nil
}
func (c *staticTestingDatabase) DeleteUserConnection(ctx context.Context, userID, connectionID string) error {
	return nil
}

func (c *staticTestingDatabase) GetAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...database.Query) ([]models.AccountSnapshot, error) {
	return nil, errors.New("GetAccountSnapshots not implemented")
}
func (c *staticTestingDatabase) CreateAccountSnapshots(ctx context.Context, snapshots ...models.AccountSnapshot) error {
	return errors.New("CreateAccountSnapshots not implemented")
}
func (c *staticTestingDatabase) GetAccountLastBattleTimes(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...database.Query) (map[string]time.Time, error) {
	return nil, errors.New("GetAccountLastBattleTimes not implemented")
}
func (c *staticTestingDatabase) GetVehicleSnapshots(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...database.Query) ([]models.VehicleSnapshot, error) {
	return nil, errors.New("GetVehicleSnapshots not implemented")
}
func (c *staticTestingDatabase) CreateAccountVehicleSnapshots(ctx context.Context, accountID string, snapshots ...models.VehicleSnapshot) error {
	return errors.New("CreateAccountVehicleSnapshots not implemented")
}
func (c *staticTestingDatabase) GetVehicleLastBattleTimes(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...database.Query) (map[string]time.Time, error) {
	return nil, errors.New("GetVehicleLastBattleTimes not implemented")
}
func (c *staticTestingDatabase) DeleteExpiredSnapshots(ctx context.Context, expiration time.Time) error {
	return errors.New("DeleteExpiredSnapshots not implemented")
}
func (c *staticTestingDatabase) CreateAccountAchievementSnapshots(ctx context.Context, accountID string, snapshots ...models.AchievementsSnapshot) error {
	return errors.New("CreateAccountAchievementSnapshots not implemented")
}
func (c *staticTestingDatabase) GetAchievementSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...database.Query) ([]models.AchievementsSnapshot, error) {
	return nil, errors.New("GetAchievementsSnapshots not implemented")
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
func (c *staticTestingDatabase) FindDiscordInteractions(ctx context.Context, opts ...database.InteractionQuery) ([]models.DiscordInteraction, error) {
	return nil, errors.New("FindDiscordInteractions not implemented")
}
func (c *staticTestingDatabase) DeleteExpiredInteractions(ctx context.Context, expiration time.Time) error {
	return errors.New("DeleteExpiredInteractions not implemented")
}

func (c *staticTestingDatabase) CreateAuthNonce(ctx context.Context, publicID, identifier string, expiresAt time.Time, meta map[string]string) (models.AuthNonce, error) {
	return models.AuthNonce{ID: "nonce1", Active: true, PublicID: "nonce1", Identifier: "ident1", ExpiresAt: time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC)}, nil
}
func (c *staticTestingDatabase) FindAuthNonce(ctx context.Context, publicID string) (models.AuthNonce, error) {
	return models.AuthNonce{ID: "nonce1", Active: true, PublicID: "nonce1", Identifier: "ident1", ExpiresAt: time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC)}, nil
}
func (c *staticTestingDatabase) SetAuthNonceActive(ctx context.Context, nonceID string, active bool) error {
	return nil
}

func (c *staticTestingDatabase) CreateSession(ctx context.Context, publicID, userID string, expiresAt time.Time, meta map[string]string) (models.Session, error) {
	return models.Session{ID: "session1", UserID: "user1", PublicID: "cookie1", ExpiresAt: time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC)}, nil
}
func (c *staticTestingDatabase) SetSessionExpiresAt(ctx context.Context, sessionID string, expiresAt time.Time) error {
	return nil
}
func (c *staticTestingDatabase) FindSession(ctx context.Context, publicID string) (models.Session, error) {
	return models.Session{ID: "session1", UserID: "user1", PublicID: "cookie1", ExpiresAt: time.Date(9999, 0, 0, 0, 0, 0, 0, time.UTC)}, nil
}
func (c *staticTestingDatabase) UserFromSession(ctx context.Context, publicID string, opts ...database.UserGetOption) (models.User, error) {
	return DefaultUserWithEdges, nil
}

func (c *staticTestingDatabase) CreateLeaderboardScores(ctx context.Context, scores ...models.LeaderboardScore) error {
	return errors.New("CreateLeaderboardScores not implemented")
}
func (c *staticTestingDatabase) GetLeaderboardScores(ctx context.Context, leaderboard string, scoreType models.ScoreType, options ...database.Query) ([]models.LeaderboardScore, error) {
	return nil, errors.New("GetLeaderboardScores not implemented")
}
func (c *staticTestingDatabase) DeleteExpiredLeaderboardScores(ctx context.Context, expiration time.Time, kind models.ScoreType) error {
	return errors.New("DeleteExpiredLeaderboardScores not implemented")
}

func (c *staticTestingDatabase) GetMap(ctx context.Context, id string) (types.Map, error) {
	return types.Map{
		ID:             "1",
		LocalizedNames: map[language.Tag]string{language.English: "Mock Map Name"},
	}, nil
}
func (c *staticTestingDatabase) UpsertMaps(ctx context.Context, maps map[string]types.Map) error {
	return errors.New("UpsertMaps not implemented")
}

func (c *staticTestingDatabase) GetWidgetSettings(ctx context.Context, settingsID string) (models.WidgetOptions, error) {
	conn, _ := DefaultUserWithEdges.Connection(models.ConnectionTypeWargaming, nil)
	return models.WidgetOptions{
		ID:        "w1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    DefaultUserWithEdges.ID,
		AccountID: conn.ID,
		Style: models.WidgetStyling{
			UnratedOverview: models.WidgetCardStyle{
				Visible:   true,
				ShowTitle: true,
			},
		},
	}, nil
}
func (c *staticTestingDatabase) GetUserWidgetSettings(ctx context.Context, userID string, referenceID []string) ([]models.WidgetOptions, error) {
	s, err := c.GetWidgetSettings(ctx, userID)
	return []models.WidgetOptions{s}, err
}
func (c *staticTestingDatabase) UpdateWidgetSettings(ctx context.Context, id string, settings models.WidgetOptions) (models.WidgetOptions, error) {
	return settings, nil
}
func (c *staticTestingDatabase) CreateWidgetSettings(ctx context.Context, userID string, settings models.WidgetOptions) (models.WidgetOptions, error) {
	settings.ID = fmt.Sprint(time.Now().Unix())
	settings.UserID = userID
	return settings, nil
}

func (c *staticTestingDatabase) GetGameModeNames(ctx context.Context, id string) (map[language.Tag]string, error) {
	return nil, nil
}
func (c *staticTestingDatabase) UpsertGameModes(ctx context.Context, modes map[string]map[language.Tag]string) (map[string]error, error) {
	return nil, nil
}

func (c *staticTestingDatabase) FindUserModerationRequests(ctx context.Context, userID string, referenceIDs []string, status []models.ModerationStatus, since time.Time) ([]models.ModerationRequest, error) {
	return nil, errors.New("FindUserModerationRequests not implemented")
}
func (c *staticTestingDatabase) CreateModerationRequest(ctx context.Context, request models.ModerationRequest) (models.ModerationRequest, error) {
	return models.ModerationRequest{}, errors.New("CreateModerationRequest not implemented")
}
func (c *staticTestingDatabase) GetModerationRequest(ctx context.Context, id string) (models.ModerationRequest, error) {
	return models.ModerationRequest{}, errors.New("GetModerationRequest not implemented")
}
func (c *staticTestingDatabase) UpdateModerationRequest(ctx context.Context, request models.ModerationRequest) (models.ModerationRequest, error) {
	return models.ModerationRequest{}, errors.New("UpdateModerationRequest not implemented")
}

func (c *staticTestingDatabase) GetUserContent(ctx context.Context, id string) (models.UserContent, error) {
	return models.UserContent{}, errors.New("GetUserContent not implemented")
}
func (c *staticTestingDatabase) GetUserContentFromRef(ctx context.Context, referenceID string, kind models.UserContentType) (models.UserContent, error) {
	return models.UserContent{}, errors.New("GetUserContentFromRef not implemented")
}
func (c *staticTestingDatabase) FindUserContentFromRefs(ctx context.Context, kind models.UserContentType, referenceIDs ...string) ([]models.UserContent, error) {
	return nil, errors.New("FindUserContentFromRefs not implemented")
}
func (c *staticTestingDatabase) CreateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	return models.UserContent{}, errors.New("CreateUserContent not implemented")
}
func (c *staticTestingDatabase) UpdateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	return models.UserContent{}, errors.New("UpdateUserContent not implemented")
}
func (c *staticTestingDatabase) UpsertUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	return models.UserContent{}, errors.New("UpsertUserContent not implemented")
}
func (c *staticTestingDatabase) DeleteUserContent(ctx context.Context, id string) error {
	return errors.New("DeleteUserContent not implemented")
}
func (c *staticTestingDatabase) GetUserRestriction(ctx context.Context, id string) (models.UserRestriction, error) {
	return models.UserRestriction{}, errors.New("GetUserRestriction not implemented")
}
func (c *staticTestingDatabase) GetUserRestrictions(ctx context.Context, userID string) ([]models.UserRestriction, error) {
	return nil, errors.New("GetUserRestrictions not implemented")
}
func (c *staticTestingDatabase) CreateUserRestriction(ctx context.Context, data models.UserRestriction) (models.UserRestriction, error) {
	return models.UserRestriction{}, errors.New("CreateUserRestriction not implemented")
}
func (c *staticTestingDatabase) UpdateUserRestriction(ctx context.Context, data models.UserRestriction) (models.UserRestriction, error) {
	return models.UserRestriction{}, errors.New("UpdateUserRestriction not implemented")
}
