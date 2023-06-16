package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/exp/slices"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
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
	request.Header.Add("x-client-id", consts.ClientID)

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
	listFailedStatusKyber := []int64{4001, 4002, 4005, 4007, 4009, 4010, 4011}
	// get list matching chain of fromTokens and toTokens
	for _, fromToken := range fromTokens {
		for _, toToken := range toTokens {
			if toToken.ChainID == fromToken.ChainID {
				// get routes
				route, err := k.GetRoute(fromToken.Address, toToken.Address, util.ConvertChainIdToChainName(int64(fromToken.ChainID)), amount)
				if err != nil {
					return nil, err
				}

				k.logger.Fields(logger.Fields{"route": route}).Info("[kyber.GetRoutes] - get route")

				// code kyber 0 means success, else failed
				if route.Code == 0 && route.Message == "successfully" {
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
				} else if slices.Contains(listFailedStatusKyber, route.Code) {
					routes = append(routes, response.ProviderSwapRoutes{
						Code:    2,
						Message: route.Message,
						Data:    response.RouteSummaryData{},
					})
				} else {
					routes = append(routes, response.ProviderSwapRoutes{
						Code:    3,
						Message: route.Message,
						Data:    response.RouteSummaryData{},
					})
				}
				k.logger.Fields(logger.Fields{"routes": routes}).Info("[kyber.GetRoutes] - get routes")

			}
		}
	}

	return routes, nil
}

func (k *KyberProvider) BuildSwapRoutes(chainName string, req *request.BuildSwapRouteRequest) (*response.BuildRoute, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	jsonBody := bytes.NewBuffer(body)

	var client = &http.Client{}
	url := fmt.Sprintf("%s/%s/api/v1/route/build", kyberBaseURL, chainName)
	k.logger.Info("check kyberswap data:")
	k.logger.Infof("url: ", url)
	k.logger.Fields(logger.Fields{"req": req}).Infof("check req payload")
	request, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("source", consts.ClientID)
	request.Header.Add("x-client-id", consts.ClientID)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &response.BuildRoute{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
