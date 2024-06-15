/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package sectors

import (
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

type GetSectorsConfig struct {
	Key    string
	Secret string
	Cmd    string
}

var (
	timeVal  string
	stockVal string
)

// sectorsCmd represents the sectors command
var SectorsCmd = &cobra.Command{
	Use:   "sectors",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Example: "  ks sectors \n" +
		"  ks sectors -t=1M \n" +
		"  ks sectors -t=1M -s=aapl \n",
	Run: func(cmd *cobra.Command, args []string) {
		ksCmd := "sectors"
		timeArg, _ := cmd.Flags().GetString("time")
		if timeArg == "" {
			timeVal = "YTD"
		} else {
			timeVal = timeArg
		}

		stockArg, _ := cmd.Flags().GetString("stock")
		if stockArg == "" {
			stockVal = "NA"
		} else {
			utils.TickerValidation(stockArg)
			stockVal = stockArg
		}

		_, key, secret := config.Init()
		cfg := GetSectorsConfig{
			Key:    key,
			Secret: secret,
			Cmd:    ksCmd,
		}

		CompareSectors(cfg, timeVal, stockVal)
	},
}

func init() {

	// Here you will define your flags and configuration settings.
	SectorsCmd.Flags().StringP("time", "t", "", "A length of time for performance comparison")
	SectorsCmd.Flags().StringP("stock", "s", "", "A stock to compare to all the sectors")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sectorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sectorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
