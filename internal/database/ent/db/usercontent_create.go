// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/ent/db/usercontent"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserContentCreate is the builder for creating a UserContent entity.
type UserContentCreate struct {
	config
	mutation *UserContentMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ucc *UserContentCreate) SetCreatedAt(i int64) *UserContentCreate {
	ucc.mutation.SetCreatedAt(i)
	return ucc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ucc *UserContentCreate) SetNillableCreatedAt(i *int64) *UserContentCreate {
	if i != nil {
		ucc.SetCreatedAt(*i)
	}
	return ucc
}

// SetUpdatedAt sets the "updated_at" field.
func (ucc *UserContentCreate) SetUpdatedAt(i int64) *UserContentCreate {
	ucc.mutation.SetUpdatedAt(i)
	return ucc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ucc *UserContentCreate) SetNillableUpdatedAt(i *int64) *UserContentCreate {
	if i != nil {
		ucc.SetUpdatedAt(*i)
	}
	return ucc
}

// SetType sets the "type" field.
func (ucc *UserContentCreate) SetType(mct models.UserContentType) *UserContentCreate {
	ucc.mutation.SetType(mct)
	return ucc
}

// SetUserID sets the "user_id" field.
func (ucc *UserContentCreate) SetUserID(s string) *UserContentCreate {
	ucc.mutation.SetUserID(s)
	return ucc
}

// SetReferenceID sets the "reference_id" field.
func (ucc *UserContentCreate) SetReferenceID(s string) *UserContentCreate {
	ucc.mutation.SetReferenceID(s)
	return ucc
}

// SetValue sets the "value" field.
func (ucc *UserContentCreate) SetValue(a any) *UserContentCreate {
	ucc.mutation.SetValue(a)
	return ucc
}

// SetMetadata sets the "metadata" field.
func (ucc *UserContentCreate) SetMetadata(m map[string]interface{}) *UserContentCreate {
	ucc.mutation.SetMetadata(m)
	return ucc
}

// SetID sets the "id" field.
func (ucc *UserContentCreate) SetID(s string) *UserContentCreate {
	ucc.mutation.SetID(s)
	return ucc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ucc *UserContentCreate) SetNillableID(s *string) *UserContentCreate {
	if s != nil {
		ucc.SetID(*s)
	}
	return ucc
}

// SetUser sets the "user" edge to the User entity.
func (ucc *UserContentCreate) SetUser(u *User) *UserContentCreate {
	return ucc.SetUserID(u.ID)
}

// Mutation returns the UserContentMutation object of the builder.
func (ucc *UserContentCreate) Mutation() *UserContentMutation {
	return ucc.mutation
}

// Save creates the UserContent in the database.
func (ucc *UserContentCreate) Save(ctx context.Context) (*UserContent, error) {
	ucc.defaults()
	return withHooks(ctx, ucc.sqlSave, ucc.mutation, ucc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ucc *UserContentCreate) SaveX(ctx context.Context) *UserContent {
	v, err := ucc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucc *UserContentCreate) Exec(ctx context.Context) error {
	_, err := ucc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucc *UserContentCreate) ExecX(ctx context.Context) {
	if err := ucc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ucc *UserContentCreate) defaults() {
	if _, ok := ucc.mutation.CreatedAt(); !ok {
		v := usercontent.DefaultCreatedAt()
		ucc.mutation.SetCreatedAt(v)
	}
	if _, ok := ucc.mutation.UpdatedAt(); !ok {
		v := usercontent.DefaultUpdatedAt()
		ucc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ucc.mutation.ID(); !ok {
		v := usercontent.DefaultID()
		ucc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucc *UserContentCreate) check() error {
	if _, ok := ucc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "UserContent.created_at"`)}
	}
	if _, ok := ucc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "UserContent.updated_at"`)}
	}
	if _, ok := ucc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`db: missing required field "UserContent.type"`)}
	}
	if v, ok := ucc.mutation.GetType(); ok {
		if err := usercontent.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserContent.type": %w`, err)}
		}
	}
	if _, ok := ucc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`db: missing required field "UserContent.user_id"`)}
	}
	if _, ok := ucc.mutation.ReferenceID(); !ok {
		return &ValidationError{Name: "reference_id", err: errors.New(`db: missing required field "UserContent.reference_id"`)}
	}
	if _, ok := ucc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`db: missing required field "UserContent.value"`)}
	}
	if _, ok := ucc.mutation.Metadata(); !ok {
		return &ValidationError{Name: "metadata", err: errors.New(`db: missing required field "UserContent.metadata"`)}
	}
	if _, ok := ucc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`db: missing required edge "UserContent.user"`)}
	}
	return nil
}

func (ucc *UserContentCreate) sqlSave(ctx context.Context) (*UserContent, error) {
	if err := ucc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ucc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserContent.ID type: %T", _spec.ID.Value)
		}
	}
	ucc.mutation.id = &_node.ID
	ucc.mutation.done = true
	return _node, nil
}

func (ucc *UserContentCreate) createSpec() (*UserContent, *sqlgraph.CreateSpec) {
	var (
		_node = &UserContent{config: ucc.config}
		_spec = sqlgraph.NewCreateSpec(usercontent.Table, sqlgraph.NewFieldSpec(usercontent.FieldID, field.TypeString))
	)
	if id, ok := ucc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ucc.mutation.CreatedAt(); ok {
		_spec.SetField(usercontent.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := ucc.mutation.UpdatedAt(); ok {
		_spec.SetField(usercontent.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := ucc.mutation.GetType(); ok {
		_spec.SetField(usercontent.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := ucc.mutation.ReferenceID(); ok {
		_spec.SetField(usercontent.FieldReferenceID, field.TypeString, value)
		_node.ReferenceID = value
	}
	if value, ok := ucc.mutation.Value(); ok {
		_spec.SetField(usercontent.FieldValue, field.TypeJSON, value)
		_node.Value = value
	}
	if value, ok := ucc.mutation.Metadata(); ok {
		_spec.SetField(usercontent.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	if nodes := ucc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercontent.UserTable,
			Columns: []string{usercontent.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserContentCreateBulk is the builder for creating many UserContent entities in bulk.
type UserContentCreateBulk struct {
	config
	err      error
	builders []*UserContentCreate
}

// Save creates the UserContent entities in the database.
func (uccb *UserContentCreateBulk) Save(ctx context.Context) ([]*UserContent, error) {
	if uccb.err != nil {
		return nil, uccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uccb.builders))
	nodes := make([]*UserContent, len(uccb.builders))
	mutators := make([]Mutator, len(uccb.builders))
	for i := range uccb.builders {
		func(i int, root context.Context) {
			builder := uccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserContentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uccb *UserContentCreateBulk) SaveX(ctx context.Context) []*UserContent {
	v, err := uccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uccb *UserContentCreateBulk) Exec(ctx context.Context) error {
	_, err := uccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uccb *UserContentCreateBulk) ExecX(ctx context.Context) {
	if err := uccb.Exec(ctx); err != nil {
		panic(err)
	}
}
