package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
)

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

func ToUserContent(record *model.UserContent) UserContent {
	c := UserContent{
		ID:          record.ID,
		Type:        UserContentType(record.Type),
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,

		Value: record.Value,
		Meta:  make(map[string]any, 0),

		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
	json.Unmarshal([]byte(record.Metadata), &c.Meta)
	return c
}

func (record *UserContent) Model() model.UserContent {
	c := model.UserContent{
		ID:        utils.StringOr(record.ID, cuid.New()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Type:        string(record.Type),
		UserID:      record.UserID,
		ReferenceID: record.ReferenceID,

		Value: record.Value,
	}
	if record.Meta != nil {
		data, _ := json.Marshal(record.Meta)
		c.Metadata = string(data)
	}
	return c
}
