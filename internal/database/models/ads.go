package models

import (
	"time"

	"golang.org/x/text/language"
)

type AdMessage struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Enabled   bool
	Weight    int

	Chance  float32
	Message map[language.Tag]string
	Meta    map[string]string
}

type AdEvent struct {
	CreatedAt time.Time

	UserID    string
	GuildID   string
	ChannelID string

	Locale      language.Tag
	AdMessageID string
	Meta        map[string]string
}
