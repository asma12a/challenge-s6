// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldEventCode holds the string denoting the event_code field in the database.
	FieldEventCode = "event_code"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldIsPublic holds the string denoting the is_public field in the database.
	FieldIsPublic = "is_public"
	// FieldIsFinished holds the string denoting the is_finished field in the database.
	FieldIsFinished = "is_finished"
	// EdgeEventType holds the string denoting the event_type edge name in mutations.
	EdgeEventType = "event_type"
	// EdgeSport holds the string denoting the sport edge name in mutations.
	EdgeSport = "sport"
	// EdgeUserStatsID holds the string denoting the user_stats_id edge name in mutations.
	EdgeUserStatsID = "user_stats_id"
	// EdgeFootEventID holds the string denoting the foot_event_id edge name in mutations.
	EdgeFootEventID = "foot_event_id"
	// EdgeBasketEventID holds the string denoting the basket_event_id edge name in mutations.
	EdgeBasketEventID = "basket_event_id"
	// EdgeTennisEventID holds the string denoting the tennis_event_id edge name in mutations.
	EdgeTennisEventID = "tennis_event_id"
	// EdgeRunningEventID holds the string denoting the running_event_id edge name in mutations.
	EdgeRunningEventID = "running_event_id"
	// EdgeTrainingEventID holds the string denoting the training_event_id edge name in mutations.
	EdgeTrainingEventID = "training_event_id"
	// Table holds the table name of the event in the database.
	Table = "events"
	// EventTypeTable is the table that holds the event_type relation/edge.
	EventTypeTable = "events"
	// EventTypeInverseTable is the table name for the EventType entity.
	// It exists in this package in order to avoid circular dependency with the "eventtype" package.
	EventTypeInverseTable = "event_types"
	// EventTypeColumn is the table column denoting the event_type relation/edge.
	EventTypeColumn = "event_type_id"
	// SportTable is the table that holds the sport relation/edge.
	SportTable = "events"
	// SportInverseTable is the table name for the Sport entity.
	// It exists in this package in order to avoid circular dependency with the "sport" package.
	SportInverseTable = "sports"
	// SportColumn is the table column denoting the sport relation/edge.
	SportColumn = "sport_id"
	// UserStatsIDTable is the table that holds the user_stats_id relation/edge.
	UserStatsIDTable = "user_stats"
	// UserStatsIDInverseTable is the table name for the UserStats entity.
	// It exists in this package in order to avoid circular dependency with the "userstats" package.
	UserStatsIDInverseTable = "user_stats"
	// UserStatsIDColumn is the table column denoting the user_stats_id relation/edge.
	UserStatsIDColumn = "event_id"
	// FootEventIDTable is the table that holds the foot_event_id relation/edge.
	FootEventIDTable = "foot_events"
	// FootEventIDInverseTable is the table name for the FootEvent entity.
	// It exists in this package in order to avoid circular dependency with the "footevent" package.
	FootEventIDInverseTable = "foot_events"
	// FootEventIDColumn is the table column denoting the foot_event_id relation/edge.
	FootEventIDColumn = "event_id"
	// BasketEventIDTable is the table that holds the basket_event_id relation/edge.
	BasketEventIDTable = "basket_events"
	// BasketEventIDInverseTable is the table name for the BasketEvent entity.
	// It exists in this package in order to avoid circular dependency with the "basketevent" package.
	BasketEventIDInverseTable = "basket_events"
	// BasketEventIDColumn is the table column denoting the basket_event_id relation/edge.
	BasketEventIDColumn = "event_id"
	// TennisEventIDTable is the table that holds the tennis_event_id relation/edge.
	TennisEventIDTable = "tennis_events"
	// TennisEventIDInverseTable is the table name for the TennisEvent entity.
	// It exists in this package in order to avoid circular dependency with the "tennisevent" package.
	TennisEventIDInverseTable = "tennis_events"
	// TennisEventIDColumn is the table column denoting the tennis_event_id relation/edge.
	TennisEventIDColumn = "event_id"
	// RunningEventIDTable is the table that holds the running_event_id relation/edge.
	RunningEventIDTable = "running_events"
	// RunningEventIDInverseTable is the table name for the RunningEvent entity.
	// It exists in this package in order to avoid circular dependency with the "runningevent" package.
	RunningEventIDInverseTable = "running_events"
	// RunningEventIDColumn is the table column denoting the running_event_id relation/edge.
	RunningEventIDColumn = "event_id"
	// TrainingEventIDTable is the table that holds the training_event_id relation/edge.
	TrainingEventIDTable = "training_events"
	// TrainingEventIDInverseTable is the table name for the TrainingEvent entity.
	// It exists in this package in order to avoid circular dependency with the "trainingevent" package.
	TrainingEventIDInverseTable = "training_events"
	// TrainingEventIDColumn is the table column denoting the training_event_id relation/edge.
	TrainingEventIDColumn = "event_id"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAddress,
	FieldEventCode,
	FieldDate,
	FieldCreatedAt,
	FieldIsPublic,
	FieldIsFinished,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"event_type_id",
	"sport_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// EventCodeValidator is a validator for the "event_code" field. It is called by the builders before save.
	EventCodeValidator func(int16) error
	// DateValidator is a validator for the "date" field. It is called by the builders before save.
	DateValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultIsPublic holds the default value on creation for the "is_public" field.
	DefaultIsPublic bool
	// DefaultIsFinished holds the default value on creation for the "is_finished" field.
	DefaultIsFinished bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)

