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
	"github.com/cufee/aftermath/internal/database/ent/db/userconnection"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserConnectionUpdate is the builder for updating UserConnection entities.
type UserConnectionUpdate struct {
	config
	hooks    []Hook
	mutation *UserConnectionMutation
}

// Where appends a list predicates to the UserConnectionUpdate builder.
func (ucu *UserConnectionUpdate) Where(ps ...predicate.UserConnection) *UserConnectionUpdate {
	ucu.mutation.Where(ps...)
	return ucu
}

// SetUpdatedAt sets the "updated_at" field.
func (ucu *UserConnectionUpdate) SetUpdatedAt(i int64) *UserConnectionUpdate {
	ucu.mutation.ResetUpdatedAt()
	ucu.mutation.SetUpdatedAt(i)
	return ucu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ucu *UserConnectionUpdate) AddUpdatedAt(i int64) *UserConnectionUpdate {
	ucu.mutation.AddUpdatedAt(i)
	return ucu
}

// SetType sets the "type" field.
func (ucu *UserConnectionUpdate) SetType(mt models.ConnectionType) *UserConnectionUpdate {
	ucu.mutation.SetType(mt)
	return ucu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ucu *UserConnectionUpdate) SetNillableType(mt *models.ConnectionType) *UserConnectionUpdate {
	if mt != nil {
		ucu.SetType(*mt)
	}
	return ucu
}

