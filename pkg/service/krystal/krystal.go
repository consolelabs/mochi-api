package krystal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/util"
)

type Krystal struct {
	config *config.Config
	logger logger.Logger
	cache  cache.Cache
}

func NewService(cfg *config.Config, l logger.Logger, cache cache.Cache) Service {
	return &Krystal{
		config: cfg,
		logger: l,
		cache:  cache,
	}
}

var (
	key = "krystal-balance-token"
)

func (k *Krystal) GetBalanceTokenByAddress(address string) (*BalanceTokenResponse, error) {
	k.logger.Debug("start krystal.GetBalanceTokenByAddress()")
	defer k.logger.Debug("end krystal.GetBalanceTokenByAddress()")

	var data BalanceTokenResponse
	// check if data cached

	cached, err := k.doCache(address)
	if err == nil && cached != "" {
		k.logger.Infof("hit cache data krystal-service, address: %s", address)
		go k.doNetwork(address, data)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return k.doNetwork(address, data)
}

func (k *Krystal) doCache(address string) (string, error) {
	return k.cache.GetString(fmt.Sprintf("%s-%s", key, strings.ToLower(address)))
}

func (k *Krystal) doNetwork(address string, data BalanceTokenResponse) (*BalanceTokenResponse, error) {
	chainIDs := []int{1, 56, 137, 43114, 25, 250, 42161, 1313161554, 8217, 10, 101}
	chainIDsStr := strings.ReplaceAll(strings.Trim(fmt.Sprint(chainIDs), "[]"), " ", ",")

	url := k.config.KrystalBaseUrl + fmt.Sprintf("/all/v1/balance/token?addresses=ethereum:%s&quoteSymbols=usd&sparkline=false&chainIds=%s", address, chainIDsStr)

	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   map[string]string{"x-rate-access-token": k.config.KrystalApiKey},
	}

	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[krystal.GetBalanceTokenByAddress] util.SendRequest() failed: %v", err)
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&data)
	k.logger.Infof("cache data krystal-service, key: %s", key)
	k.cache.Set(key+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return &data, nil
}
