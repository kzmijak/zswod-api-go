package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// CustomPage holds the schema definition for the CustomPage entity.
type CustomPage struct {
	ent.Schema
}

// Mixin of the CustomPage schema.
func (CustomPage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		IconMixin{},
		mixin.UpdateTime{},
	}
}

// Fields of the CustomPage.
func (CustomPage) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MinLen(6).MaxLen(200),
		field.String("titleNormalized").Unique(),
		field.String("content").
			SchemaType(map[string]string{
			dialect.MySQL: "mediumtext",
		}),
	}
}

// Edges of the CustomPage.
func (CustomPage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachments", Attachment.Type),
	}
}
