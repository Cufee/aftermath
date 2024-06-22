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
	"github.com/cufee/aftermath/internal/database/ent/db/vehicle"
)

// VehicleQuery is the builder for querying Vehicle entities.
type VehicleQuery struct {
	config
	ctx        *QueryContext
	order      []vehicle.OrderOption
	inters     []Interceptor
	predicates []predicate.Vehicle
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VehicleQuery builder.
func (vq *VehicleQuery) Where(ps ...predicate.Vehicle) *VehicleQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VehicleQuery) Limit(limit int) *VehicleQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VehicleQuery) Offset(offset int) *VehicleQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VehicleQuery) Unique(unique bool) *VehicleQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VehicleQuery) Order(o ...vehicle.OrderOption) *VehicleQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// First returns the first Vehicle entity from the query.
// Returns a *NotFoundError when no Vehicle was found.
func (vq *VehicleQuery) First(ctx context.Context) (*Vehicle, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vehicle.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VehicleQuery) FirstX(ctx context.Context) *Vehicle {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vehicle ID from the query.
// Returns a *NotFoundError when no Vehicle ID was found.
func (vq *VehicleQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vehicle.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VehicleQuery) FirstIDX(ctx context.Context) string {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vehicle entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vehicle entity is found.
// Returns a *NotFoundError when no Vehicle entities are found.
func (vq *VehicleQuery) Only(ctx context.Context) (*Vehicle, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vehicle.Label}
	default:
		return nil, &NotSingularError{vehicle.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VehicleQuery) OnlyX(ctx context.Context) *Vehicle {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vehicle ID in the query.
// Returns a *NotSingularError when more than one Vehicle ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VehicleQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vehicle.Label}
	default:
		err = &NotSingularError{vehicle.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VehicleQuery) OnlyIDX(ctx context.Context) string {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vehicles.
func (vq *VehicleQuery) All(ctx context.Context) ([]*Vehicle, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Vehicle, *VehicleQuery]()
	return withInterceptors[[]*Vehicle](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VehicleQuery) AllX(ctx context.Context) []*Vehicle {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vehicle IDs.
func (vq *VehicleQuery) IDs(ctx context.Context) (ids []string, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(vehicle.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VehicleQuery) IDsX(ctx context.Context) []string {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VehicleQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VehicleQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VehicleQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VehicleQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VehicleQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VehicleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VehicleQuery) Clone() *VehicleQuery {
	if vq == nil {
		return nil
	}
	return &VehicleQuery{
		config:     vq.config,
		ctx:        vq.ctx.Clone(),
		order:      append([]vehicle.OrderOption{}, vq.order...),
		inters:     append([]Interceptor{}, vq.inters...),
		predicates: append([]predicate.Vehicle{}, vq.predicates...),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt int `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vehicle.Query().
//		GroupBy(vehicle.FieldCreatedAt).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (vq *VehicleQuery) GroupBy(field string, fields ...string) *VehicleGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VehicleGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = vehicle.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt int `json:"created_at,omitempty"`
//	}
//
//	client.Vehicle.Query().
//		Select(vehicle.FieldCreatedAt).
//		Scan(ctx, &v)
func (vq *VehicleQuery) Select(fields ...string) *VehicleSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VehicleSelect{VehicleQuery: vq}
	sbuild.label = vehicle.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VehicleSelect configured with the given aggregations.
func (vq *VehicleQuery) Aggregate(fns ...AggregateFunc) *VehicleSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VehicleQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !vehicle.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VehicleQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Vehicle, error) {
	var (
		nodes = []*Vehicle{}
		_spec = vq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Vehicle).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Vehicle{config: vq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vq *VehicleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VehicleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(vehicle.Table, vehicle.Columns, sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeString))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehicle.FieldID)
		for i := range fields {
			if fields[i] != vehicle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VehicleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vehicle.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = vehicle.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VehicleGroupBy is the group-by builder for Vehicle entities.
type VehicleGroupBy struct {
	selector
	build *VehicleQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VehicleGroupBy) Aggregate(fns ...AggregateFunc) *VehicleGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VehicleGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VehicleQuery, *VehicleGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VehicleGroupBy) sqlScan(ctx context.Context, root *VehicleQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VehicleSelect is the builder for selecting fields of Vehicle entities.
type VehicleSelect struct {
	*VehicleQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VehicleSelect) Aggregate(fns ...AggregateFunc) *VehicleSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VehicleSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VehicleQuery, *VehicleSelect](ctx, vs.VehicleQuery, vs, vs.inters, v)
}

func (vs *VehicleSelect) sqlScan(ctx context.Context, root *VehicleQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
