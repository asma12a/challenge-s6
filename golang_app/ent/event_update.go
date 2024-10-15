// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/predicate"
	"github.com/asma12a/challenge-s6/ent/userstats"
	ulid "github.com/oklog/ulid/v2"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EventUpdate) SetName(s string) *EventUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EventUpdate) SetNillableName(s *string) *EventUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetAddress sets the "address" field.
func (eu *EventUpdate) SetAddress(s string) *EventUpdate {
	eu.mutation.SetAddress(s)
	return eu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (eu *EventUpdate) SetNillableAddress(s *string) *EventUpdate {
	if s != nil {
		eu.SetAddress(*s)
	}
	return eu
}

// SetEventCode sets the "event_code" field.
func (eu *EventUpdate) SetEventCode(i int16) *EventUpdate {
	eu.mutation.ResetEventCode()
	eu.mutation.SetEventCode(i)
	return eu
}

// SetNillableEventCode sets the "event_code" field if the given value is not nil.
func (eu *EventUpdate) SetNillableEventCode(i *int16) *EventUpdate {
	if i != nil {
		eu.SetEventCode(*i)
	}
	return eu
}

// AddEventCode adds i to the "event_code" field.
func (eu *EventUpdate) AddEventCode(i int16) *EventUpdate {
	eu.mutation.AddEventCode(i)
	return eu
}

// SetDate sets the "date" field.
func (eu *EventUpdate) SetDate(s string) *EventUpdate {
	eu.mutation.SetDate(s)
	return eu
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (eu *EventUpdate) SetNillableDate(s *string) *EventUpdate {
	if s != nil {
		eu.SetDate(*s)
	}
	return eu
}

// SetCreatedAt sets the "created_at" field.
func (eu *EventUpdate) SetCreatedAt(t time.Time) *EventUpdate {
	eu.mutation.SetCreatedAt(t)
	return eu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eu *EventUpdate) SetNillableCreatedAt(t *time.Time) *EventUpdate {
	if t != nil {
		eu.SetCreatedAt(*t)
	}
	return eu
}

// SetIsPublic sets the "is_public" field.
func (eu *EventUpdate) SetIsPublic(b bool) *EventUpdate {
	eu.mutation.SetIsPublic(b)
	return eu
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (eu *EventUpdate) SetNillableIsPublic(b *bool) *EventUpdate {
	if b != nil {
		eu.SetIsPublic(*b)
	}
	return eu
}

// SetIsFinished sets the "is_finished" field.
func (eu *EventUpdate) SetIsFinished(b bool) *EventUpdate {
	eu.mutation.SetIsFinished(b)
	return eu
}

// SetNillableIsFinished sets the "is_finished" field if the given value is not nil.
func (eu *EventUpdate) SetNillableIsFinished(b *bool) *EventUpdate {
	if b != nil {
		eu.SetIsFinished(*b)
	}
	return eu
}

// AddUserStatsIDIDs adds the "user_stats_id" edge to the UserStats entity by IDs.
func (eu *EventUpdate) AddUserStatsIDIDs(ids ...ulid.ULID) *EventUpdate {
	eu.mutation.AddUserStatsIDIDs(ids...)
	return eu
}

// AddUserStatsID adds the "user_stats_id" edges to the UserStats entity.
func (eu *EventUpdate) AddUserStatsID(u ...*UserStats) *EventUpdate {
	ids := make([]ulid.ULID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return eu.AddUserStatsIDIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearUserStatsID clears all "user_stats_id" edges to the UserStats entity.
func (eu *EventUpdate) ClearUserStatsID() *EventUpdate {
	eu.mutation.ClearUserStatsID()
	return eu
}

// RemoveUserStatsIDIDs removes the "user_stats_id" edge to UserStats entities by IDs.
func (eu *EventUpdate) RemoveUserStatsIDIDs(ids ...ulid.ULID) *EventUpdate {
	eu.mutation.RemoveUserStatsIDIDs(ids...)
	return eu
}

// RemoveUserStatsID removes "user_stats_id" edges to UserStats entities.
func (eu *EventUpdate) RemoveUserStatsID(u ...*UserStats) *EventUpdate {
	ids := make([]ulid.ULID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return eu.RemoveUserStatsIDIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EventUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Event.name": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Address(); ok {
		if err := event.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Event.address": %w`, err)}
		}
	}
	if v, ok := eu.mutation.EventCode(); ok {
		if err := event.EventCodeValidator(v); err != nil {
			return &ValidationError{Name: "event_code", err: fmt.Errorf(`ent: validator failed for field "Event.event_code": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Date(); ok {
		if err := event.DateValidator(v); err != nil {
			return &ValidationError{Name: "date", err: fmt.Errorf(`ent: validator failed for field "Event.date": %w`, err)}
		}
	}
	return nil
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(event.Table, event.Columns, sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(event.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Address(); ok {
		_spec.SetField(event.FieldAddress, field.TypeString, value)
	}
	if value, ok := eu.mutation.EventCode(); ok {
		_spec.SetField(event.FieldEventCode, field.TypeInt16, value)
	}
	if value, ok := eu.mutation.AddedEventCode(); ok {
		_spec.AddField(event.FieldEventCode, field.TypeInt16, value)
	}
	if value, ok := eu.mutation.Date(); ok {
		_spec.SetField(event.FieldDate, field.TypeString, value)
	}
	if value, ok := eu.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.IsPublic(); ok {
		_spec.SetField(event.FieldIsPublic, field.TypeBool, value)
	}
	if value, ok := eu.mutation.IsFinished(); ok {
		_spec.SetField(event.FieldIsFinished, field.TypeBool, value)
	}
	if eu.mutation.UserStatsIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.UserStatsIDTable,
			Columns: []string{event.UserStatsIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedUserStatsIDIDs(); len(nodes) > 0 && !eu.mutation.UserStatsIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.UserStatsIDTable,
			Columns: []string{event.UserStatsIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.UserStatsIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.UserStatsIDTable,
			Columns: []string{event.UserStatsIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetName sets the "name" field.
func (euo *EventUpdateOne) SetName(s string) *EventUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableName(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetAddress sets the "address" field.
func (euo *EventUpdateOne) SetAddress(s string) *EventUpdateOne {
	euo.mutation.SetAddress(s)
	return euo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableAddress(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetAddress(*s)
	}
	return euo
}

// SetEventCode sets the "event_code" field.
func (euo *EventUpdateOne) SetEventCode(i int16) *EventUpdateOne {
	euo.mutation.ResetEventCode()
	euo.mutation.SetEventCode(i)
	return euo
}

// SetNillableEventCode sets the "event_code" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableEventCode(i *int16) *EventUpdateOne {
	if i != nil {
		euo.SetEventCode(*i)
	}
	return euo
}

// AddEventCode adds i to the "event_code" field.
func (euo *EventUpdateOne) AddEventCode(i int16) *EventUpdateOne {
	euo.mutation.AddEventCode(i)
	return euo
}

// SetDate sets the "date" field.
func (euo *EventUpdateOne) SetDate(s string) *EventUpdateOne {
	euo.mutation.SetDate(s)
	return euo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableDate(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetDate(*s)
	}
	return euo
}

// SetCreatedAt sets the "created_at" field.
func (euo *EventUpdateOne) SetCreatedAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetCreatedAt(t)
	return euo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCreatedAt(t *time.Time) *EventUpdateOne {
	if t != nil {
		euo.SetCreatedAt(*t)
	}
	return euo
}

// SetIsPublic sets the "is_public" field.
func (euo *EventUpdateOne) SetIsPublic(b bool) *EventUpdateOne {
	euo.mutation.SetIsPublic(b)
	return euo
}

// SetNillableIsPublic sets the "is_public" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableIsPublic(b *bool) *EventUpdateOne {
	if b != nil {
		euo.SetIsPublic(*b)
	}
	return euo
}

// SetIsFinished sets the "is_finished" field.
func (euo *EventUpdateOne) SetIsFinished(b bool) *EventUpdateOne {
	euo.mutation.SetIsFinished(b)
	return euo
}

// SetNillableIsFinished sets the "is_finished" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableIsFinished(b *bool) *EventUpdateOne {
	if b != nil {
		euo.SetIsFinished(*b)
	}
	return euo
}

// AddUserStatsIDIDs adds the "user_stats_id" edge to the UserStats entity by IDs.
func (euo *EventUpdateOne) AddUserStatsIDIDs(ids ...ulid.ULID) *EventUpdateOne {
	euo.mutation.AddUserStatsIDIDs(ids...)
	return euo
}

// AddUserStatsID adds the "user_stats_id" edges to the UserStats entity.
func (euo *EventUpdateOne) AddUserStatsID(u ...*UserStats) *EventUpdateOne {
	ids := make([]ulid.ULID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return euo.AddUserStatsIDIDs(ids...)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearUserStatsID clears all "user_stats_id" edges to the UserStats entity.
func (euo *EventUpdateOne) ClearUserStatsID() *EventUpdateOne {
	euo.mutation.ClearUserStatsID()
	return euo
}

// RemoveUserStatsIDIDs removes the "user_stats_id" edge to UserStats entities by IDs.
func (euo *EventUpdateOne) RemoveUserStatsIDIDs(ids ...ulid.ULID) *EventUpdateOne {
	euo.mutation.RemoveUserStatsIDIDs(ids...)
	return euo
}

// RemoveUserStatsID removes "user_stats_id" edges to UserStats entities.
func (euo *EventUpdateOne) RemoveUserStatsID(u ...*UserStats) *EventUpdateOne {
	ids := make([]ulid.ULID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return euo.RemoveUserStatsIDIDs(ids...)
}

// Where appends a list predicates to the EventUpdate builder.
func (euo *EventUpdateOne) Where(ps ...predicate.Event) *EventUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EventUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Event.name": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Address(); ok {
		if err := event.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Event.address": %w`, err)}
		}
	}
	if v, ok := euo.mutation.EventCode(); ok {
		if err := event.EventCodeValidator(v); err != nil {
			return &ValidationError{Name: "event_code", err: fmt.Errorf(`ent: validator failed for field "Event.event_code": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Date(); ok {
		if err := event.DateValidator(v); err != nil {
			return &ValidationError{Name: "date", err: fmt.Errorf(`ent: validator failed for field "Event.date": %w`, err)}
		}
	}
	return nil
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(event.Table, event.Columns, sqlgraph.NewFieldSpec(event.FieldID, field.TypeUUID))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(event.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Address(); ok {
		_spec.SetField(event.FieldAddress, field.TypeString, value)
	}
	if value, ok := euo.mutation.EventCode(); ok {
		_spec.SetField(event.FieldEventCode, field.TypeInt16, value)
	}
	if value, ok := euo.mutation.AddedEventCode(); ok {
		_spec.AddField(event.FieldEventCode, field.TypeInt16, value)
	}
	if value, ok := euo.mutation.Date(); ok {
		_spec.SetField(event.FieldDate, field.TypeString, value)
	}
	if value, ok := euo.mutation.CreatedAt(); ok {
		_spec.SetField(event.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.IsPublic(); ok {
		_spec.SetField(event.FieldIsPublic, field.TypeBool, value)
	}
	if value, ok := euo.mutation.IsFinished(); ok {
		_spec.SetField(event.FieldIsFinished, field.TypeBool, value)
	}
	if euo.mutation.UserStatsIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.UserStatsIDTable,
			Columns: []string{event.UserStatsIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedUserStatsIDIDs(); len(nodes) > 0 && !euo.mutation.UserStatsIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.UserStatsIDTable,
			Columns: []string{event.UserStatsIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.UserStatsIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.UserStatsIDTable,
			Columns: []string{event.UserStatsIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
