package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/rjhoppe/king-spy/cmd/all"
	"github.com/rjhoppe/king-spy/cmd/buzz"
	"github.com/rjhoppe/king-spy/cmd/c2s"
	"github.com/rjhoppe/king-spy/cmd/c2t"
	"github.com/rjhoppe/king-spy/cmd/chart"
	"github.com/rjhoppe/king-spy/cmd/high"
	"github.com/rjhoppe/king-spy/cmd/low"
	"github.com/rjhoppe/king-spy/cmd/news"
	"github.com/rjhoppe/king-spy/cmd/random"
	"github.com/rjhoppe/king-spy/cmd/sectors"
	"github.com/rjhoppe/king-spy/cmd/vix"
	"github.com/rjhoppe/king-spy/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "king-spy",
	Short: "Compares a stock ticker's performance to SPY over a period of time",
	Long: "\n This CLI application compares the performance of " +
		"individual stocks or ETFs to the S&P 500. \n" +
		" This tool can help you explore which equities are currently outperforming the indexes. \n" +
		" However, in using this tool, you will find that most equities don't outperform the indexes. \n" +
		" At least not in the long haul! \n \n" + " " + color.YellowString("All hail, King SPY!"),
	Example: "  king-spy all aapl -t=1M \n" +
		"  king-spy c2s aapl -t=1Y \n" +
		"  king-spy c2t nvda amd -t=6M \n" +
		"  king-spy chart aapl \n" +
		"  king-spy high aapl \n" +
		"  king-spy low aapl \n" +
		"  king-spy news aapl \n" +
		"  king-spy random \n" +
		"  king-spy buzz \n" +
		"  king-spy sectors \n" +
		"  king-spy sectors -t=1Y -s=aapl" +
		"  king-spy vix aapl",
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
	rootCmd.AddCommand(buzz.BuzzCmd)
	rootCmd.AddCommand(sectors.SectorsCmd)
	rootCmd.AddCommand(vix.VixCmd)
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	addSubcommandPalettes()
}
