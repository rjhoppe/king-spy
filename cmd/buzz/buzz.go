package buzz

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/exp/slices"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Result struct {
	Rank           int    `json:"rank"`
	Ticker         string `json:"ticker"`
	Name           string `json:"name"`
	Mentions       int    `json:"mentions"`
	Upvotes        int    `json:"upvotes"`
	Rank24hAgo     int    `json:"rank_24h_ago"`
	Mentions24hAgo int    `json:"mentions_24h_ago"`
}

type Response struct {
	Count       int      `json:"count"`
	Pages       int      `json:"pages"`
	CurrentPage int      `json:"current_page"`
	Results     []Result `json:"results"`
}

// BuzzCmd represents the trending command
var BuzzCmd = &cobra.Command{
	Use:     "buzz",
	Short:   "Returns the top tickers mentioned on r/wallstreetbets and other social media platforms",
	Example: "  king-spy buzz",
	Run: func(cmd *cobra.Command, args []string) {
		var data Response
		filter, _ := cmd.Flags().GetString("filter")
		filterVals := []string{"options", "stocks", "crypto", "all-stocks", "investing", "Daytrading"}
		if filter == "" {
			filter = "all-stocks"
		} else {
			if !slices.Contains(filterVals, filter) {
				fmt.Println("")
				fmt.Printf("Filter: %v is not valid \n", color.RedString(filter))
				fmt.Println("Please rerun command with a valid filters or with no filter parameter")
				fmt.Println("Valid filters: options: stocks, crypto, all-stocks, investing, daytrading")
				fmt.Println("")
				os.Exit(1)
			} else if filter == "crypto" {
				filter = "Bitcoin"
			}
		}

		url := "https://apewisdom.io/api/v1.0/filter/" + filter + "/page/1"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(bodyBytes, &data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("")
		fmt.Printf("%v \n", color.YellowString("Buzz Leaderboard"))
		fmt.Println("==================================================================================")
		for i, result := range data.Results[:10] {
			var emoji rune
			if i >= 10 {
				break
			} else {

				if result.Mentions24hAgo < result.Mentions {
					emoji = '\U0001F4C8'
				} else {
					emoji = '\U0001F4C9'
				}
				// fmt.Printf("%v [%v] %v: %v mentions, comment trend: %c \n", result.Rank, color.YellowString(result.Ticker), html.UnescapeString(result.Name), result.Mentions, emoji)
				fmt.Printf("%v %v [%v]: %v mentions, comment trend: %c \n", result.Rank, color.YellowString(html.UnescapeString(result.Name)), result.Ticker, result.Mentions, emoji)
			}
		}
		fmt.Println("==================================================================================")
	},
}

func init() {
	BuzzCmd.Flags().StringP("filter", "f", "", "A filtering parameter for trending data")
}
