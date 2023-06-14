package kyber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

type kyberService struct {
	config       *config.Config
	logger       logger.Logger
	kyberBaseUrl string
}

func New(cfg *config.Config, l logger.Logger) Service {
	return &kyberService{
		config:       cfg,
		logger:       l,
		kyberBaseUrl: "https://aggregator-api.kyberswap.com",
	}
}

func (k *kyberService) GetSwapRoutesEVM(chain, fromAddress, toAddress, amount string) (*response.ProviderSwapRoutes, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/api/v1/routes?tokenIn=%s&tokenOut=%s&amountIn=%s", k.kyberBaseUrl, chain, fromAddress, toAddress, amount), nil)
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

func (k *kyberService) BuildSwapRoutes(chainName string, req *request.KyberBuildSwapRouteRequest) (*response.BuildRoute, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	jsonBody := bytes.NewBuffer(body)

	var client = &http.Client{}
	url := fmt.Sprintf("%s/%s/api/v1/route/build", k.kyberBaseUrl, chainName)
	k.logger.Info("check kyberswap data:")
	k.logger.Infof("url: ", url)
	k.logger.Fields(logger.Fields{"req": req}).Infof("check req payload")
	request, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("source", consts.ClientID)

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
