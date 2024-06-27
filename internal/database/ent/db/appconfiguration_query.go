// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/appconfiguration"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// AppConfigurationQuery is the builder for querying AppConfiguration entities.
type AppConfigurationQuery struct {
	config
	ctx        *QueryContext
	order      []appconfiguration.OrderOption
	inters     []Interceptor
	predicates []predicate.AppConfiguration
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppConfigurationQuery builder.
func (acq *AppConfigurationQuery) Where(ps ...predicate.AppConfiguration) *AppConfigurationQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *AppConfigurationQuery) Limit(limit int) *AppConfigurationQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *AppConfigurationQuery) Offset(offset int) *AppConfigurationQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *AppConfigurationQuery) Unique(unique bool) *AppConfigurationQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *AppConfigurationQuery) Order(o ...appconfiguration.OrderOption) *AppConfigurationQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// First returns the first AppConfiguration entity from the query.
// Returns a *NotFoundError when no AppConfiguration was found.
func (acq *AppConfigurationQuery) First(ctx context.Context) (*AppConfiguration, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appconfiguration.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *AppConfigurationQuery) FirstX(ctx context.Context) *AppConfiguration {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppConfiguration ID from the query.
// Returns a *NotFoundError when no AppConfiguration ID was found.
func (acq *AppConfigurationQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appconfiguration.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *AppConfigurationQuery) FirstIDX(ctx context.Context) string {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppConfiguration entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppConfiguration entity is found.
// Returns a *NotFoundError when no AppConfiguration entities are found.
func (acq *AppConfigurationQuery) Only(ctx context.Context) (*AppConfiguration, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appconfiguration.Label}
	default:
		return nil, &NotSingularError{appconfiguration.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *AppConfigurationQuery) OnlyX(ctx context.Context) *AppConfiguration {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppConfiguration ID in the query.
// Returns a *NotSingularError when more than one AppConfiguration ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *AppConfigurationQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appconfiguration.Label}
	default:
		err = &NotSingularError{appconfiguration.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *AppConfigurationQuery) OnlyIDX(ctx context.Context) string {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppConfigurations.
func (acq *AppConfigurationQuery) All(ctx context.Context) ([]*AppConfiguration, error) {
	ctx = setContextOp(ctx, acq.ctx, "All")
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AppConfiguration, *AppConfigurationQuery]()
	return withInterceptors[[]*AppConfiguration](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *AppConfigurationQuery) AllX(ctx context.Context) []*AppConfiguration {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppConfiguration IDs.
func (acq *AppConfigurationQuery) IDs(ctx context.Context) (ids []string, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, "IDs")
	if err = acq.Select(appconfiguration.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *AppConfigurationQuery) IDsX(ctx context.Context) []string {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *AppConfigurationQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, "Count")
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*AppConfigurationQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *AppConfigurationQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *AppConfigurationQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, "Exist")
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *AppConfigurationQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppConfigurationQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *AppConfigurationQuery) Clone() *AppConfigurationQuery {
	if acq == nil {
		return nil
	}
	return &AppConfigurationQuery{
		config:     acq.config,
		ctx:        acq.ctx.Clone(),
		order:      append([]appconfiguration.OrderOption{}, acq.order...),
		inters:     append([]Interceptor{}, acq.inters...),
		predicates: append([]predicate.AppConfiguration{}, acq.predicates...),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
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
//	client.AppConfiguration.Query().
//		GroupBy(appconfiguration.FieldCreatedAt).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (acq *AppConfigurationQuery) GroupBy(field string, fields ...string) *AppConfigurationGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AppConfigurationGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = appconfiguration.Label
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
//	client.AppConfiguration.Query().
//		Select(appconfiguration.FieldCreatedAt).
//		Scan(ctx, &v)
func (acq *AppConfigurationQuery) Select(fields ...string) *AppConfigurationSelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &AppConfigurationSelect{AppConfigurationQuery: acq}
	sbuild.label = appconfiguration.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AppConfigurationSelect configured with the given aggregations.
func (acq *AppConfigurationQuery) Aggregate(fns ...AggregateFunc) *AppConfigurationSelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *AppConfigurationQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !appconfiguration.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *AppConfigurationQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppConfiguration, error) {
	var (
		nodes = []*AppConfiguration{}
		_spec = acq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AppConfiguration).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AppConfiguration{config: acq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (acq *AppConfigurationQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *AppConfigurationQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(appconfiguration.Table, appconfiguration.Columns, sqlgraph.NewFieldSpec(appconfiguration.FieldID, field.TypeString))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appconfiguration.FieldID)
		for i := range fields {
			if fields[i] != appconfiguration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *AppConfigurationQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(appconfiguration.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = appconfiguration.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range acq.modifiers {
		m(selector)
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (acq *AppConfigurationQuery) Modify(modifiers ...func(s *sql.Selector)) *AppConfigurationSelect {
	acq.modifiers = append(acq.modifiers, modifiers...)
	return acq.Select()
}

// AppConfigurationGroupBy is the group-by builder for AppConfiguration entities.
type AppConfigurationGroupBy struct {
	selector
	build *AppConfigurationQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *AppConfigurationGroupBy) Aggregate(fns ...AggregateFunc) *AppConfigurationGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *AppConfigurationGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, "GroupBy")
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppConfigurationQuery, *AppConfigurationGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *AppConfigurationGroupBy) sqlScan(ctx context.Context, root *AppConfigurationQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AppConfigurationSelect is the builder for selecting fields of AppConfiguration entities.
type AppConfigurationSelect struct {
	*AppConfigurationQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *AppConfigurationSelect) Aggregate(fns ...AggregateFunc) *AppConfigurationSelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *AppConfigurationSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, "Select")
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AppConfigurationQuery, *AppConfigurationSelect](ctx, acs.AppConfigurationQuery, acs, acs.inters, v)
}

func (acs *AppConfigurationSelect) sqlScan(ctx context.Context, root *AppConfigurationQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (acs *AppConfigurationSelect) Modify(modifiers ...func(s *sql.Selector)) *AppConfigurationSelect {
	acs.modifiers = append(acs.modifiers, modifiers...)
	return acs
}
