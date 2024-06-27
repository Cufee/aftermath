package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{field.String("id").
		Unique().
		Immutable(),
		field.Time("created_at").
			Immutable().
			Default(timeNow),
		field.Time("updated_at").
			Default(timeNow).
			UpdateDefault(timeNow),
		field.Time("last_battle_time"),
		field.Time("account_created_at"),
		//
		field.String("realm").
			MinLen(2).
			MaxLen(5),
		field.String("nickname").NotEmpty(),
		field.Bool("private").
			Default(false),
		//
		field.String("clan_id").Optional(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("clan", Clan.Type).Ref("accounts").Field("clan_id").Unique(),
		edge.To("snapshots", AccountSnapshot.Type),
		edge.To("vehicle_snapshots", VehicleSnapshot.Type),
		edge.To("achievement_snapshots", AchievementsSnapshot.Type),
	}
}

func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("id", "last_battle_time"),
		index.Fields("realm"),
		index.Fields("realm", "last_battle_time"),
		index.Fields("clan_id"),
	}
}
