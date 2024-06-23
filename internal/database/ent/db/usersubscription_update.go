// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/usersubscription"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserSubscriptionUpdate is the builder for updating UserSubscription entities.
type UserSubscriptionUpdate struct {
	config
	hooks    []Hook
	mutation *UserSubscriptionMutation
}

// Where appends a list predicates to the UserSubscriptionUpdate builder.
func (usu *UserSubscriptionUpdate) Where(ps ...predicate.UserSubscription) *UserSubscriptionUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetUpdatedAt sets the "updated_at" field.
func (usu *UserSubscriptionUpdate) SetUpdatedAt(i int64) *UserSubscriptionUpdate {
	usu.mutation.ResetUpdatedAt()
	usu.mutation.SetUpdatedAt(i)
	return usu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (usu *UserSubscriptionUpdate) AddUpdatedAt(i int64) *UserSubscriptionUpdate {
	usu.mutation.AddUpdatedAt(i)
	return usu
}

// SetType sets the "type" field.
func (usu *UserSubscriptionUpdate) SetType(mt models.SubscriptionType) *UserSubscriptionUpdate {
	usu.mutation.SetType(mt)
	return usu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableType(mt *models.SubscriptionType) *UserSubscriptionUpdate {
	if mt != nil {
		usu.SetType(*mt)
	}
	return usu
}

// SetExpiresAt sets the "expires_at" field.
func (usu *UserSubscriptionUpdate) SetExpiresAt(i int64) *UserSubscriptionUpdate {
	usu.mutation.ResetExpiresAt()
	usu.mutation.SetExpiresAt(i)
	return usu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableExpiresAt(i *int64) *UserSubscriptionUpdate {
	if i != nil {
		usu.SetExpiresAt(*i)
	}
	return usu
}

// AddExpiresAt adds i to the "expires_at" field.
func (usu *UserSubscriptionUpdate) AddExpiresAt(i int64) *UserSubscriptionUpdate {
	usu.mutation.AddExpiresAt(i)
	return usu
}

// SetPermissions sets the "permissions" field.
func (usu *UserSubscriptionUpdate) SetPermissions(s string) *UserSubscriptionUpdate {
	usu.mutation.SetPermissions(s)
	return usu
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillablePermissions(s *string) *UserSubscriptionUpdate {
	if s != nil {
		usu.SetPermissions(*s)
	}
	return usu
}

// SetReferenceID sets the "reference_id" field.
func (usu *UserSubscriptionUpdate) SetReferenceID(s string) *UserSubscriptionUpdate {
	usu.mutation.SetReferenceID(s)
	return usu
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (usu *UserSubscriptionUpdate) SetNillableReferenceID(s *string) *UserSubscriptionUpdate {
	if s != nil {
		usu.SetReferenceID(*s)
	}
	return usu
}

// Mutation returns the UserSubscriptionMutation object of the builder.
func (usu *UserSubscriptionUpdate) Mutation() *UserSubscriptionMutation {
	return usu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UserSubscriptionUpdate) Save(ctx context.Context) (int, error) {
	usu.defaults()
	return withHooks(ctx, usu.sqlSave, usu.mutation, usu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserSubscriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserSubscriptionUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserSubscriptionUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usu *UserSubscriptionUpdate) defaults() {
	if _, ok := usu.mutation.UpdatedAt(); !ok {
		v := usersubscription.UpdateDefaultUpdatedAt()
		usu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usu *UserSubscriptionUpdate) check() error {
	if v, ok := usu.mutation.GetType(); ok {
		if err := usersubscription.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserSubscription.type": %w`, err)}
		}
	}
	if v, ok := usu.mutation.Permissions(); ok {
		if err := usersubscription.PermissionsValidator(v); err != nil {
			return &ValidationError{Name: "permissions", err: fmt.Errorf(`db: validator failed for field "UserSubscription.permissions": %w`, err)}
		}
	}
	if v, ok := usu.mutation.ReferenceID(); ok {
		if err := usersubscription.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "UserSubscription.reference_id": %w`, err)}
		}
	}
	if _, ok := usu.mutation.UserID(); usu.mutation.UserCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "UserSubscription.user"`)
	}
	return nil
}

func (usu *UserSubscriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := usu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersubscription.Table, usersubscription.Columns, sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeString))
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usu.mutation.UpdatedAt(); ok {
		_spec.SetField(usersubscription.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := usu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(usersubscription.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := usu.mutation.GetType(); ok {
		_spec.SetField(usersubscription.FieldType, field.TypeEnum, value)
	}
	if value, ok := usu.mutation.ExpiresAt(); ok {
		_spec.SetField(usersubscription.FieldExpiresAt, field.TypeInt64, value)
	}
	if value, ok := usu.mutation.AddedExpiresAt(); ok {
		_spec.AddField(usersubscription.FieldExpiresAt, field.TypeInt64, value)
	}
	if value, ok := usu.mutation.Permissions(); ok {
		_spec.SetField(usersubscription.FieldPermissions, field.TypeString, value)
	}
	if value, ok := usu.mutation.ReferenceID(); ok {
		_spec.SetField(usersubscription.FieldReferenceID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	usu.mutation.done = true
	return n, nil
}

// UserSubscriptionUpdateOne is the builder for updating a single UserSubscription entity.
type UserSubscriptionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserSubscriptionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (usuo *UserSubscriptionUpdateOne) SetUpdatedAt(i int64) *UserSubscriptionUpdateOne {
	usuo.mutation.ResetUpdatedAt()
	usuo.mutation.SetUpdatedAt(i)
	return usuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (usuo *UserSubscriptionUpdateOne) AddUpdatedAt(i int64) *UserSubscriptionUpdateOne {
	usuo.mutation.AddUpdatedAt(i)
	return usuo
}

// SetType sets the "type" field.
func (usuo *UserSubscriptionUpdateOne) SetType(mt models.SubscriptionType) *UserSubscriptionUpdateOne {
	usuo.mutation.SetType(mt)
	return usuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableType(mt *models.SubscriptionType) *UserSubscriptionUpdateOne {
	if mt != nil {
		usuo.SetType(*mt)
	}
	return usuo
}

// SetExpiresAt sets the "expires_at" field.
func (usuo *UserSubscriptionUpdateOne) SetExpiresAt(i int64) *UserSubscriptionUpdateOne {
	usuo.mutation.ResetExpiresAt()
	usuo.mutation.SetExpiresAt(i)
	return usuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableExpiresAt(i *int64) *UserSubscriptionUpdateOne {
	if i != nil {
		usuo.SetExpiresAt(*i)
	}
	return usuo
}

// AddExpiresAt adds i to the "expires_at" field.
func (usuo *UserSubscriptionUpdateOne) AddExpiresAt(i int64) *UserSubscriptionUpdateOne {
	usuo.mutation.AddExpiresAt(i)
	return usuo
}

// SetPermissions sets the "permissions" field.
func (usuo *UserSubscriptionUpdateOne) SetPermissions(s string) *UserSubscriptionUpdateOne {
	usuo.mutation.SetPermissions(s)
	return usuo
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillablePermissions(s *string) *UserSubscriptionUpdateOne {
	if s != nil {
		usuo.SetPermissions(*s)
	}
	return usuo
}

// SetReferenceID sets the "reference_id" field.
func (usuo *UserSubscriptionUpdateOne) SetReferenceID(s string) *UserSubscriptionUpdateOne {
	usuo.mutation.SetReferenceID(s)
	return usuo
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (usuo *UserSubscriptionUpdateOne) SetNillableReferenceID(s *string) *UserSubscriptionUpdateOne {
	if s != nil {
		usuo.SetReferenceID(*s)
	}
	return usuo
}

// Mutation returns the UserSubscriptionMutation object of the builder.
func (usuo *UserSubscriptionUpdateOne) Mutation() *UserSubscriptionMutation {
	return usuo.mutation
}

// Where appends a list predicates to the UserSubscriptionUpdate builder.
func (usuo *UserSubscriptionUpdateOne) Where(ps ...predicate.UserSubscription) *UserSubscriptionUpdateOne {
	usuo.mutation.Where(ps...)
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UserSubscriptionUpdateOne) Select(field string, fields ...string) *UserSubscriptionUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UserSubscription entity.
func (usuo *UserSubscriptionUpdateOne) Save(ctx context.Context) (*UserSubscription, error) {
	usuo.defaults()
	return withHooks(ctx, usuo.sqlSave, usuo.mutation, usuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserSubscriptionUpdateOne) SaveX(ctx context.Context) *UserSubscription {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserSubscriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserSubscriptionUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usuo *UserSubscriptionUpdateOne) defaults() {
	if _, ok := usuo.mutation.UpdatedAt(); !ok {
		v := usersubscription.UpdateDefaultUpdatedAt()
		usuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usuo *UserSubscriptionUpdateOne) check() error {
	if v, ok := usuo.mutation.GetType(); ok {
		if err := usersubscription.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserSubscription.type": %w`, err)}
		}
	}
	if v, ok := usuo.mutation.Permissions(); ok {
		if err := usersubscription.PermissionsValidator(v); err != nil {
			return &ValidationError{Name: "permissions", err: fmt.Errorf(`db: validator failed for field "UserSubscription.permissions": %w`, err)}
		}
	}
	if v, ok := usuo.mutation.ReferenceID(); ok {
		if err := usersubscription.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "UserSubscription.reference_id": %w`, err)}
		}
	}
	if _, ok := usuo.mutation.UserID(); usuo.mutation.UserCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "UserSubscription.user"`)
	}
	return nil
}

func (usuo *UserSubscriptionUpdateOne) sqlSave(ctx context.Context) (_node *UserSubscription, err error) {
	if err := usuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersubscription.Table, usersubscription.Columns, sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeString))
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "UserSubscription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersubscription.FieldID)
		for _, f := range fields {
			if !usersubscription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != usersubscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := usuo.mutation.UpdatedAt(); ok {
		_spec.SetField(usersubscription.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := usuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(usersubscription.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := usuo.mutation.GetType(); ok {
		_spec.SetField(usersubscription.FieldType, field.TypeEnum, value)
	}
	if value, ok := usuo.mutation.ExpiresAt(); ok {
		_spec.SetField(usersubscription.FieldExpiresAt, field.TypeInt64, value)
	}
	if value, ok := usuo.mutation.AddedExpiresAt(); ok {
		_spec.AddField(usersubscription.FieldExpiresAt, field.TypeInt64, value)
	}
	if value, ok := usuo.mutation.Permissions(); ok {
		_spec.SetField(usersubscription.FieldPermissions, field.TypeString, value)
	}
	if value, ok := usuo.mutation.ReferenceID(); ok {
		_spec.SetField(usersubscription.FieldReferenceID, field.TypeString, value)
	}
	_node = &UserSubscription{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersubscription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	usuo.mutation.done = true
	return _node, nil
}