// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedAt, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldExpiresAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUserID, v))
}

// PublicID applies equality check predicate on the "public_id" field. It's identical to PublicIDEQ.
func PublicID(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldPublicID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldUpdatedAt, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldExpiresAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Session {
	return predicate.Session(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldUserID, v))
}

// PublicIDEQ applies the EQ predicate on the "public_id" field.
func PublicIDEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldPublicID, v))
}

// PublicIDNEQ applies the NEQ predicate on the "public_id" field.
func PublicIDNEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldPublicID, v))
}

// PublicIDIn applies the In predicate on the "public_id" field.
func PublicIDIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldPublicID, vs...))
}

// PublicIDNotIn applies the NotIn predicate on the "public_id" field.
func PublicIDNotIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldPublicID, vs...))
}

// PublicIDGT applies the GT predicate on the "public_id" field.
func PublicIDGT(v string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldPublicID, v))
}

// PublicIDGTE applies the GTE predicate on the "public_id" field.
func PublicIDGTE(v string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldPublicID, v))
}

// PublicIDLT applies the LT predicate on the "public_id" field.
func PublicIDLT(v string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldPublicID, v))
}

// PublicIDLTE applies the LTE predicate on the "public_id" field.
func PublicIDLTE(v string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldPublicID, v))
}

// PublicIDContains applies the Contains predicate on the "public_id" field.
func PublicIDContains(v string) predicate.Session {
	return predicate.Session(sql.FieldContains(FieldPublicID, v))
}

// PublicIDHasPrefix applies the HasPrefix predicate on the "public_id" field.
func PublicIDHasPrefix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasPrefix(FieldPublicID, v))
}

// PublicIDHasSuffix applies the HasSuffix predicate on the "public_id" field.
func PublicIDHasSuffix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasSuffix(FieldPublicID, v))
}

// PublicIDEqualFold applies the EqualFold predicate on the "public_id" field.
func PublicIDEqualFold(v string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldPublicID, v))
}

// PublicIDContainsFold applies the ContainsFold predicate on the "public_id" field.
func PublicIDContainsFold(v string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldPublicID, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(sql.NotPredicates(p))
}
