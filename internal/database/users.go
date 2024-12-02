package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
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
	if !errors.Is(err, sql.ErrNoRows) {
		return models.User{}, err
	}

	options := UserQueryOptions(opts).ToOptions()
	user := m.User{
		ID:          id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Permissions: permissions.User.String(),
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

func (c *client) GetUserConnection(ctx context.Context, id string) (models.UserConnection, error) {
	stmt := t.UserConnection.
		SELECT(t.UserConnection.AllColumns).
		WHERE(t.UserConnection.ID.EQ(s.String(id)))

	var record model.UserConnection
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.UserConnection{}, err
	}

	return models.ToUserConnection(&record), nil
}

func (c *client) CreateUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	model := models.FromUserConnection(&connection)
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

// func (c *client) UpdateUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
// 	record, err := c.db.UserConnection.UpdateOneID(connection.ID).
// 		SetMetadata(connection.Metadata).
// 		SetPermissions(connection.Permissions.String()).
// 		SetReferenceID(connection.ReferenceID).
// 		SetType(connection.Type).
// 		Save(ctx)
// 	if err != nil {
// 		return models.UserConnection{}, err
// 	}
// 	return toUserConnection(record), err
// }

// func (c *client) UpsertUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
// 	if connection.ID == "" {
// 		return c.CreateUserConnection(ctx, connection)
// 	}

// 	connection, err := c.UpdateUserConnection(ctx, connection)
// 	if err != nil && !IsNotFound(err) {
// 		return models.UserConnection{}, err
// 	}
// 	if IsNotFound(err) {
// 		return c.CreateUserConnection(ctx, connection)
// 	}

// 	return connection, nil
// }

// func (c *client) DeleteUserConnection(ctx context.Context, userID, connectionID string) error {
// 	_, err := c.db.UserConnection.Delete().Where(userconnection.ID(connectionID), userconnection.UserID(userID)).Exec(ctx)
// 	return err
// }

// func (c *client) GetUserContent(ctx context.Context, id string) (models.UserContent, error) {
// 	record, err := c.db.UserContent.Get(ctx, id)
// 	if err != nil {
// 		return models.UserContent{}, err
// 	}

// 	return toUserContent(record), nil
// }

// func (c *client) GetUserContentFromRef(ctx context.Context, referenceID string, kind models.UserContentType) (models.UserContent, error) {
// 	record, err := c.db.UserContent.Query().Where(usercontent.ReferenceID(referenceID), usercontent.TypeEQ(kind)).First(ctx)
// 	if err != nil {
// 		return models.UserContent{}, err
// 	}

// 	return toUserContent(record), nil
// }

// func (c *client) FindUserContentFromRefs(ctx context.Context, kind models.UserContentType, referenceIDs ...string) ([]models.UserContent, error) {
// 	if len(referenceIDs) < 1 {
// 		return nil, errors.New("at least one reference id is required")
// 	}

// 	records, err := c.db.UserContent.Query().Where(usercontent.ReferenceIDIn(referenceIDs...), usercontent.TypeEQ(kind)).All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var content []models.UserContent
// 	for _, r := range records {
// 		content = append(content, toUserContent(r))
// 	}

// 	return content, nil
// }

// func (c *client) CreateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
// 	user, err := c.db.User.Get(ctx, content.UserID)
// 	if err != nil {
// 		return models.UserContent{}, err
// 	}

// 	record, err := c.db.UserContent.Create().
// 		SetMetadata(content.Meta).
// 		SetReferenceID(content.ReferenceID).
// 		SetType(content.Type).
// 		SetUser(user).
// 		SetValue(content.Value).
// 		Save(ctx)
// 	if err != nil {
// 		return models.UserContent{}, err
// 	}

// 	return toUserContent(record), nil
// }

// func (c *client) UpdateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
// 	record, err := c.db.UserContent.UpdateOneID(content.ID).
// 		SetMetadata(content.Meta).
// 		SetReferenceID(content.ReferenceID).
// 		SetType(content.Type).
// 		SetValue(content.Value).
// 		Save(ctx)
// 	if err != nil {
// 		return models.UserContent{}, err
// 	}

// 	return toUserContent(record), nil
// }

// func (c *client) UpsertUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
// 	id, err := c.db.UserContent.Query().Where(usercontent.UserID(content.UserID), usercontent.TypeEQ(content.Type)).FirstID(ctx)
// 	if IsNotFound(err) {
// 		return c.CreateUserContent(ctx, content)
// 	}

// 	content.ID = id
// 	return c.UpdateUserContent(ctx, content)

// }

// func (c *client) DeleteUserContent(ctx context.Context, id string) error {
// 	err := c.db.UserContent.DeleteOneID(id).Exec(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
