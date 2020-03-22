// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/arpb2/C-3PO/third_party/ent/predicate"
	"github.com/arpb2/C-3PO/third_party/ent/userlevel"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserLevelDelete is the builder for deleting a UserLevel entity.
type UserLevelDelete struct {
	config
	hooks      []Hook
	mutation   *UserLevelMutation
	predicates []predicate.UserLevel
}

// Where adds a new predicate to the delete builder.
func (uld *UserLevelDelete) Where(ps ...predicate.UserLevel) *UserLevelDelete {
	uld.predicates = append(uld.predicates, ps...)
	return uld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uld *UserLevelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uld.hooks) == 0 {
		affected, err = uld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uld.mutation = mutation
			affected, err = uld.sqlExec(ctx)
			return affected, err
		})
		for i := len(uld.hooks) - 1; i >= 0; i-- {
			mut = uld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (uld *UserLevelDelete) ExecX(ctx context.Context) int {
	n, err := uld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uld *UserLevelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userlevel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userlevel.FieldID,
			},
		},
	}
	if ps := uld.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uld.driver, _spec)
}

// UserLevelDeleteOne is the builder for deleting a single UserLevel entity.
type UserLevelDeleteOne struct {
	uld *UserLevelDelete
}

// Exec executes the deletion query.
func (uldo *UserLevelDeleteOne) Exec(ctx context.Context) error {
	n, err := uldo.uld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userlevel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uldo *UserLevelDeleteOne) ExecX(ctx context.Context) {
	uldo.uld.ExecX(ctx)
}
