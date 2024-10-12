package vix

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vixCmd represents the vix command
var VixCmd = &cobra.Command{
	Use:   "vix",
	Short: "Calculates the volatility of a stock for the past month",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running vix cmd...")
		// ksCmd := "vix"
		// ticker := args[0]
		// utils.TickerValidation(ticker)
		// ticker = strings.ToLower(ticker)
		// timeArg, _ := cmd.Flags().GetString("time")
		// if timeArg == "" {
		// 	timeVal = "1M"
		// } else {
		// 	timeVal = timeArg
		// }

		// _, key, secret := config.Init()
		// wg := sync.WaitGroup{}

		// grab 1 time value high and time value low
		// compare high to low and compare
		// https://corporatefinanceinstitute.com/resources/career-map/sell-side/capital-markets/volatility-vol/
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
