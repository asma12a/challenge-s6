package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/database"
	"github.com/asma12a/challenge-s6/database/redis"
	"github.com/asma12a/challenge-s6/handler"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	config.LoadEnvironmentFile()

	db_client := database.GetClient()
	rdb := redis.GetClient()

	defer db_client.Close()
	defer rdb.Close()

	app := fiber.New()

	// Middlewares
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(etag.New())
	app.Use(favicon.New())
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"error": "Too many requests",
			})
		},
	}))
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Routes
	api := app.Group("/api")

	handler.EventHandler(api.Group("/events"), context.Background(), *service.NewEventService(db_client), *service.NewEventTypeService(db_client), *service.NewSportService(db_client))
	handler.EventTypeHandler(api.Group("/event_types"), context.Background(), *service.NewEventTypeService(db_client))
	handler.SportHandler(api.Group("/sports"), context.Background(), *service.NewSportService(db_client))
	handler.UserHandler(api.Group("/users"), context.Background(), *service.NewUserService(db_client))
	handler.AuthHandler(api.Group("/auth"), context.Background(), *service.NewUserService(db_client), rdb)

	// Any other routes: Not Found
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Not Found",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.APIPort)))
}
