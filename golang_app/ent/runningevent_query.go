// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/predicate"
	"github.com/asma12a/challenge-s6/ent/runningevent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// RunningEventQuery is the builder for querying RunningEvent entities.
type RunningEventQuery struct {
	config
	ctx        *QueryContext
	order      []runningevent.OrderOption
	inters     []Interceptor
	predicates []predicate.RunningEvent
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RunningEventQuery builder.
func (req *RunningEventQuery) Where(ps ...predicate.RunningEvent) *RunningEventQuery {
	req.predicates = append(req.predicates, ps...)
	return req
}

// Limit the number of records to be returned by this query.
func (req *RunningEventQuery) Limit(limit int) *RunningEventQuery {
	req.ctx.Limit = &limit
	return req
}

// Offset to start from.
func (req *RunningEventQuery) Offset(offset int) *RunningEventQuery {
	req.ctx.Offset = &offset
	return req
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (req *RunningEventQuery) Unique(unique bool) *RunningEventQuery {
	req.ctx.Unique = &unique
	return req
}

// Order specifies how the records should be ordered.
func (req *RunningEventQuery) Order(o ...runningevent.OrderOption) *RunningEventQuery {
	req.order = append(req.order, o...)
	return req
}

// First returns the first RunningEvent entity from the query.
// Returns a *NotFoundError when no RunningEvent was found.
func (req *RunningEventQuery) First(ctx context.Context) (*RunningEvent, error) {
	nodes, err := req.Limit(1).All(setContextOp(ctx, req.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{runningevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (req *RunningEventQuery) FirstX(ctx context.Context) *RunningEvent {
	node, err := req.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RunningEvent ID from the query.
// Returns a *NotFoundError when no RunningEvent ID was found.
func (req *RunningEventQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = req.Limit(1).IDs(setContextOp(ctx, req.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{runningevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (req *RunningEventQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := req.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RunningEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RunningEvent entity is found.
// Returns a *NotFoundError when no RunningEvent entities are found.
func (req *RunningEventQuery) Only(ctx context.Context) (*RunningEvent, error) {
	nodes, err := req.Limit(2).All(setContextOp(ctx, req.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{runningevent.Label}
	default:
		return nil, &NotSingularError{runningevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (req *RunningEventQuery) OnlyX(ctx context.Context) *RunningEvent {
	node, err := req.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RunningEvent ID in the query.
// Returns a *NotSingularError when more than one RunningEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (req *RunningEventQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = req.Limit(2).IDs(setContextOp(ctx, req.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{runningevent.Label}
	default:
		err = &NotSingularError{runningevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (req *RunningEventQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := req.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RunningEvents.
func (req *RunningEventQuery) All(ctx context.Context) ([]*RunningEvent, error) {
	ctx = setContextOp(ctx, req.ctx, ent.OpQueryAll)
	if err := req.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RunningEvent, *RunningEventQuery]()
	return withInterceptors[[]*RunningEvent](ctx, req, qr, req.inters)
}

// AllX is like All, but panics if an error occurs.
func (req *RunningEventQuery) AllX(ctx context.Context) []*RunningEvent {
	nodes, err := req.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RunningEvent IDs.
func (req *RunningEventQuery) IDs(ctx context.Context) (ids []ulid.ID, err error) {
	if req.ctx.Unique == nil && req.path != nil {
		req.Unique(true)
	}
	ctx = setContextOp(ctx, req.ctx, ent.OpQueryIDs)
	if err = req.Select(runningevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (req *RunningEventQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := req.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (req *RunningEventQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, req.ctx, ent.OpQueryCount)
	if err := req.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, req, querierCount[*RunningEventQuery](), req.inters)
}

// CountX is like Count, but panics if an error occurs.
func (req *RunningEventQuery) CountX(ctx context.Context) int {
	count, err := req.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (req *RunningEventQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, req.ctx, ent.OpQueryExist)
	switch _, err := req.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (req *RunningEventQuery) ExistX(ctx context.Context) bool {
	exist, err := req.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RunningEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (req *RunningEventQuery) Clone() *RunningEventQuery {
	if req == nil {
		return nil
	}
	return &RunningEventQuery{
		config:     req.config,
		ctx:        req.ctx.Clone(),
		order:      append([]runningevent.OrderOption{}, req.order...),
		inters:     append([]Interceptor{}, req.inters...),
		predicates: append([]predicate.RunningEvent{}, req.predicates...),
		// clone intermediate query.
		sql:  req.sql.Clone(),
		path: req.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		EventRunningID string `json:"event_running_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RunningEvent.Query().
//		GroupBy(runningevent.FieldEventRunningID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (req *RunningEventQuery) GroupBy(field string, fields ...string) *RunningEventGroupBy {
	req.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RunningEventGroupBy{build: req}
	grbuild.flds = &req.ctx.Fields
	grbuild.label = runningevent.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		EventRunningID string `json:"event_running_id,omitempty"`
//	}
//
//	client.RunningEvent.Query().
//		Select(runningevent.FieldEventRunningID).
//		Scan(ctx, &v)
func (req *RunningEventQuery) Select(fields ...string) *RunningEventSelect {
	req.ctx.Fields = append(req.ctx.Fields, fields...)
	sbuild := &RunningEventSelect{RunningEventQuery: req}
	sbuild.label = runningevent.Label
	sbuild.flds, sbuild.scan = &req.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RunningEventSelect configured with the given aggregations.
func (req *RunningEventQuery) Aggregate(fns ...AggregateFunc) *RunningEventSelect {
	return req.Select().Aggregate(fns...)
}

func (req *RunningEventQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range req.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, req); err != nil {
				return err
			}
		}
	}
	for _, f := range req.ctx.Fields {
		if !runningevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if req.path != nil {
		prev, err := req.path(ctx)
		if err != nil {
			return err
		}
		req.sql = prev
	}
	return nil
}

func (req *RunningEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RunningEvent, error) {
	var (
		nodes = []*RunningEvent{}
		_spec = req.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RunningEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RunningEvent{config: req.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, req.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (req *RunningEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := req.querySpec()
	_spec.Node.Columns = req.ctx.Fields
	if len(req.ctx.Fields) > 0 {
		_spec.Unique = req.ctx.Unique != nil && *req.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, req.driver, _spec)
}

func (req *RunningEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(runningevent.Table, runningevent.Columns, sqlgraph.NewFieldSpec(runningevent.FieldID, field.TypeString))
	_spec.From = req.sql
	if unique := req.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if req.path != nil {
		_spec.Unique = true
	}
	if fields := req.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, runningevent.FieldID)
		for i := range fields {
			if fields[i] != runningevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := req.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := req.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := req.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := req.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (req *RunningEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(req.driver.Dialect())
	t1 := builder.Table(runningevent.Table)
	columns := req.ctx.Fields
	if len(columns) == 0 {
		columns = runningevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if req.sql != nil {
		selector = req.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if req.ctx.Unique != nil && *req.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range req.predicates {
		p(selector)
	}
	for _, p := range req.order {
		p(selector)
	}
	if offset := req.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := req.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RunningEventGroupBy is the group-by builder for RunningEvent entities.
type RunningEventGroupBy struct {
	selector
	build *RunningEventQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (regb *RunningEventGroupBy) Aggregate(fns ...AggregateFunc) *RunningEventGroupBy {
	regb.fns = append(regb.fns, fns...)
	return regb
}

// Scan applies the selector query and scans the result into the given value.
func (regb *RunningEventGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, regb.build.ctx, ent.OpQueryGroupBy)
	if err := regb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RunningEventQuery, *RunningEventGroupBy](ctx, regb.build, regb, regb.build.inters, v)
}

func (regb *RunningEventGroupBy) sqlScan(ctx context.Context, root *RunningEventQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(regb.fns))
	for _, fn := range regb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*regb.flds)+len(regb.fns))
		for _, f := range *regb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*regb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := regb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RunningEventSelect is the builder for selecting fields of RunningEvent entities.
type RunningEventSelect struct {
	*RunningEventQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (res *RunningEventSelect) Aggregate(fns ...AggregateFunc) *RunningEventSelect {
	res.fns = append(res.fns, fns...)
	return res
}

// Scan applies the selector query and scans the result into the given value.
func (res *RunningEventSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, res.ctx, ent.OpQuerySelect)
	if err := res.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RunningEventQuery, *RunningEventSelect](ctx, res.RunningEventQuery, res, res.inters, v)
}

func (res *RunningEventSelect) sqlScan(ctx context.Context, root *RunningEventQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(res.fns))
	for _, fn := range res.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*res.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := res.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
