package apilayer

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/util"
)

type APILayer struct {
	headers          map[string]string
	getTimeSeriesURL string
	getLatestRateURL string
}

func NewService(cfg *config.Config) Service {
	return &APILayer{
		headers:          map[string]string{"apiKey": cfg.APILayerAPIKey},
		getTimeSeriesURL: "https://api.apilayer.com/exchangerates_data/timeseries?start_date=%s&end_date=%s&base=%s&symbols=%s",
		getLatestRateURL: "https://api.apilayer.com/exchangerates_data/latest?base=%s&symbols=%s",
	}
}

func (s *APILayer) GetHistoricalExchangeRate(q GetHistoricalExchangeRateQuery) (*GetHistoricalExchangeRateResponse, int, error) {
	data := &GetHistoricalExchangeRateResponse{}
	req := util.SendRequestQuery{
		URL:      fmt.Sprintf(s.getTimeSeriesURL, q.StartDate, q.EndDate, q.Base, q.Target),
		Response: data,
		Headers:  s.headers,
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, statusCode, fmt.Errorf("[apilayer.GetHistoricalExchangeRate] util.SendRequest() failed: %v", err)
	}

	return data, http.StatusOK, nil
}

func (s *APILayer) GetLatestExchangeRate(q GetLatestExchangeRateQuery) (*GetLatestExchangeRateResponse, int, error) {
	data := &GetLatestExchangeRateResponse{}
	req := util.SendRequestQuery{
		URL:      fmt.Sprintf(s.getLatestRateURL, q.Base, q.Target),
		Response: data,
		Headers:  s.headers,
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, statusCode, fmt.Errorf("[apilayer.GetLatestExchangeRate] util.SendRequest() failed: %v", err)
	}

	return data, http.StatusOK, nil
}
