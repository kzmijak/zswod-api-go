package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ResetPasswordToken holds the schema definition for the ResetPasswordToken entity.
type ResetPasswordToken struct {
	ent.Schema
}

// Fields of the ResetPasswordToken.
func (ResetPasswordToken) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.Time("createdAt"),
	}
}

// Edges of the ResetPasswordToken.
func (ResetPasswordToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("resetPasswordTokens").
			Unique().
			Required(),
	}
}
