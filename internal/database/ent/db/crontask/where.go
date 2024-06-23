// Code generated by ent, DO NOT EDIT.

package crontask

import (
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.CronTask {
	return predicate.CronTask(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldUpdatedAt, v))
}

// ReferenceID applies equality check predicate on the "reference_id" field. It's identical to ReferenceIDEQ.
func ReferenceID(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldReferenceID, v))
}

// ScheduledAfter applies equality check predicate on the "scheduled_after" field. It's identical to ScheduledAfterEQ.
func ScheduledAfter(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldScheduledAfter, v))
}

// LastRun applies equality check predicate on the "last_run" field. It's identical to LastRunEQ.
func LastRun(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldLastRun, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v models.TaskType) predicate.CronTask {
	vc := v
	return predicate.CronTask(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v models.TaskType) predicate.CronTask {
	vc := v
	return predicate.CronTask(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...models.TaskType) predicate.CronTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronTask(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...models.TaskType) predicate.CronTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronTask(sql.FieldNotIn(FieldType, v...))
}

// ReferenceIDEQ applies the EQ predicate on the "reference_id" field.
func ReferenceIDEQ(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldReferenceID, v))
}

// ReferenceIDNEQ applies the NEQ predicate on the "reference_id" field.
func ReferenceIDNEQ(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldReferenceID, v))
}

// ReferenceIDIn applies the In predicate on the "reference_id" field.
func ReferenceIDIn(vs ...string) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldReferenceID, vs...))
}

// ReferenceIDNotIn applies the NotIn predicate on the "reference_id" field.
func ReferenceIDNotIn(vs ...string) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldReferenceID, vs...))
}

// ReferenceIDGT applies the GT predicate on the "reference_id" field.
func ReferenceIDGT(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldReferenceID, v))
}

// ReferenceIDGTE applies the GTE predicate on the "reference_id" field.
func ReferenceIDGTE(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldReferenceID, v))
}

// ReferenceIDLT applies the LT predicate on the "reference_id" field.
func ReferenceIDLT(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldReferenceID, v))
}

// ReferenceIDLTE applies the LTE predicate on the "reference_id" field.
func ReferenceIDLTE(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldReferenceID, v))
}

// ReferenceIDContains applies the Contains predicate on the "reference_id" field.
func ReferenceIDContains(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldContains(FieldReferenceID, v))
}

// ReferenceIDHasPrefix applies the HasPrefix predicate on the "reference_id" field.
func ReferenceIDHasPrefix(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldHasPrefix(FieldReferenceID, v))
}

// ReferenceIDHasSuffix applies the HasSuffix predicate on the "reference_id" field.
func ReferenceIDHasSuffix(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldHasSuffix(FieldReferenceID, v))
}

// ReferenceIDEqualFold applies the EqualFold predicate on the "reference_id" field.
func ReferenceIDEqualFold(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEqualFold(FieldReferenceID, v))
}

// ReferenceIDContainsFold applies the ContainsFold predicate on the "reference_id" field.
func ReferenceIDContainsFold(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldContainsFold(FieldReferenceID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v models.TaskStatus) predicate.CronTask {
	vc := v
	return predicate.CronTask(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v models.TaskStatus) predicate.CronTask {
	vc := v
	return predicate.CronTask(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...models.TaskStatus) predicate.CronTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronTask(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...models.TaskStatus) predicate.CronTask {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CronTask(sql.FieldNotIn(FieldStatus, v...))
}

// ScheduledAfterEQ applies the EQ predicate on the "scheduled_after" field.
func ScheduledAfterEQ(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldScheduledAfter, v))
}

// ScheduledAfterNEQ applies the NEQ predicate on the "scheduled_after" field.
func ScheduledAfterNEQ(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldScheduledAfter, v))
}

// ScheduledAfterIn applies the In predicate on the "scheduled_after" field.
func ScheduledAfterIn(vs ...int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldScheduledAfter, vs...))
}

// ScheduledAfterNotIn applies the NotIn predicate on the "scheduled_after" field.
func ScheduledAfterNotIn(vs ...int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldScheduledAfter, vs...))
}

// ScheduledAfterGT applies the GT predicate on the "scheduled_after" field.
func ScheduledAfterGT(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldScheduledAfter, v))
}

// ScheduledAfterGTE applies the GTE predicate on the "scheduled_after" field.
func ScheduledAfterGTE(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldScheduledAfter, v))
}

// ScheduledAfterLT applies the LT predicate on the "scheduled_after" field.
func ScheduledAfterLT(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldScheduledAfter, v))
}

// ScheduledAfterLTE applies the LTE predicate on the "scheduled_after" field.
func ScheduledAfterLTE(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldScheduledAfter, v))
}

// LastRunEQ applies the EQ predicate on the "last_run" field.
func LastRunEQ(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldLastRun, v))
}

// LastRunNEQ applies the NEQ predicate on the "last_run" field.
func LastRunNEQ(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldLastRun, v))
}

// LastRunIn applies the In predicate on the "last_run" field.
func LastRunIn(vs ...int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldLastRun, vs...))
}

// LastRunNotIn applies the NotIn predicate on the "last_run" field.
func LastRunNotIn(vs ...int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldLastRun, vs...))
}

// LastRunGT applies the GT predicate on the "last_run" field.
func LastRunGT(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldLastRun, v))
}

// LastRunGTE applies the GTE predicate on the "last_run" field.
func LastRunGTE(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldLastRun, v))
}

// LastRunLT applies the LT predicate on the "last_run" field.
func LastRunLT(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldLastRun, v))
}

// LastRunLTE applies the LTE predicate on the "last_run" field.
func LastRunLTE(v int64) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldLastRun, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CronTask) predicate.CronTask {
	return predicate.CronTask(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CronTask) predicate.CronTask {
	return predicate.CronTask(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CronTask) predicate.CronTask {
	return predicate.CronTask(sql.NotPredicates(p))
}