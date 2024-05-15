package utils

import (
	"errors"
	"io"
	"net/http"
)

func GetRequest(key string, secret string, url string) (body []byte, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("APCA-API-KEY-ID", key)
	req.Header.Add("APCA-API-SECRET-KEY", secret)
	res, err := http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, errors.New("error: unknown error occurred")
	}

	if res.StatusCode == 404 {
		return nil, errors.New("error: invalid ticker - ticker not found")
	}

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("error: reading response body")
	}

	return respBody, nil
}
