package database

import (
	"context"

	"time"

	m "github.com/cufee/aftermath/internal/database/gen/public/model"
	t "github.com/cufee/aftermath/internal/database/gen/public/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/json"
	"github.com/cufee/aftermath/internal/stats/frame"
	s "github.com/go-jet/jet/v2/postgres"
)

func (c *client) UpsertVehicleAverages(ctx context.Context, averages map[string]frame.StatsFrame) (map[string]error, error) {
	if len(averages) < 1 {
		return nil, nil
	}

	errors := make(map[string]error)
	return errors, c.withTx(ctx, func(tx *transaction) error {
		for id, data := range averages {
			encoded, err := json.Marshal(data)
			if err != nil {
				errors[id] = err
				continue
			}

			model := m.VehicleAverage{
				ID:        id,
				CreatedAt: models.TimeToString(time.Now()),
				UpdatedAt: models.TimeToString(time.Now()),
				Data:      encoded,
			}

			stmt := t.VehicleAverage.
				INSERT(t.VehicleAverage.AllColumns).
				MODEL(model).
				ON_CONFLICT(t.VehicleAverage.ID).
				DO_UPDATE(s.SET(
					t.VehicleAverage.Data.SET(t.VehicleAverage.EXCLUDED.Data),
					t.VehicleAverage.UpdatedAt.SET(t.VehicleAverage.EXCLUDED.UpdatedAt),
				))
			_, err = tx.exec(ctx, stmt)
			if err != nil {
				errors[id] = err
			}
		}
		return nil
	})
}

func (c *client) GetVehicleAverages(ctx context.Context, ids []string) (map[string]frame.StatsFrame, error) {
	if len(ids) < 1 {
		return nil, nil
	}

	var records []m.VehicleAverage
	err := c.query(ctx,
		t.VehicleAverage.
			SELECT(t.VehicleAverage.AllColumns).
			WHERE(t.VehicleAverage.ID.IN(stringsToExp(ids)...)),
		&records)
	if err != nil {
		return nil, err
	}

	averages := make(map[string]frame.StatsFrame)
	for _, a := range records {
		var frame frame.StatsFrame
		err := json.Unmarshal([]byte(a.Data), &frame)
		if err != nil {
			continue
		}
		averages[a.ID] = frame
	}
	return averages, nil
}
