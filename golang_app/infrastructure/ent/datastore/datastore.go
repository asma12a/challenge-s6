package datastore

import (
	"fmt"

	"entgo.io/ent/dialect"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/pkg/config"
	_ "github.com/lib/pq"
)

// New returns data source name
func New() string {
	config.LoadEnvironmentFile(".env")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Env.DBHost,
		config.Env.DBPort,
		config.Env.DBUser,
		config.Env.DBPass,
		config.Env.DBName,
	)

	return dsn
}

// NewClient returns an orm client
func NewClient() (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	dsn := New()

	return ent.Open(dialect.Postgres, dsn, entOptions...)
}
