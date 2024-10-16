// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/asma12a/challenge-s6/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/eventtype"
	"github.com/asma12a/challenge-s6/ent/sport"
	"github.com/asma12a/challenge-s6/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Event is the client for interacting with the Event builders.
	Event *EventClient
	// EventType is the client for interacting with the EventType builders.
	EventType *EventTypeClient
	// Sport is the client for interacting with the Sport builders.
	Sport *SportClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Event = NewEventClient(c.config)
	c.EventType = NewEventTypeClient(c.config)
	c.Sport = NewSportClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Event:     NewEventClient(cfg),
		EventType: NewEventTypeClient(cfg),
		Sport:     NewSportClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Event:     NewEventClient(cfg),
		EventType: NewEventTypeClient(cfg),
		Sport:     NewSportClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Event.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Event.Use(hooks...)
	c.EventType.Use(hooks...)
	c.Sport.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Event.Intercept(interceptors...)
	c.EventType.Intercept(interceptors...)
	c.Sport.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *EventMutation:
		return c.Event.mutate(ctx, m)
	case *EventTypeMutation:
		return c.EventType.mutate(ctx, m)
	case *SportMutation:
		return c.Sport.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// EventClient is a client for the Event schema.
type EventClient struct {
	config
}

// NewEventClient returns a client for the Event from the given config.
func NewEventClient(c config) *EventClient {
	return &EventClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `event.Hooks(f(g(h())))`.
func (c *EventClient) Use(hooks ...Hook) {
	c.hooks.Event = append(c.hooks.Event, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `event.Intercept(f(g(h())))`.
func (c *EventClient) Intercept(interceptors ...Interceptor) {
	c.inters.Event = append(c.inters.Event, interceptors...)
}

// Create returns a builder for creating a Event entity.
func (c *EventClient) Create() *EventCreate {
	mutation := newEventMutation(c.config, OpCreate)
	return &EventCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Event entities.
func (c *EventClient) CreateBulk(builders ...*EventCreate) *EventCreateBulk {
	return &EventCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EventClient) MapCreateBulk(slice any, setFunc func(*EventCreate, int)) *EventCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EventCreateBulk{err: fmt.Errorf("calling to EventClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EventCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EventCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Event.
func (c *EventClient) Update() *EventUpdate {
	mutation := newEventMutation(c.config, OpUpdate)
	return &EventUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EventClient) UpdateOne(e *Event) *EventUpdateOne {
	mutation := newEventMutation(c.config, OpUpdateOne, withEvent(e))
	return &EventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EventClient) UpdateOneID(id string) *EventUpdateOne {
	mutation := newEventMutation(c.config, OpUpdateOne, withEventID(id))
	return &EventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Event.
func (c *EventClient) Delete() *EventDelete {
	mutation := newEventMutation(c.config, OpDelete)
	return &EventDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EventClient) DeleteOne(e *Event) *EventDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EventClient) DeleteOneID(id string) *EventDeleteOne {
	builder := c.Delete().Where(event.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EventDeleteOne{builder}
}

// Query returns a query builder for Event.
func (c *EventClient) Query() *EventQuery {
	return &EventQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEvent},
		inters: c.Interceptors(),
	}
}

// Get returns a Event entity by its id.
func (c *EventClient) Get(ctx context.Context, id string) (*Event, error) {
	return c.Query().Where(event.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EventClient) GetX(ctx context.Context, id string) *Event {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEventType queries the event_type edge of a Event.
func (c *EventClient) QueryEventType(e *Event) *EventTypeQuery {
	query := (&EventTypeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, id),
			sqlgraph.To(eventtype.Table, eventtype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, event.EventTypeTable, event.EventTypeColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySport queries the sport edge of a Event.
func (c *EventClient) QuerySport(e *Event) *SportQuery {
	query := (&SportClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(event.Table, event.FieldID, id),
			sqlgraph.To(sport.Table, sport.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, event.SportTable, event.SportColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EventClient) Hooks() []Hook {
	return c.hooks.Event
}

// Interceptors returns the client interceptors.
func (c *EventClient) Interceptors() []Interceptor {
	return c.inters.Event
}

func (c *EventClient) mutate(ctx context.Context, m *EventMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EventCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EventUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EventUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EventDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Event mutation op: %q", m.Op())
	}
}

// EventTypeClient is a client for the EventType schema.
type EventTypeClient struct {
	config
}

// NewEventTypeClient returns a client for the EventType from the given config.
func NewEventTypeClient(c config) *EventTypeClient {
	return &EventTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `eventtype.Hooks(f(g(h())))`.
func (c *EventTypeClient) Use(hooks ...Hook) {
	c.hooks.EventType = append(c.hooks.EventType, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `eventtype.Intercept(f(g(h())))`.
func (c *EventTypeClient) Intercept(interceptors ...Interceptor) {
	c.inters.EventType = append(c.inters.EventType, interceptors...)
}

// Create returns a builder for creating a EventType entity.
func (c *EventTypeClient) Create() *EventTypeCreate {
	mutation := newEventTypeMutation(c.config, OpCreate)
	return &EventTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of EventType entities.
func (c *EventTypeClient) CreateBulk(builders ...*EventTypeCreate) *EventTypeCreateBulk {
	return &EventTypeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EventTypeClient) MapCreateBulk(slice any, setFunc func(*EventTypeCreate, int)) *EventTypeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EventTypeCreateBulk{err: fmt.Errorf("calling to EventTypeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EventTypeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EventTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for EventType.
func (c *EventTypeClient) Update() *EventTypeUpdate {
	mutation := newEventTypeMutation(c.config, OpUpdate)
	return &EventTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EventTypeClient) UpdateOne(et *EventType) *EventTypeUpdateOne {
	mutation := newEventTypeMutation(c.config, OpUpdateOne, withEventType(et))
	return &EventTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EventTypeClient) UpdateOneID(id string) *EventTypeUpdateOne {
	mutation := newEventTypeMutation(c.config, OpUpdateOne, withEventTypeID(id))
	return &EventTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for EventType.
func (c *EventTypeClient) Delete() *EventTypeDelete {
	mutation := newEventTypeMutation(c.config, OpDelete)
	return &EventTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EventTypeClient) DeleteOne(et *EventType) *EventTypeDeleteOne {
	return c.DeleteOneID(et.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EventTypeClient) DeleteOneID(id string) *EventTypeDeleteOne {
	builder := c.Delete().Where(eventtype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EventTypeDeleteOne{builder}
}

// Query returns a query builder for EventType.
func (c *EventTypeClient) Query() *EventTypeQuery {
	return &EventTypeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEventType},
		inters: c.Interceptors(),
	}
}

// Get returns a EventType entity by its id.
func (c *EventTypeClient) Get(ctx context.Context, id string) (*EventType, error) {
	return c.Query().Where(eventtype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EventTypeClient) GetX(ctx context.Context, id string) *EventType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEvents queries the events edge of a EventType.
func (c *EventTypeClient) QueryEvents(et *EventType) *EventQuery {
	query := (&EventClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := et.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(eventtype.Table, eventtype.FieldID, id),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, eventtype.EventsTable, eventtype.EventsColumn),
		)
		fromV = sqlgraph.Neighbors(et.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EventTypeClient) Hooks() []Hook {
	return c.hooks.EventType
}

// Interceptors returns the client interceptors.
func (c *EventTypeClient) Interceptors() []Interceptor {
	return c.inters.EventType
}

func (c *EventTypeClient) mutate(ctx context.Context, m *EventTypeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EventTypeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EventTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EventTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EventTypeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown EventType mutation op: %q", m.Op())
	}
}

// SportClient is a client for the Sport schema.
type SportClient struct {
	config
}

// NewSportClient returns a client for the Sport from the given config.
func NewSportClient(c config) *SportClient {
	return &SportClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sport.Hooks(f(g(h())))`.
func (c *SportClient) Use(hooks ...Hook) {
	c.hooks.Sport = append(c.hooks.Sport, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `sport.Intercept(f(g(h())))`.
func (c *SportClient) Intercept(interceptors ...Interceptor) {
	c.inters.Sport = append(c.inters.Sport, interceptors...)
}

// Create returns a builder for creating a Sport entity.
func (c *SportClient) Create() *SportCreate {
	mutation := newSportMutation(c.config, OpCreate)
	return &SportCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Sport entities.
func (c *SportClient) CreateBulk(builders ...*SportCreate) *SportCreateBulk {
	return &SportCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SportClient) MapCreateBulk(slice any, setFunc func(*SportCreate, int)) *SportCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SportCreateBulk{err: fmt.Errorf("calling to SportClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SportCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SportCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Sport.
func (c *SportClient) Update() *SportUpdate {
	mutation := newSportMutation(c.config, OpUpdate)
	return &SportUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SportClient) UpdateOne(s *Sport) *SportUpdateOne {
	mutation := newSportMutation(c.config, OpUpdateOne, withSport(s))
	return &SportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SportClient) UpdateOneID(id string) *SportUpdateOne {
	mutation := newSportMutation(c.config, OpUpdateOne, withSportID(id))
	return &SportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Sport.
func (c *SportClient) Delete() *SportDelete {
	mutation := newSportMutation(c.config, OpDelete)
	return &SportDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SportClient) DeleteOne(s *Sport) *SportDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SportClient) DeleteOneID(id string) *SportDeleteOne {
	builder := c.Delete().Where(sport.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SportDeleteOne{builder}
}

// Query returns a query builder for Sport.
func (c *SportClient) Query() *SportQuery {
	return &SportQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSport},
		inters: c.Interceptors(),
	}
}

// Get returns a Sport entity by its id.
func (c *SportClient) Get(ctx context.Context, id string) (*Sport, error) {
	return c.Query().Where(sport.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SportClient) GetX(ctx context.Context, id string) *Sport {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEvents queries the events edge of a Sport.
func (c *SportClient) QueryEvents(s *Sport) *EventQuery {
	query := (&EventClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(sport.Table, sport.FieldID, id),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sport.EventsTable, sport.EventsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SportClient) Hooks() []Hook {
	return c.hooks.Sport
}

// Interceptors returns the client interceptors.
func (c *SportClient) Interceptors() []Interceptor {
	return c.inters.Sport
}

func (c *SportClient) mutate(ctx context.Context, m *SportMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SportCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SportUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SportDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Sport mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Event, EventType, Sport, User []ent.Hook
	}
	inters struct {
		Event, EventType, Sport, User []ent.Interceptor
	}
)
