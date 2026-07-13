package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins for development.
		// Restrict this in production.
		return true
	},
}

func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn)

	// Register the client with the Hub.
	hub.Register(client)

	// Ensure cleanup when the function exits.
	defer func() {
		hub.Unregister(client)
		conn.Close()
	}()

	// Start the goroutine responsible for writing
	// messages to this client.
	go client.WritePump()

	// Keep reading from the socket until the client disconnects.
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
