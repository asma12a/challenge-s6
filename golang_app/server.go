package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/database"
	"github.com/asma12a/challenge-s6/database/redis"
	"github.com/asma12a/challenge-s6/handler"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/contrib/websocket"
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
	app.Use(func(c *fiber.Ctx) error {
		// add db client to context
		ctx := context.WithValue(c.Context(), "db", db_client)
		c.SetUserContext(ctx)
		return c.Next()
	})

	// Middleware pour vérifier si la requête est une demande de WebSocket
	// app.Use(func(c *fiber.Ctx) error {
	// if websocket.IsWebSocketUpgrade(c) {
	// return c.Next() // Si c'est une mise à niveau WebSocket, on continue
	// }
	// return fiber.ErrUpgradeRequired
	// })

	// Route WebSocket
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// Extraire les paramètres et afficher les informations
		log.Println("WebSocket connected!")
		log.Println("ID:", c.Params("id"))                   // Paramètre :id
		log.Println("Version:", c.Query("v"))                // Paramètre de query :v
		log.Println("Session cookie:", c.Cookies("session")) // Cookie :session

		// Lire et écrire des messages sur la connexion WebSocket
		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("Failed to read message:", err)
				break
			}
			log.Printf("Received: %s", msg)

			if err := c.WriteMessage(mt, msg); err != nil {
				log.Println("Failed to write message:", err)
				break
			}
		}
	}))

	// Routes API
	api := app.Group("/api")

	handler.EventHandler(api.Group("/events", middleware.IsAuthMiddleware), context.Background(), *service.NewEventService(db_client), *service.NewSportService(db_client))
	handler.SportHandler(api.Group("/sports", middleware.IsAuthMiddleware), context.Background(), *service.NewSportService(db_client))
	handler.UserHandler(api.Group("/users", middleware.IsAuthMiddleware), context.Background(), *service.NewUserService(db_client))
	handler.AuthHandler(api.Group("/auth"), context.Background(), *service.NewUserService(db_client), rdb)
	handler.EventTeamsHandler(api.Group("/event_teams", middleware.IsAuthMiddleware), context.Background(), *service.NewEventTeamsService(db_client), *service.NewEventService(db_client), *service.NewTeamService(db_client))
	handler.MessageHandler(api.Group("/message", middleware.IsAuthMiddleware), context.Background(), *service.NewMessageService(db_client), *service.NewEventService(db_client), *service.NewUserService(db_client))
	handler.TeamHandler(api.Group("/teams", middleware.IsAuthMiddleware), context.Background(), *service.NewTeamService(db_client))

	// Any other routes: Not Found
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Not Found",
		})
	})

	// Démarre le serveur
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.APIPort)))
}