// OrderOption defines the ordering options for the Event queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByEventCode orders the results by the event_code field.
func ByEventCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventCode, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByIsPublic orders the results by the is_public field.
func ByIsPublic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPublic, opts...).ToFunc()
}

// ByIsFinished orders the results by the is_finished field.
func ByIsFinished(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFinished, opts...).ToFunc()
}

// ByEventTypeField orders the results by event_type field.
func ByEventTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventTypeStep(), sql.OrderByField(field, opts...))
	}
}

// BySportField orders the results by sport field.
func BySportField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSportStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserStatsIDCount orders the results by user_stats_id count.
func ByUserStatsIDCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStatsIDStep(), opts...)
	}
}

// ByUserStatsID orders the results by user_stats_id terms.
func ByUserStatsID(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStatsIDStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFootEventIDCount orders the results by foot_event_id count.
func ByFootEventIDCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFootEventIDStep(), opts...)
	}
}

// ByFootEventID orders the results by foot_event_id terms.
func ByFootEventID(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFootEventIDStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByBasketEventIDCount orders the results by basket_event_id count.
func ByBasketEventIDCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBasketEventIDStep(), opts...)
	}
}

// ByBasketEventID orders the results by basket_event_id terms.
func ByBasketEventID(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBasketEventIDStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTennisEventIDCount orders the results by tennis_event_id count.
func ByTennisEventIDCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTennisEventIDStep(), opts...)
	}
}

// ByTennisEventID orders the results by tennis_event_id terms.
func ByTennisEventID(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTennisEventIDStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRunningEventIDCount orders the results by running_event_id count.
func ByRunningEventIDCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRunningEventIDStep(), opts...)
	}
}

// ByRunningEventID orders the results by running_event_id terms.
func ByRunningEventID(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRunningEventIDStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTrainingEventIDCount orders the results by training_event_id count.
func ByTrainingEventIDCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTrainingEventIDStep(), opts...)
	}
}

// ByTrainingEventID orders the results by training_event_id terms.
func ByTrainingEventID(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTrainingEventIDStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEventTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, EventTypeTable, EventTypeColumn),
	)
}
func newSportStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SportInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SportTable, SportColumn),
	)
}
func newUserStatsIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserStatsIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserStatsIDTable, UserStatsIDColumn),
	)
}
func newFootEventIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FootEventIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FootEventIDTable, FootEventIDColumn),
	)
}
func newBasketEventIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BasketEventIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, BasketEventIDTable, BasketEventIDColumn),
	)
}
func newTennisEventIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TennisEventIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TennisEventIDTable, TennisEventIDColumn),
	)
}
func newRunningEventIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RunningEventIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RunningEventIDTable, RunningEventIDColumn),
	)
}
func newTrainingEventIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TrainingEventIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TrainingEventIDTable, TrainingEventIDColumn),
	)
}