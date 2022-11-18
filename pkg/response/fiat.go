package response

type NghenhanFiatHistoricalChartResponse struct {
	Data []NghenhanFiatHistoricalChart `json:"data"`
}

type NghenhanFiatHistoricalChart struct {
	ClosePrice string  `json:"close_price"`
	HighPrice  string  `json:"high_price"`
	Interval   string  `json:"interval"`
	LowPrice   string  `json:"low_price"`
	OpenPrice  string  `json:"open_price"`
	OpenTime   int     `json:"open_time"`
	Symbol     string  `json:"symbol"`
	Volume     float64 `json:"volume"`
}
