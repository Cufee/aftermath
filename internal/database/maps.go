package database

import (
	"context"

	"github.com/cufee/aftermath-assets/types"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/gamemap"
)

func toMap(record *db.GameMap) types.Map {
	return types.Map{
		ID:              record.ID,
		GameModes:       record.GameModes,
		SupremacyPoints: record.SupremacyPoints,
		LocalizedNames:  record.LocalizedNames,
	}
}

func (c *client) UpsertMaps(ctx context.Context, maps map[string]types.Map) error {
	if len(maps) < 1 {
		return nil
	}

	var ids []string
	for id := range maps {
		ids = append(ids, id)
	}

	records, err := c.db.GameMap.Query().Where(gamemap.IDIn(ids...)).All(ctx)
	if err != nil && !IsNotFound(err) {
		return err
	}

	return c.withTx(ctx, func(tx *db.Tx) error {
		for _, r := range records {
			m, ok := maps[r.ID]
			if !ok {
				continue
			}

			err := tx.GameMap.UpdateOneID(m.ID).
				SetGameModes(m.GameModes).
				SetLocalizedNames(m.LocalizedNames).
				SetSupremacyPoints(m.SupremacyPoints).
				Exec(ctx)

			if err != nil {
				return err
			}
			delete(maps, m.ID)
		}

		var writes []*db.GameMapCreate
		for id, m := range maps {
			writes = append(writes, tx.GameMap.Create().
				SetID(id).
				SetGameModes(m.GameModes).
				SetLocalizedNames(m.LocalizedNames).
				SetSupremacyPoints(m.SupremacyPoints),
			)
		}
		return tx.GameMap.CreateBulk(writes...).Exec(ctx)
	})
}

func (c *client) GetMap(ctx context.Context, id string) (types.Map, error) {
	record, err := c.db.GameMap.Get(ctx, id)
	if err != nil {
		return types.Map{}, err
	}

	return toMap(record), nil
}
