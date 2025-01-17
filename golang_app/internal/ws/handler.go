package ws

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Message struct {
	EventID string `json:"event_id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

func WebSocketHandler(hub *Hub, eventID, userID string) func(*websocket.Conn) {

	return func(conn *websocket.Conn) {
		hub.register <- conn
		defer func() {
			hub.unregister <- conn
		}()

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Erreur de lecture:", err)
				break
			}

			var msg Message
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Printf("Erreur de décodage JSON: %v. Message reçu: %s", err, string(message))
				continue
			}

			if msg.EventID == "" || msg.UserID == "" || msg.Content == "" {
				errMessage := `{"error":"Missing event_id, user_id, or content"}`
				conn.WriteMessage(websocket.TextMessage, []byte(errMessage))
				continue
			}

			err = conn.WriteMessage(websocket.TextMessage, []byte(`{"self":true,"content":"`+msg.Content+`"}`))
			if err != nil {
				log.Println("Erreur d'envoi du message à l'expéditeur:", err)
				break
			}

			hub.BroadcastToOthers(conn, []byte(`{"self":false,"content":"`+msg.Content+`"}`))
		}
	}
}
