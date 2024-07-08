// Code generated by ent, DO NOT EDIT.

package leaderboardscore

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldUpdatedAt, v))
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldScore, v))
}

// ReferenceID applies equality check predicate on the "reference_id" field. It's identical to ReferenceIDEQ.
func ReferenceID(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldReferenceID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v models.ScoreType) predicate.LeaderboardScore {
	vc := v
	return predicate.LeaderboardScore(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v models.ScoreType) predicate.LeaderboardScore {
	vc := v
	return predicate.LeaderboardScore(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...models.ScoreType) predicate.LeaderboardScore {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LeaderboardScore(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...models.ScoreType) predicate.LeaderboardScore {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LeaderboardScore(sql.FieldNotIn(FieldType, v...))
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldScore, v))
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNEQ(FieldScore, v))
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldIn(FieldScore, vs...))
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNotIn(FieldScore, vs...))
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGT(FieldScore, v))
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGTE(FieldScore, v))
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLT(FieldScore, v))
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v float32) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLTE(FieldScore, v))
}

// ReferenceIDEQ applies the EQ predicate on the "reference_id" field.
func ReferenceIDEQ(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEQ(FieldReferenceID, v))
}

// ReferenceIDNEQ applies the NEQ predicate on the "reference_id" field.
func ReferenceIDNEQ(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNEQ(FieldReferenceID, v))
}

// ReferenceIDIn applies the In predicate on the "reference_id" field.
func ReferenceIDIn(vs ...string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldIn(FieldReferenceID, vs...))
}

// ReferenceIDNotIn applies the NotIn predicate on the "reference_id" field.
func ReferenceIDNotIn(vs ...string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldNotIn(FieldReferenceID, vs...))
}

// ReferenceIDGT applies the GT predicate on the "reference_id" field.
func ReferenceIDGT(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGT(FieldReferenceID, v))
}

// ReferenceIDGTE applies the GTE predicate on the "reference_id" field.
func ReferenceIDGTE(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldGTE(FieldReferenceID, v))
}

// ReferenceIDLT applies the LT predicate on the "reference_id" field.
func ReferenceIDLT(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLT(FieldReferenceID, v))
}

// ReferenceIDLTE applies the LTE predicate on the "reference_id" field.
func ReferenceIDLTE(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldLTE(FieldReferenceID, v))
}

// ReferenceIDContains applies the Contains predicate on the "reference_id" field.
func ReferenceIDContains(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldContains(FieldReferenceID, v))
}

// ReferenceIDHasPrefix applies the HasPrefix predicate on the "reference_id" field.
func ReferenceIDHasPrefix(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldHasPrefix(FieldReferenceID, v))
}

// ReferenceIDHasSuffix applies the HasSuffix predicate on the "reference_id" field.
func ReferenceIDHasSuffix(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldHasSuffix(FieldReferenceID, v))
}

// ReferenceIDEqualFold applies the EqualFold predicate on the "reference_id" field.
func ReferenceIDEqualFold(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldEqualFold(FieldReferenceID, v))
}

// ReferenceIDContainsFold applies the ContainsFold predicate on the "reference_id" field.
func ReferenceIDContainsFold(v string) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.FieldContainsFold(FieldReferenceID, v))
}

// LeaderboardIDEQ applies the EQ predicate on the "leaderboard_id" field.
func LeaderboardIDEQ(v models.LeaderboardID) predicate.LeaderboardScore {
	vc := v
	return predicate.LeaderboardScore(sql.FieldEQ(FieldLeaderboardID, vc))
}

// LeaderboardIDNEQ applies the NEQ predicate on the "leaderboard_id" field.
func LeaderboardIDNEQ(v models.LeaderboardID) predicate.LeaderboardScore {
	vc := v
	return predicate.LeaderboardScore(sql.FieldNEQ(FieldLeaderboardID, vc))
}

// LeaderboardIDIn applies the In predicate on the "leaderboard_id" field.
func LeaderboardIDIn(vs ...models.LeaderboardID) predicate.LeaderboardScore {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LeaderboardScore(sql.FieldIn(FieldLeaderboardID, v...))
}

// LeaderboardIDNotIn applies the NotIn predicate on the "leaderboard_id" field.
func LeaderboardIDNotIn(vs ...models.LeaderboardID) predicate.LeaderboardScore {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LeaderboardScore(sql.FieldNotIn(FieldLeaderboardID, v...))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LeaderboardScore) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LeaderboardScore) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LeaderboardScore) predicate.LeaderboardScore {
	return predicate.LeaderboardScore(sql.NotPredicates(p))
}