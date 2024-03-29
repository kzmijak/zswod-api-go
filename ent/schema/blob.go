package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Blob holds the schema definition for the Image entity.
type Blob struct {
	ent.Schema
}

// Mixin of the Blob schema.
func (Blob) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		mixin.CreateTime{},
	}
}

// Fields of the Image.
func (Blob) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("blob").
			SchemaType(map[string]string{
			dialect.MySQL: "mediumblob",
		}),
		field.String("title"),
		field.String("contentType"),
		field.Enum("type").Values("Picture", "Attachment").Default("Picture"),
		field.Bool("isPublic").Default(true),
	}
}

// Edges of the Image.
func (Blob) Edges() []ent.Edge {
	return []ent.Edge {}
}
