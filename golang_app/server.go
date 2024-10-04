package main

import (
	"fmt"
	"log"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadEnvironmentFile()

	db_client := database.GetClient()
	defer db_client.Close()

	app := fiber.New()
	app.Use(cors.New())

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.APIPort)))
}
