package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/discordinteraction"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

func (c client) CreateDiscordInteraction(ctx context.Context, data models.DiscordInteraction) error {
	user, err := c.db.User.Get(ctx, data.UserID)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}

	return c.withTx(ctx, func(tx *db.Tx) error {
		return tx.DiscordInteraction.Create().
			SetCommand(data.Command).
			SetLocale(data.Locale.String()).
			SetOptions(data.Options).
			SetReferenceID(data.ReferenceID).
			SetType(data.Type).
			SetUser(user).Exec(ctx)
	})
}

func toDiscordInteraction(record *db.DiscordInteraction) models.DiscordInteraction {
	locale, err := language.Parse(record.Locale)
	if err != nil {
		locale = language.English
	}
	return models.DiscordInteraction{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,

		UserID:      record.UserID,
		Command:     record.Command,
		ReferenceID: record.ReferenceID,

		Type:   record.Type,
		Locale: locale,

		Options: record.Options,
	}
}

func (c client) GetDiscordInteraction(ctx context.Context, referenceID string) (models.DiscordInteraction, error) {
	interaction, err := c.db.DiscordInteraction.Query().Where(discordinteraction.ReferenceID(referenceID)).First(ctx)
	if err != nil {
		return models.DiscordInteraction{}, err
	}

	return toDiscordInteraction(interaction), nil
}
