/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package random

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/rjhoppe/go-compare-to-spy/cmd/c2s"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

var spyPositive string
var tickerPositive string
var deltaPositive string
var timeVal string

// randomCmd represents the random command
var RandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Compare the performance of a random equity against the S&P 500",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		timeArg, _ := cmd.Flags().GetString("time")
		randomTick, randomTickName := utils.SelectRandomTicker()

		if timeArg == "" {
			timeVal = "YTD"
		} else {
			timeVal = timeArg
		}

		wg := new(sync.WaitGroup)
		_, key, secret, _ := config.Init()

		ch1 := make(chan float64)
		ch2 := make(chan float64)
		ch3 := make(chan float64)
		ch4 := make(chan float64)

		go c2s.GetTickerPrice(key, secret, randomTick, timeVal, "latest", ch1, wg)
		go c2s.GetTickerPrice(key, secret, randomTick, timeVal, "history", ch2, wg)
		go c2s.GetTickerPrice(key, secret, "SPY", timeVal, "latest", ch3, wg)
		go c2s.GetTickerPrice(key, secret, "SPY", timeVal, "history", ch4, wg)

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
			fmt.Println("")
			spyValC := color.New(color.FgGreen)
			spyTextC := color.YellowString("SPY")
			fmt.Printf("%v: %v performance: ", spyTextC, timeVal)
			spyValC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
		} else {
			fmt.Println("")
			spyValC := color.New(color.FgRed)
			spyTextC := color.YellowString("SPY")
			fmt.Printf("%v: %v performance: ", spyTextC, timeVal)
			spyValC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
		}

		if tickerPositive == "+" {
			tickerValC := color.New(color.FgGreen)
			fmt.Printf("%v %v: %v performance: ", color.YellowString(randomTickName), color.YellowString("("+strings.ToUpper(randomTick)+")"), timeVal)
			tickerValC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
		} else {
			tickerValC := color.New(color.FgRed)
			fmt.Printf("%v %v: %v performance: ", color.YellowString(randomTickName), color.YellowString("("+strings.ToUpper(randomTick)+")"), timeVal)
			tickerValC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
		}

		if deltaPositive == "+" {
			deltaC := color.New(color.FgGreen)
			fmt.Printf("%v %v: %v performance vs SPY: ", color.YellowString(randomTickName), color.YellowString("("+strings.ToUpper(randomTick)+")"), timeVal)
			deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
			fmt.Println("")
		} else {
			deltaC := color.New(color.FgRed)
			fmt.Printf("%v %v: %v performance vs SPY: ", color.YellowString(randomTickName), color.YellowString("("+strings.ToUpper(randomTick)+")"), timeVal)
			deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
			fmt.Println("")
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	RandomCmd.PersistentFlags().String("time", "", "A length of time for performance comparison")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
