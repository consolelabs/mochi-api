package nghenhan

import "github.com/defipod/mochi/pkg/response"

type Service interface {
	GetFiatHistoricalChart(base, target, interval string, limit int) (*response.NghenhanFiatHistoricalChartResponse, error)
}
