package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

// model UserConnection {
//   id        String   @id @default(cuid())
//   createdAt DateTime @default(now())
//   updatedAt DateTime @updatedAt

//   user   User   @relation(fields: [userId], references: [id])
//   userId String

//   type        String
//   permissions String
//   referenceId String

//   metadataEncoded Bytes

//   @@index([userId])
//   @@index([type, userId])
//   @@index([referenceId])
//   @@index([type, referenceId])
//   @@map("user_connections")
// }

// UserConnection holds the schema definition for the UserConnection entity.
type UserConnection struct {
	ent.Schema
}

// Fields of the UserConnection.
func (UserConnection) Fields() []ent.Field {
	return append(defaultFields,
		field.Enum("type").
			GoType(models.ConnectionType("")),
		field.String("user_id").Immutable(),
		field.String("reference_id"),
		field.String("permissions").Default("").Optional(),
		field.JSON("metadata", map[string]any{}).Optional(),
	)
}

// Edges of the UserConnection.
func (UserConnection) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("connections").Field("user_id").Required().Immutable().Unique(),
	}
}

func (UserConnection) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("type", "user_id"),
		index.Fields("reference_id"),
		index.Fields("type", "reference_id"),
		// index.Fields("reference_id").Edges("user").Unique(),
	}
}
