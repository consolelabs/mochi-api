package entities

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerr "github.com/defipod/mochi/pkg/model/errors"
	userwalletwatchlistitem "github.com/defipod/mochi/pkg/repo/user_wallet_watchlist_item"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/covalent"
)

func (e *Entity) GetTrackingWallets(req request.GetTrackingWalletsRequest) ([]model.UserWalletWatchlistItem, error) {
	wallets, err := e.repo.UserWalletWatchlistItem.List(req.UserID)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": req.UserID}).Error(err, "[entity.GetTrackingWallets] repo.UserWalletWatchlistItem.List() failed")
		return nil, err
	}
	chainIDs := []int{1, 56, 137, 250}
	for i, wallet := range wallets {
		for _, chainID := range chainIDs {
			res, err := e.svc.Covalent.GetHistoricalPortfolio(chainID, wallet.Address, 5)
			if err != nil {
				e.log.Fields(logger.Fields{"chainID": chainID, "addr": wallet.Address}).Error(err, "[entity.GetTrackingWallets] svc.Covalent.GetHistoricalPortfolio() failed")
				return nil, err
			}
			if res.Data.Items == nil || len(res.Data.Items) == 0 {
				continue
			}
			for _, asset := range res.Data.Items {
				if asset.Holdings == nil || len(asset.Holdings) == 0 {
					continue
				}
				latest := asset.Holdings[0]
				if strings.EqualFold(asset.ContractTickerSymbol, "icy") {
					bal, ok := new(big.Float).SetString(latest.Open.Balance)
					if ok {
						parsedBal, _ := bal.Float64()
						latest.Open.Quote = 1.5 * parsedBal / math.Pow10(asset.ContractDecimals)
					}
				}
				wallets[i].NetWorth += latest.Open.Quote
			}
		}
	}
	return wallets, nil
}

func (e *Entity) GetOneWallet(req request.GetOneWalletRequest) (*model.UserWalletWatchlistItem, error) {
	wallet, err := e.repo.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{UserID: req.UserID, Query: req.AliasOrAddress})
	if err != nil {
		e.log.Fields(logger.Fields{"userID": req.UserID}).Error(err, "[entity.GetOneWallet] repo.UserWalletWatchlistItem.GetOne() failed")
		return nil, err
	}
	return wallet, nil
}

func (e *Entity) TrackWallet(req request.TrackWalletRequest) error {
	wallet, err := e.repo.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{UserID: req.UserID, Query: req.Alias})
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TrackWallet] repo.UserWalletWatchlistItem.GetOne() failed")
		return err
	}
	if err != gorm.ErrRecordNotFound && !strings.EqualFold(wallet.Address, req.Address) {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TrackWallet] repo.UserWalletWatchlistItem.GetOne() failed")
		return baseerr.ErrConflict
	}

	err = e.repo.UserWalletWatchlistItem.Create(&model.UserWalletWatchlistItem{
		UserID:  req.UserID,
		Address: req.Address,
		Alias:   req.Alias,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TrackWallet] repo.UserWalletWatchlistItem.Create() failed")
	}
	return err
}

func (e *Entity) UntrackWallet(req request.UntrackWalletRequest) error {
	return e.repo.UserWalletWatchlistItem.Remove(userwalletwatchlistitem.DeleteQuery{
		UserID:  req.UserID,
		Address: req.Address,
		Alias:   req.Alias,
	})
}

func (e *Entity) ListWalletAssets(req request.ListWalletAssetsRequest) ([]response.WalletAssetData, error) {
	chainIDs := []int{1, 56, 137, 250}
	var assets []response.WalletAssetData
	for _, chainID := range chainIDs {
		res, err := e.svc.Covalent.GetHistoricalPortfolio(chainID, req.Address, 5)
		if err != nil {
			e.log.Fields(logger.Fields{"chainID": chainID, "address": req.Address}).Error(err, "[entity.ListWalletAssets] svc.Covalent.GetHistoricalPortfolio() failed")
			return nil, err
		}
		if res.Data.Items == nil || len(res.Data.Items) == 0 {
			continue
		}
		for _, item := range res.Data.Items {
			if item.Holdings == nil || len(item.Holdings) == 0 {
				continue
			}
			latest := item.Holdings[0]
			bal, ok := new(big.Float).SetString(latest.Open.Balance)
			if !ok {
				continue
			}
			parsedBal, _ := bal.Float64()
			assetBal := parsedBal / math.Pow10(item.ContractDecimals)
			if strings.EqualFold(item.ContractTickerSymbol, "icy") {
				latest.Open.Quote = 1.5 * assetBal
			}
			assets = append(assets, response.WalletAssetData{
				ChainID:        chainID,
				ContractName:   item.ContractName,
				ContractSymbol: item.ContractTickerSymbol,
				AssetBalance:   assetBal,
				UsdBalance:     latest.Open.Quote,
			})
		}
	}
	return assets, nil
}

