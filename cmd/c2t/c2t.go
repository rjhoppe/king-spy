package c2t

import (
	"strings"
	"sync"

	"github.com/rjhoppe/king-spy/config"
	"github.com/rjhoppe/king-spy/utils"
	"github.com/spf13/cobra"
)

var (
	ticker1Positive string
	ticker2Positive string
	deltaPositive   string
	timeVal         string
)

// Compare2TickerCmd represents the c2t command
var Compare2TickerCmd = &cobra.Command{
	Use:   "c2t",
	Short: "Compares one ticker's performance to another ticker over a specified time period",
	Example: "  king-spy c2t nvda amd \n" +
		"  king-spy c2t nvda amd -t=1M \n" +
		"  king-spy c2t nvda amd -t=1Y \n",
	Run: func(cmd *cobra.Command, args []string) {
		ksCmd := "c2t"
		ticker1 := args[0]
		ticker2 := args[1]
		utils.TickerValidation(ticker1)
		utils.TickerValidation(ticker2)
		ticker1 = strings.ToLower(ticker1)
		ticker2 = strings.ToLower(ticker2)
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

		go utils.GetTickPrice(cfg, ticker1, timeVal, "latest", ch1)
		go utils.GetTickPrice(cfg, ticker1, timeVal, "history", ch2)
		go utils.GetTickPrice(cfg, ticker2, timeVal, "latest", ch3)
		go utils.GetTickPrice(cfg, ticker2, timeVal, "history", ch4)

		wg.Wait()

		ticker1Latest := float64(<-ch1)
		ticker1Hist := float64(<-ch2)
		ticker2Latest := float64(<-ch3)
		ticker2Hist := float64(<-ch4)

		to := TickersOutput{
			ticker1Hist:   ticker1Hist,
			ticker1Latest: ticker1Latest,
			ticker2Hist:   ticker2Hist,
			ticker2Latest: ticker2Latest,
			ticker1Val:    ticker1,
			ticker2Val:    ticker2,
			timeVal:       timeVal,
		}

		formatOutputC2T(to)
	},
}

func init() {

	Compare2TickerCmd.Flags().StringP("time", "t", "", "A length of time for performance comparison")
}
