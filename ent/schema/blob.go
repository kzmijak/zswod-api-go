package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Blob holds the schema definition for the Image entity.
type Blob struct {
	ent.Schema
}

// Fields of the Image.
func (Blob) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.Bytes("blob").
			SchemaType(map[string]string{
			dialect.MySQL: "mediumblob",
		}),
		field.String("content_type"),
	}
}

// Edges of the Image.
func (Blob) Edges() []ent.Edge {
	return []ent.Edge {}
}
