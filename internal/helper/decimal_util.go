package helper

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

// StringToDecimal converts a string representation of a number to decimal.Decimal
// It handles both integer and decimal numbers, including negative values
// Returns decimal.Decimal and error if conversion fails
func StringToDecimal(value string) (decimal.Decimal, error) {
	// Handle empty string
	if value == "" {
		return decimal.Zero, fmt.Errorf("empty string cannot be converted to decimal")
	}

	// Trim any whitespace
	value = strings.TrimSpace(value)

	// Check if the string is just a minus sign
	if value == "-" {
		return decimal.Zero, fmt.Errorf("invalid number format: single minus sign")
	}

	// Convert string to decimal
	result, err := decimal.NewFromString(value)
	if err != nil {
		return decimal.Zero, fmt.Errorf("failed to convert string to decimal: %v", err)
	}

	return result, nil
}

//redeclared same function name and same objective on internal/helper/parsing_type_data_util.go
//// StringToFloat64 converts a string to float64 using decimal for precise conversion
//// Useful when you need to interface with systems requiring float64
//func StringToFloat64(value string) (float64, error) {
//	dec, err := StringToDecimal(value)
//	if err != nil {
//		return 0, err
//	}
//
//	return dec.InexactFloat64(), nil
//}
