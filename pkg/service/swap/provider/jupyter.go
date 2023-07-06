package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type JupyterProvider struct {
	config *config.Config
	logger logger.Logger
}

func NewJupyter(cfg *config.Config, l logger.Logger) Provider {
	return &JupyterProvider{
		config: cfg,
		logger: l,
	}
}

var (
	JupiterBaseUrl = "https://quote-api.jup.ag/v5"
)

func (j *JupyterProvider) GetRoute(fromToken, toToken, chain, amount string) (*response.ProviderSwapRoutes, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/quote?inputMint=%s&outputMint=%s&amount=%s&slippageBps=50&onlyDirectRoutes=false", JupiterBaseUrl, fromToken, toToken, amount), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("cannot get jupyter data")
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &response.JupiterSwapRoutesSol{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}

	// parse route
	routes := make([][]model.RouteElement, 0)
	newRoutes := make([]model.RouteElement, 0)
	swaps := res.RoutePlan
	for _, swap := range swaps {
		routeElement := new(model.RouteElement)
		routeElement.AmountOut = swap.SwapInfo.OutAmount
		routeElement.SwapAmount = swap.SwapInfo.InAmount
		routeElement.TokenIn = swap.SwapInfo.InputMint
		routeElement.TokenOut = swap.SwapInfo.OutputMint
		newRoutes = append(newRoutes, *routeElement)
	}
	routes = append(routes, newRoutes)

	return &response.ProviderSwapRoutes{
		Code:    0,
		Message: "successfully",
		Data: response.RouteSummaryData{
			RouteSummary: model.RouteSummary{
				TokenIn:      fromToken,
				AmountIn:     res.InAmount,
				AmountInUsd:  "",
				TokenOut:     toToken,
				AmountOut:    res.OutAmount,
				AmountOutUsd: "",
				Route:        routes,
			},
		},
		SwapData:   res,
		Aggregator: "jupyter",
	}, nil
}

func (j *JupyterProvider) GetRoutes(fromTokens, toTokens []model.Token, amount string) (routes []response.ProviderSwapRoutes, err error) {
	for _, fromToken := range fromTokens {
		for _, toToken := range toTokens {
			if toToken.ChainID == fromToken.ChainID {
				// get routes
				amount := util.FloatToString(amount, int64(fromToken.Decimals))
				route, err := j.GetRoute(fromToken.Address, toToken.Address, util.ConvertChainIdToChainName(int64(fromToken.ChainID)), amount)
				if err != nil {
					routes = append(routes, response.ProviderSwapRoutes{
						Code:    0,
						Message: "route not found",
						Data:    response.RouteSummaryData{},
					})
					continue
				}

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
			}
		}
	}

	return routes, nil
}

func (j *JupyterProvider) BuildSwapRoutes(chainName string, req *request.BuildSwapRouteRequest) (*response.BuildRoute, error) {
	jupiterReq := request.JupiterBuildSwapRouteRequest{
		QuoteResponse:                 req.RouteSummary,
		WrapAndUnwrapSol:              true,
		UserPublicKey:                 req.Sender,
		ComputeUnitPriceMicroLamports: "auto",
	}

	body, err := json.Marshal(jupiterReq)
	if err != nil {
		return nil, err
	}
	jsonBody := bytes.NewBuffer(body)

	var client = &http.Client{}
	url := fmt.Sprintf("%s/swap", JupiterBaseUrl)
	httpRequest, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &response.JupiterBuildRoute{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}

	quoteResp := &response.JupyterQuoteResponse{}
	quoteByte, _ := json.Marshal(req.RouteSummary)
	err = json.Unmarshal(quoteByte, quoteResp)
	if err != nil {
		return nil, err
	}

	return &response.BuildRoute{
		Code:    200,
		Message: "successfully",
		Data: response.BuildRouteData{
			AmountIn:     quoteResp.InAmount,
			AmountInUsd:  "",
			AmountOut:    quoteResp.OutAmount,
			AmountOutUsd: "",
			Data:         res.SwapTransaction,
		},
	}, nil
}
