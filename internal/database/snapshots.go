package database

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
	"github.com/cufee/aftermath/internal/database/models"
)

// --- record to model ---

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

	queryString, columns, queryArgs := vehiclesQuery(accountID, vehicleIDs, kind, vehiclesnapshot.FieldVehicleID, query)
	rows, err := c.db.VehicleSnapshot.QueryContext(ctx, queryString, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records, err := rowsToRecords[*db.VehicleSnapshot](rows, columns)
	if err != nil {
		return nil, err
	}

	var snapshots []models.VehicleSnapshot
	for _, r := range records {
		snapshots = append(snapshots, toVehicleSnapshot(r))
	}
	return snapshots, nil
}

/*
get last battle times for each vehicle in vehicle IDs for a specific account
  - passing nil vehicleIDs will return all vehicles available
*/
func (c *client) GetVehicleLastBattleTimes(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...Query) (map[string]time.Time, error) {
	var query baseQueryOptions
	for _, apply := range options {
		apply(&query)
	}
	WithSelect(vehiclesnapshot.FieldVehicleID, vehiclesnapshot.FieldLastBattleTime)(&query)

	queryString, columns, queryArgs := vehiclesQuery(accountID, vehicleIDs, kind, vehiclesnapshot.FieldVehicleID, query)
	rows, err := c.db.VehicleSnapshot.QueryContext(ctx, queryString, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records, err := rowsToRecords[*db.VehicleSnapshot](rows, columns)
	if err != nil {
		return nil, err
	}

	var lastBattleTimes = make(map[string]time.Time)
	for _, r := range records {
		lastBattleTimes[r.VehicleID] = r.LastBattleTime
	}

	return lastBattleTimes, nil
}

// create vehicle snapshots for a specific account
func (c *client) CreateAccountVehicleSnapshots(ctx context.Context, accountID string, snapshots ...models.VehicleSnapshot) error {
	if len(snapshots) < 1 {
		return nil
	}

	account, err := c.db.Account.Get(ctx, accountID)
	if err != nil {
		return err
	}

	var inserts []*db.VehicleSnapshotCreate
	for _, data := range snapshots {
		inserts = append(inserts,
			c.db.VehicleSnapshot.Create().
				SetType(data.Type).
				SetFrame(data.Stats).
				SetVehicleID(data.VehicleID).
				SetReferenceID(data.ReferenceID).
				SetCreatedAt(data.CreatedAt).
				SetBattles(int(data.Stats.Battles.Float())).
				SetLastBattleTime(data.LastBattleTime).
				SetAccount(account),
		)
	}

	for _, ops := range batch(inserts, 100) {
		err := c.withTx(ctx, func(tx *db.Tx) error {
			return tx.VehicleSnapshot.CreateBulk(ops...).Exec(ctx)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// build a complete query for vehicle snapshots
func vehiclesQuery(accountID string, vehicleIDs []string, kind models.SnapshotType, groupBy string, query baseQueryOptions) (string, []string, []any) {
	// required where constraints
	var innerWhere []*sql.Predicate
	innerWhere = append(innerWhere, sql.EQ(vehiclesnapshot.FieldType, kind), sql.EQ(vehiclesnapshot.FieldAccountID, accountID))

	// optional where constraints
	if vehicleIDs != nil {
		innerWhere = append(innerWhere, sql.In(vehiclesnapshot.FieldVehicleID, toAnySlice(vehicleIDs...)...))
	}
	if in := query.refIDIn(); in != nil {
		innerWhere = append(innerWhere, sql.In(vehiclesnapshot.FieldReferenceID, in...))
	}
	if nin := query.refIDNotIn(); nin != nil {
		innerWhere = append(innerWhere, sql.NotIn(vehiclesnapshot.FieldReferenceID, nin...))
	}

	// order and created_at constraints
	innerOrder := sql.Desc(vehiclesnapshot.FieldCreatedAt)
	if query.createdAfter != nil {
		innerWhere = append(innerWhere, sql.GT(vehiclesnapshot.FieldCreatedAt, *query.createdAfter))
		innerOrder = sql.Asc(vehiclesnapshot.FieldCreatedAt)
	}
	if query.createdBefore != nil {
		innerWhere = append(innerWhere, sql.LT(vehiclesnapshot.FieldCreatedAt, *query.createdBefore))
		innerOrder = sql.Desc(vehiclesnapshot.FieldCreatedAt)
	}

	selectFields := vehiclesnapshot.Columns
	if fields := query.selectFields(groupBy); fields != nil {
		selectFields = fields
	}

	innerQuery := sql.Select(selectFields...).From(sql.Table(vehiclesnapshot.Table))
	innerQuery = innerQuery.Where(sql.And(innerWhere...))
	innerQuery = innerQuery.OrderBy(innerOrder)

	innerQueryString, innerQueryArgs := innerQuery.Query()

	queryString, _ := sql.Select(selectFields...).FromExpr(wrap(innerQueryString)).GroupBy(groupBy).Query()
	return queryString, selectFields, innerQueryArgs
}

// --- account snapshots ---

/*
Get complete snapshots for accounts by ID, grouped by account ID
  - there are no use cases where all accounts should be returned for now, so nil slice of ids will return an error
*/
func (c *client) GetAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...Query) ([]models.AccountSnapshot, error) {
	if len(accountIDs) < 1 {
		return nil, new(db.NotFoundError)
	}

	var query baseQueryOptions
	for _, apply := range options {
		apply(&query)
	}

	queryString, columns, queryArgs := accountsQuery(accountIDs, kind, accountsnapshot.FieldAccountID, query)
	rows, err := c.db.AccountSnapshot.QueryContext(ctx, queryString, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records, err := rowsToRecords[*db.AccountSnapshot](rows, columns)
	if err != nil {
		return nil, err
	}

	var snapshots []models.AccountSnapshot
	for _, r := range records {
		snapshots = append(snapshots, toAccountSnapshot(r))
	}
	return snapshots, nil
}

/*
Get last battle times for accounts by ID, grouped by account ID
*/
func (c *client) GetAccountLastBattleTimes(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...Query) (map[string]time.Time, error) {
	if len(accountIDs) < 1 {
		return nil, new(db.NotFoundError)
	}

	var query baseQueryOptions
	for _, apply := range options {
		apply(&query)
	}
	WithSelect(accountsnapshot.FieldAccountID, accountsnapshot.FieldLastBattleTime)(&query)

	queryStr, columns, queryArgs := accountsQuery(accountIDs, kind, accountsnapshot.FieldAccountID, query)
	rows, err := c.db.AccountSnapshot.QueryContext(ctx, queryStr, queryArgs...)
	if err != nil {
		return nil, err
	}

	records, err := rowsToRecords[*db.AccountSnapshot](rows, columns)
	if err != nil {
		return nil, err
	}

	var lastBattles = make(map[string]time.Time)
	for _, r := range records {
		lastBattles[r.AccountID] = r.LastBattleTime
	}
	return lastBattles, nil
}

func (c *client) CreateAccountSnapshots(ctx context.Context, snapshots ...models.AccountSnapshot) error {
	if len(snapshots) < 1 {
		return nil
	}

	var ids []string
	for _, s := range snapshots {
		ids = append(ids, s.AccountID)
	}

	accounts, err := c.db.Account.Query().Where(account.IDIn(ids...)).All(ctx)
	if err != nil {
		return err
	}
	accountsMap := make(map[string]*db.Account)
	for _, a := range accounts {
		accountsMap[a.ID] = a
	}

	var inserts []*db.AccountSnapshotCreate
	for _, s := range snapshots {
		inserts = append(inserts,
			c.db.AccountSnapshot.Create().
				SetAccount(accountsMap[s.AccountID]). // assume its valid, transaction will fail if it's not
				SetCreatedAt(s.CreatedAt).
				SetLastBattleTime(s.LastBattleTime).
				SetRatingBattles(int(s.RatingBattles.Battles.Float())).
				SetRatingFrame(s.RatingBattles).
				SetReferenceID(s.ReferenceID).
				SetRegularBattles(int(s.RegularBattles.Battles)).
				SetRegularFrame(s.RegularBattles).
				SetType(s.Type),
		)
	}

	for _, ops := range batch(inserts, 100) {
		err := c.withTx(ctx, func(tx *db.Tx) error {
			return tx.AccountSnapshot.CreateBulk(ops...).Exec(ctx)
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// build a complete query for account snapshot
func accountsQuery(accountIDs []string, kind models.SnapshotType, groupBy string, query baseQueryOptions) (string, []string, []any) {
	// required where constraints
	var innerWhere []*sql.Predicate
	innerWhere = append(innerWhere, sql.EQ(accountsnapshot.FieldType, kind), sql.In(accountsnapshot.FieldAccountID, toAnySlice(accountIDs...)...))

	// optional where constraints
	if in := query.refIDIn(); in != nil {
		innerWhere = append(innerWhere, sql.In(accountsnapshot.FieldReferenceID, in...))
	}
	if nin := query.refIDNotIn(); nin != nil {
		innerWhere = append(innerWhere, sql.NotIn(accountsnapshot.FieldReferenceID, nin...))
	}

	// order and created_at constraints
	innerOrder := sql.Desc(accountsnapshot.FieldCreatedAt)
	if query.createdAfter != nil {
		innerWhere = append(innerWhere, sql.GT(accountsnapshot.FieldCreatedAt, *query.createdAfter))
		innerOrder = sql.Asc(accountsnapshot.FieldCreatedAt)
	}
	if query.createdBefore != nil {
		innerWhere = append(innerWhere, sql.LT(accountsnapshot.FieldCreatedAt, *query.createdBefore))
		innerOrder = sql.Desc(accountsnapshot.FieldCreatedAt)
	}

	selectFields := accountsnapshot.Columns
	if fields := query.selectFields(groupBy); fields != nil {
		selectFields = fields
	}

	innerQuery := sql.Select(selectFields...).From(sql.Table(accountsnapshot.Table))
	innerQuery = innerQuery.Where(sql.And(innerWhere...))
	innerQuery = innerQuery.OrderBy(innerOrder)

	innerQueryString, innerQueryArgs := innerQuery.Query()
	queryString, _ := sql.Select(selectFields...).FromExpr(wrap(innerQueryString)).GroupBy(groupBy).Query()

	return queryString, selectFields, innerQueryArgs
}
