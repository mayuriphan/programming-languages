package exchange

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"

	"websockets/internal/model"
)

const binanceWS = "wss://stream.binance.com:9443/ws/btcusdt@trade"

type Binance struct {
	conn *websocket.Conn
}

// Binance response
type binanceTrade struct {
	EventType string `json:"e"`
	Symbol    string `json:"s"`
	Price     string `json:"p"`
	Quantity  string `json:"q"`
}

func (b *Binance) Connect() error {

	conn, _, err := websocket.DefaultDialer.Dial(binanceWS, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to Binance: %w", err)
	}

	b.conn = conn

	fmt.Println("Connected to Binance")

	return nil
}

func (b *Binance) ReadTrades(out chan<- model.Trade) {

	defer b.conn.Close()

	for {

		_, message, err := b.conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		var trade binanceTrade

		err = json.Unmarshal(message, &trade)
		if err != nil {
			continue
		}

		out <- model.Trade{
			Exchange: "Binance",
			Symbol:   trade.Symbol,
			Price:    toFloat(trade.Price),
			Quantity: toFloat(trade.Quantity),
		}
	}
}
