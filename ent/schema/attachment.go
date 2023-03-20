package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Attachment holds the schema definition for the Attachment entity.
type Attachment struct {
	ent.Schema
}

// Mixin of the Attachment schema.
func (Attachment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		IconMixin{},
	}
}

// Fields of the Attachment.
func (Attachment) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("description").Optional(),
	}
}

// Edges of the Attachment.
func (Attachment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blob", Blob.Type).
			Required().
			Immutable().
			Unique(),
	}
}
