package database

import (
	"context"
	"errors"
	"time"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/ent/db/userconnection"
	"github.com/cufee/aftermath/internal/database/ent/db/usercontent"
	"github.com/cufee/aftermath/internal/database/ent/db/userrestriction"
	"github.com/cufee/aftermath/internal/database/ent/db/usersubscription"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
)

func toUser(record *db.User, connections []*db.UserConnection, subscriptions []*db.UserSubscription, content []*db.UserContent, restrictions []*db.UserRestriction) models.User {
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
	for _, c := range content {
		user.Uploads = append(user.Uploads, toUserContent(c))
	}
	for _, r := range restrictions {
		user.Restrictions = append(user.Restrictions, toUserRestriction(r))
	}
	return user
}

func toUserRestriction(record *db.UserRestriction) models.UserRestriction {
	return models.UserRestriction{
		ID:     record.ID,
		Type:   record.Type,
		UserID: record.UserID,

		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		ExpiresAt: record.ExpiresAt,

		ModeratorComment: record.ModeratorComment,
		PublicReason:     record.PublicReason,
		Restriction:      permissions.Parse(record.Restriction, permissions.Blank),
		Events:           record.Events,
	}
}

func toUserConnection(record *db.UserConnection) models.UserConnection {
	c := models.UserConnection{
		ID:          record.ID,
		Type:        record.Type,
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
		Metadata:    record.Metadata,
	}
	if c.Metadata == nil {
		c.Metadata = make(map[string]any)
	}
	return c
}

func toUserSubscription(record *db.UserSubscription) models.UserSubscription {
	return models.UserSubscription{
		ID:          record.ID,
		Type:        record.Type,
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		ExpiresAt:   record.ExpiresAt,
		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
	}
}

func toUserContent(record *db.UserContent) models.UserContent {
	c := models.UserContent{
		ID:          record.ID,
		Type:        record.Type,
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,

		Value: record.Value,
		Meta:  record.Metadata,

		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
	if c.Meta == nil {
		c.Meta = make(map[string]any)
	}
	return c
}

type userGetOpts struct {
	content       bool
	connections   bool
	subscriptions bool
}

type UserGetOptions []UserGetOption

func (o UserGetOptions) ToOptions() userGetOpts {
	var opts userGetOpts
	for _, apply := range o {
		apply(&opts)
	}
	return opts
}

type UserGetOption func(*userGetOpts)

func WithConnections() UserGetOption {
	return func(ugo *userGetOpts) {
		ugo.connections = true
	}
}
func WithSubscriptions() UserGetOption {
	return func(ugo *userGetOpts) {
		ugo.subscriptions = true
	}
}
func WithContent() UserGetOption {
	return func(ugo *userGetOpts) {
		ugo.content = true
	}
}

/*
Gets or creates a user with specified ID
  - assumes the ID is valid
*/
func (c *client) GetOrCreateUserByID(ctx context.Context, id string, opts ...UserGetOption) (models.User, error) {
	user, err := c.GetUserByID(ctx, id, opts...)
	if err != nil && !IsNotFound(err) {
		return models.User{}, err
	}

	if IsNotFound(err) {
		record, err := c.db.User.Create().SetID(id).SetPermissions(permissions.User.String()).Save(ctx)
		if err != nil {
			return models.User{}, err
		}
		user = toUser(record, nil, nil, nil, nil)
	}

	return user, nil
}

/*
Gets a user with specified ID
  - assumes the ID is valid
*/
func (c *client) GetUserByID(ctx context.Context, id string, opts ...UserGetOption) (models.User, error) {
	var options userGetOpts
	for _, apply := range opts {
		apply(&options)
	}

	query := c.db.User.Query().Where(user.ID(id)).WithRestrictions(func(urq *db.UserRestrictionQuery) { urq.Where(userrestriction.ExpiresAtGT(time.Now())) })
	if options.subscriptions {
		query = query.WithSubscriptions(func(usq *db.UserSubscriptionQuery) { usq.Where(usersubscription.ExpiresAtGT(time.Now())) })
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

	return toUser(record, record.Edges.Connections, record.Edges.Subscriptions, record.Edges.Content, record.Edges.Restrictions), nil
}

func (c *client) UpsertUserWithPermissions(ctx context.Context, userID string, perms permissions.Permissions) (models.User, error) {
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

	return toUser(record, nil, nil, nil, nil), nil
}

func (c *client) GetUserConnection(ctx context.Context, id string) (models.UserConnection, error) {
	record, err := c.db.UserConnection.Get(ctx, id)
	return toUserConnection(record), err
}

func (c *client) CreateUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	record, err := c.db.UserConnection.Create().
		SetUser(c.db.User.GetX(ctx, connection.UserID)).
		SetPermissions(connection.Permissions.String()).
		SetReferenceID(connection.ReferenceID).
		SetMetadata(connection.Metadata).
		SetType(connection.Type).
		Save(ctx)
	if err != nil {
		return models.UserConnection{}, err
	}
	return toUserConnection(record), err
}

func (c *client) UpdateUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
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

func (c *client) UpsertUserConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	if connection.ID == "" {
		return c.CreateUserConnection(ctx, connection)
	}

	connection, err := c.UpdateUserConnection(ctx, connection)
	if err != nil && !IsNotFound(err) {
		return models.UserConnection{}, err
	}
	if IsNotFound(err) {
		return c.CreateUserConnection(ctx, connection)
	}

	return connection, nil
}

