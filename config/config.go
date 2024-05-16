package config

import (
	"fmt"

	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/spf13/viper"
)

func Init() (*alpaca.Account, string, string, error) {
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
	// key := viper.GetString("envkey_APCA_API_KEY_ID")
	// secret := viper.GetString("envkey_APCA_API_SECRET_KEY")
	// endpoint := viper.GetString("envkey_ENDPOINT")

	client := alpaca.NewClient(alpaca.ClientOpts{
		APIKey:    key,
		APISecret: secret,
		BaseURL:   endpoint,
	})

	acct, err := client.GetAccount()
	if err != nil {
		return nil, "null", "null", fmt.Errorf("error: 401: secret %v key %v endpoint %v", secret, key, endpoint)
	}

	return acct, key, secret, nil
}
