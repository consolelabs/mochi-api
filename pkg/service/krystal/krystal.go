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

func (n *Krystal) GetBalanceTokenByAddress(address string) (*BalanceTokenResponse, error) {
	n.logger.Debug("start krystal.GetBalanceTokenByAddress()")
	defer n.logger.Debug("end krystal.GetBalanceTokenByAddress()")

	var data BalanceTokenResponse
	// check if data cached
	key := fmt.Sprintf("krystal-balance-token-%s", strings.ToLower(address))
	cached, err := n.cache.GetString(key)
	if err == nil && cached != "" {
		n.logger.Infof("hit cache data krystal-service, address: %s", address)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	chainIDs := []int{1, 56, 137, 43114, 25, 250, 42161, 1313161554, 8217, 10, 101}
	chainIDsStr := strings.ReplaceAll(strings.Trim(fmt.Sprint(chainIDs), "[]"), " ", ",")

	url := n.config.KrystalBaseUrl + fmt.Sprintf("/all/v1/balance/token?addresses=ethereum:%s&quoteSymbols=usd&sparkline=false&chainIds=%s", address, chainIDsStr)

	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   map[string]string{"x-rate-access-token": n.config.KrystalApiKey},
	}

	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[krystal.GetBalanceTokenByAddress] util.SendRequest() failed: %v", err)
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&data)
	n.logger.Infof("cache data krystal-service, key: %s", key)
	n.cache.Set(key, string(bytes), 30*time.Minute)

	return &data, nil
}
