// Code generated by ent, DO NOT EDIT.

package custompage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IconId applies equality check predicate on the "iconId" field. It's identical to IconIdEQ.
func IconId(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIconId), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNormalized applies equality check predicate on the "titleNormalized" field. It's identical to TitleNormalizedEQ.
func TitleNormalized(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleNormalized), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// IconIdEQ applies the EQ predicate on the "iconId" field.
func IconIdEQ(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIconId), v))
	})
}

// IconIdNEQ applies the NEQ predicate on the "iconId" field.
func IconIdNEQ(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIconId), v))
	})
}

// IconIdIn applies the In predicate on the "iconId" field.
func IconIdIn(vs ...string) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldIconId), v...))
	})
}

// IconIdNotIn applies the NotIn predicate on the "iconId" field.
func IconIdNotIn(vs ...string) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldIconId), v...))
	})
}

// IconIdGT applies the GT predicate on the "iconId" field.
func IconIdGT(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIconId), v))
	})
}

// IconIdGTE applies the GTE predicate on the "iconId" field.
func IconIdGTE(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIconId), v))
	})
}

// IconIdLT applies the LT predicate on the "iconId" field.
func IconIdLT(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIconId), v))
	})
}

// IconIdLTE applies the LTE predicate on the "iconId" field.
func IconIdLTE(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIconId), v))
	})
}

// IconIdContains applies the Contains predicate on the "iconId" field.
func IconIdContains(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIconId), v))
	})
}

// IconIdHasPrefix applies the HasPrefix predicate on the "iconId" field.
func IconIdHasPrefix(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIconId), v))
	})
}

// IconIdHasSuffix applies the HasSuffix predicate on the "iconId" field.
func IconIdHasSuffix(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIconId), v))
	})
}

// IconIdIsNil applies the IsNil predicate on the "iconId" field.
func IconIdIsNil() predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldIconId)))
	})
}

// IconIdNotNil applies the NotNil predicate on the "iconId" field.
func IconIdNotNil() predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldIconId)))
	})
}

// IconIdEqualFold applies the EqualFold predicate on the "iconId" field.
func IconIdEqualFold(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIconId), v))
	})
}

// IconIdContainsFold applies the ContainsFold predicate on the "iconId" field.
func IconIdContainsFold(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIconId), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// TitleNormalizedEQ applies the EQ predicate on the "titleNormalized" field.
func TitleNormalizedEQ(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedNEQ applies the NEQ predicate on the "titleNormalized" field.
func TitleNormalizedNEQ(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedIn applies the In predicate on the "titleNormalized" field.
func TitleNormalizedIn(vs ...string) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitleNormalized), v...))
	})
}

// TitleNormalizedNotIn applies the NotIn predicate on the "titleNormalized" field.
func TitleNormalizedNotIn(vs ...string) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitleNormalized), v...))
	})
}

// TitleNormalizedGT applies the GT predicate on the "titleNormalized" field.
func TitleNormalizedGT(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedGTE applies the GTE predicate on the "titleNormalized" field.
func TitleNormalizedGTE(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedLT applies the LT predicate on the "titleNormalized" field.
func TitleNormalizedLT(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedLTE applies the LTE predicate on the "titleNormalized" field.
func TitleNormalizedLTE(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedContains applies the Contains predicate on the "titleNormalized" field.
func TitleNormalizedContains(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedHasPrefix applies the HasPrefix predicate on the "titleNormalized" field.
func TitleNormalizedHasPrefix(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedHasSuffix applies the HasSuffix predicate on the "titleNormalized" field.
func TitleNormalizedHasSuffix(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedEqualFold applies the EqualFold predicate on the "titleNormalized" field.
func TitleNormalizedEqualFold(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitleNormalized), v))
	})
}

// TitleNormalizedContainsFold applies the ContainsFold predicate on the "titleNormalized" field.
func TitleNormalizedContainsFold(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitleNormalized), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.CustomPage {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// HasAttachments applies the HasEdge predicate on the "attachments" edge.
func HasAttachments() predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttachmentsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AttachmentsTable, AttachmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttachmentsWith applies the HasEdge predicate on the "attachments" edge with a given conditions (other predicates).
func HasAttachmentsWith(preds ...predicate.Attachment) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
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
func And(predicates ...predicate.CustomPage) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CustomPage) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
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
func Not(p predicate.CustomPage) predicate.CustomPage {
	return predicate.CustomPage(func(s *sql.Selector) {
		p(s.Not())
	})
}
