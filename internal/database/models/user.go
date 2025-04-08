package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/model"
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
		perms = perms.Add(c.Permissions)
	}
	for _, s := range u.Subscriptions {
		perms = perms.Add(s.Permissions)
	}
	for _, r := range u.Restrictions {
		switch r.Type {
		case RestrictionTypePartial:
			perms = perms.Remove(r.Restriction)
		default:
			return false
		}
	}
	return perms.Has(values...)
}

func (u User) Connection(kind ConnectionType, verified, selected *bool) (UserConnection, bool) {
	if len(u.Connections) < 1 {
		return UserConnection{}, false
	}

	for _, conn := range u.Connections {
		if ok := verified != nil; ok && (conn.Verified != *verified) {
			continue
		}
		if ok := selected != nil; ok && (conn.Selected != *selected) {
			continue
		}
		return conn, true
	}

	return UserConnection{}, false
}

func (u User) Subscription(kind SubscriptionType) (UserSubscription, bool) {
	valid, ok := u.FilterSubscriptions(kind)
	if !ok {
		return UserSubscription{}, false
	}
	return valid[0], true
}

func (u User) FilterSubscriptions(kind SubscriptionType) ([]UserSubscription, bool) {
	now := time.Now()
	var valid []UserSubscription
	for _, subscription := range u.Subscriptions {
		if subscription.Type == kind && subscription.ExpiresAt.After(now) {
			valid = append(valid, subscription)
		}
	}
	return valid, len(valid) > 0
}

func (u User) Content(kind UserContentType) (UserContent, bool) {
	for _, content := range u.Uploads {
		if content.Type == kind {
			return content, true
		}
	}
	return UserContent{}, false
}

func ToUser(record *model.User, connections []model.UserConnection, subscriptions []model.UserSubscription, content []model.UserContent, restrictions []model.UserRestriction) User {
	user := User{
		ID:          record.ID,
		Permissions: permissions.Parse(record.Permissions, permissions.Blank),
	}
	for _, c := range connections {
		user.Connections = append(user.Connections, ToUserConnection(&c))
	}
	for _, s := range subscriptions {
		user.Subscriptions = append(user.Subscriptions, ToUserSubscription(&s))
	}
	for _, c := range content {
		user.Uploads = append(user.Uploads, ToUserContent(&c))
	}
	for _, r := range restrictions {
		user.Restrictions = append(user.Restrictions, ToUserRestriction(&r))
	}
	return user
}
