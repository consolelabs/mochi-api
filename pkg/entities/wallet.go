package entities

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerr "github.com/defipod/mochi/pkg/model/errors"
	userwalletwatchlistitem "github.com/defipod/mochi/pkg/repo/user_wallet_watchlist_item"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/covalent"
	"github.com/defipod/mochi/pkg/service/solscan"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetTrackingWallets(req request.GetTrackingWalletsRequest) ([]model.UserWalletWatchlistItem, error) {
	if req.IsOwner {
		// if error -> logging & ignore
		if err := e.upsertVerifiedWallet(req); err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetTrackingWallets] entity.upsertVerifiedWallet() failed")
		}
	}
	listQ := userwalletwatchlistitem.ListQuery{UserID: req.UserID, IsOwner: &req.IsOwner}
	wallets, err := e.repo.UserWalletWatchlistItem.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"userID": req.UserID}).Error(err, "[entity.GetTrackingWallets] repo.UserWalletWatchlistItem.List() failed")
		return nil, err
	}
	for i, wallet := range wallets {
		// 1. solana wallet
		if wallet.Type == "sol" {
			err = e.calculateSolWalletNetWorth(&wallet)
			if err != nil {
				e.log.Fields(logger.Fields{"wallet": wallet}).Error(err, "[entity.GetTrackingWallets] entity.calculateSolanaWalletNetWorth() failed")
				continue
			}
		} else {
			// 2. eth wallet
			err := e.calculateEthWalletNetWorth(&wallet)
			if err != nil {
				e.log.Fields(logger.Fields{"wallet": wallet}).Error(err, "[entity.GetTrackingWallets] entity.calculateEthWalletNetWorth() failed")
				continue
			}
		}
		wallet.FetchedData = true
		wallets[i] = wallet
	}
	return wallets, nil
}

func (e *Entity) upsertVerifiedWallet(req request.GetTrackingWalletsRequest) error {
	userWallet, err := e.repo.UserWallet.GetOneByDiscordIDAndGuildID(req.UserID, req.GuildID)
	// query failed -> exit
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetTrackingWallets] repo.UserWallet.GetOneByDiscordIDAndGuildID() failed")
		return err
	}

	// user has linked wallet from verify channel
	existing, err := e.repo.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{UserID: req.UserID, Query: userWallet.Address})
	// if not exists -> create
	if err == gorm.ErrRecordNotFound {
		err = e.repo.UserWalletWatchlistItem.Create(&model.UserWalletWatchlistItem{
			UserID:  req.UserID,
			Address: userWallet.Address,
			Type:    "eth",
			IsOwner: true,
		})
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetTrackingWallets] repo.UserWalletWatchlistItem.Create() failed")
			return err
		}
		return nil
	}
	// query failed -> exit
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetTrackingWallets] repo.UserWalletWatchlistItem.GetOne() failed")
		return err
	}
	// if exists but is_owner = false -> update is_owner to true
	if !existing.IsOwner {
		err = e.repo.UserWalletWatchlistItem.UpdateOwnerFlag(req.UserID, userWallet.Address, true)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.GetTrackingWallets] repo.UserWalletWatchlistItem.UpdateOwnerFlag() failed")
			return err
		}
	}
	return nil
}

func (e *Entity) calculateSolWalletNetWorth(wallet *model.UserWalletWatchlistItem) error {
	solBalance, err := e.solana.Balance(wallet.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"addr": wallet.Address}).Error(err, "[entity.calculateSolWalletNetWorth] solana.Balance() failed")
		return err
	}
	prices, err := e.svc.CoinGecko.GetCoinPrice([]string{"solana"}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"id": "solana"}).Error(err, "[entity.calculateSolWalletNetWorth] svc.CoinGecko.GetCoinPrice() failed")
		return err
	}
	wallet.NetWorth += solBalance * prices["solana"]
	tokenBalances, err := e.svc.Solscan.GetTokenBalances(wallet.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"address": wallet.Address}).Error(err, "[entity.calculateSolWalletNetWorth] svc.Solscan.GetTokenBalances() failed")
		return err
	}
	for _, tb := range tokenBalances {
		metadata, err := e.svc.Solscan.GetTokenMetadata(tb.TokenAddress)
		if err != nil {
			e.log.Fields(logger.Fields{"tokenAddress": tb.TokenAddress}).Error(err, "[entity.calculateSolWalletNetWorth] svc.Solscan.GetTokenMetadata() failed")
			continue
		}
		wallet.NetWorth += tb.TokenAmount.UIAmount * metadata.Price
	}
	return nil
}

