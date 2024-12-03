package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/permissions"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
)

type ConnectionType string

const (
	ConnectionTypeWargaming = ConnectionType("wargaming")
)

// Values provides list valid values for Enum.
func (ConnectionType) Values() []string {
	var kinds []string
	for _, s := range []ConnectionType{
		ConnectionTypeWargaming,
	} {
		kinds = append(kinds, string(s))
	}
	return kinds
}

type UserConnection struct {
	ID string `json:"id"`

	Type     ConnectionType `json:"type"`
	Verified bool           `json:"verified"`
	Selected bool           `json:"selected"`

	UserID      string                  `json:"userId"`
	ReferenceID string                  `json:"referenceId"`
	Permissions permissions.Permissions `json:"permissions"`

	Metadata map[string]any `json:"metadata"`
}

func ToUserConnection(record *model.UserConnection) UserConnection {
	c := UserConnection{
		ID:          record.ID,
		Type:        ConnectionType(record.Type),
		Verified:    record.Verified,
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		Metadata:    make(map[string]any, 0),
	}
	if record.Permissions != nil {
		c.Permissions = permissions.Parse(*record.Permissions, permissions.Blank)
	}
	if c.Metadata == nil {
		c.Metadata = make(map[string]any)
	}
	return c
}

func (record *UserConnection) Model() model.UserConnection {
	perms := record.Permissions.Encode()
	c := model.UserConnection{
		ID:          utils.StringOr(record.ID, cuid.New()),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Type:        string(record.Type),
		Verified:    record.Verified,
		ReferenceID: record.ReferenceID,
		Permissions: &perms,
		UserID:      record.UserID,
	}
	if data, err := json.Marshal(record.Metadata); err == nil {
		s := string(data)
		c.Metadata = &s
	}
	return c
}
