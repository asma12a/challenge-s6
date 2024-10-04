package main

import (
	"github.com/asma12a/challenge-s6/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lpernett/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()

	app.Listen(":3001")
}
