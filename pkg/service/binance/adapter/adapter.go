package bapdater

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/defipod/mochi/pkg/response"
	butils "github.com/defipod/mochi/pkg/service/binance/utils"
)

var (
	url = "https://api.binance.com"
)

func GetApiKeyPermission(apiKey, apiSecret string) (permission *response.BinanceApiKeyPermissionResponse, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", url+"/sapi/v1/account/apiRestrictions?"+queryString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, apiKey, 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode response json
	err = json.NewDecoder(resp.Body).Decode(&permission)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

func GetUserAsset(apiKey, apiSecret string) (assets []response.BinanceUserAssetResponse, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("POST", url+"/sapi/v3/asset/getUserAsset?"+queryString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, apiKey, 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// decode response json
	err = json.Unmarshal(resBody, &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func GetFundingAsset(apiKey, apiSecret string) (assets []response.BinanceUserAssetResponse, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("POST", url+"/sapi/v1/asset/get-funding-asset?"+queryString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, apiKey, 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// decode response json
	err = json.Unmarshal(resBody, &assets)
	if err != nil {
		return nil, err
	}

	return assets, nil
}
