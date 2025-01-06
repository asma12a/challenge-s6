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

var clients = make(map[*websocket.Conn]*client)
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)

func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = &client{}
			log.Println("New client connected")

		case message := <-broadcast:
			// Broadcast the message to all connected clients
			log.Println("Broadcasting message:", message)
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) {
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						log.Println("Write error:", err)
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-unregister:
			delete(clients, connection)
			log.Println("Client disconnected")
		}
	}
}

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

	// WebSocket route to handle incoming connections
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			c.mu.Lock()
			if !c.isClosing {
				c.isClosing = true
				if err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
					log.Println("Error closing WebSocket:", err)
				}
			}
			c.mu.Unlock()
			unregister <- c
			c.Close()
		}()

		register <- c
		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				return
			}

			if messageType == websocket.TextMessage {
				// Broadcast received message to all clients
				broadcast <- string(message)
			} else {
				log.Println("Non-text message received")
			}
		}
	}))

	// Start the WebSocket hub
	go runHub()

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
