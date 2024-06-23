package config

import (
	"fmt"
	"log"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/spf13/viper"
)

func Init() (*alpaca.Account, string, string) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	// this was a config file path for testing locally
	viper.AddConfigPath(".")
	// change this to reflect wherever your .env file with Alpaca creds is located
	viper.AddConfigPath("placeholder")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("")
		fmt.Println("Config file not found: Make sure full path to .env file is specified in a viper.AddConfigPath() call")
		log.Fatal(err)
	}

	key := viper.GetString("APCA_API_KEY_ID")
	secret := viper.GetString("APCA_API_SECRET_KEY")
	endpoint := viper.GetString("ENDPOINT")

	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    key,
		APISecret: secret,
		BaseURL:   endpoint,
	})

	acct, err := client.GetAccount()
	if err != nil {
		log.Fatal("error: could not return client credentials")
	}

	return acct, key, secret
}
