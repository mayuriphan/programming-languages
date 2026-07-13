package websocket

import (
	"log"

	"websockets/internal/model"
)

type Hub struct {
	clients map[*Client]bool

	register   chan *Client
	unregister chan *Client
	broadcast  chan model.Trade
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan model.Trade),
	}
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.register:
			h.clients[client] = true
			log.Println("Client connected")

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				log.Println("Client disconnected")
			}

		case trade := <-h.broadcast:
			for client := range h.clients {

				select {

				case client.send <- trade:

				// Client is slow or disconnected.
				// Remove it so it doesn't block the hub.
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// Register adds a client to the hub.
func (h *Hub) Register(client *Client) {
	h.register <- client
}

// Unregister removes a client from the hub.
func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

// Broadcast publishes a trade to all connected clients.
func (h *Hub) Broadcast(trade model.Trade) {
	h.broadcast <- trade
}
