package apilayer

type GetHistoricalExchangeRateResponse struct {
	Success    bool                          `json:"success"`
	Timeseries bool                          `json:"timeseries"`
	StartDate  string                        `json:"start_date"`
	EndDate    string                        `json:"end_date"`
	Base       string                        `json:"base"`
	Rates      map[string]map[string]float64 `json:"rates"`
}

type GetLatestExchangeRateResponse struct {
	Success   bool               `json:"success"`
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}
