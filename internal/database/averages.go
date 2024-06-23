package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/vehicleaverage"
	"github.com/cufee/aftermath/internal/stats/frame"
)

func (c *libsqlClient) UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) (map[string]error, error) {
	if len(averages) < 1 {
		return nil, nil
	}

	var ids []string
	for id := range averages {
		ids = append(ids, id)
	}

	tx, err := c.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	existing, err := tx.VehicleAverage.Query().Where(vehicleaverage.IDIn(ids...)).All(ctx)
	if err != nil && !IsNotFound(err) {
		return nil, rollback(tx, err)
	}

	errors := make(map[string]error)
	for _, r := range existing {
		update, ok := averages[r.ID]
		if !ok {
			continue // should never happen tho
		}

		err := tx.VehicleAverage.UpdateOneID(r.ID).SetData(update).Exec(ctx)
		if err != nil {
			errors[r.ID] = err
		}

		delete(averages, r.ID)
	}

	var inserts []*db.VehicleAverageCreate
	for id, frame := range averages {
		inserts = append(inserts,
			c.db.VehicleAverage.Create().
				SetID(id).
				SetData(frame),
		)
	}

	err = tx.VehicleAverage.CreateBulk(inserts...).Exec(ctx)
	if err != nil {
		return nil, rollback(tx, err)
	}

	return nil, tx.Commit()
}

func (c *libsqlClient) GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	records, err := c.db.VehicleAverage.Query().Where(vehicleaverage.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}

	averages := make(map[string]frame.StatsFrame)
	for _, a := range records {
		averages[a.ID] = a.Data
	}
	return averages, nil
}
