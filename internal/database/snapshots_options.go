package database

import (
	"time"

	t "github.com/cufee/aftermath/internal/database/gen/table"
	"github.com/cufee/aftermath/internal/database/models"
	s "github.com/go-jet/jet/v2/sqlite"
)

type baseQueryOptions struct {
	referenceIDIn    map[string]struct{}
	referenceIDNotIn map[string]struct{}

	createdAfter  *time.Time
	createdBefore *time.Time

	columns []string
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
Constrain referenceID field for the query
  - if the final list of reference IDs is > 0, reference_id not in (ids) will be added to the query
*/
func WithReferenceIDNotIn(ids ...string) Query {
	return func(q *baseQueryOptions) {
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
func WithCreatedAfter(after time.Time) Query {
	return func(q *baseQueryOptions) {
		q.createdAfter = &after
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
func WithSelect(columns ...string) Query {
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

func (q *baseQueryOptions) refIDNotIn() []string {
	if len(q.referenceIDNotIn) == 0 {
		return nil
	}
	var ids []string
	for id := range q.referenceIDNotIn {
		ids = append(ids, id)
	}
	return ids
}

// build a complete query for vehicle snapshots
func vehiclesQuery(accountID string, vehicleIDs []string, kind models.SnapshotType, groupBy string, query baseQueryOptions) s.Statement {
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
	if nin := query.refIDNotIn(); nin != nil {
		innerWhere = append(innerWhere, column(t.VehicleSnapshot.ReferenceID).NOT_IN(stringsToExp(nin)...))
	}

	// order and created_at constraints
	innerOrder := column(t.VehicleSnapshot.CreatedAt).DESC()
	if query.createdAfter != nil {
		innerWhere = append(innerWhere, s.DateTimeColumn(t.VehicleSnapshot.CreatedAt.Name()).GT(timeValue(query.createdAfter)))
		innerOrder = column(t.VehicleSnapshot.CreatedAt).ASC()
	}
	if query.createdBefore != nil {
		innerWhere = append(innerWhere, s.DateTimeColumn(t.VehicleSnapshot.CreatedAt.Name()).LT(timeValue(query.createdBefore)))
		innerOrder = column(t.VehicleSnapshot.CreatedAt).DESC()
	}

	// var sel s.Projection = s.STAR
	// if len(query.columns)  > 0 {
	// 	sel = s.SELECT(s.StringColumn(query.columns[0]))
	// }

	innerQuery := t.VehicleSnapshot.
		SELECT(s.STAR).
		WHERE(s.AND(innerWhere...)).
		ORDER_BY(innerOrder).
		AsTable("snapshots")

	return s.
		SELECT(s.STAR).
		FROM(innerQuery).
		GROUP_BY(s.StringColumn(groupBy))
}

// build a complete query for account snapshot
func accountsQuery(accountIDs []string, kind models.SnapshotType, groupBy string, query baseQueryOptions) s.SelectStatement {
	var innerWhere []s.BoolExpression
	innerWhere = append(innerWhere, t.AccountSnapshot.Type.EQ(s.String(string(kind))), t.AccountSnapshot.AccountID.IN(stringsToExp(accountIDs)...))

	// optional where constraints
	if in := query.refIDIn(); in != nil {
		innerWhere = append(innerWhere, t.AccountSnapshot.ReferenceID.IN(stringsToExp(in)...))

	}
	if nin := query.refIDNotIn(); nin != nil {
		innerWhere = append(innerWhere, t.AccountSnapshot.ReferenceID.NOT_IN(stringsToExp(nin)...))
	}

	// order and created_at constraints
	innerOrder := s.String(t.AccountSnapshot.CreatedAt.Name()).DESC()
	if query.createdAfter != nil {
		innerWhere = append(innerWhere, s.DateTimeColumn(t.AccountSnapshot.CreatedAt.Name()).GT(timeValue(query.createdAfter)))
		innerOrder = s.String(t.AccountSnapshot.CreatedAt.Name()).ASC()
	}
	if query.createdBefore != nil {
		innerWhere = append(innerWhere, s.DateTimeColumn(t.AccountSnapshot.CreatedAt.Name()).LT(timeValue(query.createdBefore)))
		innerOrder = s.String(t.AccountSnapshot.CreatedAt.Name()).DESC()
	}

	// var sel s.Projection = t.AccountSnapshot.AllColumns
	// if query.columns != nil {
	// 	sel = query.columns
	// }

	innerQuery := t.AccountSnapshot.
		SELECT(s.STAR).
		WHERE(s.AND(innerWhere...)).
		ORDER_BY(innerOrder).
		AsTable("snapshots")

	return s.
		SELECT(s.STAR).
		FROM(innerQuery).
		GROUP_BY(s.StringColumn(groupBy))
}

type named interface {
	Name() string
}

func column(c named) s.ColumnString {
	return s.StringColumn(c.Name())
}

type formattable interface {
	Format(layout string) string
}

func timeValue(t formattable) s.DateTimeExpression {
	return s.DATETIME(t.Format(time.RFC3339))
}
