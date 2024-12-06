package database

import (
	"context"
	"errors"
	"time"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/utils"
	s "github.com/go-jet/jet/v2/sqlite"
)

type userOpts struct {
	get struct {
		content       bool
		connections   bool
		subscriptions bool
	}
	insert struct {
		username    *string
		permissions *permissions.Permissions
	}
}

type UserQueryOptions []UserQueryOption

func (o UserQueryOptions) ToOptions() userOpts {
	var opts userOpts
	for _, apply := range o {
		apply(&opts)
	}
	return opts
}

type UserQueryOption func(*userOpts)

func WithConnections() UserQueryOption {
	return func(ugo *userOpts) {
		ugo.get.connections = true
	}
}
func WithSubscriptions() UserQueryOption {
	return func(ugo *userOpts) {
		ugo.get.subscriptions = true
	}
}
func WithContent() UserQueryOption {
	return func(ugo *userOpts) {
		ugo.get.content = true
	}
}

func AddPermissions(perms permissions.Permissions) UserQueryOption {
	return func(ugo *userOpts) {
		ugo.insert.permissions = &perms
	}
}
func AddUsername(name string) UserQueryOption {
	return func(ugo *userOpts) {
		ugo.insert.username = &name
	}
}

/*
Gets or creates a user with specified ID
  - assumes the ID is valid
*/
func (c *client) GetOrCreateUserByID(ctx context.Context, id string, opts ...UserQueryOption) (models.User, error) {
	usr, err := c.GetUserByID(ctx, id, opts...)
	if err == nil {
		return usr, nil
	}
	if !IsNotFound(err) {
		return models.User{}, err
	}

	options := UserQueryOptions(opts).ToOptions()
	user := m.User{
		ID:           id,
		CreatedAt:    models.TimeToString(time.Now()),
		UpdatedAt:    models.TimeToString(time.Now()),
		Permissions:  permissions.User.String(),
		FeatureFlags: make([]byte, 0),
	}
	if options.insert.permissions != nil {
		user.Permissions = permissions.User.Add(*options.insert.permissions).String()
	}
	if options.insert.username != nil {
		user.Username = *options.insert.username
	}

	stmt := t.User.
		INSERT(t.User.AllColumns).
		MODEL(user).
		RETURNING(t.User.AllColumns)

	err = c.query(ctx, stmt, &user)
	if err != nil {
		return models.User{}, err
	}

	return models.ToUser(&user, nil, nil, nil, nil), nil
}

/*
Gets a user with specified ID
  - assumes the ID is valid
*/
func (c *client) GetUserByID(ctx context.Context, id string, opts ...UserQueryOption) (models.User, error) {
	options := UserQueryOptions(opts).ToOptions()

	var from s.ReadableTable = t.User.LEFT_JOIN(t.UserRestriction, t.UserRestriction.UserID.EQ(t.User.ID))
	projections := []s.Projection{t.UserRestriction.AllColumns}

	if options.get.subscriptions {
		from = from.LEFT_JOIN(t.UserSubscription, t.UserSubscription.UserID.EQ(t.User.ID))
		projections = append(projections, t.UserSubscription.AllColumns)
	}
	if options.get.connections {
		from = from.LEFT_JOIN(t.UserConnection, t.UserConnection.UserID.EQ(t.User.ID))
		projections = append(projections, t.UserConnection.AllColumns)
	}
	if options.get.content {
		from = from.LEFT_JOIN(t.UserContent, t.UserContent.UserID.EQ(t.User.ID))
		projections = append(projections, t.UserContent.AllColumns)
	}

	stmt := s.
		SELECT(t.User.AllColumns, projections...).
		FROM(from).
		WHERE(t.User.ID.EQ(s.String(id)))

	var record struct {
		m.User
		Restrictions []m.UserRestriction

		Content       []m.UserContent
		Connections   []m.UserConnection
		Subscriptions []m.UserSubscription
	}
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.User{}, err
	}

	return models.ToUser(&record.User, record.Connections, record.Subscriptions, record.Content, record.Restrictions), nil
}

