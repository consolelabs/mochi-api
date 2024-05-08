package badapter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	butils "github.com/defipod/mochi/pkg/service/binance/utils"
)

var (
	url       = "https://api.binance.com"
	futureUrl = "https://fapi.binance.com"
	log       = logger.NewLogrusLogger().Fields(logger.Fields{
		"component": "service.binance.adapter",
	})
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
		"timestamp":        strconv.Itoa(int(time.Now().UnixMilli())),
		"needBtcValuation": "true",
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

	resBody, err := io.ReadAll(resp.Body)
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
		"timestamp":        strconv.Itoa(int(time.Now().UnixMilli())),
		"needBtcValuation": "true",
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

	resBody, err := io.ReadAll(resp.Body)
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
		"product":   "L_DEFI",
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

	resBody, err := io.ReadAll(resp.Body)
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
	req, err := http.NewRequest("GET", url+"/sapi/v1/simple-earn/account?"+queryString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, apiKey, 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
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

	resBody, err := io.ReadAll(resp.Body)
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

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// decode response json
	err = json.Unmarshal(resBody, &fAccountBal)
	if err != nil {
		log.Fields(logger.Fields{"response": string(resBody)}).Error(err, "[GetFutureAccountBalance] failed to parse response")
		// if user has invalid api key or insufficient permission ->
		var errRes BinanceFutureErrorResponse
		json.Unmarshal(resBody, &errRes)
		if errRes.Code == -2015 {
			return nil, nil
		}
		return nil, err
	}

	return fAccountBal, nil
}

func GetFutureAccount(apiKey, apiSecret string) (fAccountBal *response.BinanceFutureAccount, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", futureUrl+"/fapi/v2/account?"+queryString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, apiKey, 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
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

func GetFutureAccountInfo(apiKey, apiSecret string) (fAccountBal []response.BinanceFuturePositionInfo, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", futureUrl+"/fapi/v2/positionRisk?"+queryString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, apiKey, 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
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

func GetTickerPrice(symbol string) (price *response.BinanceApiTickerPriceResponse, err error) {
	// http request
	req, err := http.NewRequest("GET", url+"/api/v3/ticker/price?symbol="+symbol, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, "", 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode response json
	err = json.NewDecoder(resp.Body).Decode(&price)
	if err != nil {
		return nil, err
	}

	return price, nil
}

func GetSpotTransaction(apiKey, apiSecret, symbol, startTime, endTime string) (txs []response.BinanceSpotTransactionResponse, err error) {
	q := map[string]string{
		"timestamp":  strconv.Itoa(int(time.Now().UnixMilli())),
		"startTime":  startTime,
		"endTime":    endTime,
		"symbol":     symbol,
		"limit":      "1000",
		"recvWindow": "59000",
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", url+"/api/v3/allOrders?"+queryString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, apiKey, 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.Header.Get("X-Mbx-Used-Weight-1m") != "" {
		usedWeight1M, err := strconv.Atoi(resp.Header.Get("X-Mbx-Used-Weight-1m"))
		if err != nil || usedWeight1M > 5000 {
			fmt.Printf("err: %+v, %d", err, usedWeight1M)
			time.Sleep(1 * time.Minute)
		}
	}

	// decode response json
	err = json.Unmarshal(resBody, &txs)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
