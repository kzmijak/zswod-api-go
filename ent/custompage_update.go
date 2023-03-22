// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/attachment"
	"github.com/kzmijak/zswod_api_go/ent/custompage"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// CustomPageUpdate is the builder for updating CustomPage entities.
type CustomPageUpdate struct {
	config
	hooks    []Hook
	mutation *CustomPageMutation
}

// Where appends a list predicates to the CustomPageUpdate builder.
func (cpu *CustomPageUpdate) Where(ps ...predicate.CustomPage) *CustomPageUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetIconId sets the "iconId" field.
func (cpu *CustomPageUpdate) SetIconId(s string) *CustomPageUpdate {
	cpu.mutation.SetIconId(s)
	return cpu
}

// SetNillableIconId sets the "iconId" field if the given value is not nil.
func (cpu *CustomPageUpdate) SetNillableIconId(s *string) *CustomPageUpdate {
	if s != nil {
		cpu.SetIconId(*s)
	}
	return cpu
}

// ClearIconId clears the value of the "iconId" field.
func (cpu *CustomPageUpdate) ClearIconId() *CustomPageUpdate {
	cpu.mutation.ClearIconId()
	return cpu
}

// SetUpdateTime sets the "update_time" field.
func (cpu *CustomPageUpdate) SetUpdateTime(t time.Time) *CustomPageUpdate {
	cpu.mutation.SetUpdateTime(t)
	return cpu
}

// SetOrder sets the "order" field.
func (cpu *CustomPageUpdate) SetOrder(i int) *CustomPageUpdate {
	cpu.mutation.ResetOrder()
	cpu.mutation.SetOrder(i)
	return cpu
}

// AddOrder adds i to the "order" field.
func (cpu *CustomPageUpdate) AddOrder(i int) *CustomPageUpdate {
	cpu.mutation.AddOrder(i)
	return cpu
}

// SetTitle sets the "title" field.
func (cpu *CustomPageUpdate) SetTitle(s string) *CustomPageUpdate {
	cpu.mutation.SetTitle(s)
	return cpu
}

// SetTitleNormalized sets the "titleNormalized" field.
func (cpu *CustomPageUpdate) SetTitleNormalized(s string) *CustomPageUpdate {
	cpu.mutation.SetTitleNormalized(s)
	return cpu
}

// SetContent sets the "content" field.
func (cpu *CustomPageUpdate) SetContent(s string) *CustomPageUpdate {
	cpu.mutation.SetContent(s)
	return cpu
}

// SetIsExternal sets the "isExternal" field.
func (cpu *CustomPageUpdate) SetIsExternal(b bool) *CustomPageUpdate {
	cpu.mutation.SetIsExternal(b)
	return cpu
}

// SetNillableIsExternal sets the "isExternal" field if the given value is not nil.
func (cpu *CustomPageUpdate) SetNillableIsExternal(b *bool) *CustomPageUpdate {
	if b != nil {
		cpu.SetIsExternal(*b)
	}
	return cpu
}

// ClearIsExternal clears the value of the "isExternal" field.
func (cpu *CustomPageUpdate) ClearIsExternal() *CustomPageUpdate {
	cpu.mutation.ClearIsExternal()
	return cpu
}

