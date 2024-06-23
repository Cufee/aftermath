// Code generated by ent, DO NOT EDIT.

package accountsnapshot

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldUpdatedAt, v))
}

// LastBattleTime applies equality check predicate on the "last_battle_time" field. It's identical to LastBattleTimeEQ.
func LastBattleTime(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldLastBattleTime, v))
}

// AccountID applies equality check predicate on the "account_id" field. It's identical to AccountIDEQ.
func AccountID(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldAccountID, v))
}

// ReferenceID applies equality check predicate on the "reference_id" field. It's identical to ReferenceIDEQ.
func ReferenceID(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldReferenceID, v))
}

// RatingBattles applies equality check predicate on the "rating_battles" field. It's identical to RatingBattlesEQ.
func RatingBattles(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldRatingBattles, v))
}

// RegularBattles applies equality check predicate on the "regular_battles" field. It's identical to RegularBattlesEQ.
func RegularBattles(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldRegularBattles, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v models.SnapshotType) predicate.AccountSnapshot {
	vc := v
	return predicate.AccountSnapshot(sql.FieldEQ(FieldType, vc))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v models.SnapshotType) predicate.AccountSnapshot {
	vc := v
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldType, vc))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...models.SnapshotType) predicate.AccountSnapshot {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountSnapshot(sql.FieldIn(FieldType, v...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...models.SnapshotType) predicate.AccountSnapshot {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldType, v...))
}

// LastBattleTimeEQ applies the EQ predicate on the "last_battle_time" field.
func LastBattleTimeEQ(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldLastBattleTime, v))
}

// LastBattleTimeNEQ applies the NEQ predicate on the "last_battle_time" field.
func LastBattleTimeNEQ(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldLastBattleTime, v))
}

// LastBattleTimeIn applies the In predicate on the "last_battle_time" field.
func LastBattleTimeIn(vs ...int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldIn(FieldLastBattleTime, vs...))
}

// LastBattleTimeNotIn applies the NotIn predicate on the "last_battle_time" field.
func LastBattleTimeNotIn(vs ...int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldLastBattleTime, vs...))
}

// LastBattleTimeGT applies the GT predicate on the "last_battle_time" field.
func LastBattleTimeGT(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGT(FieldLastBattleTime, v))
}

// LastBattleTimeGTE applies the GTE predicate on the "last_battle_time" field.
func LastBattleTimeGTE(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGTE(FieldLastBattleTime, v))
}

// LastBattleTimeLT applies the LT predicate on the "last_battle_time" field.
func LastBattleTimeLT(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLT(FieldLastBattleTime, v))
}

// LastBattleTimeLTE applies the LTE predicate on the "last_battle_time" field.
func LastBattleTimeLTE(v int64) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLTE(FieldLastBattleTime, v))
}

// AccountIDEQ applies the EQ predicate on the "account_id" field.
func AccountIDEQ(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldAccountID, v))
}

// AccountIDNEQ applies the NEQ predicate on the "account_id" field.
func AccountIDNEQ(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldAccountID, v))
}

// AccountIDIn applies the In predicate on the "account_id" field.
func AccountIDIn(vs ...string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldIn(FieldAccountID, vs...))
}

// AccountIDNotIn applies the NotIn predicate on the "account_id" field.
func AccountIDNotIn(vs ...string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldAccountID, vs...))
}

// AccountIDGT applies the GT predicate on the "account_id" field.
func AccountIDGT(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGT(FieldAccountID, v))
}

// AccountIDGTE applies the GTE predicate on the "account_id" field.
func AccountIDGTE(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGTE(FieldAccountID, v))
}

// AccountIDLT applies the LT predicate on the "account_id" field.
func AccountIDLT(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLT(FieldAccountID, v))
}

// AccountIDLTE applies the LTE predicate on the "account_id" field.
func AccountIDLTE(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLTE(FieldAccountID, v))
}

// AccountIDContains applies the Contains predicate on the "account_id" field.
func AccountIDContains(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldContains(FieldAccountID, v))
}

// AccountIDHasPrefix applies the HasPrefix predicate on the "account_id" field.
func AccountIDHasPrefix(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldHasPrefix(FieldAccountID, v))
}

// AccountIDHasSuffix applies the HasSuffix predicate on the "account_id" field.
func AccountIDHasSuffix(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldHasSuffix(FieldAccountID, v))
}

// AccountIDEqualFold applies the EqualFold predicate on the "account_id" field.
func AccountIDEqualFold(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEqualFold(FieldAccountID, v))
}

// AccountIDContainsFold applies the ContainsFold predicate on the "account_id" field.
func AccountIDContainsFold(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldContainsFold(FieldAccountID, v))
}

// ReferenceIDEQ applies the EQ predicate on the "reference_id" field.
func ReferenceIDEQ(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldReferenceID, v))
}

