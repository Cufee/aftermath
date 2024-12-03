package models

import (
	"slices"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
)

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

// Values provides list valid values for Enum.
func (SubscriptionType) Values() []string {
	var kinds []string
	for _, s := range AllSubscriptionTypes {
		kinds = append(kinds, string(s))
	}
	return kinds
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
}

func ToUserSubscription(record *model.UserSubscription) UserSubscription {
	return UserSubscription{
		ID:          utils.StringOr(record.ID, cuid.New()),
		Type:        SubscriptionType(record.Type),
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		ExpiresAt:   record.ExpiresAt,
		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
	}
}
