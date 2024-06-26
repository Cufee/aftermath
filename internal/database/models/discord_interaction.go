package models

import (
	"slices"
	"time"

	"golang.org/x/text/language"
)

type DiscordInteractionType string

const (
	InteractionTypeStats = "stats"
)

var discordInteractionTypes = []DiscordInteractionType{
	InteractionTypeStats,
}

// Values provides list valid values for Enum.
func (DiscordInteractionType) Values() []string {
	var kinds []string
	for _, s := range discordInteractionTypes {
		kinds = append(kinds, string(s))
	}
	return kinds
}

func (s DiscordInteractionType) Valid() bool {
	return slices.Contains(discordInteractionTypes, s)
}

type DiscordInteraction struct {
	ID        string
	CreatedAt time.Time

	UserID      string
	Command     string
	ReferenceID string

	Type   DiscordInteractionType
	Locale language.Tag

	Options DiscordInteractionOptions
}

type DiscordInteractionOptions struct {
	BackgroundImageURL string
	PeriodStart        time.Time
	AccountID          string
}
