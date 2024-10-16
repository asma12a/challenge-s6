// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/footevent"
	"github.com/asma12a/challenge-s6/ent/predicate"
)

// FootEventUpdate is the builder for updating FootEvent entities.
type FootEventUpdate struct {
	config
	hooks    []Hook
	mutation *FootEventMutation
}

// Where appends a list predicates to the FootEventUpdate builder.
func (feu *FootEventUpdate) Where(ps ...predicate.FootEvent) *FootEventUpdate {
	feu.mutation.Where(ps...)
	return feu
}

// SetEventFootID sets the "event_foot_id" field.
func (feu *FootEventUpdate) SetEventFootID(s string) *FootEventUpdate {
	feu.mutation.SetEventFootID(s)
	return feu
}

// SetNillableEventFootID sets the "event_foot_id" field if the given value is not nil.
func (feu *FootEventUpdate) SetNillableEventFootID(s *string) *FootEventUpdate {
	if s != nil {
		feu.SetEventFootID(*s)
	}
	return feu
}

// SetTeamA sets the "team_A" field.
func (feu *FootEventUpdate) SetTeamA(s string) *FootEventUpdate {
	feu.mutation.SetTeamA(s)
	return feu
}

// SetNillableTeamA sets the "team_A" field if the given value is not nil.
func (feu *FootEventUpdate) SetNillableTeamA(s *string) *FootEventUpdate {
	if s != nil {
		feu.SetTeamA(*s)
	}
	return feu
}

// SetTeamB sets the "team_B" field.
func (feu *FootEventUpdate) SetTeamB(s string) *FootEventUpdate {
	feu.mutation.SetTeamB(s)
	return feu
}

// SetNillableTeamB sets the "team_B" field if the given value is not nil.
func (feu *FootEventUpdate) SetNillableTeamB(s *string) *FootEventUpdate {
	if s != nil {
		feu.SetTeamB(*s)
	}
	return feu
}

