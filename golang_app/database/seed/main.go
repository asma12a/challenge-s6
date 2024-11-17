package main

import (
	"context"
	"log"
	"time"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/database"
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	config.LoadEnvironmentFile()

	db_client := database.GetClient()
	defer db_client.Close()

	ctx := context.Background()

	dropAllData(ctx, db_client)
	seedUsers(ctx, db_client)
	seedSports(ctx, db_client)
	seedEvents(ctx, db_client)

	log.Println("Database seeding completed!")
}

func seedSports(ctx context.Context, db_client *ent.Client) {
	sports := []ent.Sport{
		{Name: "Football", Color: "0000FF"},
		{Name: "Basketball", Color: "FFA500"},
		{Name: "Tennis", Color: "008000"},
		{Name: "Running", Color: "FF0000"},
	}

	for _, sport := range sports {
		_, err := db_client.Sport.Create().
			SetName(sport.Name).
			SetColor(sport.Color).
			Save(ctx)
		if err != nil {
			log.Fatalf("Failed creating sport: %v", err)
		}
	}

	log.Println("Sports seeded!")
}

func seedEvents(ctx context.Context, db_client *ent.Client) {
	// foot, basket, tennis, running
	sports := db_client.Sport.Query().AllX(ctx)

	for _, sport := range sports {
		_, err := db_client.Event.Create().
			SetName(gofakeit.Name()).
			SetAddress(gofakeit.Address().Address).
			SetSport(sport).
			SetEventCode("1234").
			SetDate(time.Now().Format(time.DateOnly)).
			Save(ctx)
		if err != nil {
			log.Fatalf("Failed creating event: %v", err)
		}
	}

	log.Println("Events seeded!")
}

func seedUsers(ctx context.Context, db_client *ent.Client) {

	// create admin user
	admin_user, admin_err := entity.NewUser("", "", "admin")
	if admin_err != nil {
		log.Fatalf("Failed creating admin user: %v", admin_err)
	}
	_, err := db_client.User.Create().
		SetName("admin").
		SetEmail("admin@admin.test").
		SetPassword(admin_user.Password).
		SetRoles([]string{"admin"}).
		Save(ctx)
	if err != nil {
		log.Fatalf("Failed creating admin user: %v", err)
	}

	// default user
	default_user, err := entity.NewUser("", "", "password")
	if err != nil {
		log.Fatalf("Failed creating default user: %v", err)
	}
	_, err = db_client.User.Create().
		SetName("user").
		SetEmail("user@user.test").
		SetPassword(default_user.Password).
		Save(ctx)
	if err != nil {
		log.Fatalf("Failed creating default user: %v", err)
	}

	for i := 0; i < 10; i++ {
		_, err := db_client.User.Create().
			SetName(gofakeit.Name()).
			SetEmail(gofakeit.Email()).
			SetPassword(default_user.Password).
			Save(ctx)
		if err != nil {
			log.Printf("Failed creating user: %v", err)
			continue
		}
	}

	log.Println("Users seeded!")
}

func dropAllData(ctx context.Context, db_client *ent.Client) {
	// execute sql query to drop all data
	db_client.User.Delete().ExecX(ctx)
	db_client.Sport.Delete().ExecX(ctx)
	db_client.Event.Delete().ExecX(ctx)
	log.Println("All data dropped!")
}
