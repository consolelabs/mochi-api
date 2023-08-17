package dexscreener

import (
	"fmt"

	"github.com/defipod/mochi/pkg/util"
)

const (
	baseUrl = "https://api.dexscreener.com/latest/dex"
)

type dexscreener struct {
}

func NewService() Service {
	return &dexscreener{}
}

func (d *dexscreener) Search(query string) ([]Pair, error) {
	pairResponse := PairResponse{}
	url := fmt.Sprintf("%s/search?q=%s", baseUrl, query)
	status, err := util.FetchData(url, &pairResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from dexscreener: %w", err)
	}

	if status != 200 {
		return nil, fmt.Errorf("failed to fetch data from dexscreener, status: %d", status)
	}

	return pairResponse.Pairs, nil
}

func (d *dexscreener) Get(network, address string) (*Pair, error) {
	pairResponse := PairResponse{}
	url := fmt.Sprintf("%s/pairs/%s/%s", baseUrl, network, address)
	status, err := util.FetchData(url, &pairResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from dexscreener: %w", err)
	}

	if status != 200 {
		return nil, fmt.Errorf("failed to fetch data from dexscreener, status: %d", status)
	}

	if len(pairResponse.Pairs) == 0 {
		return nil, fmt.Errorf("failed to fetch data from dexscreener, no data")
	}

	pair := pairResponse.Pairs[0]

	return &pair, nil
}
