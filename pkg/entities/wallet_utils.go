package entities

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type BinanceAsset struct {
	UserAsset         []response.BinanceUserAssetResponse `json:"userAsset"`
	FundingAsset      []response.BinanceUserAssetResponse `json:"fundingAsset"`
	StakingProductPos []response.BinanceUserAssetResponse `json:"stakingProductPos"`
	LendingAccount    []response.BinanceUserAssetResponse `json:"lendingAccount"`
}

func containsAsset(fundingAsset []response.BinanceUserAssetResponse, userAssetSymbol string) bool {
	for _, fAsset := range fundingAsset {
		if fAsset.Asset == userAssetSymbol {
			return true
		}
	}
	return false
}

func mergeAsset(userAsset, fundingAsset []response.BinanceUserAssetResponse) []response.BinanceUserAssetResponse {
	for _, uAsset := range userAsset {
		if containsAsset(fundingAsset, uAsset.Asset) {
			for i, itm := range fundingAsset {
				fAssetBtcValuation, err := strconv.ParseFloat(fundingAsset[i].BtcValuation, 64)
				if err != nil {
					continue
				}

				uAssetBtcValudation, err := strconv.ParseFloat(uAsset.BtcValuation, 64)
				if err != nil {
					continue
				}

				fAssetFree, err := strconv.ParseFloat(fundingAsset[i].Free, 64)
				if err != nil {
					continue
				}

				uAssetFree, err := strconv.ParseFloat(uAsset.Free, 64)
				if err != nil {
					continue
				}
				if itm.BtcValuation == "0" && itm.Free != "0" {
					fAssetBtcValuation = fAssetFree * uAssetBtcValudation / uAssetFree
					fundingAsset[i].BtcValuation = fmt.Sprint(fAssetBtcValuation)
				}

				if fundingAsset[i].Asset == uAsset.Asset {
					fundingAsset[i].BtcValuation = fmt.Sprint(fAssetBtcValuation + uAssetBtcValudation)
					fundingAsset[i].Free = fmt.Sprint(fAssetFree + uAssetFree)
					break
				}
			}
		} else {
			fundingAsset = append(fundingAsset, uAsset)
		}

	}

	return fundingAsset
}

func (e *Entity) GetStakingProduct(profileId, apiKey, apiSecret string) (res []response.BinanceUserAssetResponse, err error) {
	// redis cache
	value, err := e.cache.HashGet(fmt.Sprintf("binance-staking-position-%s-%s", profileId, apiKey))
	if err != nil {
		e.log.Error(err, "[entities.GetStakingProduct] Failed to get cache user data binance")
		return nil, err
	}

	if len(value) == 0 {
		pos, err := e.svc.Binance.GetStakingProductPosition(apiKey, apiSecret)
		if err != nil {
			e.log.Error(err, "[entities.GetStakingProduct] Failed to get staking product position")
			return nil, err
		}

		for _, p := range pos {
			amount, err := strconv.ParseFloat(p.Amount, 64)
			if err != nil {
				return nil, err
			}

			rewardAmt, err := strconv.ParseFloat(p.RewardAmt, 64)
			if err != nil {
				return nil, err
			}

			res = append(res, response.BinanceUserAssetResponse{
				Asset:        p.Asset,
				Free:         fmt.Sprint(amount + rewardAmt),
				BtcValuation: "0",
			})
		}

		tmp, _ := json.Marshal(res)
		encodeData := map[string]string{
			"data": string(tmp),
		}

		err = e.cache.HashSet(fmt.Sprintf("binance-staking-position-%s-%s", profileId, apiKey), encodeData, 30*time.Minute)
		if err != nil {
			e.log.Error(err, "Failed to set cache data wallet")
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(value["data"]), &res)
		if err != nil {
			return nil, err
		}
	}

	return res, err
}

func (e *Entity) GetLendingAccount(profileId, apiKey, apiSecret string) (res []response.BinanceUserAssetResponse, err error) {
	// redis cache
	value, err := e.cache.HashGet(fmt.Sprintf("binance-lending-account-%s-%s", profileId, apiKey))
	if err != nil {
		e.log.Error(err, "[entities.GetLendingAccount] Failed to get cache user data binance")
		return nil, err
	}

	if len(value) == 0 {
		lendingAcc, err := e.svc.Binance.GetLendingAccount(apiKey, apiSecret)
		if err != nil {
			e.log.Error(err, "[entities.GetLendingAccount] Failed to get lending account")
			return nil, err
		}

		for _, l := range lendingAcc.PositionAmountVos {
			res = append(res, response.BinanceUserAssetResponse{
				Asset:        l.Asset,
				Free:         l.Amount,
				BtcValuation: l.AmountInBTC,
			})
		}

		tmp, _ := json.Marshal(res)
		encodeData := map[string]string{
			"data": string(tmp),
		}

		err = e.cache.HashSet(fmt.Sprintf("binance-lending-account-%s-%s", profileId, apiKey), encodeData, 30*time.Minute)
		if err != nil {
			e.log.Error(err, "Failed to set cache data wallet")
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(value["data"]), &res)
		if err != nil {
			return nil, err
		}
	}

	return res, err
}

