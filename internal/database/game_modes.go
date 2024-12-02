package database

// import (
// 	"context"

// 	"github.com/cufee/aftermath/internal/database/ent/db"
// 	"github.com/cufee/aftermath/internal/database/ent/db/gamemode"
// 	"golang.org/x/text/language"
// )

// func (c *client) UpsertGameModes(ctx context.Context, modes map[string]map[language.Tag]string) (map[string]error, error) {
// 	if len(modes) < 1 {
// 		return nil, nil
// 	}

// 	var ids []string
// 	for id := range modes {
// 		ids = append(ids, id)
// 	}

// 	records, err := c.db.GameMode.Query().Where(gamemode.IDIn(ids...)).All(ctx)
// 	if err != nil && !IsNotFound(err) {
// 		return nil, err
// 	}

// 	errors := make(map[string]error)
// 	return errors, c.withTx(ctx, func(tx *db.Tx) error {
// 		for _, r := range records {
// 			m, ok := modes[r.ID]
// 			if !ok {
// 				continue
// 			}

// 			err := tx.GameMode.UpdateOneID(r.ID).
// 				SetLocalizedNames(m).
// 				Exec(ctx)
// 			if err != nil {
// 				errors[r.ID] = err
// 			}

// 			delete(modes, r.ID)
// 		}

// 		var writes []*db.GameModeCreate
// 		for id, v := range modes {
// 			writes = append(writes, tx.GameMode.Create().
// 				SetID(id).
// 				SetLocalizedNames(v),
// 			)
// 		}

// 		return tx.GameMode.CreateBulk(writes...).Exec(ctx)
// 	})
// }

// func (c *client) GetGameModeNames(ctx context.Context, id string) (map[language.Tag]string, error) {
// 	record, err := c.db.GameMode.Get(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return record.LocalizedNames, nil
// }
