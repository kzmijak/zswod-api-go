package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the Users.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("password"),
		field.String("email"),
	}
}

// Edges of the Users.
func (User) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("roles", Role.Type).Ref("users").Unique(),
		edge.To("resetPasswordTokens", ResetPasswordToken.Type),
	}
}
