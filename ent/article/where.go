// Code generated by ent, DO NOT EDIT.

package article

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// ArticleGUID applies equality check predicate on the "article_guid" field. It's identical to ArticleGUIDEQ.
func ArticleGUID(v uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArticleGUID), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Short applies equality check predicate on the "short" field. It's identical to ShortEQ.
func Short(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShort), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// UploadDate applies equality check predicate on the "upload_date" field. It's identical to UploadDateEQ.
func UploadDate(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUploadDate), v))
	})
}

// TitleNormalized applies equality check predicate on the "title_normalized" field. It's identical to TitleNormalizedEQ.
func TitleNormalized(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleNormalized), v))
	})
}

// ArticleGUIDEQ applies the EQ predicate on the "article_guid" field.
func ArticleGUIDEQ(v uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldArticleGUID), v))
	})
}

// ArticleGUIDNEQ applies the NEQ predicate on the "article_guid" field.
func ArticleGUIDNEQ(v uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldArticleGUID), v))
	})
}

// ArticleGUIDIn applies the In predicate on the "article_guid" field.
func ArticleGUIDIn(vs ...uuid.UUID) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldArticleGUID), v...))
	})
}

// ArticleGUIDNotIn applies the NotIn predicate on the "article_guid" field.
func ArticleGUIDNotIn(vs ...uuid.UUID) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldArticleGUID), v...))
	})
}

// ArticleGUIDGT applies the GT predicate on the "article_guid" field.
func ArticleGUIDGT(v uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldArticleGUID), v))
	})
}

// ArticleGUIDGTE applies the GTE predicate on the "article_guid" field.
func ArticleGUIDGTE(v uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldArticleGUID), v))
	})
}

// ArticleGUIDLT applies the LT predicate on the "article_guid" field.
func ArticleGUIDLT(v uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldArticleGUID), v))
	})
}

// ArticleGUIDLTE applies the LTE predicate on the "article_guid" field.
func ArticleGUIDLTE(v uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldArticleGUID), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// ShortEQ applies the EQ predicate on the "short" field.
func ShortEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShort), v))
	})
}

// ShortNEQ applies the NEQ predicate on the "short" field.
func ShortNEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShort), v))
	})
}

// ShortIn applies the In predicate on the "short" field.
func ShortIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldShort), v...))
	})
}

// ShortNotIn applies the NotIn predicate on the "short" field.
func ShortNotIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldShort), v...))
	})
}

// ShortGT applies the GT predicate on the "short" field.
func ShortGT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShort), v))
	})
}

// ShortGTE applies the GTE predicate on the "short" field.
func ShortGTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShort), v))
	})
}

// ShortLT applies the LT predicate on the "short" field.
func ShortLT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShort), v))
	})
}

// ShortLTE applies the LTE predicate on the "short" field.
func ShortLTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShort), v))
	})
}

// ShortContains applies the Contains predicate on the "short" field.
func ShortContains(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldShort), v))
	})
}

// ShortHasPrefix applies the HasPrefix predicate on the "short" field.
func ShortHasPrefix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldShort), v))
	})
}

// ShortHasSuffix applies the HasSuffix predicate on the "short" field.
func ShortHasSuffix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldShort), v))
	})
}

// ShortEqualFold applies the EqualFold predicate on the "short" field.
func ShortEqualFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldShort), v))
	})
}

// ShortContainsFold applies the ContainsFold predicate on the "short" field.
func ShortContainsFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldShort), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// UploadDateEQ applies the EQ predicate on the "upload_date" field.
func UploadDateEQ(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUploadDate), v))
	})
}

// UploadDateNEQ applies the NEQ predicate on the "upload_date" field.
func UploadDateNEQ(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUploadDate), v))
	})
}

// UploadDateIn applies the In predicate on the "upload_date" field.
func UploadDateIn(vs ...time.Time) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUploadDate), v...))
	})
}

// UploadDateNotIn applies the NotIn predicate on the "upload_date" field.
func UploadDateNotIn(vs ...time.Time) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUploadDate), v...))
	})
}

// UploadDateGT applies the GT predicate on the "upload_date" field.
func UploadDateGT(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUploadDate), v))
	})
}

// UploadDateGTE applies the GTE predicate on the "upload_date" field.
func UploadDateGTE(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUploadDate), v))
	})
}

// UploadDateLT applies the LT predicate on the "upload_date" field.
func UploadDateLT(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUploadDate), v))
	})
}

// UploadDateLTE applies the LTE predicate on the "upload_date" field.
func UploadDateLTE(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUploadDate), v))
	})
}

// TitleNormalizedEQ applies the EQ predicate on the "title_normalized" field.
func TitleNormalizedEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedNEQ applies the NEQ predicate on the "title_normalized" field.
func TitleNormalizedNEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedIn applies the In predicate on the "title_normalized" field.
func TitleNormalizedIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitleNormalized), v...))
	})
}

// TitleNormalizedNotIn applies the NotIn predicate on the "title_normalized" field.
func TitleNormalizedNotIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitleNormalized), v...))
	})
}

// TitleNormalizedGT applies the GT predicate on the "title_normalized" field.
func TitleNormalizedGT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedGTE applies the GTE predicate on the "title_normalized" field.
func TitleNormalizedGTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedLT applies the LT predicate on the "title_normalized" field.
func TitleNormalizedLT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedLTE applies the LTE predicate on the "title_normalized" field.
func TitleNormalizedLTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedContains applies the Contains predicate on the "title_normalized" field.
func TitleNormalizedContains(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedHasPrefix applies the HasPrefix predicate on the "title_normalized" field.
func TitleNormalizedHasPrefix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedHasSuffix applies the HasSuffix predicate on the "title_normalized" field.
func TitleNormalizedHasSuffix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedEqualFold applies the EqualFold predicate on the "title_normalized" field.
func TitleNormalizedEqualFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedContainsFold applies the ContainsFold predicate on the "title_normalized" field.
func TitleNormalizedContainsFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitleNormalized), v))
	})
}

// HasImages applies the HasEdge predicate on the "images" edge.
func HasImages() predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImagesTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImagesWith applies the HasEdge predicate on the "images" edge with a given conditions (other predicates).
func HasImagesWith(preds ...predicate.Image) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ImagesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ImagesTable, ImagesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
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
func Not(p predicate.Article) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		p(s.Not())
	})
}