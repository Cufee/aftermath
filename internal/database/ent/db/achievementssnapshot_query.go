// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/ent/db/achievementssnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// AchievementsSnapshotQuery is the builder for querying AchievementsSnapshot entities.
type AchievementsSnapshotQuery struct {
	config
	ctx         *QueryContext
	order       []achievementssnapshot.OrderOption
	inters      []Interceptor
	predicates  []predicate.AchievementsSnapshot
	withAccount *AccountQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AchievementsSnapshotQuery builder.
func (asq *AchievementsSnapshotQuery) Where(ps ...predicate.AchievementsSnapshot) *AchievementsSnapshotQuery {
	asq.predicates = append(asq.predicates, ps...)
	return asq
}

// Limit the number of records to be returned by this query.
func (asq *AchievementsSnapshotQuery) Limit(limit int) *AchievementsSnapshotQuery {
	asq.ctx.Limit = &limit
	return asq
}

// Offset to start from.
func (asq *AchievementsSnapshotQuery) Offset(offset int) *AchievementsSnapshotQuery {
	asq.ctx.Offset = &offset
	return asq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (asq *AchievementsSnapshotQuery) Unique(unique bool) *AchievementsSnapshotQuery {
	asq.ctx.Unique = &unique
	return asq
}

// Order specifies how the records should be ordered.
func (asq *AchievementsSnapshotQuery) Order(o ...achievementssnapshot.OrderOption) *AchievementsSnapshotQuery {
	asq.order = append(asq.order, o...)
	return asq
}

// QueryAccount chains the current query on the "account" edge.
func (asq *AchievementsSnapshotQuery) QueryAccount() *AccountQuery {
	query := (&AccountClient{config: asq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := asq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(achievementssnapshot.Table, achievementssnapshot.FieldID, selector),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, achievementssnapshot.AccountTable, achievementssnapshot.AccountColumn),
		)
		fromU = sqlgraph.SetNeighbors(asq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AchievementsSnapshot entity from the query.
// Returns a *NotFoundError when no AchievementsSnapshot was found.
func (asq *AchievementsSnapshotQuery) First(ctx context.Context) (*AchievementsSnapshot, error) {
	nodes, err := asq.Limit(1).All(setContextOp(ctx, asq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{achievementssnapshot.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (asq *AchievementsSnapshotQuery) FirstX(ctx context.Context) *AchievementsSnapshot {
	node, err := asq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AchievementsSnapshot ID from the query.
// Returns a *NotFoundError when no AchievementsSnapshot ID was found.
func (asq *AchievementsSnapshotQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = asq.Limit(1).IDs(setContextOp(ctx, asq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{achievementssnapshot.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (asq *AchievementsSnapshotQuery) FirstIDX(ctx context.Context) string {
	id, err := asq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AchievementsSnapshot entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AchievementsSnapshot entity is found.
// Returns a *NotFoundError when no AchievementsSnapshot entities are found.
func (asq *AchievementsSnapshotQuery) Only(ctx context.Context) (*AchievementsSnapshot, error) {
	nodes, err := asq.Limit(2).All(setContextOp(ctx, asq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{achievementssnapshot.Label}
	default:
		return nil, &NotSingularError{achievementssnapshot.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (asq *AchievementsSnapshotQuery) OnlyX(ctx context.Context) *AchievementsSnapshot {
	node, err := asq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AchievementsSnapshot ID in the query.
// Returns a *NotSingularError when more than one AchievementsSnapshot ID is found.
// Returns a *NotFoundError when no entities are found.
func (asq *AchievementsSnapshotQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = asq.Limit(2).IDs(setContextOp(ctx, asq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{achievementssnapshot.Label}
	default:
		err = &NotSingularError{achievementssnapshot.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (asq *AchievementsSnapshotQuery) OnlyIDX(ctx context.Context) string {
	id, err := asq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AchievementsSnapshots.
func (asq *AchievementsSnapshotQuery) All(ctx context.Context) ([]*AchievementsSnapshot, error) {
	ctx = setContextOp(ctx, asq.ctx, "All")
	if err := asq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AchievementsSnapshot, *AchievementsSnapshotQuery]()
	return withInterceptors[[]*AchievementsSnapshot](ctx, asq, qr, asq.inters)
}

// AllX is like All, but panics if an error occurs.
func (asq *AchievementsSnapshotQuery) AllX(ctx context.Context) []*AchievementsSnapshot {
	nodes, err := asq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AchievementsSnapshot IDs.
func (asq *AchievementsSnapshotQuery) IDs(ctx context.Context) (ids []string, err error) {
	if asq.ctx.Unique == nil && asq.path != nil {
		asq.Unique(true)
	}
	ctx = setContextOp(ctx, asq.ctx, "IDs")
	if err = asq.Select(achievementssnapshot.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (asq *AchievementsSnapshotQuery) IDsX(ctx context.Context) []string {
	ids, err := asq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (asq *AchievementsSnapshotQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, asq.ctx, "Count")
	if err := asq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, asq, querierCount[*AchievementsSnapshotQuery](), asq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (asq *AchievementsSnapshotQuery) CountX(ctx context.Context) int {
	count, err := asq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (asq *AchievementsSnapshotQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, asq.ctx, "Exist")
	switch _, err := asq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (asq *AchievementsSnapshotQuery) ExistX(ctx context.Context) bool {
	exist, err := asq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AchievementsSnapshotQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (asq *AchievementsSnapshotQuery) Clone() *AchievementsSnapshotQuery {
	if asq == nil {
		return nil
	}
	return &AchievementsSnapshotQuery{
		config:      asq.config,
		ctx:         asq.ctx.Clone(),
		order:       append([]achievementssnapshot.OrderOption{}, asq.order...),
		inters:      append([]Interceptor{}, asq.inters...),
		predicates:  append([]predicate.AchievementsSnapshot{}, asq.predicates...),
		withAccount: asq.withAccount.Clone(),
		// clone intermediate query.
		sql:  asq.sql.Clone(),
		path: asq.path,
	}
}

// WithAccount tells the query-builder to eager-load the nodes that are connected to
// the "account" edge. The optional arguments are used to configure the query builder of the edge.
func (asq *AchievementsSnapshotQuery) WithAccount(opts ...func(*AccountQuery)) *AchievementsSnapshotQuery {
	query := (&AccountClient{config: asq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	asq.withAccount = query
	return asq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AchievementsSnapshot.Query().
//		GroupBy(achievementssnapshot.FieldCreatedAt).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (asq *AchievementsSnapshotQuery) GroupBy(field string, fields ...string) *AchievementsSnapshotGroupBy {
	asq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AchievementsSnapshotGroupBy{build: asq}
	grbuild.flds = &asq.ctx.Fields
	grbuild.label = achievementssnapshot.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//	}
//
//	client.AchievementsSnapshot.Query().
//		Select(achievementssnapshot.FieldCreatedAt).
//		Scan(ctx, &v)
func (asq *AchievementsSnapshotQuery) Select(fields ...string) *AchievementsSnapshotSelect {
	asq.ctx.Fields = append(asq.ctx.Fields, fields...)
	sbuild := &AchievementsSnapshotSelect{AchievementsSnapshotQuery: asq}
	sbuild.label = achievementssnapshot.Label
	sbuild.flds, sbuild.scan = &asq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AchievementsSnapshotSelect configured with the given aggregations.
func (asq *AchievementsSnapshotQuery) Aggregate(fns ...AggregateFunc) *AchievementsSnapshotSelect {
	return asq.Select().Aggregate(fns...)
}

func (asq *AchievementsSnapshotQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range asq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, asq); err != nil {
				return err
			}
		}
	}
	for _, f := range asq.ctx.Fields {
		if !achievementssnapshot.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if asq.path != nil {
		prev, err := asq.path(ctx)
		if err != nil {
			return err
		}
		asq.sql = prev
	}
	return nil
}

func (asq *AchievementsSnapshotQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AchievementsSnapshot, error) {
	var (
		nodes       = []*AchievementsSnapshot{}
		_spec       = asq.querySpec()
		loadedTypes = [1]bool{
			asq.withAccount != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AchievementsSnapshot).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AchievementsSnapshot{config: asq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, asq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := asq.withAccount; query != nil {
		if err := asq.loadAccount(ctx, query, nodes, nil,
			func(n *AchievementsSnapshot, e *Account) { n.Edges.Account = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (asq *AchievementsSnapshotQuery) loadAccount(ctx context.Context, query *AccountQuery, nodes []*AchievementsSnapshot, init func(*AchievementsSnapshot), assign func(*AchievementsSnapshot, *Account)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*AchievementsSnapshot)
	for i := range nodes {
		fk := nodes[i].AccountID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(account.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "account_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (asq *AchievementsSnapshotQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := asq.querySpec()
	_spec.Node.Columns = asq.ctx.Fields
	if len(asq.ctx.Fields) > 0 {
		_spec.Unique = asq.ctx.Unique != nil && *asq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, asq.driver, _spec)
}

func (asq *AchievementsSnapshotQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(achievementssnapshot.Table, achievementssnapshot.Columns, sqlgraph.NewFieldSpec(achievementssnapshot.FieldID, field.TypeString))
	_spec.From = asq.sql
	if unique := asq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if asq.path != nil {
		_spec.Unique = true
	}
	if fields := asq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, achievementssnapshot.FieldID)
		for i := range fields {
			if fields[i] != achievementssnapshot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if asq.withAccount != nil {
			_spec.Node.AddColumnOnce(achievementssnapshot.FieldAccountID)
		}
	}
	if ps := asq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := asq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := asq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := asq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (asq *AchievementsSnapshotQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(asq.driver.Dialect())
	t1 := builder.Table(achievementssnapshot.Table)
	columns := asq.ctx.Fields
	if len(columns) == 0 {
		columns = achievementssnapshot.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if asq.sql != nil {
		selector = asq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if asq.ctx.Unique != nil && *asq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range asq.predicates {
		p(selector)
	}
	for _, p := range asq.order {
		p(selector)
	}
	if offset := asq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := asq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AchievementsSnapshotGroupBy is the group-by builder for AchievementsSnapshot entities.
type AchievementsSnapshotGroupBy struct {
	selector
	build *AchievementsSnapshotQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (asgb *AchievementsSnapshotGroupBy) Aggregate(fns ...AggregateFunc) *AchievementsSnapshotGroupBy {
	asgb.fns = append(asgb.fns, fns...)
	return asgb
}

// Scan applies the selector query and scans the result into the given value.
func (asgb *AchievementsSnapshotGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, asgb.build.ctx, "GroupBy")
	if err := asgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AchievementsSnapshotQuery, *AchievementsSnapshotGroupBy](ctx, asgb.build, asgb, asgb.build.inters, v)
}

func (asgb *AchievementsSnapshotGroupBy) sqlScan(ctx context.Context, root *AchievementsSnapshotQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(asgb.fns))
	for _, fn := range asgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*asgb.flds)+len(asgb.fns))
		for _, f := range *asgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*asgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AchievementsSnapshotSelect is the builder for selecting fields of AchievementsSnapshot entities.
type AchievementsSnapshotSelect struct {
	*AchievementsSnapshotQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ass *AchievementsSnapshotSelect) Aggregate(fns ...AggregateFunc) *AchievementsSnapshotSelect {
	ass.fns = append(ass.fns, fns...)
	return ass
}

// Scan applies the selector query and scans the result into the given value.
func (ass *AchievementsSnapshotSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ass.ctx, "Select")
	if err := ass.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AchievementsSnapshotQuery, *AchievementsSnapshotSelect](ctx, ass.AchievementsSnapshotQuery, ass, ass.inters, v)
}

func (ass *AchievementsSnapshotSelect) sqlScan(ctx context.Context, root *AchievementsSnapshotQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ass.fns))
	for _, fn := range ass.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ass.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}