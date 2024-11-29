package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

// WebSocketHandler handles incoming WebSocket connections
func WebSocketHandler(hub *Hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		// Register the client in the Hub
		hub.register <- conn

		defer func() {
			// Unregister the client upon disconnection
			hub.unregister <- conn
		}()

		// Listen for messages from the client
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			// Broadcast the message to all clients
			hub.Broadcast(message)
		}
	}
}
