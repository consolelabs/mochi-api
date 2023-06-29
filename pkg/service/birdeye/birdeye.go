package birdeye

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

type birdeye struct {
	config *config.Config
	logger logger.Logger
	cache  cache.Cache
}

func NewService(cfg *config.Config, l logger.Logger, cache cache.Cache) Service {
	return &birdeye{
		config: cfg,
		logger: l,
		cache:  cache,
	}
}

var (
	publicBirdeye        = "https://public-api.birdeye.so"
	birdeyeTokenPriceKey = "birdeye-token-price"
)

func (b *birdeye) GetTokenPrice(address string) (*TokenPrice, error) {
	b.logger.Debug("start skymavis.GetAddressFarming()")
	defer b.logger.Debug("end skymavis.GetAddressFarming()")

	var data TokenPrice
	// check if data cached

	cached, err := b.doCacheTokenPrice(address)
	if err == nil && cached != "" {
		b.logger.Infof("hit cache data birdeye-service, address: %s", address)
		go b.doNetworkTokenPrice(address)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return b.doNetworkTokenPrice(address)
}

func (b *birdeye) fetchBirdeyeData(url string, v any) error {
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &v,
		Headers:   map[string]string{"Content-Type": "application/json"},
	}

	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch Birdeye data")
	}

	return nil
}

func (b *birdeye) doCacheTokenPrice(address string) (string, error) {
	return b.cache.GetString(fmt.Sprintf("%s-%s", birdeyeTokenPriceKey, strings.ToLower(address)))
}

func (b *birdeye) doNetworkTokenPrice(address string) (*TokenPrice, error) {
	var res TokenPrice
	url := fmt.Sprintf("%s/public/price?address=%s", publicBirdeye, address)
	err := b.fetchBirdeyeData(url, &res)
	if err != nil {
		b.logger.Fields(logger.Fields{"url": url}).Error(err, "[birdeye.GetTokenPrice] b.fetchBirdeyeData() failed")
		return nil, err
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	b.logger.Infof("cache data birdeye-service, key: %s", birdeyeTokenPriceKey)
	b.cache.Set(birdeyeTokenPriceKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return &res, nil
}
