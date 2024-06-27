package database

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
	"github.com/cufee/aftermath/internal/database/models"
)

type getSnapshotQuery struct {
	vehicleIDs []string

	createdAfter  *time.Time
	createdBefore *time.Time
}

type SnapshotQuery func(*getSnapshotQuery)

func WithVehicleIDs(ids []string) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		q.vehicleIDs = ids
	}
}
func WithCreatedAfter(after time.Time) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		q.createdAfter = &after
	}
}
func WithCreatedBefore(before time.Time) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		q.createdBefore = &before
	}
}

func toVehicleSnapshot(record *db.VehicleSnapshot) models.VehicleSnapshot {
	return models.VehicleSnapshot{
		ID:             record.ID,
		Type:           record.Type,
		CreatedAt:      record.CreatedAt,
		LastBattleTime: record.LastBattleTime,
		ReferenceID:    record.ReferenceID,
		AccountID:      record.AccountID,
		VehicleID:      record.VehicleID,
		Stats:          record.Frame,
	}
}

func (c *client) CreateAccountVehicleSnapshots(ctx context.Context, accountID string, snapshots ...models.VehicleSnapshot) (map[string]error, error) {
	if len(snapshots) < 1 {
		return nil, nil
	}

	var errors = make(map[string]error)
	for _, data := range snapshots {
		// make a transaction per write to avoid locking for too long
		err := c.withTx(ctx, func(tx *db.Tx) error {
			err := c.db.VehicleSnapshot.Create().
				SetType(data.Type).
				SetFrame(data.Stats).
				SetVehicleID(data.VehicleID).
				SetReferenceID(data.ReferenceID).
				SetCreatedAt(data.CreatedAt).
				SetBattles(int(data.Stats.Battles.Float())).
				SetLastBattleTime(data.LastBattleTime).
				SetAccount(c.db.Account.GetX(ctx, accountID)).
				Exec(ctx)
			if err != nil {
				errors[data.VehicleID] = err
			}
			return nil
		})
		if err != nil {
			errors[data.VehicleID] = err
		}
	}

	if len(errors) > 0 {
		return errors, nil
	}
	return nil, nil
}