func (e *Entity) calculateEthWalletNetWorth(wallet *model.UserWalletWatchlistItem) error {
	chainIDs := []int{1, 56, 137, 250, 2020}
	for _, chainID := range chainIDs {
		res, err := e.svc.Covalent.GetTokenBalances(chainID, wallet.Address, 3)
		if err != nil {
			e.log.Fields(logger.Fields{"chainID": chainID, "addr": wallet.Address}).Error(err, "[entity.calculateEthWalletNetWorth] svc.Covalent.GetTokenBalances() failed")
			return err
		}
		if res.Data.Items == nil || len(res.Data.Items) == 0 {
			continue
		}
		for _, item := range res.Data.Items {
			if item.Type != "cryptocurrency" {
				continue
			}
			_, quote := e.calculateTokenBalance(item, chainID)
			wallet.NetWorth += quote
		}
	}
	return nil
}

func (e *Entity) calculateTokenBalance(item covalent.TokenBalanceItem, chainID int) (bal, quote float64) {
	balance, ok := new(big.Float).SetString(item.Balance)
	if !ok {
		return
	}
	parsedBal, _ := balance.Float64()
	bal = parsedBal / math.Pow10(item.ContractDecimals)
	if strings.EqualFold(item.ContractTickerSymbol, "icy") && chainID == 137 {
		quote = 1.5 * bal
	} else {
		quote = item.Quote
	}
	return
}

func (e *Entity) GetOneWallet(req request.GetOneWalletRequest) (*model.UserWalletWatchlistItem, error) {
	req.Standardize()
	wallet, err := e.repo.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{UserID: req.UserID, Query: req.AliasOrAddress})
	if err != nil {
		e.log.Fields(logger.Fields{"userID": req.UserID}).Error(err, "[entity.GetOneWallet] repo.UserWalletWatchlistItem.GetOne() failed")
		return nil, err
	}
	return wallet, nil
}

func (e *Entity) TrackWallet(req request.TrackWalletRequest) error {
	if req.Alias != "" {
		wallet, err := e.repo.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{UserID: req.UserID, Query: req.Alias})
		if err != nil && err != gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TrackWallet] repo.UserWalletWatchlistItem.GetOne() failed")
			return err
		}
		if err == nil && !strings.EqualFold(wallet.Address, req.Address) {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TrackWallet] alias has been used for another address")
			return baseerr.ErrConflict
		}
		if err == nil {
			req.Alias = wallet.Alias
		}
	}

	err := e.repo.UserWalletWatchlistItem.Create(&model.UserWalletWatchlistItem{
		UserID:  req.UserID,
		Address: req.Address,
		Alias:   req.Alias,
		Type:    req.Type,
		IsOwner: req.IsOwner,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TrackWallet] repo.UserWalletWatchlistItem.Create() failed")
		return err
	}
	if req.ChannelID == "" || req.MessageID == "" {
		return nil
	}
	err = e.notifyWalletAddition(req)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.TrackWallet] entity.notifyWalletAddition() failed")
		return err
	}
	return nil
}

