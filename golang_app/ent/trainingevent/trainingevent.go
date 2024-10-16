// Code generated by ent, DO NOT EDIT.

package trainingevent

import (
	"entgo.io/ent/dialect/sql"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

const (
	// Label holds the string label denoting the trainingevent type in the database.
	Label = "training_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventTrainingID holds the string denoting the event_training_id field in the database.
	FieldEventTrainingID = "event_training_id"
	// FieldTeamID holds the string denoting the team_id field in the database.
	FieldTeamID = "team_id"
	// Table holds the table name of the trainingevent in the database.
	Table = "training_events"
)

// Columns holds all SQL columns for trainingevent fields.
var Columns = []string{
	FieldID,
	FieldEventTrainingID,
	FieldTeamID,
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
	// EventTrainingIDValidator is a validator for the "event_training_id" field. It is called by the builders before save.
	EventTrainingIDValidator func(string) error
	// TeamIDValidator is a validator for the "team_id" field. It is called by the builders before save.
	TeamIDValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)

// OrderOption defines the ordering options for the TrainingEvent queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEventTrainingID orders the results by the event_training_id field.
func ByEventTrainingID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventTrainingID, opts...).ToFunc()
}

// ByTeamID orders the results by the team_id field.
func ByTeamID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeamID, opts...).ToFunc()
}
