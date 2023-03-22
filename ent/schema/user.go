package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User schema.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
	}
}

// Fields of the Users.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("password"),
		field.String("email"),
		field.String("firstName").Optional(),
		field.String("lastName").Optional(),
		field.Enum("role").Values("Admin", "Teacher", "LegalGuardian", "Student", "Unknown"),
	}
}

// Edges of the Users.
func (User) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("galleries", Gallery.Type),
		edge.To("articles", Article.Type).Unique(),
		edge.To("avatar", Image.Type).Unique(),
		edge.To("resetPasswordTokens", ResetPasswordToken.Type),
	}
}
