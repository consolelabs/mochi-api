package coingecko

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	errs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

func (c *CoinGecko) doCacheCoinByContract(platformId, contractAddress string) (string, error) {
	return c.cache.GetString(fmt.Sprintf("%s-%s-%s", coingeckoGetTokenByContractKey, platformId, strings.ToLower(contractAddress)))
}

func (c *CoinGecko) doCacheCoinsMarketData(ids []string, sparkline bool, page, pageSize string) (string, error) {
	return c.cache.GetString(fmt.Sprintf("%s-%s-%v", coingeckoCoinsMarketDataKey, strings.Join(ids, "-"), sparkline))
}

func (c *CoinGecko) doNetworkCoinByContract(platformId, contractAddress string, retry int) (*response.GetCoinByContractResponseData, error) {
	endpoint := fmt.Sprintf(c.getCoinByContract, platformId, contractAddress)
	res := &response.GetCoinByContractResponseData{}
	status, err := util.FetchData(endpoint, &res)
	if status == http.StatusNotFound {
		bytes, _ := json.Marshal("404")
		c.cache.Set(fmt.Sprintf("%s-%s-%s", coingeckoGetTokenByContractKey, platformId, strings.ToLower(contractAddress)), string(bytes), 7*24*time.Hour)
		return nil, errors.New("404 contract not found")
	}
	if err != nil || status != http.StatusOK {
		if retry == 0 {
			c.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
				Message: fmt.Sprintf("[API mochi] - Coingecko - doNetworkCoinByContract failed - %v", err),
				Tags:    sentryTags,
				Extra: map[string]interface{}{
					"platformId":      platformId,
					"contractAddress": contractAddress,
					"retry":           retry,
				},
			})
			return nil, fmt.Errorf("%d - %v", status, err)
		} else {
			return c.GetCoinByContract(platformId, contractAddress, retry-1)
		}
	}

	// cache solana-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	c.cache.Set(fmt.Sprintf("%s-%s-%s", coingeckoGetTokenByContractKey, platformId, strings.ToLower(contractAddress)), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (c *CoinGecko) doNetworkCoinsMarketData(ids []string, sparkline bool, page, pageSize string) ([]response.CoinMarketItemData, error, int) {
	res := make([]response.CoinMarketItemData, 0)
	var resTmp []response.CoinMarketItemDataRes

	statusCode, err := util.FetchData(fmt.Sprintf(c.getCoinsMarketData, strings.Join(ids, ","), pageSize, page, sparkline), &resTmp)
	if err != nil || statusCode != http.StatusOK {
		if statusCode != http.StatusNotFound {
			c.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
				Message: fmt.Sprintf("[API mochi] - Coingecko - GetCoinsMarketData failed %v", err),
				Tags:    sentryTags,
				Extra: map[string]interface{}{
					"ids":       ids,
					"page":      page,
					"pageSize":  pageSize,
					"sparkline": sparkline,
				},
			})
			return nil, err, statusCode
		}
		return nil, errs.ErrRecordNotFound, statusCode
	}

	for _, r := range resTmp {
		res = append(res, r.ToCoinMarketItemData())
	}

	bytes, _ := json.Marshal(&res)
	c.cache.Set(fmt.Sprintf("%s-%s-%v", coingeckoCoinsMarketDataKey, strings.Join(ids, "-"), sparkline), string(bytes), 7*24*time.Hour)

	return res, nil, http.StatusOK
}
