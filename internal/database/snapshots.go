package database

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db"
	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/achievementssnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
	"github.com/cufee/aftermath/internal/database/models"
)

type getSnapshotQuery struct {
	referenceIDIn    map[string]struct{}
	referenceIDNotIn map[string]struct{}

	createdAfter  *time.Time
	createdBefore *time.Time

	fields []string
}

type SnapshotQuery func(*getSnapshotQuery)

// --- snapshot query options ---

/*
Constrain referenceID field for the query
  - if the final list of reference IDs is > 0, reference_id in (ids) will be added to the query
*/
func WithReferenceIDIn(ids ...string) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		if q.referenceIDIn == nil {
			q.referenceIDIn = make(map[string]struct{})
		}
		for _, id := range ids {
			q.referenceIDIn[id] = struct{}{}
		}
	}
}

/*
Constrain referenceID field for the query
  - if the final list of reference IDs is > 0, reference_id not in (ids) will be added to the query
*/
func WithReferenceIDNotIn(ids ...string) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		if q.referenceIDNotIn == nil {
			q.referenceIDNotIn = make(map[string]struct{})
		}
		for _, id := range ids {
			q.referenceIDNotIn[id] = struct{}{}
		}
	}
}

/*
Adds a created_at lt constraint
  - if this constraint is set, records will be sorted by created_at ASC
*/
func WithCreatedAfter(after time.Time) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		q.createdAfter = &after
	}
}

/*
Adds a created_at gt constraint
  - if this constraint is set, records will be sorted by created_at DESC
*/
func WithCreatedBefore(before time.Time) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		q.createdBefore = &before
	}
}

/*
Set fields that will be selected.
  - Some fields like id, created_at, updated_at will always be selected
  - Passing 0 length fields will result in select *
*/
func WithSelect(fields ...string) SnapshotQuery {
	return func(q *getSnapshotQuery) {
		// make sure fields are unique
		fieldsSet := make(map[string]struct{})
		for _, field := range append(q.fields, fields...) {
			fieldsSet[field] = struct{}{}
		}

		// passing 0 length fields will result in select *
		q.fields = nil
		for field := range fieldsSet {
			q.fields = append(q.fields, field)
		}
	}
}

/*
Returns a slice of fields for a select statement with all duplicates removed and required fields added
  - if query.fields is nil, returns a nil slice
*/
func (q *getSnapshotQuery) selectFields(required ...string) []string {
	if q.fields == nil {
		return nil
	}

	// maker sure all required fields are part of the slice
	var requiredFields = make(map[string]struct{})
	for _, field := range required {
		requiredFields[field] = struct{}{}
	}

	for _, field := range q.fields {
		for required := range requiredFields {
			if field == required {
				delete(requiredFields, required)
			}
		}
	}
	for field := range requiredFields {
		q.fields = append(q.fields, field)
	}

	return q.fields
}

func (q *getSnapshotQuery) refIDIn() []any {
	if len(q.referenceIDIn) == 0 {
		return nil
	}
	var ids []any
	for id := range q.referenceIDIn {
		ids = append(ids, id)
	}
	return ids
}

func (q *getSnapshotQuery) refIDNotIn() []any {
	if len(q.referenceIDNotIn) == 0 {
		return nil
	}
	var ids []any
	for id := range q.referenceIDNotIn {
		ids = append(ids, id)
	}
	return ids
}

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

func toAchievementsSnapshot(record *db.AchievementsSnapshot) models.AchievementsSnapshot {
	return models.AchievementsSnapshot{
		ID:             record.ID,
		Type:           record.Type,
		CreatedAt:      record.CreatedAt,
		LastBattleTime: record.LastBattleTime,
		ReferenceID:    record.ReferenceID,
		AccountID:      record.AccountID,
		Battles:        record.Battles,
		Data:           record.Data,
	}
}

// --- vehicle snapshots ---

