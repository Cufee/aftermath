package database

import (
	"context"
	"encoding/json"
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

/*
Gets or creates a user with specified ID
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

	model, err := c.prisma.User.FindUnique(db.User.ID.Equals(id)).With(fields...).Exec(ctx)
	if err != nil {
		if !db.IsErrNotFound(err) {
			return User{}, err
		}
		model, err = c.prisma.User.CreateOne(db.User.ID.Set(id), db.User.Permissions.Set(permissions.User.Encode())).Exec(ctx)
		if err != nil {
			return User{}, err
		}
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
