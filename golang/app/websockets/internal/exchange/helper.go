package exchange

import "strconv"

func toFloat(value string) float64 {

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return f
}
