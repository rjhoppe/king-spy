/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package high

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/fatih/color"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

var timeVal string

func GetHigh(key string, secret string, ticker string, timeVal string, cmdArgs string) {
	var (
		startTime string
		endTime   string
		timeframe string
		iterator  int
	)

	curTime := time.Now()
	switch timeVal {
	case "1M":
		pastTimeVal := curTime.AddDate(0, -1, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		endTime = curTime.Format(time.RFC3339)
		timeframe = "1D"
		iterator = 10
	case "3M":
		pastTimeVal := curTime.AddDate(0, -3, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		endTime = curTime.Format(time.RFC3339)
		timeframe = "1W"
		iterator = 10
	case "6M":
		pastTimeVal := curTime.AddDate(0, 6, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		endTime = curTime.Format(time.RFC3339)
		timeframe = "1W"
		iterator = 22
	case "1Y":
		pastTimeVal := curTime.AddDate(-1, 0, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		endTime = curTime.Format(time.RFC3339)
		timeframe = "1M"
		iterator = 11
	default:
		pastTimeVal := curTime.AddDate(-1, 0, 0)
		startTime = pastTimeVal.Format(time.RFC3339)
		endTime = curTime.Format(time.RFC3339)
		timeframe = "1M"
		timeVal = "1Y"
		iterator = 11
	}

	url := "https://data.alpaca.markets/v2/stocks/" + strings.ToUpper(ticker) + "/bars?timeframe=" + timeframe + "&start=" + startTime + "&end=" + endTime + "&limit=11&adjustment=raw&feed=iex&sort=asc"

	body, _ := utils.GetRequest(key, secret, url)
	var highestVal float64
	var highestDate string

	i := 0
	highestVal = 0.0
	for i < iterator {
		arrayVal := fmt.Sprintf("[%v]", i)
		nextHighPrice, err := jsonparser.GetFloat(body, "bars", arrayVal, "h")
		if err != nil {
			panic(err)
		}
		nextHighDate, err := jsonparser.GetString(body, "bars", arrayVal, "t")
		if err != nil {
			panic(err)
		} else if highestVal == 0.0 {
			highestVal = nextHighPrice
			highestDate = nextHighDate
		} else {
			if highestVal < nextHighPrice {
				highestVal = nextHighPrice
				highestDate = nextHighDate
			}
			i++
		}
	}

	curPriceUrl := "https://data.alpaca.markets/v2/stocks/" + ticker + "/trades/latest?feed=iex"
	curPriceBody, _ := utils.GetRequest(key, secret, curPriceUrl)
	curPrice, err := jsonparser.GetFloat(curPriceBody, "trade", "p")
	if err != nil {
		panic(err)
	}

	priceDiff := (highestVal - curPrice)
	percDiff := (priceDiff / highestVal) * 100
	// priceColor := color.New(color.FgRed)

	if cmdArgs == "high" {
		fmt.Println("")
	}

	fmt.Printf("The highest price of %v in the last %v time period was: %v on %v \n", color.YellowString(strings.ToUpper(ticker)), timeVal, color.GreenString("$"+strconv.FormatFloat(highestVal, 'f', 2, 64)), highestDate[:10])
	fmt.Printf("Price decrease off %v high: %v which is a %v decrease. \n", timeVal, color.RedString("-$"+strconv.FormatFloat(priceDiff, 'f', 2, 64)), color.RedString(strconv.FormatFloat(percDiff, 'f', 2, 64)+"%"))
	fmt.Println("")
}

// highCmd represents the high command
var HighCmd = &cobra.Command{
	Use:   "high",
	Short: "Returns a ticker's percentage and dollar decrease from a recent high",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := args[0]
		utils.TickerValidation(ticker)
		ticker = strings.ToLower(ticker)
		cmdArgs := os.Args[1]
		_, key, secret, _ := config.Init()

		timeArg, _ := cmd.Flags().GetString("time")

		if timeArg != "1Y" && timeArg != "6M" && timeArg != "3M" && timeArg != "1M" {
			fmt.Println("")
			flagVal := color.YellowString("--time={timeframe}")
			fmt.Printf("Timeframe not recognized or not provided. Use the %v flag to provide a timeframe. \n", flagVal)
			fmt.Println("The recognized timeframes are: 3Y, 1Y, 6M, 3M, 1M")
			fmt.Println("Defaulting to 1Y timeframe.")
		} else if timeArg == "12M" {
			timeVal = "1Y"
		} else {
			timeVal = timeArg
		}

		GetHigh(key, secret, ticker, timeVal, cmdArgs)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	HighCmd.Flags().StringP("time", "t", "", "A length of time for performance comparison")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// highCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// highCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
