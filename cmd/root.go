/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/rjhoppe/go-compare-to-spy/cmd/c2s"
	"github.com/rjhoppe/go-compare-to-spy/cmd/h2l"
	"github.com/rjhoppe/go-compare-to-spy/cmd/news"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-compare-to-spy",
	Short: "Compares a stock ticker's performance to SPY over a period of time",
	Long: `Can also perform other actions such as return ticker info and calculate recent lows to
	current price and recent hights to the current, lower price`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandPalettes() {
	rootCmd.AddCommand(c2s.CompareSpyCmd)
	rootCmd.AddCommand(news.NewsCmd)
	rootCmd.AddCommand(h2l.High2LowCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-compare-to-spy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubcommandPalettes()
}


