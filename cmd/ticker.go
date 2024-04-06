package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var tickerCmd = &cobra.Command{
	Use:   "ticker",
	Short: "searches price info for a ticker",
	Long:  `This subcommand says hello`,
	Run: func(cmd *cobra.Command, ticker []string) {

		if len(ticker) > 4 {
			fmt.Errorf("%v: Invalid ticker. Try again.", ticker)
			return
		}

		fmt.Println("hello called")
	},
}

func init() {
	RootCmd.AddCommand(tickerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
