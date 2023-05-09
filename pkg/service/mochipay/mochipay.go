package mochipay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
)

var supportedPlatforms = []string{
	"discord",
	"evm-chain",
	"sol-chain",
}

type MochiPay struct {
	config *config.Config
	logger logger.Logger
}

func NewService(cfg *config.Config, l logger.Logger) Service {
	return &MochiPay{
		config: cfg,
		logger: l,
	}
}

func (m *MochiPay) SwapMochiPay(req request.KyberSwapRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	jsonBody := bytes.NewBuffer(payload)

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/swap", m.config.MochiPayServerHost)
	request, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		errBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		errResponse := &ErrorResponse{}
		err = json.Unmarshal(errBody, &errResponse)
		if err != nil {
			return err
		}

		err = fmt.Errorf(errResponse.Msg)
		return err
	}

	return nil
}

func (m *MochiPay) TransferVaultMochiPay(req request.MochiPayVaultRequest) (*VaultResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	jsonBody := bytes.NewBuffer(payload)

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/vault/transfer", m.config.MochiPayServerHost)
	request, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		errResponse := &ErrorResponse{}
		err = json.Unmarshal(errBody, &errResponse)
		if err != nil {
			return nil, err
		}

		err = fmt.Errorf(errResponse.Msg)
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &VaultResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return res, nil
}

func (m *MochiPay) GetBalance(profileId, token, chainId string) (*GetBalanceDataResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/mochi-wallet/%s/balances/%s/%s", m.config.MochiPayServerHost, profileId, token, chainId)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		errResponse := &ErrorResponse{}
		err = json.Unmarshal(errBody, &errResponse)
		if err != nil {
			return nil, err
		}

		err = fmt.Errorf(errResponse.Msg)
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &GetBalanceDataResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return res, nil
}

func (m *MochiPay) Transfer(req request.MochiPayTransferRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	jsonBody := bytes.NewBuffer(payload)

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/transfer", m.config.MochiPayServerHost)
	request, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		errBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		errResponse := &ErrorResponse{}
		err = json.Unmarshal(errBody, &errResponse)
		if err != nil {
			return err
		}

		err = fmt.Errorf(errResponse.Msg)
		return err
	}
	return nil
}

func (m *MochiPay) CreateToken(req CreateTokenRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	jsonBody := bytes.NewBuffer(payload)

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/%s/tokens", m.config.MochiPayServerHost, req.ChainId)
	request, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		errBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		errResponse := &ErrorResponse{}
		err = json.Unmarshal(errBody, &errResponse)
		if err != nil {
			return err
		}

		err = fmt.Errorf(errResponse.Msg)
		return err
	}

	return nil
}

func (m *MochiPay) ListTokens(symbol string) ([]Token, error) {
	url := fmt.Sprintf("%s/api/v1/tokens", m.config.MochiPayServerHost)
	if symbol != "" {
		url += "?symbol=" + symbol
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &ListTokensResponse{}
	err = json.Unmarshal(responseBody, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (m *MochiPay) GetToken(symbol, chainId string) (*Token, error) {
	url := fmt.Sprintf("%s/api/v1/%s/tokens/%s", m.config.MochiPayServerHost, chainId, symbol)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &GetTokenResponse{}
	err = json.Unmarshal(responseBody, res)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}