func (c *client) GetVehicleSnapshots(ctx context.Context, accountID, referenceID string, kind models.SnapshotType, options ...SnapshotQuery) ([]models.VehicleSnapshot, error) {
	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	// this query is impossible to build using the db.VehicleSnapshots methods, so we do it manually
	var innerWhere []*sql.Predicate
	innerWhere = append(innerWhere,
		sql.EQ(vehiclesnapshot.FieldType, kind),
		sql.EQ(vehiclesnapshot.FieldAccountID, accountID),
		sql.EQ(vehiclesnapshot.FieldReferenceID, referenceID),
	)
	innerOrder := sql.Desc(vehiclesnapshot.FieldCreatedAt)

	if query.createdAfter != nil {
		innerWhere = append(innerWhere, sql.GT(vehiclesnapshot.FieldCreatedAt, *query.createdAfter))
		innerOrder = sql.Asc(vehiclesnapshot.FieldCreatedAt)
	}
	if query.createdBefore != nil {
		innerWhere = append(innerWhere, sql.LT(vehiclesnapshot.FieldCreatedAt, *query.createdBefore))
		innerOrder = sql.Desc(vehiclesnapshot.FieldCreatedAt)
	}
	if len(query.vehicleIDs) > 0 {
		var ids []any
		for _, id := range query.vehicleIDs {
			ids = append(ids, id)
		}
		innerWhere = append(innerWhere, sql.In(vehiclesnapshot.FieldVehicleID, ids...))
	}

	innerQuery := sql.Select(vehiclesnapshot.Columns...).From(sql.Table(vehiclesnapshot.Table))
	innerQuery = innerQuery.Where(sql.And(innerWhere...))
	innerQuery = innerQuery.OrderBy(innerOrder)

	innerQueryString, innerQueryArgs := innerQuery.Query()

	// wrap the inner query in a GROUP BY
	wrapper := &sql.Builder{}
	wrapped := wrapper.Wrap(func(b *sql.Builder) { b.WriteString(innerQueryString) })
	queryString, _ := sql.Select("*").FromExpr(wrapped).GroupBy(vehiclesnapshot.FieldVehicleID).Query()

	rows, err := c.db.VehicleSnapshot.QueryContext(ctx, queryString, innerQueryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snapshots []models.VehicleSnapshot
	for rows.Next() {
		var record db.VehicleSnapshot

		values, err := record.ScanValues(vehiclesnapshot.Columns)
		if err != nil {
			return nil, err
		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		err = record.AssignValues(vehiclesnapshot.Columns, values)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, toVehicleSnapshot(&record))
	}

	return snapshots, nil
}

func toAccountSnapshot(record *db.AccountSnapshot) models.AccountSnapshot {
	return models.AccountSnapshot{
		ID:             record.ID,
		Type:           record.Type,
		AccountID:      record.AccountID,
		ReferenceID:    record.ReferenceID,
		RatingBattles:  record.RatingFrame,
		RegularBattles: record.RegularFrame,
		CreatedAt:      record.CreatedAt,
		LastBattleTime: record.LastBattleTime,
	}
}

func (c *client) CreateAccountSnapshots(ctx context.Context, snapshots ...models.AccountSnapshot) (map[string]error, error) {
	if len(snapshots) < 1 {
		return nil, nil
	}

	var errors = make(map[string]error)
	for _, snapshot := range snapshots {
		// make a transaction per write to avoid locking for too long
		err := c.withTx(ctx, func(tx *db.Tx) error {
			for _, s := range snapshots {
				err := c.db.AccountSnapshot.Create().
					SetAccount(c.db.Account.GetX(ctx, s.AccountID)). // we assume the account exists here
					SetCreatedAt(s.CreatedAt).
					SetLastBattleTime(s.LastBattleTime).
					SetRatingBattles(int(s.RatingBattles.Battles.Float())).
					SetRatingFrame(s.RatingBattles).
					SetReferenceID(s.ReferenceID).
					SetRegularBattles(int(s.RegularBattles.Battles)).
					SetRegularFrame(s.RegularBattles).
					SetType(s.Type).Exec(ctx)
				if err != nil {
					errors[s.AccountID] = err
				}
			}
			return nil
		})
		if err != nil {
			errors[snapshot.AccountID] = err
		}
	}
	if len(errors) > 0 {
		return errors, nil
	}
	return nil, nil
}

func (c *client) GetLastAccountSnapshots(ctx context.Context, accountID string, limit int) ([]models.AccountSnapshot, error) {
	records, err := c.db.AccountSnapshot.Query().Where(accountsnapshot.AccountID(accountID)).Order(accountsnapshot.ByCreatedAt(sql.OrderDesc())).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}

	var snapshots []models.AccountSnapshot
	for _, r := range records {
		snapshots = append(snapshots, toAccountSnapshot(r))
	}

	return snapshots, nil
}

func (c *client) GetAccountSnapshot(ctx context.Context, accountID, referenceID string, kind models.SnapshotType, options ...SnapshotQuery) (models.AccountSnapshot, error) {
	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	var where []predicate.AccountSnapshot
	order := accountsnapshot.ByCreatedAt(sql.OrderDesc())
	where = append(where, accountsnapshot.AccountID(accountID), accountsnapshot.ReferenceID(referenceID), accountsnapshot.TypeEQ(kind))
	if query.createdAfter != nil {
		order = accountsnapshot.ByCreatedAt(sql.OrderAsc())
		where = append(where, accountsnapshot.CreatedAtGT(*query.createdAfter))
	}
	if query.createdBefore != nil {
		where = append(where, accountsnapshot.CreatedAtLT(*query.createdBefore))
	}

	record, err := c.db.AccountSnapshot.Query().Where(where...).Order(order).First(ctx)
	if err != nil {
		return models.AccountSnapshot{}, err
	}

	return toAccountSnapshot(record), nil
}

func (c *client) GetManyAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...SnapshotQuery) ([]models.AccountSnapshot, error) {
	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	var where []predicate.AccountSnapshot
	where = append(where, accountsnapshot.AccountIDIn(accountIDs...), accountsnapshot.TypeEQ(kind))
	if query.createdAfter != nil {
		where = append(where, accountsnapshot.CreatedAtGT(*query.createdAfter))
	}
	if query.createdBefore != nil {
		where = append(where, accountsnapshot.CreatedAtLT(*query.createdAfter))
	}

	records, err := c.db.AccountSnapshot.Query().Where(where...).All(ctx)
	if err != nil {
		return nil, err
	}

	var snapshots []models.AccountSnapshot
	for _, r := range records {
		snapshots = append(snapshots, toAccountSnapshot(r))
	}

	return snapshots, nil

}
