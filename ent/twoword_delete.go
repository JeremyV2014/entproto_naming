// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/twoword"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TwoWordDelete is the builder for deleting a TwoWord entity.
type TwoWordDelete struct {
	config
	hooks    []Hook
	mutation *TwoWordMutation
}

// Where appends a list predicates to the TwoWordDelete builder.
func (twd *TwoWordDelete) Where(ps ...predicate.TwoWord) *TwoWordDelete {
	twd.mutation.Where(ps...)
	return twd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (twd *TwoWordDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(twd.hooks) == 0 {
		affected, err = twd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TwoWordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			twd.mutation = mutation
			affected, err = twd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(twd.hooks) - 1; i >= 0; i-- {
			if twd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = twd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, twd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (twd *TwoWordDelete) ExecX(ctx context.Context) int {
	n, err := twd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (twd *TwoWordDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: twoword.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: twoword.FieldID,
			},
		},
	}
	if ps := twd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, twd.driver, _spec)
}

// TwoWordDeleteOne is the builder for deleting a single TwoWord entity.
type TwoWordDeleteOne struct {
	twd *TwoWordDelete
}

// Exec executes the deletion query.
func (twdo *TwoWordDeleteOne) Exec(ctx context.Context) error {
	n, err := twdo.twd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{twoword.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (twdo *TwoWordDeleteOne) ExecX(ctx context.Context) {
	twdo.twd.ExecX(ctx)
}
