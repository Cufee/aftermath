// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/applicationcommand"
)

// ApplicationCommandCreate is the builder for creating a ApplicationCommand entity.
type ApplicationCommandCreate struct {
	config
	mutation *ApplicationCommandMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (acc *ApplicationCommandCreate) SetCreatedAt(i int) *ApplicationCommandCreate {
	acc.mutation.SetCreatedAt(i)
	return acc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acc *ApplicationCommandCreate) SetNillableCreatedAt(i *int) *ApplicationCommandCreate {
	if i != nil {
		acc.SetCreatedAt(*i)
	}
	return acc
}

// SetUpdatedAt sets the "updated_at" field.
func (acc *ApplicationCommandCreate) SetUpdatedAt(i int) *ApplicationCommandCreate {
	acc.mutation.SetUpdatedAt(i)
	return acc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acc *ApplicationCommandCreate) SetNillableUpdatedAt(i *int) *ApplicationCommandCreate {
	if i != nil {
		acc.SetUpdatedAt(*i)
	}
	return acc
}

// SetName sets the "name" field.
func (acc *ApplicationCommandCreate) SetName(s string) *ApplicationCommandCreate {
	acc.mutation.SetName(s)
	return acc
}

// SetVersion sets the "version" field.
func (acc *ApplicationCommandCreate) SetVersion(s string) *ApplicationCommandCreate {
	acc.mutation.SetVersion(s)
	return acc
}

// SetOptionsHash sets the "options_hash" field.
func (acc *ApplicationCommandCreate) SetOptionsHash(s string) *ApplicationCommandCreate {
	acc.mutation.SetOptionsHash(s)
	return acc
}

// SetID sets the "id" field.
func (acc *ApplicationCommandCreate) SetID(s string) *ApplicationCommandCreate {
	acc.mutation.SetID(s)
	return acc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (acc *ApplicationCommandCreate) SetNillableID(s *string) *ApplicationCommandCreate {
	if s != nil {
		acc.SetID(*s)
	}
	return acc
}

// Mutation returns the ApplicationCommandMutation object of the builder.
func (acc *ApplicationCommandCreate) Mutation() *ApplicationCommandMutation {
	return acc.mutation
}

// Save creates the ApplicationCommand in the database.
func (acc *ApplicationCommandCreate) Save(ctx context.Context) (*ApplicationCommand, error) {
	acc.defaults()
	return withHooks(ctx, acc.sqlSave, acc.mutation, acc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acc *ApplicationCommandCreate) SaveX(ctx context.Context) *ApplicationCommand {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *ApplicationCommandCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *ApplicationCommandCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *ApplicationCommandCreate) defaults() {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		v := applicationcommand.DefaultCreatedAt()
		acc.mutation.SetCreatedAt(v)
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		v := applicationcommand.DefaultUpdatedAt()
		acc.mutation.SetUpdatedAt(v)
	}
	if _, ok := acc.mutation.ID(); !ok {
		v := applicationcommand.DefaultID()
		acc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *ApplicationCommandCreate) check() error {
	if _, ok := acc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "ApplicationCommand.created_at"`)}
	}
	if _, ok := acc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "ApplicationCommand.updated_at"`)}
	}
	if _, ok := acc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "ApplicationCommand.name"`)}
	}
	if v, ok := acc.mutation.Name(); ok {
		if err := applicationcommand.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "ApplicationCommand.name": %w`, err)}
		}
	}
	if _, ok := acc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`db: missing required field "ApplicationCommand.version"`)}
	}
	if v, ok := acc.mutation.Version(); ok {
		if err := applicationcommand.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`db: validator failed for field "ApplicationCommand.version": %w`, err)}
		}
	}
	if _, ok := acc.mutation.OptionsHash(); !ok {
		return &ValidationError{Name: "options_hash", err: errors.New(`db: missing required field "ApplicationCommand.options_hash"`)}
	}
	if v, ok := acc.mutation.OptionsHash(); ok {
		if err := applicationcommand.OptionsHashValidator(v); err != nil {
			return &ValidationError{Name: "options_hash", err: fmt.Errorf(`db: validator failed for field "ApplicationCommand.options_hash": %w`, err)}
		}
	}
	return nil
}

func (acc *ApplicationCommandCreate) sqlSave(ctx context.Context) (*ApplicationCommand, error) {
	if err := acc.check(); err != nil {
		return nil, err
	}
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ApplicationCommand.ID type: %T", _spec.ID.Value)
		}
	}
	acc.mutation.id = &_node.ID
	acc.mutation.done = true
	return _node, nil
}

func (acc *ApplicationCommandCreate) createSpec() (*ApplicationCommand, *sqlgraph.CreateSpec) {
	var (
		_node = &ApplicationCommand{config: acc.config}
		_spec = sqlgraph.NewCreateSpec(applicationcommand.Table, sqlgraph.NewFieldSpec(applicationcommand.FieldID, field.TypeString))
	)
	if id, ok := acc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := acc.mutation.CreatedAt(); ok {
		_spec.SetField(applicationcommand.FieldCreatedAt, field.TypeInt, value)
		_node.CreatedAt = value
	}
	if value, ok := acc.mutation.UpdatedAt(); ok {
		_spec.SetField(applicationcommand.FieldUpdatedAt, field.TypeInt, value)
		_node.UpdatedAt = value
	}
	if value, ok := acc.mutation.Name(); ok {
		_spec.SetField(applicationcommand.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := acc.mutation.Version(); ok {
		_spec.SetField(applicationcommand.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := acc.mutation.OptionsHash(); ok {
		_spec.SetField(applicationcommand.FieldOptionsHash, field.TypeString, value)
		_node.OptionsHash = value
	}
	return _node, _spec
}

// ApplicationCommandCreateBulk is the builder for creating many ApplicationCommand entities in bulk.
type ApplicationCommandCreateBulk struct {
	config
	err      error
	builders []*ApplicationCommandCreate
}

// Save creates the ApplicationCommand entities in the database.
func (accb *ApplicationCommandCreateBulk) Save(ctx context.Context) ([]*ApplicationCommand, error) {
	if accb.err != nil {
		return nil, accb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*ApplicationCommand, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApplicationCommandMutation)
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
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *ApplicationCommandCreateBulk) SaveX(ctx context.Context) []*ApplicationCommand {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *ApplicationCommandCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *ApplicationCommandCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
