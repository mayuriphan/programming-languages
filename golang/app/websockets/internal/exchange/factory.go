package exchange

import "fmt"

func New(name string) (Exchange, error) {

	switch name {

	case "binance":
		return &Binance{}, nil

	default:
		return nil, fmt.Errorf("unsupported exchange: %s", name)
	}
}
