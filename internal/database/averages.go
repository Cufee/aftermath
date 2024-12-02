package database

// import (
// 	"context"
// 	"time"

// 	m "github.com/cufee/aftermath/internal/database/gen/model"
// 	t "github.com/cufee/aftermath/internal/database/gen/table"
// 	"github.com/cufee/aftermath/internal/database/models"
// 	s "github.com/go-jet/jet/v2/sqlite"
// )

// func (c *client) UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) (map[string]error, error) {
// 	if len(averages) < 1 {
// 		return nil, nil
// 	}

// 	var ids []string
// 	for id := range averages {
// 		ids = append(ids, id)
// 	}

// 	existing, err := c.db.VehicleAverage.Query().Where(vehicleaverage.IDIn(ids...)).All(ctx)
// 	if err != nil && !IsNotFound(err) {
// 		return nil, err
// 	}

// 	errors := make(map[string]error)
// 	return errors, c.withTx(ctx, func(tx *db.Tx) error {
// 		for _, r := range existing {
// 			update, ok := averages[r.ID]
// 			if !ok {
// 				continue // should never happen tho
// 			}

// 			err := tx.VehicleAverage.UpdateOneID(r.ID).SetData(update).Exec(ctx)
// 			if err != nil {
// 				errors[r.ID] = err
// 			}

// 			delete(averages, r.ID)
// 		}

// 		var writes []*db.VehicleAverageCreate
// 		for id, frame := range averages {
// 			writes = append(writes, tx.VehicleAverage.Create().
// 				SetID(id).
// 				SetData(frame),
// 			)
// 		}

// 		return tx.VehicleAverage.CreateBulk(writes...).Exec(ctx)
// 	})
// }

// func (c *client) GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error) {
// 	if len(ids) < 1 {
// 		return nil, nil
// 	}

// 	records, err := c.db.VehicleAverage.Query().Where(vehicleaverage.IDIn(ids...)).All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	averages := make(map[string]frame.StatsFrame)
// 	for _, a := range records {
// 		averages[a.ID] = a.Data
// 	}
// 	return averages, nil
// }
