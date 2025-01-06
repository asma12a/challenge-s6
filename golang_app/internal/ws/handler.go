package ws

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

// Message struct pour vérifier les données reçues
type Message struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

// WebSocketHandler handles incoming WebSocket connections
func WebSocketHandler(hub *Hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
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

			// Décoder le message JSON sans validation de parsing
			var msg Message
			err = json.Unmarshal(message, &msg)

			// Vérifier si toutes les informations sont présentes
			if msg.EventID == "" || msg.UserID == "" || msg.Content == "" {
				// Envoi d'une erreur si l'une des informations est manquante
				errMessage := `{"error":"Missing event_id, user_id, or content"}`
				conn.WriteMessage(websocket.TextMessage, []byte(errMessage))
				continue
			}

			// Envoi immédiat à l'expéditeur (self: true)
			err = conn.WriteMessage(websocket.TextMessage, []byte(`{"self":true,"content":"`+msg.Content+`"}`))
			if err != nil {
				log.Println("Error sending message to sender:", err)
				break
			}

			// Diffusion aux autres clients (self: false)
			hub.BroadcastToOthers(conn, []byte(`{"self":false,"content":"`+msg.Content+`"}`))
		}
	}
}
