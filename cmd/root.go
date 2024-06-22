package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/rjhoppe/go-compare-to-spy/cmd/all"
	"github.com/rjhoppe/go-compare-to-spy/cmd/c2s"
	"github.com/rjhoppe/go-compare-to-spy/cmd/c2t"
	"github.com/rjhoppe/go-compare-to-spy/cmd/chart"
	"github.com/rjhoppe/go-compare-to-spy/cmd/high"
	"github.com/rjhoppe/go-compare-to-spy/cmd/low"
	"github.com/rjhoppe/go-compare-to-spy/cmd/news"
	"github.com/rjhoppe/go-compare-to-spy/cmd/random"
	"github.com/rjhoppe/go-compare-to-spy/cmd/sectors"
	"github.com/rjhoppe/go-compare-to-spy/cmd/wsb"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ks",
	Short: "Compares a stock ticker's performance to SPY over a period of time",
	Long: "\n This CLI application compares the performance of " +
		"individual stocks or ETFs to the S&P 500. \n" +
		" This tool can help you explore which equities are currently outperforming the indexes. \n" +
		" However, in using this tool, you will find that most equities don't outperform the indexes. \n" +
		" At least not over the long haul! \n \n" + " " + color.YellowString("All hail, King SPY!"),
	Example: "  ks all aapl -t=1M \n" +
		"  ks c2s aapl -t=1Y \n" +
		"  ks c2t nvda amd -t=6M \n" +
		"  ks chart aapl \n" +
		"  ks high aapl \n" +
		"  ks low aapl \n" +
		"  ks news aapl \n" +
		"  ks random \n" +
		"  ks wsb \n" +
		"  ks sectors \n" +
		"  ks sectors -t=1Y -s=aapl",
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  rootCmd,
		Headings: cc.HiGreen + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
	})

	if len(os.Args) < 2 {
		fmt.Println("")
		utils.AsciiTitleText()
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(c2s.Compare2SpyCmd)
	rootCmd.AddCommand(c2t.Compare2TickerCmd)
	rootCmd.AddCommand(news.NewsCmd)
	rootCmd.AddCommand(low.LowCmd)
	rootCmd.AddCommand(high.HighCmd)
	rootCmd.AddCommand(random.RandomCmd)
	rootCmd.AddCommand(chart.ChartCmd)
	rootCmd.AddCommand(all.AllCmd)
	rootCmd.AddCommand(wsb.WsbCmd)
	rootCmd.AddCommand(sectors.SectorsCmd)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	addSubcommandPalettes()
}
