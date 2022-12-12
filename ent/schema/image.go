package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("image_guid", uuid.New()),
		field.Bytes("blob"),
		field.String("content_type"),
		field.String("title"),
		field.String("alt"),
		field.Time("upload_date"),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("article", Article.Type).
			Ref("images").Unique(),
	}
}