package models

import (
	"time"

	"golang.org/x/text/language"
)

type DiscordInteractionType string

const (
	InteractionTypeUnknown          = "unknown"
	InteractionTypeModal            = "modal"
	InteractionTypeCommand          = "command"
	InteractionTypeFollowUp         = "follow_up"
	InteractionTypeComponent        = "component"
	InteractionTypeAutocomplete     = "autocomplete"
	InteractionTypeAutomatedMessage = "automated_message"
)

// Values provides list valid values for Enum.
func (DiscordInteractionType) Values() []string {
	return []string{
		InteractionTypeUnknown,
		InteractionTypeModal,
		InteractionTypeCommand,
		InteractionTypeFollowUp,
		InteractionTypeComponent,
		InteractionTypeAutocomplete,
		InteractionTypeAutomatedMessage,
	}
}

type DiscordInteraction struct {
	ID        string
	CreatedAt time.Time

	Snowflake string
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