func (c *client) DeleteUserConnection(ctx context.Context, userID, connectionID string) error {
	_, err := c.db.UserConnection.Delete().Where(userconnection.ID(connectionID), userconnection.UserID(userID)).Exec(ctx)
	return err
}

func (c *client) GetUserContent(ctx context.Context, id string) (models.UserContent, error) {
	record, err := c.db.UserContent.Get(ctx, id)
	if err != nil {
		return models.UserContent{}, err
	}

	return toUserContent(record), nil
}

func (c *client) GetUserContentFromRef(ctx context.Context, referenceID string, kind models.UserContentType) (models.UserContent, error) {
	record, err := c.db.UserContent.Query().Where(usercontent.ReferenceID(referenceID), usercontent.TypeEQ(kind)).First(ctx)
	if err != nil {
		return models.UserContent{}, err
	}

	return toUserContent(record), nil
}

func (c *client) FindUserContentFromRefs(ctx context.Context, kind models.UserContentType, referenceIDs ...string) ([]models.UserContent, error) {
	if len(referenceIDs) < 1 {
		return nil, errors.New("at least one reference id is required")
	}

	records, err := c.db.UserContent.Query().Where(usercontent.ReferenceIDIn(referenceIDs...), usercontent.TypeEQ(kind)).All(ctx)
	if err != nil {
		return nil, err
	}

	var content []models.UserContent
	for _, r := range records {
		content = append(content, toUserContent(r))
	}

	return content, nil
}

func (c *client) CreateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	user, err := c.db.User.Get(ctx, content.UserID)
	if err != nil {
		return models.UserContent{}, err
	}

	record, err := c.db.UserContent.Create().
		SetMetadata(content.Meta).
		SetReferenceID(content.ReferenceID).
		SetType(content.Type).
		SetUser(user).
		SetValue(content.Value).
		Save(ctx)
	if err != nil {
		return models.UserContent{}, err
	}

	return toUserContent(record), nil
}

func (c *client) UpdateUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	record, err := c.db.UserContent.UpdateOneID(content.ID).
		SetMetadata(content.Meta).
		SetReferenceID(content.ReferenceID).
		SetType(content.Type).
		SetValue(content.Value).
		Save(ctx)
	if err != nil {
		return models.UserContent{}, err
	}

	return toUserContent(record), nil
}

func (c *client) UpsertUserContent(ctx context.Context, content models.UserContent) (models.UserContent, error) {
	id, err := c.db.UserContent.Query().Where(usercontent.UserID(content.UserID), usercontent.TypeEQ(content.Type)).FirstID(ctx)
	if IsNotFound(err) {
		return c.CreateUserContent(ctx, content)
	}

	content.ID = id
	return c.UpdateUserContent(ctx, content)

}

func (c *client) DeleteUserContent(ctx context.Context, id string) error {
	err := c.db.UserContent.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
