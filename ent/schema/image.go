package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Mixin of the Image schema.
func (Image) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		mixin.CreateTime{},
	}
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("alt").Optional(),
		field.Int("Order").Optional().Min(0),
		field.String("src").Immutable(),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("gallery", Gallery.Type).Ref("images").
			Unique(),
	}
}