func (e *Entity) notifyWalletAddition(req request.TrackWalletRequest) error {
	description := fmt.Sprintf("Wallet `%s` was added succesfully!\n%s You can set a label for your wallet or add more wallets.", req.Address, util.GetEmoji("pointingdown"))
	embed := &discordgo.MessageEmbed{
		Color: 0x5cd97d,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Successfully set!",
			IconURL: "https://cdn.discordapp.com/emojis/933341948402618378.png?size=240&quality=lossless",
		},
		Description: description,
	}
	_, err := e.discord.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Channel: req.ChannelID,
		ID:      req.MessageID,
		Embeds:  []*discordgo.MessageEmbed{embed},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "View Wallet",
						Style:    discordgo.SecondaryButton,
						CustomID: fmt.Sprintf("wallet_view_details-%s", req.Address),
						Emoji: discordgo.ComponentEmoji{
							Name: "wallet",
							ID:   util.GetEmojiID("wallet"),
						},
					},
					discordgo.Button{
						Label:    "Rename Label",
						Style:    discordgo.SecondaryButton,
						CustomID: fmt.Sprintf("wallet_rename-%s-%s", req.UserID, req.Address),
						Emoji: discordgo.ComponentEmoji{
							Name: "pencil",
							ID:   util.GetEmojiID("pencil"),
						},
					},
					discordgo.Button{
						Label:    "Add More",
						Style:    discordgo.SecondaryButton,
						CustomID: fmt.Sprintf("wallet_add_more-%s", req.UserID),
						Emoji: discordgo.ComponentEmoji{
							Name: "plus",
							ID:   util.GetEmojiID("plus"),
						},
					},
				},
			},
		},
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.notifyWalletAddition] discord.ChannelMessageEditComplex() failed")
		return err
	}
	return nil
}

func (e *Entity) UntrackWallet(req request.UntrackWalletRequest) error {
	return e.repo.UserWalletWatchlistItem.Remove(userwalletwatchlistitem.DeleteQuery{
		UserID:  req.UserID,
		Address: req.Address,
		Alias:   req.Alias,
	})
}

func (e *Entity) ListWalletAssets(req request.ListWalletAssetsRequest) ([]response.WalletAssetData, string, string, error) {
	req.Standardize()
	if req.Type == "sol" {
		return e.listSolWalletAssets(req)
	}
	if req.Type == "sui" {
		return e.listSuiWalletAssets(req)
	}
	return e.listEthWalletAssets(req)
}

func (e *Entity) listEthWalletAssets(req request.ListWalletAssetsRequest) ([]response.WalletAssetData, string, string, error) {
	address, err := util.ConvertToChecksumAddr(req.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listEthWalletAssets] util.ConvertToChecksumAddr() failed")
		return nil, "", "", err
	}

	// redis cache
	value, err := e.cache.HashGet(address + "-eth")
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to get cache data wallet")
		return nil, "", "", err
	}

	chainIDs := []int{1, 56, 137, 250, 2020}
	assets := make([]response.WalletAssetData, 0)
	if len(value) == 0 {
		for _, chainID := range chainIDs {
			res, err := e.svc.Covalent.GetTokenBalances(chainID, address, 3)
			if err != nil {
				e.log.Fields(logger.Fields{"chainID": chainID, "address": address}).Error(err, "[entity.listEthWalletAssets] svc.Covalent.GetTokenBalances() failed")
				return nil, "", "", err
			}
			if res.Data.Items == nil || len(res.Data.Items) == 0 {
				continue
			}
			for _, item := range res.Data.Items {
				if item.Type != "cryptocurrency" {
					continue
				}
				bal, quote := e.calculateTokenBalance(item, chainID)
				// filter out dusty tokens
				if quote < 0.001 {
					continue
				}
				assets = append(assets, response.WalletAssetData{
					ChainID:        chainID,
					ContractName:   item.ContractName,
					ContractSymbol: item.ContractTickerSymbol,
					AssetBalance:   bal,
					UsdBalance:     quote,
					Token: response.AssetToken{
						Name:    item.ContractName,
						Symbol:  item.ContractTickerSymbol,
						Decimal: int64(item.ContractDecimals),
						Price:   item.QuoteRate,
						Native:  item.NativeToken,
						Chain: response.AssetTokenChain{
							Name: res.Data.ChainName,
						},
					},
					Amount: util.FloatToString(fmt.Sprint(bal), int64(item.ContractDecimals)),
				})
			}
		}

		encodeData := make(map[string]string)
		if len(assets) == 0 {
			encodeData["empty"] = "empty"
		}

		for _, asset := range assets {
			encodeData[fmt.Sprintf("%s-%s-%d-%d-%f-%v-%s", asset.ContractName, asset.ContractSymbol, asset.ChainID, asset.Token.Decimal, asset.Token.Price, asset.Token.Native, asset.Token.Chain.Name)] = fmt.Sprintf("%f-%f", asset.AssetBalance, asset.UsdBalance)
		}

		err := e.cache.HashSet(address+"-eth", encodeData, 6*time.Hour)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to set cache data wallet")
			return nil, "", "", err
		}

	} else {
		for k, v := range value {
			if k == "empty" {
				break
			}

			key := strings.Split(k, "-")
			value := strings.Split(v, "-")
			chainId, _ := strconv.Atoi(key[2])
			decimal, _ := strconv.Atoi(key[3])
			price, _ := strconv.ParseFloat(key[4], 64)
			native, _ := strconv.ParseBool(key[5])
			assetBalance, _ := strconv.ParseFloat(value[0], 64)
			usdBalance, _ := strconv.ParseFloat(value[1], 64)
			// filter out dusty tokens
			if usdBalance < 0.001 {
				continue
			}

			assets = append(assets, response.WalletAssetData{
				ContractName:   key[0],
				ContractSymbol: key[1],
				ChainID:        chainId,
				AssetBalance:   assetBalance,
				UsdBalance:     usdBalance,
				Token: response.AssetToken{
					Name:    key[0],
					Symbol:  key[1],
					Decimal: int64(decimal),
					Price:   price,
					Native:  native,
					Chain: response.AssetTokenChain{
						Name: key[6],
					},
				},
				Amount: util.FloatToString(fmt.Sprint(assetBalance), int64(decimal)),
			})
		}
	}

	// calculate pnl
	pnl, latestSnapshotBal, err := e.calculateWalletSnapshot(address, true, assets)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listEthWalletAssets] calculateWalletSnapshot() failed")
		return assets, "", "", nil
	}

	return assets, pnl, latestSnapshotBal, nil
}

