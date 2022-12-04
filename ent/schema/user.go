package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the Users.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("username"),
		field.String("password"),
		field.String("email"),
	}
}

// Edges of the Users.
func (User) Edges() []ent.Edge {
	return nil
}
