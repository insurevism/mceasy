package helper

import (
	"github.com/labstack/gommon/log"
	"strconv"
)

func StringToFloat64(s string) float64 {
	if s == "" {
		return float64(0)
	}

	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Errorf("failed parsing StringToFloat64 because of error: %v", err)
		return float64(0)
	}

	return value
}
