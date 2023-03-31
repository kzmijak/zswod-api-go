// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// BlobUpdate is the builder for updating Blob entities.
type BlobUpdate struct {
	config
	hooks    []Hook
	mutation *BlobMutation
}

// Where appends a list predicates to the BlobUpdate builder.
func (bu *BlobUpdate) Where(ps ...predicate.Blob) *BlobUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetBlob sets the "blob" field.
func (bu *BlobUpdate) SetBlob(b []byte) *BlobUpdate {
	bu.mutation.SetBlob(b)
	return bu
}

// SetContentType sets the "contentType" field.
func (bu *BlobUpdate) SetContentType(s string) *BlobUpdate {
	bu.mutation.SetContentType(s)
	return bu
}

// SetType sets the "type" field.
func (bu *BlobUpdate) SetType(b blob.Type) *BlobUpdate {
	bu.mutation.SetType(b)
	return bu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bu *BlobUpdate) SetNillableType(b *blob.Type) *BlobUpdate {
	if b != nil {
		bu.SetType(*b)
	}
	return bu
}

// SetIsPublic sets the "isPublic" field.
func (bu *BlobUpdate) SetIsPublic(b bool) *BlobUpdate {
	bu.mutation.SetIsPublic(b)
	return bu
}

// SetNillableIsPublic sets the "isPublic" field if the given value is not nil.
func (bu *BlobUpdate) SetNillableIsPublic(b *bool) *BlobUpdate {
	if b != nil {
		bu.SetIsPublic(*b)
	}
	return bu
}

// Mutation returns the BlobMutation object of the builder.
func (bu *BlobUpdate) Mutation() *BlobMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlobUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlobUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlobUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlobUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlobUpdate) check() error {
	if v, ok := bu.mutation.GetType(); ok {
		if err := blob.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Blob.type": %w`, err)}
		}
	}
	return nil
}

func (bu *BlobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blob.Table,
			Columns: blob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Blob(); ok {
		_spec.SetField(blob.FieldBlob, field.TypeBytes, value)
	}
	if value, ok := bu.mutation.ContentType(); ok {
		_spec.SetField(blob.FieldContentType, field.TypeString, value)
	}
	if value, ok := bu.mutation.GetType(); ok {
		_spec.SetField(blob.FieldType, field.TypeEnum, value)
	}
	if value, ok := bu.mutation.IsPublic(); ok {
		_spec.SetField(blob.FieldIsPublic, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blob.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// BlobUpdateOne is the builder for updating a single Blob entity.
type BlobUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlobMutation
}

// SetBlob sets the "blob" field.
func (buo *BlobUpdateOne) SetBlob(b []byte) *BlobUpdateOne {
	buo.mutation.SetBlob(b)
	return buo
}

// SetContentType sets the "contentType" field.
func (buo *BlobUpdateOne) SetContentType(s string) *BlobUpdateOne {
	buo.mutation.SetContentType(s)
	return buo
}

// SetType sets the "type" field.
func (buo *BlobUpdateOne) SetType(b blob.Type) *BlobUpdateOne {
	buo.mutation.SetType(b)
	return buo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (buo *BlobUpdateOne) SetNillableType(b *blob.Type) *BlobUpdateOne {
	if b != nil {
		buo.SetType(*b)
	}
	return buo
}

// SetIsPublic sets the "isPublic" field.
func (buo *BlobUpdateOne) SetIsPublic(b bool) *BlobUpdateOne {
	buo.mutation.SetIsPublic(b)
	return buo
}

// SetNillableIsPublic sets the "isPublic" field if the given value is not nil.
func (buo *BlobUpdateOne) SetNillableIsPublic(b *bool) *BlobUpdateOne {
	if b != nil {
		buo.SetIsPublic(*b)
	}
	return buo
}

// Mutation returns the BlobMutation object of the builder.
func (buo *BlobUpdateOne) Mutation() *BlobMutation {
	return buo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlobUpdateOne) Select(field string, fields ...string) *BlobUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blob entity.
func (buo *BlobUpdateOne) Save(ctx context.Context) (*Blob, error) {
	var (
		err  error
		node *Blob
	)
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Blob)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlobMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlobUpdateOne) SaveX(ctx context.Context) *Blob {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlobUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlobUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlobUpdateOne) check() error {
	if v, ok := buo.mutation.GetType(); ok {
		if err := blob.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Blob.type": %w`, err)}
		}
	}
	return nil
}

func (buo *BlobUpdateOne) sqlSave(ctx context.Context) (_node *Blob, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blob.Table,
			Columns: blob.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blob.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blob.FieldID)
		for _, f := range fields {
			if !blob.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blob.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.Blob(); ok {
		_spec.SetField(blob.FieldBlob, field.TypeBytes, value)
	}
	if value, ok := buo.mutation.ContentType(); ok {
		_spec.SetField(blob.FieldContentType, field.TypeString, value)
	}
	if value, ok := buo.mutation.GetType(); ok {
		_spec.SetField(blob.FieldType, field.TypeEnum, value)
	}
	if value, ok := buo.mutation.IsPublic(); ok {
		_spec.SetField(blob.FieldIsPublic, field.TypeBool, value)
	}
	_node = &Blob{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blob.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
