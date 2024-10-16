// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/eventtype"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/userstats"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *EventCreate) SetName(s string) *EventCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetAddress sets the "address" field.
func (ec *EventCreate) SetAddress(s string) *EventCreate {
	ec.mutation.SetAddress(s)
	return ec
}

// SetEventCode sets the "event_code" field.
func (ec *EventCreate) SetEventCode(i int16) *EventCreate {
	ec.mutation.SetEventCode(i)
	return ec
}

// SetDate sets the "date" field.
func (ec *EventCreate) SetDate(s string) *EventCreate {
	ec.mutation.SetDate(s)
	return ec
}

// SetCreatedAt sets the "created_at" field.
func (ec *EventCreate) SetCreatedAt(t time.Time) *EventCreate {
	ec.mutation.SetCreatedAt(t)
	return ec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ec *EventCreate) SetNillableCreatedAt(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetCreatedAt(*t)
	}
	return ec
}

// SetIsPublic sets the "is_public" field.
func (ec *EventCreate) SetIsPublic(b bool) *EventCreate {
	ec.mutation.SetIsPublic(b)
	return ec
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (ec *EventCreate) SetNillableIsPublic(b *bool) *EventCreate {
	if b != nil {
		ec.SetIsPublic(*b)
	}
	return ec
}

// SetIsFinished sets the "is_finished" field.
func (ec *EventCreate) SetIsFinished(b bool) *EventCreate {
	ec.mutation.SetIsFinished(b)
	return ec
}

// SetNillableIsFinished sets the "is_finished" field if the given value is not nil.
func (ec *EventCreate) SetNillableIsFinished(b *bool) *EventCreate {
	if b != nil {
		ec.SetIsFinished(*b)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *EventCreate) SetID(u ulid.ID) *EventCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EventCreate) SetNillableID(u *ulid.ID) *EventCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// AddUserStatsIDIDs adds the "user_stats_id" edge to the UserStats entity by IDs.
func (ec *EventCreate) AddUserStatsIDIDs(ids ...ulid.ID) *EventCreate {
	ec.mutation.AddUserStatsIDIDs(ids...)
	return ec
}

// AddUserStatsID adds the "user_stats_id" edges to the UserStats entity.
func (ec *EventCreate) AddUserStatsID(u ...*UserStats) *EventCreate {
	ids := make([]ulid.ID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ec.AddUserStatsIDIDs(ids...)
}

// SetEventTypeID sets the "event_type" edge to the EventType entity by ID.
func (ec *EventCreate) SetEventTypeID(id string) *EventCreate {
	ec.mutation.SetEventTypeID(id)
	return ec
}

// SetNillableEventTypeID sets the "event_type" edge to the EventType entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableEventTypeID(id *string) *EventCreate {
	if id != nil {
		ec = ec.SetEventTypeID(*id)
	}
	return ec
}

// SetEventType sets the "event_type" edge to the EventType entity.
func (ec *EventCreate) SetEventType(e *EventType) *EventCreate {
	return ec.SetEventTypeID(e.ID)
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EventCreate) defaults() {
	if _, ok := ec.mutation.CreatedAt(); !ok {
		v := event.DefaultCreatedAt()
		ec.mutation.SetCreatedAt(v)
	}
	if _, ok := ec.mutation.IsPublic(); !ok {
		v := event.DefaultIsPublic
		ec.mutation.SetIsPublic(v)
	}
	if _, ok := ec.mutation.IsFinished(); !ok {
		v := event.DefaultIsFinished
		ec.mutation.SetIsFinished(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := event.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Event.name"`)}
	}
	if v, ok := ec.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Event.name": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Event.address"`)}
	}
	if v, ok := ec.mutation.Address(); ok {
		if err := event.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Event.address": %w`, err)}
		}
	}
	if _, ok := ec.mutation.EventCode(); !ok {
		return &ValidationError{Name: "event_code", err: errors.New(`ent: missing required field "Event.event_code"`)}
	}
	if v, ok := ec.mutation.EventCode(); ok {
		if err := event.EventCodeValidator(v); err != nil {
			return &ValidationError{Name: "event_code", err: fmt.Errorf(`ent: validator failed for field "Event.event_code": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Event.date"`)}
	}
	if v, ok := ec.mutation.Date(); ok {
		if err := event.DateValidator(v); err != nil {
			return &ValidationError{Name: "date", err: fmt.Errorf(`ent: validator failed for field "Event.date": %w`, err)}
		}
	}
	if _, ok := ec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Event.created_at"`)}
	}
	if _, ok := ec.mutation.IsPublic(); !ok {
		return &ValidationError{Name: "is_public", err: errors.New(`ent: missing required field "Event.is_public"`)}
	}
	if _, ok := ec.mutation.IsFinished(); !ok {
		return &ValidationError{Name: "is_finished", err: errors.New(`ent: missing required field "Event.is_finished"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*ulid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(event.Table, sqlgraph.NewFieldSpec(event.FieldID, field.TypeString))
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.SetField(event.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ec.mutation.Address(); ok {
		_spec.SetField(event.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := ec.mutation.EventCode(); ok {
		_spec.SetField(event.FieldEventCode, field.TypeInt16, value)
		_node.EventCode = value
	}
	if value, ok := ec.mutation.Date(); ok {
		_spec.SetField(event.FieldDate, field.TypeString, value)
		_node.Date = value
	}
	if value, ok := ec.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ec.mutation.IsPublic(); ok {
		_spec.SetField(event.FieldIsPublic, field.TypeBool, value)
		_node.IsPublic = value
	}
	if value, ok := ec.mutation.IsFinished(); ok {
		_spec.SetField(event.FieldIsFinished, field.TypeBool, value)
		_node.IsFinished = value
	}
	if nodes := ec.mutation.UserStatsIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.UserStatsIDTable,
			Columns: []string{event.UserStatsIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EventTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.EventTypeTable,
			Columns: []string{event.EventTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(eventtype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.event_type_event = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	err      error
	builders []*EventCreate
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
