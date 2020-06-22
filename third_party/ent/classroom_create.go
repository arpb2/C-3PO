// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arpb2/C-3PO/third_party/ent/classroom"
	"github.com/arpb2/C-3PO/third_party/ent/level"
	"github.com/arpb2/C-3PO/third_party/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ClassroomCreate is the builder for creating a Classroom entity.
type ClassroomCreate struct {
	config
	mutation *ClassroomMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (cc *ClassroomCreate) SetCreatedAt(t time.Time) *ClassroomCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (cc *ClassroomCreate) SetNillableCreatedAt(t *time.Time) *ClassroomCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the updated_at field.
func (cc *ClassroomCreate) SetUpdatedAt(t time.Time) *ClassroomCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (cc *ClassroomCreate) SetNillableUpdatedAt(t *time.Time) *ClassroomCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the id field.
func (cc *ClassroomCreate) SetID(u uint) *ClassroomCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetTeacherID sets the teacher edge to User by id.
func (cc *ClassroomCreate) SetTeacherID(id uint) *ClassroomCreate {
	cc.mutation.SetTeacherID(id)
	return cc
}

// SetTeacher sets the teacher edge to User.
func (cc *ClassroomCreate) SetTeacher(u *User) *ClassroomCreate {
	return cc.SetTeacherID(u.ID)
}

// AddStudentIDs adds the students edge to User by ids.
func (cc *ClassroomCreate) AddStudentIDs(ids ...uint) *ClassroomCreate {
	cc.mutation.AddStudentIDs(ids...)
	return cc
}

// AddStudents adds the students edges to User.
func (cc *ClassroomCreate) AddStudents(u ...*User) *ClassroomCreate {
	ids := make([]uint, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddStudentIDs(ids...)
}

// SetLevelID sets the level edge to Level by id.
func (cc *ClassroomCreate) SetLevelID(id uint) *ClassroomCreate {
	cc.mutation.SetLevelID(id)
	return cc
}

// SetNillableLevelID sets the level edge to Level by id if the given value is not nil.
func (cc *ClassroomCreate) SetNillableLevelID(id *uint) *ClassroomCreate {
	if id != nil {
		cc = cc.SetLevelID(*id)
	}
	return cc
}

// SetLevel sets the level edge to Level.
func (cc *ClassroomCreate) SetLevel(l *Level) *ClassroomCreate {
	return cc.SetLevelID(l.ID)
}

// Save creates the Classroom in the database.
func (cc *ClassroomCreate) Save(ctx context.Context) (*Classroom, error) {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := classroom.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := classroom.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.TeacherID(); !ok {
		return nil, errors.New("ent: missing required edge \"teacher\"")
	}
	var (
		err  error
		node *Classroom
	)
	if len(cc.hooks) == 0 {
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ClassroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ClassroomCreate) SaveX(ctx context.Context) *Classroom {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *ClassroomCreate) sqlSave(ctx context.Context) (*Classroom, error) {
	var (
		c     = &Classroom{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: classroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: classroom.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		c.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: classroom.FieldCreatedAt,
		})
		c.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: classroom.FieldUpdatedAt,
		})
		c.UpdatedAt = value
	}
	if nodes := cc.mutation.TeacherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   classroom.TeacherTable,
			Columns: []string{classroom.TeacherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.StudentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classroom.StudentsTable,
			Columns: []string{classroom.StudentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.LevelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   classroom.LevelTable,
			Columns: []string{classroom.LevelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint,
					Column: level.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if c.ID == 0 {
		id := _spec.ID.Value.(int64)
		c.ID = uint(id)
	}
	return c, nil
}
