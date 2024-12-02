package database

// import (
// 	"context"

// 	"github.com/cufee/aftermath/internal/database/ent/db"
// 	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
// 	"github.com/cufee/aftermath/internal/database/ent/db/widgetsettings"
// 	"github.com/cufee/aftermath/internal/database/models"
// )

// func toWidgetOptions(record *db.WidgetSettings) models.WidgetOptions {
// 	return models.WidgetOptions{
// 		ID:        record.ID,
// 		CreatedAt: record.CreatedAt,
// 		UpdatedAt: record.UpdatedAt,

// 		Title:        record.Title,
// 		UserID:       record.UserID,
// 		AccountID:    record.ReferenceID,
// 		SessionFrom:  record.SessionFrom,
// 		SessionRefID: record.SessionReferenceID,

// 		Style: record.Styles,
// 		Meta:  record.Metadata,
// 	}
// }

// func (c *client) GetWidgetSettings(ctx context.Context, settingsID string) (models.WidgetOptions, error) {
// 	record, err := c.db.WidgetSettings.Get(ctx, settingsID)
// 	if err != nil {
// 		return models.WidgetOptions{}, err
// 	}
// 	return toWidgetOptions(record), nil
// }

// func (c *client) GetUserWidgetSettings(ctx context.Context, userID string, referenceID []string) ([]models.WidgetOptions, error) {
// 	var where = []predicate.WidgetSettings{widgetsettings.UserID(userID)}
// 	if referenceID != nil {
// 		where = append(where, widgetsettings.ReferenceIDIn(referenceID...))
// 	}

// 	records, err := c.db.WidgetSettings.Query().Where(where...).All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var options []models.WidgetOptions
// 	for _, r := range records {
// 		options = append(options, toWidgetOptions(r))
// 	}
// 	return options, nil
// }

// func (c *client) CreateWidgetSettings(ctx context.Context, userID string, settings models.WidgetOptions) (models.WidgetOptions, error) {
// 	user, err := c.db.User.Get(ctx, userID)
// 	if err != nil {
// 		return models.WidgetOptions{}, err
// 	}

// 	created, err := c.db.WidgetSettings.Create().
// 		SetTitle(settings.Title).
// 		SetMetadata(settings.Meta).
// 		SetReferenceID(settings.AccountID).
// 		SetSessionFrom(settings.SessionFrom).
// 		SetSessionReferenceID(settings.SessionRefID).
// 		SetStyles(settings.Style).
// 		SetUser(user).
// 		Save(ctx)
// 	if err != nil {
// 		return models.WidgetOptions{}, err
// 	}

// 	return toWidgetOptions(created), nil
// }

// func (c *client) UpdateWidgetSettings(ctx context.Context, id string, settings models.WidgetOptions) (models.WidgetOptions, error) {
// 	updated, err := c.db.WidgetSettings.UpdateOneID(id).
// 		SetTitle(settings.Title).
// 		SetMetadata(settings.Meta).
// 		SetReferenceID(settings.AccountID).
// 		SetSessionFrom(settings.SessionFrom).
// 		SetSessionReferenceID(settings.SessionRefID).
// 		SetStyles(settings.Style).
// 		Save(ctx)
// 	if err != nil {
// 		return models.WidgetOptions{}, err
// 	}

// 	return toWidgetOptions(updated), nil
// }
