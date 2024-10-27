// Code generated by ent, DO NOT EDIT.

package entschema

import (
	"context"
	"example/be/gen/entschema/photographcity"
	"example/be/gen/entschema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PhotographCityDelete is the builder for deleting a PhotographCity entity.
type PhotographCityDelete struct {
	config
	hooks    []Hook
	mutation *PhotographCityMutation
}

// Where appends a list predicates to the PhotographCityDelete builder.
func (pcd *PhotographCityDelete) Where(ps ...predicate.PhotographCity) *PhotographCityDelete {
	pcd.mutation.Where(ps...)
	return pcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pcd *PhotographCityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pcd.sqlExec, pcd.mutation, pcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pcd *PhotographCityDelete) ExecX(ctx context.Context) int {
	n, err := pcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pcd *PhotographCityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(photographcity.Table, sqlgraph.NewFieldSpec(photographcity.FieldID, field.TypeInt))
	if ps := pcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pcd.mutation.done = true
	return affected, err
}

// PhotographCityDeleteOne is the builder for deleting a single PhotographCity entity.
type PhotographCityDeleteOne struct {
	pcd *PhotographCityDelete
}

// Where appends a list predicates to the PhotographCityDelete builder.
func (pcdo *PhotographCityDeleteOne) Where(ps ...predicate.PhotographCity) *PhotographCityDeleteOne {
	pcdo.pcd.mutation.Where(ps...)
	return pcdo
}

// Exec executes the deletion query.
func (pcdo *PhotographCityDeleteOne) Exec(ctx context.Context) error {
	n, err := pcdo.pcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{photographcity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pcdo *PhotographCityDeleteOne) ExecX(ctx context.Context) {
	if err := pcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
