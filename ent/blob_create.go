// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/blob"
)

// BlobCreate is the builder for creating a Blob entity.
type BlobCreate struct {
	config
	mutation *BlobMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (bc *BlobCreate) SetCreateTime(t time.Time) *BlobCreate {
	bc.mutation.SetCreateTime(t)
	return bc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (bc *BlobCreate) SetNillableCreateTime(t *time.Time) *BlobCreate {
	if t != nil {
		bc.SetCreateTime(*t)
	}
	return bc
}

// SetBlob sets the "blob" field.
func (bc *BlobCreate) SetBlob(b []byte) *BlobCreate {
	bc.mutation.SetBlob(b)
	return bc
}

// SetTitle sets the "title" field.
func (bc *BlobCreate) SetTitle(s string) *BlobCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetContentType sets the "contentType" field.
func (bc *BlobCreate) SetContentType(s string) *BlobCreate {
	bc.mutation.SetContentType(s)
	return bc
}

// SetType sets the "type" field.
func (bc *BlobCreate) SetType(b blob.Type) *BlobCreate {
	bc.mutation.SetType(b)
	return bc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (bc *BlobCreate) SetNillableType(b *blob.Type) *BlobCreate {
	if b != nil {
		bc.SetType(*b)
	}
	return bc
}

// SetIsPublic sets the "isPublic" field.
func (bc *BlobCreate) SetIsPublic(b bool) *BlobCreate {
	bc.mutation.SetIsPublic(b)
	return bc
}

// SetNillableIsPublic sets the "isPublic" field if the given value is not nil.
func (bc *BlobCreate) SetNillableIsPublic(b *bool) *BlobCreate {
	if b != nil {
		bc.SetIsPublic(*b)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BlobCreate) SetID(u uuid.UUID) *BlobCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bc *BlobCreate) SetNillableID(u *uuid.UUID) *BlobCreate {
	if u != nil {
		bc.SetID(*u)
	}
	return bc
}

// Mutation returns the BlobMutation object of the builder.
func (bc *BlobCreate) Mutation() *BlobMutation {
	return bc.mutation
}

// Save creates the Blob in the database.
func (bc *BlobCreate) Save(ctx context.Context) (*Blob, error) {
	var (
		err  error
		node *Blob
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (bc *BlobCreate) SaveX(ctx context.Context) *Blob {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlobCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlobCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlobCreate) defaults() {
	if _, ok := bc.mutation.CreateTime(); !ok {
		v := blob.DefaultCreateTime()
		bc.mutation.SetCreateTime(v)
	}
	if _, ok := bc.mutation.GetType(); !ok {
		v := blob.DefaultType
		bc.mutation.SetType(v)
	}
	if _, ok := bc.mutation.IsPublic(); !ok {
		v := blob.DefaultIsPublic
		bc.mutation.SetIsPublic(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := blob.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlobCreate) check() error {
	if _, ok := bc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Blob.create_time"`)}
	}
	if _, ok := bc.mutation.Blob(); !ok {
		return &ValidationError{Name: "blob", err: errors.New(`ent: missing required field "Blob.blob"`)}
	}
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Blob.title"`)}
	}
	if _, ok := bc.mutation.ContentType(); !ok {
		return &ValidationError{Name: "contentType", err: errors.New(`ent: missing required field "Blob.contentType"`)}
	}
	if _, ok := bc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Blob.type"`)}
	}
	if v, ok := bc.mutation.GetType(); ok {
		if err := blob.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Blob.type": %w`, err)}
		}
	}
	if _, ok := bc.mutation.IsPublic(); !ok {
		return &ValidationError{Name: "isPublic", err: errors.New(`ent: missing required field "Blob.isPublic"`)}
	}
	return nil
}

func (bc *BlobCreate) sqlSave(ctx context.Context) (*Blob, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (bc *BlobCreate) createSpec() (*Blob, *sqlgraph.CreateSpec) {
	var (
		_node = &Blob{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bc.mutation.CreateTime(); ok {
		_spec.SetField(blob.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := bc.mutation.Blob(); ok {
		_spec.SetField(blob.FieldBlob, field.TypeBytes, value)
		_node.Blob = value
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(blob.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.ContentType(); ok {
		_spec.SetField(blob.FieldContentType, field.TypeString, value)
		_node.ContentType = value
	}
	if value, ok := bc.mutation.GetType(); ok {
		_spec.SetField(blob.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := bc.mutation.IsPublic(); ok {
		_spec.SetField(blob.FieldIsPublic, field.TypeBool, value)
		_node.IsPublic = value
	}
	return _node, _spec
}

// BlobCreateBulk is the builder for creating many Blob entities in bulk.
type BlobCreateBulk struct {
	config
	builders []*BlobCreate
}

// Save creates the Blob entities in the database.
func (bcb *BlobCreateBulk) Save(ctx context.Context) ([]*Blob, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blob, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlobMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlobCreateBulk) SaveX(ctx context.Context) []*Blob {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlobCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlobCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
