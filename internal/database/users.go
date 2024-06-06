package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database/prisma/db"
	"github.com/cufee/aftermath/internal/permissions"
)

type User struct {
	ID string

	Permissions permissions.Permissions

	Connections   []UserConnection
	Subscriptions []UserSubscription
}

func (u User) Connection(kind ConnectionType) (UserConnection, bool) {
	for _, connection := range u.Connections {
		if connection.Type == kind {
			return connection, true
		}
	}
	return UserConnection{}, false
}

func (u User) Subscription(kind SubscriptionType) (UserSubscription, bool) {
	for _, subscription := range u.Subscriptions {
		if subscription.Type == kind {
			return subscription, true
		}
	}
	return UserSubscription{}, false
}

type ConnectionType string

const (
	ConnectionTypeWargaming = ConnectionType("wargaming")
)

type UserConnection struct {
	ID string `json:"id"`

	Type ConnectionType `json:"type"`

	UserID      string                  `json:"userId"`
	ReferenceID string                  `json:"referenceId"`
	Permissions permissions.Permissions `json:"permissions"`

	Metadata map[string]any `json:"metadata"`
}

func (c UserConnection) FromModel(model *db.UserConnectionModel) UserConnection {
	c.ID = model.ID
	c.UserID = model.UserID
	c.Metadata = make(map[string]any)
	c.ReferenceID = model.ReferenceID
	c.Permissions = permissions.Parse(model.Permissions, permissions.Blank)

	c.Type = ConnectionType(model.Type)
	if model.MetadataEncoded != "" {
		_ = json.Unmarshal([]byte(model.MetadataEncoded), &c.Metadata)
	}
	return c
}

type SubscriptionType string

func (s SubscriptionType) GetPermissions() permissions.Permissions {
	switch s {
	case SubscriptionTypePlus:
		return permissions.SubscriptionAftermathPlus
	case SubscriptionTypePro:
		return permissions.SubscriptionAftermathPro
	case SubscriptionTypeProClan:
		return permissions.SubscriptionAftermathPro
	default:
		return permissions.User
	}
}

// Paid
const SubscriptionTypePro = SubscriptionType("aftermath-pro")
const SubscriptionTypeProClan = SubscriptionType("aftermath-pro-clan")
const SubscriptionTypePlus = SubscriptionType("aftermath-plus")

// Misc
const SubscriptionTypeSupporter = SubscriptionType("supporter")
const SubscriptionTypeVerifiedClan = SubscriptionType("verified-clan")

// Moderators
const SubscriptionTypeServerModerator = SubscriptionType("server-moderator")
const SubscriptionTypeContentModerator = SubscriptionType("content-moderator")

// Special
const SubscriptionTypeDeveloper = SubscriptionType("developer")
const SubscriptionTypeServerBooster = SubscriptionType("server-booster")
const SubscriptionTypeContentTranslator = SubscriptionType("content-translator")

var AllSubscriptionTypes = []SubscriptionType{
	SubscriptionTypePro,
	SubscriptionTypeProClan,
	SubscriptionTypePlus,
	SubscriptionTypeSupporter,
	SubscriptionTypeVerifiedClan,
	SubscriptionTypeServerModerator,
	SubscriptionTypeContentModerator,
	SubscriptionTypeDeveloper,
	SubscriptionTypeServerBooster,
	SubscriptionTypeContentTranslator,
}

func (s SubscriptionType) Valid() bool {
	return slices.Contains(AllSubscriptionTypes, s)
}

