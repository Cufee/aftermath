package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/permissions"
)

// func (c UserConnection) FromModel(model *db.UserConnectionModel) UserConnection {
// 	c.ID = model.ID
// 	c.UserID = model.UserID
// 	c.Metadata = make(map[string]any)
// 	c.ReferenceID = model.ReferenceID
// 	c.Permissions = permissions.Parse(model.Permissions, permissions.Blank)

// 	c.Type = ConnectionType(model.Type)
// 	if model.MetadataEncoded != nil {
// 		_ = encoding.DecodeGob(model.MetadataEncoded, &c.Metadata)
// 	}
// 	return c
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
	// user, err := c.GetUserByID(ctx, id, opts...)
	// if err != nil {
	// 	if database.IsNotFound(err) {
	// 		model, err := c.prisma.User.CreateOne(db.User.ID.Set(id), db.User.Permissions.Set(permissions.User.Encode())).Exec(ctx)
	// 		if err != nil {
	// 			return User{}, err
	// 		}
	// 		user.ID = model.ID
	// 		user.Permissions = permissions.Parse(model.Permissions, permissions.User)
	// 	}
	// 	return user, nil
	// }

	return models.User{}, nil
}

/*
Gets a user with specified ID
  - assumes the ID is valid
*/
func (c *libsqlClient) GetUserByID(ctx context.Context, id string, opts ...userGetOption) (models.User, error) {
	// var options userGetOpts
	// for _, apply := range opts {
	// 	apply(&options)
	// }

	// var fields []db.UserRelationWith
	// if options.subscriptions {
	// 	fields = append(fields, db.User.Subscriptions.Fetch())
	// }
	// if options.connections {
	// 	fields = append(fields, db.User.Connections.Fetch())
	// }
	// if options.content {
	// 	fields = append(fields, db.User.Content.Fetch())
	// }

	// model, err := c.prisma.User.FindUnique(db.User.ID.Equals(id)).With(fields...).Exec(ctx)
	// if err != nil {
	// 	return User{}, err
	// }

	var user models.User
	// user.ID = model.ID
	// user.Permissions = permissions.Parse(model.Permissions, permissions.User)

	// if options.connections {
	// 	for _, cModel := range model.Connections() {
	// 		user.Connections = append(user.Connections, UserConnection{}.FromModel(&cModel))
	// 	}
	// }

	return user, nil
}

func (c *libsqlClient) UpsertUserWithPermissions(ctx context.Context, userID string, perms permissions.Permissions) (models.User, error) {
	// model, err := c.prisma.User.UpsertOne(db.User.ID.Equals(userID)).
	// 	Create(db.User.ID.Set(userID), db.User.Permissions.Set(perms.String())).
	// 	Update(db.User.Permissions.Set(perms.String())).Exec(ctx)
	// if err != nil {
	// 	return User{}, err
	// }

	var user models.User
	// user.ID = model.ID
	// user.Permissions = permissions.Parse(model.Permissions, permissions.User)
	return user, nil
}

func (c *libsqlClient) UpdateConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	// if connection.ReferenceID == "" {
	// 	return UserConnection{}, errors.New("connection referenceID cannot be left blank")
	// }

	// encoded, err := encoding.EncodeGob(connection.Metadata)
	// if err != nil {
	// 	return UserConnection{}, fmt.Errorf("failed to encode metadata: %w", err)
	// }

	// model, err := c.prisma.UserConnection.FindUnique(db.UserConnection.ID.Equals(connection.ID)).Update(
	// 	db.UserConnection.ReferenceID.Set(connection.ReferenceID),
	// 	db.UserConnection.Permissions.Set(connection.Permissions.Encode()),
	// 	db.UserConnection.MetadataEncoded.Set(encoded),
	// ).Exec(ctx)
	// if err != nil {
	// 	return UserConnection{}, err
	// }

	// return connection.FromModel(model), nil
	return models.UserConnection{}, nil
}

func (c *libsqlClient) UpsertConnection(ctx context.Context, connection models.UserConnection) (models.UserConnection, error) {
	// if connection.ID != "" {
	// 	return c.UpdateConnection(ctx, connection)
	// }
	// if connection.UserID == "" {
	// 	return UserConnection{}, errors.New("connection userID cannot be left blank")
	// }
	// if connection.ReferenceID == "" {
	// 	return UserConnection{}, errors.New("connection referenceID cannot be left blank")
	// }
	// if connection.Type == "" {
	// 	return UserConnection{}, errors.New("connection Type cannot be left blank")
	// }

	// encoded, err := encoding.EncodeGob(connection.Metadata)
	// if err != nil {
	// 	return UserConnection{}, fmt.Errorf("failed to encode metadata: %w", err)
	// }

	// model, err := c.prisma.UserConnection.CreateOne(
	// 	db.UserConnection.User.Link(db.User.ID.Equals(connection.UserID)),
	// 	db.UserConnection.Type.Set(string(connection.Type)),
	// 	db.UserConnection.Permissions.Set(connection.Permissions.Encode()),
	// 	db.UserConnection.ReferenceID.Set(connection.ReferenceID),
	// 	db.UserConnection.MetadataEncoded.Set(encoded),
	// ).Exec(ctx)
	// if err != nil {
	// 	return UserConnection{}, err
	// }

	// return connection.FromModel(model), nil
	return models.UserConnection{}, nil
}