type connectionOpts struct {
	get struct {
		kind     *models.ConnectionType
		verified *bool
		selected *bool
	}
}

type connectionQueryOptions []ConnectionQueryOption

type ConnectionQueryOption func(*connectionOpts)

func (o connectionQueryOptions) ToOptions() connectionOpts {
	var opts connectionOpts
	for _, apply := range o {
		apply(&opts)
	}
	return opts
}

func ConnectionType(kind models.ConnectionType) ConnectionQueryOption {
	return func(ugo *connectionOpts) {
		ugo.get.kind = &kind
	}
}

func ConnectionVerified(verified bool) ConnectionQueryOption {
	return func(ugo *connectionOpts) {
		ugo.get.verified = utils.Pointer(verified)
	}
}

func ConnectionSelected(selected bool) ConnectionQueryOption {
	return func(ugo *connectionOpts) {
		ugo.get.selected = utils.Pointer(selected)
	}
}

func (c *client) GetUserConnection(ctx context.Context, id string) (models.UserConnection, error) {
	stmt := t.UserConnection.
		SELECT(t.UserConnection.AllColumns).
		WHERE(t.UserConnection.ID.EQ(s.String(id)))

	var record m.UserConnection
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.UserConnection{}, err
	}

	return models.ToUserConnection(&record), nil
}

func (c *client) FindUserConnections(ctx context.Context, userID string, opts ...ConnectionQueryOption) ([]models.UserConnection, error) {
	options := connectionQueryOptions(opts).ToOptions()

	where := []s.BoolExpression{t.UserConnection.UserID.EQ(s.String(userID))}
	if options.get.kind != nil {
		where = append(where, t.UserConnection.Type.EQ(s.String(string(*options.get.kind))))
	}
	if options.get.verified != nil {
		where = append(where, t.UserConnection.Verified.EQ(s.Bool(*options.get.verified)))
	}
	if options.get.selected != nil {
		where = append(where, t.UserConnection.Selected.EQ(s.Bool(*options.get.selected)))
	}

	stmt := t.UserConnection.
		SELECT(t.UserConnection.AllColumns).
		WHERE(s.AND(where...))

	var records []m.UserConnection
	err := c.query(ctx, stmt, &records)
	if err != nil {
		return nil, err
	}

	var connections []models.UserConnection
	for _, r := range records {
		connections = append(connections, models.ToUserConnection(&r))
	}

	return connections, nil
}

func (c *client) CreateUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	model := connection.Model()
	stmt := t.UserConnection.
		INSERT(t.UserConnection.AllColumns).
		MODEL(model).
		RETURNING(t.UserConnection.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.UserConnection{}, err
	}

	return models.ToUserConnection(&model), err
}

func (c *client) UpdateUserConnection(ctx context.Context, id string, connection models.UserConnection) (models.UserConnection, error) {
	model := connection.Model()
	stmt := t.UserConnection.
		UPDATE(
			t.UserConnection.UpdatedAt,
			t.UserConnection.Type,
			t.UserConnection.Verified,
			t.UserConnection.Selected,
			t.UserConnection.ReferenceID,
			t.UserConnection.Permissions,
			t.UserConnection.Metadata,
		).
		MODEL(model).
		WHERE(t.UserConnection.ID.EQ(s.String(id))).
		RETURNING(t.UserConnection.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.UserConnection{}, err
	}

	return models.ToUserConnection(&model), nil

}

func (c *client) UpsertUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	if connection.ID == "" {
		return c.CreateUserConnection(ctx, connection)
	}

	connection, err := c.UpdateUserConnection(ctx, connection.ID, connection)
	if IsNotFound(err) {
		return c.CreateUserConnection(ctx, connection)
	}
	return connection, err
}

