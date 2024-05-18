package utils

import (
	"fmt"
	"sync"

	"github.com/buger/jsonparser"
)

type GetTickerPriceConfig struct {
	Key    string
	Secret string
	Wg     *sync.WaitGroup
	Cmd    string
}

func GetTickPrice(cfg GetTickerPriceConfig, ticker string, timeVal string, urlType string, ch chan float64) {
	cfg.Wg.Add(1)
	defer cfg.Wg.Done()

	t := TimeAssignVals{
		TimeVal: timeVal,
		Ticker:  ticker,
		Cmd:     cfg.Cmd,
		UrlType: urlType,
	}

	u := AssignTime(t)
	url := AssignUrl(t, u)
	body, _ := GetRequest(cfg.Key, cfg.Secret, url)

	if urlType == "history" {
		tickerPrice, err := jsonparser.GetFloat(body, "trades", "[0]", "p")
		if err != nil {
			fmt.Printf("Error: Could not parse ticker asking price. %v %v", err, ticker)
			panic(err)
		}
		ch <- tickerPrice
	} else {
		tickerPrice, err := jsonparser.GetFloat(body, "trade", "p")
		if err != nil {
			fmt.Printf("Error: Could not parse ticker asking price. %v %v", err, ticker)
			panic(err)
		}
		ch <- tickerPrice
	}
}
