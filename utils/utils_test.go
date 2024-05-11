package utils_test

import (
	"testing"

	"github.com/rjhoppe/go-compare-to-spy/utils"
)

func TestCheckTickerBadCharPass(t *testing.T) {
	err := utils.CheckTickerBadChars("AAPL")
	if err != nil {
		t.Errorf("Valid ticker input threw string validation error")
	}
}

func TestCheckTickerBadCharIntFail(t *testing.T) {
	expectedErr := "error: ticker input value contains a number"
	err := utils.CheckTickerBadChars("AB12")
	if err.Error() != expectedErr {
		t.Error("Invalid ticker containing numeric values not properly handled")
	}
}

func TestCheckTickerBadCharSymFail(t *testing.T) {
	expectedErr := "error: ticker input value contains a symbol"
	err := utils.CheckTickerBadChars("AB|*")
	if err.Error() != expectedErr {
		t.Error("Invalid ticker containing symbol values not properly handled")
	}
}

func TestCheckTickerBadCharBothFail(t *testing.T) {
	expectedErr := "error: ticker input value contains a number"
	err := utils.CheckTickerBadChars("AB2`")
	if err.Error() != expectedErr {
		t.Error("Invalid ticker containing a symbol and numeric value not handled properly")
	}
}