/*
Get complete vehicle snapshots fot each vehicle ID in vehicle IDs for a specific account
  - passing nil vehicleIDs will return all vehicles available
*/
func (c *client) GetVehicleSnapshots(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...SnapshotQuery) ([]models.VehicleSnapshot, error) {
	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	queryString, queryArgs := vehiclesQuery(accountID, vehicleIDs, kind, vehiclesnapshot.FieldVehicleID, query)
	rows, err := c.db.VehicleSnapshot.QueryContext(ctx, queryString, queryArgs...)
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

/*
get last battle times for each vehicle in vehicle IDs for a specific account
  - passing nil vehicleIDs will return all vehicles available
*/
func (c *client) GetVehicleLastBattleTimes(ctx context.Context, accountID string, vehicleIDs []string, kind models.SnapshotType, options ...SnapshotQuery) (map[string]time.Time, error) {
	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}
	WithSelect(vehiclesnapshot.FieldVehicleID, vehiclesnapshot.FieldLastBattleTime)(&query)

	queryString, queryArgs := vehiclesQuery(accountID, vehicleIDs, kind, vehiclesnapshot.FieldVehicleID, query)
	rows, err := c.db.VehicleSnapshot.QueryContext(ctx, queryString, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lastBattleTimes = make(map[string]time.Time)
	for rows.Next() {
		var record db.VehicleSnapshot

		values, err := record.ScanValues([]string{vehiclesnapshot.FieldVehicleID, vehiclesnapshot.FieldLastBattleTime})
		if err != nil {
			return nil, err
		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		err = record.AssignValues([]string{vehiclesnapshot.FieldVehicleID, vehiclesnapshot.FieldLastBattleTime}, values)
		if err != nil {
			return nil, err
		}
		lastBattleTimes[record.VehicleID] = record.LastBattleTime
	}

	return lastBattleTimes, nil
}

// create vehicle snapshots for a specific account
func (c *client) CreateAccountVehicleSnapshots(ctx context.Context, accountID string, snapshots ...models.VehicleSnapshot) (map[string]error, error) {
	if len(snapshots) < 1 {
		return nil, new(db.NotFoundError)
	}

	account, err := c.db.Account.Get(ctx, accountID)
	if err != nil {
		return nil, err
	}

	var errors = make(map[string]error)
	for _, data := range snapshots {
		// make a transaction per write to avoid locking for too long
		err := c.withTx(ctx, func(tx *db.Tx) error {
			return c.db.VehicleSnapshot.Create().
				SetType(data.Type).
				SetFrame(data.Stats).
				SetVehicleID(data.VehicleID).
				SetReferenceID(data.ReferenceID).
				SetCreatedAt(data.CreatedAt).
				SetBattles(int(data.Stats.Battles.Float())).
				SetLastBattleTime(data.LastBattleTime).
				SetAccount(account).
				Exec(ctx)
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

// build a complete query for vehicle snapshots
func vehiclesQuery(accountID string, vehicleIDs []string, kind models.SnapshotType, groupBy string, query getSnapshotQuery) (string, []any) {
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

	innerQuery := sql.Select(vehiclesnapshot.Columns...).From(sql.Table(vehiclesnapshot.Table))
	innerQuery = innerQuery.Where(sql.And(innerWhere...))
	innerQuery = innerQuery.OrderBy(innerOrder)

	innerQueryString, innerQueryArgs := innerQuery.Query()

	// wrap the inner query in a GROUP BY
	wrapper := &sql.Builder{}
	wrapped := wrapper.Wrap(func(b *sql.Builder) { b.WriteString(innerQueryString) })

	selector := sql.Select("*")
	if fields := query.selectFields(groupBy); fields != nil {
		// make sure FieldVehicleID is selected, we need it for Group By
		selector = sql.Select(fields...)
	}
	queryString, _ := selector.FromExpr(wrapped).GroupBy(groupBy).Query()
	return queryString, innerQueryArgs
}

// --- account snapshots ---

/*
Get complete snapshots for accounts by ID, grouped by account ID
  - there are no use cases where all accounts should be returned for now, so nil slice of ids will return an error
*/
func (c *client) GetAccountSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...SnapshotQuery) ([]models.AccountSnapshot, error) {
	if len(accountIDs) < 1 {
		return nil, new(db.NotFoundError)
	}

	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	queryString, queryArgs := accountsQuery(accountIDs, kind, accountsnapshot.FieldAccountID, query)
	rows, err := c.db.AccountSnapshot.QueryContext(ctx, queryString, queryArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snapshots []models.AccountSnapshot
	for rows.Next() {
		var record db.AccountSnapshot

		values, err := record.ScanValues(accountsnapshot.Columns)
		if err != nil {
			return nil, err
		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		err = record.AssignValues(accountsnapshot.Columns, values)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, toAccountSnapshot(&record))
	}

	return snapshots, nil
}

/*
Get last battle times for accounts by ID, grouped by account ID
*/
func (c *client) GetAccountLastBattleTimes(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...SnapshotQuery) (map[string]time.Time, error) {
	if len(accountIDs) < 1 {
		return nil, new(db.NotFoundError)
	}

	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}
	WithSelect(accountsnapshot.FieldAccountID, accountsnapshot.FieldLastBattleTime)(&query)

	queryStr, queryArgs := accountsQuery(accountIDs, kind, accountsnapshot.FieldAccountID, query)
	rows, err := c.db.AccountSnapshot.QueryContext(ctx, queryStr, queryArgs...)
	if err != nil {
		return nil, err
	}

	var lastBattles = make(map[string]time.Time)
	for rows.Next() {
		var record db.AccountSnapshot

		values, err := record.ScanValues([]string{accountsnapshot.FieldAccountID, accountsnapshot.FieldLastBattleTime})
		if err != nil {
			return nil, err
		}
		if err := rows.Scan(values...); err != nil {
			return nil, err
		}

		err = record.AssignValues([]string{accountsnapshot.FieldAccountID, accountsnapshot.FieldLastBattleTime}, values)
		if err != nil {
			return nil, err
		}

		lastBattles[record.ReferenceID] = record.LastBattleTime
	}

	return lastBattles, nil
}

func (c *client) CreateAccountSnapshots(ctx context.Context, snapshots ...models.AccountSnapshot) (map[string]error, error) {
	if len(snapshots) < 1 {
		return nil, new(db.NotFoundError)
	}

	var errors = make(map[string]error)
	for _, s := range snapshots {
		account, err := c.db.Account.Get(ctx, s.AccountID)
		if err != nil {
			errors[s.AccountID] = err
			continue
		}
		// make a transaction per write to avoid locking for too long
		err = c.withTx(ctx, func(tx *db.Tx) error {
			return c.db.AccountSnapshot.Create().
				SetAccount(account).
				SetCreatedAt(s.CreatedAt).
				SetLastBattleTime(s.LastBattleTime).
				SetRatingBattles(int(s.RatingBattles.Battles.Float())).
				SetRatingFrame(s.RatingBattles).
				SetReferenceID(s.ReferenceID).
				SetRegularBattles(int(s.RegularBattles.Battles)).
				SetRegularFrame(s.RegularBattles).
				SetType(s.Type).Exec(ctx)
		})
		if err != nil {
			errors[s.AccountID] = err
		}
	}
	if len(errors) > 0 {
		return errors, nil
	}
	return nil, nil
}

// build a complete query for account snapshot
func accountsQuery(accountIDs []string, kind models.SnapshotType, groupBy string, query getSnapshotQuery) (string, []any) {
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

	innerQuery := sql.Select(accountsnapshot.Columns...).From(sql.Table(accountsnapshot.Table))
	innerQuery = innerQuery.Where(sql.And(innerWhere...))
	innerQuery = innerQuery.OrderBy(innerOrder)

	innerQueryString, innerQueryArgs := innerQuery.Query()

	// wrap the inner query in a GROUP BY
	wrapper := &sql.Builder{}
	wrapped := wrapper.Wrap(func(b *sql.Builder) { b.WriteString(innerQueryString) })

	selector := sql.Select("*")
	if fields := query.selectFields(groupBy); fields != nil {
		// make sure FieldVehicleID is selected, we need it for Group By
		selector = sql.Select(fields...)
	}
	queryString, _ := selector.FromExpr(wrapped).GroupBy(groupBy).Query()
	return queryString, innerQueryArgs
}

// --- achievement snapshots ---

func (c *client) GetAchievementSnapshots(ctx context.Context, accountIDs []string, kind models.SnapshotType, options ...SnapshotQuery) ([]models.AchievementsSnapshot, error) {
	if len(accountIDs) < 1 {
		return nil, new(db.NotFoundError)
	}

	var query getSnapshotQuery
	for _, apply := range options {
		apply(&query)
	}

	var where []predicate.AchievementsSnapshot
	where = append(where, achievementssnapshot.AccountIDIn(accountIDs...), achievementssnapshot.TypeEQ(kind))
	if query.createdAfter != nil {
		where = append(where, achievementssnapshot.CreatedAtGT(*query.createdAfter))
	}
	if query.createdBefore != nil {
		where = append(where, achievementssnapshot.CreatedAtLT(*query.createdAfter))
	}

	records, err := c.db.AchievementsSnapshot.Query().Select(query.selectFields()...).Where(where...).All(ctx)
	if err != nil {
		return nil, err
	}

	var snapshots []models.AchievementsSnapshot
	for _, r := range records {
		snapshots = append(snapshots, toAchievementsSnapshot(r))
	}

	return snapshots, nil
}

func (c *client) CreateAccountAchievementSnapshots(ctx context.Context, accountID string, snapshots ...models.AchievementsSnapshot) (map[string]error, error) {
	if len(snapshots) < 1 {
		return nil, new(db.NotFoundError)
	}

	account, err := c.db.Account.Get(ctx, accountID)
	if err != nil {
		return nil, err
	}

	var errors = make(map[string]error)
	for _, data := range snapshots {
		// make a transaction per write to avoid locking for too long
		err := c.withTx(ctx, func(tx *db.Tx) error {
			return c.db.AchievementsSnapshot.Create().
				SetType(data.Type).
				SetData(data.Data).
				SetBattles(data.Battles).
				SetCreatedAt(data.CreatedAt).
				SetReferenceID(data.ReferenceID).
				SetLastBattleTime(data.LastBattleTime).
				SetAccount(account).
				Exec(ctx)
		})
		if err != nil {
			errors[data.ReferenceID] = err
		}
	}

	if len(errors) > 0 {
		return errors, nil
	}
	return nil, nil
}
