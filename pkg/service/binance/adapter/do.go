package badapter

import (
	"net/http"
	"time"

	"github.com/k0kubun/pp/v3"
)

func do(req *http.Request, apiKey string, retry int) (*http.Response, error) {
	req.Header.Add("X-MBX-APIKEY", apiKey)
	req.Header.Add("Content-Type", "application/json")

	// http request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	pp.Println("check resp binance", resp.Header)

	if resp.StatusCode == 429 {
		time.Sleep(time.Duration(retry) * 5 * time.Second)
		return do(req, apiKey, retry+1)
	}

	return resp, nil
}
