package main

import (
	"context"
	"log"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/database"
	"github.com/asma12a/challenge-s6/ent/migrate"
)

func main() {
	config.LoadEnvironmentFile()

	db_client := database.GetClient()
	defer db_client.Close()

	err := db_client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	)
	if err != nil {
		log.Fatalf("Error creating schema: %v", err)
	}
	log.Println("Database migrated!")
}
