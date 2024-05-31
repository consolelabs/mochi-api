package adapter

import (
	"errors"
	"net/http"
	"time"
)

func do3(req *http.Request, apiKey string) (*http.Response, error) {
	const retries = 3
	return do(req, apiKey, retries)
}

func do(req *http.Request, apiKey string, retry int) (*http.Response, error) {
	if retry == 0 {
		return nil, errors.New("retry limit exceeded")
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("X-DUNE-API-KEY", apiKey)
	req.Header.Add("Content-Type", "application/json")

	// http request
	client := &http.Client{Timeout: 2 * time.Minute}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 429 {
		time.Sleep(5 * time.Second)
		return do(req, apiKey, retry-1)
	}

	return resp, nil
}
