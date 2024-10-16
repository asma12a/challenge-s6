// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "event_code", Type: field.TypeInt16},
		{Name: "date", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "is_public", Type: field.TypeBool, Default: false},
		{Name: "is_finished", Type: field.TypeBool, Default: false},
		{Name: "event_type_event", Type: field.TypeString, Nullable: true},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "events_event_types_event",
				Columns:    []*schema.Column{EventsColumns[8]},
				RefColumns: []*schema.Column{EventTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EventTypesColumns holds the columns for the "event_types" table.
	EventTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
	}
	// EventTypesTable holds the schema information for the "event_types" table.
	EventTypesTable = &schema.Table{
		Name:       "event_types",
		Columns:    EventTypesColumns,
		PrimaryKey: []*schema.Column{EventTypesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Default: "user"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserStatsColumns holds the columns for the "user_stats" table.
	UserStatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "user_id", Type: field.TypeString},
		{Name: "event_id", Type: field.TypeString},
	}
	// UserStatsTable holds the schema information for the "user_stats" table.
	UserStatsTable = &schema.Table{
		Name:       "user_stats",
		Columns:    UserStatsColumns,
		PrimaryKey: []*schema.Column{UserStatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_stats_events_user_stats_id",
				Columns:    []*schema.Column{UserStatsColumns[2]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "user_stats_users_user_stats",
				Columns:    []*schema.Column{UserStatsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		EventTypesTable,
		UsersTable,
		UserStatsTable,
	}
)

func init() {
	EventsTable.ForeignKeys[0].RefTable = EventTypesTable
	UserStatsTable.ForeignKeys[0].RefTable = EventsTable
	UserStatsTable.ForeignKeys[1].RefTable = UsersTable
}
