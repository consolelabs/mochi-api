package nghenhan

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

type Nghenhan struct {
	baseUrl string
	sentry  sentrygo.Service
}

func NewService(sentry sentrygo.Service) Service {
	return &Nghenhan{
		baseUrl: "https://cex.console.so/api/v1",
		sentry:  sentry,
	}
}

var (
	sentryTags = map[string]string{
		"type": "system",
	}
)

func (n *Nghenhan) GetFiatHistoricalChart(base, target, interval string, limit int) (*response.NghenhanFiatHistoricalChartResponse, error) {
	tmpBase := base
	tmpTarg := target
	url := n.baseUrl + fmt.Sprintf("/rate?base=%s&target=%s&interval=%s&limit=%v", tmpBase, tmpTarg, interval, limit)
	data := response.NghenhanFiatHistoricalChartResponse{}
	req := util.SendRequestQuery{
		URL:      url,
		Response: &data,
		Headers:  map[string]string{"Content-Type": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		retry := 3
		for retry > 0 {
			statusCode, err = util.SendRequest(req)
			if err == nil && statusCode == http.StatusOK {
				break
			}
			retry--
		}

		if retry == 0 {
			n.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
				Message: fmt.Sprintf("[API mochi] - Nghenhan - GetFiatHistoricalChart failed %v", err),
				Tags:    sentryTags,
				Extra: map[string]interface{}{
					"base":     base,
					"target":   target,
					"interval": interval,
					"limit":    limit,
				},
			})
			return &response.NghenhanFiatHistoricalChartResponse{
				Data: []response.NghenhanFiatHistoricalChart{},
			}, nil
		}
	}
	return &data, nil
}
