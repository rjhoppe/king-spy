package sectors

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/fatih/color"
	"github.com/rjhoppe/king-spy/data"
	"github.com/rjhoppe/king-spy/utils"
)

type Sector struct {
	Name  string
	Value float64
}

func CompareSectors(cfg GetSectorsConfig, timeVal string, stockVal string) {
	sectorPerf := []Sector{}
	wg := sync.WaitGroup{}
	ch1 := make(chan float64)
	ch2 := make(chan float64)
	for k, v := range data.SectorData {
		ticker := k
		go getSectorCurPrice(cfg, ticker, ch1)
		go getSectorHistPrice(cfg, ticker, timeVal, ch2)
		tickerLatest := float64(<-ch1)
		tickerHistPrice := float64(<-ch2)
		wg.Wait()

		tickPerf := ((tickerLatest - tickerHistPrice) / tickerHistPrice) * 100
		sectorStringVal := fmt.Sprintf("%v [%v]", v, k)
		sectorPerf = append(sectorPerf, Sector{sectorStringVal, tickPerf})
	}

	if stockVal != "NA" {
		stockName := data.TickerData[(strings.ToUpper(stockVal))]
		go getSectorCurPrice(cfg, stockVal, ch1)
		go getSectorHistPrice(cfg, stockVal, timeVal, ch2)
		tickerLatest := float64(<-ch1)
		tickerHistPrice := float64(<-ch2)
		wg.Wait()

		tickPerf := ((tickerLatest - tickerHistPrice) / tickerHistPrice) * 100
		if stockName != "" {
			sectorPerf = append(sectorPerf, Sector{color.YellowString(stockName + " " + "[" + strings.ToUpper(stockVal) + "]"), tickPerf})
		} else {
			sectorPerf = append(sectorPerf, Sector{color.YellowString(strings.ToUpper(stockVal)), tickPerf})
		}
	}

	sort.Slice(sectorPerf, func(i, j int) bool {
		return sectorPerf[i].Value > sectorPerf[j].Value
	})

	fmt.Println("")
	fmt.Printf("%v (%v) \n", color.YellowString("Top Performing Sectors"), color.YellowString(timeVal))
	fmt.Println("==================================================================================")
	for i := range sectorPerf {
		if sectorPerf[i].Value > 0 {
			sectorValColored := color.New(color.FgGreen)
			fmt.Printf("%v: ", sectorPerf[i].Name)
			sectorValColored.Printf("+%.2f%% \n", sectorPerf[i].Value)
		} else {
			sectorValColored := color.New(color.FgRed)
			fmt.Printf("%v: ", sectorPerf[i].Name)
			sectorValColored.Printf("%.2f%% \n", sectorPerf[i].Value)
		}
	}
	fmt.Println("==================================================================================")
}

func getSectorCurPrice(cfg GetSectorsConfig, ticker string, ch chan float64) {
	t := utils.TimeAssignVals{
		TimeVal: "NA",
		Ticker:  ticker,
		Cmd:     "sectors",
		UrlType: "latest",
	}
	u := utils.AssignTime(t)
	url := utils.AssignUrl(t, u)
	body, _ := utils.GetRequest(cfg.Key, cfg.Secret, url)

	tickerPrice, err := jsonparser.GetFloat(body, "trade", "p")
	if err != nil {
		fmt.Println("")
		fmt.Printf("Data retrieval error: ensure ticker: %v has existed longer than timeframe \n", strings.ToUpper(ticker))
		log.Fatal(err)
	}

	ch <- tickerPrice
}

func getSectorHistPrice(cfg GetSectorsConfig, ticker string, timeVal string, ch chan float64) {
	t := utils.TimeAssignVals{
		TimeVal: timeVal,
		Ticker:  ticker,
		Cmd:     "sectors",
		UrlType: "history",
	}
	u := utils.AssignTime(t)
	url := utils.AssignUrl(t, u)
	body, _ := utils.GetRequest(cfg.Key, cfg.Secret, url)

	tickerPrice, err := jsonparser.GetFloat(body, "trades", "[0]", "p")
	if err != nil {
		fmt.Println("")
		fmt.Printf("Data retrieval error: ensure ticker: %v has existed longer than timeframe \n", strings.ToUpper(ticker))
		log.Fatal(err)
	}
	ch <- tickerPrice
}
