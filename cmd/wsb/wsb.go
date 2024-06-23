package wsb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type WSBData struct {
	No_Of_Comments  int
	Sentiment       string
	Sentiment_Score float64
	Ticker          string
}

// wsbCmd represents the wsb command
var WsbCmd = &cobra.Command{
	Use:     "wsb",
	Short:   "Returns the top tickers mentioned on r/wallstreetbets and the related sentiment for each",
	Example: "  king-spy wsb",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			stocks    []WSBData
			sentColor string
		)

		resp, err := http.Get("https://tradestie.com/api/v1/apps/reddit")
		if err != nil {
			log.Fatal(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(bodyBytes, &stocks)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("")
		fmt.Printf("%v \n", color.YellowString("WSB Leaderboard"))
		fmt.Println("==================================================================================")
		for i := range stocks[:10] {
			if stocks[i].Sentiment == "Bullish" {
				sentColor = color.GreenString(stocks[i].Sentiment)
			} else {
				sentColor = color.RedString(stocks[i].Sentiment)
			}
			spacing := formatOutput(stocks[i].Ticker)
			fmt.Println(stocks[i].Ticker + spacing + sentColor + " - Comments(" + strconv.Itoa(stocks[i].No_Of_Comments) + ")")
		}
		fmt.Println("==================================================================================")
	},
}
