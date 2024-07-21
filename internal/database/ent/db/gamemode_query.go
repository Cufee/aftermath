// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/gamemode"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// GameModeQuery is the builder for querying GameMode entities.
type GameModeQuery struct {
	config
	ctx        *QueryContext
	order      []gamemode.OrderOption
	inters     []Interceptor
	predicates []predicate.GameMode
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GameModeQuery builder.
func (gmq *GameModeQuery) Where(ps ...predicate.GameMode) *GameModeQuery {
	gmq.predicates = append(gmq.predicates, ps...)
	return gmq
}

// Limit the number of records to be returned by this query.
func (gmq *GameModeQuery) Limit(limit int) *GameModeQuery {
	gmq.ctx.Limit = &limit
	return gmq
}

// Offset to start from.
func (gmq *GameModeQuery) Offset(offset int) *GameModeQuery {
	gmq.ctx.Offset = &offset
	return gmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gmq *GameModeQuery) Unique(unique bool) *GameModeQuery {
	gmq.ctx.Unique = &unique
	return gmq
}

// Order specifies how the records should be ordered.
func (gmq *GameModeQuery) Order(o ...gamemode.OrderOption) *GameModeQuery {
	gmq.order = append(gmq.order, o...)
	return gmq
}

// First returns the first GameMode entity from the query.
// Returns a *NotFoundError when no GameMode was found.
func (gmq *GameModeQuery) First(ctx context.Context) (*GameMode, error) {
	nodes, err := gmq.Limit(1).All(setContextOp(ctx, gmq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gamemode.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gmq *GameModeQuery) FirstX(ctx context.Context) *GameMode {
	node, err := gmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GameMode ID from the query.
// Returns a *NotFoundError when no GameMode ID was found.
func (gmq *GameModeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gmq.Limit(1).IDs(setContextOp(ctx, gmq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gamemode.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gmq *GameModeQuery) FirstIDX(ctx context.Context) string {
	id, err := gmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GameMode entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GameMode entity is found.
// Returns a *NotFoundError when no GameMode entities are found.
func (gmq *GameModeQuery) Only(ctx context.Context) (*GameMode, error) {
	nodes, err := gmq.Limit(2).All(setContextOp(ctx, gmq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gamemode.Label}
	default:
		return nil, &NotSingularError{gamemode.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gmq *GameModeQuery) OnlyX(ctx context.Context) *GameMode {
	node, err := gmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GameMode ID in the query.
// Returns a *NotSingularError when more than one GameMode ID is found.
// Returns a *NotFoundError when no entities are found.
func (gmq *GameModeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gmq.Limit(2).IDs(setContextOp(ctx, gmq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gamemode.Label}
	default:
		err = &NotSingularError{gamemode.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gmq *GameModeQuery) OnlyIDX(ctx context.Context) string {
	id, err := gmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GameModes.
func (gmq *GameModeQuery) All(ctx context.Context) ([]*GameMode, error) {
	ctx = setContextOp(ctx, gmq.ctx, "All")
	if err := gmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GameMode, *GameModeQuery]()
	return withInterceptors[[]*GameMode](ctx, gmq, qr, gmq.inters)
}

// AllX is like All, but panics if an error occurs.
func (gmq *GameModeQuery) AllX(ctx context.Context) []*GameMode {
	nodes, err := gmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GameMode IDs.
func (gmq *GameModeQuery) IDs(ctx context.Context) (ids []string, err error) {
	if gmq.ctx.Unique == nil && gmq.path != nil {
		gmq.Unique(true)
	}
	ctx = setContextOp(ctx, gmq.ctx, "IDs")
	if err = gmq.Select(gamemode.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gmq *GameModeQuery) IDsX(ctx context.Context) []string {
	ids, err := gmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gmq *GameModeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gmq.ctx, "Count")
	if err := gmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gmq, querierCount[*GameModeQuery](), gmq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (gmq *GameModeQuery) CountX(ctx context.Context) int {
	count, err := gmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gmq *GameModeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, gmq.ctx, "Exist")
	switch _, err := gmq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (gmq *GameModeQuery) ExistX(ctx context.Context) bool {
	exist, err := gmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GameModeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gmq *GameModeQuery) Clone() *GameModeQuery {
	if gmq == nil {
		return nil
	}
	return &GameModeQuery{
		config:     gmq.config,
		ctx:        gmq.ctx.Clone(),
		order:      append([]gamemode.OrderOption{}, gmq.order...),
		inters:     append([]Interceptor{}, gmq.inters...),
		predicates: append([]predicate.GameMode{}, gmq.predicates...),
		// clone intermediate query.
		sql:  gmq.sql.Clone(),
		path: gmq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GameMode.Query().
//		GroupBy(gamemode.FieldCreatedAt).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (gmq *GameModeQuery) GroupBy(field string, fields ...string) *GameModeGroupBy {
	gmq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GameModeGroupBy{build: gmq}
	grbuild.flds = &gmq.ctx.Fields
	grbuild.label = gamemode.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.GameMode.Query().
//		Select(gamemode.FieldCreatedAt).
//		Scan(ctx, &v)
func (gmq *GameModeQuery) Select(fields ...string) *GameModeSelect {
	gmq.ctx.Fields = append(gmq.ctx.Fields, fields...)
	sbuild := &GameModeSelect{GameModeQuery: gmq}
	sbuild.label = gamemode.Label
	sbuild.flds, sbuild.scan = &gmq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GameModeSelect configured with the given aggregations.
func (gmq *GameModeQuery) Aggregate(fns ...AggregateFunc) *GameModeSelect {
	return gmq.Select().Aggregate(fns...)
}

func (gmq *GameModeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gmq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gmq); err != nil {
				return err
			}
		}
	}
	for _, f := range gmq.ctx.Fields {
		if !gamemode.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if gmq.path != nil {
		prev, err := gmq.path(ctx)
		if err != nil {
			return err
		}
		gmq.sql = prev
	}
	return nil
}

func (gmq *GameModeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GameMode, error) {
	var (
		nodes = []*GameMode{}
		_spec = gmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GameMode).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GameMode{config: gmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gmq.modifiers) > 0 {
		_spec.Modifiers = gmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gmq *GameModeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gmq.querySpec()
	if len(gmq.modifiers) > 0 {
		_spec.Modifiers = gmq.modifiers
	}
	_spec.Node.Columns = gmq.ctx.Fields
	if len(gmq.ctx.Fields) > 0 {
		_spec.Unique = gmq.ctx.Unique != nil && *gmq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gmq.driver, _spec)
}

func (gmq *GameModeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(gamemode.Table, gamemode.Columns, sqlgraph.NewFieldSpec(gamemode.FieldID, field.TypeString))
	_spec.From = gmq.sql
	if unique := gmq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if gmq.path != nil {
		_spec.Unique = true
	}
	if fields := gmq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gamemode.FieldID)
		for i := range fields {
			if fields[i] != gamemode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gmq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gmq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gmq *GameModeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gmq.driver.Dialect())
	t1 := builder.Table(gamemode.Table)
	columns := gmq.ctx.Fields
	if len(columns) == 0 {
		columns = gamemode.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gmq.sql != nil {
		selector = gmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gmq.ctx.Unique != nil && *gmq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gmq.modifiers {
		m(selector)
	}
	for _, p := range gmq.predicates {
		p(selector)
	}
	for _, p := range gmq.order {
		p(selector)
	}
	if offset := gmq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gmq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gmq *GameModeQuery) Modify(modifiers ...func(s *sql.Selector)) *GameModeSelect {
	gmq.modifiers = append(gmq.modifiers, modifiers...)
	return gmq.Select()
}

// GameModeGroupBy is the group-by builder for GameMode entities.
type GameModeGroupBy struct {
	selector
	build *GameModeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gmgb *GameModeGroupBy) Aggregate(fns ...AggregateFunc) *GameModeGroupBy {
	gmgb.fns = append(gmgb.fns, fns...)
	return gmgb
}

// Scan applies the selector query and scans the result into the given value.
func (gmgb *GameModeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gmgb.build.ctx, "GroupBy")
	if err := gmgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GameModeQuery, *GameModeGroupBy](ctx, gmgb.build, gmgb, gmgb.build.inters, v)
}

func (gmgb *GameModeGroupBy) sqlScan(ctx context.Context, root *GameModeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gmgb.fns))
	for _, fn := range gmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gmgb.flds)+len(gmgb.fns))
		for _, f := range *gmgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gmgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gmgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GameModeSelect is the builder for selecting fields of GameMode entities.
type GameModeSelect struct {
	*GameModeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gms *GameModeSelect) Aggregate(fns ...AggregateFunc) *GameModeSelect {
	gms.fns = append(gms.fns, fns...)
	return gms
}

// Scan applies the selector query and scans the result into the given value.
func (gms *GameModeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gms.ctx, "Select")
	if err := gms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GameModeQuery, *GameModeSelect](ctx, gms.GameModeQuery, gms, gms.inters, v)
}

func (gms *GameModeSelect) sqlScan(ctx context.Context, root *GameModeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gms.fns))
	for _, fn := range gms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gms *GameModeSelect) Modify(modifiers ...func(s *sql.Selector)) *GameModeSelect {
	gms.modifiers = append(gms.modifiers, modifiers...)
	return gms
}