// Mutation returns the FootEventMutation object of the builder.
func (feu *FootEventUpdate) Mutation() *FootEventMutation {
	return feu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (feu *FootEventUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, feu.sqlSave, feu.mutation, feu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (feu *FootEventUpdate) SaveX(ctx context.Context) int {
	affected, err := feu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (feu *FootEventUpdate) Exec(ctx context.Context) error {
	_, err := feu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (feu *FootEventUpdate) ExecX(ctx context.Context) {
	if err := feu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (feu *FootEventUpdate) check() error {
	if v, ok := feu.mutation.EventFootID(); ok {
		if err := footevent.EventFootIDValidator(v); err != nil {
			return &ValidationError{Name: "event_foot_id", err: fmt.Errorf(`ent: validator failed for field "FootEvent.event_foot_id": %w`, err)}
		}
	}
	if v, ok := feu.mutation.TeamA(); ok {
		if err := footevent.TeamAValidator(v); err != nil {
			return &ValidationError{Name: "team_A", err: fmt.Errorf(`ent: validator failed for field "FootEvent.team_A": %w`, err)}
		}
	}
	if v, ok := feu.mutation.TeamB(); ok {
		if err := footevent.TeamBValidator(v); err != nil {
			return &ValidationError{Name: "team_B", err: fmt.Errorf(`ent: validator failed for field "FootEvent.team_B": %w`, err)}
		}
	}
	return nil
}

func (feu *FootEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := feu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(footevent.Table, footevent.Columns, sqlgraph.NewFieldSpec(footevent.FieldID, field.TypeString))
	if ps := feu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := feu.mutation.EventFootID(); ok {
		_spec.SetField(footevent.FieldEventFootID, field.TypeString, value)
	}
	if value, ok := feu.mutation.TeamA(); ok {
		_spec.SetField(footevent.FieldTeamA, field.TypeString, value)
	}
	if value, ok := feu.mutation.TeamB(); ok {
		_spec.SetField(footevent.FieldTeamB, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, feu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{footevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	feu.mutation.done = true
	return n, nil
}

// FootEventUpdateOne is the builder for updating a single FootEvent entity.
type FootEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FootEventMutation
}

// SetEventFootID sets the "event_foot_id" field.
func (feuo *FootEventUpdateOne) SetEventFootID(s string) *FootEventUpdateOne {
	feuo.mutation.SetEventFootID(s)
	return feuo
}

// SetNillableEventFootID sets the "event_foot_id" field if the given value is not nil.
func (feuo *FootEventUpdateOne) SetNillableEventFootID(s *string) *FootEventUpdateOne {
	if s != nil {
		feuo.SetEventFootID(*s)
	}
	return feuo
}

// SetTeamA sets the "team_A" field.
func (feuo *FootEventUpdateOne) SetTeamA(s string) *FootEventUpdateOne {
	feuo.mutation.SetTeamA(s)
	return feuo
}

// SetNillableTeamA sets the "team_A" field if the given value is not nil.
func (feuo *FootEventUpdateOne) SetNillableTeamA(s *string) *FootEventUpdateOne {
	if s != nil {
		feuo.SetTeamA(*s)
	}
	return feuo
}

// SetTeamB sets the "team_B" field.
func (feuo *FootEventUpdateOne) SetTeamB(s string) *FootEventUpdateOne {
	feuo.mutation.SetTeamB(s)
	return feuo
}

// SetNillableTeamB sets the "team_B" field if the given value is not nil.
func (feuo *FootEventUpdateOne) SetNillableTeamB(s *string) *FootEventUpdateOne {
	if s != nil {
		feuo.SetTeamB(*s)
	}
	return feuo
}

// Mutation returns the FootEventMutation object of the builder.
func (feuo *FootEventUpdateOne) Mutation() *FootEventMutation {
	return feuo.mutation
}

// Where appends a list predicates to the FootEventUpdate builder.
func (feuo *FootEventUpdateOne) Where(ps ...predicate.FootEvent) *FootEventUpdateOne {
	feuo.mutation.Where(ps...)
	return feuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (feuo *FootEventUpdateOne) Select(field string, fields ...string) *FootEventUpdateOne {
	feuo.fields = append([]string{field}, fields...)
	return feuo
}

// Save executes the query and returns the updated FootEvent entity.
func (feuo *FootEventUpdateOne) Save(ctx context.Context) (*FootEvent, error) {
	return withHooks(ctx, feuo.sqlSave, feuo.mutation, feuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (feuo *FootEventUpdateOne) SaveX(ctx context.Context) *FootEvent {
	node, err := feuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (feuo *FootEventUpdateOne) Exec(ctx context.Context) error {
	_, err := feuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (feuo *FootEventUpdateOne) ExecX(ctx context.Context) {
	if err := feuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (feuo *FootEventUpdateOne) check() error {
	if v, ok := feuo.mutation.EventFootID(); ok {
		if err := footevent.EventFootIDValidator(v); err != nil {
			return &ValidationError{Name: "event_foot_id", err: fmt.Errorf(`ent: validator failed for field "FootEvent.event_foot_id": %w`, err)}
		}
	}
	if v, ok := feuo.mutation.TeamA(); ok {
		if err := footevent.TeamAValidator(v); err != nil {
			return &ValidationError{Name: "team_A", err: fmt.Errorf(`ent: validator failed for field "FootEvent.team_A": %w`, err)}
		}
	}
	if v, ok := feuo.mutation.TeamB(); ok {
		if err := footevent.TeamBValidator(v); err != nil {
			return &ValidationError{Name: "team_B", err: fmt.Errorf(`ent: validator failed for field "FootEvent.team_B": %w`, err)}
		}
	}
	return nil
}

func (feuo *FootEventUpdateOne) sqlSave(ctx context.Context) (_node *FootEvent, err error) {
	if err := feuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(footevent.Table, footevent.Columns, sqlgraph.NewFieldSpec(footevent.FieldID, field.TypeString))
	id, ok := feuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FootEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := feuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, footevent.FieldID)
		for _, f := range fields {
			if !footevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != footevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := feuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := feuo.mutation.EventFootID(); ok {
		_spec.SetField(footevent.FieldEventFootID, field.TypeString, value)
	}
	if value, ok := feuo.mutation.TeamA(); ok {
		_spec.SetField(footevent.FieldTeamA, field.TypeString, value)
	}
	if value, ok := feuo.mutation.TeamB(); ok {
		_spec.SetField(footevent.FieldTeamB, field.TypeString, value)
	}
	_node = &FootEvent{config: feuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, feuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{footevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	feuo.mutation.done = true
	return _node, nil
}
