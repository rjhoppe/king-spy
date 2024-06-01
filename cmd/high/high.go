/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package high

import (
	"fmt"
	"os"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/fatih/color"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

var timeVal string

func GetHigh(key string, secret string, ticker string, timeVal string, cmdArgs string) {
	var (
		highestVal  float64
		highestDate string
	)

	t := utils.TimeAssignVals{
		TimeVal: timeVal,
		Ticker:  ticker,
		Cmd:     "high",
		UrlType: "",
	}

	u := utils.AssignTime(t)
	url := utils.AssignUrl(t, u)
	body, _ := utils.GetRequest(key, secret, url)

	i := 0
	highestVal = 0.0
	for i < u.Iterator {
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

	// Refactor this to use gettickerprice func
	// is this really necessary?
	curPriceUrl := "https://data.alpaca.markets/v2/stocks/" + ticker + "/trades/latest?feed=iex"
	curPriceBody, _ := utils.GetRequest(key, secret, curPriceUrl)
	curPrice, err := jsonparser.GetFloat(curPriceBody, "trade", "p")
	if err != nil {
		panic(err)
	}

	priceDiff := (highestVal - curPrice)
	percDiff := (priceDiff / highestVal) * 100
	// priceColor := color.New(color.FgRed)

	h := HighOutput{
		ticker:      ticker,
		timeVal:     timeVal,
		priceDiff:   priceDiff,
		highestVal:  highestVal,
		highestDate: highestDate,
		percDiff:    percDiff,
		cmdArgs:     cmdArgs,
	}

	formatOutputHigh(h)
	// if cmdArgs == "high" {
	// 	fmt.Println("")
	// }
	// fmt.Println("==================================================================================")
	// fmt.Printf("The highest price of %v in the last %v time period was: %v on %v \n", color.YellowString(strings.ToUpper(ticker)), timeVal, color.GreenString("$"+strconv.FormatFloat(highestVal, 'f', 2, 64)), highestDate[:10])
	// fmt.Printf("Price decrease off %v high: %v which is a %v decrease. \n", timeVal, color.RedString("-$"+strconv.FormatFloat(priceDiff, 'f', 2, 64)), color.RedString(strconv.FormatFloat(percDiff, 'f', 2, 64)+"%"))
	// fmt.Println("==================================================================================")
}

// highCmd represents the high command
var HighCmd = &cobra.Command{
	Use:   "high",
	Short: "Returns a ticker's percentage and dollar decrease from a recent high",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// ksCmd := "high"
		ticker := args[0]
		utils.TickerValidation(ticker)
		ticker = strings.ToLower(ticker)
		cmdArgs := os.Args[1]
		_, key, secret := config.Init()

		timeArg, _ := cmd.Flags().GetString("time")

		if timeArg != "1Y" && timeArg != "6M" && timeArg != "3M" && timeArg != "1M" {
			fmt.Println("")
			flagVal := color.YellowString("--time={timeframe}")
			fmt.Printf("Timeframe not recognized or not provided. Use the %v flag to provide a timeframe. \n", flagVal)
			fmt.Println("The recognized timeframes are: 3Y, 1Y, 6M, 3M, 1M")
			fmt.Println("Defaulting to 1Y timeframe.")
			timeVal = "1Y"
		} else if timeArg == "12M" {
			timeVal = "1Y"
		} else {
			timeVal = timeArg
		}

		GetHigh(key, secret, ticker, timeVal, cmdArgs)
		// c2s.GetTickerPrice(key, secret, ticker, "NA", "latest", ch1, &wg1, ksCmd)
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
