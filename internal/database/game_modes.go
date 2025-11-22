package database

import (
	"context"
	"time"

	m "github.com/cufee/aftermath/internal/database/gen/public/model"
	t "github.com/cufee/aftermath/internal/database/gen/public/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/json"
	s "github.com/go-jet/jet/v2/postgres"
	"golang.org/x/text/language"
)

func (c *client) UpsertGameModes(ctx context.Context, modes map[string]map[language.Tag]string) error {
	if len(modes) < 1 {
		return nil
	}

	return c.withTx(ctx, func(tx *transaction) error {
		for id, locales := range modes {
			encoded, err := json.Marshal(locales)
			if err != nil {
				return err
			}

			model := m.GameMode{
				ID:             id,
				CreatedAt:      models.TimeToString(time.Now()),
				UpdatedAt:      models.TimeToString(time.Now()),
				LocalizedNames: encoded,
			}

			stmt := t.GameMode.
				INSERT(t.GameMode.AllColumns).
				MODEL(model).
				ON_CONFLICT(t.GameMode.ID).
				DO_UPDATE(s.SET(
					t.GameMode.UpdatedAt.SET(t.GameMode.EXCLUDED.UpdatedAt),
					t.GameMode.LocalizedNames.SET(t.GameMode.EXCLUDED.LocalizedNames),
				))
			_, err = tx.exec(ctx, stmt)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (c *client) GetGameModeNames(ctx context.Context, id string) (map[language.Tag]string, error) {
	var record m.GameMode
	err := c.query(ctx, t.GameMode.SELECT(t.GameMode.AllColumns).WHERE(t.GameMode.ID.EQ(s.String(id))), &record)
	if err != nil {
		return nil, err
	}

	var names map[language.Tag]string
	err = json.Unmarshal([]byte(record.LocalizedNames), &names)
	if err != nil {
		return nil, err
	}

	return names, nil
}