type UserSubscription struct {
	ID          string
	Type        SubscriptionType
	UserID      string
	ExpiresAt   time.Time
	ReferenceID string
	Permissions permissions.Permissions

	// id        String   @id @default(cuid())
	// createdAt DateTime @default(now())
	// updatedAt DateTime @updatedAt

	// user   User   @relation(fields: [userId], references: [id])
	// userId String

	// type        String
	// expiresAt   DateTime
	// referenceId String
	// permissions String?

	db.UserSubscriptionModel
}

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
func (c *client) GetOrCreateUserByID(ctx context.Context, id string, opts ...userGetOption) (User, error) {
	user, err := c.GetUserByID(ctx, id, opts...)
	if err != nil {
		if db.IsErrNotFound(err) {
			model, err := c.Raw.User.CreateOne(db.User.ID.Set(id), db.User.Permissions.Set(permissions.User.Encode())).Exec(ctx)
			if err != nil {
				return User{}, err
			}
			user.ID = model.ID
			user.Permissions = permissions.Parse(model.Permissions, permissions.User)
		}
		return user, nil
	}

	return user, nil
}

/*
Gets a user with specified ID
  - assumes the ID is valid
*/
func (c *client) GetUserByID(ctx context.Context, id string, opts ...userGetOption) (User, error) {
	var options userGetOpts
	for _, apply := range opts {
		apply(&options)
	}

	var fields []db.UserRelationWith
	if options.subscriptions {
		fields = append(fields, db.User.Subscriptions.Fetch())
	}
	if options.connections {
		fields = append(fields, db.User.Connections.Fetch())
	}
	if options.content {
		fields = append(fields, db.User.Content.Fetch())
	}

	model, err := c.Raw.User.FindUnique(db.User.ID.Equals(id)).With(fields...).Exec(ctx)
	if err != nil {
		return User{}, err
	}

	var user User
	user.ID = model.ID
	user.Permissions = permissions.Parse(model.Permissions, permissions.User)

	if options.connections {
		for _, cModel := range model.Connections() {
			user.Connections = append(user.Connections, UserConnection{}.FromModel(&cModel))
		}
	}

	return user, nil
}

func (c *client) UpdateConnection(ctx context.Context, connection UserConnection) (UserConnection, error) {
	if connection.ReferenceID == "" {
		return UserConnection{}, errors.New("connection referenceID cannot be left blank")
	}

	encoded, err := json.Marshal(connection.Metadata)
	if err != nil {
		return UserConnection{}, fmt.Errorf("failed to encode metadata: %w", err)
	}

	model, err := c.Raw.UserConnection.FindUnique(db.UserConnection.ID.Equals(connection.ID)).Update(
		db.UserConnection.ReferenceID.Set(connection.ReferenceID),
		db.UserConnection.Permissions.Set(connection.Permissions.Encode()),
		db.UserConnection.MetadataEncoded.Set(string(encoded)),
	).Exec(ctx)
	if err != nil {
		return UserConnection{}, err
	}

	return connection.FromModel(model), nil
}

func (c *client) UpsertConnection(ctx context.Context, connection UserConnection) (UserConnection, error) {
	if connection.ID != "" {
		return c.UpdateConnection(ctx, connection)
	}
	if connection.UserID == "" {
		return UserConnection{}, errors.New("connection userID cannot be left blank")
	}
	if connection.ReferenceID == "" {
		return UserConnection{}, errors.New("connection referenceID cannot be left blank")
	}
	if connection.Type == "" {
		return UserConnection{}, errors.New("connection Type cannot be left blank")
	}

	encoded, err := json.Marshal(connection.Metadata)
	if err != nil {
		return UserConnection{}, fmt.Errorf("failed to encode metadata: %w", err)
	}

	model, err := c.Raw.UserConnection.CreateOne(
		db.UserConnection.User.Link(db.User.ID.Equals(connection.UserID)),
		db.UserConnection.Type.Set(string(connection.Type)),
		db.UserConnection.Permissions.Set(connection.Permissions.Encode()),
		db.UserConnection.ReferenceID.Set(connection.ReferenceID),
		db.UserConnection.MetadataEncoded.Set(string(encoded)),
	).Exec(ctx)
	if err != nil {
		return UserConnection{}, err
	}

	return connection.FromModel(model), nil
}
