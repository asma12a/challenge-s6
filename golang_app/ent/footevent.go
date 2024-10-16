// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/asma12a/challenge-s6/ent/footevent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// FootEvent is the model entity for the FootEvent schema.
type FootEvent struct {
	config `json:"-"`
	// ID of the ent.
	ID ulid.ID `json:"id,omitempty"`
	// EventFootID holds the value of the "event_foot_id" field.
	EventFootID string `json:"event_foot_id,omitempty"`
	// TeamA holds the value of the "team_A" field.
	TeamA string `json:"team_A,omitempty"`
	// TeamB holds the value of the "team_B" field.
	TeamB         string `json:"team_B,omitempty"`
	event_foot_id *ulid.ID
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FootEvent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case footevent.FieldEventFootID, footevent.FieldTeamA, footevent.FieldTeamB:
			values[i] = new(sql.NullString)
		case footevent.FieldID:
			values[i] = new(ulid.ID)
		case footevent.ForeignKeys[0]: // event_foot_id
			values[i] = &sql.NullScanner{S: new(ulid.ID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FootEvent fields.
func (fe *FootEvent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case footevent.FieldID:
			if value, ok := values[i].(*ulid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				fe.ID = *value
			}
		case footevent.FieldEventFootID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field event_foot_id", values[i])
			} else if value.Valid {
				fe.EventFootID = value.String
			}
		case footevent.FieldTeamA:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_A", values[i])
			} else if value.Valid {
				fe.TeamA = value.String
			}
		case footevent.FieldTeamB:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field team_B", values[i])
			} else if value.Valid {
				fe.TeamB = value.String
			}
		case footevent.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field event_foot_id", values[i])
			} else if value.Valid {
				fe.event_foot_id = new(ulid.ID)
				*fe.event_foot_id = *value.S.(*ulid.ID)
			}
		default:
			fe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FootEvent.
// This includes values selected through modifiers, order, etc.
func (fe *FootEvent) Value(name string) (ent.Value, error) {
	return fe.selectValues.Get(name)
}

// Update returns a builder for updating this FootEvent.
// Note that you need to call FootEvent.Unwrap() before calling this method if this FootEvent
// was returned from a transaction, and the transaction was committed or rolled back.
func (fe *FootEvent) Update() *FootEventUpdateOne {
	return NewFootEventClient(fe.config).UpdateOne(fe)
}

// Unwrap unwraps the FootEvent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fe *FootEvent) Unwrap() *FootEvent {
	_tx, ok := fe.config.driver.(*txDriver)
	if !ok {
		panic("ent: FootEvent is not a transactional entity")
	}
	fe.config.driver = _tx.drv
	return fe
}

// String implements the fmt.Stringer.
func (fe *FootEvent) String() string {
	var builder strings.Builder
	builder.WriteString("FootEvent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fe.ID))
	builder.WriteString("event_foot_id=")
	builder.WriteString(fe.EventFootID)
	builder.WriteString(", ")
	builder.WriteString("team_A=")
	builder.WriteString(fe.TeamA)
	builder.WriteString(", ")
	builder.WriteString("team_B=")
	builder.WriteString(fe.TeamB)
	builder.WriteByte(')')
	return builder.String()
}

// FootEvents is a parsable slice of FootEvent.
type FootEvents []*FootEvent
