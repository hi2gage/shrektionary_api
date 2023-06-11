// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"shrektionary_api/ent/definition"
	"shrektionary_api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DefinitionDelete is the builder for deleting a Definition entity.
type DefinitionDelete struct {
	config
	hooks    []Hook
	mutation *DefinitionMutation
}

// Where appends a list predicates to the DefinitionDelete builder.
func (dd *DefinitionDelete) Where(ps ...predicate.Definition) *DefinitionDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DefinitionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DefinitionDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DefinitionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(definition.Table, sqlgraph.NewFieldSpec(definition.FieldID, field.TypeInt))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DefinitionDeleteOne is the builder for deleting a single Definition entity.
type DefinitionDeleteOne struct {
	dd *DefinitionDelete
}

// Where appends a list predicates to the DefinitionDelete builder.
func (ddo *DefinitionDeleteOne) Where(ps ...predicate.Definition) *DefinitionDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DefinitionDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{definition.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DefinitionDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}