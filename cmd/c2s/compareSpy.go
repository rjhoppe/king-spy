/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package c2s

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/buger/jsonparser"
	"github.com/fatih/color"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/spf13/cobra"
)

var spyPositive string
var tickerPositive string
var deltaPositive string
var timeVal string
// var timeOptions [5]string

func getTickerPrice(key string, secret string, ticker string, timeVal string, urlType string, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	var url string
	curTime := time.Now()
	var startTime string
	var endTime string

	if (urlType == "history") {
		switch timeVal {
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
		// url = "https://data.alpaca.markets/v2/stocks/" + ticker + "/trades?limit=1&start=2023-04-25T13:59:52-04:00&end=2023-04-25T18:59:52-04:00&feed=iex&currency=USD"
		// use asof
		url = "https://data.alpaca.markets/v2/stocks/" + ticker + "/trades?limit=1&start=" + startTime + "&end=" + endTime + "&feed=iex&currency=USD"
		// fmt.Println(url)
	} else {
		url = "https://data.alpaca.markets/v2/stocks/" + ticker + "/trades/latest?feed=iex"
	} 

	// https://data.alpaca.markets/v2/stocks/SPY/trades?limit=1&start=2024-04-27T00:00:45-04:00&end=2023-04-30T00:42:45-04:00&feed=iex&currency=USD

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", key)
	req.Header.Add("APCA-API-SECRET-KEY", secret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: Could not get data. %v", err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: Could not read response body. %v", err)
	}
	
	if (urlType == "history") {
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

// compareSpyCmd represents the compareSpy command
var CompareSpyCmd = &cobra.Command{
	Use:   "c2s",
	Short: "Compares a ticker's performance to spy",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
			// timeOptions = [5]string{"1M", "6M", "YTD", "1Y", "5Y"}
			ticker := args[0]
			timeArg, _ := cmd.Flags().GetString("time")
			
			if timeArg == "" {
				timeVal = "YTD"
			} else {
				timeVal = timeArg
			}

			wg := new(sync.WaitGroup)
			_, key, secret := config.Init()

			ch1 := make(chan float64)
			ch2 := make(chan float64)
			ch3 := make(chan float64)
			ch4 := make(chan float64)

			wg.Add(4)
			go getTickerPrice(key, secret, ticker, timeVal, "latest", ch1, wg)
			go getTickerPrice(key, secret, ticker, timeVal, "history", ch2, wg)
			go getTickerPrice(key, secret, "SPY", timeVal, "latest", ch3, wg)
			go getTickerPrice(key, secret, "SPY", timeVal, "history", ch4, wg)

			spyHist := float64(<-ch4)
			spyLatest := float64(<-ch3)
			tickerHist := float64(<-ch2)
			tickerLatest := float64(<-ch1)
			spyPerf := ((spyLatest - spyHist) / spyHist) * 100
			if (spyPerf > 0) {
				spyPositive = "+"
			} else {
				spyPositive = ""
			}

			tickerPerf := ((tickerLatest - tickerHist) / tickerHist) * 100
			if (tickerPerf > 0) {
				tickerPositive = "+"
			} else {
				tickerPositive = ""
			}

			deltaPerf := tickerPerf - spyPerf
			if (deltaPerf > 0) {
				deltaPositive = "+"
			} else {
				deltaPositive = ""
			}

			if (spyPositive == "+") {
				spyC := color.New(color.FgGreen)
				fmt.Printf("SPY performance %v: ", timeVal)
				spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
			} else {
				spyC := color.New(color.FgRed)
				fmt.Printf("SPY performance %v: ", timeVal)
				spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
			}

			if (tickerPositive == "+") {
				tickerC := color.New(color.FgGreen)
				fmt.Printf("%v performance %v: ", strings.ToUpper(ticker), timeVal)
				tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
			} else {
				tickerC := color.New(color.FgRed)
				fmt.Printf("%v performance %v: ", strings.ToUpper(ticker), timeVal)
				tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
			}

			if (deltaPositive == "+") {
				deltaC := color.New(color.FgGreen)
				fmt.Printf("%v performance vs SPY: ", strings.ToUpper(ticker))
				deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
			} else {
				deltaC := color.New(color.FgRed)
				fmt.Printf("%v performance vs SPY: ", strings.ToUpper(ticker))
				deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
			}

			wg.Wait()
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	CompareSpyCmd.PersistentFlags().String("time", "", "A length of time for performance comparison")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compareSpyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