func (e *Entity) ListWalletTxns(req request.ListWalletTransactionsRequest) ([]response.WalletTransactionData, error) {
	chainIDs := []int{1, 56, 137, 250}
	var txns []response.WalletTransactionData
	for _, chainID := range chainIDs {
		res, err := e.svc.Covalent.GetTransactionsByAddress(chainID, req.Address, 5, 5)
		if err != nil {
			e.log.Fields(logger.Fields{"chainID": chainID, "address": req.Address}).Error(err, "[entity.ListWalletTxns] svc.Covalent.GetTransactionsByAddress() failed")
			return nil, err
		}
		if res.Data.Items == nil || len(res.Data.Items) == 0 {
			continue
		}
		for _, item := range res.Data.Items {
			tx := response.WalletTransactionData{
				ChainID:    chainID,
				From:       item.FromAddress,
				To:         item.ToAddress,
				TxHash:     item.TxHash,
				SignedAt:   item.BlockSignedAt,
				Successful: item.Successful,
			}
			//
			if err := e.parseCovalentTxData(item, &tx, req.Address); err != nil {
				e.log.Fields(logger.Fields{"item": item, "tx": tx, "address": req.Address}).Error(err, "[entity.ListWalletTxns] entity.parseCovalentTxData() failed")
			}
			txns = append(txns, tx)
		}
	}
	sort.Slice(txns, func(i, j int) bool {
		return txns[i].SignedAt.Unix() > txns[j].SignedAt.Unix()
	})
	return txns, nil
}

func (e *Entity) parseCovalentTxData(tx covalent.TransactionItemData, res *response.WalletTransactionData, addr string) error {
	chain, _ := e.repo.Chain.GetByID(res.ChainID)
	nativeSymbol := chain.Currency
	scanBaseUrl := strings.Replace(chain.TxBaseURL, "/tx", "", 1)
	res.ScanBaseUrl = scanBaseUrl
	// 1. transfer native token
	if tx.ValueQuote > 0 {
		value, ok := new(big.Float).SetString(tx.Value)
		if !ok {
			err := fmt.Errorf("invalid tx amount %s", tx.Value)
			e.log.Errorf(err, "[getTxDetails] invalid native amount %s", tx.Value)
			return err
		}
		if value.Cmp(big.NewFloat(0)) == 0 {
			err := errors.New("zero tx amount")
			e.log.Errorf(err, "[getTxDetails] zero transaction amount %s", tx.Value)
			return nil
		}
		// res.Type = "transfer_native"
		res.HasTransfer = true
		amount, _ := value.Float64()
		amount /= math.Pow10(18)
		if strings.EqualFold(tx.FromAddress, addr) {
			amount = 0 - amount
		}
		res.Actions = append(res.Actions, response.WalletTransactionAction{
			Amount:         amount,
			Unit:           nativeSymbol,
			NativeTransfer: true,
		})
		return nil
	}

	if tx.LogEvents == nil || len(tx.LogEvents) == 0 {
		return nil
	}

	// 2. transaction with log events
	events := tx.LogEvents
	transferHandlers := map[string]func(string, covalent.LogEvent, *response.WalletTransactionAction){
		"Transfer(indexed address from, indexed address to, uint256 value)":                                                        e.handleErc20Transfer,
		"Transfer(indexed address from, indexed address to, indexed uint256 tokenId)":                                              e.handleErc721Transfer,
		"TransferSingle(indexed address _operator, indexed address _from, indexed address _to, uint256 _id, uint256 _amount)":      e.handleErc1155TransferSingle,
		"TransferBatch(indexed address _operator, indexed address _from, indexed address _to, uint256[] _ids, uint256[] _amounts)": e.handleErc1155TransferBatch,
	}
	var actions []response.WalletTransactionAction
	for _, ev := range events {
		name := ev.Decoded.Name
		signature := ev.Decoded.Signature
		if name == "" || signature == "" {
			continue
		}
		contract := &response.ContractMetadata{
			Address: ev.SenderAddress,
			Symbol:  ev.SenderContractTickerSymbol,
			Name:    ev.SenderName,
		}
		action := &response.WalletTransactionAction{Contract: contract, Name: name, Signature: signature}
		_, isTransfer := transferHandlers[signature]
		if isTransfer {
			res.HasTransfer = isTransfer
		}
		if isTransfer {
			handler, ok := transferHandlers[signature]
			if ok {
				// batch case: 1 action -> 1 token
				if strings.Contains(signature, "Batch") {
					type valueObj struct {
						Value string
					}
					_ids := ev.Decoded.Params[3].Value.([]valueObj)
					_amounts := ev.Decoded.Params[4].Value.([]valueObj)
					for i, _id := range _ids {
						if contract.Name == "" {
							action.Unit = _id.Value
							amt, err := strconv.ParseFloat(_amounts[i].Value, 10)
							if err != nil {
								continue
							}
							action.Amount = amt
						}
					}
				}
				handler(addr, ev, action)
				if !strings.EqualFold(addr, action.From) && !strings.EqualFold(addr, action.To) {
					// not relevant -> exclude action
					continue
				}
			} else {
				// no transfer handler -> exclude action
				continue
			}
		}
		actions = append(actions, *action)
	}
	for _, action := range actions {
		_, isTransfer := transferHandlers[action.Signature]
		if res.HasTransfer {
			if isTransfer {
				res.Actions = append(res.Actions, action)
			}
			continue
		}
		res.Actions = append(res.Actions, action)
	}
	return nil
}

