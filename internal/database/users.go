package database

import (
	"slices"

	"github.com/cufee/aftermath-core/permissions/v2"
	"github.com/cufee/aftermath/internal/database/prisma/db"
)

type User struct {
	ID string

	Connections   []UserConnection
	Subscriptions []UserSubscription
}

type ConnectionType string

const (
	ConnectionTypeDiscord   = ConnectionType("discord")
	ConnectionTypeWargaming = ConnectionType("wargaming")
)

type UserConnection struct {
	ID string `json:"id"`

	ConnectionType ConnectionType `json:"type"`

	UserID      string `json:"userId"`
	ReferenceID string `json:"referenceId"`
	Permissions string `bson:"permissions" json:"permissions"`

	Metadata map[string]any `bson:"metadata" json:"metadata"`
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
	db.UserSubscriptionModel
}

func (sub UserSubscription) Type() SubscriptionType {
	return SubscriptionType(sub.UserSubscriptionModel.Type)
}
