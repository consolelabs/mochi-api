package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
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

	if err != nil {
		return nil, err
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
		Aggregator: "jupiter",
	}, nil
}

func (j *JupyterProvider) GetRoutes(fromTokens, toTokens []model.Token, amount string) ([]response.ProviderSwapRoutes, error) {
	return nil, nil
}

func (j *JupyterProvider) BuildSwapRoutes(chainName string, req *request.BuildSwapRouteRequest) (*response.BuildRoute, error) {
	jupiterReq := request.JupiterBuildSwapRouteRequest{
		QuoteResponse:                 req.SwapData,
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

	return &response.BuildRoute{
		Code:    200,
		Message: "successfully",
		Data: response.BuildRouteData{
			AmountIn:     req.RouteSummary.AmountIn,
			AmountInUsd:  "",
			AmountOut:    req.RouteSummary.AmountOut,
			AmountOutUsd: "",
			Data:         res.SwapTransaction,
		},
	}, nil
}
