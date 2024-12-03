package database

import (
	"context"

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

// func (c *client) UpsertMaps(ctx context.Context, maps map[string]types.Map) error {
// 	if len(maps) < 1 {
// 		return nil
// 	}

// 	var ids []string
// 	for id := range maps {
// 		ids = append(ids, id)
// 	}

// 	records, err := c.db.GameMap.Query().Where(gamemap.IDIn(ids...)).All(ctx)
// 	if err != nil && !IsNotFound(err) {
// 		return err
// 	}

// 	return c.withTx(ctx, func(tx *db.Tx) error {
// 		for _, r := range records {
// 			m, ok := maps[r.ID]
// 			if !ok {
// 				continue
// 			}

// 			err := tx.GameMap.UpdateOneID(m.ID).
// 				SetGameModes(m.GameModes).
// 				SetLocalizedNames(m.LocalizedNames).
// 				SetSupremacyPoints(m.SupremacyPoints).
// 				Exec(ctx)

// 			if err != nil {
// 				return err
// 			}
// 			delete(maps, m.ID)
// 		}

// 		var writes []*db.GameMapCreate
// 		for id, m := range maps {
// 			writes = append(writes, tx.GameMap.Create().
// 				SetID(id).
// 				SetGameModes(m.GameModes).
// 				SetLocalizedNames(m.LocalizedNames).
// 				SetSupremacyPoints(m.SupremacyPoints),
// 			)
// 		}
// 		return tx.GameMap.CreateBulk(writes...).Exec(ctx)
// 	})
// }

func (c *client) GetMap(ctx context.Context, id string) (types.Map, error) {
	var record m.GameMap
	err := c.query(ctx, t.GameMap.SELECT(t.GameMap.AllColumns).WHERE(t.GameMap.ID.EQ(s.String(id))), &record)
	if err != nil {
		return types.Map{}, err
	}
	return toMap(&record), nil
}
