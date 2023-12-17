package covalent

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

var (
	retryMap map[string]string = make(map[string]string)
)

func (c *Covalent) doCacheSolanaTokenBalances(address string) (string, error) {
	return c.cache.GetString(fmt.Sprintf("%s-%s", covalentSolanaTokenBalanceKey, strings.ToLower(address)))
}

func (c *Covalent) doNetworkSolanaTokenBalances(chainName string, address string, retry int) (*GetTokenBalancesResponse, error) {
	endpoint := fmt.Sprintf("/%s/address/%s/balances_v2/?no-spam=true&no-nft-fetch=true&nft=false", chainName, address)
	res := &GetTokenBalancesResponse{}
	code, err := c.fetchCovalentData(endpoint, res)
	if err != nil {
		c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.chainName] util.FetchData() failed")
		c.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Covalent - doNetworkTokenBalances failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"chainName": chainName,
				"address":   address,
				"retry":     retry,
			},
		})
		return nil, err
	}
	if res.Error {
		if retry == 0 {
			c.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
				Message: fmt.Sprintf("[API mochi] - Covalent - doNetworkTokenBalances failed %v", err),
				Tags:    sentryTags,
				Extra: map[string]interface{}{
					"chainName": chainName,
					"address":   address,
					"retry":     retry,
				},
			})
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

func (c *Covalent) doCacheTokenBalances(chainID int, address string) (string, error) {
	return c.cache.GetString(fmt.Sprintf("%s-%d-%s", covalentTokenBalanceKey, chainID, strings.ToLower(address)))
}

func (c *Covalent) doCacheTransaction(chainID int, address string) (string, error) {
	return c.cache.GetString(fmt.Sprintf("%s-%d-%s", covalentTransactionKey, chainID, strings.ToLower(address)))
}

func (c *Covalent) doNetworkTransaction(chainID int, address string, size int, retry int) (*GetTransactionsResponse, error) {
	endpoint := fmt.Sprintf("/%d/address/%s/transactions_v2/?page-size=%d&no-spam=true&no-nft-fetch=true&nft=false", chainID, address, size)
	res := &GetTransactionsResponse{}
	code, err := c.fetchCovalentData(endpoint, res)
	if err != nil {
		c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.GetTransactionsByAddress] util.FetchData() failed")
		return nil, err
	}
	if res.Error {
		if retry == 0 {
			c.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
				Message: fmt.Sprintf("[API mochi] - Covalent - doNetworkTransaction failed %v", err),
				Tags:    sentryTags,
				Extra: map[string]interface{}{
					"chainID": chainID,
					"address": address,
					"size":    size,
					"retry":   retry,
				},
			})
			return nil, fmt.Errorf("%d - %s", res.ErrorCode, res.ErrorMessage)
		} else {
			return c.GetTransactionsByAddress(chainID, address, size, retry-1)
		}
	}

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	c.logger.Infof("cache data covalent-service, key: %s", covalentTokenBalanceKey)
	c.cache.Set(fmt.Sprintf("%s-%d-%s", covalentTransactionKey, chainID, strings.ToLower(address)), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (c *Covalent) doNetworkTokenBalances(chainID int, address string, retry int) (*GetTokenBalancesResponse, error) {
	endpoint := fmt.Sprintf("/%d/address/%s/balances_v2/?no-spam=true&no-nft-fetch=true&nft=false", chainID, address)
	res := &GetTokenBalancesResponse{}
	code, err := c.fetchCovalentData(endpoint, res)
	if err != nil {
		c.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Covalent - doNetworkTokenBalances failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"chainID": chainID,
				"address": address,
				"retry":   retry,
			},
		})
		c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.GetTokenBalances] util.FetchData() failed")
		return nil, err
	}

	if res.Error {
		if res.ErrorCode == http.StatusNotAcceptable {
			//TODO: predictably timeout -> should ignore now to avoid missing data from other chains. Will be fixed in the future
			c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.fetchCovalentData] Endpoint will predictably time out")
			return res, nil
		}

		c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.GetTokenBalances] cannot get data from covalent, retrying ...")
		retryTime := retryMap[fmt.Sprintf("TokenBalances-%d-%s", chainID, address)]
		t, _ := time.Parse(time.RFC3339, retryTime)
		now := time.Now()

		// current temp solution to fix covalent exceed quota
		// check if last time retry < now + 30 minutes -> allow retry else ignore
		// TODO: need real solution for this
		if now.Sub(t) > 30*time.Minute {
			c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.GetTokenBalances] time passed, allow retry")
			// temp fix for covalent bug
			retryMap[fmt.Sprintf("TokenBalances-%d-%s", chainID, address)] = time.Now().Format(time.RFC3339)

			if retry == 0 {
				c.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
					Message: fmt.Sprintf("[API mochi] - Covalent - doNetworkTokenBalances failed %v", err),
					Tags:    sentryTags,
					Extra: map[string]interface{}{
						"chainID": chainID,
						"address": address,
						"retry":   retry,
					},
				})

				return nil, fmt.Errorf("%d - %s", res.ErrorCode, res.ErrorMessage)
			} else {
				return c.GetTokenBalances(chainID, address, retry-1)
			}
		} else {
			c.logger.Fields(logger.Fields{"endpoint": endpoint, "code": code}).Error(err, "[covalent.GetTokenBalances] time not passed, ignore retry")
			return &GetTokenBalancesResponse{Data: &GetTokenBalancesData{Items: nil}}, nil
		}

	}

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	c.logger.Infof("cache data covalent-service, key: %s", covalentTokenBalanceKey)
	c.cache.Set(fmt.Sprintf("%s-%d-%s", covalentTokenBalanceKey, chainID, strings.ToLower(address)), string(bytes), 7*24*time.Hour)

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
