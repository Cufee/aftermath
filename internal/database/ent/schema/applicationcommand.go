package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// model ApplicationCommand {
//   id        String   @id
//   createdAt DateTime @default(now())
//   updatedAt DateTime @updatedAt

//   name        String
//   version     String
//   optionsHash String

//   @@index([optionsHash])
//   @@map("application_commands")
// }

// ApplicationCommand holds the schema definition for the ApplicationCommand entity.
type ApplicationCommand struct {
	ent.Schema
}

// Fields of the ApplicationCommand.
func (ApplicationCommand) Fields() []ent.Field {
	return append(defaultFields,
		field.String("name").Unique().NotEmpty(),
		field.String("version").NotEmpty(),
		field.String("options_hash").NotEmpty(),
	)
}

// Edges of the ApplicationCommand.
func (ApplicationCommand) Edges() []ent.Edge {
	return nil
}

func (ApplicationCommand) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("options_hash"),
	}
}
