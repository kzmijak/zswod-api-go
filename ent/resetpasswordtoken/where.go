// Code generated by ent, DO NOT EDIT.

package resetpasswordtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// ExpiryTime applies equality check predicate on the "expiryTime" field. It's identical to ExpiryTimeEQ.
func ExpiryTime(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiryTime), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ResetPasswordToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.ResetPasswordToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.ResetPasswordToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.ResetPasswordToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// ExpiryTimeEQ applies the EQ predicate on the "expiryTime" field.
func ExpiryTimeEQ(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiryTime), v))
	})
}

// ExpiryTimeNEQ applies the NEQ predicate on the "expiryTime" field.
func ExpiryTimeNEQ(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiryTime), v))
	})
}

// ExpiryTimeIn applies the In predicate on the "expiryTime" field.
func ExpiryTimeIn(vs ...time.Time) predicate.ResetPasswordToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExpiryTime), v...))
	})
}

// ExpiryTimeNotIn applies the NotIn predicate on the "expiryTime" field.
func ExpiryTimeNotIn(vs ...time.Time) predicate.ResetPasswordToken {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExpiryTime), v...))
	})
}

// ExpiryTimeGT applies the GT predicate on the "expiryTime" field.
func ExpiryTimeGT(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiryTime), v))
	})
}

// ExpiryTimeGTE applies the GTE predicate on the "expiryTime" field.
func ExpiryTimeGTE(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiryTime), v))
	})
}

// ExpiryTimeLT applies the LT predicate on the "expiryTime" field.
func ExpiryTimeLT(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiryTime), v))
	})
}

// ExpiryTimeLTE applies the LTE predicate on the "expiryTime" field.
func ExpiryTimeLTE(v time.Time) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiryTime), v))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ResetPasswordToken) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ResetPasswordToken) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
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
func Not(p predicate.ResetPasswordToken) predicate.ResetPasswordToken {
	return predicate.ResetPasswordToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
