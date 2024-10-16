// Code generated by ent, DO NOT EDIT.

package tennisevent

import (
	"entgo.io/ent/dialect/sql"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

const (
	// Label holds the string label denoting the tennisevent type in the database.
	Label = "tennis_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEventTennisID holds the string denoting the event_tennis_id field in the database.
	FieldEventTennisID = "event_tennis_id"
	// FieldTeamA holds the string denoting the team_a field in the database.
	FieldTeamA = "team_a"
	// FieldTeamB holds the string denoting the team_b field in the database.
	FieldTeamB = "team_b"
	// Table holds the table name of the tennisevent in the database.
	Table = "tennis_events"
)

// Columns holds all SQL columns for tennisevent fields.
var Columns = []string{
	FieldID,
	FieldEventTennisID,
	FieldTeamA,
	FieldTeamB,
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
	// EventTennisIDValidator is a validator for the "event_tennis_id" field. It is called by the builders before save.
	EventTennisIDValidator func(string) error
	// TeamAValidator is a validator for the "team_A" field. It is called by the builders before save.
	TeamAValidator func(string) error
	// TeamBValidator is a validator for the "team_B" field. It is called by the builders before save.
	TeamBValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() ulid.ID
)

// OrderOption defines the ordering options for the TennisEvent queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEventTennisID orders the results by the event_tennis_id field.
func ByEventTennisID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventTennisID, opts...).ToFunc()
}

// ByTeamA orders the results by the team_A field.
func ByTeamA(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeamA, opts...).ToFunc()
}

// ByTeamB orders the results by the team_B field.
func ByTeamB(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTeamB, opts...).ToFunc()
}
