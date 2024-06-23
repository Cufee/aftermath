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
		CreatedAt:      time.Unix(record.CreatedAt, 0),
		LastBattleTime: time.Unix(record.LastBattleTime, 0),
		ReferenceID:    record.ReferenceID,
		AccountID:      record.AccountID,
		VehicleID:      record.VehicleID,
		Stats:          record.Frame,
	}
}

func (c *client) CreateVehicleSnapshots(ctx context.Context, snapshots ...models.VehicleSnapshot) error {
	if len(snapshots) < 1 {
		return nil
	}

	var inserts []*db.VehicleSnapshotCreate
	for _, data := range snapshots {
		inserts = append(inserts,
			c.db.VehicleSnapshot.Create().
				SetType(data.Type).
				SetFrame(data.Stats).
				SetVehicleID(data.VehicleID).
				SetReferenceID(data.ReferenceID).
				SetBattles(int(data.Stats.Battles.Float())).
				SetLastBattleTime(data.LastBattleTime.Unix()).
				SetAccount(c.db.Account.GetX(ctx, data.AccountID)),
		)
	}

	return c.db.VehicleSnapshot.CreateBulk(inserts...).Exec(ctx)
}

func (c *client) GetVehicleSnapshots(ctx context.Context, accountID, referenceID string, kind models.SnapshotType, options ...SnapshotQuery) ([]models.VehicleSnapshot, error) {
	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	var where []predicate.VehicleSnapshot
	where = append(where, vehiclesnapshot.AccountID(accountID))
	where = append(where, vehiclesnapshot.ReferenceID(referenceID))
	where = append(where, vehiclesnapshot.TypeEQ(kind))

	order := vehiclesnapshot.ByCreatedAt(sql.OrderDesc())
	if query.createdAfter != nil {
		order = vehiclesnapshot.ByCreatedAt(sql.OrderAsc())
		where = append(where, vehiclesnapshot.CreatedAtGT(query.createdAfter.Unix()))
	}
	if query.createdBefore != nil {
		where = append(where, vehiclesnapshot.CreatedAtLT(query.createdBefore.Unix()))
	}
	if query.vehicleIDs != nil {
		where = append(where, vehiclesnapshot.VehicleIDIn(query.vehicleIDs...))
	}

	var records db.VehicleSnapshots
	err := c.db.VehicleSnapshot.Query().Where(where...).Order(order).GroupBy(vehiclesnapshot.FieldVehicleID).Aggregate(func(s *sql.Selector) string { return s.Select("*").String() }).Scan(ctx, &records)
	if err != nil {
		return nil, err
	}
	var snapshots []models.VehicleSnapshot
	for _, r := range records {
		snapshots = append(snapshots, toVehicleSnapshot(r))
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
		CreatedAt:      time.Unix(record.CreatedAt, 0),
		LastBattleTime: time.Unix(record.LastBattleTime, 0),
	}
}

func (c *client) CreateAccountSnapshots(ctx context.Context, snapshots ...models.AccountSnapshot) error {
	if len(snapshots) < 1 {
		return nil
	}

	var inserts []*db.AccountSnapshotCreate
	for _, s := range snapshots {
		inserts = append(inserts,
			c.db.AccountSnapshot.Create().
				SetAccount(c.db.Account.GetX(ctx, s.AccountID)). // we assume the account exists here
				SetCreatedAt(s.CreatedAt.Unix()).
				SetLastBattleTime(s.LastBattleTime.Unix()).
				SetRatingBattles(int(s.RatingBattles.Battles.Float())).
				SetRatingFrame(s.RatingBattles).
				SetReferenceID(s.ReferenceID).
				SetRegularBattles(int(s.RegularBattles.Battles)).
				SetRegularFrame(s.RegularBattles).
				SetType(s.Type),
		)
	}

	return c.db.AccountSnapshot.CreateBulk(inserts...).Exec(ctx)
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
		where = append(where, accountsnapshot.CreatedAtGT(query.createdAfter.Unix()))
	}
	if query.createdBefore != nil {
		where = append(where, accountsnapshot.CreatedAtLT(query.createdBefore.Unix()))
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
		where = append(where, accountsnapshot.CreatedAtGT(query.createdAfter.Unix()))
	}
	if query.createdBefore != nil {
		where = append(where, accountsnapshot.CreatedAtLT(query.createdAfter.Unix()))
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
