package tchou

import (
	"strconv"
)

func softParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return float64(0)
	}
	return f
}