func (e *Entity) calculateWalletSnapshot(address string, isEvm bool, assets []response.WalletAssetData) (string, string, error) {
	totalAmount := sumBal(assets)
	// store snapshot whenever used wallet service
	_, err := e.repo.WalletSnapshot.Create(&model.WalletSnapshot{
		WalletAddress:   address,
		IsEvm:           isEvm,
		TotalUsdBalance: fmt.Sprintf("%.4f", totalAmount),
		SnapshotTime:    time.Now(),
	})
	if err != nil {
		e.log.Fields(logger.Fields{"address": address}).Error(err, "[entity.calculateWalletSnapshot] repo.WalletSnapshot.Create() failed")
		return "", "", err
	}

	// get snapshot in 8 hour
	snapshots, err := e.repo.WalletSnapshot.GetSnapshotInTime(address, time.Now().Add(-8*time.Hour))
	if err != nil {
		e.log.Fields(logger.Fields{"address": address}).Error(err, "[entity.calculateWalletSnapshot] repo.WalletSnapshot.GetSnapshotInTime() failed")
		return "", "", err
	}

	// this means in last 8 hour no data, get latest data we have in db
	latestSnapshotBal := 0.0
	if len(snapshots) == 1 {
		latestSnapshot, err := e.repo.WalletSnapshot.GetLatestInPast(address, time.Now().Add(-8*time.Hour))
		if err != nil {
			e.log.Fields(logger.Fields{"address": address}).Error(err, "[entity.calculateWalletSnapshot] repo.WalletSnapshot.GetLatestInPast() failed")
			return "", "", err
		}

		// this is the first time user add data to snapshot
		if len(latestSnapshot) == 0 {
			return fmt.Sprintf("%.4f", totalAmount-latestSnapshotBal), fmt.Sprintf("%.4f", latestSnapshotBal), nil
		}

		latestSnapshotBal, _ = strconv.ParseFloat(latestSnapshot[0].TotalUsdBalance, 64)
	} else {
		latestSnapshotBal, _ = strconv.ParseFloat(snapshots[0].TotalUsdBalance, 64)
	}

	return fmt.Sprintf("%.4f", totalAmount-latestSnapshotBal), fmt.Sprintf("%.4f", latestSnapshotBal), nil
}

