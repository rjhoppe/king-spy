package utils

import (
	"math/rand"

	"github.com/rjhoppe/king-spy/data"
)

func SelectRandomTicker() (string, string) {
	i := 0
	tickerList := make([]string, len(data.TickerData))
	for k := range data.TickerData {
		tickerList[i] = k
		i++
	}

	randomIndex := rand.Intn(len(tickerList))
	randTickVal := tickerList[randomIndex]
	randTickName := data.TickerData[randTickVal]
	return randTickVal, randTickName
}
