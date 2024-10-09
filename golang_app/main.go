package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asma12a/challenge-s6/api/handler"
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/infrastructure/ent/datastore"
	"github.com/asma12a/challenge-s6/infrastructure/ent/repository"
	"github.com/asma12a/challenge-s6/pkg/config"
	"github.com/asma12a/challenge-s6/usecase/event"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/lib/pq"
)

func main() {

	config.LoadEnvironmentFile(".env")
	client := newDBClient()
	defer client.Close()

	// Creates a new Fiber instance.
	app := fiber.New(fiber.Config{
		AppName:      "Fiber Ent Clean Architecture",
		ServerHeader: "Fiber",
	})

	// Use global middlewares.
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Create repositories.
	eventRepository := repository.NewEventRepoEnt(client)

	// Create all of our services.
	eventService := event.NewService(eventRepository)

	api := app.Group("/api")

	// Prepare our endpoints for the API.
	handler.NewEventHandler(api.Group("/v1/events"), context.Background(), eventService)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 3000.s
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.APIPort)))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening Posgres client: %v", err)
	}

	return client
}
