package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("title").MinLen(6).MaxLen(32),
		field.String("short").MinLen(12).MaxLen(128),
		field.String("content"),
		field.Time("uploadDate"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Image.Type),
		edge.To("titleNormalized", ArticleTitleGuid.Type).Unique(),
	}
}
