package exchange

import "websockets/internal/model"

type Exchange interface {
	Connect() error
	ReadTrades(chan<- model.Trade)
}
