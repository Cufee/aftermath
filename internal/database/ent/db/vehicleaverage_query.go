// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/vehicleaverage"
)

// VehicleAverageQuery is the builder for querying VehicleAverage entities.
type VehicleAverageQuery struct {
	config
	ctx        *QueryContext
	order      []vehicleaverage.OrderOption
	inters     []Interceptor
	predicates []predicate.VehicleAverage
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VehicleAverageQuery builder.
func (vaq *VehicleAverageQuery) Where(ps ...predicate.VehicleAverage) *VehicleAverageQuery {
	vaq.predicates = append(vaq.predicates, ps...)
	return vaq
}

// Limit the number of records to be returned by this query.
func (vaq *VehicleAverageQuery) Limit(limit int) *VehicleAverageQuery {
	vaq.ctx.Limit = &limit
	return vaq
}

// Offset to start from.
func (vaq *VehicleAverageQuery) Offset(offset int) *VehicleAverageQuery {
	vaq.ctx.Offset = &offset
	return vaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vaq *VehicleAverageQuery) Unique(unique bool) *VehicleAverageQuery {
	vaq.ctx.Unique = &unique
	return vaq
}

// Order specifies how the records should be ordered.
func (vaq *VehicleAverageQuery) Order(o ...vehicleaverage.OrderOption) *VehicleAverageQuery {
	vaq.order = append(vaq.order, o...)
	return vaq
}

// First returns the first VehicleAverage entity from the query.
// Returns a *NotFoundError when no VehicleAverage was found.
func (vaq *VehicleAverageQuery) First(ctx context.Context) (*VehicleAverage, error) {
	nodes, err := vaq.Limit(1).All(setContextOp(ctx, vaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vehicleaverage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vaq *VehicleAverageQuery) FirstX(ctx context.Context) *VehicleAverage {
	node, err := vaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VehicleAverage ID from the query.
// Returns a *NotFoundError when no VehicleAverage ID was found.
func (vaq *VehicleAverageQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vaq.Limit(1).IDs(setContextOp(ctx, vaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vehicleaverage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vaq *VehicleAverageQuery) FirstIDX(ctx context.Context) string {
	id, err := vaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VehicleAverage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VehicleAverage entity is found.
// Returns a *NotFoundError when no VehicleAverage entities are found.
func (vaq *VehicleAverageQuery) Only(ctx context.Context) (*VehicleAverage, error) {
	nodes, err := vaq.Limit(2).All(setContextOp(ctx, vaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vehicleaverage.Label}
	default:
		return nil, &NotSingularError{vehicleaverage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vaq *VehicleAverageQuery) OnlyX(ctx context.Context) *VehicleAverage {
	node, err := vaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VehicleAverage ID in the query.
// Returns a *NotSingularError when more than one VehicleAverage ID is found.
// Returns a *NotFoundError when no entities are found.
func (vaq *VehicleAverageQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vaq.Limit(2).IDs(setContextOp(ctx, vaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vehicleaverage.Label}
	default:
		err = &NotSingularError{vehicleaverage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vaq *VehicleAverageQuery) OnlyIDX(ctx context.Context) string {
	id, err := vaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VehicleAverages.
func (vaq *VehicleAverageQuery) All(ctx context.Context) ([]*VehicleAverage, error) {
	ctx = setContextOp(ctx, vaq.ctx, "All")
	if err := vaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*VehicleAverage, *VehicleAverageQuery]()
	return withInterceptors[[]*VehicleAverage](ctx, vaq, qr, vaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vaq *VehicleAverageQuery) AllX(ctx context.Context) []*VehicleAverage {
	nodes, err := vaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VehicleAverage IDs.
func (vaq *VehicleAverageQuery) IDs(ctx context.Context) (ids []string, err error) {
	if vaq.ctx.Unique == nil && vaq.path != nil {
		vaq.Unique(true)
	}
	ctx = setContextOp(ctx, vaq.ctx, "IDs")
	if err = vaq.Select(vehicleaverage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vaq *VehicleAverageQuery) IDsX(ctx context.Context) []string {
	ids, err := vaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vaq *VehicleAverageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vaq.ctx, "Count")
	if err := vaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vaq, querierCount[*VehicleAverageQuery](), vaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vaq *VehicleAverageQuery) CountX(ctx context.Context) int {
	count, err := vaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vaq *VehicleAverageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vaq.ctx, "Exist")
	switch _, err := vaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vaq *VehicleAverageQuery) ExistX(ctx context.Context) bool {
	exist, err := vaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VehicleAverageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vaq *VehicleAverageQuery) Clone() *VehicleAverageQuery {
	if vaq == nil {
		return nil
	}
	return &VehicleAverageQuery{
		config:     vaq.config,
		ctx:        vaq.ctx.Clone(),
		order:      append([]vehicleaverage.OrderOption{}, vaq.order...),
		inters:     append([]Interceptor{}, vaq.inters...),
		predicates: append([]predicate.VehicleAverage{}, vaq.predicates...),
		// clone intermediate query.
		sql:  vaq.sql.Clone(),
		path: vaq.path,
	}
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
//	client.VehicleAverage.Query().
//		GroupBy(vehicleaverage.FieldCreatedAt).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (vaq *VehicleAverageQuery) GroupBy(field string, fields ...string) *VehicleAverageGroupBy {
	vaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VehicleAverageGroupBy{build: vaq}
	grbuild.flds = &vaq.ctx.Fields
	grbuild.label = vehicleaverage.Label
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
//	client.VehicleAverage.Query().
//		Select(vehicleaverage.FieldCreatedAt).
//		Scan(ctx, &v)
func (vaq *VehicleAverageQuery) Select(fields ...string) *VehicleAverageSelect {
	vaq.ctx.Fields = append(vaq.ctx.Fields, fields...)
	sbuild := &VehicleAverageSelect{VehicleAverageQuery: vaq}
	sbuild.label = vehicleaverage.Label
	sbuild.flds, sbuild.scan = &vaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VehicleAverageSelect configured with the given aggregations.
func (vaq *VehicleAverageQuery) Aggregate(fns ...AggregateFunc) *VehicleAverageSelect {
	return vaq.Select().Aggregate(fns...)
}

func (vaq *VehicleAverageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vaq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vaq); err != nil {
				return err
			}
		}
	}
	for _, f := range vaq.ctx.Fields {
		if !vehicleaverage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if vaq.path != nil {
		prev, err := vaq.path(ctx)
		if err != nil {
			return err
		}
		vaq.sql = prev
	}
	return nil
}

func (vaq *VehicleAverageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VehicleAverage, error) {
	var (
		nodes = []*VehicleAverage{}
		_spec = vaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*VehicleAverage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &VehicleAverage{config: vaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vaq *VehicleAverageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vaq.querySpec()
	_spec.Node.Columns = vaq.ctx.Fields
	if len(vaq.ctx.Fields) > 0 {
		_spec.Unique = vaq.ctx.Unique != nil && *vaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vaq.driver, _spec)
}

func (vaq *VehicleAverageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(vehicleaverage.Table, vehicleaverage.Columns, sqlgraph.NewFieldSpec(vehicleaverage.FieldID, field.TypeString))
	_spec.From = vaq.sql
	if unique := vaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vaq.path != nil {
		_spec.Unique = true
	}
	if fields := vaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehicleaverage.FieldID)
		for i := range fields {
			if fields[i] != vehicleaverage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vaq *VehicleAverageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vaq.driver.Dialect())
	t1 := builder.Table(vehicleaverage.Table)
	columns := vaq.ctx.Fields
	if len(columns) == 0 {
		columns = vehicleaverage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vaq.sql != nil {
		selector = vaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vaq.ctx.Unique != nil && *vaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vaq.predicates {
		p(selector)
	}
	for _, p := range vaq.order {
		p(selector)
	}
	if offset := vaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VehicleAverageGroupBy is the group-by builder for VehicleAverage entities.
type VehicleAverageGroupBy struct {
	selector
	build *VehicleAverageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vagb *VehicleAverageGroupBy) Aggregate(fns ...AggregateFunc) *VehicleAverageGroupBy {
	vagb.fns = append(vagb.fns, fns...)
	return vagb
}

// Scan applies the selector query and scans the result into the given value.
func (vagb *VehicleAverageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vagb.build.ctx, "GroupBy")
	if err := vagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VehicleAverageQuery, *VehicleAverageGroupBy](ctx, vagb.build, vagb, vagb.build.inters, v)
}

func (vagb *VehicleAverageGroupBy) sqlScan(ctx context.Context, root *VehicleAverageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vagb.fns))
	for _, fn := range vagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vagb.flds)+len(vagb.fns))
		for _, f := range *vagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VehicleAverageSelect is the builder for selecting fields of VehicleAverage entities.
type VehicleAverageSelect struct {
	*VehicleAverageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vas *VehicleAverageSelect) Aggregate(fns ...AggregateFunc) *VehicleAverageSelect {
	vas.fns = append(vas.fns, fns...)
	return vas
}

// Scan applies the selector query and scans the result into the given value.
func (vas *VehicleAverageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vas.ctx, "Select")
	if err := vas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VehicleAverageQuery, *VehicleAverageSelect](ctx, vas.VehicleAverageQuery, vas, vas.inters, v)
}

func (vas *VehicleAverageSelect) sqlScan(ctx context.Context, root *VehicleAverageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vas.fns))
	for _, fn := range vas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}