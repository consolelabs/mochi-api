package provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type KyberProvider struct {
	config *config.Config
	logger logger.Logger
}

func NewKyber(cfg *config.Config, l logger.Logger) Provider {
	return &KyberProvider{
		config: cfg,
		logger: l,
	}
}

var (
	kyberBaseURL = "https://aggregator-api.kyberswap.com"
)

func (k *KyberProvider) GetRoute(fromToken, toToken, chain, amount string) (*response.KyberSwapRoutes, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/api/v1/routes?tokenIn=%s&tokenOut=%s&amountIn=%s", kyberBaseURL, chain, fromToken, toToken, amount), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("clientData", fmt.Sprintf("{\"source\": \"%s\"}", consts.ClientID))

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &response.KyberSwapRoutes{}

	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k *KyberProvider) GetRoutes(fromTokens, toTokens []model.Token, amount string) (routes []response.KyberSwapRoutes, err error) {
	// get list matching chain of fromTokens and toTokens
	for _, fromToken := range fromTokens {
		for _, toToken := range toTokens {
			if toToken.ChainID == fromToken.ChainID {
				// get routes
				route, err := k.GetRoute(fromToken.Address, toToken.Address, util.ConvertChainIdToChainName(int64(fromToken.ChainID)), amount)
				if err != nil {
					return nil, err
				}

				// code kyber 0 means success, else failed
				if route.Code == 0 {
					routes = append(routes, *route)
				}

			}
		}
	}

	return routes, nil
}
