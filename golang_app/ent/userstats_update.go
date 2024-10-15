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
	"github.com/asma12a/challenge-s6/ent/userstats"
	"github.com/google/uuid"
)

// UserStatsUpdate is the builder for updating UserStats entities.
type UserStatsUpdate struct {
	config
	hooks    []Hook
	mutation *UserStatsMutation
}

// Where appends a list predicates to the UserStatsUpdate builder.
func (usu *UserStatsUpdate) Where(ps ...predicate.UserStats) *UserStatsUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetUserID sets the "user_id" field.
func (usu *UserStatsUpdate) SetUserID(s string) *UserStatsUpdate {
	usu.mutation.SetUserID(s)
	return usu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (usu *UserStatsUpdate) SetNillableUserID(s *string) *UserStatsUpdate {
	if s != nil {
		usu.SetUserID(*s)
	}
	return usu
}

// SetEventID sets the "event_id" field.
func (usu *UserStatsUpdate) SetEventID(u uuid.UUID) *UserStatsUpdate {
	usu.mutation.SetEventID(u)
	return usu
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (usu *UserStatsUpdate) SetNillableEventID(u *uuid.UUID) *UserStatsUpdate {
	if u != nil {
		usu.SetEventID(*u)
	}
	return usu
}

// Mutation returns the UserStatsMutation object of the builder.
func (usu *UserStatsUpdate) Mutation() *UserStatsMutation {
	return usu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UserStatsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, usu.sqlSave, usu.mutation, usu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserStatsUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserStatsUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserStatsUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (usu *UserStatsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userstats.Table, userstats.Columns, sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeUUID))
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.UserID(); ok {
		_spec.SetField(userstats.FieldUserID, field.TypeString, value)
	}
	if value, ok := usu.mutation.EventID(); ok {
		_spec.SetField(userstats.FieldEventID, field.TypeUUID, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	usu.mutation.done = true
	return n, nil
}

// UserStatsUpdateOne is the builder for updating a single UserStats entity.
type UserStatsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserStatsMutation
}

// SetUserID sets the "user_id" field.
func (usuo *UserStatsUpdateOne) SetUserID(s string) *UserStatsUpdateOne {
	usuo.mutation.SetUserID(s)
	return usuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (usuo *UserStatsUpdateOne) SetNillableUserID(s *string) *UserStatsUpdateOne {
	if s != nil {
		usuo.SetUserID(*s)
	}
	return usuo
}

// SetEventID sets the "event_id" field.
func (usuo *UserStatsUpdateOne) SetEventID(u uuid.UUID) *UserStatsUpdateOne {
	usuo.mutation.SetEventID(u)
	return usuo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (usuo *UserStatsUpdateOne) SetNillableEventID(u *uuid.UUID) *UserStatsUpdateOne {
	if u != nil {
		usuo.SetEventID(*u)
	}
	return usuo
}

// Mutation returns the UserStatsMutation object of the builder.
func (usuo *UserStatsUpdateOne) Mutation() *UserStatsMutation {
	return usuo.mutation
}

// Where appends a list predicates to the UserStatsUpdate builder.
func (usuo *UserStatsUpdateOne) Where(ps ...predicate.UserStats) *UserStatsUpdateOne {
	usuo.mutation.Where(ps...)
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UserStatsUpdateOne) Select(field string, fields ...string) *UserStatsUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UserStats entity.
func (usuo *UserStatsUpdateOne) Save(ctx context.Context) (*UserStats, error) {
	return withHooks(ctx, usuo.sqlSave, usuo.mutation, usuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserStatsUpdateOne) SaveX(ctx context.Context) *UserStats {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserStatsUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserStatsUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (usuo *UserStatsUpdateOne) sqlSave(ctx context.Context) (_node *UserStats, err error) {
	_spec := sqlgraph.NewUpdateSpec(userstats.Table, userstats.Columns, sqlgraph.NewFieldSpec(userstats.FieldID, field.TypeUUID))
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserStats.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userstats.FieldID)
		for _, f := range fields {
			if !userstats.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userstats.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usuo.mutation.UserID(); ok {
		_spec.SetField(userstats.FieldUserID, field.TypeString, value)
	}
	if value, ok := usuo.mutation.EventID(); ok {
		_spec.SetField(userstats.FieldEventID, field.TypeUUID, value)
	}
	_node = &UserStats{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userstats.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	usuo.mutation.done = true
	return _node, nil
}
