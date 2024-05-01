/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/spf13/cobra"
)

func getNews(key string, secret string, ticker string) {
	url := "https://data.alpaca.markets/v1beta1/news?symbols=" + strings.ToUpper(ticker)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", key)
	req.Header.Add("APCA-API-SECRET-KEY", secret)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error: Could not get data. %v", err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: Could not read response body. %v", err)
	}

	headline := make([]string, 0)

	i := 0
	for i < 6 {
		arrayVal := fmt.Sprintf("[%v]", i)
		nextHeadline, err := jsonparser.GetString(body, "news", arrayVal, "headline")
		if err != nil {
			fmt.Printf("Error: Could not parse news source. %v", err)
			panic(err)
		} else {
			headline = append(headline, nextHeadline)
			i++
		}
	}

	for i := 1; i < len(headline); i++ {
		fmt.Println(headline[i])
	}
}

// newsCmd represents the news command
var NewsCmd = &cobra.Command{
	Use:   "news",
	Short: "Get the most recent news for a given ticker",
	Long: `Returns the most recent news headline for a supplied ticker`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := strings.ToLower(args[0])
		_, key, secret := config.Init()
		getNews(key, secret, ticker)
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
