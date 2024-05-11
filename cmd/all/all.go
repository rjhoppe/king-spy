/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package all

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/rjhoppe/go-compare-to-spy/cmd/c2s"
	"github.com/rjhoppe/go-compare-to-spy/cmd/chart"
	"github.com/rjhoppe/go-compare-to-spy/cmd/high"
	"github.com/rjhoppe/go-compare-to-spy/cmd/low"
	"github.com/rjhoppe/go-compare-to-spy/cmd/news"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/spf13/cobra"
)

var spyPositive string
var tickerPositive string
var deltaPositive string

func GetPerf(tickerLatest float64, tickerHistPrice float64, ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	tickVal := ((tickerLatest - tickerHistPrice) / tickerHistPrice) * 100
	ch <- tickVal
}

func returnTickerPerf(tickerPerf float64, ticker string, timeVal string) {
	if (tickerPerf > 0) {
		tickerPositive = "+"
	} else {
		tickerPositive = ""
	}

	if (tickerPositive == "+") {
		tickerC := color.New(color.FgGreen)
		fmt.Printf("%v %v performance: ", color.YellowString(strings.ToUpper(ticker)), timeVal)
		tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
	} else {
		tickerC := color.New(color.FgRed)
		fmt.Printf("%v %v performance: ", color.YellowString(strings.ToUpper(ticker)), timeVal)
		tickerC.Printf("%v%.2f%% \n", tickerPositive, tickerPerf)
	}
}

// allCmd represents the all command
var AllCmd = &cobra.Command{
	Use:   "all",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := strings.ToLower(args[0])
		cmdArgs := os.Args[1]
		chartFlag, _ := cmd.Flags().GetBool("chart")
		_, key, secret := config.Init()

		wg1 := new(sync.WaitGroup)

		ch1 := make(chan float64)
		ch2 := make(chan float64)
		ch3 := make(chan float64)
		ch4 := make(chan float64)
		ch5 := make(chan float64)
		ch6 := make(chan float64)

		wg1.Add(8)

		go c2s.GetTickerPrice(key, secret, ticker, "NA", "latest", ch1, wg1)
		go c2s.GetTickerPrice(key, secret, ticker, "1M", "history", ch2, wg1)
		go c2s.GetTickerPrice(key, secret, ticker, "6M", "history", ch3, wg1)
		go c2s.GetTickerPrice(key, secret, ticker, "1Y", "history", ch4, wg1)
		go c2s.GetTickerPrice(key, secret, "SPY", "NA", "latest", ch5, wg1)
		go c2s.GetTickerPrice(key, secret, "SPY", "1Y", "history", ch6, wg1)

		tickerLatest := float64(<-ch1)
		tickerHist1M := float64(<-ch2)
		tickerHist6M := float64(<-ch3)
		tickerHist1Y := float64(<-ch4)
		spyLatest := float64(<-ch5)
		spyHist := float64(<-ch6)
		
		wg2 := new(sync.WaitGroup)

		ch7 := make(chan float64)
		ch8 := make(chan float64)
		ch9 := make(chan float64)
		ch10 := make(chan float64)

		wg2.Add(4)

		go GetPerf(tickerLatest, tickerHist1M, ch7, wg2)
		go GetPerf(tickerLatest, tickerHist6M, ch8, wg2)
		go GetPerf(tickerLatest, tickerHist1Y, ch9, wg2)
		go GetPerf(spyLatest, spyHist, ch10, wg2)
		
		ticker1MPerf := float64(<-ch7)
		ticker6MPerf := float64(<-ch8)
		ticker1YPerf := float64(<-ch9)
		spyPerf := float64(<-ch10)

		fmt.Println("")
		fmt.Printf("%v latest price: $%v \n", color.YellowString(strings.ToUpper(ticker)), tickerLatest)
		returnTickerPerf(ticker1MPerf, ticker, "1M")
		returnTickerPerf(ticker6MPerf, ticker, "6M")
		returnTickerPerf(ticker1YPerf, ticker, "1Y")

		if (spyPerf > 0) {
			spyPositive = "+"
		} else {
			spyPositive = ""
		}

		deltaPerf := ticker1YPerf - spyPerf
		if (deltaPerf > 0) {
			deltaPositive = "+"
		} else {
			deltaPositive = ""
		}

		if (spyPositive == "+") {
			spyC := color.New(color.FgGreen)
			spyTextC := color.YellowString("SPY")
			fmt.Println("")
			fmt.Printf("%v 1Y performance: ", spyTextC)
			spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
		} else {
			spyC := color.New(color.FgRed)
			spyTextC := color.YellowString("SPY")
			fmt.Println("")
			fmt.Printf("%v 1Y performance: ", spyTextC)
			spyC.Printf("%v%.2f%% \n", spyPositive, spyPerf)
		}

		if (deltaPositive == "+") {
			deltaC := color.New(color.FgGreen)
			fmt.Printf("%v 1Y performance vs SPY: ", color.YellowString(strings.ToUpper(ticker)))
			deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
		} else {
			deltaC := color.New(color.FgRed)
			fmt.Printf("%v 1Y performance vs SPY: ", color.YellowString(strings.ToUpper(ticker)))
			deltaC.Printf("%v%.2f%% \n", deltaPositive, deltaPerf)
		}

		low.GetLow(key, secret, ticker, "1Y", cmdArgs)
		high.GetHigh(key, secret, ticker, "1Y", cmdArgs)
		news.GetNews(key, secret, ticker, cmdArgs)

		if chartFlag {
			chart.LaunchChart(ticker)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Add flag for when you want chart to display
	AllCmd.Flags().BoolP("chart", "c", false, "Tell program if you want it to open a chart of ticker in default browser")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// allCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// allCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
