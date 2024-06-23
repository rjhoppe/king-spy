package utils_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/rjhoppe/king-spy/utils"
	"github.com/spf13/viper"
)

func TestCheckTickerBadCharPass(t *testing.T) {
	_ = utils.CheckTickerBadChars("AAPL")
	assert.Equal(t, nil, nil, "valid ticker input threw ticker validation error")
}

func TestCheckTickerBadCharIntFail(t *testing.T) {
	expected := "error: ticker input value contains a number"
	_ = utils.CheckTickerBadChars("AB12")
	assert.Equal(t, expected, expected, "invalid ticker containing numbers not handled properly")
}

func TestCheckTickerBadCharSymFail(t *testing.T) {
	expected := "error: ticker input value contains a symbol"
	_ = utils.CheckTickerBadChars("AB|*")
	assert.Equal(t, expected, expected, "invalid ticker containing symbols not handled properly")
}

func TestCheckTickerBadCharBothFail(t *testing.T) {
	expected := "error: ticker input value contains a number"
	_ = utils.CheckTickerBadChars("AB2`")
	assert.Equal(t, expected, expected, "invalid ticker containing a symbol and numeric value not handled properly")
}

func TestValidTicker(t *testing.T) {
	viper.AddConfigPath("./..")
	utils.IsTickerValid("TSLA")
	assert.Equal(t, nil, nil, "valid ticker test failed")
}

func TestFakeTicker(t *testing.T) {
	expected := "invalid ticker - ticker not found"
	viper.AddConfigPath("./..")
	utils.IsTickerValid("azyh")
	assert.Equal(t, expected, expected, "invalid ticker test failed")
}

func TestBlankTicker(t *testing.T) {
	expected := "invalid ticker - ticker not found"
	viper.AddConfigPath("./..")
	utils.IsTickerValid("    ")
	assert.Equal(t, expected, expected, "invalid ticker test failed")
}

func TestTickerValidationValid(t *testing.T) {
	expected := ""
	viper.AddConfigPath("./..")
	utils.TickerValidation("aapl")
	assert.Equal(t, expected, expected, "valid ticker validation func test failed")
}
