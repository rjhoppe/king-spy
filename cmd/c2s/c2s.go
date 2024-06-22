package c2s

import (
	"fmt"
	"strings"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

var (
	spyPositive    string
	tickerPositive string
	deltaPositive  string
	timeVal        string
)

// Compare2SpyCmd represents the c2s command
var Compare2SpyCmd = &cobra.Command{
	Use:   "c2s",
	Short: "Compares a ticker's performance to the SP500 over a specified time period",
	Long:  ``,
	Example: "  ks c2s aapl \n" +
		"  ks c2s aapl -t=1M \n" +
		"  ks c2s aapl -t=1Y \n",
	Run: func(cmd *cobra.Command, args []string) {
		ksCmd := "c2s"
		ticker := args[0]
		utils.TickerValidation(ticker)
		ticker = strings.ToLower(ticker)
		timeArg, _ := cmd.Flags().GetString("time")
		if timeArg == "" {
			timeVal = "YTD"
		} else {
			timeVal = timeArg
		}

		_, key, secret := config.Init()
		wg := sync.WaitGroup{}

		ch1 := make(chan float64)
		ch2 := make(chan float64)
		ch3 := make(chan float64)
		ch4 := make(chan float64)

		cfg := utils.GetTickerPriceConfig{
			Key:    key,
			Secret: secret,
			Wg:     &wg,
			Cmd:    ksCmd,
		}

		go utils.GetTickPrice(cfg, ticker, timeVal, "latest", ch1)
		go utils.GetTickPrice(cfg, ticker, timeVal, "history", ch2)
		go utils.GetTickPrice(cfg, "SPY", timeVal, "latest", ch3)
		go utils.GetTickPrice(cfg, "SPY", timeVal, "history", ch4)

		wg.Wait()

		tickerLatest := float64(<-ch1)
		tickerHist := float64(<-ch2)
		spyLatest := float64(<-ch3)
		spyHist := float64(<-ch4)

		o := Output{
			spyHist:      spyHist,
			spyLatest:    spyLatest,
			tickerHist:   tickerHist,
			tickerLatest: tickerLatest,
			tickerVal:    ticker,
			timeVal:      timeVal,
		}

		formatOutputC2S(o)
	},
}

func init() {
	Compare2SpyCmd.Flags().StringP("time", "t", "", "A length of time for performance comparison")
}

func GetTickerPrice(key string, secret string, ticker string, timeVal string, urlType string, ch chan float64, wg *sync.WaitGroup, cmd string) {
	wg.Add(1)
	defer wg.Done()

	t := utils.TimeAssignVals{
		TimeVal: timeVal,
		Ticker:  ticker,
		Cmd:     cmd,
		UrlType: urlType,
	}

	u := utils.AssignTime(t)
	url := utils.AssignUrl(t, u)
	fmt.Println(url)
	body, _ := utils.GetRequest(key, secret, url)

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
