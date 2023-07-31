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
	url       = "https://api.binance.com"
	futureUrl = "https://fapi.binance.com"
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

func GetStakingProductPosition(apiKey, apiSecret string) (pos []response.BinanceStakingProductPosition, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
		"product":   "STAKING",
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", url+"/sapi/v1/staking/position?"+queryString, nil)
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
	err = json.Unmarshal(resBody, &pos)
	if err != nil {
		return nil, err
	}

	return pos, nil
}

func GetLendingAccount(apiKey, apiSecret string) (lendingAcc *response.BinanceLendingAccount, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", url+"/sapi/v1/lending/union/account?"+queryString, nil)
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
	err = json.Unmarshal(resBody, &lendingAcc)
	if err != nil {
		return nil, err
	}

	return lendingAcc, nil
}

func GetSimpleEarnAccount(apiKey, apiSecret string) (simpleEarn *response.BinanceSimpleEarnAccount, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", url+"/sapi/v1/simple-earn/account?"+queryString, nil)
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
	err = json.Unmarshal(resBody, &simpleEarn)
	if err != nil {
		return nil, err
	}

	return simpleEarn, nil
}

func GetFutureAccountBalance(apiKey, apiSecret string) (fAccountBal []response.BinanceFutureAccountBalance, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", futureUrl+"/fapi/v2/balance?"+queryString, nil)
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
	err = json.Unmarshal(resBody, &fAccountBal)
	if err != nil {
		return nil, err
	}

	return fAccountBal, nil
}
