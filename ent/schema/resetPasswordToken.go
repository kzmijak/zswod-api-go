package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// ResetPasswordToken holds the schema definition for the ResetPasswordToken entity.
type ResetPasswordToken struct {
	ent.Schema
}

func (ResetPasswordToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		mixin.Time{},
	}
}

// Fields of the ResetPasswordToken.
func (ResetPasswordToken) Fields() []ent.Field {
	return []ent.Field{
		field.Time("expiryTime"),
	}
}

// Edges of the ResetPasswordToken.
func (ResetPasswordToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("resetPasswordToken").
			Unique().
			Required(),
	}
}
