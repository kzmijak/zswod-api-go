// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kzmijak/zswod_api_go/ent/articletitleguid"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// ArticleTitleGuidDelete is the builder for deleting a ArticleTitleGuid entity.
type ArticleTitleGuidDelete struct {
	config
	hooks    []Hook
	mutation *ArticleTitleGuidMutation
}

// Where appends a list predicates to the ArticleTitleGuidDelete builder.
func (atgd *ArticleTitleGuidDelete) Where(ps ...predicate.ArticleTitleGuid) *ArticleTitleGuidDelete {
	atgd.mutation.Where(ps...)
	return atgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (atgd *ArticleTitleGuidDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(atgd.hooks) == 0 {
		affected, err = atgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleTitleGuidMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			atgd.mutation = mutation
			affected, err = atgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atgd.hooks) - 1; i >= 0; i-- {
			if atgd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (atgd *ArticleTitleGuidDelete) ExecX(ctx context.Context) int {
	n, err := atgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (atgd *ArticleTitleGuidDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: articletitleguid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: articletitleguid.FieldID,
			},
		},
	}
	if ps := atgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, atgd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// ArticleTitleGuidDeleteOne is the builder for deleting a single ArticleTitleGuid entity.
type ArticleTitleGuidDeleteOne struct {
	atgd *ArticleTitleGuidDelete
}

// Exec executes the deletion query.
func (atgdo *ArticleTitleGuidDeleteOne) Exec(ctx context.Context) error {
	n, err := atgdo.atgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{articletitleguid.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (atgdo *ArticleTitleGuidDeleteOne) ExecX(ctx context.Context) {
	atgdo.atgd.ExecX(ctx)
}
