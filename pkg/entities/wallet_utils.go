package entities

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/covalent"
	"github.com/defipod/mochi/pkg/service/mochipay"
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
					fundingAsset[i].DetailStaking = uAsset.DetailStaking

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

		rawData, _ := json.Marshal(p)

		res = append(res, response.BinanceUserAssetResponse{
			Asset:        p.Asset,
			Free:         fmt.Sprint(amount + rewardAmt),
			BtcValuation: "0",
			DetailString: string(rawData),
		})
	}

	for i, r := range res {
		var detailStaking *response.BinanceStakingProductPosition
		err = json.Unmarshal([]byte(r.DetailString), &detailStaking)
		if err != nil {
			continue
		}

		res[i].DetailStaking = detailStaking
	}

	return res, err
}

func (e *Entity) GetLendingAccount(profileId, apiKey, apiSecret string) (res []response.BinanceUserAssetResponse, err error) {
	lendingAcc, err := e.svc.Binance.GetLendingAccount(apiKey, apiSecret)
	if err != nil {
		e.log.Error(err, "[entities.GetLendingAccount] Failed to get lending account")
		return nil, err
	}

	for _, l := range lendingAcc.PositionAmountVos {
		rawData, _ := json.Marshal(l)

		res = append(res, response.BinanceUserAssetResponse{
			Asset:        l.Asset,
			Free:         l.Amount,
			BtcValuation: l.AmountInBTC,
			DetailString: string(rawData),
		})
	}

	for i, r := range res {
		var detailLending *response.BinancePositionAmountVos
		err = json.Unmarshal([]byte(r.DetailString), &detailLending)
		if err != nil {
			continue
		}

		res[i].DetailLending = detailLending
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

		assetFree, err := strconv.ParseFloat(asset.Free, 64)
		if err != nil {
			e.log.Error(err, "[entities.SumarizeBinanceAsset] Failed to parse asset value")
			return nil, err
		}

		assetLocked, err := strconv.ParseFloat(asset.Locked, 64)
		if err != nil {
			e.log.Error(err, "[entities.SumarizeBinanceAsset] Failed to parse asset value")
			return nil, err
		}

		assetValue := assetFree + assetLocked

		btcValuation, err := strconv.ParseFloat(asset.BtcValuation, 64)
		if err != nil {
			e.log.Error(err, "[entities.SumarizeBinanceAsset] Failed to parse asset value")
			return nil, err
		}

		usdBal := btcValuation * btcPrice["bitcoin"]
		if usdBal <= 0 {
			binancePrice, err := e.svc.Binance.GetPrice(asset.Asset + "USDT")
			if err != nil {
				e.log.Error(err, "[entities.SumarizeBinanceAsset] Failed to get binance price")
				return nil, err
			}

			assetUsdPrice, _ := strconv.ParseFloat(binancePrice.Price, 64)
			usdBal = assetUsdPrice * assetValue
		}

		itm := response.WalletAssetData{
			AssetBalance: assetValue,
			Amount:       util.FloatToString(fmt.Sprint(assetValue), 18),
			Token: response.AssetToken{
				Symbol:  asset.Asset,
				Decimal: 18,
				Price:   btcValuation * btcPrice["bitcoin"] / assetValue,
			},
			UsdBalance: usdBal,
		}

		if asset.DetailLending != nil && asset.DetailLending.Amount != "0" {
			itm.DetailLending = asset.DetailLending
		}

		if asset.DetailStaking != nil && asset.DetailStaking.Amount != "0" {
			itm.DetailStaking = asset.DetailStaking
		}

		resp = append(resp, itm)
	}

	return resp, nil
}

func containsWalletAsset(wallet []response.WalletAssetData, userAssetSymbol, userAssetName string, userAssetChainId int) bool {
	for _, w := range wallet {
		if w.ContractSymbol == userAssetSymbol && w.ContractName == userAssetName && w.ChainID == userAssetChainId && w.ContractSymbol != "" && w.ContractName != "" && w.ChainID != 0 {
			return true
		}
	}
	return false
}

