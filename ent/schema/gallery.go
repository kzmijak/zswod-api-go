package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Gallery holds the schema definition for the Image entity.
type Gallery struct {
	ent.Schema
}

// Fields of the Gallery.
func (Gallery) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("title"),
		field.Time("createdAt"),
	}
}

// Edges of the Gallery.
func (Gallery) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("images", Image.Type),
		edge.To("article", Article.Type),
	}
}
