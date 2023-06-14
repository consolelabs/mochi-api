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

func (k *KyberProvider) GetRoute(fromToken, toToken, chain, amount string) (*response.ProviderSwapRoutes, error) {
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

	res := &response.ProviderSwapRoutes{}

	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (k *KyberProvider) GetRoutes(fromTokens, toTokens []model.Token, amount string) (routes []response.ProviderSwapRoutes, err error) {
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
					route.Data.TokenIn = response.RouteToken{
						Address:     fromToken.Address,
						ChainId:     int64(fromToken.ChainID),
						ChainName:   util.ConvertChainIdToChainName(int64(fromToken.ChainID)),
						Decimals:    int64(fromToken.Decimals),
						Symbol:      fromToken.Symbol,
						Name:        fromToken.Name,
						CoingeckoId: fromToken.CoinGeckoID,
					}

					route.Data.TokenOut = response.RouteToken{
						Address:     toToken.Address,
						ChainId:     int64(toToken.ChainID),
						ChainName:   util.ConvertChainIdToChainName(int64(toToken.ChainID)),
						Decimals:    int64(toToken.Decimals),
						Symbol:      toToken.Symbol,
						Name:        toToken.Name,
						CoingeckoId: toToken.CoinGeckoID,
					}

					route.Code = 1

					routes = append(routes, *route)
				} else if route.Code == 4008 {
					routes = append(routes, response.ProviderSwapRoutes{
						Code:    0,
						Message: route.Message,
						Data:    response.RouteSummaryData{},
					})
				} else {
					routes = append(routes, response.ProviderSwapRoutes{
						Code:    2,
						Message: route.Message,
						Data:    response.RouteSummaryData{},
					})
				}

			}
		}
	}

	return routes, nil
}
