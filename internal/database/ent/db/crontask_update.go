// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/crontask"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
)

// CronTaskUpdate is the builder for updating CronTask entities.
type CronTaskUpdate struct {
	config
	hooks     []Hook
	mutation  *CronTaskMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CronTaskUpdate builder.
func (ctu *CronTaskUpdate) Where(ps ...predicate.CronTask) *CronTaskUpdate {
	ctu.mutation.Where(ps...)
	return ctu
}

// SetUpdatedAt sets the "updated_at" field.
func (ctu *CronTaskUpdate) SetUpdatedAt(t time.Time) *CronTaskUpdate {
	ctu.mutation.SetUpdatedAt(t)
	return ctu
}

// SetType sets the "type" field.
func (ctu *CronTaskUpdate) SetType(mt models.TaskType) *CronTaskUpdate {
	ctu.mutation.SetType(mt)
	return ctu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ctu *CronTaskUpdate) SetNillableType(mt *models.TaskType) *CronTaskUpdate {
	if mt != nil {
		ctu.SetType(*mt)
	}
	return ctu
}

// SetReferenceID sets the "reference_id" field.
func (ctu *CronTaskUpdate) SetReferenceID(s string) *CronTaskUpdate {
	ctu.mutation.SetReferenceID(s)
	return ctu
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (ctu *CronTaskUpdate) SetNillableReferenceID(s *string) *CronTaskUpdate {
	if s != nil {
		ctu.SetReferenceID(*s)
	}
	return ctu
}

// SetTargets sets the "targets" field.
func (ctu *CronTaskUpdate) SetTargets(s []string) *CronTaskUpdate {
	ctu.mutation.SetTargets(s)
	return ctu
}

// AppendTargets appends s to the "targets" field.
func (ctu *CronTaskUpdate) AppendTargets(s []string) *CronTaskUpdate {
	ctu.mutation.AppendTargets(s)
	return ctu
}

// SetStatus sets the "status" field.
func (ctu *CronTaskUpdate) SetStatus(ms models.TaskStatus) *CronTaskUpdate {
	ctu.mutation.SetStatus(ms)
	return ctu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ctu *CronTaskUpdate) SetNillableStatus(ms *models.TaskStatus) *CronTaskUpdate {
	if ms != nil {
		ctu.SetStatus(*ms)
	}
	return ctu
}

// SetScheduledAfter sets the "scheduled_after" field.
func (ctu *CronTaskUpdate) SetScheduledAfter(t time.Time) *CronTaskUpdate {
	ctu.mutation.SetScheduledAfter(t)
	return ctu
}

// SetNillableScheduledAfter sets the "scheduled_after" field if the given value is not nil.
func (ctu *CronTaskUpdate) SetNillableScheduledAfter(t *time.Time) *CronTaskUpdate {
	if t != nil {
		ctu.SetScheduledAfter(*t)
	}
	return ctu
}

// SetLastRun sets the "last_run" field.
func (ctu *CronTaskUpdate) SetLastRun(t time.Time) *CronTaskUpdate {
	ctu.mutation.SetLastRun(t)
	return ctu
}

// SetNillableLastRun sets the "last_run" field if the given value is not nil.
func (ctu *CronTaskUpdate) SetNillableLastRun(t *time.Time) *CronTaskUpdate {
	if t != nil {
		ctu.SetLastRun(*t)
	}
	return ctu
}

// SetLogs sets the "logs" field.
func (ctu *CronTaskUpdate) SetLogs(ml []models.TaskLog) *CronTaskUpdate {
	ctu.mutation.SetLogs(ml)
	return ctu
}

// AppendLogs appends ml to the "logs" field.
func (ctu *CronTaskUpdate) AppendLogs(ml []models.TaskLog) *CronTaskUpdate {
	ctu.mutation.AppendLogs(ml)
	return ctu
}

// SetData sets the "data" field.
func (ctu *CronTaskUpdate) SetData(m map[string]interface{}) *CronTaskUpdate {
	ctu.mutation.SetData(m)
	return ctu
}

// Mutation returns the CronTaskMutation object of the builder.
func (ctu *CronTaskUpdate) Mutation() *CronTaskMutation {
	return ctu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ctu *CronTaskUpdate) Save(ctx context.Context) (int, error) {
	ctu.defaults()
	return withHooks(ctx, ctu.sqlSave, ctu.mutation, ctu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctu *CronTaskUpdate) SaveX(ctx context.Context) int {
	affected, err := ctu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ctu *CronTaskUpdate) Exec(ctx context.Context) error {
	_, err := ctu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctu *CronTaskUpdate) ExecX(ctx context.Context) {
	if err := ctu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctu *CronTaskUpdate) defaults() {
	if _, ok := ctu.mutation.UpdatedAt(); !ok {
		v := crontask.UpdateDefaultUpdatedAt()
		ctu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctu *CronTaskUpdate) check() error {
	if v, ok := ctu.mutation.GetType(); ok {
		if err := crontask.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "CronTask.type": %w`, err)}
		}
	}
	if v, ok := ctu.mutation.ReferenceID(); ok {
		if err := crontask.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "CronTask.reference_id": %w`, err)}
		}
	}
	if v, ok := ctu.mutation.Status(); ok {
		if err := crontask.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "CronTask.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ctu *CronTaskUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CronTaskUpdate {
	ctu.modifiers = append(ctu.modifiers, modifiers...)
	return ctu
}

func (ctu *CronTaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ctu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(crontask.Table, crontask.Columns, sqlgraph.NewFieldSpec(crontask.FieldID, field.TypeString))
	if ps := ctu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctu.mutation.UpdatedAt(); ok {
		_spec.SetField(crontask.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ctu.mutation.GetType(); ok {
		_spec.SetField(crontask.FieldType, field.TypeEnum, value)
	}
	if value, ok := ctu.mutation.ReferenceID(); ok {
		_spec.SetField(crontask.FieldReferenceID, field.TypeString, value)
	}
	if value, ok := ctu.mutation.Targets(); ok {
		_spec.SetField(crontask.FieldTargets, field.TypeJSON, value)
	}
	if value, ok := ctu.mutation.AppendedTargets(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, crontask.FieldTargets, value)
		})
	}
	if value, ok := ctu.mutation.Status(); ok {
		_spec.SetField(crontask.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ctu.mutation.ScheduledAfter(); ok {
		_spec.SetField(crontask.FieldScheduledAfter, field.TypeTime, value)
	}
	if value, ok := ctu.mutation.LastRun(); ok {
		_spec.SetField(crontask.FieldLastRun, field.TypeTime, value)
	}
	if value, ok := ctu.mutation.Logs(); ok {
		_spec.SetField(crontask.FieldLogs, field.TypeJSON, value)
	}
	if value, ok := ctu.mutation.AppendedLogs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, crontask.FieldLogs, value)
		})
	}
	if value, ok := ctu.mutation.Data(); ok {
		_spec.SetField(crontask.FieldData, field.TypeJSON, value)
	}
	_spec.AddModifiers(ctu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ctu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{crontask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ctu.mutation.done = true
	return n, nil
}

// CronTaskUpdateOne is the builder for updating a single CronTask entity.
type CronTaskUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CronTaskMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (ctuo *CronTaskUpdateOne) SetUpdatedAt(t time.Time) *CronTaskUpdateOne {
	ctuo.mutation.SetUpdatedAt(t)
	return ctuo
}

// SetType sets the "type" field.
func (ctuo *CronTaskUpdateOne) SetType(mt models.TaskType) *CronTaskUpdateOne {
	ctuo.mutation.SetType(mt)
	return ctuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ctuo *CronTaskUpdateOne) SetNillableType(mt *models.TaskType) *CronTaskUpdateOne {
	if mt != nil {
		ctuo.SetType(*mt)
	}
	return ctuo
}

// SetReferenceID sets the "reference_id" field.
func (ctuo *CronTaskUpdateOne) SetReferenceID(s string) *CronTaskUpdateOne {
	ctuo.mutation.SetReferenceID(s)
	return ctuo
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (ctuo *CronTaskUpdateOne) SetNillableReferenceID(s *string) *CronTaskUpdateOne {
	if s != nil {
		ctuo.SetReferenceID(*s)
	}
	return ctuo
}

// SetTargets sets the "targets" field.
func (ctuo *CronTaskUpdateOne) SetTargets(s []string) *CronTaskUpdateOne {
	ctuo.mutation.SetTargets(s)
	return ctuo
}

// AppendTargets appends s to the "targets" field.
func (ctuo *CronTaskUpdateOne) AppendTargets(s []string) *CronTaskUpdateOne {
	ctuo.mutation.AppendTargets(s)
	return ctuo
}

// SetStatus sets the "status" field.
func (ctuo *CronTaskUpdateOne) SetStatus(ms models.TaskStatus) *CronTaskUpdateOne {
	ctuo.mutation.SetStatus(ms)
	return ctuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ctuo *CronTaskUpdateOne) SetNillableStatus(ms *models.TaskStatus) *CronTaskUpdateOne {
	if ms != nil {
		ctuo.SetStatus(*ms)
	}
	return ctuo
}

// SetScheduledAfter sets the "scheduled_after" field.
func (ctuo *CronTaskUpdateOne) SetScheduledAfter(t time.Time) *CronTaskUpdateOne {
	ctuo.mutation.SetScheduledAfter(t)
	return ctuo
}

// SetNillableScheduledAfter sets the "scheduled_after" field if the given value is not nil.
func (ctuo *CronTaskUpdateOne) SetNillableScheduledAfter(t *time.Time) *CronTaskUpdateOne {
	if t != nil {
		ctuo.SetScheduledAfter(*t)
	}
	return ctuo
}

// SetLastRun sets the "last_run" field.
func (ctuo *CronTaskUpdateOne) SetLastRun(t time.Time) *CronTaskUpdateOne {
	ctuo.mutation.SetLastRun(t)
	return ctuo
}

// SetNillableLastRun sets the "last_run" field if the given value is not nil.
func (ctuo *CronTaskUpdateOne) SetNillableLastRun(t *time.Time) *CronTaskUpdateOne {
	if t != nil {
		ctuo.SetLastRun(*t)
	}
	return ctuo
}

// SetLogs sets the "logs" field.
func (ctuo *CronTaskUpdateOne) SetLogs(ml []models.TaskLog) *CronTaskUpdateOne {
	ctuo.mutation.SetLogs(ml)
	return ctuo
}

// AppendLogs appends ml to the "logs" field.
func (ctuo *CronTaskUpdateOne) AppendLogs(ml []models.TaskLog) *CronTaskUpdateOne {
	ctuo.mutation.AppendLogs(ml)
	return ctuo
}

// SetData sets the "data" field.
func (ctuo *CronTaskUpdateOne) SetData(m map[string]interface{}) *CronTaskUpdateOne {
	ctuo.mutation.SetData(m)
	return ctuo
}

// Mutation returns the CronTaskMutation object of the builder.
func (ctuo *CronTaskUpdateOne) Mutation() *CronTaskMutation {
	return ctuo.mutation
}

// Where appends a list predicates to the CronTaskUpdate builder.
func (ctuo *CronTaskUpdateOne) Where(ps ...predicate.CronTask) *CronTaskUpdateOne {
	ctuo.mutation.Where(ps...)
	return ctuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ctuo *CronTaskUpdateOne) Select(field string, fields ...string) *CronTaskUpdateOne {
	ctuo.fields = append([]string{field}, fields...)
	return ctuo
}

// Save executes the query and returns the updated CronTask entity.
func (ctuo *CronTaskUpdateOne) Save(ctx context.Context) (*CronTask, error) {
	ctuo.defaults()
	return withHooks(ctx, ctuo.sqlSave, ctuo.mutation, ctuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ctuo *CronTaskUpdateOne) SaveX(ctx context.Context) *CronTask {
	node, err := ctuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ctuo *CronTaskUpdateOne) Exec(ctx context.Context) error {
	_, err := ctuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctuo *CronTaskUpdateOne) ExecX(ctx context.Context) {
	if err := ctuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ctuo *CronTaskUpdateOne) defaults() {
	if _, ok := ctuo.mutation.UpdatedAt(); !ok {
		v := crontask.UpdateDefaultUpdatedAt()
		ctuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctuo *CronTaskUpdateOne) check() error {
	if v, ok := ctuo.mutation.GetType(); ok {
		if err := crontask.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "CronTask.type": %w`, err)}
		}
	}
	if v, ok := ctuo.mutation.ReferenceID(); ok {
		if err := crontask.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "CronTask.reference_id": %w`, err)}
		}
	}
	if v, ok := ctuo.mutation.Status(); ok {
		if err := crontask.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "CronTask.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ctuo *CronTaskUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CronTaskUpdateOne {
	ctuo.modifiers = append(ctuo.modifiers, modifiers...)
	return ctuo
}

func (ctuo *CronTaskUpdateOne) sqlSave(ctx context.Context) (_node *CronTask, err error) {
	if err := ctuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(crontask.Table, crontask.Columns, sqlgraph.NewFieldSpec(crontask.FieldID, field.TypeString))
	id, ok := ctuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "CronTask.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ctuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, crontask.FieldID)
		for _, f := range fields {
			if !crontask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != crontask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ctuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ctuo.mutation.UpdatedAt(); ok {
		_spec.SetField(crontask.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ctuo.mutation.GetType(); ok {
		_spec.SetField(crontask.FieldType, field.TypeEnum, value)
	}
	if value, ok := ctuo.mutation.ReferenceID(); ok {
		_spec.SetField(crontask.FieldReferenceID, field.TypeString, value)
	}
	if value, ok := ctuo.mutation.Targets(); ok {
		_spec.SetField(crontask.FieldTargets, field.TypeJSON, value)
	}
	if value, ok := ctuo.mutation.AppendedTargets(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, crontask.FieldTargets, value)
		})
	}
	if value, ok := ctuo.mutation.Status(); ok {
		_spec.SetField(crontask.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ctuo.mutation.ScheduledAfter(); ok {
		_spec.SetField(crontask.FieldScheduledAfter, field.TypeTime, value)
	}
	if value, ok := ctuo.mutation.LastRun(); ok {
		_spec.SetField(crontask.FieldLastRun, field.TypeTime, value)
	}
	if value, ok := ctuo.mutation.Logs(); ok {
		_spec.SetField(crontask.FieldLogs, field.TypeJSON, value)
	}
	if value, ok := ctuo.mutation.AppendedLogs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, crontask.FieldLogs, value)
		})
	}
	if value, ok := ctuo.mutation.Data(); ok {
		_spec.SetField(crontask.FieldData, field.TypeJSON, value)
	}
	_spec.AddModifiers(ctuo.modifiers...)
	_node = &CronTask{config: ctuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ctuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{crontask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ctuo.mutation.done = true
	return _node, nil
}
