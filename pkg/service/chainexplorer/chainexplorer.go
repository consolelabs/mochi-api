package chainexplorer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

type chainExplorer struct {
	cfg config.Config
	log logger.Logger
}

func NewService(cfg config.Config, log logger.Logger) Service {
	return &chainExplorer{
		cfg: cfg,
		log: log,
	}
}

var (
	listChainSupportGasTracker = []string{"ftm", "bsc", "ether", "polygon"}
)

func (c *chainExplorer) GetGasTracker() ([]response.GasTrackerResponse, error) {
	var chainExplorerBaseURL, apiKey string
	gasTrackerResp := make([]response.GasTrackerResponse, 0)
	for _, chain := range listChainSupportGasTracker {
		if chain == "ether" {
			chainExplorerBaseURL = fmt.Sprintf("https://api.%sscan.io/api", chain)
		} else {
			chainExplorerBaseURL = fmt.Sprintf("https://api.%sscan.com/api", chain)
		}

		apiKey = c.getChainApiKey(chain)

		gasTracker, err := c.executeGetGasTracker(chainExplorerBaseURL, apiKey)
		if err != nil {
			c.log.Fields(logger.Fields{"chain": chain}).Error(err, "failed to get gas tracker")
			return nil, err
		}
		gasTrackerResp = append(gasTrackerResp, response.GasTrackerResponse{
			Chain:           strings.ToUpper(chain[0:3]),
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
	case "ether":
		return c.cfg.ChainExplorer.EtherScanApiKey
	case "polygon":
		return c.cfg.ChainExplorer.PolygonScanApiKey
	default:
		return ""
	}
}
func (c *chainExplorer) executeGetGasTracker(url, apiKey string) (*response.ChainExplorerGasTracker, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s?module=gastracker&action=gasoracle&apikey=%s", url, apiKey), nil)
	if err != nil {
		return nil, err
	}

	responseURL, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer responseURL.Body.Close()
	resBody, err := ioutil.ReadAll(responseURL.Body)
	if err != nil {
		return nil, err
	}

	res := &response.ChainExplorerGasTracker{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