// SetLink sets the "link" field.
func (cpu *CustomPageUpdate) SetLink(s string) *CustomPageUpdate {
	cpu.mutation.SetLink(s)
	return cpu
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (cpu *CustomPageUpdate) SetNillableLink(s *string) *CustomPageUpdate {
	if s != nil {
		cpu.SetLink(*s)
	}
	return cpu
}

// ClearLink clears the value of the "link" field.
func (cpu *CustomPageUpdate) ClearLink() *CustomPageUpdate {
	cpu.mutation.ClearLink()
	return cpu
}

// SetSection sets the "section" field.
func (cpu *CustomPageUpdate) SetSection(s string) *CustomPageUpdate {
	cpu.mutation.SetSection(s)
	return cpu
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by IDs.
func (cpu *CustomPageUpdate) AddAttachmentIDs(ids ...uuid.UUID) *CustomPageUpdate {
	cpu.mutation.AddAttachmentIDs(ids...)
	return cpu
}

// AddAttachments adds the "attachments" edges to the Attachment entity.
func (cpu *CustomPageUpdate) AddAttachments(a ...*Attachment) *CustomPageUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpu.AddAttachmentIDs(ids...)
}

// Mutation returns the CustomPageMutation object of the builder.
func (cpu *CustomPageUpdate) Mutation() *CustomPageMutation {
	return cpu.mutation
}

// ClearAttachments clears all "attachments" edges to the Attachment entity.
func (cpu *CustomPageUpdate) ClearAttachments() *CustomPageUpdate {
	cpu.mutation.ClearAttachments()
	return cpu
}

// RemoveAttachmentIDs removes the "attachments" edge to Attachment entities by IDs.
func (cpu *CustomPageUpdate) RemoveAttachmentIDs(ids ...uuid.UUID) *CustomPageUpdate {
	cpu.mutation.RemoveAttachmentIDs(ids...)
	return cpu
}

// RemoveAttachments removes "attachments" edges to Attachment entities.
func (cpu *CustomPageUpdate) RemoveAttachments(a ...*Attachment) *CustomPageUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpu.RemoveAttachmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *CustomPageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cpu.defaults()
	if len(cpu.hooks) == 0 {
		if err = cpu.check(); err != nil {
			return 0, err
		}
		affected, err = cpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomPageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpu.check(); err != nil {
				return 0, err
			}
			cpu.mutation = mutation
			affected, err = cpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cpu.hooks) - 1; i >= 0; i-- {
			if cpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *CustomPageUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *CustomPageUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *CustomPageUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpu *CustomPageUpdate) defaults() {
	if _, ok := cpu.mutation.UpdateTime(); !ok {
		v := custompage.UpdateDefaultUpdateTime()
		cpu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpu *CustomPageUpdate) check() error {
	if v, ok := cpu.mutation.Title(); ok {
		if err := custompage.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "CustomPage.title": %w`, err)}
		}
	}
	return nil
}

func (cpu *CustomPageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   custompage.Table,
			Columns: custompage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: custompage.FieldID,
			},
		},
	}
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpu.mutation.IconId(); ok {
		_spec.SetField(custompage.FieldIconId, field.TypeString, value)
	}
	if cpu.mutation.IconIdCleared() {
		_spec.ClearField(custompage.FieldIconId, field.TypeString)
	}
	if value, ok := cpu.mutation.UpdateTime(); ok {
		_spec.SetField(custompage.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cpu.mutation.Order(); ok {
		_spec.SetField(custompage.FieldOrder, field.TypeInt, value)
	}
	if value, ok := cpu.mutation.AddedOrder(); ok {
		_spec.AddField(custompage.FieldOrder, field.TypeInt, value)
	}
	if value, ok := cpu.mutation.Title(); ok {
		_spec.SetField(custompage.FieldTitle, field.TypeString, value)
	}
	if value, ok := cpu.mutation.TitleNormalized(); ok {
		_spec.SetField(custompage.FieldTitleNormalized, field.TypeString, value)
	}
	if value, ok := cpu.mutation.Content(); ok {
		_spec.SetField(custompage.FieldContent, field.TypeString, value)
	}
	if value, ok := cpu.mutation.IsExternal(); ok {
		_spec.SetField(custompage.FieldIsExternal, field.TypeBool, value)
	}
	if cpu.mutation.IsExternalCleared() {
		_spec.ClearField(custompage.FieldIsExternal, field.TypeBool)
	}
	if value, ok := cpu.mutation.Link(); ok {
		_spec.SetField(custompage.FieldLink, field.TypeString, value)
	}
	if cpu.mutation.LinkCleared() {
		_spec.ClearField(custompage.FieldLink, field.TypeString)
	}
	if value, ok := cpu.mutation.Section(); ok {
		_spec.SetField(custompage.FieldSection, field.TypeString, value)
	}
	if cpu.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custompage.AttachmentsTable,
			Columns: []string{custompage.AttachmentsColumn},
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
	if nodes := cpu.mutation.RemovedAttachmentsIDs(); len(nodes) > 0 && !cpu.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custompage.AttachmentsTable,
			Columns: []string{custompage.AttachmentsColumn},
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
	if nodes := cpu.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custompage.AttachmentsTable,
			Columns: []string{custompage.AttachmentsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{custompage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CustomPageUpdateOne is the builder for updating a single CustomPage entity.
type CustomPageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomPageMutation
}

// SetIconId sets the "iconId" field.
func (cpuo *CustomPageUpdateOne) SetIconId(s string) *CustomPageUpdateOne {
	cpuo.mutation.SetIconId(s)
	return cpuo
}

// SetNillableIconId sets the "iconId" field if the given value is not nil.
func (cpuo *CustomPageUpdateOne) SetNillableIconId(s *string) *CustomPageUpdateOne {
	if s != nil {
		cpuo.SetIconId(*s)
	}
	return cpuo
}

// ClearIconId clears the value of the "iconId" field.
func (cpuo *CustomPageUpdateOne) ClearIconId() *CustomPageUpdateOne {
	cpuo.mutation.ClearIconId()
	return cpuo
}

// SetUpdateTime sets the "update_time" field.
func (cpuo *CustomPageUpdateOne) SetUpdateTime(t time.Time) *CustomPageUpdateOne {
	cpuo.mutation.SetUpdateTime(t)
	return cpuo
}

// SetOrder sets the "order" field.
func (cpuo *CustomPageUpdateOne) SetOrder(i int) *CustomPageUpdateOne {
	cpuo.mutation.ResetOrder()
	cpuo.mutation.SetOrder(i)
	return cpuo
}

// AddOrder adds i to the "order" field.
func (cpuo *CustomPageUpdateOne) AddOrder(i int) *CustomPageUpdateOne {
	cpuo.mutation.AddOrder(i)
	return cpuo
}

// SetTitle sets the "title" field.
func (cpuo *CustomPageUpdateOne) SetTitle(s string) *CustomPageUpdateOne {
	cpuo.mutation.SetTitle(s)
	return cpuo
}

// SetTitleNormalized sets the "titleNormalized" field.
func (cpuo *CustomPageUpdateOne) SetTitleNormalized(s string) *CustomPageUpdateOne {
	cpuo.mutation.SetTitleNormalized(s)
	return cpuo
}

// SetContent sets the "content" field.
func (cpuo *CustomPageUpdateOne) SetContent(s string) *CustomPageUpdateOne {
	cpuo.mutation.SetContent(s)
	return cpuo
}

// SetIsExternal sets the "isExternal" field.
func (cpuo *CustomPageUpdateOne) SetIsExternal(b bool) *CustomPageUpdateOne {
	cpuo.mutation.SetIsExternal(b)
	return cpuo
}

// SetNillableIsExternal sets the "isExternal" field if the given value is not nil.
func (cpuo *CustomPageUpdateOne) SetNillableIsExternal(b *bool) *CustomPageUpdateOne {
	if b != nil {
		cpuo.SetIsExternal(*b)
	}
	return cpuo
}

// ClearIsExternal clears the value of the "isExternal" field.
func (cpuo *CustomPageUpdateOne) ClearIsExternal() *CustomPageUpdateOne {
	cpuo.mutation.ClearIsExternal()
	return cpuo
}

// SetLink sets the "link" field.
func (cpuo *CustomPageUpdateOne) SetLink(s string) *CustomPageUpdateOne {
	cpuo.mutation.SetLink(s)
	return cpuo
}

// SetNillableLink sets the "link" field if the given value is not nil.
func (cpuo *CustomPageUpdateOne) SetNillableLink(s *string) *CustomPageUpdateOne {
	if s != nil {
		cpuo.SetLink(*s)
	}
	return cpuo
}

// ClearLink clears the value of the "link" field.
func (cpuo *CustomPageUpdateOne) ClearLink() *CustomPageUpdateOne {
	cpuo.mutation.ClearLink()
	return cpuo
}

// SetSection sets the "section" field.
func (cpuo *CustomPageUpdateOne) SetSection(s string) *CustomPageUpdateOne {
	cpuo.mutation.SetSection(s)
	return cpuo
}

// AddAttachmentIDs adds the "attachments" edge to the Attachment entity by IDs.
func (cpuo *CustomPageUpdateOne) AddAttachmentIDs(ids ...uuid.UUID) *CustomPageUpdateOne {
	cpuo.mutation.AddAttachmentIDs(ids...)
	return cpuo
}

// AddAttachments adds the "attachments" edges to the Attachment entity.
func (cpuo *CustomPageUpdateOne) AddAttachments(a ...*Attachment) *CustomPageUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpuo.AddAttachmentIDs(ids...)
}

// Mutation returns the CustomPageMutation object of the builder.
func (cpuo *CustomPageUpdateOne) Mutation() *CustomPageMutation {
	return cpuo.mutation
}

// ClearAttachments clears all "attachments" edges to the Attachment entity.
func (cpuo *CustomPageUpdateOne) ClearAttachments() *CustomPageUpdateOne {
	cpuo.mutation.ClearAttachments()
	return cpuo
}

// RemoveAttachmentIDs removes the "attachments" edge to Attachment entities by IDs.
func (cpuo *CustomPageUpdateOne) RemoveAttachmentIDs(ids ...uuid.UUID) *CustomPageUpdateOne {
	cpuo.mutation.RemoveAttachmentIDs(ids...)
	return cpuo
}

// RemoveAttachments removes "attachments" edges to Attachment entities.
func (cpuo *CustomPageUpdateOne) RemoveAttachments(a ...*Attachment) *CustomPageUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cpuo.RemoveAttachmentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *CustomPageUpdateOne) Select(field string, fields ...string) *CustomPageUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated CustomPage entity.
func (cpuo *CustomPageUpdateOne) Save(ctx context.Context) (*CustomPage, error) {
	var (
		err  error
		node *CustomPage
	)
	cpuo.defaults()
	if len(cpuo.hooks) == 0 {
		if err = cpuo.check(); err != nil {
			return nil, err
		}
		node, err = cpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CustomPageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cpuo.check(); err != nil {
				return nil, err
			}
			cpuo.mutation = mutation
			node, err = cpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cpuo.hooks) - 1; i >= 0; i-- {
			if cpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CustomPage)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CustomPageMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *CustomPageUpdateOne) SaveX(ctx context.Context) *CustomPage {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *CustomPageUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *CustomPageUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpuo *CustomPageUpdateOne) defaults() {
	if _, ok := cpuo.mutation.UpdateTime(); !ok {
		v := custompage.UpdateDefaultUpdateTime()
		cpuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpuo *CustomPageUpdateOne) check() error {
	if v, ok := cpuo.mutation.Title(); ok {
		if err := custompage.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "CustomPage.title": %w`, err)}
		}
	}
	return nil
}

