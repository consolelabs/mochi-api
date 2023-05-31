package sui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
	"io/ioutil"
	"net/http"
)

type SuiService struct {
	config *config.Config
	logger logger.Logger
}

func New(cfg *config.Config, l logger.Logger) Service {
	return &SuiService{
		config: cfg,
		logger: l,
	}
}

func (s *SuiService) GetBalance(address string) (*response.SuiAllBalance, error) {
	var client = &http.Client{}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "suix_getAllBalances",
		"params":  []string{address},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	jsonBody := bytes.NewBuffer(body)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s", s.config.Sui.Rpc), jsonBody)
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
	res := &response.SuiAllBalance{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SuiService) GetCoinMetadata(coinType string) (*response.SuiCoinMetadata, error) {
	var client = &http.Client{}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "suix_getCoinMetadata",
		"params":  []string{coinType},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	jsonBody := bytes.NewBuffer(body)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s", s.config.Sui.Rpc), jsonBody)
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
	res := &response.SuiCoinMetadata{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SuiService) GetAddressAssets(address string) ([]response.WalletAssetData, error) {
	walletAssetList := make([]response.WalletAssetData, 0)
	allBalanceRes, err := s.GetBalance(address)
	if err != nil {
		return []response.WalletAssetData{}, err
	}

	balanceTokenList := allBalanceRes.Result

	for _, token := range balanceTokenList {
		tokenMetadata, err := s.GetCoinMetadata(token.CoinType)
		if err != nil {
			return nil, err
		}

		native := false
		if token.CoinType == "0x2::sui::SUI" {
			native = true
		}

		assetToken := response.AssetToken{
			Name:    tokenMetadata.Result.Name,
			Symbol:  tokenMetadata.Result.Symbol,
			Decimal: int64(tokenMetadata.Result.Decimals),
			Price:   0,
			Native:  native,
			Chain: response.AssetTokenChain{
				Name: "sui",
			},
		}

		walletAssetToken := response.WalletAssetData{
			ChainID:        9996,
			ContractName:   assetToken.Name,
			ContractSymbol: assetToken.Symbol,
			AssetBalance:   util.CalculateTokenBalance(token.TotalBalance, int(assetToken.Decimal)),
			UsdBalance:     0,
			Token:          assetToken,
			Amount:         token.TotalBalance,
		}
		walletAssetList = append(walletAssetList, walletAssetToken)
	}
	return walletAssetList, nil
}
