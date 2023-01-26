// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/article"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/ent/image"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// ImageUpdate is the builder for updating Image entities.
type ImageUpdate struct {
	config
	hooks    []Hook
	mutation *ImageMutation
}

// Where appends a list predicates to the ImageUpdate builder.
func (iu *ImageUpdate) Where(ps ...predicate.Image) *ImageUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetTitle sets the "title" field.
func (iu *ImageUpdate) SetTitle(s string) *ImageUpdate {
	iu.mutation.SetTitle(s)
	return iu
}

// SetAlt sets the "alt" field.
func (iu *ImageUpdate) SetAlt(s string) *ImageUpdate {
	iu.mutation.SetAlt(s)
	return iu
}

// SetIsPreview sets the "isPreview" field.
func (iu *ImageUpdate) SetIsPreview(b bool) *ImageUpdate {
	iu.mutation.SetIsPreview(b)
	return iu
}

// SetArticleID sets the "article" edge to the Article entity by ID.
func (iu *ImageUpdate) SetArticleID(id uuid.UUID) *ImageUpdate {
	iu.mutation.SetArticleID(id)
	return iu
}

// SetNillableArticleID sets the "article" edge to the Article entity by ID if the given value is not nil.
func (iu *ImageUpdate) SetNillableArticleID(id *uuid.UUID) *ImageUpdate {
	if id != nil {
		iu = iu.SetArticleID(*id)
	}
	return iu
}

// SetArticle sets the "article" edge to the Article entity.
func (iu *ImageUpdate) SetArticle(a *Article) *ImageUpdate {
	return iu.SetArticleID(a.ID)
}

// SetBlobID sets the "blob" edge to the Blob entity by ID.
func (iu *ImageUpdate) SetBlobID(id uuid.UUID) *ImageUpdate {
	iu.mutation.SetBlobID(id)
	return iu
}

// SetBlob sets the "blob" edge to the Blob entity.
func (iu *ImageUpdate) SetBlob(b *Blob) *ImageUpdate {
	return iu.SetBlobID(b.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (iu *ImageUpdate) Mutation() *ImageMutation {
	return iu.mutation
}

// ClearArticle clears the "article" edge to the Article entity.
func (iu *ImageUpdate) ClearArticle() *ImageUpdate {
	iu.mutation.ClearArticle()
	return iu
}

// ClearBlob clears the "blob" edge to the Blob entity.
func (iu *ImageUpdate) ClearBlob() *ImageUpdate {
	iu.mutation.ClearBlob()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ImageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ImageUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ImageUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ImageUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ImageUpdate) check() error {
	if _, ok := iu.mutation.BlobID(); iu.mutation.BlobCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Image.blob"`)
	}
	return nil
}

func (iu *ImageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   image.Table,
			Columns: image.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: image.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Title(); ok {
		_spec.SetField(image.FieldTitle, field.TypeString, value)
	}
	if value, ok := iu.mutation.Alt(); ok {
		_spec.SetField(image.FieldAlt, field.TypeString, value)
	}
	if value, ok := iu.mutation.IsPreview(); ok {
		_spec.SetField(image.FieldIsPreview, field.TypeBool, value)
	}
	if iu.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.ArticleTable,
			Columns: []string{image.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.ArticleTable,
			Columns: []string{image.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.BlobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.BlobTable,
			Columns: []string{image.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.BlobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.BlobTable,
			Columns: []string{image.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ImageUpdateOne is the builder for updating a single Image entity.
type ImageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ImageMutation
}

// SetTitle sets the "title" field.
func (iuo *ImageUpdateOne) SetTitle(s string) *ImageUpdateOne {
	iuo.mutation.SetTitle(s)
	return iuo
}

// SetAlt sets the "alt" field.
func (iuo *ImageUpdateOne) SetAlt(s string) *ImageUpdateOne {
	iuo.mutation.SetAlt(s)
	return iuo
}

// SetIsPreview sets the "isPreview" field.
func (iuo *ImageUpdateOne) SetIsPreview(b bool) *ImageUpdateOne {
	iuo.mutation.SetIsPreview(b)
	return iuo
}

// SetArticleID sets the "article" edge to the Article entity by ID.
func (iuo *ImageUpdateOne) SetArticleID(id uuid.UUID) *ImageUpdateOne {
	iuo.mutation.SetArticleID(id)
	return iuo
}

// SetNillableArticleID sets the "article" edge to the Article entity by ID if the given value is not nil.
func (iuo *ImageUpdateOne) SetNillableArticleID(id *uuid.UUID) *ImageUpdateOne {
	if id != nil {
		iuo = iuo.SetArticleID(*id)
	}
	return iuo
}

// SetArticle sets the "article" edge to the Article entity.
func (iuo *ImageUpdateOne) SetArticle(a *Article) *ImageUpdateOne {
	return iuo.SetArticleID(a.ID)
}

// SetBlobID sets the "blob" edge to the Blob entity by ID.
func (iuo *ImageUpdateOne) SetBlobID(id uuid.UUID) *ImageUpdateOne {
	iuo.mutation.SetBlobID(id)
	return iuo
}

// SetBlob sets the "blob" edge to the Blob entity.
func (iuo *ImageUpdateOne) SetBlob(b *Blob) *ImageUpdateOne {
	return iuo.SetBlobID(b.ID)
}

// Mutation returns the ImageMutation object of the builder.
func (iuo *ImageUpdateOne) Mutation() *ImageMutation {
	return iuo.mutation
}

// ClearArticle clears the "article" edge to the Article entity.
func (iuo *ImageUpdateOne) ClearArticle() *ImageUpdateOne {
	iuo.mutation.ClearArticle()
	return iuo
}

// ClearBlob clears the "blob" edge to the Blob entity.
func (iuo *ImageUpdateOne) ClearBlob() *ImageUpdateOne {
	iuo.mutation.ClearBlob()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ImageUpdateOne) Select(field string, fields ...string) *ImageUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Image entity.
func (iuo *ImageUpdateOne) Save(ctx context.Context) (*Image, error) {
	var (
		err  error
		node *Image
	)
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ImageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, iuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Image)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ImageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ImageUpdateOne) SaveX(ctx context.Context) *Image {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ImageUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ImageUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ImageUpdateOne) check() error {
	if _, ok := iuo.mutation.BlobID(); iuo.mutation.BlobCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Image.blob"`)
	}
	return nil
}

func (iuo *ImageUpdateOne) sqlSave(ctx context.Context) (_node *Image, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   image.Table,
			Columns: image.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: image.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Image.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, image.FieldID)
		for _, f := range fields {
			if !image.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != image.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Title(); ok {
		_spec.SetField(image.FieldTitle, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Alt(); ok {
		_spec.SetField(image.FieldAlt, field.TypeString, value)
	}
	if value, ok := iuo.mutation.IsPreview(); ok {
		_spec.SetField(image.FieldIsPreview, field.TypeBool, value)
	}
	if iuo.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.ArticleTable,
			Columns: []string{image.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.ArticleTable,
			Columns: []string{image.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.BlobCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.BlobTable,
			Columns: []string{image.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.BlobIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   image.BlobTable,
			Columns: []string{image.BlobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Image{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{image.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
