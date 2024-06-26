package models

import (
	"github.com/cufee/aftermath/internal/permissions"
)

type User struct {
	ID string

	Permissions permissions.Permissions

	Uploads       []UserContent
	Connections   []UserConnection
	Restrictions  []UserRestriction
	Subscriptions []UserSubscription
}

func (u User) HasPermission(value permissions.Permissions) bool {
	perms := u.Permissions
	for _, c := range u.Connections {
		perms.Add(c.Permissions)
	}
	for _, s := range u.Subscriptions {
		perms.Add(s.Permissions)
	}
	for _, r := range u.Restrictions {
		switch r.Type {
		case RestrictionTypePartial:
			perms.Remove(r.Restriction)
		default:
			return false
		}
	}
	return perms.Has(value)
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

func (u User) Content(kind UserContentType) (UserContent, bool) {
	for _, content := range u.Uploads {
		if content.Type == kind {
			return content, true
		}
	}
	return UserContent{}, false
}
