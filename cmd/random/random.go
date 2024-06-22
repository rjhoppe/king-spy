package random

import (
	"sync"

	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

var (
	spyPositive    string
	tickerPositive string
	deltaPositive  string
	timeVal        string
)

var RandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Compare the performance of a random equity against the S&P 500",
	Long: `The random cmd pulls in a random ticker from the data/tickers.go file
and then compares it to the S&P 500 (SPY). This cmd can take in a time variable
with the "-t" flag.`,
	Example: "  ks random \n" +
		"  ks random -t=1Y",
	Run: func(cmd *cobra.Command, args []string) {
		ksCmd := "random"
		timeArg, _ := cmd.Flags().GetString("time")
		randomTick, randomTickName := utils.SelectRandomTicker()
		if timeArg == "" {
			timeVal = "YTD"
		} else {
			timeVal = timeArg
		}

		wg := sync.WaitGroup{}
		_, key, secret := config.Init()

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

		go utils.GetTickPrice(cfg, randomTick, timeVal, "latest", ch1)
		go utils.GetTickPrice(cfg, randomTick, timeVal, "history", ch2)
		go utils.GetTickPrice(cfg, "SPY", timeVal, "latest", ch3)
		go utils.GetTickPrice(cfg, "SPY", timeVal, "history", ch4)

		wg.Wait()

		spyHist := float64(<-ch4)
		spyLatest := float64(<-ch3)
		tickerHist := float64(<-ch2)
		tickerLatest := float64(<-ch1)

		r := RandomOutput{
			spyHist:      spyHist,
			spyLatest:    spyLatest,
			tickerHist:   tickerHist,
			tickerLatest: tickerLatest,
			tickerVal:    randomTick,
			tickerName:   randomTickName,
			timeVal:      timeVal,
		}

		formatOutputRandom(r)
	},
}

func init() {
	RandomCmd.PersistentFlags().String("time", "", "A length of time for performance comparison")
}
