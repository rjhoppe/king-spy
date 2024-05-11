package utils_test

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"

// 	"github.com/rjhoppe/go-compare-to-spy/config"
// 	"github.com/rjhoppe/go-compare-to-spy/utils"
// )

// func TestGetRequestsTypeCheck(t *testing.T) {
// 	_, key, secret := config.Init()
// 	url := "https://data.alpaca.markets/v2/stocks/AAPL/trades/latest?feed=iex"
// 	body := utils.GetRequest(key, secret, url)
// 	bodyType := reflect.TypeOf(body).Kind()
// 	if reflect.TypeOf(body) != reflect. {
// 		log.Error("GetRequest func returned type other than []byte")
// 	}
// }

// func TestGetRequestsBadTickerType(t *testing.T) {
// 	_, key, secret := config.Init()
// 	url := "https://data.alpaca.markets/v2/stocks/4830/trades/latest?feed=iex"
// 	body := utils.GetRequest(key, secret, url)
// }

// func TestGetRequestsBadTicker(t *testing.T) {
// 	_, key, secret := config.Init()
// 	url := "https://data.alpaca.markets/v2/stocks/AAXMP/trades/latest?feed=iex"
// 	body := utils.GetRequest(key, secret, url)
// }
