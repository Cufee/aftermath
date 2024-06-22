// Code generated by ent, DO NOT EDIT.

package crontask

import (
	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
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
func CreatedAt(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldUpdatedAt, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldType, v))
}

// ReferenceID applies equality check predicate on the "reference_id" field. It's identical to ReferenceIDEQ.
func ReferenceID(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldReferenceID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldStatus, v))
}

// ScheduledAfter applies equality check predicate on the "scheduled_after" field. It's identical to ScheduledAfterEQ.
func ScheduledAfter(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldScheduledAfter, v))
}

// LastRun applies equality check predicate on the "last_run" field. It's identical to LastRunEQ.
func LastRun(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldLastRun, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldContainsFold(FieldType, v))
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
func StatusEQ(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.CronTask {
	return predicate.CronTask(sql.FieldContainsFold(FieldStatus, v))
}

// ScheduledAfterEQ applies the EQ predicate on the "scheduled_after" field.
func ScheduledAfterEQ(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldScheduledAfter, v))
}

// ScheduledAfterNEQ applies the NEQ predicate on the "scheduled_after" field.
func ScheduledAfterNEQ(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldScheduledAfter, v))
}

// ScheduledAfterIn applies the In predicate on the "scheduled_after" field.
func ScheduledAfterIn(vs ...int) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldScheduledAfter, vs...))
}

// ScheduledAfterNotIn applies the NotIn predicate on the "scheduled_after" field.
func ScheduledAfterNotIn(vs ...int) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldScheduledAfter, vs...))
}

// ScheduledAfterGT applies the GT predicate on the "scheduled_after" field.
func ScheduledAfterGT(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldScheduledAfter, v))
}

// ScheduledAfterGTE applies the GTE predicate on the "scheduled_after" field.
func ScheduledAfterGTE(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldScheduledAfter, v))
}

// ScheduledAfterLT applies the LT predicate on the "scheduled_after" field.
func ScheduledAfterLT(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldScheduledAfter, v))
}

// ScheduledAfterLTE applies the LTE predicate on the "scheduled_after" field.
func ScheduledAfterLTE(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldLTE(FieldScheduledAfter, v))
}

// LastRunEQ applies the EQ predicate on the "last_run" field.
func LastRunEQ(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldEQ(FieldLastRun, v))
}

// LastRunNEQ applies the NEQ predicate on the "last_run" field.
func LastRunNEQ(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldNEQ(FieldLastRun, v))
}

// LastRunIn applies the In predicate on the "last_run" field.
func LastRunIn(vs ...int) predicate.CronTask {
	return predicate.CronTask(sql.FieldIn(FieldLastRun, vs...))
}

// LastRunNotIn applies the NotIn predicate on the "last_run" field.
func LastRunNotIn(vs ...int) predicate.CronTask {
	return predicate.CronTask(sql.FieldNotIn(FieldLastRun, vs...))
}

// LastRunGT applies the GT predicate on the "last_run" field.
func LastRunGT(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldGT(FieldLastRun, v))
}

// LastRunGTE applies the GTE predicate on the "last_run" field.
func LastRunGTE(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldGTE(FieldLastRun, v))
}

// LastRunLT applies the LT predicate on the "last_run" field.
func LastRunLT(v int) predicate.CronTask {
	return predicate.CronTask(sql.FieldLT(FieldLastRun, v))
}

// LastRunLTE applies the LTE predicate on the "last_run" field.
func LastRunLTE(v int) predicate.CronTask {
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
