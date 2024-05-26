package utils

import "strings"

func AssignUrl(t TimeAssignVals, u UrlAssignVals) (url string) {
	switch t.Cmd {
	case "all", "c2s", "random":
		if t.UrlType == "history" {
			url = "https://data.alpaca.markets/v2/stocks/" + t.Ticker + "/trades?limit=1&start=" + u.StartTime + "&end=" + u.EndTime + "&feed=iex&currency=USD"
		} else {
			url = "https://data.alpaca.markets/v2/stocks/" + t.Ticker + "/trades/latest?feed=iex"
		}
	case "high":
		url = "https://data.alpaca.markets/v2/stocks/" + strings.ToUpper(t.Ticker) + "/bars?timeframe=" + u.Timeframe + "&start=" + u.StartTime + "&end=" + u.EndTime + "&adjustment=raw&feed=iex&sort=asc"
	case "low":
		url = "https://data.alpaca.markets/v2/stocks/" + strings.ToUpper(t.Ticker) + "/bars?timeframe=" + u.Timeframe + "&start=" + u.StartTime + "&end=" + u.EndTime + "&adjustment=raw&feed=iex&sort=asc"
	}

	return url
}
