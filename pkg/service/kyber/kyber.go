package kyber

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
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

func (k *kyberService) GetSwapRoutes(chain, fromAddress, toAddress, amount string) (*response.KyberSwapRoutes, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/api/v1/routes?tokenIn=%s&tokenOut=%s&amountIn=%s", k.kyberBaseUrl, chain, fromAddress, toAddress, amount), nil)
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
	res := &response.KyberSwapRoutes{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
