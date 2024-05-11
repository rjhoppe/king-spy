/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package chart

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

func LaunchChart(ticker string) {
	classicUrl := "https://stockcharts.com/h-sc/ui?s="
	newUrl := "https://stockcharts.com/sc3/ui/?s="

	err := browser.OpenURL(classicUrl + ticker)
	if err != nil {
		err := browser.OpenURL(newUrl + ticker)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Launching chart in the default browser...")
}

// chartCmd represents the chart command
var ChartCmd = &cobra.Command{
	Use:   "chart",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := args[0]
		utils.CheckTickerBadChars(ticker)
		ticker = strings.ToLower(ticker)
		LaunchChart(ticker)
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