func (cpuo *CustomPageUpdateOne) sqlSave(ctx context.Context) (_node *CustomPage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   custompage.Table,
			Columns: custompage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: custompage.FieldID,
			},
		},
	}
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CustomPage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, custompage.FieldID)
		for _, f := range fields {
			if !custompage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != custompage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpuo.mutation.IconId(); ok {
		_spec.SetField(custompage.FieldIconId, field.TypeString, value)
	}
	if cpuo.mutation.IconIdCleared() {
		_spec.ClearField(custompage.FieldIconId, field.TypeString)
	}
	if value, ok := cpuo.mutation.UpdateTime(); ok {
		_spec.SetField(custompage.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cpuo.mutation.Order(); ok {
		_spec.SetField(custompage.FieldOrder, field.TypeInt, value)
	}
	if value, ok := cpuo.mutation.AddedOrder(); ok {
		_spec.AddField(custompage.FieldOrder, field.TypeInt, value)
	}
	if value, ok := cpuo.mutation.Title(); ok {
		_spec.SetField(custompage.FieldTitle, field.TypeString, value)
	}
	if value, ok := cpuo.mutation.TitleNormalized(); ok {
		_spec.SetField(custompage.FieldTitleNormalized, field.TypeString, value)
	}
	if value, ok := cpuo.mutation.Content(); ok {
		_spec.SetField(custompage.FieldContent, field.TypeString, value)
	}
	if value, ok := cpuo.mutation.IsExternal(); ok {
		_spec.SetField(custompage.FieldIsExternal, field.TypeBool, value)
	}
	if cpuo.mutation.IsExternalCleared() {
		_spec.ClearField(custompage.FieldIsExternal, field.TypeBool)
	}
	if value, ok := cpuo.mutation.Link(); ok {
		_spec.SetField(custompage.FieldLink, field.TypeString, value)
	}
	if cpuo.mutation.LinkCleared() {
		_spec.ClearField(custompage.FieldLink, field.TypeString)
	}
	if value, ok := cpuo.mutation.Section(); ok {
		_spec.SetField(custompage.FieldSection, field.TypeString, value)
	}
	if cpuo.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custompage.AttachmentsTable,
			Columns: []string{custompage.AttachmentsColumn},
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
	if nodes := cpuo.mutation.RemovedAttachmentsIDs(); len(nodes) > 0 && !cpuo.mutation.AttachmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custompage.AttachmentsTable,
			Columns: []string{custompage.AttachmentsColumn},
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
	if nodes := cpuo.mutation.AttachmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   custompage.AttachmentsTable,
			Columns: []string{custompage.AttachmentsColumn},
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
	_node = &CustomPage{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{custompage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
