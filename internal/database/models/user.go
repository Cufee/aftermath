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

func (u User) HasPermission(values ...permissions.Permissions) bool {
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
	return perms.Has(values...)
}

func (u User) Connection(kind ConnectionType, conditions map[string]any) (UserConnection, bool) {
	valid, ok := u.FilterConnections(kind, conditions)
	if !ok {
		return UserConnection{}, false
	}
	return valid[0], true
}

func (u User) FilterConnections(kind ConnectionType, conditions map[string]any) ([]UserConnection, bool) {
	var valid []UserConnection

outerLoop:
	for _, connection := range u.Connections {
		if connection.Type == kind {
			for key, value := range conditions {
				if connection.Metadata[key] != value {
					continue outerLoop
				}
			}
			valid = append(valid, connection)
		}
	}

	return valid, len(valid) > 0
}

func (u User) Subscription(kind SubscriptionType) (UserSubscription, bool) {
	valid, ok := u.FilterSubscriptions(kind)
	if !ok {
		return UserSubscription{}, false
	}
	return valid[0], true
}

func (u User) FilterSubscriptions(kind SubscriptionType) ([]UserSubscription, bool) {
	var valid []UserSubscription
	for _, subscription := range u.Subscriptions {
		if subscription.Type == kind {
			valid = append(valid, subscription)
		}
	}
	return valid, len(valid) > 0
}

func (u User) Content(kind UserContentType) (UserContent, bool) {
	valid, ok := u.FilterContent(kind)
	if !ok {
		return UserContent{}, false
	}
	return valid[0], true
}
func (u User) FilterContent(kind UserContentType) ([]UserContent, bool) {
	var valid []UserContent
	for _, content := range u.Uploads {
		if content.Type == kind {
			valid = append(valid, content)
		}
	}
	return valid, len(valid) > 0
}