func (e *Entity) listSolWalletAssets(req request.ListWalletAssetsRequest) ([]response.WalletAssetData, string, string, error) {
	// redis cache
	value, err := e.cache.HashGet(req.Address + "-sol")
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to get cache data wallet")
		return nil, "", "", err
	}

	chainIDs := []int{999}
	assets := make([]response.WalletAssetData, 0)
	if len(value) == 0 {
		for _, chainID := range chainIDs {
			res, err := e.svc.Covalent.GetSolanaTokenBalances("solana-mainnet", req.Address, 3)
			if err != nil {
				e.log.Fields(logger.Fields{"chainID": chainID, "address": req.Address}).Error(err, "[entity.listSolWalletAssets] svc.Covalent.GetTokenBalances() failed")
				return nil, "", "", err
			}
			if res.Data.Items == nil || len(res.Data.Items) == 0 {
				continue
			}
			for _, item := range res.Data.Items {
				if item.Type != "cryptocurrency" {
					continue
				}
				bal, quote := e.calculateTokenBalance(item, chainID)
				assets = append(assets, response.WalletAssetData{
					ChainID:        chainID,
					ContractName:   item.ContractName,
					ContractSymbol: item.ContractTickerSymbol,
					AssetBalance:   bal,
					UsdBalance:     quote,
					Token: response.AssetToken{
						Name:    item.ContractName,
						Symbol:  item.ContractTickerSymbol,
						Decimal: int64(item.ContractDecimals),
						Price:   item.QuoteRate,
						Native:  item.NativeToken,
						Chain: response.AssetTokenChain{
							Name: res.Data.ChainName,
						},
					},
					Amount: util.FloatToString(fmt.Sprint(bal), int64(item.ContractDecimals)),
				})
			}
		}

		encodeData := make(map[string]string)
		if len(assets) == 0 {
			encodeData["empty"] = "empty"
		}

		for _, asset := range assets {
			encodeData[fmt.Sprintf("%s-%s-%d-%d-%f-%v-%s", asset.ContractName, asset.ContractSymbol, asset.ChainID, asset.Token.Decimal, asset.Token.Price, asset.Token.Native, asset.Token.Chain.Name)] = fmt.Sprintf("%f-%f", asset.AssetBalance, asset.UsdBalance)
		}

		err := e.cache.HashSet(req.Address+"-sol", encodeData, 6*time.Hour)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to set cache data wallet")
			return nil, "", "", err
		}

	} else {
		for k, v := range value {
			if k == "empty" {
				break
			}

			key := strings.Split(k, "-")
			value := strings.Split(v, "-")
			chainId, _ := strconv.Atoi(key[2])
			decimal, _ := strconv.Atoi(key[3])
			price, _ := strconv.ParseFloat(key[4], 64)
			native, _ := strconv.ParseBool(key[5])
			assetBalance, _ := strconv.ParseFloat(value[0], 64)
			usdBalance, _ := strconv.ParseFloat(value[1], 64)

			assets = append(assets, response.WalletAssetData{
				ContractName:   key[0],
				ContractSymbol: key[1],
				ChainID:        chainId,
				AssetBalance:   assetBalance,
				UsdBalance:     usdBalance,
				Token: response.AssetToken{
					Name:    key[0],
					Symbol:  key[1],
					Decimal: int64(decimal),
					Price:   price,
					Native:  native,
					Chain: response.AssetTokenChain{
						Name: key[6],
					},
				},
				Amount: util.FloatToString(fmt.Sprint(assetBalance), int64(decimal)),
			})
		}
	}

	// calculate pnl
	pnl, latestSnapshotBal, err := e.calculateWalletSnapshot(req.Address, false, assets)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listEthWalletAssets] calculateWalletSnapshot() failed")
		return assets, "", "", nil
	}

	return assets, pnl, latestSnapshotBal, nil
}

