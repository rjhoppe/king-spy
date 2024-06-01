/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package wsb

import (
	"encoding/json"
	"fmt"
	"io"
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
	Use:   "wsb",
	Short: "Returns the top tickers mentioned on WSB and the related sentiment",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			stocks    []WSBData
			sentColor string
		)

		resp, err := http.Get("https://tradestie.com/api/v1/apps/reddit")
		if err != nil {
			// log.Fatal(err)
			panic(err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			// log.Fatal(err)
			panic(err)
		}

		err = json.Unmarshal(bodyBytes, &stocks)
		if err != nil {
			panic(err)
		}

		fmt.Println("")
		fmt.Printf("%v \n", color.YellowString("WSB Leaderboard"))
		fmt.Println("=================================")
		for i := range stocks[:10] {
			if stocks[i].Sentiment == "Bullish" {
				sentColor = color.GreenString(stocks[i].Sentiment)
			} else {
				sentColor = color.RedString(stocks[i].Sentiment)
			}
			spacing := formatOutput(stocks[i].Ticker)
			fmt.Println(stocks[i].Ticker + spacing + sentColor + " - Comments(" + strconv.Itoa(stocks[i].No_Of_Comments) + ")")
		}
		fmt.Println("=================================")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wsbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wsbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
