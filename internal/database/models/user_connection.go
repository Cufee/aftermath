package models

import "github.com/cufee/aftermath/internal/permissions"

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

	Type ConnectionType `json:"type"`

	UserID      string                  `json:"userId"`
	ReferenceID string                  `json:"referenceId"`
	Permissions permissions.Permissions `json:"permissions"`

	Metadata map[string]any `json:"metadata"`
}
