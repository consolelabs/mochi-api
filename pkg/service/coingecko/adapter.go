package coingecko

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

func (c *CoinGecko) doCacheCoinByContract(platformId, contractAddress string) (string, error) {
	return c.cache.GetString(fmt.Sprintf("%s-%s-%s", coingeckoGetTokenByContractKey, platformId, strings.ToLower(contractAddress)))
}

func (c *CoinGecko) doNetworkCoinByContract(platformId, contractAddress string, retry int) (*response.GetCoinByContractResponseData, error) {
	endpoint := fmt.Sprintf(c.getCoinByContract, platformId, contractAddress)
	res := &response.GetCoinByContractResponseData{}
	status, err := util.FetchData(endpoint, &res)
	if err != nil || status != http.StatusOK {
		if retry == 0 {
			return nil, fmt.Errorf("%d - %s", status, err)
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
