package apilayer

type Service interface {
	GetHistoricalExchangeRate(q GetHistoricalExchangeRateQuery) (*GetHistoricalExchangeRateResponse, int, error)
	GetLatestExchangeRate(q GetLatestExchangeRateQuery) (*GetLatestExchangeRateResponse, int, error)
}
