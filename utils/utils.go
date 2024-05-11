package utils

import (
	"errors"
	"math"
	"strings"
)

// Not sure this is used
func Round(x float64) float64 {
	return math.Round(x*100) / 100
}

func CheckTickerBadChars(x string) error {
	intVals := [10]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	specialChars := "!@#$%^&*()-_+={}[]|;:'<>?/~`"
	for _, i := range intVals {
		check := strings.Contains(x, i)
		if check {
			return errors.New("error: ticker input value contains a number")
		}
	}

	if strings.ContainsAny(x, specialChars) {
		return errors.New("error: ticker input value contains a symbol")
	}

	return nil
}
