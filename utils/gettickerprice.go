package utils

import (
	"fmt"
	"log"
	"strings"
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
			fmt.Println("")
			fmt.Printf("Data retrieval error: ensure ticker %v has existed longer than timeframe \n", strings.ToUpper(ticker))
			fmt.Println("Ticker needs to have been publicly traded for longer than a year to run the implicit 'YTD' timeframe")
			log.Fatal(err)
		}
		ch <- tickerPrice
	} else {
		tickerPrice, err := jsonparser.GetFloat(body, "trade", "p")
		if err != nil {
			fmt.Println("")
			fmt.Printf("Data retrieval error: ensure ticker %v has existed longer than timeframe \n", strings.ToUpper(ticker))
			log.Fatal(err)
		}
		ch <- tickerPrice
	}
}
