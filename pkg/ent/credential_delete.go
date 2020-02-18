// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/arpb2/C-3PO/pkg/ent/credential"
	"github.com/arpb2/C-3PO/pkg/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// CredentialDelete is the builder for deleting a Credential entity.
type CredentialDelete struct {
	config
	predicates []predicate.Credential
}

// Where adds a new predicate to the delete builder.
func (cd *CredentialDelete) Where(ps ...predicate.Credential) *CredentialDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CredentialDelete) Exec(ctx context.Context) (int, error) {
	return cd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CredentialDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CredentialDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: credential.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: credential.FieldID,
			},
		},
	}
	if ps := cd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
}

// CredentialDeleteOne is the builder for deleting a single Credential entity.
type CredentialDeleteOne struct {
	cd *CredentialDelete
}

// Exec executes the deletion query.
func (cdo *CredentialDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{credential.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CredentialDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
