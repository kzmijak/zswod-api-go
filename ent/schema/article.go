package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.String("title").MinLen(6).MaxLen(200),
		field.String("titleNormalized").Unique(),
		field.String("short").MinLen(12).MaxLen(300),
		field.String("content").
			SchemaType(map[string]string{
			dialect.MySQL: "mediumtext",
		}),
		field.Time("uploadDate"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("gallery", Gallery.Type).Ref("article").Unique(),
	}
}
