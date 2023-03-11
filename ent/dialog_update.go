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
	"github.com/qyqx233/chat-go-api/ent/dialog"
	"github.com/qyqx233/chat-go-api/ent/predicate"
)

// DialogUpdate is the builder for updating Dialog entities.
type DialogUpdate struct {
	config
	hooks    []Hook
	mutation *DialogMutation
}

// Where appends a list predicates to the DialogUpdate builder.
func (du *DialogUpdate) Where(ps ...predicate.Dialog) *DialogUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetSid sets the "sid" field.
func (du *DialogUpdate) SetSid(i int) *DialogUpdate {
	du.mutation.ResetSid()
	du.mutation.SetSid(i)
	return du
}

// AddSid adds i to the "sid" field.
func (du *DialogUpdate) AddSid(i int) *DialogUpdate {
	du.mutation.AddSid(i)
	return du
}

// SetUser sets the "user" field.
func (du *DialogUpdate) SetUser(s string) *DialogUpdate {
	du.mutation.SetUser(s)
	return du
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (du *DialogUpdate) SetNillableUser(s *string) *DialogUpdate {
	if s != nil {
		du.SetUser(*s)
	}
	return du
}

// SetAssistant sets the "assistant" field.
func (du *DialogUpdate) SetAssistant(s string) *DialogUpdate {
	du.mutation.SetAssistant(s)
	return du
}

// SetNillableAssistant sets the "assistant" field if the given value is not nil.
func (du *DialogUpdate) SetNillableAssistant(s *string) *DialogUpdate {
	if s != nil {
		du.SetAssistant(*s)
	}
	return du
}

// SetError sets the "error" field.
func (du *DialogUpdate) SetError(s string) *DialogUpdate {
	du.mutation.SetError(s)
	return du
}

// SetNillableError sets the "error" field if the given value is not nil.
func (du *DialogUpdate) SetNillableError(s *string) *DialogUpdate {
	if s != nil {
		du.SetError(*s)
	}
	return du
}

// SetCreatedAt sets the "created_at" field.
func (du *DialogUpdate) SetCreatedAt(t time.Time) *DialogUpdate {
	du.mutation.SetCreatedAt(t)
	return du
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (du *DialogUpdate) SetNillableCreatedAt(t *time.Time) *DialogUpdate {
	if t != nil {
		du.SetCreatedAt(*t)
	}
	return du
}

// Mutation returns the DialogMutation object of the builder.
func (du *DialogUpdate) Mutation() *DialogMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DialogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, DialogMutation](ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DialogUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DialogUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DialogUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DialogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(dialog.Table, dialog.Columns, sqlgraph.NewFieldSpec(dialog.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Sid(); ok {
		_spec.SetField(dialog.FieldSid, field.TypeInt, value)
	}
	if value, ok := du.mutation.AddedSid(); ok {
		_spec.AddField(dialog.FieldSid, field.TypeInt, value)
	}
	if value, ok := du.mutation.User(); ok {
		_spec.SetField(dialog.FieldUser, field.TypeString, value)
	}
	if value, ok := du.mutation.Assistant(); ok {
		_spec.SetField(dialog.FieldAssistant, field.TypeString, value)
	}
	if value, ok := du.mutation.Error(); ok {
		_spec.SetField(dialog.FieldError, field.TypeString, value)
	}
	if value, ok := du.mutation.CreatedAt(); ok {
		_spec.SetField(dialog.FieldCreatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dialog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DialogUpdateOne is the builder for updating a single Dialog entity.
type DialogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DialogMutation
}

// SetSid sets the "sid" field.
func (duo *DialogUpdateOne) SetSid(i int) *DialogUpdateOne {
	duo.mutation.ResetSid()
	duo.mutation.SetSid(i)
	return duo
}

// AddSid adds i to the "sid" field.
func (duo *DialogUpdateOne) AddSid(i int) *DialogUpdateOne {
	duo.mutation.AddSid(i)
	return duo
}

// SetUser sets the "user" field.
func (duo *DialogUpdateOne) SetUser(s string) *DialogUpdateOne {
	duo.mutation.SetUser(s)
	return duo
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (duo *DialogUpdateOne) SetNillableUser(s *string) *DialogUpdateOne {
	if s != nil {
		duo.SetUser(*s)
	}
	return duo
}

// SetAssistant sets the "assistant" field.
func (duo *DialogUpdateOne) SetAssistant(s string) *DialogUpdateOne {
	duo.mutation.SetAssistant(s)
	return duo
}

// SetNillableAssistant sets the "assistant" field if the given value is not nil.
func (duo *DialogUpdateOne) SetNillableAssistant(s *string) *DialogUpdateOne {
	if s != nil {
		duo.SetAssistant(*s)
	}
	return duo
}

// SetError sets the "error" field.
func (duo *DialogUpdateOne) SetError(s string) *DialogUpdateOne {
	duo.mutation.SetError(s)
	return duo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (duo *DialogUpdateOne) SetNillableError(s *string) *DialogUpdateOne {
	if s != nil {
		duo.SetError(*s)
	}
	return duo
}

// SetCreatedAt sets the "created_at" field.
func (duo *DialogUpdateOne) SetCreatedAt(t time.Time) *DialogUpdateOne {
	duo.mutation.SetCreatedAt(t)
	return duo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (duo *DialogUpdateOne) SetNillableCreatedAt(t *time.Time) *DialogUpdateOne {
	if t != nil {
		duo.SetCreatedAt(*t)
	}
	return duo
}

// Mutation returns the DialogMutation object of the builder.
func (duo *DialogUpdateOne) Mutation() *DialogMutation {
	return duo.mutation
}

// Where appends a list predicates to the DialogUpdate builder.
func (duo *DialogUpdateOne) Where(ps ...predicate.Dialog) *DialogUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DialogUpdateOne) Select(field string, fields ...string) *DialogUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dialog entity.
func (duo *DialogUpdateOne) Save(ctx context.Context) (*Dialog, error) {
	return withHooks[*Dialog, DialogMutation](ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DialogUpdateOne) SaveX(ctx context.Context) *Dialog {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DialogUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DialogUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DialogUpdateOne) sqlSave(ctx context.Context) (_node *Dialog, err error) {
	_spec := sqlgraph.NewUpdateSpec(dialog.Table, dialog.Columns, sqlgraph.NewFieldSpec(dialog.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dialog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dialog.FieldID)
		for _, f := range fields {
			if !dialog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dialog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Sid(); ok {
		_spec.SetField(dialog.FieldSid, field.TypeInt, value)
	}
	if value, ok := duo.mutation.AddedSid(); ok {
		_spec.AddField(dialog.FieldSid, field.TypeInt, value)
	}
	if value, ok := duo.mutation.User(); ok {
		_spec.SetField(dialog.FieldUser, field.TypeString, value)
	}
	if value, ok := duo.mutation.Assistant(); ok {
		_spec.SetField(dialog.FieldAssistant, field.TypeString, value)
	}
	if value, ok := duo.mutation.Error(); ok {
		_spec.SetField(dialog.FieldError, field.TypeString, value)
	}
	if value, ok := duo.mutation.CreatedAt(); ok {
		_spec.SetField(dialog.FieldCreatedAt, field.TypeTime, value)
	}
	_node = &Dialog{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dialog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
