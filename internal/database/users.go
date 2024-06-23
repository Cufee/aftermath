package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
)

func toUser(record *db.User, connections []*db.UserConnection, subscriptions []*db.UserSubscription, content []*db.UserContent) models.User {
	user := models.User{
		ID:          record.ID,
		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
	}
	for _, c := range connections {
		user.Connections = append(user.Connections, toUserConnection(c))
	}
	for _, s := range subscriptions {
		user.Subscriptions = append(user.Subscriptions, toUserSubscription(s))
	}
	// for _, s := range content {
	// 	user.Subscriptions = append(user.Subscriptions, toUserSubscription(s))
	// }
	return user
}

func toUserConnection(record *db.UserConnection) models.UserConnection {
	return models.UserConnection{
		ID:          record.ID,
		Type:        record.Type,
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
		Metadata:    record.Metadata,
	}
}

func toUserSubscription(record *db.UserSubscription) models.UserSubscription {
	return models.UserSubscription{
		ID:          record.ID,
		Type:        record.Type,
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		ExpiresAt:   time.Unix(record.ExpiresAt, 0),
		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
	}
}

// func toUserContent(record *db.UserSubscription) models.UserConnection {
// 	return models.UserSubscription{
// 		ID:          record.ID,
// 		Type:        record.Type,
// 		UserID:      record.UserID,
// 		ReferenceID: record.ReferenceID,
// 		ExpiresAt:   time.Unix(record.ExpiresAt, 0),
// 		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
// 	}
// }

type userGetOpts struct {
	content       bool
	connections   bool
	subscriptions bool
}

type userGetOption func(*userGetOpts)

func WithConnections() userGetOption {
	return func(ugo *userGetOpts) {
		ugo.connections = true
	}
}
func WithSubscriptions() userGetOption {
	return func(ugo *userGetOpts) {
		ugo.subscriptions = true
	}
}
func WithContent() userGetOption {
	return func(ugo *userGetOpts) {
		ugo.content = true
	}
}

/*
Gets or creates a user with specified ID
  - assumes the ID is valid
*/
func (c *libsqlClient) GetOrCreateUserByID(ctx context.Context, id string, opts ...userGetOption) (models.User, error) {
	user, err := c.GetUserByID(ctx, id, opts...)
	if err != nil && !IsNotFound(err) {
		return models.User{}, err
	}

	if IsNotFound(err) {
		record, err := c.db.User.Create().SetID(id).SetPermissions(permissions.User.String()).Save(ctx)
		if err != nil {
			return models.User{}, err
		}
		user = toUser(record, nil, nil, nil)
	}

	return user, nil
}

/*
Gets a user with specified ID
  - assumes the ID is valid
*/
func (c *libsqlClient) GetUserByID(ctx context.Context, id string, opts ...userGetOption) (models.User, error) {
	var options userGetOpts
	for _, apply := range opts {
		apply(&options)
	}

	query := c.db.User.Query().Where(user.ID(id))
	if options.subscriptions {
		query = query.WithSubscriptions()
	}
	if options.connections {
		query = query.WithConnections()
	}
	if options.content {
		query.WithContent()
	}

	record, err := query.Only(ctx)
	if err != nil {
		return models.User{}, err
	}

	return toUser(record, record.Edges.Connections, record.Edges.Subscriptions, record.Edges.Content), nil
}

func (c *libsqlClient) UpsertUserWithPermissions(ctx context.Context, userID string, perms permissions.Permissions) (models.User, error) {
	record, err := c.db.User.UpdateOneID(userID).SetPermissions(perms.String()).Save(ctx)
	if err != nil && !IsNotFound(err) {
		return models.User{}, err
	}

	if IsNotFound(err) {
		record, err = c.db.User.Create().SetID(userID).SetPermissions(perms.String()).Save(ctx)
		if err != nil {
			return models.User{}, err
		}
	}

	return toUser(record, nil, nil, nil), nil
}

func (c *libsqlClient) CreateConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	record, err := c.db.UserConnection.Create().
		SetUserID(connection.UserID).
		SetMetadata(connection.Metadata).
		SetPermissions(connection.Permissions.String()).
		SetReferenceID(connection.ReferenceID).
		SetType(connection.Type).
		Save(ctx)
	if err != nil {
		return models.UserConnection{}, err
	}
	return toUserConnection(record), err
}

func (c *libsqlClient) UpdateConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	record, err := c.db.UserConnection.UpdateOneID(connection.ID).
		SetMetadata(connection.Metadata).
		SetPermissions(connection.Permissions.String()).
		SetReferenceID(connection.ReferenceID).
		SetType(connection.Type).
		Save(ctx)
	if err != nil {
		return models.UserConnection{}, err
	}
	return toUserConnection(record), err
}

func (c *libsqlClient) UpsertConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	if connection.ID == "" {
		return c.CreateConnection(ctx, connection)
	}

	connection, err := c.UpdateConnection(ctx, connection)
	if err != nil && !IsNotFound(err) {
		return models.UserConnection{}, err
	}
	if IsNotFound(err) {
		return c.CreateConnection(ctx, connection)
	}

	return connection, nil
}
