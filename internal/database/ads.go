package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/models"
	"golang.org/x/text/language"
)

func toAdMessage(r *db.AdMessage) models.AdMessage {
	return models.AdMessage{
		ID:        r.ID,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Enabled:   r.Enabled,
		Weight:    r.Weight,
		Chance:    r.Chance,
		Message:   r.Message,
	}
}

func toAdEvent(r *db.AdEvent) models.AdEvent {
	tag, err := language.Parse(r.Locale)
	if err != nil {
		tag = language.English
	}
	return models.AdEvent{
		CreatedAt:   r.CreatedAt,
		UserID:      r.UserID,
		GuildID:     r.GuildID,
		ChannelID:   r.ChannelID,
		Locale:      tag,
		AdMessageID: r.MessageID,
	}
}

func (c *client) CreateAdMessage(ctx context.Context, message models.AdMessage) (models.AdMessage, error) {
	return models.AdMessage{}, nil
}

func (c *client) GetAdMessage(ctx context.Context, messageID string) (models.AdMessage, error) {
	return models.AdMessage{}, nil
}
