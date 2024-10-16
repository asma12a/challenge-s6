// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldEventTypeID holds the string denoting the event_type_id field in the database.
	FieldEventTypeID = "event_type_id"
	// FieldSportID holds the string denoting the sport_id field in the database.
	FieldSportID = "sport_id"
	// EdgeEventType holds the string denoting the event_type edge name in mutations.
	EdgeEventType = "event_type"
	// EdgeSport holds the string denoting the sport edge name in mutations.
	EdgeSport = "sport"
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
	FieldEventTypeID,
	FieldSportID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// EventTypeIDValidator is a validator for the "event_type_id" field. It is called by the builders before save.
	EventTypeIDValidator func(string) error
	// SportIDValidator is a validator for the "sport_id" field. It is called by the builders before save.
	SportIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
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

// ByEventTypeID orders the results by the event_type_id field.
func ByEventTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventTypeID, opts...).ToFunc()
}

// BySportID orders the results by the sport_id field.
func BySportID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSportID, opts...).ToFunc()
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
