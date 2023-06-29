package covalent

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
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
	endpoint := fmt.Sprintf("/%d/address/%s/balances_v2/?no-spam=true&no-nft-fetch=true&nft=false", chainID, address)
	res := &GetTokenBalancesResponse{}
	code, err := c.fetchCovalentData(endpoint, res)
	if err != nil {
		c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.GetTokenBalances] util.FetchData() failed")
		return nil, err
	}
	if res.Error {
		if res.ErrorCode == http.StatusNotAcceptable {
			//TODO: predictably timeout -> should ignore now to avoid missing data from other chains. Will be fixed in the future
			c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.fetchCovalentData] Endpoint will predictably time out")
			return res, nil
		}

		if retry == 0 {
			return nil, fmt.Errorf("%d - %s", res.ErrorCode, res.ErrorMessage)
		} else {
			return c.GetTokenBalances(chainID, address, retry-1)
		}
	}
	return res, nil
}

func (c *Covalent) GetSolanaTokenBalances(chainName string, address string, retry int) (*GetTokenBalancesResponse, error) {
	c.logger.Debug("start Covalent.GetSolanaTokenBalances()")
	defer c.logger.Debug("end Covalent.GetSolanaTokenBalances()")

	var data GetTokenBalancesResponse
	// check if data cached

	cached, err := c.doCacheSolanaTokenBalances(address)
	if err == nil && cached != "" {
		c.logger.Infof("hit cache data krystal-service, address: %s", address)
		defer c.doNetworkSolanaTokenBalances(chainName, address, retry)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return c.doNetworkSolanaTokenBalances(chainName, address, retry)
}

func (c *Covalent) doCacheSolanaTokenBalances(address string) (string, error) {
	return c.cache.GetString(fmt.Sprintf("%s-%s", covalentSolanaTokenBalanceKey, strings.ToLower(address)))
}

func (c *Covalent) doNetworkSolanaTokenBalances(chainName string, address string, retry int) (*GetTokenBalancesResponse, error) {
	endpoint := fmt.Sprintf("/%s/address/%s/balances_v2/?no-spam=true&no-nft-fetch=true&nft=false", chainName, address)
	res := &GetTokenBalancesResponse{}
	code, err := c.fetchCovalentData(endpoint, res)
	if err != nil {
		c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.chainName] util.FetchData() failed")
		return nil, err
	}
	if res.Error {
		if retry == 0 {
			return nil, fmt.Errorf("%d - %s", res.ErrorCode, res.ErrorMessage)
		} else {
			return c.GetSolanaTokenBalances(chainName, address, retry-1)
		}
	}

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	c.logger.Infof("cache data covalent-service, key: %s", covalentSolanaTokenBalanceKey)
	c.cache.Set(covalentSolanaTokenBalanceKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (c *Covalent) fetchCovalentData(endpoint string, parseForm interface{}) (int, error) {
	var success bool
	for i, key := range c.config.CovalentAPIKeys {
		if key == "" {
			c.logger.Info("[covalent.fetchCovalentData] env COVALENT_API_KEYS has not been set")
			continue
		}
		url := c.getFullUrl(endpoint, i)
		code, err := util.FetchData(url, parseForm)
		if code == 402 {
			c.logger.Fields(logger.Fields{"code": code}).Infof("[covalent.fetchCovalentData] Exceed limit API key at index %d", i)
			continue
		}
		if err != nil {
			c.logger.Fields(logger.Fields{"code": code}).Error(err, "[covalent.fetchCovalentData] util.FetchData() failed")
			return code, err
		}
		// shift usable key to first idx, save time for later requests
		c.config.CovalentAPIKeys[0], c.config.CovalentAPIKeys[i] = c.config.CovalentAPIKeys[i], c.config.CovalentAPIKeys[0]
		success = true
		break
	}
	if !success {
		return http.StatusPaymentRequired, errors.New("all API keys may exceed their limit")
	}
	return http.StatusOK, nil
}
