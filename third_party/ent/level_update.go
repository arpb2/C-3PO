// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/arpb2/C-3PO/third_party/ent/level"
	"github.com/arpb2/C-3PO/third_party/ent/predicate"
	"github.com/arpb2/C-3PO/third_party/ent/schema"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// LevelUpdate is the builder for updating Level entities.
type LevelUpdate struct {
	config
	hooks      []Hook
	mutation   *LevelMutation
	predicates []predicate.Level
}

// Where adds a new predicate for the builder.
func (lu *LevelUpdate) Where(ps ...predicate.Level) *LevelUpdate {
	lu.predicates = append(lu.predicates, ps...)
	return lu
}

// SetUpdatedAt sets the updated_at field.
func (lu *LevelUpdate) SetUpdatedAt(t time.Time) *LevelUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetName sets the name field.
func (lu *LevelUpdate) SetName(s string) *LevelUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetDescription sets the description field.
func (lu *LevelUpdate) SetDescription(s string) *LevelUpdate {
	lu.mutation.SetDescription(s)
	return lu
}

// SetDefinition sets the definition field.
func (lu *LevelUpdate) SetDefinition(sd *schema.LevelDefinition) *LevelUpdate {
	lu.mutation.SetDefinition(sd)
	return lu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (lu *LevelUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := level.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
	if v, ok := lu.mutation.Name(); ok {
		if err := level.NameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if v, ok := lu.mutation.Description(); ok {
		if err := level.DescriptionValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"description\": %v", err)
		}
	}
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LevelUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LevelUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LevelUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LevelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   level.Table,
			Columns: level.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: level.FieldID,
			},
		},
	}
	if ps := lu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: level.FieldUpdatedAt,
		})
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: level.FieldName,
		})
	}
	if value, ok := lu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: level.FieldDescription,
		})
	}
	if value, ok := lu.mutation.Definition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: level.FieldDefinition,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{level.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LevelUpdateOne is the builder for updating a single Level entity.
type LevelUpdateOne struct {
	config
	hooks    []Hook
	mutation *LevelMutation
}

// SetUpdatedAt sets the updated_at field.
func (luo *LevelUpdateOne) SetUpdatedAt(t time.Time) *LevelUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetName sets the name field.
func (luo *LevelUpdateOne) SetName(s string) *LevelUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetDescription sets the description field.
func (luo *LevelUpdateOne) SetDescription(s string) *LevelUpdateOne {
	luo.mutation.SetDescription(s)
	return luo
}

// SetDefinition sets the definition field.
func (luo *LevelUpdateOne) SetDefinition(sd *schema.LevelDefinition) *LevelUpdateOne {
	luo.mutation.SetDefinition(sd)
	return luo
}

// Save executes the query and returns the updated entity.
func (luo *LevelUpdateOne) Save(ctx context.Context) (*Level, error) {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := level.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
	if v, ok := luo.mutation.Name(); ok {
		if err := level.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if v, ok := luo.mutation.Description(); ok {
		if err := level.DescriptionValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"description\": %v", err)
		}
	}
	var (
		err  error
		node *Level
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LevelUpdateOne) SaveX(ctx context.Context) *Level {
	l, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return l
}

// Exec executes the query on the entity.
func (luo *LevelUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LevelUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LevelUpdateOne) sqlSave(ctx context.Context) (l *Level, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   level.Table,
			Columns: level.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: level.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Level.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: level.FieldUpdatedAt,
		})
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: level.FieldName,
		})
	}
	if value, ok := luo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: level.FieldDescription,
		})
	}
	if value, ok := luo.mutation.Definition(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: level.FieldDefinition,
		})
	}
	l = &Level{config: luo.config}
	_spec.Assign = l.assignValues
	_spec.ScanValues = l.scanValues()
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{level.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return l, nil
}
