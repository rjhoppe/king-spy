/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/fatih/color"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/rjhoppe/go-compare-to-spy/cmd/all"
	"github.com/rjhoppe/go-compare-to-spy/cmd/c2s"
	"github.com/rjhoppe/go-compare-to-spy/cmd/chart"
	"github.com/rjhoppe/go-compare-to-spy/cmd/high"
	"github.com/rjhoppe/go-compare-to-spy/cmd/low"
	"github.com/rjhoppe/go-compare-to-spy/cmd/news"
	"github.com/rjhoppe/go-compare-to-spy/cmd/random"
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
	// Example: "example [sub command]",
	Example: "  ks all aapl -t=1M \n" +
		"  ks c2s aapl -t=1Y \n" +
		"  ks chart aapl \n" +
		"  ks high aapl \n" +
		"  ks low aapl \n" +
		"  ks news aapl \n" +
		"  ks random \n" +
		"  ks wsb",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  rootCmd,
		Headings: cc.HiGreen + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
	})

	if len(os.Args) < 2 {
		utils.AsciiTitleText()
	}
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(c2s.CompareSpyCmd)
	rootCmd.AddCommand(news.NewsCmd)
	rootCmd.AddCommand(low.LowCmd)
	rootCmd.AddCommand(high.HighCmd)
	rootCmd.AddCommand(random.RandomCmd)
	rootCmd.AddCommand(chart.ChartCmd)
	rootCmd.AddCommand(all.AllCmd)
	rootCmd.AddCommand(wsb.WsbCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-compare-to-spy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubcommandPalettes()
	// input := os.Stdin
	// fmt.Println(input)
	// utils.AsciiTitleText()
}
