package entities

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
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
			assets = append(assets, response.WalletAssetData{
				ChainID:        chainID,
				ContractName:   item.ContractName,
				ContractSymbol: item.ContractTickerSymbol,
				AssetBalance:   parsedBal / math.Pow10(item.ContractDecimals),
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
		chain, _ := e.repo.Chain.GetByID(chainID)
		nativeSymbol := chain.Currency
		txBaseUrl := chain.TxBaseURL
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
				ChainID:      chainID,
				From:         item.FromAddress,
				To:           item.ToAddress,
				TxHash:       item.TxHash,
				SignedAt:     item.BlockSignedAt,
				NativeSymbol: nativeSymbol,
				TxBaseUrl:    txBaseUrl,
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
	// default tx type
	res.Type = "contract_interaction"
	// 1. transfer native token
	if tx.LogEvents == nil || len(tx.LogEvents) == 0 {
		amount, ok := new(big.Float).SetString(tx.Value)
		if !ok {
			err := fmt.Errorf("invalid tx amount %s", tx.Value)
			e.log.Errorf(err, "[getTxDetails] invalid native amount %s", tx.Value)
			return err
		}
		if amount.Cmp(big.NewFloat(0)) == 0 {
			err := errors.New("zero tx amount")
			e.log.Errorf(err, "[getTxDetails] zero transaction amount %s", tx.Value)
			return nil
		}
		res.Type = "transfer_native"
		res.Amount, _ = amount.Float64()
		res.Amount /= math.Pow10(18)
		if strings.EqualFold(tx.FromAddress, addr) {
			res.Amount = 0 - res.Amount
		}
		return nil
	}

	// 2. other tx types (transfer nft, transfer erc20)
	// res.NftIDs = make([]string, 0)
	events := tx.LogEvents
	firstEv := events[0]
	res.Contract = &response.ContractMetadata{
		Address: firstEv.SenderAddress,
		Symbol:  firstEv.SenderContractTickerSymbol,
		Name:    firstEv.SenderName,
	}
	if strings.EqualFold(firstEv.Decoded.Name, "Transfer") {
		res.From = firstEv.Decoded.Params[0].Value.(string)
		res.To = firstEv.Decoded.Params[1].Value.(string)
		if strings.EqualFold(firstEv.Decoded.Signature, "Transfer(indexed address from, indexed address to, uint256 value)") {
			res.Type = "transfer_erc20"
			amount, _ := new(big.Float).SetString(firstEv.Decoded.Params[2].Value.(string))
			res.Amount, _ = amount.Float64()
			res.Amount /= math.Pow10(firstEv.SenderContractDecimals)
		} else if strings.EqualFold(firstEv.Decoded.Signature, "Transfer(indexed address from, indexed address to, indexed uint256 tokenId)") {
			res.Type = "transfer_nft"
			res.Amount = 1
			res.NftIDs = append(res.NftIDs, firstEv.Decoded.Params[2].Value.(string))
		}
		// res.Amount = ev.Decoded.Params[2].Value.
		//  if ev.Decoded.Signature == "Transfer(indexed address from, indexed address to, uint256 value)" {
		// 	res.Type = "transfer_erc20"
		// 	res.From = ev.Decoded.Params[0].Value.(string)
		// 	res.To = ev.Decoded.Params[1].Value.(string)
		// 	res.Amount = ev.Decoded.Params[2].Value.

		// } else if ev.Decoded.Signature == "Transfer(indexed address from, indexed address to, indexed uint256 tokenId)" {
		// 	res.Type = "transfer_nft"
		// 	res.From = ev.Decoded.Params[0].Value.(string)
		// 	res.To = ev.Decoded.Params[1].Value.(string)
		// }
		// for _, p := range firstEv.Decoded.Params {
		// 	if strings.EqualFold(p.Name, "from") {
		// 		res.From = p.Value.(string)
		// 	} else if strings.EqualFold(p.Name, "to") {
		// 		res.To = p.Value.(string)
		// 	} else if strings.EqualFold(p.Name, "value") {
		// 		// res.Type = "transfer_erc20"
		// 		amount, _ := new(big.Float).SetString(p.Value.(string))
		// 		res.Amount, _ = amount.Float64()
		// 		res.Amount /= math.Pow10(firstEv.SenderContractDecimals)
		// 	} else if strings.EqualFold(p.Name, "tokenId") {
		// 		// res.Type = "transfer_nft"
		// 		res.Amount += 1
		// 		res.NftIDs = append(res.NftIDs, p.Value.(string))
		// 	}
		// }
		if strings.EqualFold(res.From, addr) {
			res.Amount = 0 - res.Amount
		}
		// if len(res.NftIDs) > 0 {
		// 	res.Type = "transfer_nft"
		// } else {
		// 	res.Type = "transfer_erc20"
		// }
	} else if strings.EqualFold(firstEv.Decoded.Name, "Approval") || strings.EqualFold(firstEv.Decoded.Name, "ApprovalForAll") {
		res.Type = "approval"
	}
	// for _, ev := range events {
	// if strings.EqualFold(ev.Decoded.Name, "Transfer") {
	// 	//  if ev.Decoded.Signature == "Transfer(indexed address from, indexed address to, uint256 value)" {
	// 	// 	res.Type = "transfer_erc20"
	// 	// 	res.From = ev.Decoded.Params[0].Value.(string)
	// 	// 	res.To = ev.Decoded.Params[1].Value.(string)
	// 	// 	res.Amount = ev.Decoded.Params[2].Value.

	// 	// } else if ev.Decoded.Signature == "Transfer(indexed address from, indexed address to, indexed uint256 tokenId)" {
	// 	// 	res.Type = "transfer_nft"
	// 	// 	res.From = ev.Decoded.Params[0].Value.(string)
	// 	// 	res.To = ev.Decoded.Params[1].Value.(string)
	// 	// }
	// 	for _, p := range ev.Decoded.Params {
	// 		if strings.EqualFold(p.Name, "from") {
	// 			res.From = p.Value.(string)
	// 		} else if strings.EqualFold(p.Name, "to") {
	// 			res.To = p.Value.(string)
	// 		} else if strings.EqualFold(p.Name, "value") {
	// 			// res.Type = "transfer_erc20"
	// 			amount, _ := new(big.Float).SetString(p.Value.(string))
	// 			res.Amount, _ = amount.Float64()
	// 			res.Amount /= math.Pow10(ev.SenderContractDecimals)
	// 		} else if strings.EqualFold(p.Name, "tokenId") {
	// 			// res.Type = "transfer_nft"
	// 			res.Amount += 1
	// 			res.NftIDs = append(res.NftIDs, p.Value.(string))
	// 		}
	// 	}
	// 	if strings.EqualFold(res.From, addr) {
	// 		res.Amount = 0 - res.Amount
	// 	}
	// 	if len(res.NftIDs) > 0 {
	// 		res.Type = "transfer_nft"
	// 	} else {
	// 		res.Type = "transfer_erc20"
	// 	}
	// } else if strings.EqualFold(ev.Decoded.Name, "Approval") || strings.EqualFold(ev.Decoded.Name, "ApprovalForAll") {
	// 	res.Type = "approval"
	// } else {
	// 	res.Type = "contract_interaction"
	// }
	// }

	// handle case transfer multiple nft
	if res.Type == "transfer_nft" && len(events) > 1 {
		for _, ev := range events[1:] {
			if strings.EqualFold(ev.Decoded.Signature, "Transfer(indexed address from, indexed address to, indexed uint256 tokenId)") {
				res.NftIDs = append(res.NftIDs, ev.Decoded.Params[2].Value.(string))
			}
		}
		res.Amount = float64(len(res.NftIDs))
	}
	return nil
}
