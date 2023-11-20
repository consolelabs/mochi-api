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
	"github.com/defipod/mochi/pkg/util"
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

func (m *MochiPay) SwapMochiPay(req request.MochiPaySwapRequest) error {
	payload, err := json.Marshal(req)
	if err != nil {
		return err
	}

	jsonBody := bytes.NewBuffer(payload)

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/mochi-wallet/swap", m.config.MochiPayServerHost)
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

func (m *MochiPay) GetListChains() (*GetChainDataResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/chains", m.config.MochiPayServerHost)
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

	res := &GetChainDataResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return res, nil
}

func (m *MochiPay) GetListBalances(profileId string) (*GetBalanceDataResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/mochi-wallet/%s/balances", m.config.MochiPayServerHost, profileId)
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

func (m *MochiPay) Transfer(req request.MochiPayTransferRequest) (*TipResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	jsonBody := bytes.NewBuffer(payload)

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/transfer", m.config.MochiPayServerHost)
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

	res := &TipResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return res, nil
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

func (m *MochiPay) CreateBatchToken(req CreateBatchTokenRequest) ([]Token, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	jsonBody := bytes.NewBuffer(payload)

	client := &http.Client{}
	url := fmt.Sprintf("%s/api/v1/tokens", m.config.MochiPayServerHost)
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

func (m *MochiPay) GetTokenByProperties(req TokenProperties) (*Token, error) {
	url := fmt.Sprintf("%s/api/v1/%s/tokens?address=%s", m.config.MochiPayServerHost, req.ChainId, req.Address)
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

func (m *MochiPay) TransferV2(req TransferV2Request) (*TransferV2Response, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res := &TransferV2Response{}
	status, err := util.SendRequest(util.SendRequestQuery{
		URL:       fmt.Sprintf("%s/api/v2/transfer", m.config.MochiPayServerHost),
		Body:      bytes.NewBuffer(payload),
		Method:    "POST",
		ParseForm: res,
	})
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("transfer failed with status %d", status)
	}

	return res, nil
}

func (m *MochiPay) ApplicationTransfer(req ApplicationTransferRequest) (*ApplicationTransferResponse, error) {
	payload, err := json.Marshal(req.Metadata)
	if err != nil {
		return nil, err
	}

	res := &ApplicationTransferResponse{}
	status, err := util.SendRequest(util.SendRequestQuery{
		URL:       fmt.Sprintf("%s/api/v1/applications/%s/transfer", m.config.MochiPayServerHost, req.AppId),
		Body:      bytes.NewBuffer(payload),
		Method:    "POST",
		ParseForm: res,
		Headers: map[string]string{
			"x-application": req.Header.Application,
			"x-message":     req.Header.Message,
			"x-signature":   req.Header.Signature,
		},
	})
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("transfer failed with status %d", status)
	}

	return res, nil
}

func (m *MochiPay) GetProfileCustodialWallets(profileID string) (any, error) {
	status, err := util.SendRequest(util.SendRequestQuery{
		URL:    fmt.Sprintf("%s/api/v1/in-app-wallets/get-by-profile/%s", m.config.MochiPayServerHost, profileID),
		Method: "POST",
		// ParseForm: res,
	})
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("util.SendRequest() failed with status %d", status)
	}

	return nil, nil
}

func (m *MochiPay) GetProfileKrystalEarnBalances(profileID string) (any, error) {
	status, err := util.SendRequest(util.SendRequestQuery{
		URL: fmt.Sprintf("%s/api/v1/earns/krystal/earn-balances?profile_id=%s", m.config.MochiPayServerHost, profileID),
		// ParseForm: res,
	})
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("util.SendRequest() failed with status %d", status)
	}

	return nil, nil
}
