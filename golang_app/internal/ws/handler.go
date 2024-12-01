package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

// WebSocketHandler handles incoming WebSocket connections
func WebSocketHandler(hub *Hub) func(*websocket.Conn) {
	return func(conn *websocket.Conn) {
		// Authentifier l'utilisateur avant d'autoriser la connexion WebSocket
		//	req := conn.Locals("request").(*fiber.Request)
		// token := string(req.Header.Peek("Authorization"))
		// if token == "" {
		// 	conn.WriteMessage(websocket.TextMessage, []byte("Erreur: Pas d'authentification"))
		// 	conn.Close()
		// 	return
		// }

		// // Valider le token ici, par exemple avec votre middleware JWT
		// if !isValidToken(token) {
		// 	conn.WriteMessage(websocket.TextMessage, []byte("Erreur: Token invalide"))
		// 	conn.Close()
		// 	return
		// }

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

			// Enregistrez le message dans la base de données
			// storeMessage(db, message)

			// Diffuser le message à tous les clients
			hub.Broadcast(message)
		}
	}
}

func isValidToken(token string) bool {
	// Implémentez ici la logique pour valider le token JWT
	return true
}

// func storeMessage(db *gorm.DB, message []byte) {
// 	// Créez une structure Message avec les informations nécessaires
// 	msg := models.Message{
// 		EventID: 1, // L'ID de l'événement ou à récupérer via la connexion
// 		Content: string(message),
// 	}

// 	// Insérez le message dans la base de données
// 	if err := db.Create(&msg).Error; err != nil {
// 		log.Println("Erreur lors du stockage du message:", err)
// 	}
// }
