/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package c2s

import (
	"fmt"
	"strings"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/fatih/color"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

// type TimeAssignVals struct {
// 	timeVal string
// 	ticker string
// 	cmd string
// }

var (
	spyPositive    string
	tickerPositive string
	deltaPositive  string
	timeVal        string
)

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

		go GetTickerPrice(key, secret, ticker, timeVal, "latest", ch1, &wg, ksCmd)
		go GetTickerPrice(key, secret, ticker, timeVal, "history", ch2, &wg, ksCmd)
		go GetTickerPrice(key, secret, "SPY", timeVal, "latest", ch3, &wg, ksCmd)
		go GetTickerPrice(key, secret, "SPY", timeVal, "history", ch4, &wg, ksCmd)

		wg.Wait()

		spyHist := float64(<-ch4)
		spyLatest := float64(<-ch3)
		tickerHist := float64(<-ch2)
		tickerLatest := float64(<-ch1)

		spyPerf := ((spyLatest - spyHist) / spyHist) * 100
		if spyPerf > 0 {
			spyPositive = "+"
		} else {
			spyPositive = ""
		}

		tickerPerf := ((tickerLatest - tickerHist) / tickerHist) * 100
		if tickerPerf > 0 {
			tickerPositive = "+"
		} else {
			tickerPositive = ""
		}

		deltaPerf := tickerPerf - spyPerf
		if deltaPerf > 0 {
			deltaPositive = "+"
		} else {
			deltaPositive = ""
		}

		if spyPositive == "+" {
			spyC := color.New(color.FgGreen)
			spyTextC := color.YellowString("SPY")
			fmt.Println("")
			fmt.Printf("%v %v performance: ", spyTextC, timeVal)
			spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
		} else {
			spyC := color.New(color.FgRed)
			spyTextC := color.YellowString("SPY")
			fmt.Println("")
			fmt.Printf("%v %v performance: ", spyTextC, timeVal)
			spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
		}

		if tickerPositive == "+" {
			tickerC := color.New(color.FgGreen)
			fmt.Printf("%v %v performance: ", color.YellowString(strings.ToUpper(ticker)), timeVal)
			tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
		} else {
			tickerC := color.New(color.FgRed)
			fmt.Printf("%v %v performance: ", color.YellowString(strings.ToUpper(ticker)), timeVal)
			tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
		}

		if deltaPositive == "+" {
			deltaC := color.New(color.FgGreen)
			fmt.Printf("%v %v performance vs SPY: ", color.YellowString(strings.ToUpper(ticker)), timeVal)
			deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
			fmt.Println("")
		} else {
			deltaC := color.New(color.FgRed)
			fmt.Printf("%v %v performance vs SPY: ", color.YellowString(strings.ToUpper(ticker)), timeVal)
			deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
			fmt.Println("")
		}
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
