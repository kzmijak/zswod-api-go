package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Mixin of the Article schema.
func (Article) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		mixin.Time{},
	}
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MinLen(6).MaxLen(200),
		field.String("titleNormalized").Unique(),
		field.String("short").MinLen(12).MaxLen(300),
		field.String("content").
			SchemaType(map[string]string{
			dialect.MySQL: "mediumtext",
		}),
		field.Enum("status").Values("Draft", "Review", "Published", "Removed", "Unknown"),

	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("gallery", Gallery.Type).Required().Unique(),
		edge.From("author", User.Type).Ref("articles").Required().Unique(),
		edge.To("attachments", Attachment.Type),
	}
}
