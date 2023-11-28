package chainexplorer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

type chainExplorer struct {
	cfg    config.Config
	log    logger.Logger
	cache  cache.Cache
	sentry sentrygo.Service
}

func NewService(cfg config.Config, log logger.Logger, cache cache.Cache, sentry sentrygo.Service) Service {
	return &chainExplorer{
		cfg:    cfg,
		log:    log,
		cache:  cache,
		sentry: sentry,
	}
}

var (
	sentryTags = map[string]string{
		"type": "system",
	}
)

func (c *chainExplorer) GetGasTracker(listChain []model.Chain) ([]response.GasTrackerResponse, error) {
	apiKey := ""
	gasTrackerResp := make([]response.GasTrackerResponse, 0)
	for _, chain := range listChain {
		apiKey = c.getChainApiKey(chain.ShortName)

		url := fmt.Sprintf("%smodule=gastracker&action=gasoracle&apikey=%s", chain.APIBaseURL, apiKey)
		gasTracker, err := c.executeGetGasTracker(url)
		if err != nil {
			c.log.Fields(logger.Fields{"chain": chain}).Error(err, "failed to get gas tracker")
			return nil, err
		}
		gasTrackerResp = append(gasTrackerResp, response.GasTrackerResponse{
			Chain:           chain.Name,
			SafeGasPrice:    gasTracker.Result.SafeGasPrice,
			ProposeGasPrice: gasTracker.Result.ProposeGasPrice,
			FastGasPrice:    gasTracker.Result.FastGasPrice,
			// not have this data, temp hardcode
			EstSafeTime:    "180",
			EstProposeTime: "180",
			EstFastTime:    "30",
		})
	}
	return gasTrackerResp, nil
}

func (c *chainExplorer) getChainApiKey(chain string) string {
	switch chain {
	case "ftm":
		return c.cfg.ChainExplorer.FtmScanApiKey
	case "bsc":
		return c.cfg.ChainExplorer.BscScanApiKey
	case "eth":
		return c.cfg.ChainExplorer.EtherScanApiKey
	case "polygon":
		return c.cfg.ChainExplorer.PolygonScanApiKey
	default:
		return ""
	}
}

func (c *chainExplorer) executeGetGasTracker(url string) (*response.ChainExplorerGasTracker, error) {
	resp := &response.ChainExplorerGasTracker{}
	cached, err := c.doCacheGasTracker(url)
	if err == nil && cached != "" {
		go c.doNetworkGetGasTracker(url, resp)
		return resp, json.Unmarshal([]byte(cached), resp)
	}

	err = util.RetryRequest(func() error {
		return c.doNetworkGetGasTracker(url, resp)
	}, 5, 2*time.Second)
	if err != nil {
		c.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - ChainExplorer - doNetWorkGetGasTracker failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		c.log.Error(err, "[chainexplorer.executeGasTracker] c.doNetworkGetGastracker() failed")
		return nil, err
	}

	return resp, nil
}

func (c *chainExplorer) doNetworkGetGasTracker(url string, resp interface{}) error {
	query := util.SendRequestQuery{
		URL:       url,
		ParseForm: resp,
	}
	statusCode, err := util.SendRequest(query)
	if err != nil {
		return fmt.Errorf("send request failed: %v", err)
	}

	if statusCode != http.StatusOK {
		return fmt.Errorf("get gas tracker from explorer failed, status code: %d", statusCode)
	}

	// cache data
	bytes, _ := json.Marshal(resp)
	c.cache.Set(url, string(bytes), 1*time.Hour)

	return nil
}

func (c *chainExplorer) doCacheGasTracker(url string) (string, error) {
	return c.cache.GetString(url)
}
