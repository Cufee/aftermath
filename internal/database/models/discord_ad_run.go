package models

import (
	"strings"
	"time"

	"github.com/cufee/aftermath/internal/database/gen/public/model"
	"github.com/cufee/aftermath/internal/json"
	"github.com/guregu/null/v6"
	"golang.org/x/text/language"
)

type DiscordAdRun struct {
	ID         int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CampaignID string
	ContentID  string
	GuildID    null.String
	ChannelID  string
	MessageID  string
	Locale     language.Tag
	Tags       []string
	Metadata   map[string]string
}

func ToDiscordAdRun(record *model.DiscordAdRun) DiscordAdRun {
	locale, err := language.Parse(record.Locale)
	if err != nil {
		locale = language.English
	}
	var meta map[string]string
	if record.Metadata != nil {
		json.Unmarshal(record.Metadata, &meta)
	}
	return DiscordAdRun{
		ID:         int(record.ID),
		CreatedAt:  StringToTime(record.CreatedAt),
		UpdatedAt:  StringToTime(record.UpdatedAt),
		CampaignID: record.CampaignID,
		ContentID:  record.ContentID,
		GuildID:    null.NewString(record.GuildID, record.GuildID != ""),
		ChannelID:  record.ChannelID,
		MessageID:  record.MessageID,
		Locale:     locale,
		Tags:       strings.Split(record.Tags, ","),
		Metadata:   meta,
	}
}

func (m *DiscordAdRun) Model() model.DiscordAdRun {
	locale := m.Locale.String()
	meta, _ := json.Marshal(m.Metadata)
	return model.DiscordAdRun{
		CreatedAt:  TimeToString(time.Now()),
		UpdatedAt:  TimeToString(time.Now()),
		CampaignID: m.CampaignID,
		ContentID:  m.ContentID,
		GuildID:    m.GuildID.String,
		ChannelID:  m.ChannelID,
		MessageID:  m.MessageID,
		Locale:     locale,
		Tags:       strings.Join(m.Tags, ","),
		Metadata:   meta,
	}
}