func (e *Entity) listSuiWalletAssets(req request.ListWalletAssetsRequest) ([]response.WalletAssetData, string, string, error) {
	// redis cache
	value, err := e.cache.HashGet(req.Address + "-sui")
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to get cache data wallet")
		return nil, "", "", err
	}

	assets := make([]response.WalletAssetData, 0)
	if len(value) == 0 {
		assets, err = e.svc.Sui.GetAddressAssets(req.Address)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to set get sui assets wallet")
			return nil, "", "", err
		}

		encodeData := make(map[string]string)
		if len(assets) == 0 {
			encodeData["empty"] = "empty"
		}

		for _, asset := range assets {
			encodeData[fmt.Sprintf("%s-%s-%d-%d-%f-%v-%s", asset.ContractName, asset.ContractSymbol, asset.ChainID, asset.Token.Decimal, asset.Token.Price, asset.Token.Native, asset.Token.Chain.Name)] = fmt.Sprintf("%f-%f", asset.AssetBalance, asset.UsdBalance)
		}

		err := e.cache.HashSet(req.Address+"-sol", encodeData, 6*time.Hour)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to set cache data wallet")
			return nil, "", "", err
		}

	} else {
		for k, v := range value {
			if k == "empty" {
				break
			}

			key := strings.Split(k, "-")
			value := strings.Split(v, "-")
			chainId, _ := strconv.Atoi(key[2])
			decimal, _ := strconv.Atoi(key[3])
			price, _ := strconv.ParseFloat(key[4], 64)
			native, _ := strconv.ParseBool(key[5])
			assetBalance, _ := strconv.ParseFloat(value[0], 64)
			usdBalance, _ := strconv.ParseFloat(value[1], 64)

			assets = append(assets, response.WalletAssetData{
				ContractName:   key[0],
				ContractSymbol: key[1],
				ChainID:        chainId,
				AssetBalance:   assetBalance,
				UsdBalance:     usdBalance,
				Token: response.AssetToken{
					Name:    key[0],
					Symbol:  key[1],
					Decimal: int64(decimal),
					Price:   price,
					Native:  native,
					Chain: response.AssetTokenChain{
						Name: key[6],
					},
				},
				Amount: util.FloatToString(fmt.Sprint(assetBalance), int64(decimal)),
			})
		}
	}

	// calculate pnl
	pnl, latestSnapshotBal, err := e.calculateWalletSnapshot(req.Address, false, assets)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listEthWalletAssets] calculateWalletSnapshot() failed")
		return assets, "", "", nil
	}

	return assets, pnl, latestSnapshotBal, nil
}

func (e *Entity) ListWalletTxns(req request.ListWalletTransactionsRequest) ([]response.WalletTransactionData, error) {
	req.Standardize()
	if req.Type == "sol" {
		return e.listSolWalletTxns(req)
	}
	return e.listEthWalletTxns(req)
}

func (e *Entity) listEthWalletTxns(req request.ListWalletTransactionsRequest) ([]response.WalletTransactionData, error) {
	chainIDs := []int{1, 56, 137, 250, 2020}
	txns := make([]response.WalletTransactionData, 0)
	for _, chainID := range chainIDs {
		res, err := e.svc.Covalent.GetTransactionsByAddress(chainID, req.Address, 5, 5)
		if err != nil {
			e.log.Fields(logger.Fields{"chainID": chainID, "address": req.Address}).Error(err, "[entity.listEthWalletTxns] svc.Covalent.GetTransactionsByAddress() failed")
			return nil, err
		}
		if res.Data.Items == nil || len(res.Data.Items) == 0 {
			continue
		}
		for _, item := range res.Data.Items {
			tx := response.WalletTransactionData{
				ChainID:    chainID,
				TxHash:     item.TxHash,
				SignedAt:   item.BlockSignedAt,
				Successful: item.Successful,
			}
			//
			if err := e.parseCovalentTxData(item, &tx, req.Address); err != nil {
				e.log.Fields(logger.Fields{"item": item, "tx": tx, "address": req.Address}).Error(err, "[entity.listEthWalletTxns] entity.parseCovalentTxData() failed")
			}
			txns = append(txns, tx)
		}
	}
	sort.Slice(txns, func(i, j int) bool {
		return txns[i].SignedAt.Unix() > txns[j].SignedAt.Unix()
	})
	return txns, nil
}