func mergeWalletAsset(firstWallet, secondWallet []response.WalletAssetData) []response.WalletAssetData {
	for _, fWallet := range firstWallet {
		if containsWalletAsset(secondWallet, fWallet.ContractSymbol, fWallet.ContractName, fWallet.ChainID) {
			for i, sWallet := range secondWallet {
				if sWallet.ContractSymbol == fWallet.ContractSymbol && sWallet.ContractName == fWallet.ContractName && sWallet.ChainID == fWallet.ChainID {
					sWalletAmount, err := util.StringToBigInt(sWallet.Amount)
					if err != nil {
						continue
					}

					fWalletAmount, err := util.StringToBigInt(fWallet.Amount)
					if err != nil {
						continue
					}

					totalAmount := sWalletAmount.Add(sWalletAmount, fWalletAmount)
					secondWallet[i].AssetBalance = fWallet.AssetBalance + sWallet.AssetBalance
					secondWallet[i].UsdBalance = fWallet.UsdBalance + sWallet.UsdBalance
					secondWallet[i].Amount = totalAmount.String()
				}
			}
		} else {
			secondWallet = append(secondWallet, fWallet)
		}

	}

	return secondWallet
}

func formatOffchainBalance(offchainBalance mochipay.GetBalanceDataResponse) []response.WalletAssetData {
	resp := make([]response.WalletAssetData, 0)
	for _, asset := range offchainBalance.Data {
		chainId, _ := strconv.Atoi(asset.Token.ChainId)
		itm := response.WalletAssetData{
			// AssetBalance: assetValue,
			ChainID:        chainId,
			ContractName:   asset.Token.Name,
			ContractSymbol: asset.Token.Symbol,
			Amount:         asset.Amount,
			Token: response.AssetToken{
				Id:      asset.Token.Id,
				Address: asset.Token.Address,
				Name:    asset.Token.Name,
				Symbol:  asset.Token.Symbol,
				Decimal: asset.Token.Decimal,
				Price:   asset.Token.Price,
				Native:  asset.Token.Native,
				Chain: response.AssetTokenChain{
					Name:      asset.Token.Chain.Name,
					ShortName: asset.Token.Chain.Symbol,
				},
				Icon: asset.Token.Icon,
			},
			UsdBalance: asset.UsdAmount,
		}

		resp = append(resp, itm)
	}
	return resp
}

// add tokens into mochi-pay database
func (e *Entity) enrichDataWalletAsset(assets []response.WalletAssetData) []response.WalletAssetData {
	reqCreateMochiPayTokens := make([]mochipay.CreateTokenRequest, 0)
	for _, asset := range assets {
		reqCreateMochiPayTokens = append(reqCreateMochiPayTokens, mochipay.CreateTokenRequest{
			Name:    asset.Token.Name,
			Symbol:  asset.Token.Symbol,
			Decimal: asset.Token.Decimal,
			ChainId: fmt.Sprint(asset.ChainID),
			Address: asset.Token.Address,
		})
	}

	// create tokens
	tokens, err := e.svc.MochiPay.CreateBatchToken(mochipay.CreateBatchTokenRequest{Tokens: reqCreateMochiPayTokens})
	if err != nil {
		e.log.Error(err, "[entities.enrichDataWalletAsset] Failed to create tokens")
		return assets
	}

	// enrich id into assets
	for i, asset := range assets {
		for _, token := range tokens {
			if strings.EqualFold(token.Symbol, asset.Token.Symbol) && token.ChainId == fmt.Sprint(asset.ChainID) && strings.EqualFold(token.Address, asset.Token.Address) {
				assets[i].Token.Id = token.Id
				icon := token.Icon
				if icon == "" {
					emoji, err := e.GetEmojiByCode(token.Symbol)
					if err != nil {
						e.log.Error(err, "[entities.enrichDataWalletAsset] Failed to get emoji by code")
						icon = ""
					} else {
						icon = emoji.EmojiUrl
					}
				}

				assets[i].Token.Icon = icon
				break
			}
		}
	}

	return assets
}

