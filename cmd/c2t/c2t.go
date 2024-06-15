/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package c2t

import (
	"strings"
	"sync"

	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

var (
	ticker1Positive string
	ticker2Positive string
	deltaPositive   string
	timeVal         string
)

// c2tCmd represents the c2t command
var Compare2TickerCmd = &cobra.Command{
	Use:   "c2t",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c2tCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c2tCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}