func (e *Entity) listSolWalletTxns(req request.ListWalletTransactionsRequest) ([]response.WalletTransactionData, error) {
	res := make([]response.WalletTransactionData, 0)
	txns, err := e.svc.Solscan.GetTransactions(req.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"address": req.Address}).Error(err, "[entity.listSolWalletTxns] svc.Solscan.GetTransactions() failed")
		return nil, err
	}
	for _, item := range txns {
		data := response.WalletTransactionData{
			ChainID:     999,
			TxHash:      item.TxHash,
			ScanBaseUrl: "https://solscan.io",
			SignedAt:    time.Unix(int64(item.BlockTime), 0),
			Successful:  false,
		}
		if strings.EqualFold(item.Status, "success") {
			data.Successful = true
		}
		tx, err := e.svc.Solscan.GetTxDetails(item.TxHash)
		if err != nil {
			e.log.Fields(logger.Fields{"txHash": item.TxHash}).Error(err, "[entity.listSolWalletTxns] svc.Solscan.GetTxDetails() failed")
			return nil, err
		}
		e.handleSolTransfers(req.Address, tx, &data)
		e.handleSolTokenTransfers(req.Address, tx, &data)
		res = append(res, data)
	}
	return res, nil
}

func (e *Entity) handleSolTransfers(address string, tx *solscan.TransactionDetailsResponse, res *response.WalletTransactionData) {
	for _, transfer := range tx.SolTransfers {
		value := new(big.Float).SetInt64(int64(transfer.Amount))
		amount, _ := value.Float64()
		amount /= math.Pow10(9)
		if strings.EqualFold(transfer.Source, address) {
			amount = 0 - amount
		}
		res.Actions = append(res.Actions, response.WalletTransactionAction{
			Amount:         amount,
			NativeTransfer: true,
			Unit:           "SOL",
			From:           transfer.Source,
			To:             transfer.Destination,
		})
		res.HasTransfer = true
	}
}

func (e *Entity) handleSolTokenTransfers(address string, tx *solscan.TransactionDetailsResponse, data *response.WalletTransactionData) {
	for _, transfer := range tx.TokenTransfers {
		value, ok := new(big.Float).SetString(transfer.Amount)
		if !ok {
			return
		}
		amount, _ := value.Float64()
		amount /= math.Pow10(transfer.Token.Decimals)
		if strings.EqualFold(transfer.SourceOwner, address) {
			amount = 0 - amount
		}
		action := response.WalletTransactionAction{
			Amount: amount,
			Unit:   transfer.Token.Symbol,
			From:   transfer.SourceOwner,
			To:     transfer.DestinationOwner,
			Contract: &response.ContractMetadata{
				Address: transfer.Token.Address,
			},
		}
		if action.Unit == "" {
			contract, err := e.svc.Solscan.GetTokenMetadata(transfer.Token.Address)
			if err != nil {
				e.log.Fields(logger.Fields{"txHash": tx.TxHash}).Error(err, "[entity.listSolWalletTxns] svc.Solscan.GetTokenMetadata() failed")
			} else {
				action.Unit = contract.Symbol
			}
		}
		data.Actions = append(data.Actions, action)
		data.HasTransfer = true
	}
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
			From:           tx.FromAddress,
			To:             tx.ToAddress,
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

func (e *Entity) GenerateWalletVerification(req request.GenerateWalletVerificationRequest) (string, error) {
	code := uuid.New().String()
	err := e.repo.DiscordWalletVerification.UpsertOne(
		model.DiscordWalletVerification{
			Code:          code,
			UserDiscordID: req.UserID,
			GuildID:       "",
			CreatedAt:     time.Now(),
			ChannelID:     req.ChannelID,
			MessageID:     req.MessageID,
		},
	)
	if err != nil {
		e.log.Error(err, "[entity.GenerateWalletVerification] repo.DiscordWalletVerification.UpsertOne failed")
	}
	return code, err
}
