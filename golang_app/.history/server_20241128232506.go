package main

import (
	"context"
	"fmt"
	"log"
	"sync"

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

type client struct {
	isClosing bool
	mu        sync.Mutex
}

var (
	clients    = make(map[*websocket.Conn]*client)
	clientsMu  sync.Mutex
	register   = make(chan *websocket.Conn)
	broadcast  = make(chan string)
	unregister = make(chan *websocket.Conn)
)

func runHub() {
	for {
		select {
		case connection := <-register:
			clientsMu.Lock()
			clients[connection] = &client{}
			clientsMu.Unlock()
			log.Printf("New client connected: %p", connection)

		case message := <-broadcast:
			// Broadcast the message to all connected clients
			clientsMu.Lock()
			log.Printf("Broadcasting message to %d clients: %s", len(clients), message)
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) {
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						log.Printf("Write error for client %p: %v", connection, err)
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}
			clientsMu.Unlock()

		case connection := <-unregister:
			delete(clients, connection)
			log.Printf("Client disconnected: %p. Total clients: %d", connection, len(clients))
			clientsMu.Lock()
			delete(clients, connection)
			clientsMu.Unlock()
			log.Printf("Client disconnected: %p", connection)
		}
	}
}

func main() {
	config.LoadEnvironmentFile()

	dbClient := database.GetClient()
	rdb := redis.GetClient()

	defer dbClient.Close()
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
		// Add db client to context
		ctx := context.WithValue(c.Context(), "db", dbClient)
		c.SetUserContext(ctx)
		return c.Next()
	})

	// WebSocket route to handle incoming connections
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			clientObj := clients[c]
			clientObj.mu.Lock()
			defer clientObj.mu.Unlock()

			if !clientObj.isClosing {
				clientObj.isClosing = true
				unregister <- c // Signale la déconnexion
			}
			c.Close()
		}()

		register <- c
		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				clientObj := clients[c]
				clientObj.mu.Lock()
				if !clientObj.isClosing {
					clientObj.isClosing = true
					unregister <- c // Signale la déconnexion
				}
				clientObj.mu.Unlock()
				log.Printf("Read error for client %p: %v", c, err)
				return
			}

			if messageType == websocket.TextMessage {
				broadcast <- string(message)
			} else {
				log.Printf("Non-text message received from client %p", c)
			}
		}
	}))

	// Start the WebSocket hub
	go runHub()

	// Routes API
	api := app.Group("/api")

	handler.EventHandler(api.Group("/events", middleware.IsAuthMiddleware), context.Background(), *service.NewEventService(dbClient), *service.NewSportService(dbClient))
	handler.SportHandler(api.Group("/sports", middleware.IsAuthMiddleware), context.Background(), *service.NewSportService(dbClient))
	handler.UserHandler(api.Group("/users", middleware.IsAuthMiddleware), context.Background(), *service.NewUserService(dbClient))
	handler.AuthHandler(api.Group("/auth"), context.Background(), *service.NewUserService(dbClient), rdb)
	handler.EventTeamsHandler(api.Group("/event_teams", middleware.IsAuthMiddleware), context.Background(), *service.NewEventTeamsService(dbClient), *service.NewEventService(dbClient), *service.NewTeamService(dbClient))
	handler.MessageHandler(api.Group("/message", middleware.IsAuthMiddleware), context.Background(), *service.NewMessageService(dbClient), *service.NewEventService(dbClient), *service.NewUserService(dbClient))
	handler.TeamHandler(api.Group("/teams", middleware.IsAuthMiddleware), context.Background(), *service.NewTeamService(dbClient))

	// Any other routes: Not Found
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"error": "Not Found",
		})
	})

	// Start the server
	log.Fatal(app.Listen(fmt.Sprintf(":%s", config.Env.APIPort)))
}
