// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/moderationrequest"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
)

// ModerationRequestQuery is the builder for querying ModerationRequest entities.
type ModerationRequestQuery struct {
	config
	ctx           *QueryContext
	order         []moderationrequest.OrderOption
	inters        []Interceptor
	predicates    []predicate.ModerationRequest
	withModerator *UserQuery
	withRequestor *UserQuery
	modifiers     []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ModerationRequestQuery builder.
func (mrq *ModerationRequestQuery) Where(ps ...predicate.ModerationRequest) *ModerationRequestQuery {
	mrq.predicates = append(mrq.predicates, ps...)
	return mrq
}

// Limit the number of records to be returned by this query.
func (mrq *ModerationRequestQuery) Limit(limit int) *ModerationRequestQuery {
	mrq.ctx.Limit = &limit
	return mrq
}

// Offset to start from.
func (mrq *ModerationRequestQuery) Offset(offset int) *ModerationRequestQuery {
	mrq.ctx.Offset = &offset
	return mrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mrq *ModerationRequestQuery) Unique(unique bool) *ModerationRequestQuery {
	mrq.ctx.Unique = &unique
	return mrq
}

// Order specifies how the records should be ordered.
func (mrq *ModerationRequestQuery) Order(o ...moderationrequest.OrderOption) *ModerationRequestQuery {
	mrq.order = append(mrq.order, o...)
	return mrq
}

// QueryModerator chains the current query on the "moderator" edge.
func (mrq *ModerationRequestQuery) QueryModerator() *UserQuery {
	query := (&UserClient{config: mrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moderationrequest.Table, moderationrequest.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, moderationrequest.ModeratorTable, moderationrequest.ModeratorColumn),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRequestor chains the current query on the "requestor" edge.
func (mrq *ModerationRequestQuery) QueryRequestor() *UserQuery {
	query := (&UserClient{config: mrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(moderationrequest.Table, moderationrequest.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, moderationrequest.RequestorTable, moderationrequest.RequestorColumn),
		)
		fromU = sqlgraph.SetNeighbors(mrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ModerationRequest entity from the query.
// Returns a *NotFoundError when no ModerationRequest was found.
func (mrq *ModerationRequestQuery) First(ctx context.Context) (*ModerationRequest, error) {
	nodes, err := mrq.Limit(1).All(setContextOp(ctx, mrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{moderationrequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mrq *ModerationRequestQuery) FirstX(ctx context.Context) *ModerationRequest {
	node, err := mrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ModerationRequest ID from the query.
// Returns a *NotFoundError when no ModerationRequest ID was found.
func (mrq *ModerationRequestQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mrq.Limit(1).IDs(setContextOp(ctx, mrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{moderationrequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mrq *ModerationRequestQuery) FirstIDX(ctx context.Context) string {
	id, err := mrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ModerationRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ModerationRequest entity is found.
// Returns a *NotFoundError when no ModerationRequest entities are found.
func (mrq *ModerationRequestQuery) Only(ctx context.Context) (*ModerationRequest, error) {
	nodes, err := mrq.Limit(2).All(setContextOp(ctx, mrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{moderationrequest.Label}
	default:
		return nil, &NotSingularError{moderationrequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mrq *ModerationRequestQuery) OnlyX(ctx context.Context) *ModerationRequest {
	node, err := mrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ModerationRequest ID in the query.
// Returns a *NotSingularError when more than one ModerationRequest ID is found.
// Returns a *NotFoundError when no entities are found.
func (mrq *ModerationRequestQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mrq.Limit(2).IDs(setContextOp(ctx, mrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{moderationrequest.Label}
	default:
		err = &NotSingularError{moderationrequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mrq *ModerationRequestQuery) OnlyIDX(ctx context.Context) string {
	id, err := mrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ModerationRequests.
func (mrq *ModerationRequestQuery) All(ctx context.Context) ([]*ModerationRequest, error) {
	ctx = setContextOp(ctx, mrq.ctx, "All")
	if err := mrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ModerationRequest, *ModerationRequestQuery]()
	return withInterceptors[[]*ModerationRequest](ctx, mrq, qr, mrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mrq *ModerationRequestQuery) AllX(ctx context.Context) []*ModerationRequest {
	nodes, err := mrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ModerationRequest IDs.
func (mrq *ModerationRequestQuery) IDs(ctx context.Context) (ids []string, err error) {
	if mrq.ctx.Unique == nil && mrq.path != nil {
		mrq.Unique(true)
	}
	ctx = setContextOp(ctx, mrq.ctx, "IDs")
	if err = mrq.Select(moderationrequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mrq *ModerationRequestQuery) IDsX(ctx context.Context) []string {
	ids, err := mrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mrq *ModerationRequestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mrq.ctx, "Count")
	if err := mrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mrq, querierCount[*ModerationRequestQuery](), mrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mrq *ModerationRequestQuery) CountX(ctx context.Context) int {
	count, err := mrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mrq *ModerationRequestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mrq.ctx, "Exist")
	switch _, err := mrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mrq *ModerationRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := mrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ModerationRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mrq *ModerationRequestQuery) Clone() *ModerationRequestQuery {
	if mrq == nil {
		return nil
	}
	return &ModerationRequestQuery{
		config:        mrq.config,
		ctx:           mrq.ctx.Clone(),
		order:         append([]moderationrequest.OrderOption{}, mrq.order...),
		inters:        append([]Interceptor{}, mrq.inters...),
		predicates:    append([]predicate.ModerationRequest{}, mrq.predicates...),
		withModerator: mrq.withModerator.Clone(),
		withRequestor: mrq.withRequestor.Clone(),
		// clone intermediate query.
		sql:  mrq.sql.Clone(),
		path: mrq.path,
	}
}

// WithModerator tells the query-builder to eager-load the nodes that are connected to
// the "moderator" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *ModerationRequestQuery) WithModerator(opts ...func(*UserQuery)) *ModerationRequestQuery {
	query := (&UserClient{config: mrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mrq.withModerator = query
	return mrq
}

// WithRequestor tells the query-builder to eager-load the nodes that are connected to
// the "requestor" edge. The optional arguments are used to configure the query builder of the edge.
func (mrq *ModerationRequestQuery) WithRequestor(opts ...func(*UserQuery)) *ModerationRequestQuery {
	query := (&UserClient{config: mrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mrq.withRequestor = query
	return mrq
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
//	client.ModerationRequest.Query().
//		GroupBy(moderationrequest.FieldCreatedAt).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (mrq *ModerationRequestQuery) GroupBy(field string, fields ...string) *ModerationRequestGroupBy {
	mrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ModerationRequestGroupBy{build: mrq}
	grbuild.flds = &mrq.ctx.Fields
	grbuild.label = moderationrequest.Label
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
//	client.ModerationRequest.Query().
//		Select(moderationrequest.FieldCreatedAt).
//		Scan(ctx, &v)
func (mrq *ModerationRequestQuery) Select(fields ...string) *ModerationRequestSelect {
	mrq.ctx.Fields = append(mrq.ctx.Fields, fields...)
	sbuild := &ModerationRequestSelect{ModerationRequestQuery: mrq}
	sbuild.label = moderationrequest.Label
	sbuild.flds, sbuild.scan = &mrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ModerationRequestSelect configured with the given aggregations.
func (mrq *ModerationRequestQuery) Aggregate(fns ...AggregateFunc) *ModerationRequestSelect {
	return mrq.Select().Aggregate(fns...)
}

func (mrq *ModerationRequestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mrq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mrq); err != nil {
				return err
			}
		}
	}
	for _, f := range mrq.ctx.Fields {
		if !moderationrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if mrq.path != nil {
		prev, err := mrq.path(ctx)
		if err != nil {
			return err
		}
		mrq.sql = prev
	}
	return nil
}

func (mrq *ModerationRequestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ModerationRequest, error) {
	var (
		nodes       = []*ModerationRequest{}
		_spec       = mrq.querySpec()
		loadedTypes = [2]bool{
			mrq.withModerator != nil,
			mrq.withRequestor != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ModerationRequest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ModerationRequest{config: mrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(mrq.modifiers) > 0 {
		_spec.Modifiers = mrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mrq.withModerator; query != nil {
		if err := mrq.loadModerator(ctx, query, nodes, nil,
			func(n *ModerationRequest, e *User) { n.Edges.Moderator = e }); err != nil {
			return nil, err
		}
	}
	if query := mrq.withRequestor; query != nil {
		if err := mrq.loadRequestor(ctx, query, nodes, nil,
			func(n *ModerationRequest, e *User) { n.Edges.Requestor = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mrq *ModerationRequestQuery) loadModerator(ctx context.Context, query *UserQuery, nodes []*ModerationRequest, init func(*ModerationRequest), assign func(*ModerationRequest, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ModerationRequest)
	for i := range nodes {
		if nodes[i].ModeratorID == nil {
			continue
		}
		fk := *nodes[i].ModeratorID
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
			return fmt.Errorf(`unexpected foreign-key "moderator_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mrq *ModerationRequestQuery) loadRequestor(ctx context.Context, query *UserQuery, nodes []*ModerationRequest, init func(*ModerationRequest), assign func(*ModerationRequest, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*ModerationRequest)
	for i := range nodes {
		fk := nodes[i].RequestorID
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
			return fmt.Errorf(`unexpected foreign-key "requestor_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mrq *ModerationRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mrq.querySpec()
	if len(mrq.modifiers) > 0 {
		_spec.Modifiers = mrq.modifiers
	}
	_spec.Node.Columns = mrq.ctx.Fields
	if len(mrq.ctx.Fields) > 0 {
		_spec.Unique = mrq.ctx.Unique != nil && *mrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mrq.driver, _spec)
}

func (mrq *ModerationRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(moderationrequest.Table, moderationrequest.Columns, sqlgraph.NewFieldSpec(moderationrequest.FieldID, field.TypeString))
	_spec.From = mrq.sql
	if unique := mrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mrq.path != nil {
		_spec.Unique = true
	}
	if fields := mrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, moderationrequest.FieldID)
		for i := range fields {
			if fields[i] != moderationrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mrq.withModerator != nil {
			_spec.Node.AddColumnOnce(moderationrequest.FieldModeratorID)
		}
		if mrq.withRequestor != nil {
			_spec.Node.AddColumnOnce(moderationrequest.FieldRequestorID)
		}
	}
	if ps := mrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mrq *ModerationRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mrq.driver.Dialect())
	t1 := builder.Table(moderationrequest.Table)
	columns := mrq.ctx.Fields
	if len(columns) == 0 {
		columns = moderationrequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mrq.sql != nil {
		selector = mrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mrq.ctx.Unique != nil && *mrq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range mrq.modifiers {
		m(selector)
	}
	for _, p := range mrq.predicates {
		p(selector)
	}
	for _, p := range mrq.order {
		p(selector)
	}
	if offset := mrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mrq *ModerationRequestQuery) Modify(modifiers ...func(s *sql.Selector)) *ModerationRequestSelect {
	mrq.modifiers = append(mrq.modifiers, modifiers...)
	return mrq.Select()
}

// ModerationRequestGroupBy is the group-by builder for ModerationRequest entities.
type ModerationRequestGroupBy struct {
	selector
	build *ModerationRequestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mrgb *ModerationRequestGroupBy) Aggregate(fns ...AggregateFunc) *ModerationRequestGroupBy {
	mrgb.fns = append(mrgb.fns, fns...)
	return mrgb
}

// Scan applies the selector query and scans the result into the given value.
func (mrgb *ModerationRequestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mrgb.build.ctx, "GroupBy")
	if err := mrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ModerationRequestQuery, *ModerationRequestGroupBy](ctx, mrgb.build, mrgb, mrgb.build.inters, v)
}

func (mrgb *ModerationRequestGroupBy) sqlScan(ctx context.Context, root *ModerationRequestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mrgb.fns))
	for _, fn := range mrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mrgb.flds)+len(mrgb.fns))
		for _, f := range *mrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ModerationRequestSelect is the builder for selecting fields of ModerationRequest entities.
type ModerationRequestSelect struct {
	*ModerationRequestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mrs *ModerationRequestSelect) Aggregate(fns ...AggregateFunc) *ModerationRequestSelect {
	mrs.fns = append(mrs.fns, fns...)
	return mrs
}

// Scan applies the selector query and scans the result into the given value.
func (mrs *ModerationRequestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mrs.ctx, "Select")
	if err := mrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ModerationRequestQuery, *ModerationRequestSelect](ctx, mrs.ModerationRequestQuery, mrs, mrs.inters, v)
}

func (mrs *ModerationRequestSelect) sqlScan(ctx context.Context, root *ModerationRequestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mrs.fns))
	for _, fn := range mrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mrs *ModerationRequestSelect) Modify(modifiers ...func(s *sql.Selector)) *ModerationRequestSelect {
	mrs.modifiers = append(mrs.modifiers, modifiers...)
	return mrs
}
