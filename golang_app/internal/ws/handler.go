package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

// WebSocketHandler handles incoming WebSocket connections
func WebSocketHandler(hub *Hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		// TO DO : ajouter logique d'authentification

		// Si l'utilisateur est authentifié, alors procéder
		hub.register <- conn
		defer func() {
			hub.unregister <- conn
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}
			hub.Broadcast(message)
		}
	}
}
