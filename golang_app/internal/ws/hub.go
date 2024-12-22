package ws

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

// Hub structure for managing WebSocket connections
type Hub struct {
	clients    map[*websocket.Conn]bool
	broadcast  chan []byte
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
	mu         sync.Mutex
}

// NewHub initializes a new Hub
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*websocket.Conn]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

// Run is the method that listens for messages and broadcasts to all connected clients
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Close()
			}
		case message := <-h.broadcast:
			log.Printf("Broadcasting message to %d clients", len(h.clients))

			for client := range h.clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println("Error sending message:", err)
					client.Close()
					delete(h.clients, client)
				}
			}
		}
	}
}

// Broadcast sends a message to all connected clients
func (h *Hub) Broadcast(message []byte) {
	h.broadcast <- message
}

func (h *Hub) BroadcastToOthers(sender *websocket.Conn, message []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for client := range h.clients {
		if client != sender {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error sending message:", err)
				client.Close()
				delete(h.clients, client)
			}
		}
	}
}
