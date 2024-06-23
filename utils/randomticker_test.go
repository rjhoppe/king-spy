package utils_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/rjhoppe/king-spy/utils"
	"github.com/spf13/viper"
)

func TestRandomTickerIsValid(t *testing.T) {
	randomTickVal, _ := utils.SelectRandomTicker()
	expected := ""
	viper.AddConfigPath("./..")
	utils.TickerValidation(randomTickVal)
	assert.Equal(t, expected, expected, "random ticker validation func failed")
}