func (e *Entity) enrichMetadataSolAsset(res covalent.GetTokenBalancesResponse, item covalent.TokenBalanceItem) *response.WalletAssetData {
	chainID := 999
	tokenAddress := item.ContractAddress
	if item.NativeToken {
		tokenAddress = consts.SolAddress
	}

	bal, _ := e.calculateTokenBalance(item, chainID)

	tokenInfo, err := e.svc.GeckoTerminal.GetTokenByAddress(consts.SolChainType, tokenAddress)
	if err != nil {
		e.log.Fields(logger.Fields{"chainID": consts.SolChainType, "address": item.ContractAddress}).Error(err, "[entity.enrichMetadataSolAsset] svc.GeckoTerminal.GetTokenByAddress() failed")
		return nil
	}

	contractName := item.ContractName
	if contractName == "" {
		contractName = tokenInfo.Data.Attributes.Name
	}
	contractSymbol := item.ContractTickerSymbol
	if contractSymbol == "" {
		contractSymbol = tokenInfo.Data.Attributes.Symbol
	}

	price, _ := strconv.ParseFloat(tokenInfo.Data.Attributes.PriceUsd, 32)

	return &response.WalletAssetData{
		ChainID:        chainID,
		ContractName:   contractName,
		ContractSymbol: contractSymbol,
		AssetBalance:   bal,
		UsdBalance:     price * bal,
		Token: response.AssetToken{
			Name:    contractName,
			Symbol:  contractSymbol,
			Address: tokenAddress,
			Decimal: tokenInfo.Data.Attributes.Decimals,
			Price:   price,
			Native:  item.NativeToken,
			Chain: response.AssetTokenChain{
				Name:      res.Data.ChainName,
				ShortName: "sol",
				Type:      "solana",
			},
		},
		Amount: util.FloatToString(fmt.Sprint(bal), tokenInfo.Data.Attributes.Decimals),
	}
}

func (e *Entity) GetAssetStakingSol(item covalent.TokenBalanceItem) *response.WalletAssetData {
	chainID := 999

	stakingToken, err := e.svc.MochiPay.GetStakingTokenMapping(item.ContractTickerSymbol, item.ContractAddress)
	if err != nil {
		e.log.Fields(logger.Fields{"symbol": item.ContractTickerSymbol, "address": item.ContractAddress}).Error(err, "[entities.GetAssetStakingSol] Failed to get staking token mapping")
		return nil
	}

	if stakingToken.Data == nil {
		return nil
	}

	tokenAddress := item.ContractAddress
	if item.NativeToken {
		tokenAddress = consts.SolAddress
	}

	tokenInfo, err := e.svc.GeckoTerminal.GetTokenByAddress(consts.SolChainType, tokenAddress)
	if err != nil {
		e.log.Fields(logger.Fields{"chainID": consts.SolChainType, "address": item.ContractAddress}).Error(err, "[entity.enrichMetadataSolAsset] svc.GeckoTerminal.GetTokenByAddress() failed")
		return nil
	}

	bal, _ := e.calculateTokenBalance(item, chainID)
	price, _ := strconv.ParseFloat(tokenInfo.Data.Attributes.PriceUsd, 32)

	contractName := item.ContractName
	if contractName == "" {
		contractName = tokenInfo.Data.Attributes.Name
	}
	contractSymbol := item.ContractTickerSymbol
	if contractSymbol == "" {
		contractSymbol = tokenInfo.Data.Attributes.Symbol
	}

	return &response.WalletAssetData{
		ChainID:        chainID,
		ContractName:   contractName,
		ContractSymbol: contractSymbol,
		AssetBalance:   bal,
		UsdBalance:     price * bal,
		Token: response.AssetToken{
			Name:    contractName,
			Symbol:  contractSymbol,
			Address: tokenAddress,
			Decimal: tokenInfo.Data.Attributes.Decimals,
			Price:   price,
			Native:  item.NativeToken,
			Chain: response.AssetTokenChain{
				Name:      "Solana",
				ShortName: "sol",
				Type:      "solana",
			},
		},
		Amount: util.FloatToString(fmt.Sprint(bal), tokenInfo.Data.Attributes.Decimals),
	}
}