func (*Entity) handleErc20Transfer(address string, ev covalent.LogEvent, action *response.WalletTransactionAction) {
	action.From = ev.Decoded.Params[0].Value.(string)
	action.To = ev.Decoded.Params[1].Value.(string)
	value, _ := new(big.Float).SetString(ev.Decoded.Params[2].Value.(string))
	action.Amount, _ = value.Float64()
	action.Amount = action.Amount / math.Pow10(ev.SenderContractDecimals)
	if strings.EqualFold(action.From, address) {
		action.Amount = 0 - action.Amount
	}
	action.Unit = action.Contract.Symbol
}

func (*Entity) handleErc721Transfer(address string, ev covalent.LogEvent, action *response.WalletTransactionAction) {
	action.From = ev.Decoded.Params[0].Value.(string)
	action.To = ev.Decoded.Params[1].Value.(string)
	tokenID := ev.Decoded.Params[2].Value.(string)
	action.Amount = 1
	if strings.EqualFold(action.From, address) {
		action.Amount = -1
	}
	if action.Contract.Symbol == "" {
		action.Unit = fmt.Sprintf("ERC-721 [%s]", action.Unit)
	} else {
		action.Unit = fmt.Sprintf("%s [%s]", action.Contract.Symbol, tokenID)
	}
}

func (*Entity) handleErc1155TransferSingle(address string, ev covalent.LogEvent, action *response.WalletTransactionAction) {
	action.From = ev.Decoded.Params[1].Value.(string)
	action.To = ev.Decoded.Params[2].Value.(string)
	tokenID := ev.Decoded.Params[3].Value.(string)
	_amount, _ := new(big.Float).SetString(ev.Decoded.Params[4].Value.(string))
	action.Amount, _ = _amount.Float64()
	if strings.EqualFold(action.From, address) {
		action.Amount = 0 - action.Amount
	}
	if action.Contract.Symbol == "" {
		action.Unit = fmt.Sprintf("ERC-1155 [%s]", action.Unit)
	} else {
		action.Unit = fmt.Sprintf("%s [%s]", action.Contract.Symbol, tokenID)
	}
}

func (*Entity) handleErc1155TransferBatch(address string, ev covalent.LogEvent, action *response.WalletTransactionAction) {
	action.From = ev.Decoded.Params[1].Value.(string)
	action.To = ev.Decoded.Params[2].Value.(string)
	if action.Contract.Symbol == "" {
		action.Unit = fmt.Sprintf("ERC-1155 [%s]", action.Unit)
	} else {
		action.Unit = fmt.Sprintf("%s [%s]", action.Contract.Symbol, action.Unit)
	}
}