// SetReferenceID sets the "reference_id" field.
func (ucu *UserConnectionUpdate) SetReferenceID(s string) *UserConnectionUpdate {
	ucu.mutation.SetReferenceID(s)
	return ucu
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (ucu *UserConnectionUpdate) SetNillableReferenceID(s *string) *UserConnectionUpdate {
	if s != nil {
		ucu.SetReferenceID(*s)
	}
	return ucu
}

// SetPermissions sets the "permissions" field.
func (ucu *UserConnectionUpdate) SetPermissions(s string) *UserConnectionUpdate {
	ucu.mutation.SetPermissions(s)
	return ucu
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (ucu *UserConnectionUpdate) SetNillablePermissions(s *string) *UserConnectionUpdate {
	if s != nil {
		ucu.SetPermissions(*s)
	}
	return ucu
}

// ClearPermissions clears the value of the "permissions" field.
func (ucu *UserConnectionUpdate) ClearPermissions() *UserConnectionUpdate {
	ucu.mutation.ClearPermissions()
	return ucu
}

// SetMetadata sets the "metadata" field.
func (ucu *UserConnectionUpdate) SetMetadata(m map[string]interface{}) *UserConnectionUpdate {
	ucu.mutation.SetMetadata(m)
	return ucu
}

// ClearMetadata clears the value of the "metadata" field.
func (ucu *UserConnectionUpdate) ClearMetadata() *UserConnectionUpdate {
	ucu.mutation.ClearMetadata()
	return ucu
}

// Mutation returns the UserConnectionMutation object of the builder.
func (ucu *UserConnectionUpdate) Mutation() *UserConnectionMutation {
	return ucu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucu *UserConnectionUpdate) Save(ctx context.Context) (int, error) {
	ucu.defaults()
	return withHooks(ctx, ucu.sqlSave, ucu.mutation, ucu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucu *UserConnectionUpdate) SaveX(ctx context.Context) int {
	affected, err := ucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucu *UserConnectionUpdate) Exec(ctx context.Context) error {
	_, err := ucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucu *UserConnectionUpdate) ExecX(ctx context.Context) {
	if err := ucu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucu *UserConnectionUpdate) defaults() {
	if _, ok := ucu.mutation.UpdatedAt(); !ok {
		v := userconnection.UpdateDefaultUpdatedAt()
		ucu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucu *UserConnectionUpdate) check() error {
	if v, ok := ucu.mutation.GetType(); ok {
		if err := userconnection.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserConnection.type": %w`, err)}
		}
	}
	if _, ok := ucu.mutation.UserID(); ucu.mutation.UserCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "UserConnection.user"`)
	}
	return nil
}

func (ucu *UserConnectionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ucu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userconnection.Table, userconnection.Columns, sqlgraph.NewFieldSpec(userconnection.FieldID, field.TypeString))
	if ps := ucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucu.mutation.UpdatedAt(); ok {
		_spec.SetField(userconnection.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ucu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(userconnection.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ucu.mutation.GetType(); ok {
		_spec.SetField(userconnection.FieldType, field.TypeEnum, value)
	}
	if value, ok := ucu.mutation.ReferenceID(); ok {
		_spec.SetField(userconnection.FieldReferenceID, field.TypeString, value)
	}
	if value, ok := ucu.mutation.Permissions(); ok {
		_spec.SetField(userconnection.FieldPermissions, field.TypeString, value)
	}
	if ucu.mutation.PermissionsCleared() {
		_spec.ClearField(userconnection.FieldPermissions, field.TypeString)
	}
	if value, ok := ucu.mutation.Metadata(); ok {
		_spec.SetField(userconnection.FieldMetadata, field.TypeJSON, value)
	}
	if ucu.mutation.MetadataCleared() {
		_spec.ClearField(userconnection.FieldMetadata, field.TypeJSON)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userconnection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ucu.mutation.done = true
	return n, nil
}

// UserConnectionUpdateOne is the builder for updating a single UserConnection entity.
type UserConnectionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserConnectionMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ucuo *UserConnectionUpdateOne) SetUpdatedAt(i int64) *UserConnectionUpdateOne {
	ucuo.mutation.ResetUpdatedAt()
	ucuo.mutation.SetUpdatedAt(i)
	return ucuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ucuo *UserConnectionUpdateOne) AddUpdatedAt(i int64) *UserConnectionUpdateOne {
	ucuo.mutation.AddUpdatedAt(i)
	return ucuo
}

// SetType sets the "type" field.
func (ucuo *UserConnectionUpdateOne) SetType(mt models.ConnectionType) *UserConnectionUpdateOne {
	ucuo.mutation.SetType(mt)
	return ucuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ucuo *UserConnectionUpdateOne) SetNillableType(mt *models.ConnectionType) *UserConnectionUpdateOne {
	if mt != nil {
		ucuo.SetType(*mt)
	}
	return ucuo
}

// SetReferenceID sets the "reference_id" field.
func (ucuo *UserConnectionUpdateOne) SetReferenceID(s string) *UserConnectionUpdateOne {
	ucuo.mutation.SetReferenceID(s)
	return ucuo
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (ucuo *UserConnectionUpdateOne) SetNillableReferenceID(s *string) *UserConnectionUpdateOne {
	if s != nil {
		ucuo.SetReferenceID(*s)
	}
	return ucuo
}

// SetPermissions sets the "permissions" field.
func (ucuo *UserConnectionUpdateOne) SetPermissions(s string) *UserConnectionUpdateOne {
	ucuo.mutation.SetPermissions(s)
	return ucuo
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (ucuo *UserConnectionUpdateOne) SetNillablePermissions(s *string) *UserConnectionUpdateOne {
	if s != nil {
		ucuo.SetPermissions(*s)
	}
	return ucuo
}

// ClearPermissions clears the value of the "permissions" field.
func (ucuo *UserConnectionUpdateOne) ClearPermissions() *UserConnectionUpdateOne {
	ucuo.mutation.ClearPermissions()
	return ucuo
}

// SetMetadata sets the "metadata" field.
func (ucuo *UserConnectionUpdateOne) SetMetadata(m map[string]interface{}) *UserConnectionUpdateOne {
	ucuo.mutation.SetMetadata(m)
	return ucuo
}

// ClearMetadata clears the value of the "metadata" field.
func (ucuo *UserConnectionUpdateOne) ClearMetadata() *UserConnectionUpdateOne {
	ucuo.mutation.ClearMetadata()
	return ucuo
}

// Mutation returns the UserConnectionMutation object of the builder.
func (ucuo *UserConnectionUpdateOne) Mutation() *UserConnectionMutation {
	return ucuo.mutation
}

// Where appends a list predicates to the UserConnectionUpdate builder.
func (ucuo *UserConnectionUpdateOne) Where(ps ...predicate.UserConnection) *UserConnectionUpdateOne {
	ucuo.mutation.Where(ps...)
	return ucuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucuo *UserConnectionUpdateOne) Select(field string, fields ...string) *UserConnectionUpdateOne {
	ucuo.fields = append([]string{field}, fields...)
	return ucuo
}

// Save executes the query and returns the updated UserConnection entity.
func (ucuo *UserConnectionUpdateOne) Save(ctx context.Context) (*UserConnection, error) {
	ucuo.defaults()
	return withHooks(ctx, ucuo.sqlSave, ucuo.mutation, ucuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucuo *UserConnectionUpdateOne) SaveX(ctx context.Context) *UserConnection {
	node, err := ucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucuo *UserConnectionUpdateOne) Exec(ctx context.Context) error {
	_, err := ucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucuo *UserConnectionUpdateOne) ExecX(ctx context.Context) {
	if err := ucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucuo *UserConnectionUpdateOne) defaults() {
	if _, ok := ucuo.mutation.UpdatedAt(); !ok {
		v := userconnection.UpdateDefaultUpdatedAt()
		ucuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucuo *UserConnectionUpdateOne) check() error {
	if v, ok := ucuo.mutation.GetType(); ok {
		if err := userconnection.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserConnection.type": %w`, err)}
		}
	}
	if _, ok := ucuo.mutation.UserID(); ucuo.mutation.UserCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "UserConnection.user"`)
	}
	return nil
}

func (ucuo *UserConnectionUpdateOne) sqlSave(ctx context.Context) (_node *UserConnection, err error) {
	if err := ucuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userconnection.Table, userconnection.Columns, sqlgraph.NewFieldSpec(userconnection.FieldID, field.TypeString))
	id, ok := ucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "UserConnection.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userconnection.FieldID)
		for _, f := range fields {
			if !userconnection.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != userconnection.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userconnection.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ucuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(userconnection.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ucuo.mutation.GetType(); ok {
		_spec.SetField(userconnection.FieldType, field.TypeEnum, value)
	}
	if value, ok := ucuo.mutation.ReferenceID(); ok {
		_spec.SetField(userconnection.FieldReferenceID, field.TypeString, value)
	}
	if value, ok := ucuo.mutation.Permissions(); ok {
		_spec.SetField(userconnection.FieldPermissions, field.TypeString, value)
	}
	if ucuo.mutation.PermissionsCleared() {
		_spec.ClearField(userconnection.FieldPermissions, field.TypeString)
	}
	if value, ok := ucuo.mutation.Metadata(); ok {
		_spec.SetField(userconnection.FieldMetadata, field.TypeJSON, value)
	}
	if ucuo.mutation.MetadataCleared() {
		_spec.ClearField(userconnection.FieldMetadata, field.TypeJSON)
	}
	_node = &UserConnection{config: ucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userconnection.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucuo.mutation.done = true
	return _node, nil
}