// ReferenceIDNEQ applies the NEQ predicate on the "reference_id" field.
func ReferenceIDNEQ(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldReferenceID, v))
}

// ReferenceIDIn applies the In predicate on the "reference_id" field.
func ReferenceIDIn(vs ...string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldIn(FieldReferenceID, vs...))
}

// ReferenceIDNotIn applies the NotIn predicate on the "reference_id" field.
func ReferenceIDNotIn(vs ...string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldReferenceID, vs...))
}

// ReferenceIDGT applies the GT predicate on the "reference_id" field.
func ReferenceIDGT(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGT(FieldReferenceID, v))
}

// ReferenceIDGTE applies the GTE predicate on the "reference_id" field.
func ReferenceIDGTE(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGTE(FieldReferenceID, v))
}

// ReferenceIDLT applies the LT predicate on the "reference_id" field.
func ReferenceIDLT(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLT(FieldReferenceID, v))
}

// ReferenceIDLTE applies the LTE predicate on the "reference_id" field.
func ReferenceIDLTE(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLTE(FieldReferenceID, v))
}

// ReferenceIDContains applies the Contains predicate on the "reference_id" field.
func ReferenceIDContains(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldContains(FieldReferenceID, v))
}

// ReferenceIDHasPrefix applies the HasPrefix predicate on the "reference_id" field.
func ReferenceIDHasPrefix(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldHasPrefix(FieldReferenceID, v))
}

// ReferenceIDHasSuffix applies the HasSuffix predicate on the "reference_id" field.
func ReferenceIDHasSuffix(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldHasSuffix(FieldReferenceID, v))
}

// ReferenceIDEqualFold applies the EqualFold predicate on the "reference_id" field.
func ReferenceIDEqualFold(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEqualFold(FieldReferenceID, v))
}

// ReferenceIDContainsFold applies the ContainsFold predicate on the "reference_id" field.
func ReferenceIDContainsFold(v string) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldContainsFold(FieldReferenceID, v))
}

// RatingBattlesEQ applies the EQ predicate on the "rating_battles" field.
func RatingBattlesEQ(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldRatingBattles, v))
}

// RatingBattlesNEQ applies the NEQ predicate on the "rating_battles" field.
func RatingBattlesNEQ(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldRatingBattles, v))
}

// RatingBattlesIn applies the In predicate on the "rating_battles" field.
func RatingBattlesIn(vs ...int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldIn(FieldRatingBattles, vs...))
}

// RatingBattlesNotIn applies the NotIn predicate on the "rating_battles" field.
func RatingBattlesNotIn(vs ...int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldRatingBattles, vs...))
}

// RatingBattlesGT applies the GT predicate on the "rating_battles" field.
func RatingBattlesGT(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGT(FieldRatingBattles, v))
}

// RatingBattlesGTE applies the GTE predicate on the "rating_battles" field.
func RatingBattlesGTE(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGTE(FieldRatingBattles, v))
}

// RatingBattlesLT applies the LT predicate on the "rating_battles" field.
func RatingBattlesLT(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLT(FieldRatingBattles, v))
}

// RatingBattlesLTE applies the LTE predicate on the "rating_battles" field.
func RatingBattlesLTE(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLTE(FieldRatingBattles, v))
}

// RegularBattlesEQ applies the EQ predicate on the "regular_battles" field.
func RegularBattlesEQ(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldEQ(FieldRegularBattles, v))
}

// RegularBattlesNEQ applies the NEQ predicate on the "regular_battles" field.
func RegularBattlesNEQ(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNEQ(FieldRegularBattles, v))
}

// RegularBattlesIn applies the In predicate on the "regular_battles" field.
func RegularBattlesIn(vs ...int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldIn(FieldRegularBattles, vs...))
}

// RegularBattlesNotIn applies the NotIn predicate on the "regular_battles" field.
func RegularBattlesNotIn(vs ...int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldNotIn(FieldRegularBattles, vs...))
}

// RegularBattlesGT applies the GT predicate on the "regular_battles" field.
func RegularBattlesGT(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGT(FieldRegularBattles, v))
}

// RegularBattlesGTE applies the GTE predicate on the "regular_battles" field.
func RegularBattlesGTE(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldGTE(FieldRegularBattles, v))
}

// RegularBattlesLT applies the LT predicate on the "regular_battles" field.
func RegularBattlesLT(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLT(FieldRegularBattles, v))
}

// RegularBattlesLTE applies the LTE predicate on the "regular_battles" field.
func RegularBattlesLTE(v int) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.FieldLTE(FieldRegularBattles, v))
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.AccountSnapshot {
	return predicate.AccountSnapshot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(func(s *sql.Selector) {
		step := newAccountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccountSnapshot) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccountSnapshot) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccountSnapshot) predicate.AccountSnapshot {
	return predicate.AccountSnapshot(sql.NotPredicates(p))
}
