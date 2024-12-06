package database

import (
	"time"

	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/sqlite"
)

type baseQueryOptions struct {
	referenceIDIn map[string]struct{}

	createdBefore *time.Time

	columns s.ColumnList
}

type Query func(*baseQueryOptions)

// --- snapshot query options ---

/*
Constrain referenceID field for the query
  - if the final list of reference IDs is > 0, reference_id in (ids) will be added to the query
*/
func WithReferenceIDIn(ids ...string) Query {
	return func(q *baseQueryOptions) {
		if q.referenceIDIn == nil {
			q.referenceIDIn = make(map[string]struct{})
		}
		for _, id := range ids {
			q.referenceIDIn[id] = struct{}{}
		}
	}
}

/*
Adds a created_at gt constraint
  - if this constraint is set, records will be sorted by created_at DESC
*/
func WithCreatedBefore(before time.Time) Query {
	return func(q *baseQueryOptions) {
		q.createdBefore = &before
	}
}

/*
Adds columns used for SELECT
*/
func withSelect(columns s.ColumnList) Query {
	return func(q *baseQueryOptions) {
		q.columns = columns
	}
}

func (q *baseQueryOptions) refIDIn() []string {
	if len(q.referenceIDIn) == 0 {
		return nil
	}
	var ids []string
	for id := range q.referenceIDIn {
		ids = append(ids, id)
	}
	return ids
}

// build a complete query for vehicle snapshots
func vehiclesQuery(accountID string, vehicleIDs []string, kind models.SnapshotType, groupBy s.GroupByClause, query baseQueryOptions) s.Statement {
	// required where constraints
	var innerWhere []s.BoolExpression
	innerWhere = append(innerWhere, column(t.VehicleSnapshot.Type).EQ(s.String(string(kind))), column(t.VehicleSnapshot.AccountID).EQ(s.String(accountID)))

	// optional where constraints
	if vehicleIDs != nil {
		innerWhere = append(innerWhere, column(t.VehicleSnapshot.VehicleID).IN(stringsToExp(vehicleIDs)...))
	}
	if in := query.refIDIn(); in != nil {
		innerWhere = append(innerWhere, column(t.VehicleSnapshot.ReferenceID).IN(stringsToExp(in)...))
	}

	// order and created_at constraints
	innerOrder := column(t.VehicleSnapshot.CreatedAt).DESC()
	if query.createdBefore != nil {
		innerWhere = append(innerWhere, column(t.VehicleSnapshot.CreatedAt).LT(timeToField(*query.createdBefore)))
	}

	if query.columns == nil {
		query.columns = t.VehicleSnapshot.AllColumns
	}

	innerQuery := t.VehicleSnapshot.
		SELECT(s.STAR).
		WHERE(s.AND(innerWhere...)).
		ORDER_BY(innerOrder).
		AsTable("vehicle_snapshot")

	return s.
		SELECT(query.columns).
		FROM(innerQuery).
		GROUP_BY(groupBy)
}

// build a complete query for account snapshot
func accountsQuery(accountIDs []string, kind models.SnapshotType, groupBy s.GroupByClause, query baseQueryOptions) s.SelectStatement {
	var innerWhere []s.BoolExpression
	innerWhere = append(innerWhere, column(t.AccountSnapshot.Type).EQ(s.String(string(kind))), column(t.AccountSnapshot.AccountID).IN(stringsToExp(accountIDs)...))

	// optional where constraints
	if in := query.refIDIn(); in != nil {
		innerWhere = append(innerWhere, column(t.AccountSnapshot.ReferenceID).IN(stringsToExp(in)...))

	}

	// order and created_at constraints
	innerOrder := column(t.AccountSnapshot.CreatedAt).DESC()
	if query.createdBefore != nil {
		innerWhere = append(innerWhere, column(t.AccountSnapshot.CreatedAt).LT(timeToField(*query.createdBefore)))
	}

	if query.columns == nil {
		query.columns = t.AccountSnapshot.AllColumns
	}

	innerQuery := t.AccountSnapshot.
		SELECT(s.STAR).
		WHERE(s.AND(innerWhere...)).
		ORDER_BY(innerOrder).
		AsTable("account_snapshot")

	return s.
		SELECT(query.columns).
		FROM(innerQuery).
		GROUP_BY(groupBy)
}

type named interface {
	Name() string
}

func column(c named) s.ColumnString {
	return s.StringColumn(c.Name())
}
