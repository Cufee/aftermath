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

func (c *libsqlClient) UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) error {
	if len(vehicles) < 1 {
		return nil
	}

	var ids []string
	for id := range vehicles {
		ids = append(ids, id)
	}

	tx, err := c.db.Tx(ctx)
	if err != nil {
		return err
	}

	records, err := tx.Vehicle.Query().Where(vehicle.IDIn(ids...)).All(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	for _, r := range records {
		v, ok := vehicles[r.ID]
		if !ok {
			continue
		}

		err := tx.Vehicle.UpdateOneID(v.ID).
			SetTier(v.Tier).
			SetLocalizedNames(v.LocalizedNames).
			Exec(ctx)
		if err != nil {
			return rollback(tx, err)
		}

		delete(vehicles, v.ID)
	}

	var inserts []*db.VehicleCreate
	for id, v := range vehicles {
		inserts = append(inserts,
			c.db.Vehicle.Create().
				SetID(id).
				SetTier(v.Tier).
				SetLocalizedNames(v.LocalizedNames),
		)
	}

	err = tx.Vehicle.CreateBulk(inserts...).Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}

	return tx.Commit()
}

func (c *libsqlClient) GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	records, err := c.db.Vehicle.Query().Where(vehicle.IDIn(ids...)).All(ctx)
	if err != nil && !IsNotFound(err) {
		return nil, err
	}

	vehicles := make(map[string]models.Vehicle)
	for _, r := range records {
		vehicles[r.ID] = toVehicle(r)
	}

	return vehicles, nil
}
