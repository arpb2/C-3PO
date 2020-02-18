// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arpb2/C-3PO/pkg/ent/credential"
	"github.com/arpb2/C-3PO/pkg/ent/user"
	"github.com/arpb2/C-3PO/pkg/ent/userlevel"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	id          *uint
	email       *string
	name        *string
	surname     *string
	created_at  *time.Time
	updated_at  *time.Time
	levels      map[int]struct{}
	credentials map[int]struct{}
}

// SetEmail sets the email field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.email = &s
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// SetSurname sets the surname field.
func (uc *UserCreate) SetSurname(s string) *UserCreate {
	uc.surname = &s
	return uc
}

// SetCreatedAt sets the created_at field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.created_at = &t
	return uc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the updated_at field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.updated_at = &t
	return uc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetID sets the id field.
func (uc *UserCreate) SetID(u uint) *UserCreate {
	uc.id = &u
	return uc
}

// AddLevelIDs adds the levels edge to UserLevel by ids.
func (uc *UserCreate) AddLevelIDs(ids ...int) *UserCreate {
	if uc.levels == nil {
		uc.levels = make(map[int]struct{})
	}
	for i := range ids {
		uc.levels[ids[i]] = struct{}{}
	}
	return uc
}

// AddLevels adds the levels edges to UserLevel.
func (uc *UserCreate) AddLevels(u ...*UserLevel) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddLevelIDs(ids...)
}

// SetCredentialsID sets the credentials edge to Credential by id.
func (uc *UserCreate) SetCredentialsID(id int) *UserCreate {
	if uc.credentials == nil {
		uc.credentials = make(map[int]struct{})
	}
	uc.credentials[id] = struct{}{}
	return uc
}

// SetNillableCredentialsID sets the credentials edge to Credential by id if the given value is not nil.
func (uc *UserCreate) SetNillableCredentialsID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetCredentialsID(*id)
	}
	return uc
}

// SetCredentials sets the credentials edge to Credential.
func (uc *UserCreate) SetCredentials(c *Credential) *UserCreate {
	return uc.SetCredentialsID(c.ID)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.email == nil {
		return nil, errors.New("ent: missing required field \"email\"")
	}
	if err := user.EmailValidator(*uc.email); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"email\": %v", err)
	}
	if uc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := user.NameValidator(*uc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	if uc.surname == nil {
		return nil, errors.New("ent: missing required field \"surname\"")
	}
	if err := user.SurnameValidator(*uc.surname); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"surname\": %v", err)
	}
	if uc.created_at == nil {
		v := user.DefaultCreatedAt()
		uc.created_at = &v
	}
	if uc.updated_at == nil {
		v := user.DefaultUpdatedAt()
		uc.updated_at = &v
	}
	if len(uc.credentials) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"credentials\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint,
				Column: user.FieldID,
			},
		}
	)
	if value := uc.id; value != nil {
		u.ID = *value
		_spec.ID.Value = *value
	}
	if value := uc.email; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldEmail,
		})
		u.Email = *value
	}
	if value := uc.name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
		u.Name = *value
	}
	if value := uc.surname; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldSurname,
		})
		u.Surname = *value
	}
	if value := uc.created_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldCreatedAt,
		})
		u.CreatedAt = *value
	}
	if value := uc.updated_at; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: user.FieldUpdatedAt,
		})
		u.UpdatedAt = *value
	}
	if nodes := uc.levels; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.LevelsTable,
			Columns: []string{user.LevelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: userlevel.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.credentials; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CredentialsTable,
			Columns: []string{user.CredentialsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: credential.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if u.ID == 0 {
		id := _spec.ID.Value.(int64)
		u.ID = uint(id)
	}
	return u, nil
}
