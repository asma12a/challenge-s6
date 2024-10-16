// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/asma12a/challenge-s6/ent/predicate"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// ID filters vertices based on their ID field.
func ID(id ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id ulid.ID) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldName, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldAddress, v))
}

// EventCode applies equality check predicate on the "event_code" field. It's identical to EventCodeEQ.
func EventCode(v int16) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventCode, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// IsPublic applies equality check predicate on the "is_public" field. It's identical to IsPublicEQ.
func IsPublic(v bool) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldIsPublic, v))
}

// IsFinished applies equality check predicate on the "is_finished" field. It's identical to IsFinishedEQ.
func IsFinished(v bool) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldIsFinished, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldName, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldAddress, v))
}

// EventCodeEQ applies the EQ predicate on the "event_code" field.
func EventCodeEQ(v int16) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldEventCode, v))
}

// EventCodeNEQ applies the NEQ predicate on the "event_code" field.
func EventCodeNEQ(v int16) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldEventCode, v))
}

// EventCodeIn applies the In predicate on the "event_code" field.
func EventCodeIn(vs ...int16) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldEventCode, vs...))
}

// EventCodeNotIn applies the NotIn predicate on the "event_code" field.
func EventCodeNotIn(vs ...int16) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldEventCode, vs...))
}

// EventCodeGT applies the GT predicate on the "event_code" field.
func EventCodeGT(v int16) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldEventCode, v))
}

// EventCodeGTE applies the GTE predicate on the "event_code" field.
func EventCodeGTE(v int16) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldEventCode, v))
}

// EventCodeLT applies the LT predicate on the "event_code" field.
func EventCodeLT(v int16) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldEventCode, v))
}

// EventCodeLTE applies the LTE predicate on the "event_code" field.
func EventCodeLTE(v int16) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldEventCode, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v string) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...string) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v string) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v string) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v string) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v string) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldDate, v))
}

// DateContains applies the Contains predicate on the "date" field.
func DateContains(v string) predicate.Event {
	return predicate.Event(sql.FieldContains(FieldDate, v))
}

// DateHasPrefix applies the HasPrefix predicate on the "date" field.
func DateHasPrefix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasPrefix(FieldDate, v))
}

// DateHasSuffix applies the HasSuffix predicate on the "date" field.
func DateHasSuffix(v string) predicate.Event {
	return predicate.Event(sql.FieldHasSuffix(FieldDate, v))
}

// DateEqualFold applies the EqualFold predicate on the "date" field.
func DateEqualFold(v string) predicate.Event {
	return predicate.Event(sql.FieldEqualFold(FieldDate, v))
}

// DateContainsFold applies the ContainsFold predicate on the "date" field.
func DateContainsFold(v string) predicate.Event {
	return predicate.Event(sql.FieldContainsFold(FieldDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Event {
	return predicate.Event(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Event {
	return predicate.Event(sql.FieldLTE(FieldCreatedAt, v))
}

// IsPublicEQ applies the EQ predicate on the "is_public" field.
func IsPublicEQ(v bool) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldIsPublic, v))
}

// IsPublicNEQ applies the NEQ predicate on the "is_public" field.
func IsPublicNEQ(v bool) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldIsPublic, v))
}

// IsFinishedEQ applies the EQ predicate on the "is_finished" field.
func IsFinishedEQ(v bool) predicate.Event {
	return predicate.Event(sql.FieldEQ(FieldIsFinished, v))
}

// IsFinishedNEQ applies the NEQ predicate on the "is_finished" field.
func IsFinishedNEQ(v bool) predicate.Event {
	return predicate.Event(sql.FieldNEQ(FieldIsFinished, v))
}

// HasEventType applies the HasEdge predicate on the "event_type" edge.
func HasEventType() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EventTypeTable, EventTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventTypeWith applies the HasEdge predicate on the "event_type" edge with a given conditions (other predicates).
func HasEventTypeWith(preds ...predicate.EventType) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newEventTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSport applies the HasEdge predicate on the "sport" edge.
func HasSport() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SportTable, SportColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSportWith applies the HasEdge predicate on the "sport" edge with a given conditions (other predicates).
func HasSportWith(preds ...predicate.Sport) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newSportStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserStatsID applies the HasEdge predicate on the "user_stats_id" edge.
func HasUserStatsID() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UserStatsIDTable, UserStatsIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserStatsIDWith applies the HasEdge predicate on the "user_stats_id" edge with a given conditions (other predicates).
func HasUserStatsIDWith(preds ...predicate.UserStats) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newUserStatsIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFootEventID applies the HasEdge predicate on the "foot_event_id" edge.
func HasFootEventID() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FootEventIDTable, FootEventIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFootEventIDWith applies the HasEdge predicate on the "foot_event_id" edge with a given conditions (other predicates).
func HasFootEventIDWith(preds ...predicate.FootEvent) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newFootEventIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBasketEventID applies the HasEdge predicate on the "basket_event_id" edge.
func HasBasketEventID() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, BasketEventIDTable, BasketEventIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBasketEventIDWith applies the HasEdge predicate on the "basket_event_id" edge with a given conditions (other predicates).
func HasBasketEventIDWith(preds ...predicate.BasketEvent) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newBasketEventIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTennisEventID applies the HasEdge predicate on the "tennis_event_id" edge.
func HasTennisEventID() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TennisEventIDTable, TennisEventIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTennisEventIDWith applies the HasEdge predicate on the "tennis_event_id" edge with a given conditions (other predicates).
func HasTennisEventIDWith(preds ...predicate.TennisEvent) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newTennisEventIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRunningEventID applies the HasEdge predicate on the "running_event_id" edge.
func HasRunningEventID() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RunningEventIDTable, RunningEventIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRunningEventIDWith applies the HasEdge predicate on the "running_event_id" edge with a given conditions (other predicates).
func HasRunningEventIDWith(preds ...predicate.RunningEvent) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newRunningEventIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTrainingEventID applies the HasEdge predicate on the "training_event_id" edge.
func HasTrainingEventID() predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TrainingEventIDTable, TrainingEventIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTrainingEventIDWith applies the HasEdge predicate on the "training_event_id" edge with a given conditions (other predicates).
func HasTrainingEventIDWith(preds ...predicate.TrainingEvent) predicate.Event {
	return predicate.Event(func(s *sql.Selector) {
		step := newTrainingEventIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Event) predicate.Event {
	return predicate.Event(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Event) predicate.Event {
	return predicate.Event(sql.NotPredicates(p))
}