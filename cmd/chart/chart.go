package chart

import (
	"fmt"
	"strings"

	"github.com/pkg/browser"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

// ChartCmd represents the chart command
var ChartCmd = &cobra.Command{
	Use:     "chart",
	Short:   "Opens a one year chart for a ticker in your default browser",
	Example: "  ks chart aapl",
	Run: func(cmd *cobra.Command, args []string) {
		ticker := args[0]
		utils.TickerValidation(ticker)
		ticker = strings.ToLower(ticker)
		LaunchChart(ticker)
	},
}

func LaunchChart(ticker string) {
	classicUrl := "https://stockcharts.com/h-sc/ui?s="
	newUrl := "https://stockcharts.com/sc3/ui/?s="

	fmt.Println("Launching chart in the default browser...")
	err := browser.OpenURL(classicUrl + ticker)
	if err != nil {
		err := browser.OpenURL(newUrl + ticker)
		if err != nil {
			panic(err)
		}
	}
}
