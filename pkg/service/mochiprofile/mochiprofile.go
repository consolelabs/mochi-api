package mochiprofile

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/util"
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

func (m *MochiProfile) GetByDiscordID(discordID string, noFetchAmount bool) (*GetProfileResponse, error) {
	url := fmt.Sprintf("%s/api/v1/profiles/get-by-discord/%s?no-fetch-amount=%v", m.config.MochiProfileServerHost, discordID, noFetchAmount)
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

	res := &GetProfileResponse{}
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

func (m *MochiProfile) GetByID(profileID string) (*GetProfileResponse, error) {
	url := fmt.Sprintf("%s/api/v1/profiles/%s", m.config.MochiProfileServerHost, profileID)

	res := GetProfileResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &res,
		Headers:   map[string]string{"Content-Type": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[mochiprofile.GetByID] util.SendRequest() failed: %v", err)
	}
	return &res, nil
}
