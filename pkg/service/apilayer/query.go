package apilayer

type GetHistoricalExchangeRateQuery struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Base      string `json:"base"`
	Target    string `json:"target"`
}

type GetLatestExchangeRateQuery struct {
	Base   string `json:"base"`
	Target string `json:"target"`
}
