package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/model"
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
		Selected:    record.Selected,
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,
		Permissions: permissions.Blank,
	}
	json.Unmarshal(record.Metadata, &c.Metadata)

	if record.Permissions != nil {
		c.Permissions = permissions.Parse(*record.Permissions, c.Permissions)
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
		CreatedAt:   TimeToString(time.Now()),
		UpdatedAt:   TimeToString(time.Now()),
		Type:        string(record.Type),
		Verified:    record.Verified,
		Selected:    record.Selected,
		ReferenceID: record.ReferenceID,
		Permissions: &perms,
		UserID:      record.UserID,
	}
	c.Metadata, _ = json.Marshal(record.Metadata)
	return c
}
