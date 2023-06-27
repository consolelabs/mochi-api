package krystal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/util"
)

type Krystal struct {
	config *config.Config
	logger logger.Logger
}

func NewService(cfg *config.Config, l logger.Logger) Service {
	return &Krystal{
		config: cfg,
		logger: l,
	}
}

func (n *Krystal) GetBalanceTokenByAddress(address string, chainIDs []int) (*BalanceTokenResponse, error) {
	chainIDsStr := strings.ReplaceAll(strings.Trim(fmt.Sprint(chainIDs), "[]"), " ", ",")
	url := n.config.KrystalBaseUrl + fmt.Sprintf("/all/v1/balance/token?addresses=ethereum:%s&quoteSymbols=usd&sparkline=false&chainIds=%s", address, chainIDsStr)
	data := BalanceTokenResponse{}

	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   map[string]string{"x-rate-access-token": n.config.KrystalApiKey},
	}

	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[krystal.GetBalanceTokenByAddress] util.SendRequest() failed: %v", err)
	}

	return &data, nil
}
