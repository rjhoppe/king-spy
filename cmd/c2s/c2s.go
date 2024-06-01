/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package c2s

import (
	"fmt"
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

// compareSpyCmd represents the compareSpy command
var CompareSpyCmd = &cobra.Command{
	Use:   "c2s",
	Short: "Compares a ticker's performance to the SP500 over a specified time period",
	Long:  ``,
	Example: "  ks c2s aapl \n" +
		"  ks c2s aapl -t=1M \n" +
		"  ks c2s aapl -t=1Y \n",
	Run: func(cmd *cobra.Command, args []string) {
		// timeOptions = [5]string{"1M", "6M", "YTD", "1Y", "5Y"}
		ksCmd := "c2s"
		ticker := args[0]
		utils.TickerValidation(ticker)
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

		// go GetTickerPrice(key, secret, ticker, timeVal, "latest", ch1, &wg, ksCmd)
		// go GetTickerPrice(key, secret, ticker, timeVal, "history", ch2, &wg, ksCmd)
		// go GetTickerPrice(key, secret, "SPY", timeVal, "latest", ch3, &wg, ksCmd)
		// go GetTickerPrice(key, secret, "SPY", timeVal, "history", ch4, &wg, ksCmd)

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
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	CompareSpyCmd.Flags().StringP("time", "t", "", "A length of time for performance comparison")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareSpyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
