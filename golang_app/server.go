package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/database"
	"github.com/asma12a/challenge-s6/database/redis"
	"github.com/asma12a/challenge-s6/handler"
	"github.com/asma12a/challenge-s6/internal/ws"
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
	// Charger les variables d'environnement
	config.LoadEnvironmentFile()

	// Initialiser les clients de la base de données et Redis
	dbClient := database.GetClient()
	rdb := redis.GetClient()

	// Assurez-vous de fermer les connexions lorsque le programme se termine
	defer dbClient.Close()
	defer rdb.Close()

	// Créer une nouvelle instance de l'application Fiber
	app := fiber.New()

	// Middleware
	app.Use(cors.New())     // Activer CORS
	app.Use(compress.New()) // Compression des réponses
	app.Use(etag.New())     // Gestion des ETag
	app.Use(favicon.New())  // Gestion du favicon
	app.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"error": "Too many requests",
			})
		},
	})) // Limiter le nombre de requêtes
	app.Use(logger.New())    // Logger des requêtes
	app.Use(recover.New())   // Récupérer après des erreurs
	app.Use(requestid.New()) // Générer un identifiant unique pour chaque requête
	app.Use(func(c *fiber.Ctx) error {
		// Ajouter le client DB au contexte de la requête
		ctx := context.WithValue(c.Context(), "db", dbClient)
		c.SetUserContext(ctx)
		return c.Next()
	})

	// Créer un Hub WebSocket
	hub := ws.NewHub()
	go hub.Run() // Lancer le Hub dans une goroutine

	// Route WebSocket
	app.Get("/ws", websocket.New(ws.WebSocketHandler(hub)))

	// Routes API (sans authentification middleware)
	api := app.Group("/api")
	handler.EventHandler(api.Group("/events"), context.Background(), *service.NewEventService(dbClient), *service.NewSportService(dbClient))
	handler.SportHandler(api.Group("/sports"), context.Background(), *service.NewSportService(dbClient))
	handler.UserHandler(api.Group("/users"), context.Background(), *service.NewUserService(dbClient))
	handler.AuthHandler(api.Group("/auth"), context.Background(), *service.NewUserService(dbClient), rdb)
	handler.MessageHandler(api.Group("/message"), context.Background(), *service.NewMessageService(dbClient), *service.NewEventService(dbClient), *service.NewUserService(dbClient))
	handler.TeamHandler(api.Group("/teams"), context.Background(), *service.NewTeamService(dbClient))

	// Route de gestion des erreurs (Not Found)
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Not Found",
		})
	})

	// Démarrer le serveur Fiber
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.APIPort)))
}
