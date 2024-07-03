// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldUpdatedAt, v))
}

// LastBattleTime applies equality check predicate on the "last_battle_time" field. It's identical to LastBattleTimeEQ.
func LastBattleTime(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldLastBattleTime, v))
}

// AccountCreatedAt applies equality check predicate on the "account_created_at" field. It's identical to AccountCreatedAtEQ.
func AccountCreatedAt(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccountCreatedAt, v))
}

// Realm applies equality check predicate on the "realm" field. It's identical to RealmEQ.
func Realm(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldRealm, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldNickname, v))
}

// Private applies equality check predicate on the "private" field. It's identical to PrivateEQ.
func Private(v bool) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPrivate, v))
}

// ClanID applies equality check predicate on the "clan_id" field. It's identical to ClanIDEQ.
func ClanID(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldClanID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldUpdatedAt, v))
}

// LastBattleTimeEQ applies the EQ predicate on the "last_battle_time" field.
func LastBattleTimeEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldLastBattleTime, v))
}

// LastBattleTimeNEQ applies the NEQ predicate on the "last_battle_time" field.
func LastBattleTimeNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldLastBattleTime, v))
}

// LastBattleTimeIn applies the In predicate on the "last_battle_time" field.
func LastBattleTimeIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldLastBattleTime, vs...))
}

// LastBattleTimeNotIn applies the NotIn predicate on the "last_battle_time" field.
func LastBattleTimeNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldLastBattleTime, vs...))
}

// LastBattleTimeGT applies the GT predicate on the "last_battle_time" field.
func LastBattleTimeGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldLastBattleTime, v))
}

// LastBattleTimeGTE applies the GTE predicate on the "last_battle_time" field.
func LastBattleTimeGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldLastBattleTime, v))
}

// LastBattleTimeLT applies the LT predicate on the "last_battle_time" field.
func LastBattleTimeLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldLastBattleTime, v))
}

// LastBattleTimeLTE applies the LTE predicate on the "last_battle_time" field.
func LastBattleTimeLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldLastBattleTime, v))
}

// AccountCreatedAtEQ applies the EQ predicate on the "account_created_at" field.
func AccountCreatedAtEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldAccountCreatedAt, v))
}

// AccountCreatedAtNEQ applies the NEQ predicate on the "account_created_at" field.
func AccountCreatedAtNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldAccountCreatedAt, v))
}

// AccountCreatedAtIn applies the In predicate on the "account_created_at" field.
func AccountCreatedAtIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldAccountCreatedAt, vs...))
}

// AccountCreatedAtNotIn applies the NotIn predicate on the "account_created_at" field.
func AccountCreatedAtNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldAccountCreatedAt, vs...))
}

// AccountCreatedAtGT applies the GT predicate on the "account_created_at" field.
func AccountCreatedAtGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldAccountCreatedAt, v))
}

// AccountCreatedAtGTE applies the GTE predicate on the "account_created_at" field.
func AccountCreatedAtGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldAccountCreatedAt, v))
}

// AccountCreatedAtLT applies the LT predicate on the "account_created_at" field.
func AccountCreatedAtLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldAccountCreatedAt, v))
}

// AccountCreatedAtLTE applies the LTE predicate on the "account_created_at" field.
func AccountCreatedAtLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldAccountCreatedAt, v))
}

// RealmEQ applies the EQ predicate on the "realm" field.
func RealmEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldRealm, v))
}

// RealmNEQ applies the NEQ predicate on the "realm" field.
func RealmNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldRealm, v))
}

// RealmIn applies the In predicate on the "realm" field.
func RealmIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldRealm, vs...))
}

// RealmNotIn applies the NotIn predicate on the "realm" field.
func RealmNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldRealm, vs...))
}

// RealmGT applies the GT predicate on the "realm" field.
func RealmGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldRealm, v))
}

// RealmGTE applies the GTE predicate on the "realm" field.
func RealmGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldRealm, v))
}

// RealmLT applies the LT predicate on the "realm" field.
func RealmLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldRealm, v))
}

// RealmLTE applies the LTE predicate on the "realm" field.
func RealmLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldRealm, v))
}

// RealmContains applies the Contains predicate on the "realm" field.
func RealmContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldRealm, v))
}

