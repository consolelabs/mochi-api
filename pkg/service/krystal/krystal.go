package krystal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

type Krystal struct {
	config *config.Config
	logger logger.Logger
	cache  cache.Cache
	sentry sentrygo.Service
}

func NewService(cfg *config.Config, l logger.Logger, cache cache.Cache, sentry sentrygo.Service) Service {
	return &Krystal{
		config: cfg,
		logger: l,
		cache:  cache,
		sentry: sentry,
	}
}

var (
	sentryTags = map[string]string{
		"type": "system",
	}
)

const (
	tokenBalanceKey = "krystal-balance-token"
)

func (k *Krystal) GetBalanceTokenByAddress(address string) (*BalanceTokenResponse, error) {
	k.logger.Debug("start krystal.GetBalanceTokenByAddress()")
	defer k.logger.Debug("end krystal.GetBalanceTokenByAddress()")

	var data BalanceTokenResponse
	// check if data cached

	cached, err := k.doCache(address)
	if err == nil && cached != "" {
		k.logger.Infof("hit cache data krystal-service, address: %s", address)
		go k.doNetwork(address, data)
		return &data, json.Unmarshal([]byte(cached), &data)
	}

	// call network
	return k.doNetwork(address, data)
}

func (k *Krystal) GetEarningOptions(platforms, chainIds, types, statuses, address string) (*GetEarningOptionsResponse, error) {
	resp := &GetEarningOptionsResponse{}
	url := k.config.KrystalBaseUrl + fmt.Sprintf("/all/v1/earning/options?platforms=%s&chainIds=%s&types=%s&statuses=%s&address=%s", platforms, chainIds, types, statuses, address)
	req := util.SendRequestQuery{
		URL:      url,
		Response: resp,
		Headers: map[string]string{
			"accept":              "application/json",
			"x-rate-access-token": k.config.KrystalApiKey,
		},
	}

	cached, err := k.cache.GetString(url)
	if err == nil && cached != "" {
		k.logger.Infof("hit cache data krystal-service, url: %s", url)
		go k.doNetworkGeneric(req, resp)
		return resp, json.Unmarshal([]byte(cached), resp)
	}

	if err := k.doNetworkGeneric(req, resp); err != nil {
		k.logger.Error(err, "[krystal.GetEarningOptions] k.doNetworkGeneric() failed")
		return nil, err
	}

	return resp, nil
}

func (k *Krystal) BuildStakeTx(req BuildStakeTxReq) (*BuildTxResp, error) {
	res, err := k.buildTx("buildStakeTx", req)
	if err != nil {
		k.logger.Fields(logger.Fields{"request": req}).Error(err, "[krystal.BuildStakeTx] k.buildTx() failed")
		return nil, err
	}
	return res, nil
}

func (k *Krystal) BuildUnstakeTx(req BuildUnstakeTxReq) (*BuildTxResp, error) {
	res, err := k.buildTx("buildUnstakeTx", req)
	if err != nil {
		k.logger.Fields(logger.Fields{"request": req}).Error(err, "[krystal.BuildUnstakeTx] k.buildTx() failed")
		return nil, err
	}
	return res, nil
}

func (k *Krystal) doCache(address string) (string, error) {
	return k.cache.GetString(fmt.Sprintf("%s-%s", tokenBalanceKey, strings.ToLower(address)))
}

func (k *Krystal) doNetwork(address string, data BalanceTokenResponse) (*BalanceTokenResponse, error) {
	chainIDs := []int{1, 10, 25, 56, 101, 137, 199, 250, 324, 2000, 5000, 8217, 8453, 42161, 43114, 1313161554}
	chainIDsStr := strings.ReplaceAll(strings.Trim(fmt.Sprint(chainIDs), "[]"), " ", ",")

	url := k.config.KrystalBaseUrl + fmt.Sprintf("/all/v1/balance/token?addresses=ethereum:%s&quoteSymbols=usd&sparkline=false&chainIds=%s", address, chainIDsStr)

	req := util.SendRequestQuery{
		URL:      url,
		Response: &data,
		Headers:  map[string]string{"x-rate-access-token": k.config.KrystalApiKey},
	}

	statusCode, err := util.SendRequest(req)
	if err != nil {
		k.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Krystal - doNetWork failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"url": url,
			},
		})
		k.logger.Fields(
			logger.Fields{
				"address": address,
				"url":     url,
			},
		).Errorf(err, "krystal.GetBalanceTokenByAddress() failed, status code: %d", statusCode)
		return nil, fmt.Errorf("[krystal.GetBalanceTokenByAddress] util.SendRequest() failed: %v", err)
	}

	if statusCode != http.StatusOK {
		k.logger.Infof("krystal.GetBalanceTokenByAddress() failed, status code: %d", statusCode)
		return &data, nil
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&data)
	k.logger.Infof("cache data krystal-service, key: %s", tokenBalanceKey)
	k.cache.Set(tokenBalanceKey+"-"+strings.ToLower(address), string(bytes), 15*time.Minute)

	return &data, nil
}

func (k *Krystal) doNetworkGeneric(req util.SendRequestQuery, response interface{}) error {
	statusCode, err := util.SendRequest(req)
	if err != nil {
		return fmt.Errorf("[krystal.doNetworkGeneric] util.SendRequest() failed: %v", err)
	}

	if statusCode != http.StatusOK {
		err = fmt.Errorf("krystal.doNetworkGeneric() failed, status code: %d", statusCode)
		k.logger.Error(err, "krystal.doNetworkGeneric() failed")
		k.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Krystal - doNetWorkGeneric failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"req": req,
			},
		})
		return err
	}

	// cache data
	// if error occurs -> ignore
	cacheKey := req.URL
	bytes, _ := json.Marshal(&req.Response)
	k.logger.Infof("cache data krystal-service, key: %s", cacheKey)
	k.cache.Set(cacheKey, string(bytes), 7*24*time.Hour)

	return nil

}

func (k *Krystal) buildTx(path string, req interface{}) (*BuildTxResp, error) {
	v, err := json.Marshal(req)
	if err != nil {
		k.logger.Fields(logger.Fields{"request": req}).Error(err, "[krystal.buildTx] json.Marshal() failed")
		return nil, err
	}
	body := bytes.NewBuffer(v)
	res := &BuildTxResp{}
	status, err := util.SendRequest(util.SendRequestQuery{
		URL:    fmt.Sprintf("%s/all/v1/earning/%s", k.config.KrystalBaseUrl, path),
		Method: "POST",
		Headers: map[string]string{
			"Accept":              "application/json",
			"Content-Type":        "application/json",
			"x-rate-access-token": k.config.KrystalApiKey,
		},
		Body:     body,
		Response: res,
	})
	if err != nil {
		k.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Krystal - buildTx failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"req": req,
			},
		})
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", status)
	}
	return res, nil
}
