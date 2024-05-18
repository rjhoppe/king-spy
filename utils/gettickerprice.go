package utils

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/buger/jsonparser"
)

type GetTickerPriceConfig struct {
	key     string
	secret  string
	ticker  string
	timeVal string
	urlType string
	ch      chan float64
	wg      *sync.WaitGroup
}

func GetTickPrice(cfg GetTickerPriceConfig) {
	var (
		url       string
		startTime string
		endTime   string
	)

	cfg.wg.Add(1)
	defer cfg.wg.Done()
	curTime := time.Now()
	if cfg.urlType == "history" {
		switch cfg.timeVal {
		case "1M":
			pastTimeVal := curTime.AddDate(0, -1, 0)
			startTime = pastTimeVal.Format(time.RFC3339)
			endTimeVal := pastTimeVal.Add(72 * time.Hour)
			endTime = endTimeVal.Format(time.RFC3339)
		case "6M":
			pastTimeVal := curTime.AddDate(0, -6, 0)
			startTime = pastTimeVal.Format(time.RFC3339)
			endTimeVal := pastTimeVal.Add(72 * time.Hour)
			endTime = endTimeVal.Format(time.RFC3339)
		case "1Y":
			pastTimeVal := curTime.AddDate(-1, 0, 0)
			startTime = pastTimeVal.Format(time.RFC3339)
			endTimeVal := pastTimeVal.Add(72 * time.Hour)
			endTime = endTimeVal.Format(time.RFC3339)
		case "3Y":
			pastTimeVal := curTime.AddDate(-3, 0, 0)
			startTime = pastTimeVal.Format(time.RFC3339)
			endTimeVal := pastTimeVal.Add(72 * time.Hour)
			endTime = endTimeVal.Format(time.RFC3339)
		// default is YTD
		default:
			curYear, _, _ := time.Now().Date()
			curYearString := strconv.Itoa(curYear)
			startTime = curYearString + "-01-01T00:00:00-00:00"
			endTime = curTime.Format(time.RFC3339)
		}
		url = "https://data.alpaca.markets/v2/stocks/" + cfg.ticker + "/trades?limit=1&start=" + startTime + "&end=" + endTime + "&feed=iex&currency=USD"
	} else {
		url = "https://data.alpaca.markets/v2/stocks/" + cfg.ticker + "/trades/latest?feed=iex"
	}

	body, _ := GetRequest(cfg.key, cfg.secret, url)
	if cfg.urlType == "history" {
		tickerPrice, err := jsonparser.GetFloat(body, "trades", "[0]", "p")
		if err != nil {
			fmt.Printf("Error: Could not parse ticker asking price. %v %v", err, cfg.ticker)
			panic(err)
		}
		cfg.ch <- tickerPrice
	} else {
		tickerPrice, err := jsonparser.GetFloat(body, "trade", "p")
		if err != nil {
			fmt.Printf("Error: Could not parse ticker asking price. %v %v", err, cfg.ticker)
			panic(err)
		}
		cfg.ch <- tickerPrice
	}
}
