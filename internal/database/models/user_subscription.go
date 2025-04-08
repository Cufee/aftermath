package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/model"
	"github.com/cufee/aftermath/internal/json"
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
		return permissions.Blank
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

// Manually assigned
const SubscriptionTypeThumbsCounter = SubscriptionType("thumbs-up-counter")

type UserSubscription struct {
	ID          string
	Type        SubscriptionType
	UserID      string
	ExpiresAt   time.Time
	ReferenceID string
	Permissions permissions.Permissions
	Meta        map[string]any
}

func ToUserSubscription(record *model.UserSubscription) UserSubscription {
	sub := UserSubscription{
		ID:          record.ID,
		Type:        SubscriptionType(record.Type),
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		ExpiresAt:   StringToTime(record.ExpiresAt),
		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
	}

	json.Unmarshal(record.Metadata, &sub.Meta)
	if sub.Meta == nil {
		sub.Meta = make(map[string]any, 0)
	}

	return sub
}

func (record *UserSubscription) Model() model.UserSubscription {
	sub := model.UserSubscription{
		ID:          utils.StringOr(record.ID, cuid.New()),
		CreatedAt:   TimeToString(time.Now()),
		UpdatedAt:   TimeToString(time.Now()),
		Type:        string(record.Type),
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		ExpiresAt:   TimeToString(record.ExpiresAt),
		Permissions: record.Permissions.Encode(),
	}
	sub.Metadata, _ = json.Marshal(record.Meta)
	return sub
}
