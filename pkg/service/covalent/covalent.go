package covalent

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Covalent struct {
	getTokenURL string
}

func NewService(cfg *config.Config) Service {
	return &Covalent{
		getTokenURL: "https://api.covalenthq.com/v1/pricing/historical_by_addresses_v2/%d/%s/%s/?key=" + cfg.CovalentAPIKey,
	}
}

func (c Covalent) GetHistoricalTokenPrices(chainID int, currency string, address string) (*response.HistoricalTokenPricesResponse, error, int) {
	data := &response.HistoricalTokenPricesResponse{}
	statusCode, err := util.FetchData(fmt.Sprintf(c.getTokenURL, chainID, currency, address), data)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch token data of %s: %v", currency, err), statusCode
	}

	return data, nil, http.StatusOK
}
