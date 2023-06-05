package jupiter

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

type jupiterService struct {
	config         *config.Config
	logger         logger.Logger
	jupiterBaseUrl string
}

func New(cfg *config.Config, l logger.Logger) Service {
	return &jupiterService{
		config:         cfg,
		logger:         l,
		jupiterBaseUrl: "https://quote-api.jup.ag/v5",
	}
}

func (k *jupiterService) GetSwapRoutesSolana(chain, fromAddress, toAddress, amount string) (*response.KyberSwapRoutes, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/quote?inputMint=%s&outputMint=%s&amount=%s&slippageBps=50&onlyDirectRoutes=false", k.jupiterBaseUrl, fromAddress, toAddress, amount), nil)
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

	return &response.KyberSwapRoutes{
		Code:    0,
		Message: "successfully",
		Data: response.RouteSummaryData{
			RouteSummary: model.RouteSummary{
				TokenIn:      fromAddress,
				AmountIn:     res.InAmount,
				AmountInUsd:  "",
				TokenOut:     toAddress,
				AmountOut:    res.OutAmount,
				AmountOutUsd: "",
				Route:        routes,
			},
			SwapData: res,
		},
	}, nil
}

func (k *jupiterService) BuildSwapRoutes(chainName string, req *request.JupiterBuildSwapRouteRequest) (*response.BuildRoute, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	jsonBody := bytes.NewBuffer(body)

	var client = &http.Client{}
	url := fmt.Sprintf("%s/swap", k.jupiterBaseUrl)
	k.logger.Info("check jupiter swap data:")
	k.logger.Infof("url: ", url)
	k.logger.Infof("jsonBody: ", jsonBody)
	k.logger.Fields(logger.Fields{"req": req}).Infof("check req payload")
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

	quoteResponse, _ := req.QuoteResponse.(request.QuoteResponse)
	return &response.BuildRoute{
		Code:    200,
		Message: "successfully",
		Data: response.BuildRouteData{
			AmountIn:     quoteResponse.InAmount,
			AmountInUsd:  "",
			AmountOut:    quoteResponse.OutAmount,
			AmountOutUsd: "",
			Data:         res.SwapTransaction,
		},
	}, nil
}
