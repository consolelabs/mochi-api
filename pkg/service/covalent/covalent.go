package covalent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

var networks = map[int]string{
	1:     "eth-mainnet",
	56:    "bsc-mainnet",
	137:   "matic-mainnet",
	250:   "fantom-mainnet",
	999:   "solana-mainnet",
	42161: "arbitrum-mainnet",
}

var (
	covalentSolanaTokenBalanceKey = "covalent-solana-balance-token"
	covalentTokenBalanceKey       = "covalent-balance-token"
	covalentTransactionKey        = "covalent-transaction"
)

type Covalent struct {
	config *config.Config
	logger logger.Logger
	cache  cache.Cache
}

func NewService(cfg *config.Config, l logger.Logger, cache cache.Cache) Service {
	return &Covalent{
		config: cfg,
		logger: l,
		cache:  cache,
	}
}

func (c *Covalent) getFullUrl(endpoint string, idx int) string {
	url := c.config.CovalentBaseUrl + endpoint
	if strings.Contains(url, "?") {
		url += "&key="
	} else {
		url += "?key="
	}
	return url + c.config.CovalentAPIKeys[idx]
}

func (c *Covalent) GetHistoricalTokenPrices(chainID int, currency string, address string) (*response.HistoricalTokenPricesResponse, error, int) {
	chainName, ok := networks[chainID]
	if !ok {
		chainName = fmt.Sprint(chainID)
	}
	endpoint := fmt.Sprintf("/pricing/historical_by_addresses_v2/%s/%s/%s/?no-spam=true&no-nft-fetch=true&nft=false", chainName, currency, address)
	res := &response.HistoricalTokenPricesResponse{}
	code, err := c.fetchCovalentData(endpoint, res)
	if err != nil || code != http.StatusOK {
		c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.GetTransactionsByAddress] util.FetchData() failed")
		return nil, fmt.Errorf("failed to fetch token data of %s: %v", currency, err), code
	}
	return res, nil, http.StatusOK
}

func (c *Covalent) GetTransactionsByAddress(chainID int, address string, size int, retry int) (*GetTransactionsResponse, error) {
	c.logger.Debug("start Covalent.GetTransactionsByAddress()")
	defer c.logger.Debug("end Covalent.GetTransactionsByAddress()")

	var data GetTransactionsResponse
	// check if data cached

	cached, err := c.doCacheTransaction(chainID, address)
	if err == nil && cached != "" {
		c.logger.Infof("hit cache data covalent-service, address: %s", address)
		go c.doNetworkTransaction(chainID, address, size, retry)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	go c.doNetworkTransaction(chainID, address, size, retry)
	return nil, nil
}

func (c *Covalent) GetTransactionsByAddressV3(chainID int, address string, size int, retry int) (*GetTransactionsResponse, error) {
	endpoint := fmt.Sprintf("/%d/address/%s/transactions_v2/?page-size=%d&no-spam=true&no-nft-fetch=true&nft=false", chainID, address, size)
	res := &GetTransactionsResponse{}
	code, err := c.fetchCovalentData(endpoint, res)
	if err != nil {
		c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.GetTransactionsByAddress] util.FetchData() failed")
		return nil, err
	}
	if res.Error {
		if retry == 0 {
			return nil, fmt.Errorf("%d - %s", res.ErrorCode, res.ErrorMessage)
		} else {
			return c.GetTransactionsByAddress(chainID, address, size, retry-1)
		}
	}
	return res, nil
}

func (c *Covalent) GetTokenBalances(chainID int, address string, retry int) (*GetTokenBalancesResponse, error) {
	c.logger.Debug("start Covalent.GetTokenBalances()")
	defer c.logger.Debug("end Covalent.GetTokenBalances()")

	var data GetTokenBalancesResponse
	// check if data cached

	cached, err := c.doCacheTokenBalances(chainID, address)
	if err == nil && cached != "" {
		c.logger.Infof("hit cache data covalent-service, address: %s", address)
		go c.doNetworkTokenBalances(chainID, address, retry)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return c.doNetworkTokenBalances(chainID, address, retry)
}

func (c *Covalent) GetSolanaTokenBalances(chainName string, address string, retry int) (*GetTokenBalancesResponse, error) {
	c.logger.Debug("start Covalent.GetSolanaTokenBalances()")
	defer c.logger.Debug("end Covalent.GetSolanaTokenBalances()")

	var data GetTokenBalancesResponse
	// check if data cached

	cached, err := c.doCacheSolanaTokenBalances(address)
	if err == nil && cached != "" {
		c.logger.Infof("hit cache data covalent-service, address: %s", address)
		go c.doNetworkSolanaTokenBalances(chainName, address, retry)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return c.doNetworkSolanaTokenBalances(chainName, address, retry)
}
