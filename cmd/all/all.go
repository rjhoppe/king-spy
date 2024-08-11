package all

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/rjhoppe/king-spy/cmd/chart"
	"github.com/rjhoppe/king-spy/cmd/high"
	"github.com/rjhoppe/king-spy/cmd/low"
	"github.com/rjhoppe/king-spy/cmd/news"
	"github.com/rjhoppe/king-spy/cmd/sectors"
	"github.com/rjhoppe/king-spy/config"
	"github.com/rjhoppe/king-spy/utils"
	"github.com/spf13/cobra"
)

var (
	spyPositive    string
	tickerPositive string
	deltaPositive  string
)

// AllCmd represents the all command
var AllCmd = &cobra.Command{
	Use:   "all",
	Short: "Runs the c2s, high, low, sectors, and news cmds for a single ticker",
	Long:  `This command packages cmds c2s, high, low, and news together for a single ticker and returns the results`,
	Example: "  king-spy all aapl \n" +
		"  king-spy all aapl -c",
	Run: func(cmd *cobra.Command, args []string) {
		ksCmd := "all"
		ticker := args[0]
		utils.TickerValidation(ticker)
		ticker = strings.ToLower(ticker)
		cmdArgs := os.Args[1]
		chartFlag, _ := cmd.Flags().GetBool("chart")
		_, key, secret := config.Init()

		wg1 := sync.WaitGroup{}

		ch1 := make(chan float64)
		ch2 := make(chan float64)
		ch3 := make(chan float64)
		ch4 := make(chan float64)
		ch5 := make(chan float64)
		ch6 := make(chan float64)

		cfg := utils.GetTickerPriceConfig{
			Key:    key,
			Secret: secret,
			Wg:     &wg1,
			Cmd:    ksCmd,
		}

		go utils.GetTickPrice(cfg, ticker, "NA", "latest", ch1)
		go utils.GetTickPrice(cfg, ticker, "1M", "history", ch2)
		go utils.GetTickPrice(cfg, ticker, "6M", "history", ch3)
		go utils.GetTickPrice(cfg, ticker, "1Y", "history", ch4)
		go utils.GetTickPrice(cfg, "SPY", "NA", "latest", ch5)
		go utils.GetTickPrice(cfg, "SPY", "1Y", "history", ch6)

		wg1.Wait()

		tickerLatest := float64(<-ch1)
		tickerHist1M := float64(<-ch2)
		tickerHist6M := float64(<-ch3)
		tickerHist1Y := float64(<-ch4)
		spyLatest := float64(<-ch5)
		spyHist := float64(<-ch6)

		wg2 := sync.WaitGroup{}

		ch7 := make(chan float64)
		ch8 := make(chan float64)
		ch9 := make(chan float64)
		ch10 := make(chan float64)

		go GetPerf(tickerLatest, tickerHist1M, ch7, &wg2)
		go GetPerf(tickerLatest, tickerHist6M, ch8, &wg2)
		go GetPerf(tickerLatest, tickerHist1Y, ch9, &wg2)
		go GetPerf(spyLatest, spyHist, ch10, &wg2)

		wg2.Wait()

		ticker1MPerf := float64(<-ch7)
		ticker6MPerf := float64(<-ch8)
		ticker1YPerf := float64(<-ch9)
		spyPerf := float64(<-ch10)

		fmt.Println("")
		fmt.Printf("%v %v \n", color.YellowString(strings.ToUpper(ticker)), color.YellowString("Performance Overview"))
		fmt.Println("==================================================================================")
		fmt.Printf("%v latest price:   $%v \n", color.YellowString(strings.ToUpper(ticker)), tickerLatest)
		returnTickerPerf(ticker1MPerf, ticker, "1M")
		returnTickerPerf(ticker6MPerf, ticker, "6M")
		returnTickerPerf(ticker1YPerf, ticker, "1Y")

		inputs := AllInputs{
			spyPerf:      spyPerf,
			ticker1YPerf: ticker1YPerf,
			ticker:       ticker,
		}

		sectorCfg := sectors.GetSectorsConfig{
			Key:    key,
			Secret: secret,
			Cmd:    "sectors",
		}

		FormatAll(inputs)
		low.GetLow(key, secret, ticker, "1Y", cmdArgs)
		high.GetHigh(key, secret, ticker, "1Y", cmdArgs)
		news.GetNews(key, secret, ticker, cmdArgs)
		sectors.CompareSectors(sectorCfg, "1Y", ticker)
		fmt.Println("")

		if chartFlag {
			chart.LaunchChart(ticker)
		}
	},
}

func init() {
	AllCmd.Flags().BoolP("chart", "c", false, "Tells the program that you want it to open a chart of ticker in default browser")
}

func GetPerf(tickerLatest float64, tickerHistPrice float64, ch chan float64, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	tickVal := ((tickerLatest - tickerHistPrice) / tickerHistPrice) * 100
	ch <- tickVal
}

func returnTickerPerf(tickerPerf float64, ticker string, timeVal string) {
	if tickerPerf > 0 {
		tickerPositive = "+"
	} else {
		tickerPositive = ""
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
}
