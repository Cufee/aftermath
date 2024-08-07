// Code generated by ent, DO NOT EDIT.

package db

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Permissions holds the value of the "permissions" field.
	Permissions string `json:"permissions,omitempty"`
	// FeatureFlags holds the value of the "feature_flags" field.
	FeatureFlags []string `json:"feature_flags,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// DiscordInteractions holds the value of the discord_interactions edge.
	DiscordInteractions []*DiscordInteraction `json:"discord_interactions,omitempty"`
	// Subscriptions holds the value of the subscriptions edge.
	Subscriptions []*UserSubscription `json:"subscriptions,omitempty"`
	// Connections holds the value of the connections edge.
	Connections []*UserConnection `json:"connections,omitempty"`
	// Widgets holds the value of the widgets edge.
	Widgets []*WidgetSettings `json:"widgets,omitempty"`
	// Content holds the value of the content edge.
	Content []*UserContent `json:"content,omitempty"`
	// Sessions holds the value of the sessions edge.
	Sessions []*Session `json:"sessions,omitempty"`
	// ModerationRequests holds the value of the moderation_requests edge.
	ModerationRequests []*ModerationRequest `json:"moderation_requests,omitempty"`
	// ModerationActions holds the value of the moderation_actions edge.
	ModerationActions []*ModerationRequest `json:"moderation_actions,omitempty"`
	// Restrictions holds the value of the restrictions edge.
	Restrictions []*UserRestriction `json:"restrictions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [9]bool
}

// DiscordInteractionsOrErr returns the DiscordInteractions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DiscordInteractionsOrErr() ([]*DiscordInteraction, error) {
	if e.loadedTypes[0] {
		return e.DiscordInteractions, nil
	}
	return nil, &NotLoadedError{edge: "discord_interactions"}
}

// SubscriptionsOrErr returns the Subscriptions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SubscriptionsOrErr() ([]*UserSubscription, error) {
	if e.loadedTypes[1] {
		return e.Subscriptions, nil
	}
	return nil, &NotLoadedError{edge: "subscriptions"}
}

// ConnectionsOrErr returns the Connections value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ConnectionsOrErr() ([]*UserConnection, error) {
	if e.loadedTypes[2] {
		return e.Connections, nil
	}
	return nil, &NotLoadedError{edge: "connections"}
}

// WidgetsOrErr returns the Widgets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WidgetsOrErr() ([]*WidgetSettings, error) {
	if e.loadedTypes[3] {
		return e.Widgets, nil
	}
	return nil, &NotLoadedError{edge: "widgets"}
}

// ContentOrErr returns the Content value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ContentOrErr() ([]*UserContent, error) {
	if e.loadedTypes[4] {
		return e.Content, nil
	}
	return nil, &NotLoadedError{edge: "content"}
}

// SessionsOrErr returns the Sessions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SessionsOrErr() ([]*Session, error) {
	if e.loadedTypes[5] {
		return e.Sessions, nil
	}
	return nil, &NotLoadedError{edge: "sessions"}
}

// ModerationRequestsOrErr returns the ModerationRequests value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ModerationRequestsOrErr() ([]*ModerationRequest, error) {
	if e.loadedTypes[6] {
		return e.ModerationRequests, nil
	}
	return nil, &NotLoadedError{edge: "moderation_requests"}
}

// ModerationActionsOrErr returns the ModerationActions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ModerationActionsOrErr() ([]*ModerationRequest, error) {
	if e.loadedTypes[7] {
		return e.ModerationActions, nil
	}
	return nil, &NotLoadedError{edge: "moderation_actions"}
}

// RestrictionsOrErr returns the Restrictions value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RestrictionsOrErr() ([]*UserRestriction, error) {
	if e.loadedTypes[8] {
		return e.Restrictions, nil
	}
	return nil, &NotLoadedError{edge: "restrictions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldFeatureFlags:
			values[i] = new([]byte)
		case user.FieldID, user.FieldUsername, user.FieldPermissions:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPermissions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field permissions", values[i])
			} else if value.Valid {
				u.Permissions = value.String
			}
		case user.FieldFeatureFlags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field feature_flags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.FeatureFlags); err != nil {
					return fmt.Errorf("unmarshal field feature_flags: %w", err)
				}
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryDiscordInteractions queries the "discord_interactions" edge of the User entity.
func (u *User) QueryDiscordInteractions() *DiscordInteractionQuery {
	return NewUserClient(u.config).QueryDiscordInteractions(u)
}

// QuerySubscriptions queries the "subscriptions" edge of the User entity.
func (u *User) QuerySubscriptions() *UserSubscriptionQuery {
	return NewUserClient(u.config).QuerySubscriptions(u)
}

// QueryConnections queries the "connections" edge of the User entity.
func (u *User) QueryConnections() *UserConnectionQuery {
	return NewUserClient(u.config).QueryConnections(u)
}

// QueryWidgets queries the "widgets" edge of the User entity.
func (u *User) QueryWidgets() *WidgetSettingsQuery {
	return NewUserClient(u.config).QueryWidgets(u)
}

// QueryContent queries the "content" edge of the User entity.
func (u *User) QueryContent() *UserContentQuery {
	return NewUserClient(u.config).QueryContent(u)
}

// QuerySessions queries the "sessions" edge of the User entity.
func (u *User) QuerySessions() *SessionQuery {
	return NewUserClient(u.config).QuerySessions(u)
}

// QueryModerationRequests queries the "moderation_requests" edge of the User entity.
func (u *User) QueryModerationRequests() *ModerationRequestQuery {
	return NewUserClient(u.config).QueryModerationRequests(u)
}

// QueryModerationActions queries the "moderation_actions" edge of the User entity.
func (u *User) QueryModerationActions() *ModerationRequestQuery {
	return NewUserClient(u.config).QueryModerationActions(u)
}

// QueryRestrictions queries the "restrictions" edge of the User entity.
func (u *User) QueryRestrictions() *UserRestrictionQuery {
	return NewUserClient(u.config).QueryRestrictions(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("db: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("permissions=")
	builder.WriteString(u.Permissions)
	builder.WriteString(", ")
	builder.WriteString("feature_flags=")
	builder.WriteString(fmt.Sprintf("%v", u.FeatureFlags))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
