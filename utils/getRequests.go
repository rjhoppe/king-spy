package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetRequest(key string, secret string, url string) (body []byte) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", key)
	req.Header.Add("APCA-API-SECRET-KEY", secret)

	res, err := http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		fmt.Printf("Error: Could not get data. %v", err)
		return
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: Could not read response body. %v", err)
	}

	return respBody
}
