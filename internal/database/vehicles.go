package database

import (
	"context"

	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/vehicle"
	"github.com/cufee/aftermath/internal/database/models"
)

func toVehicle(record *db.Vehicle) models.Vehicle {
	return models.Vehicle{
		ID:             record.ID,
		Tier:           record.Tier,
		LocalizedNames: record.LocalizedNames,
	}
}

func (c *client) UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) (map[string]error, error) {
	if len(vehicles) < 1 {
		return nil, nil
	}

	var ids []string
	for id := range vehicles {
		ids = append(ids, id)
	}

	records, err := c.db.Vehicle.Query().Where(vehicle.IDIn(ids...)).All(ctx)
	if err != nil && !IsNotFound(err) {
		return nil, err
	}

	errors := make(map[string]error)
	return errors, c.withTx(ctx, func(tx *db.Tx) error {
		for _, r := range records {
			v, ok := vehicles[r.ID]
			if !ok {
				continue
			}

			err := c.db.Vehicle.UpdateOneID(v.ID).
				SetTier(v.Tier).
				SetLocalizedNames(v.LocalizedNames).
				Exec(ctx)
			if err != nil {
				errors[v.ID] = err
			}

			delete(vehicles, v.ID)
		}

		var writes []*db.VehicleCreate
		for id, v := range vehicles {
			writes = append(writes, c.db.Vehicle.Create().
				SetID(id).
				SetTier(v.Tier).
				SetLocalizedNames(v.LocalizedNames),
			)
		}

		return c.db.Vehicle.CreateBulk(writes...).Exec(ctx)
	})
}

func (c *client) GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	records, err := c.db.Vehicle.Query().Where(vehicle.IDIn(ids...)).All(ctx)
	if err != nil {
		return nil, err
	}

	vehicles := make(map[string]models.Vehicle)
	for _, r := range records {
		vehicles[r.ID] = toVehicle(r)
	}

	return vehicles, nil
}
