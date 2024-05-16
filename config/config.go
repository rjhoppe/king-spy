package config

import (
	"log"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/spf13/viper"
)

func Init() (*alpaca.Account, string, string) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
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
