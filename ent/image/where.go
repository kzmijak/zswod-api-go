// Code generated by ent, DO NOT EDIT.

package image

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Alt applies equality check predicate on the "alt" field. It's identical to AltEQ.
func Alt(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAlt), v))
	})
}

// IsPreview applies equality check predicate on the "isPreview" field. It's identical to IsPreviewEQ.
func IsPreview(v bool) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPreview), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// AltEQ applies the EQ predicate on the "alt" field.
func AltEQ(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAlt), v))
	})
}

// AltNEQ applies the NEQ predicate on the "alt" field.
func AltNEQ(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAlt), v))
	})
}

// AltIn applies the In predicate on the "alt" field.
func AltIn(vs ...string) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAlt), v...))
	})
}

// AltNotIn applies the NotIn predicate on the "alt" field.
func AltNotIn(vs ...string) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAlt), v...))
	})
}

// AltGT applies the GT predicate on the "alt" field.
func AltGT(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAlt), v))
	})
}

// AltGTE applies the GTE predicate on the "alt" field.
func AltGTE(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAlt), v))
	})
}

// AltLT applies the LT predicate on the "alt" field.
func AltLT(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAlt), v))
	})
}

// AltLTE applies the LTE predicate on the "alt" field.
func AltLTE(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAlt), v))
	})
}

// AltContains applies the Contains predicate on the "alt" field.
func AltContains(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAlt), v))
	})
}

// AltHasPrefix applies the HasPrefix predicate on the "alt" field.
func AltHasPrefix(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAlt), v))
	})
}

// AltHasSuffix applies the HasSuffix predicate on the "alt" field.
func AltHasSuffix(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAlt), v))
	})
}

// AltEqualFold applies the EqualFold predicate on the "alt" field.
func AltEqualFold(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAlt), v))
	})
}

// AltContainsFold applies the ContainsFold predicate on the "alt" field.
func AltContainsFold(v string) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAlt), v))
	})
}

// IsPreviewEQ applies the EQ predicate on the "isPreview" field.
func IsPreviewEQ(v bool) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIsPreview), v))
	})
}

// IsPreviewNEQ applies the NEQ predicate on the "isPreview" field.
func IsPreviewNEQ(v bool) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIsPreview), v))
	})
}

// HasArticle applies the HasEdge predicate on the "article" edge.
func HasArticle() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticleTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArticleTable, ArticleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticleWith applies the HasEdge predicate on the "article" edge with a given conditions (other predicates).
func HasArticleWith(preds ...predicate.Article) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ArticleInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArticleTable, ArticleColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBlob applies the HasEdge predicate on the "blob" edge.
func HasBlob() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlobTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BlobTable, BlobColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBlobWith applies the HasEdge predicate on the "blob" edge with a given conditions (other predicates).
func HasBlobWith(preds ...predicate.Blob) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BlobInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BlobTable, BlobColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		p(s.Not())
	})
}
