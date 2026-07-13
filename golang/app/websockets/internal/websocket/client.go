package websocket

import (
	"log"

	"github.com/gorilla/websocket"

	"websockets/internal/model"
)

type Client struct {
	conn *websocket.Conn
	send chan model.Trade
}

func NewClient(conn *websocket.Conn) *Client {

	return &Client{
		conn: conn,
		send: make(chan model.Trade),
	}
}

func (c *Client) WritePump() {

	defer c.conn.Close()

	for trade := range c.send {

		err := c.conn.WriteJSON(trade)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
