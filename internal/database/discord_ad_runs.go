package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/models"
	"github.com/pkg/errors"
)

func (c client) CreateDiscordAdRun(ctx context.Context, data models.DiscordAdRun) (models.DiscordAdRun, error) {
	//
	return models.DiscordAdRun{}, errors.New("not implemented")
}

func (c client) GetChannelLastAdRun(ctx context.Context, channelID string) (time.Time, error) {
	//
	return time.Now(), nil
}
