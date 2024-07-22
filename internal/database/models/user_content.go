package models

import "time"

type UserContentType string

const (
	UserContentTypeInModeration       = UserContentType("in-moderation")
	UserContentTypeClanBackground     = UserContentType("clan-background-image")
	UserContentTypePersonalBackground = UserContentType("personal-background-image")
)

func (t UserContentType) Valid() bool {
	switch t {
	case UserContentTypeClanBackground:
		return true
	case UserContentTypePersonalBackground:
		return true
	default:
		return false
	}
}

// Values provides list valid values for Enum.
func (UserContentType) Values() []string {
	var kinds []string
	for _, s := range []UserContentType{
		UserContentTypePersonalBackground,
		UserContentTypeClanBackground,
		UserContentTypeInModeration,
	} {
		kinds = append(kinds, string(s))
	}
	return kinds
}

type UserContent struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	Type        UserContentType
	UserID      string
	ReferenceID string

	Value string
	Meta  map[string]any
}
