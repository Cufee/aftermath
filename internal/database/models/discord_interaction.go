package models

import (
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/model"
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
	InteractionTypeGatewayEvent     = "gateway_event"
)

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
		CreatedAt: StringToTime(record.CreatedAt),

		Result:    record.Result,
		UserID:    record.UserID,
		GuildID:   record.GuildID,
		Snowflake: record.Snowflake,
		ChannelID: record.ChannelID,
		MessageID: record.MessageID,

		Locale:  locale,
		Type:    DiscordInteractionType(record.Type),
		EventID: record.EventID,
	}
	json.Unmarshal(record.Metadata, &i.Meta)

	if i.Meta == nil {
		i.Meta = make(map[string]any, 0)
	}
	return i
}

func (record *DiscordInteraction) Model() model.DiscordInteraction {
	i := model.DiscordInteraction{
		ID:        utils.StringOr(record.ID, cuid.New()),
		CreatedAt: TimeToString(time.Now()),
		UpdatedAt: TimeToString(time.Now()),

		Result:    record.Result,
		UserID:    record.UserID,
		GuildID:   record.GuildID,
		Snowflake: record.Snowflake,
		ChannelID: record.ChannelID,
		MessageID: record.MessageID,

		Locale:   record.Locale.String(),
		Type:     string(record.Type),
		EventID:  record.EventID,
		Metadata: make([]byte, 0),
	}
	i.Metadata, _ = json.Marshal(record.Meta)

	return i
}
