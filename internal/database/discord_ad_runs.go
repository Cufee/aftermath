package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/database/models"

	m "github.com/cufee/aftermath/internal/database/gen/public/model"
	t "github.com/cufee/aftermath/internal/database/gen/public/table"
	s "github.com/go-jet/jet/v2/postgres"
)

func (c client) CreateDiscordAdRun(ctx context.Context, data models.DiscordAdRun) (models.DiscordAdRun, error) {
	model := data.Model()
	stmt := t.DiscordAdRun.
		INSERT(t.DiscordAdRun.AllColumns.Except(t.DiscordAdRun.ID)).
		MODEL(model).
		RETURNING(t.DiscordAdRun.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.DiscordAdRun{}, err
	}
	return models.ToDiscordAdRun(&model), nil
}

func (c client) GetChannelLastAdRun(ctx context.Context, channelID string) (time.Time, error) {
	stmt := t.DiscordAdRun.
		SELECT(t.DiscordAdRun.CreatedAt).
		ORDER_BY(t.DiscordAdRun.CreatedAt.DESC()).
		WHERE(t.DiscordAdRun.ChannelID.EQ(s.String(channelID)))

	var result m.DiscordAdRun
	err := c.query(ctx, stmt, &result)
	if err != nil && !IsNotFound(err) {
		return time.Time{}, err
	}

	return models.StringToTime(result.CreatedAt), nil
}

func (c client) UpdateDiscordAdRun(ctx context.Context, data models.DiscordAdRun) (models.DiscordAdRun, error) {
	model := data.Model()
	stmt := t.DiscordAdRun.
		UPDATE(t.DiscordAdRun.AllColumns.Except(t.DiscordAdRun.ID, t.DiscordAdRun.CreatedAt)).
		WHERE(t.DiscordAdRun.ID.EQ(s.Int(model.ID))).
		MODEL(model).
		RETURNING(t.DiscordAdRun.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.DiscordAdRun{}, err
	}

	return models.ToDiscordAdRun(&model), nil
}
