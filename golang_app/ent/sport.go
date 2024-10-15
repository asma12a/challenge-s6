// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/asma12a/challenge-s6/ent/sport"
)

// Sport is the model entity for the Sport schema.
type Sport struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SportQuery when eager-loading is set.
	Edges        SportEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SportEdges holds the relations/edges for other nodes in the graph.
type SportEdges struct {
	// Event holds the value of the event edge.
	Event []*Event `json:"event,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading.
func (e SportEdges) EventOrErr() ([]*Event, error) {
	if e.loadedTypes[0] {
		return e.Event, nil
	}
	return nil, &NotLoadedError{edge: "event"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sport) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sport.FieldID, sport.FieldName, sport.FieldImageURL:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sport fields.
func (s *Sport) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sport.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				s.ID = value.String
			}
		case sport.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case sport.FieldImageURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image_url", values[i])
			} else if value.Valid {
				s.ImageURL = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Sport.
// This includes values selected through modifiers, order, etc.
func (s *Sport) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryEvent queries the "event" edge of the Sport entity.
func (s *Sport) QueryEvent() *EventQuery {
	return NewSportClient(s.config).QueryEvent(s)
}

// Update returns a builder for updating this Sport.
// Note that you need to call Sport.Unwrap() before calling this method if this Sport
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sport) Update() *SportUpdateOne {
	return NewSportClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Sport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sport) Unwrap() *Sport {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sport is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sport) String() string {
	var builder strings.Builder
	builder.WriteString("Sport(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("image_url=")
	builder.WriteString(s.ImageURL)
	builder.WriteByte(')')
	return builder.String()
}

// Sports is a parsable slice of Sport.
type Sports []*Sport
