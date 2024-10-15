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
	// EdgeUserStatsID holds the string denoting the user_stats_id edge name in mutations.
	EdgeUserStatsID = "user_stats_id"
	// Table holds the table name of the event in the database.
	Table = "events"
	// UserStatsIDTable is the table that holds the user_stats_id relation/edge.
	UserStatsIDTable = "user_stats"
	// UserStatsIDInverseTable is the table name for the UserStats entity.
	// It exists in this package in order to avoid circular dependency with the "userstats" package.
	UserStatsIDInverseTable = "user_stats"
	// UserStatsIDColumn is the table column denoting the user_stats_id relation/edge.
	UserStatsIDColumn = "event_id"
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
func newUserStatsIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserStatsIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserStatsIDTable, UserStatsIDColumn),
	)
}
