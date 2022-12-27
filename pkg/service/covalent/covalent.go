package covalent

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Covalent struct {
	config *config.Config
	logger logger.Logger
}

func NewService(cfg *config.Config, l logger.Logger) Service {
	return &Covalent{
		config: cfg,
		logger: l,
	}
}

func (c *Covalent) getFullUrl(urlPath string) string {
	url := c.config.CovalentBaseUrl + urlPath
	if strings.Contains(url, "?") {
		url += "&key="
	} else {
		url += "?key="
	}
	return url + c.config.CovalentAPIKey
}

func (c *Covalent) GetHistoricalTokenPrices(chainID int, currency string, address string) (*response.HistoricalTokenPricesResponse, error, int) {
	data := &response.HistoricalTokenPricesResponse{}
	url := c.getFullUrl(fmt.Sprintf("/pricing/historical_by_addresses_v2/%d/%s/%s/", chainID, currency, address))
	statusCode, err := util.FetchData(url, data)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch token data of %s: %v", currency, err), statusCode
	}

	return data, nil, http.StatusOK
}

func (c *Covalent) GetTransactionsByAddress(chainID int, address string) (*GetTransactionsResponse, error) {
	url := c.getFullUrl(fmt.Sprintf("/%d/address/%s/transactions_v2/?page-size=1000", chainID, address))
	res := &GetTransactionsResponse{}
	statusCode, err := util.FetchData(url, res)
	if err != nil {
		c.logger.Fields(logger.Fields{"url": url, "status": statusCode}).Error(err, "[covalent.GetTransactionsByAddress] utils.MakeRequest() failed")
		return nil, err
	}
	return res, nil
}
