// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/eventtype"
	"github.com/asma12a/challenge-s6/ent/predicate"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// EventTypeQuery is the builder for querying EventType entities.
type EventTypeQuery struct {
	config
	ctx        *QueryContext
	order      []eventtype.OrderOption
	inters     []Interceptor
	predicates []predicate.EventType
	withEvents *EventQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EventTypeQuery builder.
func (etq *EventTypeQuery) Where(ps ...predicate.EventType) *EventTypeQuery {
	etq.predicates = append(etq.predicates, ps...)
	return etq
}

// Limit the number of records to be returned by this query.
func (etq *EventTypeQuery) Limit(limit int) *EventTypeQuery {
	etq.ctx.Limit = &limit
	return etq
}

// Offset to start from.
func (etq *EventTypeQuery) Offset(offset int) *EventTypeQuery {
	etq.ctx.Offset = &offset
	return etq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (etq *EventTypeQuery) Unique(unique bool) *EventTypeQuery {
	etq.ctx.Unique = &unique
	return etq
}

// Order specifies how the records should be ordered.
func (etq *EventTypeQuery) Order(o ...eventtype.OrderOption) *EventTypeQuery {
	etq.order = append(etq.order, o...)
	return etq
}

// QueryEvents chains the current query on the "events" edge.
func (etq *EventTypeQuery) QueryEvents() *EventQuery {
	query := (&EventClient{config: etq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eventtype.Table, eventtype.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, eventtype.EventsTable, eventtype.EventsColumn),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EventType entity from the query.
// Returns a *NotFoundError when no EventType was found.
func (etq *EventTypeQuery) First(ctx context.Context) (*EventType, error) {
	nodes, err := etq.Limit(1).All(setContextOp(ctx, etq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{eventtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (etq *EventTypeQuery) FirstX(ctx context.Context) *EventType {
	node, err := etq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EventType ID from the query.
// Returns a *NotFoundError when no EventType ID was found.
func (etq *EventTypeQuery) FirstID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = etq.Limit(1).IDs(setContextOp(ctx, etq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{eventtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (etq *EventTypeQuery) FirstIDX(ctx context.Context) ulid.ID {
	id, err := etq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EventType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EventType entity is found.
// Returns a *NotFoundError when no EventType entities are found.
func (etq *EventTypeQuery) Only(ctx context.Context) (*EventType, error) {
	nodes, err := etq.Limit(2).All(setContextOp(ctx, etq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{eventtype.Label}
	default:
		return nil, &NotSingularError{eventtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (etq *EventTypeQuery) OnlyX(ctx context.Context) *EventType {
	node, err := etq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EventType ID in the query.
// Returns a *NotSingularError when more than one EventType ID is found.
// Returns a *NotFoundError when no entities are found.
func (etq *EventTypeQuery) OnlyID(ctx context.Context) (id ulid.ID, err error) {
	var ids []ulid.ID
	if ids, err = etq.Limit(2).IDs(setContextOp(ctx, etq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = &NotSingularError{eventtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (etq *EventTypeQuery) OnlyIDX(ctx context.Context) ulid.ID {
	id, err := etq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EventTypes.
func (etq *EventTypeQuery) All(ctx context.Context) ([]*EventType, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryAll)
	if err := etq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EventType, *EventTypeQuery]()
	return withInterceptors[[]*EventType](ctx, etq, qr, etq.inters)
}

// AllX is like All, but panics if an error occurs.
func (etq *EventTypeQuery) AllX(ctx context.Context) []*EventType {
	nodes, err := etq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EventType IDs.
func (etq *EventTypeQuery) IDs(ctx context.Context) (ids []ulid.ID, err error) {
	if etq.ctx.Unique == nil && etq.path != nil {
		etq.Unique(true)
	}
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryIDs)
	if err = etq.Select(eventtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (etq *EventTypeQuery) IDsX(ctx context.Context) []ulid.ID {
	ids, err := etq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (etq *EventTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryCount)
	if err := etq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, etq, querierCount[*EventTypeQuery](), etq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (etq *EventTypeQuery) CountX(ctx context.Context) int {
	count, err := etq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (etq *EventTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryExist)
	switch _, err := etq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (etq *EventTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := etq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EventTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (etq *EventTypeQuery) Clone() *EventTypeQuery {
	if etq == nil {
		return nil
	}
	return &EventTypeQuery{
		config:     etq.config,
		ctx:        etq.ctx.Clone(),
		order:      append([]eventtype.OrderOption{}, etq.order...),
		inters:     append([]Interceptor{}, etq.inters...),
		predicates: append([]predicate.EventType{}, etq.predicates...),
		withEvents: etq.withEvents.Clone(),
		// clone intermediate query.
		sql:  etq.sql.Clone(),
		path: etq.path,
	}
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EventTypeQuery) WithEvents(opts ...func(*EventQuery)) *EventTypeQuery {
	query := (&EventClient{config: etq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	etq.withEvents = query
	return etq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EventType.Query().
//		GroupBy(eventtype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (etq *EventTypeQuery) GroupBy(field string, fields ...string) *EventTypeGroupBy {
	etq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EventTypeGroupBy{build: etq}
	grbuild.flds = &etq.ctx.Fields
	grbuild.label = eventtype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.EventType.Query().
//		Select(eventtype.FieldName).
//		Scan(ctx, &v)
func (etq *EventTypeQuery) Select(fields ...string) *EventTypeSelect {
	etq.ctx.Fields = append(etq.ctx.Fields, fields...)
	sbuild := &EventTypeSelect{EventTypeQuery: etq}
	sbuild.label = eventtype.Label
	sbuild.flds, sbuild.scan = &etq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EventTypeSelect configured with the given aggregations.
func (etq *EventTypeQuery) Aggregate(fns ...AggregateFunc) *EventTypeSelect {
	return etq.Select().Aggregate(fns...)
}

func (etq *EventTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range etq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, etq); err != nil {
				return err
			}
		}
	}
	for _, f := range etq.ctx.Fields {
		if !eventtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if etq.path != nil {
		prev, err := etq.path(ctx)
		if err != nil {
			return err
		}
		etq.sql = prev
	}
	return nil
}

func (etq *EventTypeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EventType, error) {
	var (
		nodes       = []*EventType{}
		_spec       = etq.querySpec()
		loadedTypes = [1]bool{
			etq.withEvents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EventType).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EventType{config: etq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, etq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := etq.withEvents; query != nil {
		if err := etq.loadEvents(ctx, query, nodes,
			func(n *EventType) { n.Edges.Events = []*Event{} },
			func(n *EventType, e *Event) { n.Edges.Events = append(n.Edges.Events, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (etq *EventTypeQuery) loadEvents(ctx context.Context, query *EventQuery, nodes []*EventType, init func(*EventType), assign func(*EventType, *Event)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[ulid.ID]*EventType)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Event(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(eventtype.EventsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.event_type_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "event_type_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "event_type_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (etq *EventTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := etq.querySpec()
	_spec.Node.Columns = etq.ctx.Fields
	if len(etq.ctx.Fields) > 0 {
		_spec.Unique = etq.ctx.Unique != nil && *etq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, etq.driver, _spec)
}

func (etq *EventTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(eventtype.Table, eventtype.Columns, sqlgraph.NewFieldSpec(eventtype.FieldID, field.TypeString))
	_spec.From = etq.sql
	if unique := etq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if etq.path != nil {
		_spec.Unique = true
	}
	if fields := etq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventtype.FieldID)
		for i := range fields {
			if fields[i] != eventtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := etq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := etq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := etq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := etq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (etq *EventTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(etq.driver.Dialect())
	t1 := builder.Table(eventtype.Table)
	columns := etq.ctx.Fields
	if len(columns) == 0 {
		columns = eventtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if etq.sql != nil {
		selector = etq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if etq.ctx.Unique != nil && *etq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range etq.predicates {
		p(selector)
	}
	for _, p := range etq.order {
		p(selector)
	}
	if offset := etq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := etq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EventTypeGroupBy is the group-by builder for EventType entities.
type EventTypeGroupBy struct {
	selector
	build *EventTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (etgb *EventTypeGroupBy) Aggregate(fns ...AggregateFunc) *EventTypeGroupBy {
	etgb.fns = append(etgb.fns, fns...)
	return etgb
}

// Scan applies the selector query and scans the result into the given value.
func (etgb *EventTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, etgb.build.ctx, ent.OpQueryGroupBy)
	if err := etgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EventTypeQuery, *EventTypeGroupBy](ctx, etgb.build, etgb, etgb.build.inters, v)
}

func (etgb *EventTypeGroupBy) sqlScan(ctx context.Context, root *EventTypeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(etgb.fns))
	for _, fn := range etgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*etgb.flds)+len(etgb.fns))
		for _, f := range *etgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*etgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := etgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EventTypeSelect is the builder for selecting fields of EventType entities.
type EventTypeSelect struct {
	*EventTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ets *EventTypeSelect) Aggregate(fns ...AggregateFunc) *EventTypeSelect {
	ets.fns = append(ets.fns, fns...)
	return ets
}

// Scan applies the selector query and scans the result into the given value.
func (ets *EventTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ets.ctx, ent.OpQuerySelect)
	if err := ets.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EventTypeQuery, *EventTypeSelect](ctx, ets.EventTypeQuery, ets, ets.inters, v)
}

func (ets *EventTypeSelect) sqlScan(ctx context.Context, root *EventTypeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ets.fns))
	for _, fn := range ets.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ets.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
