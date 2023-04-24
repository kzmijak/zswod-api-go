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
		mixin.Time{},
	}
}

// Fields of the CustomPage.
func (CustomPage) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MinLen(3).MaxLen(36),
		field.String("url").Unique(),
		field.String("content").
			SchemaType(map[string]string{
			dialect.MySQL: "mediumtext",
		}),
		field.Bool("isExternal").Optional(),
		field.String("link").Optional(),
		field.String("section").MinLen(3).MaxLen(36),
	}
}

// Edges of the CustomPage.
func (CustomPage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attachments", Attachment.Type),
	}
}
