package models

import (
	"time"

	"golang.org/x/text/language"
)

type DiscordInteractionType string

const (
	InteractionTypeModal            = "modal"
	InteractionTypeCommand          = "command"
	InteractionTypeComponent        = "component"
	InteractionTypeAutocomplete     = "autocomplete"
	InteractionTypeAutomatedMessage = "automated_message"
)

// Values provides list valid values for Enum.
func (DiscordInteractionType) Values() []string {
	return []string{
		InteractionTypeModal,
		InteractionTypeCommand,
		InteractionTypeComponent,
		InteractionTypeAutocomplete,
		InteractionTypeAutomatedMessage,
	}
}

type DiscordInteraction struct {
	ID        string
	CreatedAt time.Time

	Result    string
	UserID    string
	GuildID   string
	ChannelID string
	MessageID string

	EventID string
	Locale  language.Tag
	Type    DiscordInteractionType

	Meta map[string]any
}
