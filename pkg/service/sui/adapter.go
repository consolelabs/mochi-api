package sui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/sentrygo"
	"github.com/defipod/mochi/pkg/util"
)

func (s *SuiService) doCacheBalance(address string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", suiBalanceKey, strings.ToLower(address)))
}

func (s *SuiService) doCacheCoinMetadata(coinType string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", suiCoinMetadataKey, strings.ToLower(coinType)))
}

func (s *SuiService) doCacheAddressAssets(address string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", suiAddressAssetsKey, strings.ToLower(address)))
}

func (s *SuiService) doCacheTransactionBlock(txnHash string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", suiTransactionBlockKey, strings.ToLower(txnHash)))
}

func (s *SuiService) doCacheAddressTxn(address string) (string, error) {
	return s.cache.GetString(fmt.Sprintf("%s-%s", suiAddressTxnKey, strings.ToLower(address)))
}

func (s *SuiService) doNetworkBalance(address string) (*response.SuiAllBalance, error) {
	var client = &http.Client{}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "suix_getAllBalances",
		"params":  []string{address},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	jsonBody := bytes.NewBuffer(body)

	request, err := http.NewRequest("POST", s.config.Sui.Rpc, jsonBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		s.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Sui - doNetWorkBalance failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"address": address,
			},
		})
		return nil, err
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &response.SuiAllBalance{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}

	// cache sui-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.cache.Set(suiBalanceKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (s *SuiService) doNetworkCoinMetadata(coinType string) (*response.SuiCoinMetadata, error) {
	var client = &http.Client{}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "suix_getCoinMetadata",
		"params":  []string{coinType},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	jsonBody := bytes.NewBuffer(body)
	request, err := http.NewRequest("POST", s.config.Sui.Rpc, jsonBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		s.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Sui - doNetWorkCoinMetadaata failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"coinType": coinType,
			},
		})
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &response.SuiCoinMetadata{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}

	// cache sui-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.cache.Set(suiCoinMetadataKey+"-"+strings.ToLower(coinType), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (s *SuiService) doNetworkAddressAssets(address string) ([]response.WalletAssetData, error) {
	walletAssetList := make([]response.WalletAssetData, 0)
	allBalanceRes, err := s.GetBalance(address)
	if err != nil {
		return []response.WalletAssetData{}, err
	}

	balanceTokenList := allBalanceRes.Result

	for _, token := range balanceTokenList {
		tokenMetadata, err := s.GetCoinMetadata(token.CoinType)
		if err != nil {
			return []response.WalletAssetData{}, err
		}

		native := false
		if token.CoinType == "0x2::sui::SUI" {
			native = true
		}

		assetToken := response.AssetToken{
			Name:    tokenMetadata.Result.Name,
			Symbol:  tokenMetadata.Result.Symbol,
			Decimal: int64(tokenMetadata.Result.Decimals),
			Price:   0,
			Native:  native,
			Chain: response.AssetTokenChain{
				Name:      "sui",
				ShortName: "sui",
			},
		}

		walletAssetToken := response.WalletAssetData{
			ChainID:        9996,
			ContractName:   assetToken.Name,
			ContractSymbol: assetToken.Symbol,
			AssetBalance:   util.CalculateTokenBalance(token.TotalBalance, int(assetToken.Decimal)),
			UsdBalance:     0,
			Token:          assetToken,
			Amount:         token.TotalBalance,
		}
		walletAssetList = append(walletAssetList, walletAssetToken)
	}

	// cache sui-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&walletAssetList)
	s.cache.Set(suiAddressAssetsKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return walletAssetList, nil
}

func (s *SuiService) doNetworkTransactionBlock(address string) (*response.SuiTransactionBlock, error) {
	var client = &http.Client{}

	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      "4",
		"method":  "suix_queryTransactionBlocks",
		"params": []interface{}{
			map[string]interface{}{
				"filter": map[string]interface{}{
					"ToAddress": address,
				},
				"options": map[string]interface{}{
					"showBalanceChanges": true,
					"showObjectChanges":  true,
				},
			},
			nil,
			10,
			true,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	jsonBody := bytes.NewBuffer(body)
	request, err := http.NewRequest("POST", fmt.Sprintf("%s", s.config.Sui.Rpc), jsonBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		s.sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
			Message: fmt.Sprintf("[API mochi] - Sui - doNetWorkTransactionBlock failed %v", err),
			Tags:    sentryTags,
			Extra: map[string]interface{}{
				"address": address,
			},
		})
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &response.SuiTransactionBlock{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}

	// cache sui-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&res)
	s.cache.Set(suiTransactionBlockKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return res, nil
}

func (s *SuiService) doNetworkAddressTxn(address string) ([]response.WalletTransactionData, error) {
	walletTransactionDataList := make([]response.WalletTransactionData, 0)
	transactionBlockList, err := s.GetTransactionBlock(address)
	if err != nil {
		return []response.WalletTransactionData{}, err
	}

	for _, transactionBlock := range transactionBlockList.Result.Data {
		transactionData := response.WalletTransactionData{}
		for _, object := range transactionBlock.ObjectChanges {
			if object.Type == "created" && (object.Sender == address || object.Owner.AddressOwner == address) && object.Sender != object.Owner.AddressOwner {
				actions := make([]response.WalletTransactionAction, 0)
				transactionData.HasTransfer = true
				transactionData.ChainID = 9996
				transactionData.TxHash = transactionBlock.Digest
				transactionData.ScanBaseUrl = "https://suiexplorer.com"
				transactionData.Successful = true

				timeTxn, err := strconv.Atoi(transactionBlock.TimestampMs)
				if err != nil {
					return []response.WalletTransactionData{}, err
				}

				transactionData.SignedAt = time.UnixMilli(int64(timeTxn))
				for _, balance := range transactionBlock.BalanceChanges {
					action := response.WalletTransactionAction{}
					if balance.Owner.AddressOwner == object.Owner.AddressOwner {
						action.From = object.Sender
						action.To = object.Owner.AddressOwner
						tokenMetadata, err := s.GetCoinMetadata(balance.CoinType)
						if err != nil {
							return []response.WalletTransactionData{}, err
						}

						native := false
						if balance.CoinType == "0x2::sui::SUI" {
							native = true
						}

						action.Name = tokenMetadata.Result.Name
						action.NativeTransfer = native
						action.Contract = &response.ContractMetadata{
							Name:    tokenMetadata.Result.Name,
							Address: balance.CoinType,
							Symbol:  tokenMetadata.Result.Symbol,
						}
						action.Unit = tokenMetadata.Result.Symbol
						action.Amount = util.CalculateTokenBalance(balance.Amount, tokenMetadata.Result.Decimals)
						actions = append(actions, action)
					}
				}
				transactionData.Actions = actions
				if len(actions) == 0 {
					transactionData.Actions = nil
				}

				walletTransactionDataList = append(walletTransactionDataList, transactionData)
				break
			}
		}
	}

	// cache sui-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&walletTransactionDataList)
	s.cache.Set(suiAddressTxnKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return walletTransactionDataList, nil
}
