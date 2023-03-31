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

func (k *kyberService) GetSwapRoutes(amount string) (*response.KyberSwapRoutes, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/ethereum/api/v1/routes?tokenIn=0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE&tokenOut=0xdAC17F958D2ee523a2206206994597C13D831ec7&amountIn=%s", k.kyberBaseUrl, amount), nil)
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
