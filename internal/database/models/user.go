package models

import "github.com/cufee/aftermath/internal/permissions"

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
