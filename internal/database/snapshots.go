package database

import (
	"context"
	"time"

	m "github.com/cufee/aftermath/internal/database/gen/model"
	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/utils"
	s "github.com/go-jet/jet/v2/sqlite"
	"github.com/pkg/errors"
)

// --- vehicle snapshots ---

/*
Get complete vehicle snapshots fot each vehicle ID in vehicle IDs for a specific account
  - passing nil vehicleIDs will return all vehicles available
*/
func (c *client) GetVehicleSnapshots(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...Query) ([]models.VehicleSnapshot, error) {
	var query baseQueryOptions
	for _, apply := range options {
		apply(&query)
	}

	stmt := vehiclesQuery(accountID, vehicleIDs, kind, t.VehicleSnapshot.VehicleID, query)
	rs, err := c.rows(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var snapshots []models.VehicleSnapshot
	for rs.Next() {
		var record m.VehicleSnapshot
		if err := rs.Scan(&record); err != nil {
			return nil, errors.New("failed to scan row: " + err.Error())
		}
		snapshots = append(snapshots, models.ToVehicleSnapshot(&record))
	}

	return snapshots, nil
}

// CreateVehicleSnapshots inserts many vehicle snapshots
func (c *client) CreateVehicleSnapshots(ctx context.Context, snapshots ...*models.VehicleSnapshot) error {
	if len(snapshots) < 1 {
		return nil
	}

	for _, batch := range utils.Batch(snapshots, 100) {
		var models []m.VehicleSnapshot
		for _, item := range batch {
			models = append(models, item.Model())
		}

		err := c.withTx(ctx, func(tx *transaction) error {
			_, err := tx.exec(ctx, t.VehicleSnapshot.INSERT(t.VehicleSnapshot.AllColumns).MODELS(models))
			return err
		})

		if err != nil {
			return err
		}
	}
	return nil
}

// --- account snapshots ---

/*
GetAccountSnapshots returns complete snapshots for accounts by ID, grouped by account ID
  - there are no use cases where all accounts should be returned for now, so nil slice of ids will return an error
*/
func (c *client) GetAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...Query) ([]models.AccountSnapshot, error) {
	if len(accountIDs) < 1 {
		return nil, ErrNotFound
	}

	var query baseQueryOptions
	for _, apply := range options {
		apply(&query)
	}

	stmt := accountsQuery(accountIDs, kind, t.AccountSnapshot.AccountID, query)
	rs, err := c.rows(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var snapshots []models.AccountSnapshot
	for rs.Next() {
		var snapshot m.AccountSnapshot
		err := rs.Scan(&snapshot)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, models.ToAccountSnapshot(&snapshot))
	}
	return snapshots, nil
}

/*
Get last battle times for accounts by ID, grouped by account ID
*/
func (c *client) GetAccountLastBattleTimes(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...Query) (map[string]time.Time, error) {
	if len(accountIDs) < 1 {
		return nil, ErrNotFound
	}

	var query baseQueryOptions
	for _, apply := range options {
		apply(&query)
	}
	withSelect(s.ColumnList{t.AccountSnapshot.AccountID, t.AccountSnapshot.LastBattleTime})(&query)

	stmt := accountsQuery(accountIDs, kind, t.AccountSnapshot.AccountID, query)
	rs, err := c.rows(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var snapshots []m.AccountSnapshot
	for rs.Next() {
		var snapshot m.AccountSnapshot
		err := rs.Scan(&snapshot)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, snapshot)
	}

	var lastBattles = make(map[string]time.Time)
	for _, r := range snapshots {
		lastBattles[r.AccountID] = models.StringToTime(r.LastBattleTime)
	}
	return lastBattles, nil
}

func (c *client) CreateAccountSnapshots(ctx context.Context, snapshots ...*models.AccountSnapshot) error {
	if len(snapshots) < 1 {
		return nil
	}

	for _, batch := range utils.Batch(snapshots, 100) {
		var models []m.AccountSnapshot
		for _, item := range batch {
			models = append(models, item.Model())
		}

		err := c.withTx(ctx, func(tx *transaction) error {
			_, err := tx.exec(ctx, t.AccountSnapshot.INSERT(t.AccountSnapshot.AllColumns).MODELS(models))
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}
