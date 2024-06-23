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
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/ent/db/userconnection"
)

// UserConnectionQuery is the builder for querying UserConnection entities.
type UserConnectionQuery struct {
	config
	ctx        *QueryContext
	order      []userconnection.OrderOption
	inters     []Interceptor
	predicates []predicate.UserConnection
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserConnectionQuery builder.
func (ucq *UserConnectionQuery) Where(ps ...predicate.UserConnection) *UserConnectionQuery {
	ucq.predicates = append(ucq.predicates, ps...)
	return ucq
}

// Limit the number of records to be returned by this query.
func (ucq *UserConnectionQuery) Limit(limit int) *UserConnectionQuery {
	ucq.ctx.Limit = &limit
	return ucq
}

// Offset to start from.
func (ucq *UserConnectionQuery) Offset(offset int) *UserConnectionQuery {
	ucq.ctx.Offset = &offset
	return ucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucq *UserConnectionQuery) Unique(unique bool) *UserConnectionQuery {
	ucq.ctx.Unique = &unique
	return ucq
}

// Order specifies how the records should be ordered.
func (ucq *UserConnectionQuery) Order(o ...userconnection.OrderOption) *UserConnectionQuery {
	ucq.order = append(ucq.order, o...)
	return ucq
}

// QueryUser chains the current query on the "user" edge.
func (ucq *UserConnectionQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: ucq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userconnection.Table, userconnection.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userconnection.UserTable, userconnection.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserConnection entity from the query.
// Returns a *NotFoundError when no UserConnection was found.
func (ucq *UserConnectionQuery) First(ctx context.Context) (*UserConnection, error) {
	nodes, err := ucq.Limit(1).All(setContextOp(ctx, ucq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userconnection.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucq *UserConnectionQuery) FirstX(ctx context.Context) *UserConnection {
	node, err := ucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserConnection ID from the query.
// Returns a *NotFoundError when no UserConnection ID was found.
func (ucq *UserConnectionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ucq.Limit(1).IDs(setContextOp(ctx, ucq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userconnection.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucq *UserConnectionQuery) FirstIDX(ctx context.Context) string {
	id, err := ucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserConnection entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserConnection entity is found.
// Returns a *NotFoundError when no UserConnection entities are found.
func (ucq *UserConnectionQuery) Only(ctx context.Context) (*UserConnection, error) {
	nodes, err := ucq.Limit(2).All(setContextOp(ctx, ucq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userconnection.Label}
	default:
		return nil, &NotSingularError{userconnection.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucq *UserConnectionQuery) OnlyX(ctx context.Context) *UserConnection {
	node, err := ucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserConnection ID in the query.
// Returns a *NotSingularError when more than one UserConnection ID is found.
// Returns a *NotFoundError when no entities are found.
func (ucq *UserConnectionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ucq.Limit(2).IDs(setContextOp(ctx, ucq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userconnection.Label}
	default:
		err = &NotSingularError{userconnection.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucq *UserConnectionQuery) OnlyIDX(ctx context.Context) string {
	id, err := ucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserConnections.
func (ucq *UserConnectionQuery) All(ctx context.Context) ([]*UserConnection, error) {
	ctx = setContextOp(ctx, ucq.ctx, "All")
	if err := ucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserConnection, *UserConnectionQuery]()
	return withInterceptors[[]*UserConnection](ctx, ucq, qr, ucq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ucq *UserConnectionQuery) AllX(ctx context.Context) []*UserConnection {
	nodes, err := ucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserConnection IDs.
func (ucq *UserConnectionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if ucq.ctx.Unique == nil && ucq.path != nil {
		ucq.Unique(true)
	}
	ctx = setContextOp(ctx, ucq.ctx, "IDs")
	if err = ucq.Select(userconnection.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucq *UserConnectionQuery) IDsX(ctx context.Context) []string {
	ids, err := ucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucq *UserConnectionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ucq.ctx, "Count")
	if err := ucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ucq, querierCount[*UserConnectionQuery](), ucq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ucq *UserConnectionQuery) CountX(ctx context.Context) int {
	count, err := ucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucq *UserConnectionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ucq.ctx, "Exist")
	switch _, err := ucq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ucq *UserConnectionQuery) ExistX(ctx context.Context) bool {
	exist, err := ucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserConnectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucq *UserConnectionQuery) Clone() *UserConnectionQuery {
	if ucq == nil {
		return nil
	}
	return &UserConnectionQuery{
		config:     ucq.config,
		ctx:        ucq.ctx.Clone(),
		order:      append([]userconnection.OrderOption{}, ucq.order...),
		inters:     append([]Interceptor{}, ucq.inters...),
		predicates: append([]predicate.UserConnection{}, ucq.predicates...),
		withUser:   ucq.withUser.Clone(),
		// clone intermediate query.
		sql:  ucq.sql.Clone(),
		path: ucq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserConnectionQuery) WithUser(opts ...func(*UserQuery)) *UserConnectionQuery {
	query := (&UserClient{config: ucq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ucq.withUser = query
	return ucq
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
//	client.UserConnection.Query().
//		GroupBy(userconnection.FieldCreatedAt).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (ucq *UserConnectionQuery) GroupBy(field string, fields ...string) *UserConnectionGroupBy {
	ucq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserConnectionGroupBy{build: ucq}
	grbuild.flds = &ucq.ctx.Fields
	grbuild.label = userconnection.Label
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
//	client.UserConnection.Query().
//		Select(userconnection.FieldCreatedAt).
//		Scan(ctx, &v)
func (ucq *UserConnectionQuery) Select(fields ...string) *UserConnectionSelect {
	ucq.ctx.Fields = append(ucq.ctx.Fields, fields...)
	sbuild := &UserConnectionSelect{UserConnectionQuery: ucq}
	sbuild.label = userconnection.Label
	sbuild.flds, sbuild.scan = &ucq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserConnectionSelect configured with the given aggregations.
func (ucq *UserConnectionQuery) Aggregate(fns ...AggregateFunc) *UserConnectionSelect {
	return ucq.Select().Aggregate(fns...)
}

func (ucq *UserConnectionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ucq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ucq); err != nil {
				return err
			}
		}
	}
	for _, f := range ucq.ctx.Fields {
		if !userconnection.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if ucq.path != nil {
		prev, err := ucq.path(ctx)
		if err != nil {
			return err
		}
		ucq.sql = prev
	}
	return nil
}

func (ucq *UserConnectionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserConnection, error) {
	var (
		nodes       = []*UserConnection{}
		_spec       = ucq.querySpec()
		loadedTypes = [1]bool{
			ucq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserConnection).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserConnection{config: ucq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ucq.withUser; query != nil {
		if err := ucq.loadUser(ctx, query, nodes, nil,
			func(n *UserConnection, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ucq *UserConnectionQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserConnection, init func(*UserConnection), assign func(*UserConnection, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserConnection)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ucq *UserConnectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucq.querySpec()
	_spec.Node.Columns = ucq.ctx.Fields
	if len(ucq.ctx.Fields) > 0 {
		_spec.Unique = ucq.ctx.Unique != nil && *ucq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ucq.driver, _spec)
}

func (ucq *UserConnectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userconnection.Table, userconnection.Columns, sqlgraph.NewFieldSpec(userconnection.FieldID, field.TypeString))
	_spec.From = ucq.sql
	if unique := ucq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ucq.path != nil {
		_spec.Unique = true
	}
	if fields := ucq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userconnection.FieldID)
		for i := range fields {
			if fields[i] != userconnection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ucq.withUser != nil {
			_spec.Node.AddColumnOnce(userconnection.FieldUserID)
		}
	}
	if ps := ucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucq *UserConnectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucq.driver.Dialect())
	t1 := builder.Table(userconnection.Table)
	columns := ucq.ctx.Fields
	if len(columns) == 0 {
		columns = userconnection.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucq.sql != nil {
		selector = ucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ucq.ctx.Unique != nil && *ucq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ucq.predicates {
		p(selector)
	}
	for _, p := range ucq.order {
		p(selector)
	}
	if offset := ucq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserConnectionGroupBy is the group-by builder for UserConnection entities.
type UserConnectionGroupBy struct {
	selector
	build *UserConnectionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucgb *UserConnectionGroupBy) Aggregate(fns ...AggregateFunc) *UserConnectionGroupBy {
	ucgb.fns = append(ucgb.fns, fns...)
	return ucgb
}

// Scan applies the selector query and scans the result into the given value.
func (ucgb *UserConnectionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucgb.build.ctx, "GroupBy")
	if err := ucgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserConnectionQuery, *UserConnectionGroupBy](ctx, ucgb.build, ucgb, ucgb.build.inters, v)
}

func (ucgb *UserConnectionGroupBy) sqlScan(ctx context.Context, root *UserConnectionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ucgb.fns))
	for _, fn := range ucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ucgb.flds)+len(ucgb.fns))
		for _, f := range *ucgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ucgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserConnectionSelect is the builder for selecting fields of UserConnection entities.
type UserConnectionSelect struct {
	*UserConnectionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ucs *UserConnectionSelect) Aggregate(fns ...AggregateFunc) *UserConnectionSelect {
	ucs.fns = append(ucs.fns, fns...)
	return ucs
}

// Scan applies the selector query and scans the result into the given value.
func (ucs *UserConnectionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucs.ctx, "Select")
	if err := ucs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserConnectionQuery, *UserConnectionSelect](ctx, ucs.UserConnectionQuery, ucs, ucs.inters, v)
}

func (ucs *UserConnectionSelect) sqlScan(ctx context.Context, root *UserConnectionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ucs.fns))
	for _, fn := range ucs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ucs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}