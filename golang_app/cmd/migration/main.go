package main

import (
	"context"

	"log"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/migrate"
	"github.com/asma12a/challenge-s6/infrastructure/ent/datastore"
	"github.com/asma12a/challenge-s6/pkg/config"
)

func main() {
	config.LoadEnvironmentFile(".env")

	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening Postgres client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
