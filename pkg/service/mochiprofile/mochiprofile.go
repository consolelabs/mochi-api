package mochiprofile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model/errors"
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
	url := fmt.Sprintf("%s/api/v1/profiles/get-by-discord/%s?no_fetch_amount=%v", m.config.MochiProfileServerHost, discordID, noFetchAmount)
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

func (m *MochiProfile) GetByID(profileID, authorization string) (*GetProfileResponse, error) {
	url := fmt.Sprintf("%s/api/v1/profiles/%s", m.config.MochiProfileServerHost, profileID)

	res := GetProfileResponse{}
	req := util.SendRequestQuery{
		URL:      url,
		Response: &res,
		Headers:  map[string]string{"Content-Type": "application/json", "Authorization": "Bearer " + authorization},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[mochiprofile.GetByID] util.SendRequest() failed: %v", err)
	}
	return &res, nil
}

func (m *MochiProfile) GetAllEvmAccount() (res []*EvmAssociatedAccount, err error) {
	url := fmt.Sprintf("%s/api/v1/associated-accounts/get-evm-account", m.config.MochiProfileServerHost)

	req := util.SendRequestQuery{
		URL:      url,
		Response: &res,
		Headers:  map[string]string{"Content-Type": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[mochiprofile.GetByID] util.SendRequest() failed: %v", err)
	}
	return res, nil
}

type AssociateDexRequest struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

func (m *MochiProfile) AssociateDex(profileId, platform, apiKey, apiSecret string) error {
	body, err := json.Marshal(AssociateDexRequest{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	})
	if err != nil {
		return err
	}

	jsonBody := bytes.NewBuffer(body)

	url := fmt.Sprintf("%s/api/v1/profiles/%s/accounts/connect-dex/%s", m.config.MochiProfileServerHost, profileId, platform)
	request, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.ErrProfile
	}

	return nil
}

func (m *MochiProfile) UnlinkDex(profileId, platform string) error {
	url := fmt.Sprintf("%s/api/v1/profiles/%s/accounts/disconnect-dex/%s", m.config.MochiProfileServerHost, profileId, platform)
	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.ErrProfile
	}

	return nil
}

func (m *MochiProfile) GetOnboardingStatus(profileId string) (res *OnboardingStatusResponse, err error) {
	url := fmt.Sprintf("%s/api/v1/profiles/%s/onboarding-status", m.config.MochiProfileServerHost, profileId)
	req := util.SendRequestQuery{
		URL:      url,
		Response: &res,
		Headers:  map[string]string{"Content-Type": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[mochiprofile.GetOnboardingStatus] util.SendRequest() failed: %v", err)
	}
	return res, nil
}

func (m *MochiProfile) MarkUserDidOnboarding(profileId string) error {
	url := fmt.Sprintf("%s/api/v1/profiles/%s/onboarding-status", m.config.MochiProfileServerHost, profileId)
	request, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.ErrProfile
	}

	return nil
}

func (m *MochiProfile) GetByTelegramID(telegramID string, noFetchAmount bool) (*GetProfileResponse, error) {
	url := fmt.Sprintf("%s/api/v1/profiles/get-by-telegram/%s?no_fetch_amount=%v", m.config.MochiProfileServerHost, telegramID, noFetchAmount)
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

func (m *MochiProfile) GetProfileActivities(profileID string) (any, error) {
	url := fmt.Sprintf("%s/api/v1/profiles/%s/activities?actions=9|10&page=0&size=12&status=new", m.config.MochiProfileServerHost, profileID)

	res := GetProfileResponse{}
	req := util.SendRequestQuery{
		URL:     url,
		Headers: map[string]string{"Content-Type": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("util.SendRequest() failed: %v", err)
	}
	return &res, nil
}

func (m *MochiProfile) GetByIds(profileIds []string) ([]GetProfileResponse, error) {
	payload, err := json.Marshal(struct {
		Ids []string `json:"ids"`
	}{Ids: profileIds})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/profiles/list-by-ids", m.config.MochiProfileServerHost)
	res := []GetProfileResponse{}
	req := util.SendRequestQuery{
		URL:      url,
		Method:   "POST",
		Headers:  map[string]string{"Content-Type": "application/json"},
		Body:     bytes.NewBuffer(payload),
		Response: &res,
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("util.SendRequest() failed: %v", err)
	}
	return res, nil
}

// maximize 50 id per request
func (m *MochiProfile) GetByDiscordIds(discordIds []string) ([]GetProfileResponse, error) {
	payload, err := json.Marshal(struct {
		Ids []string `json:"ids"`
	}{Ids: discordIds})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/v1/profiles/list-by-discord-ids", m.config.MochiProfileServerHost)
	res := []GetProfileResponse{}
	req := util.SendRequestQuery{
		URL:      url,
		Method:   "POST",
		Headers:  map[string]string{"Content-Type": "application/json"},
		Body:     bytes.NewBuffer(payload),
		Response: &res,
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("util.SendRequest() failed: %v", err)
	}
	return res, nil
}
