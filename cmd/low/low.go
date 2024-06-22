package low

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

func GetLow(key string, secret string, ticker string, timeVal string, cmdArgs string) {
	var (
		lowestVal  float64
		lowestDate string
	)

	t := utils.TimeAssignVals{
		TimeVal: timeVal,
		Ticker:  ticker,
		Cmd:     "low",
		UrlType: "",
	}

	u := utils.AssignTime(t)
	url := utils.AssignUrl(t, u)
	body, _ := utils.GetRequest(key, secret, url)

	i := 0
	lowestVal = 0.0
	for i < u.Iterator {
		arrayVal := fmt.Sprintf("[%v]", i)
		nextLowPrice, err := jsonparser.GetFloat(body, "bars", arrayVal, "l")
		if err != nil {
			panic(err)
		}
		nextLowDate, err := jsonparser.GetString(body, "bars", arrayVal, "t")
		if err != nil {
			panic(err)
		} else if lowestVal == 0.0 {
			lowestVal = nextLowPrice
			lowestDate = nextLowDate
		} else {
			if lowestVal > nextLowPrice {
				lowestVal = nextLowPrice
				lowestDate = nextLowDate
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

	priceDiff := (curPrice - lowestVal)
	percDiff := (priceDiff / lowestVal) * 100

	l := LowOutput{
		ticker:     ticker,
		timeVal:    timeVal,
		priceDiff:  priceDiff,
		lowestVal:  lowestVal,
		lowestDate: lowestDate,
		percDiff:   percDiff,
		cmdArgs:    cmdArgs,
	}

	formatOutputLow(l)
}

// LowCmd represents the low command
var LowCmd = &cobra.Command{
	Use:   "low",
	Short: "Returns a ticker's percentage and dollar increase from a recent low",
	Example: "  ks low aapl \n" +
		"  ks low aapl -t=6M \n",
	Run: func(cmd *cobra.Command, args []string) {
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
			fmt.Println("Defaulting to 1Y timeframe")
			timeVal = "1Y"
		} else if timeArg == "12M" {
			timeVal = "1Y"
		} else {
			timeVal = timeArg
		}

		GetLow(key, secret, ticker, timeVal, cmdArgs)
	},
}

func init() {
	LowCmd.Flags().StringP("time", "t", "", "A length of time for performance comparison")
}
