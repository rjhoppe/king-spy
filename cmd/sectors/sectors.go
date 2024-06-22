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

// SectorsCmd represents the sectors command
var SectorsCmd = &cobra.Command{
	Use:   "sectors",
	Short: "Returns the performance of various sectors over a time period",
	Long: `The sectors cmd can also take the optional 't' flag which allows you
	to specify a time period to benchmark sector performance against. This cmd can
	also take the 's' stock flag which allows you to compare all the sectors against
	the performance of a particular equity.`,
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
	SectorsCmd.Flags().StringP("time", "t", "", "A length of time for performance comparison")
	SectorsCmd.Flags().StringP("stock", "s", "", "A stock to compare to all the sectors")
}
