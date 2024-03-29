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
	"github.com/kzmijak/zswod_api_go/ent/predicate"
	"github.com/kzmijak/zswod_api_go/ent/resetpasswordtoken"
	"github.com/kzmijak/zswod_api_go/ent/user"
)

// ResetPasswordTokenUpdate is the builder for updating ResetPasswordToken entities.
type ResetPasswordTokenUpdate struct {
	config
	hooks    []Hook
	mutation *ResetPasswordTokenMutation
}

// Where appends a list predicates to the ResetPasswordTokenUpdate builder.
func (rptu *ResetPasswordTokenUpdate) Where(ps ...predicate.ResetPasswordToken) *ResetPasswordTokenUpdate {
	rptu.mutation.Where(ps...)
	return rptu
}

// SetUpdateTime sets the "update_time" field.
func (rptu *ResetPasswordTokenUpdate) SetUpdateTime(t time.Time) *ResetPasswordTokenUpdate {
	rptu.mutation.SetUpdateTime(t)
	return rptu
}

// SetExpiryTime sets the "expiryTime" field.
func (rptu *ResetPasswordTokenUpdate) SetExpiryTime(t time.Time) *ResetPasswordTokenUpdate {
	rptu.mutation.SetExpiryTime(t)
	return rptu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (rptu *ResetPasswordTokenUpdate) SetOwnerID(id uuid.UUID) *ResetPasswordTokenUpdate {
	rptu.mutation.SetOwnerID(id)
	return rptu
}

// SetOwner sets the "owner" edge to the User entity.
func (rptu *ResetPasswordTokenUpdate) SetOwner(u *User) *ResetPasswordTokenUpdate {
	return rptu.SetOwnerID(u.ID)
}

// Mutation returns the ResetPasswordTokenMutation object of the builder.
func (rptu *ResetPasswordTokenUpdate) Mutation() *ResetPasswordTokenMutation {
	return rptu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (rptu *ResetPasswordTokenUpdate) ClearOwner() *ResetPasswordTokenUpdate {
	rptu.mutation.ClearOwner()
	return rptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rptu *ResetPasswordTokenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	rptu.defaults()
	if len(rptu.hooks) == 0 {
		if err = rptu.check(); err != nil {
			return 0, err
		}
		affected, err = rptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResetPasswordTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rptu.check(); err != nil {
				return 0, err
			}
			rptu.mutation = mutation
			affected, err = rptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rptu.hooks) - 1; i >= 0; i-- {
			if rptu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rptu *ResetPasswordTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := rptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rptu *ResetPasswordTokenUpdate) Exec(ctx context.Context) error {
	_, err := rptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rptu *ResetPasswordTokenUpdate) ExecX(ctx context.Context) {
	if err := rptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rptu *ResetPasswordTokenUpdate) defaults() {
	if _, ok := rptu.mutation.UpdateTime(); !ok {
		v := resetpasswordtoken.UpdateDefaultUpdateTime()
		rptu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rptu *ResetPasswordTokenUpdate) check() error {
	if _, ok := rptu.mutation.OwnerID(); rptu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ResetPasswordToken.owner"`)
	}
	return nil
}

func (rptu *ResetPasswordTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resetpasswordtoken.Table,
			Columns: resetpasswordtoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: resetpasswordtoken.FieldID,
			},
		},
	}
	if ps := rptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rptu.mutation.UpdateTime(); ok {
		_spec.SetField(resetpasswordtoken.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := rptu.mutation.ExpiryTime(); ok {
		_spec.SetField(resetpasswordtoken.FieldExpiryTime, field.TypeTime, value)
	}
	if rptu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resetpasswordtoken.OwnerTable,
			Columns: []string{resetpasswordtoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rptu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resetpasswordtoken.OwnerTable,
			Columns: []string{resetpasswordtoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resetpasswordtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ResetPasswordTokenUpdateOne is the builder for updating a single ResetPasswordToken entity.
type ResetPasswordTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResetPasswordTokenMutation
}

// SetUpdateTime sets the "update_time" field.
func (rptuo *ResetPasswordTokenUpdateOne) SetUpdateTime(t time.Time) *ResetPasswordTokenUpdateOne {
	rptuo.mutation.SetUpdateTime(t)
	return rptuo
}

// SetExpiryTime sets the "expiryTime" field.
func (rptuo *ResetPasswordTokenUpdateOne) SetExpiryTime(t time.Time) *ResetPasswordTokenUpdateOne {
	rptuo.mutation.SetExpiryTime(t)
	return rptuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (rptuo *ResetPasswordTokenUpdateOne) SetOwnerID(id uuid.UUID) *ResetPasswordTokenUpdateOne {
	rptuo.mutation.SetOwnerID(id)
	return rptuo
}

// SetOwner sets the "owner" edge to the User entity.
func (rptuo *ResetPasswordTokenUpdateOne) SetOwner(u *User) *ResetPasswordTokenUpdateOne {
	return rptuo.SetOwnerID(u.ID)
}

// Mutation returns the ResetPasswordTokenMutation object of the builder.
func (rptuo *ResetPasswordTokenUpdateOne) Mutation() *ResetPasswordTokenMutation {
	return rptuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (rptuo *ResetPasswordTokenUpdateOne) ClearOwner() *ResetPasswordTokenUpdateOne {
	rptuo.mutation.ClearOwner()
	return rptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rptuo *ResetPasswordTokenUpdateOne) Select(field string, fields ...string) *ResetPasswordTokenUpdateOne {
	rptuo.fields = append([]string{field}, fields...)
	return rptuo
}

// Save executes the query and returns the updated ResetPasswordToken entity.
func (rptuo *ResetPasswordTokenUpdateOne) Save(ctx context.Context) (*ResetPasswordToken, error) {
	var (
		err  error
		node *ResetPasswordToken
	)
	rptuo.defaults()
	if len(rptuo.hooks) == 0 {
		if err = rptuo.check(); err != nil {
			return nil, err
		}
		node, err = rptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResetPasswordTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rptuo.check(); err != nil {
				return nil, err
			}
			rptuo.mutation = mutation
			node, err = rptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rptuo.hooks) - 1; i >= 0; i-- {
			if rptuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rptuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rptuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ResetPasswordToken)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ResetPasswordTokenMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rptuo *ResetPasswordTokenUpdateOne) SaveX(ctx context.Context) *ResetPasswordToken {
	node, err := rptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rptuo *ResetPasswordTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := rptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rptuo *ResetPasswordTokenUpdateOne) ExecX(ctx context.Context) {
	if err := rptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rptuo *ResetPasswordTokenUpdateOne) defaults() {
	if _, ok := rptuo.mutation.UpdateTime(); !ok {
		v := resetpasswordtoken.UpdateDefaultUpdateTime()
		rptuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rptuo *ResetPasswordTokenUpdateOne) check() error {
	if _, ok := rptuo.mutation.OwnerID(); rptuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ResetPasswordToken.owner"`)
	}
	return nil
}

func (rptuo *ResetPasswordTokenUpdateOne) sqlSave(ctx context.Context) (_node *ResetPasswordToken, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resetpasswordtoken.Table,
			Columns: resetpasswordtoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: resetpasswordtoken.FieldID,
			},
		},
	}
	id, ok := rptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ResetPasswordToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resetpasswordtoken.FieldID)
		for _, f := range fields {
			if !resetpasswordtoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != resetpasswordtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rptuo.mutation.UpdateTime(); ok {
		_spec.SetField(resetpasswordtoken.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := rptuo.mutation.ExpiryTime(); ok {
		_spec.SetField(resetpasswordtoken.FieldExpiryTime, field.TypeTime, value)
	}
	if rptuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resetpasswordtoken.OwnerTable,
			Columns: []string{resetpasswordtoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rptuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   resetpasswordtoken.OwnerTable,
			Columns: []string{resetpasswordtoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ResetPasswordToken{config: rptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resetpasswordtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
