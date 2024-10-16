// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/predicate"
	"github.com/asma12a/challenge-s6/ent/trainingevent"
)

// TrainingEventUpdate is the builder for updating TrainingEvent entities.
type TrainingEventUpdate struct {
	config
	hooks    []Hook
	mutation *TrainingEventMutation
}

// Where appends a list predicates to the TrainingEventUpdate builder.
func (teu *TrainingEventUpdate) Where(ps ...predicate.TrainingEvent) *TrainingEventUpdate {
	teu.mutation.Where(ps...)
	return teu
}

// SetEventTrainingID sets the "event_training_id" field.
func (teu *TrainingEventUpdate) SetEventTrainingID(s string) *TrainingEventUpdate {
	teu.mutation.SetEventTrainingID(s)
	return teu
}

// SetNillableEventTrainingID sets the "event_training_id" field if the given value is not nil.
func (teu *TrainingEventUpdate) SetNillableEventTrainingID(s *string) *TrainingEventUpdate {
	if s != nil {
		teu.SetEventTrainingID(*s)
	}
	return teu
}

// SetTeamID sets the "team_id" field.
func (teu *TrainingEventUpdate) SetTeamID(s string) *TrainingEventUpdate {
	teu.mutation.SetTeamID(s)
	return teu
}

// SetNillableTeamID sets the "team_id" field if the given value is not nil.
func (teu *TrainingEventUpdate) SetNillableTeamID(s *string) *TrainingEventUpdate {
	if s != nil {
		teu.SetTeamID(*s)
	}
	return teu
}

// Mutation returns the TrainingEventMutation object of the builder.
func (teu *TrainingEventUpdate) Mutation() *TrainingEventMutation {
	return teu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (teu *TrainingEventUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, teu.sqlSave, teu.mutation, teu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (teu *TrainingEventUpdate) SaveX(ctx context.Context) int {
	affected, err := teu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (teu *TrainingEventUpdate) Exec(ctx context.Context) error {
	_, err := teu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (teu *TrainingEventUpdate) ExecX(ctx context.Context) {
	if err := teu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (teu *TrainingEventUpdate) check() error {
	if v, ok := teu.mutation.EventTrainingID(); ok {
		if err := trainingevent.EventTrainingIDValidator(v); err != nil {
			return &ValidationError{Name: "event_training_id", err: fmt.Errorf(`ent: validator failed for field "TrainingEvent.event_training_id": %w`, err)}
		}
	}
	if v, ok := teu.mutation.TeamID(); ok {
		if err := trainingevent.TeamIDValidator(v); err != nil {
			return &ValidationError{Name: "team_id", err: fmt.Errorf(`ent: validator failed for field "TrainingEvent.team_id": %w`, err)}
		}
	}
	return nil
}

func (teu *TrainingEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := teu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(trainingevent.Table, trainingevent.Columns, sqlgraph.NewFieldSpec(trainingevent.FieldID, field.TypeString))
	if ps := teu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := teu.mutation.EventTrainingID(); ok {
		_spec.SetField(trainingevent.FieldEventTrainingID, field.TypeString, value)
	}
	if value, ok := teu.mutation.TeamID(); ok {
		_spec.SetField(trainingevent.FieldTeamID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, teu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trainingevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	teu.mutation.done = true
	return n, nil
}

// TrainingEventUpdateOne is the builder for updating a single TrainingEvent entity.
type TrainingEventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TrainingEventMutation
}

// SetEventTrainingID sets the "event_training_id" field.
func (teuo *TrainingEventUpdateOne) SetEventTrainingID(s string) *TrainingEventUpdateOne {
	teuo.mutation.SetEventTrainingID(s)
	return teuo
}

// SetNillableEventTrainingID sets the "event_training_id" field if the given value is not nil.
func (teuo *TrainingEventUpdateOne) SetNillableEventTrainingID(s *string) *TrainingEventUpdateOne {
	if s != nil {
		teuo.SetEventTrainingID(*s)
	}
	return teuo
}

// SetTeamID sets the "team_id" field.
func (teuo *TrainingEventUpdateOne) SetTeamID(s string) *TrainingEventUpdateOne {
	teuo.mutation.SetTeamID(s)
	return teuo
}

// SetNillableTeamID sets the "team_id" field if the given value is not nil.
func (teuo *TrainingEventUpdateOne) SetNillableTeamID(s *string) *TrainingEventUpdateOne {
	if s != nil {
		teuo.SetTeamID(*s)
	}
	return teuo
}

// Mutation returns the TrainingEventMutation object of the builder.
func (teuo *TrainingEventUpdateOne) Mutation() *TrainingEventMutation {
	return teuo.mutation
}

// Where appends a list predicates to the TrainingEventUpdate builder.
func (teuo *TrainingEventUpdateOne) Where(ps ...predicate.TrainingEvent) *TrainingEventUpdateOne {
	teuo.mutation.Where(ps...)
	return teuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (teuo *TrainingEventUpdateOne) Select(field string, fields ...string) *TrainingEventUpdateOne {
	teuo.fields = append([]string{field}, fields...)
	return teuo
}

// Save executes the query and returns the updated TrainingEvent entity.
func (teuo *TrainingEventUpdateOne) Save(ctx context.Context) (*TrainingEvent, error) {
	return withHooks(ctx, teuo.sqlSave, teuo.mutation, teuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (teuo *TrainingEventUpdateOne) SaveX(ctx context.Context) *TrainingEvent {
	node, err := teuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (teuo *TrainingEventUpdateOne) Exec(ctx context.Context) error {
	_, err := teuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (teuo *TrainingEventUpdateOne) ExecX(ctx context.Context) {
	if err := teuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (teuo *TrainingEventUpdateOne) check() error {
	if v, ok := teuo.mutation.EventTrainingID(); ok {
		if err := trainingevent.EventTrainingIDValidator(v); err != nil {
			return &ValidationError{Name: "event_training_id", err: fmt.Errorf(`ent: validator failed for field "TrainingEvent.event_training_id": %w`, err)}
		}
	}
	if v, ok := teuo.mutation.TeamID(); ok {
		if err := trainingevent.TeamIDValidator(v); err != nil {
			return &ValidationError{Name: "team_id", err: fmt.Errorf(`ent: validator failed for field "TrainingEvent.team_id": %w`, err)}
		}
	}
	return nil
}

func (teuo *TrainingEventUpdateOne) sqlSave(ctx context.Context) (_node *TrainingEvent, err error) {
	if err := teuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(trainingevent.Table, trainingevent.Columns, sqlgraph.NewFieldSpec(trainingevent.FieldID, field.TypeString))
	id, ok := teuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TrainingEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := teuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trainingevent.FieldID)
		for _, f := range fields {
			if !trainingevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != trainingevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := teuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := teuo.mutation.EventTrainingID(); ok {
		_spec.SetField(trainingevent.FieldEventTrainingID, field.TypeString, value)
	}
	if value, ok := teuo.mutation.TeamID(); ok {
		_spec.SetField(trainingevent.FieldTeamID, field.TypeString, value)
	}
	_node = &TrainingEvent{config: teuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, teuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trainingevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	teuo.mutation.done = true
	return _node, nil
}
