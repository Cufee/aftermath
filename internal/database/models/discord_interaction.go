package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/utils"
	"github.com/lucsky/cuid"
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

func ToDiscordInteraction(record *model.DiscordInteraction) DiscordInteraction {
	locale, err := language.Parse(record.Locale)
	if err != nil {
		locale = language.English
	}
	i := DiscordInteraction{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,

		Result:    record.Result,
		UserID:    record.UserID,
		GuildID:   record.GuildID,
		Snowflake: record.Snowflake,
		ChannelID: record.ChannelID,
		MessageID: record.MessageID,

		Locale:  locale,
		Type:    DiscordInteractionType(record.Type),
		EventID: record.EventID,
		Meta:    make(map[string]any, 0),
	}
	json.Unmarshal([]byte(record.Metadata), &i.Meta)
	return i
}

func FromDiscordInteraction(record *DiscordInteraction) model.DiscordInteraction {
	i := model.DiscordInteraction{
		ID:        utils.StringOr(record.ID, cuid.New()),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),

		Result:    record.Result,
		UserID:    record.UserID,
		GuildID:   record.GuildID,
		Snowflake: record.Snowflake,
		ChannelID: record.ChannelID,
		MessageID: record.MessageID,

		Locale:  record.Locale.String(),
		Type:    string(record.Type),
		EventID: record.EventID,
	}
	if record.Meta != nil {
		data, _ := json.Marshal(record.Meta)
		i.Metadata = string(data)
	}
	return i
}
