package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath-assets/types"
	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/json"
	s "github.com/go-jet/jet/v2/sqlite"
)

func toMap(record *m.GameMap) types.Map {
	data := types.Map{
		ID:              record.ID,
		Key:             record.ID,
		SupremacyPoints: int(record.SupremacyPoints),
	}
	json.Unmarshal([]byte(record.GameModes), &data.GameModes)
	json.Unmarshal([]byte(record.LocalizedNames), &data.LocalizedNames)
	return data
}

func (c *client) UpsertMaps(ctx context.Context, maps map[string]types.Map) error {
	if len(maps) < 1 {
		return nil
	}

	return c.withTx(ctx, func(tx *transaction) error {
		for id, data := range maps {
			model := m.GameMap{
				ID:              id,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
				SupremacyPoints: int32(data.SupremacyPoints),
			}
			modes, err := json.Marshal(data.GameModes)
			if err != nil {
				return err
			}
			model.GameModes = modes
			names, err := json.Marshal(data.LocalizedNames)
			if err != nil {
				return err
			}
			model.LocalizedNames = names

			stmt := t.GameMap.
				INSERT(t.GameMap.AllColumns).
				MODEL(model).
				ON_CONFLICT(t.GameMap.ID).
				DO_UPDATE(s.SET(
					t.GameMap.UpdatedAt.SET(t.GameMap.EXCLUDED.UpdatedAt),
					t.GameMap.GameModes.SET(t.GameMap.EXCLUDED.GameModes),
					t.GameMap.LocalizedNames.SET(t.GameMap.EXCLUDED.LocalizedNames),
					t.GameMap.SupremacyPoints.SET(t.GameMap.EXCLUDED.SupremacyPoints),
				))

			_, err = tx.exec(ctx, stmt)
			if err != nil {
				return err
			}
		}
		return nil
	})

}

func (c *client) GetMap(ctx context.Context, id string) (types.Map, error) {
	var record m.GameMap
	err := c.query(ctx, t.GameMap.SELECT(t.GameMap.AllColumns).WHERE(t.GameMap.ID.EQ(s.String(id))), &record)
	if err != nil {
		return types.Map{}, err
	}
	return toMap(&record), nil
}
