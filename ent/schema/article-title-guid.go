package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ArticleTitleGuid holds the schema definition for the ArticleTitleGuid entity.
type ArticleTitleGuid struct {
	ent.Schema
}

// Fields of the ArticleTitleGuid.
func (ArticleTitleGuid) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Unique(),
		field.String("titleNormalized").Unique(),
	}
}

// Edges of the ArticleTitleGuid.
func (ArticleTitleGuid) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).Ref("titleNormalized").Unique().Required(),
	}
}
