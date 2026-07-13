package main

import (
	"log"
	"net/http"

	"websockets/internal/exchange"
	"websockets/internal/model"
	ws "websockets/internal/websocket"
)

func main() {

	// Create the Hub
	hub := ws.NewHub()
	go hub.Run()

	// Create a trade channel
	tradeChan := make(chan model.Trade)

	// Create exchange using Factory
	ex, err := exchange.New("binance")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Binance
	if err := ex.Connect(); err != nil {
		log.Fatal(err)
	}

	// Read trades
	go ex.ReadTrades(tradeChan)

	// Forward trades to Hub
	go func() {
		for trade := range tradeChan {
			hub.Broadcast(trade)
		}
	}()

	// WebSocket endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWS(hub, w, r)
	})

	log.Println("Server listening on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