func (c *client) DeleteUserConnection(ctx context.Context, userID, connectionID string) error {
	stmt := t.UserConnection.DELETE().WHERE(s.AND(t.UserConnection.ID.EQ(s.String(connectionID)), t.UserConnection.UserID.EQ(s.String(userID))))
	_, err := c.exec(ctx, stmt)
	return err
}

func (c *client) GetUserContent(ctx context.Context, id string) (models.UserContent, error) {
	stmt := t.UserContent.
		SELECT(t.UserContent.AllColumns).
		WHERE(t.UserContent.ID.EQ(s.String(id)))

	var record m.UserContent
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.UserContent{}, err
	}

	return models.ToUserContent(&record), nil

}

func (c *client) GetUserContentFromRef(ctx context.Context, referenceID string, kind models.UserContentType) (models.UserContent, error) {
	stmt := t.UserContent.
		SELECT(t.UserContent.AllColumns).
		WHERE(s.AND(
			t.UserContent.Type.EQ(s.String(string(kind))),
			t.UserContent.ReferenceID.EQ(s.String(referenceID)),
		))

	var record m.UserContent
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.UserContent{}, err
	}

	return models.ToUserContent(&record), nil
}

func (c *client) FindUserContentFromRefs(ctx context.Context, kind models.UserContentType, referenceIDs ...string) ([]models.UserContent, error) {
	if len(referenceIDs) < 1 {
		return nil, errors.New("at least one reference id is required")
	}

	stmt := t.UserContent.
		SELECT(t.UserContent.AllColumns).
		WHERE(s.AND(
			t.UserContent.Type.EQ(s.String(string(kind))),
			t.UserContent.ReferenceID.IN(stringsToExp(referenceIDs)...),
		))

	var record []m.UserContent
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return nil, err
	}

	var content []models.UserContent
	for _, c := range record {
		content = append(content, models.ToUserContent(&c))
	}

	return content, nil
}

func (c *client) CreateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	model := content.Model()
	stmt := t.UserContent.
		INSERT(t.UserContent.AllColumns).
		MODEL(model).
		RETURNING(t.UserContent.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.UserContent{}, err
	}

	return models.ToUserContent(&model), nil
}

func (c *client) UpdateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	model := content.Model()
	stmt := t.UserContent.
		UPDATE(
			t.UserContent.UpdatedAt,
			t.UserContent.Type,
			t.UserContent.UserID,
			t.UserContent.ReferenceID,
			t.UserContent.Value,
			t.UserContent.Metadata,
		).
		MODEL(model).
		RETURNING(t.UserContent.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.UserContent{}, err
	}

	return models.ToUserContent(&model), nil
}

func (c *client) UpsertUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	if content.ID != "" {
		return c.UpdateUserContent(ctx, content)
	}
	return c.CreateUserContent(ctx, content)
}

func (c *client) DeleteUserContent(ctx context.Context, id string) error {
	_, err := c.exec(ctx, t.UserContent.DELETE().WHERE(t.UserContent.ID.EQ(s.String(id))))
	return err
}

func (c *client) CreateUserSubscription(ctx context.Context, subscription models.UserSubscription) (models.UserSubscription, error) {
	model := subscription.Model()
	stmt := t.UserSubscription.
		INSERT(t.UserSubscription.AllColumns).
		MODEL(model).
		RETURNING(t.UserSubscription.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.UserSubscription{}, err
	}

	return models.ToUserSubscription(&model), nil
}

func (c *client) UpdateUserSubscription(ctx context.Context, id string, subscription models.UserSubscription) (models.UserSubscription, error) {
	model := subscription.Model()
	stmt := t.UserSubscription.
		UPDATE(t.UserSubscription.AllColumns.Except(
			t.UserSubscription.ID,
			t.UserSubscription.Type,
			t.UserSubscription.UserID,
			t.UserSubscription.CreatedAt,
		)).
		MODEL(model).
		WHERE(t.UserSubscription.ID.EQ(s.String(id))).
		RETURNING(t.UserSubscription.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.UserSubscription{}, err
	}

	return models.ToUserSubscription(&model), nil
}
