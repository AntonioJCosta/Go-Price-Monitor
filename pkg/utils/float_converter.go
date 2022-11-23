package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Float64Converter is a function that converts a string to a float64
func Float64Converter(s string) (float64, error) {
	s = strings.ReplaceAll(s, "R$", "")
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", ".")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("Erro ao converter string para float64: %w", err)
	}
	return f, nil
}
