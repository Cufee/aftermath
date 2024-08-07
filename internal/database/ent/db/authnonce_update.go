// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/authnonce"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// AuthNonceUpdate is the builder for updating AuthNonce entities.
type AuthNonceUpdate struct {
	config
	hooks     []Hook
	mutation  *AuthNonceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AuthNonceUpdate builder.
func (anu *AuthNonceUpdate) Where(ps ...predicate.AuthNonce) *AuthNonceUpdate {
	anu.mutation.Where(ps...)
	return anu
}

// SetUpdatedAt sets the "updated_at" field.
func (anu *AuthNonceUpdate) SetUpdatedAt(t time.Time) *AuthNonceUpdate {
	anu.mutation.SetUpdatedAt(t)
	return anu
}

// SetActive sets the "active" field.
func (anu *AuthNonceUpdate) SetActive(b bool) *AuthNonceUpdate {
	anu.mutation.SetActive(b)
	return anu
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (anu *AuthNonceUpdate) SetNillableActive(b *bool) *AuthNonceUpdate {
	if b != nil {
		anu.SetActive(*b)
	}
	return anu
}

// SetMetadata sets the "metadata" field.
func (anu *AuthNonceUpdate) SetMetadata(m map[string]string) *AuthNonceUpdate {
	anu.mutation.SetMetadata(m)
	return anu
}

// Mutation returns the AuthNonceMutation object of the builder.
func (anu *AuthNonceUpdate) Mutation() *AuthNonceMutation {
	return anu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (anu *AuthNonceUpdate) Save(ctx context.Context) (int, error) {
	anu.defaults()
	return withHooks(ctx, anu.sqlSave, anu.mutation, anu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (anu *AuthNonceUpdate) SaveX(ctx context.Context) int {
	affected, err := anu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (anu *AuthNonceUpdate) Exec(ctx context.Context) error {
	_, err := anu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (anu *AuthNonceUpdate) ExecX(ctx context.Context) {
	if err := anu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (anu *AuthNonceUpdate) defaults() {
	if _, ok := anu.mutation.UpdatedAt(); !ok {
		v := authnonce.UpdateDefaultUpdatedAt()
		anu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (anu *AuthNonceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AuthNonceUpdate {
	anu.modifiers = append(anu.modifiers, modifiers...)
	return anu
}

func (anu *AuthNonceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(authnonce.Table, authnonce.Columns, sqlgraph.NewFieldSpec(authnonce.FieldID, field.TypeString))
	if ps := anu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := anu.mutation.UpdatedAt(); ok {
		_spec.SetField(authnonce.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := anu.mutation.Active(); ok {
		_spec.SetField(authnonce.FieldActive, field.TypeBool, value)
	}
	if value, ok := anu.mutation.Metadata(); ok {
		_spec.SetField(authnonce.FieldMetadata, field.TypeJSON, value)
	}
	_spec.AddModifiers(anu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, anu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authnonce.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	anu.mutation.done = true
	return n, nil
}

// AuthNonceUpdateOne is the builder for updating a single AuthNonce entity.
type AuthNonceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AuthNonceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (anuo *AuthNonceUpdateOne) SetUpdatedAt(t time.Time) *AuthNonceUpdateOne {
	anuo.mutation.SetUpdatedAt(t)
	return anuo
}

// SetActive sets the "active" field.
func (anuo *AuthNonceUpdateOne) SetActive(b bool) *AuthNonceUpdateOne {
	anuo.mutation.SetActive(b)
	return anuo
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (anuo *AuthNonceUpdateOne) SetNillableActive(b *bool) *AuthNonceUpdateOne {
	if b != nil {
		anuo.SetActive(*b)
	}
	return anuo
}

// SetMetadata sets the "metadata" field.
func (anuo *AuthNonceUpdateOne) SetMetadata(m map[string]string) *AuthNonceUpdateOne {
	anuo.mutation.SetMetadata(m)
	return anuo
}

// Mutation returns the AuthNonceMutation object of the builder.
func (anuo *AuthNonceUpdateOne) Mutation() *AuthNonceMutation {
	return anuo.mutation
}

// Where appends a list predicates to the AuthNonceUpdate builder.
func (anuo *AuthNonceUpdateOne) Where(ps ...predicate.AuthNonce) *AuthNonceUpdateOne {
	anuo.mutation.Where(ps...)
	return anuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (anuo *AuthNonceUpdateOne) Select(field string, fields ...string) *AuthNonceUpdateOne {
	anuo.fields = append([]string{field}, fields...)
	return anuo
}

// Save executes the query and returns the updated AuthNonce entity.
func (anuo *AuthNonceUpdateOne) Save(ctx context.Context) (*AuthNonce, error) {
	anuo.defaults()
	return withHooks(ctx, anuo.sqlSave, anuo.mutation, anuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (anuo *AuthNonceUpdateOne) SaveX(ctx context.Context) *AuthNonce {
	node, err := anuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (anuo *AuthNonceUpdateOne) Exec(ctx context.Context) error {
	_, err := anuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (anuo *AuthNonceUpdateOne) ExecX(ctx context.Context) {
	if err := anuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (anuo *AuthNonceUpdateOne) defaults() {
	if _, ok := anuo.mutation.UpdatedAt(); !ok {
		v := authnonce.UpdateDefaultUpdatedAt()
		anuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (anuo *AuthNonceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AuthNonceUpdateOne {
	anuo.modifiers = append(anuo.modifiers, modifiers...)
	return anuo
}

func (anuo *AuthNonceUpdateOne) sqlSave(ctx context.Context) (_node *AuthNonce, err error) {
	_spec := sqlgraph.NewUpdateSpec(authnonce.Table, authnonce.Columns, sqlgraph.NewFieldSpec(authnonce.FieldID, field.TypeString))
	id, ok := anuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "AuthNonce.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := anuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authnonce.FieldID)
		for _, f := range fields {
			if !authnonce.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != authnonce.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := anuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := anuo.mutation.UpdatedAt(); ok {
		_spec.SetField(authnonce.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := anuo.mutation.Active(); ok {
		_spec.SetField(authnonce.FieldActive, field.TypeBool, value)
	}
	if value, ok := anuo.mutation.Metadata(); ok {
		_spec.SetField(authnonce.FieldMetadata, field.TypeJSON, value)
	}
	_spec.AddModifiers(anuo.modifiers...)
	_node = &AuthNonce{config: anuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, anuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authnonce.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	anuo.mutation.done = true
	return _node, nil
}
