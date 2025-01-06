package database

import (
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/ent"
	_ "github.com/asma12a/challenge-s6/ent/runtime"
	_ "github.com/lib/pq" // Required for "postgres" driver
)

// Returns a ent ORM client
func GetClient() *ent.Client {

	env := os.Getenv("ENV")
	sslMode := "disable"

	if env == "production" {
		sslMode = "enable"

	}
	// Postgres DSN
	db_url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		config.Env.DBUser,
		config.Env.DBPass,
		config.Env.DBHost,
		config.Env.DBPort,
		config.Env.DBName,
		sslMode,
	)
	var entOptions []ent.Option
	// entOptions = append(entOptions, ent.Debug()) // Display DB debug logs

	db_client, err := ent.Open(dialect.Postgres, db_url, entOptions...)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Connected to database")

	return db_client
}
