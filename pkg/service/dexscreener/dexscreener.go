package dexscreener

import (
	"fmt"

	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

const (
	baseUrl = "https://api.dexscreener.com/latest/dex"
)

type dexscreener struct {
	sentry sentrygo.Service
}

func NewService(sentry sentrygo.Service) Service {
	return &dexscreener{
		sentry: sentry,
	}
}

var (
	sentryTags = map[string]string{
		"type": "system",
	}
)

func (d *dexscreener) Search(query string) ([]Pair, error) {
	pairResponse := PairResponse{}
	url := fmt.Sprintf("%s/search?q=%s", baseUrl, query)
	status, err := util.FetchData(url, &pairResponse)
	if err != nil {
		d.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - DexScreener - Search failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"query": query,
			},
		})
		return nil, fmt.Errorf("failed to fetch data from dexscreener: %w", err)
	}

	if status != 200 {
		d.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - DexScreener - Search failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"query": query,
			},
		})
		return nil, fmt.Errorf("failed to fetch data from dexscreener, status: %d", status)
	}

	return pairResponse.Pairs, nil
}

func (d *dexscreener) GetByChainAndPairAddress(network, pairAddr string) (*Pair, error) {
	pairResponse := PairResponse{}
	url := fmt.Sprintf("%s/pairs/%s/%s", baseUrl, network, pairAddr)
	status, err := util.FetchData(url, &pairResponse)
	if err != nil {
		d.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - DexScreener - GetByChainAndPairAddress failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"network": network,
				"address": pairAddr,
			},
		})
		return nil, fmt.Errorf("failed to fetch data from dexscreener: %w", err)
	}

	if status != 200 {
		d.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - DexScreener - GetByChainAndPairAddress failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"network": network,
				"address": pairAddr,
			},
		})
		return nil, fmt.Errorf("failed to fetch data from dexscreener, status: %d", status)
	}

	if len(pairResponse.Pairs) == 0 {
		return nil, fmt.Errorf("failed to fetch data from dexscreener, no data")
	}

	pair := pairResponse.Pairs[0]

	return &pair, nil
}

func (d *dexscreener) GetByTokenAddress(tokenAddr string) ([]Pair, error) {
	pairResponse := PairResponse{}
	url := fmt.Sprintf("%s/tokens/%s", baseUrl, tokenAddr)
	status, err := util.FetchData(url, &pairResponse)
	if err != nil {
		d.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - DexScreener - GetByTokenAddress failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"address": tokenAddr,
			},
		})
		return nil, fmt.Errorf("failed to fetch data from dexscreener: %w", err)
	}

	if status != 200 {
		d.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - DexScreener - GetByTokenAddress failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"address": tokenAddr,
			},
		})
		return nil, fmt.Errorf("failed to fetch data from dexscreener, status: %d", status)
	}

	return pairResponse.Pairs, nil
}
