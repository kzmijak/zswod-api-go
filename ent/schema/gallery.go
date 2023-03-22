package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Gallery holds the schema definition for the Image entity.
type Gallery struct {
	ent.Schema
}

// Mixin of the Gallery schema.
func (Gallery) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		mixin.Time{},
	}
}

// Fields of the Gallery.
func (Gallery) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
	}
}

// Edges of the Gallery.
func (Gallery) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("images", Image.Type),
		edge.From("article", Article.Type).Ref("gallery").Unique(),
		edge.From("author", User.Type).Ref("galleries").Required().Unique(),
	}
}
