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
func ID(id uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNormalized applies equality check predicate on the "titleNormalized" field. It's identical to TitleNormalizedEQ.
func TitleNormalized(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleNormalized), v))
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

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
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

// TitleNormalizedEQ applies the EQ predicate on the "titleNormalized" field.
func TitleNormalizedEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedNEQ applies the NEQ predicate on the "titleNormalized" field.
func TitleNormalizedNEQ(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedIn applies the In predicate on the "titleNormalized" field.
func TitleNormalizedIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitleNormalized), v...))
	})
}

// TitleNormalizedNotIn applies the NotIn predicate on the "titleNormalized" field.
func TitleNormalizedNotIn(vs ...string) predicate.Article {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitleNormalized), v...))
	})
}

// TitleNormalizedGT applies the GT predicate on the "titleNormalized" field.
func TitleNormalizedGT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedGTE applies the GTE predicate on the "titleNormalized" field.
func TitleNormalizedGTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedLT applies the LT predicate on the "titleNormalized" field.
func TitleNormalizedLT(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedLTE applies the LTE predicate on the "titleNormalized" field.
func TitleNormalizedLTE(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedContains applies the Contains predicate on the "titleNormalized" field.
func TitleNormalizedContains(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedHasPrefix applies the HasPrefix predicate on the "titleNormalized" field.
func TitleNormalizedHasPrefix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedHasSuffix applies the HasSuffix predicate on the "titleNormalized" field.
func TitleNormalizedHasSuffix(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedEqualFold applies the EqualFold predicate on the "titleNormalized" field.
func TitleNormalizedEqualFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedContainsFold applies the ContainsFold predicate on the "titleNormalized" field.
func TitleNormalizedContainsFold(v string) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitleNormalized), v))
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

// HasGallery applies the HasEdge predicate on the "gallery" edge.
func HasGallery() predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GalleryTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, GalleryTable, GalleryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGalleryWith applies the HasEdge predicate on the "gallery" edge with a given conditions (other predicates).
func HasGalleryWith(preds ...predicate.Gallery) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GalleryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, GalleryTable, GalleryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AuthorTable, AuthorPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, AuthorTable, AuthorPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAttachments applies the HasEdge predicate on the "attachments" edge.
func HasAttachments() predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentsTable, AttachmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentsWith applies the HasEdge predicate on the "attachments" edge with a given conditions (other predicates).
func HasAttachmentsWith(preds ...predicate.Attachment) predicate.Article {
	return predicate.Article(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentsTable, AttachmentsColumn),
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
