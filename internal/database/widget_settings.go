package database

import (
	"context"

	m "github.com/cufee/aftermath/internal/database/gen/public/model"
	t "github.com/cufee/aftermath/internal/database/gen/public/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/postgres"
)

func (c *client) GetWidgetSettings(ctx context.Context, settingsID string) (models.WidgetOptions, error) {
	stmt := t.WidgetSettings.
		SELECT(t.WidgetSettings.AllColumns).
		WHERE(t.WidgetSettings.ID.EQ(s.String(settingsID)))

	var record m.WidgetSettings
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return models.WidgetOptions{}, err
	}
	return models.ToWidgetOptions(&record), nil
}

func (c *client) GetUserWidgetSettings(ctx context.Context, userID string, referenceID []string) ([]models.WidgetOptions, error) {
	where := []s.BoolExpression{t.WidgetSettings.UserID.EQ(s.String(userID))}
	if referenceID != nil {
		where = append(where, t.WidgetSettings.ReferenceID.IN(stringsToExp(referenceID)...))
	}

	stmt := t.WidgetSettings.
		SELECT(t.WidgetSettings.AllColumns).
		WHERE(s.AND(where...))

	var record []m.WidgetSettings
	err := c.query(ctx, stmt, &record)
	if err != nil {
		return nil, err
	}

	var settings []models.WidgetOptions
	for _, r := range record {
		settings = append(settings, models.ToWidgetOptions(&r))
	}

	return settings, nil
}

func (c *client) CreateWidgetSettings(ctx context.Context, userID string, settings models.WidgetOptions) (models.WidgetOptions, error) {
	settings.UserID = userID
	model := settings.Model()

	stmt := t.WidgetSettings.
		INSERT(t.WidgetSettings.AllColumns).
		MODEL(model).
		RETURNING(t.WidgetSettings.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.WidgetOptions{}, err
	}

	return models.ToWidgetOptions(&model), nil
}

func (c *client) UpdateWidgetSettings(ctx context.Context, id string, settings models.WidgetOptions) (models.WidgetOptions, error) {
	model := settings.Model()

	stmt := t.WidgetSettings.
		UPDATE(
			t.WidgetSettings.UpdatedAt,
			t.WidgetSettings.ReferenceID,
			t.WidgetSettings.Title,
			t.WidgetSettings.SessionFrom,
			t.WidgetSettings.Metadata,
			t.WidgetSettings.Styles,
			t.WidgetSettings.UserID,
			t.WidgetSettings.SessionReferenceID,
		).
		MODEL(model).
		WHERE(t.WidgetSettings.ID.EQ(s.String(id))).
		RETURNING(t.WidgetSettings.AllColumns)

	err := c.query(ctx, stmt, &model)
	if err != nil {
		return models.WidgetOptions{}, err
	}

	return models.ToWidgetOptions(&model), nil
}
