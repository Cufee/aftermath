package database

import (
	"context"
	"time"

	"github.com/cufee/aftermath/internal/json"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/sqlite"
)

func (c *client) UpsertVehicles(ctx context.Context, vehicles map[string]models.Vehicle) (map[string]error, error) {
	if len(vehicles) < 1 {
		return nil, nil
	}

	errors := make(map[string]error, len(vehicles))
	return errors, c.withTx(ctx, func(tx *transaction) error {
		for id, data := range vehicles {
			model := m.Vehicle{
				ID:        id,
				CreatedAt: models.TimeToString(time.Now()),
				UpdatedAt: models.TimeToString(time.Now()),
				Tier:      int32(data.Tier),
			}
			names, err := json.Marshal(data.LocalizedNames)
			if err != nil {
				errors[id] = err
				continue
			}
			model.LocalizedNames = names

			stmt := t.Vehicle.
				INSERT(t.Vehicle.AllColumns).
				MODEL(model).
				ON_CONFLICT(t.Vehicle.ID).
				DO_UPDATE(s.SET(
					t.Vehicle.LocalizedNames.SET(t.Vehicle.EXCLUDED.LocalizedNames),
					t.Vehicle.UpdatedAt.SET(t.Vehicle.EXCLUDED.UpdatedAt),
					t.Vehicle.Tier.SET(t.Vehicle.EXCLUDED.Tier),
				))

			_, errors[id] = tx.exec(ctx, stmt)
		}
		return nil
	})
}

func (c *client) GetVehicles(ctx context.Context, ids []string) (map[string]models.Vehicle, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	var records []m.Vehicle
	err := c.query(ctx, t.Vehicle.SELECT(t.Vehicle.AllColumns).WHERE(t.Vehicle.ID.IN(stringsToExp(ids)...)), &records)
	if err != nil {
		return nil, err
	}

	vehicles := make(map[string]models.Vehicle)
	for _, r := range records {
		vehicles[r.ID] = models.ToVehicle(&r)
	}

	return vehicles, nil
}

func (c *client) GetAllVehicles(ctx context.Context) (map[string]models.Vehicle, error) {
	var records []m.Vehicle
	err := c.query(ctx, t.Vehicle.SELECT(t.Vehicle.AllColumns), &records)
	if err != nil {
		return nil, err
	}

	vehicles := make(map[string]models.Vehicle)
	for _, r := range records {
		vehicles[r.ID] = models.ToVehicle(&r)
	}

	return vehicles, nil
}
