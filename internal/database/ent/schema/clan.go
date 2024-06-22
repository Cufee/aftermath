package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// model Clan {
//   id        String   @id
//   createdAt DateTime
//   updatedAt DateTime @updatedAt

//   tag      String
//   name     String
//   emblemId String @default("")

//   accounts      Account[]
//   membersString String

//   recordUpdatedAt DateTime @updatedAt

//   @@index([tag])
//   @@map("account_clans")
// }

// Clan holds the schema definition for the Clan entity.
type Clan struct {
	ent.Schema
}

// Fields of the Clan.
func (Clan) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable(),
		field.Int("created_at").
			Immutable().
			DefaultFunc(timeNow),
		field.Int("updated_at").
			DefaultFunc(timeNow).
			UpdateDefault(timeNow),
		//
		field.String("tag").NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("emblem_id").Default("").Optional(),
		field.Strings("members"),
	}
}

// Edges of the Clan.
func (Clan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accounts", Account.Type),
	}
}

func (Clan) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tag"),
		index.Fields("name"),
	}
}
