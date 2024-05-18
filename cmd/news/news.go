/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package news

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/fatih/color"
	"github.com/rjhoppe/go-compare-to-spy/config"
	"github.com/rjhoppe/go-compare-to-spy/utils"
	"github.com/spf13/cobra"
)

// newsCmd represents the news command
var NewsCmd = &cobra.Command{
	Use:   "news",
	Short: "Get the most recent headlines for a specified ticker",
	Long:  `Returns the 5 most recent news headline for a supplied ticker`,
	Run: func(cmd *cobra.Command, args []string) {
		ticker := args[0]
		utils.TickerValidation(ticker)
		ticker = strings.ToLower(ticker)
		cmdArgs := os.Args[1]
		_, key, secret := config.Init()
		GetNews(key, secret, ticker, cmdArgs)
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

func GetNews(key string, secret string, ticker string, cmdArgs string) {
	url := "https://data.alpaca.markets/v1beta1/news?symbols=" + strings.ToUpper(ticker)
	body, _ := utils.GetRequest(key, secret, url)
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

	if cmdArgs == "news" {
		fmt.Println("")
	}

	for i := 1; i < len(headline); i++ {
		j := strconv.Itoa(i)
		fmt.Printf("%v %v: %v", color.YellowString("Headline"), color.YellowString(j), headline[i])
		fmt.Println("")
	}
	fmt.Println("")
}
