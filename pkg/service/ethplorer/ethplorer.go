package ethplorer

import (
	"fmt"

	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

const (
	baseUrl = "https://api.ethplorer.io"
	apiKey  = "freekey"
)

type ethplorer struct {
	sentry sentrygo.Service
}

func NewService(sentry sentrygo.Service) Service {
	return &ethplorer{
		sentry: sentry,
	}
}

var (
	sentryTags = map[string]string{
		"type": "system",
	}
)

func (e *ethplorer) GetTopTokenHolders(address string, limit int) (*TokenHoldersResponse, error) {
	resp := &TokenHoldersResponse{}

	url := fmt.Sprintf("%s/getTopTokenHolders/%s?apiKey=%s&limit=%d", baseUrl, address, apiKey, limit)

	status, err := util.FetchData(url, resp)
	if err != nil {
		e.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - EthExplorer - GetTopTokenHolders failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		return nil, fmt.Errorf("fetch data failed: %w", err)
	}

	if status != 200 {
		e.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - EthExplorer - GetTopTokenHolders failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		return nil, fmt.Errorf("fetch data failed: status %d", status)
	}

	return resp, nil
}

func (e *ethplorer) GetTokenInfo(address string) (*TokenInfoResponse, error) {
	resp := &TokenInfoResponse{}

	url := fmt.Sprintf("%s/getTokenInfo/%s?apiKey=%s", baseUrl, address, apiKey)

	status, err := util.FetchData(url, resp)
	if err != nil {
		e.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - EthExplorer - GetTopTokenHolders failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		return nil, fmt.Errorf("fetch data failed: %w", err)
	}

	if status != 200 {
		e.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - EthExplorer - GetTopTokenHolders failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		return nil, fmt.Errorf("fetch data failed: status %d", status)
	}

	return resp, nil
}
