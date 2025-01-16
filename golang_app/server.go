//	@title			Challenge S6 API
//	@version		1.0
//	@description	API pour gérer des groupes de personnes autour d'une thématique
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Support Technique
//	@contact.url	http://www.example.com/support
//	@contact.email	support@example.com

//	@license.name	MIT
//	@license.url	https://opensource.org/licenses/MIT

// @host		localhost:3001
// @BasePath	/api
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/database"
	"github.com/asma12a/challenge-s6/database/redis"
	_ "github.com/asma12a/challenge-s6/docs"
	"github.com/asma12a/challenge-s6/handler"
	"github.com/asma12a/challenge-s6/internal/ws"
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
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	config.LoadEnvironmentFile()

	dbClient := database.GetClient()
	rdb := redis.GetClient()

	defer dbClient.Close()
	defer rdb.Close()

	app := fiber.New()

	app.Static("/images", "./images")

	// Middleware
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
	app.Use(requestid.New()) // Générer un identifiant unique pour chaque requête
	app.Use(func(c *fiber.Ctx) error {
		// Ajouter le client DB au contexte de la requête
		ctx := context.WithValue(c.Context(), "db", dbClient)
		c.SetUserContext(ctx)
		return c.Next()
	})

	// Route pour afficher Swagger UI
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Créer un Hub WebSocket
	hub := ws.NewHub()
	go hub.Run()

	// Route WebSocket
	app.Get("/ws", websocket.New(ws.WebSocketHandler(hub)))

	notificationService, err := service.NewNotificationService()
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation du service de notification : %v", err)
	}

	eventService := service.NewEventService(dbClient)

	// Lancer le cron pour les notifications dans une goroutine
	ctx := context.Background() // Contexte global
	go handler.NotifyPlayersBeforeEventCron(ctx, *notificationService, *eventService, rdb)

	// Routes API (sans authentification middleware)
	api := app.Group("/api")
	handler.NotificationHandler(api.Group("/notifications", middleware.IsAuthMiddleware), context.Background(), *notificationService, *service.NewEventService(dbClient), rdb)
	handler.EventHandler(api.Group("/events", middleware.IsAuthMiddleware), context.Background(), *service.NewEventService(dbClient), *service.NewSportService(dbClient), *service.NewTeamService(dbClient), *service.NewTeamUserService(dbClient))
	handler.SportHandler(api.Group("/sports", middleware.IsAuthMiddleware), context.Background(), *service.NewSportService(dbClient))
	handler.UserHandler(api.Group("/users", middleware.IsAuthMiddleware), context.Background(), *service.NewUserService(dbClient))
	handler.AuthHandler(api.Group("/auth"), context.Background(), *service.NewUserService(dbClient), *service.NewTeamUserService(dbClient), rdb, *notificationService)
	handler.MessageHandler(api.Group("/message", middleware.IsAuthMiddleware), context.Background(), *service.NewMessageService(dbClient), *service.NewEventService(dbClient), *service.NewUserService(dbClient))
	handler.SportStatLabelsHandler(api.Group("/sportstatlabels", middleware.IsAuthMiddleware), context.Background(), *service.NewSportStatLabelsService(dbClient), *service.NewSportService(dbClient), *service.NewEventService(dbClient), *service.NewUserService(dbClient), *notificationService, rdb)
	handler.ActionLogHandler(api.Group("/actionlogs", middleware.IsAuthMiddleware), context.Background(), *service.NewActionLogService(dbClient), *service.NewUserService(dbClient))
	userService := service.NewUserService(dbClient)
	oauthHandler := handler.NewOAuthHandler(userService)

	app.Get("/auth/google/login", oauthHandler.OAuthLoginHandler)
	app.Get("/auth/google/callback", oauthHandler.OAuthCallbackHandler)

	// Route de gestion des erreurs (Not Found)
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Route not Found",
		})
	})

	// Serveur Fiber
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.APIPort)))
}