func (e *Entity) SummarizeAsset(profileId, apiKey, apiSecret string) (finalAsset []response.BinanceUserAssetResponse, err error) {
	userAsset, err := e.svc.Binance.GetUserAsset(apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId, "apiKey": apiKey, "apiSecret": apiSecret}).Error(err, "[entities.SumarizeBinanceAsset] Failed to get user asset binance")
		return nil, err
	}

	// get funding asset
	fundingAsset, err := e.svc.Binance.GetFundingAsset(apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId, "apiKey": apiKey, "apiSecret": apiSecret}).Error(err, "[entities.GetBinanceAssets] Failed to get binance funding asset")
		return nil, err
	}

	// get staking position asset
	pos, err := e.GetStakingProduct(profileId, apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId, "apiKey": apiKey, "apiSecret": apiSecret}).Error(err, "[entities.GetBinanceAssets] Failed to get binance staking position asset")
		return nil, err
	}

	// get lending acc asset
	lendingAcc, err := e.GetLendingAccount(profileId, apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId, "apiKey": apiKey, "apiSecret": apiSecret}).Error(err, "[entities.GetBinanceAssets] Failed to get binance lending account asset")
		return nil, err
	}

	// merge 2 list asset
	userFundingAsset := mergeAsset(userAsset, fundingAsset)
	userFundingPosAsset := mergeAsset(userFundingAsset, pos)
	finalAsset = mergeAsset(userFundingPosAsset, lendingAcc)

	return finalAsset, nil
}

func (e *Entity) SummarizeFundingAsset(profileId, apiKey, apiSecret string) (userFundingAsset []response.BinanceUserAssetResponse, err error) {
	userAsset, err := e.svc.Binance.GetUserAsset(apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId, "apiKey": apiKey, "apiSecret": apiSecret}).Error(err, "[entities.SumarizeBinanceAsset] Failed to get user asset binance")
		return nil, err
	}

	// get funding asset
	fundingAsset, err := e.svc.Binance.GetFundingAsset(apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId, "apiKey": apiKey, "apiSecret": apiSecret}).Error(err, "[entities.GetBinanceAssets] Failed to get binance funding asset")
		return nil, err
	}

	userFundingAsset = mergeAsset(userAsset, fundingAsset)

	return userFundingAsset, nil
}

func (e *Entity) SummarizeEarnAsset(profileId, apiKey, apiSecret string) (earnAsset []response.BinanceUserAssetResponse, err error) {
	// get staking position asset
	pos, err := e.GetStakingProduct(profileId, apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId, "apiKey": apiKey, "apiSecret": apiSecret}).Error(err, "[entities.GetBinanceAssets] Failed to get binance staking position asset")
		return nil, err
	}

	// get lending acc asset
	lendingAcc, err := e.GetLendingAccount(profileId, apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"profileId": profileId, "apiKey": apiKey, "apiSecret": apiSecret}).Error(err, "[entities.GetBinanceAssets] Failed to get binance lending account asset")
		return nil, err
	}

	earnAsset = mergeAsset(pos, lendingAcc)

	return earnAsset, nil
}

func (e *Entity) FormatAsset(assets []response.BinanceUserAssetResponse) ([]response.WalletAssetData, error) {
	// btc price
	btcPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{"bitcoin"}, "usd")
	if err != nil {
		e.log.Error(err, "[entities.SumarizeBinanceAsset] Failed to get btc price")
		return nil, err
	}

	resp := make([]response.WalletAssetData, 0)
	for _, asset := range assets {
		// filter dust
		if asset.Free == "0" {
			continue
		}

		assetValue, err := strconv.ParseFloat(asset.Free, 64)
		if err != nil {
			e.log.Error(err, "[entities.SumarizeBinanceAsset] Failed to parse asset value")
			return nil, err
		}

		btcValuation, err := strconv.ParseFloat(asset.BtcValuation, 64)
		if err != nil {
			e.log.Error(err, "[entities.SumarizeBinanceAsset] Failed to parse asset value")
			return nil, err
		}

		resp = append(resp, response.WalletAssetData{
			AssetBalance: assetValue,
			Amount:       util.FloatToString(fmt.Sprint(assetValue), 18),
			Token: response.AssetToken{
				Symbol:  asset.Asset,
				Decimal: 18,
				Price:   btcValuation * btcPrice["bitcoin"] / assetValue,
			},
		})
	}

	return resp, nil
}
