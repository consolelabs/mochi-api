package mochiprofile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
)

var supportedPlatforms = []string{
	"discord",
	"evm-chain",
	"sol-chain",
}

type MochiProfile struct {
	config *config.Config
	logger logger.Logger
}

func NewService(cfg *config.Config, l logger.Logger) Service {
	return &MochiProfile{
		config: cfg,
		logger: l,
	}
}

func (m *MochiProfile) GetByDiscordID(discordID string) (*GetProfileByDiscordResponse, error) {
	url := fmt.Sprintf("%s/api/v1/profiles/get-by-discord/%s", m.config.MochiProfileServerHost, discordID)
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

	res := &GetProfileByDiscordResponse{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return res, nil
}

func (m *MochiProfile) GetApiKeyByProfileID(profileID string) (*ProfileApiKeyResponse, error) {
	url := fmt.Sprintf("%s/api/v1/api-key/%s", m.config.MochiProfileServerHost, profileID)
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

	if response.StatusCode != http.StatusOK {
		errBody, err := io.ReadAll(response.Body)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &ProfileApiKeyResponseData{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return &res.Data, nil
}

func (m *MochiProfile) CreateProfileApiKey(profileAccessToken string) (*ProfileApiKeyResponse, error) {
	url := fmt.Sprintf("%s/api/v1/api-key/me", m.config.MochiProfileServerHost)
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", profileAccessToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errBody, err := io.ReadAll(response.Body)
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &ProfileApiKeyResponseData{}
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return &res.Data, nil
}
