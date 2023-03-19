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
	"github.com/kzmijak/zswod_api_go/ent/attachment"
	"github.com/kzmijak/zswod_api_go/ent/blob"
	"github.com/kzmijak/zswod_api_go/ent/image"
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

// SetTitle sets the "title" field.
func (bu *BlobUpdate) SetTitle(s string) *BlobUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetAlt sets the "alt" field.
func (bu *BlobUpdate) SetAlt(s string) *BlobUpdate {
	bu.mutation.SetAlt(s)
	return bu
}

// SetContentType sets the "contentType" field.
func (bu *BlobUpdate) SetContentType(s string) *BlobUpdate {
	bu.mutation.SetContentType(s)
	return bu
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by IDs.
func (bu *BlobUpdate) AddAttachmentIDs(ids ...uuid.UUID) *BlobUpdate {
	bu.mutation.AddAttachmentIDs(ids...)
	return bu
}

// AddAttachments adds the "attachments" edges to the Attachment entity.
func (bu *BlobUpdate) AddAttachments(a ...*Attachment) *BlobUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.AddAttachmentIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (bu *BlobUpdate) AddImageIDs(ids ...uuid.UUID) *BlobUpdate {
	bu.mutation.AddImageIDs(ids...)
	return bu
}

// AddImages adds the "images" edges to the Image entity.
func (bu *BlobUpdate) AddImages(i ...*Image) *BlobUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bu.AddImageIDs(ids...)
}

// Mutation returns the BlobMutation object of the builder.
func (bu *BlobUpdate) Mutation() *BlobMutation {
	return bu.mutation
}

// ClearAttachments clears all "attachments" edges to the Attachment entity.
func (bu *BlobUpdate) ClearAttachments() *BlobUpdate {
	bu.mutation.ClearAttachments()
	return bu
}

// RemoveAttachmentIDs removes the "attachments" edge to Attachment entities by IDs.
func (bu *BlobUpdate) RemoveAttachmentIDs(ids ...uuid.UUID) *BlobUpdate {
	bu.mutation.RemoveAttachmentIDs(ids...)
	return bu
}

// RemoveAttachments removes "attachments" edges to Attachment entities.
func (bu *BlobUpdate) RemoveAttachments(a ...*Attachment) *BlobUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return bu.RemoveAttachmentIDs(ids...)
}

// ClearImages clears all "images" edges to the Image entity.
func (bu *BlobUpdate) ClearImages() *BlobUpdate {
	bu.mutation.ClearImages()
	return bu
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (bu *BlobUpdate) RemoveImageIDs(ids ...uuid.UUID) *BlobUpdate {
	bu.mutation.RemoveImageIDs(ids...)
	return bu
}

// RemoveImages removes "images" edges to Image entities.
func (bu *BlobUpdate) RemoveImages(i ...*Image) *BlobUpdate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bu.RemoveImageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlobUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(blob.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Alt(); ok {
		_spec.SetField(blob.FieldAlt, field.TypeString, value)
	}
	if value, ok := bu.mutation.ContentType(); ok {
		_spec.SetField(blob.FieldContentType, field.TypeString, value)
	}
	if bu.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.AttachmentsTable,
			Columns: []string{blob.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedAttachmentsIDs(); len(nodes) > 0 && !bu.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.AttachmentsTable,
			Columns: []string{blob.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.AttachmentsTable,
			Columns: []string{blob.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.ImagesTable,
			Columns: []string{blob.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedImagesIDs(); len(nodes) > 0 && !bu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.ImagesTable,
			Columns: []string{blob.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.ImagesTable,
			Columns: []string{blob.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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

// SetTitle sets the "title" field.
func (buo *BlobUpdateOne) SetTitle(s string) *BlobUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetAlt sets the "alt" field.
func (buo *BlobUpdateOne) SetAlt(s string) *BlobUpdateOne {
	buo.mutation.SetAlt(s)
	return buo
}

// SetContentType sets the "contentType" field.
func (buo *BlobUpdateOne) SetContentType(s string) *BlobUpdateOne {
	buo.mutation.SetContentType(s)
	return buo
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by IDs.
func (buo *BlobUpdateOne) AddAttachmentIDs(ids ...uuid.UUID) *BlobUpdateOne {
	buo.mutation.AddAttachmentIDs(ids...)
	return buo
}

// AddAttachments adds the "attachments" edges to the Attachment entity.
func (buo *BlobUpdateOne) AddAttachments(a ...*Attachment) *BlobUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.AddAttachmentIDs(ids...)
}

// AddImageIDs adds the "images" edge to the Image entity by IDs.
func (buo *BlobUpdateOne) AddImageIDs(ids ...uuid.UUID) *BlobUpdateOne {
	buo.mutation.AddImageIDs(ids...)
	return buo
}

// AddImages adds the "images" edges to the Image entity.
func (buo *BlobUpdateOne) AddImages(i ...*Image) *BlobUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return buo.AddImageIDs(ids...)
}

// Mutation returns the BlobMutation object of the builder.
func (buo *BlobUpdateOne) Mutation() *BlobMutation {
	return buo.mutation
}

// ClearAttachments clears all "attachments" edges to the Attachment entity.
func (buo *BlobUpdateOne) ClearAttachments() *BlobUpdateOne {
	buo.mutation.ClearAttachments()
	return buo
}

// RemoveAttachmentIDs removes the "attachments" edge to Attachment entities by IDs.
func (buo *BlobUpdateOne) RemoveAttachmentIDs(ids ...uuid.UUID) *BlobUpdateOne {
	buo.mutation.RemoveAttachmentIDs(ids...)
	return buo
}

// RemoveAttachments removes "attachments" edges to Attachment entities.
func (buo *BlobUpdateOne) RemoveAttachments(a ...*Attachment) *BlobUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return buo.RemoveAttachmentIDs(ids...)
}

// ClearImages clears all "images" edges to the Image entity.
func (buo *BlobUpdateOne) ClearImages() *BlobUpdateOne {
	buo.mutation.ClearImages()
	return buo
}

// RemoveImageIDs removes the "images" edge to Image entities by IDs.
func (buo *BlobUpdateOne) RemoveImageIDs(ids ...uuid.UUID) *BlobUpdateOne {
	buo.mutation.RemoveImageIDs(ids...)
	return buo
}

// RemoveImages removes "images" edges to Image entities.
func (buo *BlobUpdateOne) RemoveImages(i ...*Image) *BlobUpdateOne {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return buo.RemoveImageIDs(ids...)
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
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(blob.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Alt(); ok {
		_spec.SetField(blob.FieldAlt, field.TypeString, value)
	}
	if value, ok := buo.mutation.ContentType(); ok {
		_spec.SetField(blob.FieldContentType, field.TypeString, value)
	}
	if buo.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.AttachmentsTable,
			Columns: []string{blob.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedAttachmentsIDs(); len(nodes) > 0 && !buo.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.AttachmentsTable,
			Columns: []string{blob.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.AttachmentsTable,
			Columns: []string{blob.AttachmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: attachment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.ImagesTable,
			Columns: []string{blob.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedImagesIDs(); len(nodes) > 0 && !buo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.ImagesTable,
			Columns: []string{blob.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   blob.ImagesTable,
			Columns: []string{blob.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
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