// RealmHasPrefix applies the HasPrefix predicate on the "realm" field.
func RealmHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldRealm, v))
}

// RealmHasSuffix applies the HasSuffix predicate on the "realm" field.
func RealmHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldRealm, v))
}

// RealmEqualFold applies the EqualFold predicate on the "realm" field.
func RealmEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldRealm, v))
}

// RealmContainsFold applies the ContainsFold predicate on the "realm" field.
func RealmContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldRealm, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldNickname, v))
}

// PrivateEQ applies the EQ predicate on the "private" field.
func PrivateEQ(v bool) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPrivate, v))
}

// PrivateNEQ applies the NEQ predicate on the "private" field.
func PrivateNEQ(v bool) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldPrivate, v))
}

// ClanIDEQ applies the EQ predicate on the "clan_id" field.
func ClanIDEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldClanID, v))
}

// ClanIDNEQ applies the NEQ predicate on the "clan_id" field.
func ClanIDNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldClanID, v))
}

// ClanIDIn applies the In predicate on the "clan_id" field.
func ClanIDIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldClanID, vs...))
}

// ClanIDNotIn applies the NotIn predicate on the "clan_id" field.
func ClanIDNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldClanID, vs...))
}

// ClanIDGT applies the GT predicate on the "clan_id" field.
func ClanIDGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldClanID, v))
}

// ClanIDGTE applies the GTE predicate on the "clan_id" field.
func ClanIDGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldClanID, v))
}

// ClanIDLT applies the LT predicate on the "clan_id" field.
func ClanIDLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldClanID, v))
}

// ClanIDLTE applies the LTE predicate on the "clan_id" field.
func ClanIDLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldClanID, v))
}

// ClanIDContains applies the Contains predicate on the "clan_id" field.
func ClanIDContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldClanID, v))
}

// ClanIDHasPrefix applies the HasPrefix predicate on the "clan_id" field.
func ClanIDHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldClanID, v))
}

// ClanIDHasSuffix applies the HasSuffix predicate on the "clan_id" field.
func ClanIDHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldClanID, v))
}

// ClanIDIsNil applies the IsNil predicate on the "clan_id" field.
func ClanIDIsNil() predicate.Account {
	return predicate.Account(sql.FieldIsNull(FieldClanID))
}

// ClanIDNotNil applies the NotNil predicate on the "clan_id" field.
func ClanIDNotNil() predicate.Account {
	return predicate.Account(sql.FieldNotNull(FieldClanID))
}

// ClanIDEqualFold applies the EqualFold predicate on the "clan_id" field.
func ClanIDEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldClanID, v))
}

// ClanIDContainsFold applies the ContainsFold predicate on the "clan_id" field.
func ClanIDContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldClanID, v))
}

// HasClan applies the HasEdge predicate on the "clan" edge.
func HasClan() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ClanTable, ClanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasClanWith applies the HasEdge predicate on the "clan" edge with a given conditions (other predicates).
func HasClanWith(preds ...predicate.Clan) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newClanStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAchievementSnapshots applies the HasEdge predicate on the "achievement_snapshots" edge.
func HasAchievementSnapshots() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AchievementSnapshotsTable, AchievementSnapshotsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAchievementSnapshotsWith applies the HasEdge predicate on the "achievement_snapshots" edge with a given conditions (other predicates).
func HasAchievementSnapshotsWith(preds ...predicate.AchievementsSnapshot) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newAchievementSnapshotsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVehicleSnapshots applies the HasEdge predicate on the "vehicle_snapshots" edge.
func HasVehicleSnapshots() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VehicleSnapshotsTable, VehicleSnapshotsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVehicleSnapshotsWith applies the HasEdge predicate on the "vehicle_snapshots" edge with a given conditions (other predicates).
func HasVehicleSnapshotsWith(preds ...predicate.VehicleSnapshot) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newVehicleSnapshotsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccountSnapshots applies the HasEdge predicate on the "account_snapshots" edge.
func HasAccountSnapshots() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountSnapshotsTable, AccountSnapshotsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountSnapshotsWith applies the HasEdge predicate on the "account_snapshots" edge with a given conditions (other predicates).
func HasAccountSnapshotsWith(preds ...predicate.AccountSnapshot) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newAccountSnapshotsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(sql.NotPredicates(p))
}
