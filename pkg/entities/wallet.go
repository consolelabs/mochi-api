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

	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerr "github.com/defipod/mochi/pkg/model/errors"
	userwalletwatchlistitem "github.com/defipod/mochi/pkg/repo/user_wallet_watchlist_item"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/covalent"
	"github.com/defipod/mochi/pkg/service/krystal"
	"github.com/defipod/mochi/pkg/service/solscan"
	"github.com/defipod/mochi/pkg/util"
)

func (e *Entity) GetTrackingWallets(req request.GetTrackingWalletsRequest) (*model.UserWalletWatchlist, error) {
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
		if wallet.ChainType == model.ChainTypeSolana {
			err = e.calculateSolWalletNetWorth(&wallet)
			if err != nil {
				e.log.Fields(logger.Fields{"wallet": wallet}).Error(err, "[entity.GetTrackingWallets] entity.calculateSolanaWalletNetWorth() failed")
				continue
			}
		} else {
			// 2. EVM wallet
			err := e.calculateEvmWalletNetWorth(&wallet)
			if err != nil {
				e.log.Fields(logger.Fields{"wallet": wallet}).Error(err, "[entity.GetTrackingWallets] entity.calculateEvmWalletNetWorth() failed")
				continue
			}
		}

		wallet.FetchedData = true
		wallets[i] = wallet
	}

	var result model.UserWalletWatchlist
	for _, wallet := range wallets {
		switch wallet.Type {
		case model.TrackingTypeFollow:
			result.Following = append(result.Following, wallet)
		case model.TrackingTypeTrack:
			result.Tracking = append(result.Tracking, wallet)
		case model.TrackingTypeCopy:
			result.Copying = append(result.Copying, wallet)
		}
	}

	return &result, nil
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
		err = e.repo.UserWalletWatchlistItem.Upsert(&model.UserWalletWatchlistItem{
			UserID:  req.UserID,
			Address: userWallet.Address,
			Type:    "evm",
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

func (e *Entity) calculateEvmWalletNetWorth(wallet *model.UserWalletWatchlistItem) error {
	res, err := e.svc.Krystal.GetBalanceTokenByAddress(wallet.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"address": wallet.Address}).Error(err, "[entity.calculateEvmWalletNetWorth] svc.Krystal.GetBalanceTokenByAddress() failed")
		return err
	}

	if res == nil || res.Data == nil || len(res.Data) == 0 {
		return nil
	}

	for _, item := range res.Data {
		for _, bal := range item.Balances {
			_, quote := e.calculateEvmTokenBalance(bal, item.ChainId)
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

func (e *Entity) TrackWallet(mod model.UserWalletWatchlistItem, channelID, messageID string) error {
	if mod.Alias != "" {
		wallet, err := e.repo.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{UserID: mod.UserID, Query: mod.Alias})
		if err != nil && err != gorm.ErrRecordNotFound {
			e.log.Fields(logger.Fields{"mod": mod}).Error(err, "[entity.TrackWallet] repo.UserWalletWatchlistItem.GetOne() failed")
			return err
		}
		if err == nil && !strings.EqualFold(wallet.Address, mod.Address) {
			e.log.Fields(logger.Fields{"mod": mod}).Error(err, "[entity.TrackWallet] alias has been used for another address")
			return baseerr.ErrConflict
		}
		if err == nil {
			mod.Alias = wallet.Alias
		}
	}

	existItem, err := e.repo.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{UserID: mod.UserID, Query: mod.Address})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		e.log.Fields(logger.Fields{"mod": mod}).Error(err, "[entity.TrackWallet] repo.UserWalletWatchlistItem.GetOne() failed")
		return err
	}
	if existItem.ChainType != "" && !strings.EqualFold(existItem.ChainType.String(), mod.ChainType.String()) {
		e.log.Fields(logger.Fields{"mod": mod}).Error(err, "[entity.TrackWallet] wallet chain type does not match")
		return baseerr.ErrChainTypeConflict
	}

	err = e.repo.UserWalletWatchlistItem.Upsert(&mod)
	if err != nil {
		e.log.Fields(logger.Fields{"mod": mod}).Error(err, "[entity.TrackWallet] repo.UserWalletWatchlistItem.Create() failed")
		return err
	}

	if channelID == "" || messageID == "" {
		return nil
	}

	err = e.notifyWalletAddition(mod, channelID, messageID)
	if err != nil {
		e.log.Fields(logger.Fields{"req": mod}).Error(err, "[entity.TrackWallet] entity.notifyWalletAddition() failed")
		return err
	}
	return nil
}

func (e *Entity) notifyWalletAddition(mod model.UserWalletWatchlistItem, channelID, messageID string) error {
	description := fmt.Sprintf("Wallet `%s` was added succesfully!\n%s You can set a label for your wallet or add more wallets.", mod.Address, util.GetEmoji("pointingdown"))
	embed := &discordgo.MessageEmbed{
		Color: 0x5cd97d,
		Author: &discordgo.MessageEmbedAuthor{
			Name:    "Successfully set!",
			IconURL: "https://cdn.discordapp.com/emojis/933341948402618378.png?size=240&quality=lossless",
		},
		Description: description,
	}
	_, err := e.discord.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Channel: channelID,
		ID:      messageID,
		Embeds:  []*discordgo.MessageEmbed{embed},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "View Wallet",
						Style:    discordgo.SecondaryButton,
						CustomID: fmt.Sprintf("wallet_view_details-%s", mod.Address),
						Emoji: discordgo.ComponentEmoji{
							Name: "wallet",
							ID:   util.GetEmojiID("wallet"),
						},
					},
					discordgo.Button{
						Label:    "Rename Label",
						Style:    discordgo.SecondaryButton,
						CustomID: fmt.Sprintf("wallet_rename-%s-%s", mod.UserID, mod.Address),
						Emoji: discordgo.ComponentEmoji{
							Name: "pencil",
							ID:   util.GetEmojiID("pencil"),
						},
					},
					discordgo.Button{
						Label:    "Add More",
						Style:    discordgo.SecondaryButton,
						CustomID: fmt.Sprintf("wallet_add_more-%s", mod.UserID),
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
		e.log.Fields(logger.Fields{"req": mod}).Error(err, "[entity.notifyWalletAddition] discord.ChannelMessageEditComplex() failed")
		return nil
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

	// Solana
	if req.Type == "sol" {
		return e.listSolWalletAssets(req)
	}

	// SUI
	if req.Type == "sui" {
		return e.listSuiWalletAssets(req)
	}

	// Ronin
	if req.Type == "ron" {
		return e.listRoninWalletAssets(req)
	}

	// EVM-based chains
	return e.listEvmWalletAssets(req)
}

func (e *Entity) listRoninWalletAssets(req request.ListWalletAssetsRequest) (data []response.WalletAssetData, pnl string, latestSnapshotBal string, err error) {
	chainID := 2020
	res, err := e.svc.Covalent.GetTokenBalances(chainID, req.Address, 3)
	if err != nil {
		e.log.Fields(logger.Fields{"addr": req.Address}).Error(err, "[entity.listRoninWalletAssets] svc.Covalent.GetTokenBalances() failed")
		return
	}

	if res == nil || res.Data == nil || res.Data.Items == nil || len(res.Data.Items) == 0 {
		return
	}

	for _, item := range res.Data.Items {
		if item.Type != "cryptocurrency" {
			continue
		}
		bal, quote := e.calculateTokenBalance(item, 2020)
		// filter out dusty tokens
		if quote < 0.001 {
			continue
		}

		data = append(data, response.WalletAssetData{
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
					Name:      res.Data.ChainName,
					ShortName: "RON",
				},
			},
			Amount: util.FloatToString(fmt.Sprint(bal), int64(item.ContractDecimals)),
		})
	}

	// calculate pnl
	pnl, latestSnapshotBal, err = e.calculateWalletSnapshot(req.Address, true, data)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listRoninWalletAssets] calculateWalletSnapshot() failed")
		return nil, "", "", err
	}

	return
}

func (e *Entity) listEvmWalletAssets(req request.ListWalletAssetsRequest) ([]response.WalletAssetData, string, string, error) {
	address, err := util.ConvertToChecksumAddr(req.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listEvmWalletAssets] util.ConvertToChecksumAddr() failed")
		return nil, "", "", err
	}

	// get all tokens balances by address & chainIds
	res, err := e.svc.Krystal.GetBalanceTokenByAddress(address)
	if err != nil {
		e.log.Fields(logger.Fields{"address": address}).Error(err, "[entity.listEvmWalletAssets] svc.Krystal.GetBalanceTokenByAddress() failed")
		return nil, "", "", err
	}

	assets := make([]response.WalletAssetData, 0)

	for _, item := range res.Data {
		chain, err := e.repo.Chain.GetByID(item.ChainId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			e.log.Fields(logger.Fields{"chainID": item.ChainId}).Error(err, "[entity.listEvmWalletAssets] repo.Chain.GetByID() failed")
			return nil, "", "", err
		}

		for _, bal := range item.Balances {
			assetBal, quote := e.calculateEvmTokenBalance(bal, item.ChainId)
			// filter out dusty tokens
			if quote < 0.001 {
				continue
			}
			assets = append(assets, response.WalletAssetData{
				ChainID:        item.ChainId,
				ContractName:   bal.Token.Name,
				ContractSymbol: bal.Token.Symbol,
				AssetBalance:   assetBal,
				UsdBalance:     quote,
				Token: response.AssetToken{
					Name:    bal.Token.Symbol,
					Symbol:  bal.Token.Name,
					Decimal: int64(bal.Token.Decimals),
					Price:   bal.Quotes.Usd.Price,
					Native:  bal.TokenType == "NATIVE",
					Chain: response.AssetTokenChain{
						Name:      item.ChainName,
						ShortName: chain.ShortName,
					},
				},
				Amount: util.FloatToString(fmt.Sprint(assetBal), int64(bal.Token.Decimals)),
			})
		}

	}

	// calculate pnl
	pnl, latestSnapshotBal, err := e.calculateWalletSnapshot(address, true, assets)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listEvmWalletAssets] calculateWalletSnapshot() failed")
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
	chainIDs := []int{999}
	assets := make([]response.WalletAssetData, 0)
	for _, chainID := range chainIDs {
		// get chain
		chain, err := e.repo.Chain.GetByID(chainID)
		if err != nil {
			e.log.Fields(logger.Fields{"chainID": chainID}).Error(err, "[entity.listSolWalletAssets] repo.Chain.GetByID() failed")
			return nil, "", "", err
		}

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

			tokenAddress := item.ContractAddress
			if item.NativeToken {
				tokenAddress = consts.SolAddress
			}

			bal, _ := e.calculateTokenBalance(item, chainID)

			tokenPrice, err := e.svc.Birdeye.GetTokenPrice(tokenAddress)
			if err != nil {
				e.log.Fields(logger.Fields{"chainID": chainID, "address": req.Address}).Error(err, "[entity.listSolWalletAssets] svc.Birdeye.GetTokenPrice() failed")
				continue
			}

			assets = append(assets, response.WalletAssetData{
				ChainID:        chainID,
				ContractName:   item.ContractName,
				ContractSymbol: item.ContractTickerSymbol,
				AssetBalance:   bal,
				UsdBalance:     tokenPrice.Data.Value * bal,
				Token: response.AssetToken{
					Name:    item.ContractName,
					Symbol:  item.ContractTickerSymbol,
					Decimal: int64(item.ContractDecimals),
					Price:   tokenPrice.Data.Value,
					Native:  item.NativeToken,
					Chain: response.AssetTokenChain{
						Name:      res.Data.ChainName,
						ShortName: chain.ShortName,
					},
				},
				Amount: util.FloatToString(fmt.Sprint(bal), int64(item.ContractDecimals)),
			})
		}
	}

	// calculate pnl
	pnl, latestSnapshotBal, err := e.calculateWalletSnapshot(req.Address, false, assets)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listSolWalletAssets] calculateWalletSnapshot() failed")
		return assets, "", "", nil
	}

	return assets, pnl, latestSnapshotBal, nil
}

func (e *Entity) listSuiWalletAssets(req request.ListWalletAssetsRequest) ([]response.WalletAssetData, string, string, error) {
	assets, err := e.svc.Sui.GetAddressAssets(req.Address)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to set get sui assets wallet")
		return nil, "", "", err
	}

	for index, asset := range assets {
		price, err := e.GetTokenPrice(asset.Token.Symbol, asset.Token.Name)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to set get price token from sui wallet")
			continue
		}
		assets[index].Token.Price = *price
		assets[index].UsdBalance = assets[index].Token.Price * assets[index].AssetBalance
	}

	// calculate pnl
	pnl, latestSnapshotBal, err := e.calculateWalletSnapshot(req.Address, false, assets)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listSuiWalletAssets] calculateWalletSnapshot() failed")
		return assets, "", "", nil
	}

	return assets, pnl, latestSnapshotBal, nil
}

func (e *Entity) ListWalletTxns(req request.ListWalletTransactionsRequest) ([]response.WalletTransactionData, error) {
	req.Standardize()

	// Solana
	if req.Type == "sol" {
		return e.listSolWalletTxns(req)
	}

	// Sui
	if req.Type == "sui" {
		return e.listSuiWalletTxns(req)
	}

	// ronin or EVM-based wallets
	return e.listEvmWalletTxns(req)
}

func (e *Entity) listEvmWalletTxns(req request.ListWalletTransactionsRequest) ([]response.WalletTransactionData, error) {
	chainIDs := []int{1, 56, 137, 250, 2020}
	txns := make([]response.WalletTransactionData, 0)
	for _, chainID := range chainIDs {
		res, err := e.svc.Covalent.GetTransactionsByAddress(chainID, req.Address, 5, 5)
		if err != nil {
			e.log.Fields(logger.Fields{"chainID": chainID, "address": req.Address}).Error(err, "[entity.listEvmWalletTxns] svc.Covalent.GetTransactionsByAddress() failed")
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
				e.log.Fields(logger.Fields{"item": item, "tx": tx, "address": req.Address}).Error(err, "[entity.listEvmWalletTxns] entity.parseCovalentTxData() failed")
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

func (e *Entity) listSuiWalletTxns(req request.ListWalletTransactionsRequest) ([]response.WalletTransactionData, error) {
	res, err := e.svc.Sui.GetAddressTxn(req.Address)
	if err != nil {
		return []response.WalletTransactionData{}, err
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
					_ids := ev.Decoded.Params[3].Value.([]interface{})
					_amounts := ev.Decoded.Params[4].Value.([]interface{})
					if contract.Name == "" {
						action.Unit = _ids[0].(map[string]interface{})["value"].(string)
						amt, err := strconv.ParseFloat(_amounts[0].(map[string]interface{})["value"].(string), 10)
						if err != nil {
							continue
						}
						action.Amount = amt
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

func (e *Entity) SumarizeBinanceAsset(req request.BinanceRequest) (*response.WalletBinanceResponse, error) {
	// redis cache
	value, err := e.cache.HashGet(fmt.Sprintf("binance-assets-%s-%s", req.Id, req.ApiKey))
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.SumarizeBinanceAsset] Failed to get cache user data binance")
		return nil, err
	}

	totalAssetValue := 0.0
	if len(value) == 0 {
		asset, err := e.SummarizeFundingAsset(req.Id, req.ApiKey, req.ApiSecret)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.SumarizeBinanceAsset] Failed to get binance asset")
			return nil, err
		}

		for _, asset := range asset {
			assetValue, err := strconv.ParseFloat(asset.BtcValuation, 64)
			if err != nil {
				e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.SumarizeBinanceAsset] Failed to parse asset value")
				return nil, err
			}

			totalAssetValue += assetValue
		}

		encodeData := map[string]string{
			"total_asset": fmt.Sprint(totalAssetValue),
		}

		err = e.cache.HashSet(fmt.Sprintf("binance-assets-%s-%s", req.Id, req.ApiKey), encodeData, 30*time.Second)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "Failed to set cache data wallet")
			return nil, err
		}
	} else {
		totalAssetValue, err = strconv.ParseFloat(value["total_asset"], 64)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.SumarizeBinanceAsset] Failed to parse asset value")
			return nil, err
		}
	}

	// btc price
	btcPrice, err := e.svc.CoinGecko.GetCoinPrice([]string{"bitcoin"}, "usd")
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.SumarizeBinanceAsset] Failed to get btc price")
		return nil, err
	}

	return &response.WalletBinanceResponse{
		TotalBtc: totalAssetValue,
		Price:    btcPrice["bitcoin"],
	}, err
}

func (e *Entity) GetBinanceAssets(req request.GetBinanceAssetsRequest) (*response.GetBinanceAsset, error) {
	profile, err := e.svc.MochiProfile.GetByID(req.Id)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.GetBinanceAssets] Failed to get profile")
		return nil, err
	}

	apiKey, apiSecret := "", ""
	for _, acc := range profile.AssociatedAccounts {
		if acc.Platform == consts.PlatformBinance {
			apiKey = acc.PlatformIdentifier
			apiSecret = acc.PlatformMetadata.ApiSecret
		}
	}

	if apiKey == "" || apiSecret == "" {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.GetBinanceAssets] Failed to get api key or api secret")
		return nil, baseerr.ErrProfileNotLinkBinance
	}

	// get data asset from binance or cache
	fundingAsset, err := e.SummarizeFundingAsset(req.Id, apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.GetBinanceAssets] Failed to get binance asset")
		return nil, err
	}

	earnAsset, err := e.SummarizeEarnAsset(req.Id, apiKey, apiSecret)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.GetBinanceAssets] Failed to get binance asset")
		return nil, err
	}

	// format asset
	formatFundingAsset, err := e.FormatAsset(fundingAsset)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.GetBinanceAssets] Failed to format asset")
		return nil, err
	}

	formatEarnAsset, err := e.FormatAsset(earnAsset)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entities.GetBinanceAssets] Failed to format asset")
		return nil, err
	}

	return &response.GetBinanceAsset{
		Asset: formatFundingAsset,
		Earn:  formatEarnAsset,
	}, nil
}

func (e *Entity) ListWalletFarmings(req request.ListWalletAssetsRequest) ([]response.LiquidityPosition, error) {
	req.Standardize()

	// Ronin
	if req.Type == "ron" {
		return e.listRoninFarmings(req)
	}

	// TODO: support Solana, Sui & other EVM-based chains
	return nil, nil
}

func (e *Entity) listRoninFarmings(req request.ListWalletAssetsRequest) ([]response.LiquidityPosition, error) {

	l := e.log.Fields(logger.Fields{"req": req})

	// no cache -> re-fetch
	res, err := e.svc.Skymavis.GetAddressFarming(strings.ToLower(req.Address))
	if err != nil {
		l.Error(err, "[entity.listRoninFarmings] svc.Skymavis.GetAddressFarming() failed")
		return nil, err
	}

	if res == nil || res.Data == nil {
		l.Error(err, "[entity.listRoninFarmings] svc.Skymavis.GetAddressFarming() response nil")
		return nil, nil
	}

	rewards, err := e.svc.Ronin.GetLpPendingRewards(req.Address)
	if err != nil {
		l.Error(err, "[entity.listRoninFarmings] svc.Ronin.GetLpPendingRewards() failed")
		return nil, err
	}

	for i, p := range res.Data.LiquidityPositions {
		totalLiquidity, err := strconv.ParseFloat(p.Pair.ReserveUSD, 64)
		if err != nil {
			l.Error(err, "[entity.listRoninFarmings] parse reserveUSD failed")
			return nil, err
		}

		totalSupply, err := strconv.ParseFloat(p.Pair.TotalSupply, 64)
		if err != nil {
			l.Error(err, "[entity.listRoninFarmings] parse supply failed")
			return nil, err
		}

		lpTokenBalance, err := strconv.ParseFloat(p.LiquidityTokenBalance, 64)
		if err != nil {
			l.Error(err, "[entity.listRoninFarmings] parse lp token balance failed")
			return nil, err
		}

		token0Price, err := strconv.ParseFloat(p.Pair.Token0.TokenDayData[0].PriceUSD, 64)
		if err != nil {
			l.Error(err, "[entity.listRoninFarmings] parse token0 price failed")
			return nil, err
		}

		token1Price, err := strconv.ParseFloat(p.Pair.Token1.TokenDayData[0].PriceUSD, 64)
		if err != nil {
			l.Error(err, "[entity.listRoninFarmings] parse token1 price failed")
			return nil, err
		}

		lpTokenWorth := totalLiquidity / totalSupply
		lpLiquidityWorth := lpTokenWorth * lpTokenBalance
		token0Balance := lpLiquidityWorth / 2 / token0Price
		token1Balance := lpLiquidityWorth / 2 / token1Price
		res.Data.LiquidityPositions[i].Pair.Token0.Balance = token0Balance
		res.Data.LiquidityPositions[i].Pair.Token1.Balance = token1Balance

		poolReward := rewards[strings.ToLower(p.Pair.ID)]
		// reward token is either one of token pair
		rewardToken := p.Pair.Token0
		if strings.EqualFold(poolReward.RewardToken, p.Pair.Token1.ID) {
			rewardToken = p.Pair.Token1
		}

		res.Data.LiquidityPositions[i].Reward = response.WalletFarmingReward{
			Amount: poolReward.Reward,
			Token:  rewardToken,
		}
	}

	return res.Data.LiquidityPositions, nil
}

func (e *Entity) ListWalletStakings(req request.ListWalletAssetsRequest) ([]response.WalletStakingData, error) {
	req.Standardize()

	// Ronin
	if req.Type == "ron" {
		return e.listRoninStakings(req)
	}

	// TODO: support Solana, Sui & other EVM-based chains
	return nil, nil
}

func (e *Entity) listRoninStakings(req request.ListWalletAssetsRequest) ([]response.WalletStakingData, error) {
	var res []response.WalletStakingData

	l := e.log.Fields(logger.Fields{"req": req})

	// no cache -> re-fetch
	axsStakingAmount, err := e.svc.Ronin.GetAxsStakingAmount(req.Address)
	if err != nil {
		l.Error(err, "[entity.listRoninWalletStakings] svc.Ronin.GetAxsStakingAmount failed")
		return nil, err
	}

	axsRewards, err := e.svc.Ronin.GetAxsPendingRewards(req.Address)
	if err != nil {
		l.Error(err, "[entity.listRoninWalletStakings] svc.Ronin.GetAxsPendingRewards failed")
		return nil, err
	}

	ronStakingAmount, err := e.svc.Ronin.GetRonStakingAmount(req.Address)
	if err != nil {
		l.Error(err, "[entity.listRoninWalletStakings] svc.Ronin.GetRonStakingAmount failed")
		return nil, err
	}

	ronRewrds, err := e.svc.Ronin.GetRonPendingRewards(req.Address)
	if err != nil {
		l.Error(err, "[entity.listRoninWalletStakings] svc.Ronin.GetRonPendingRewards failed")
		return nil, err
	}

	// axsData, err, _ := e.svc.CoinGecko.GetCoin("axie-infinity")
	// if err != nil {
	// 	l.Error(err, "[entity.listRoninWalletStakings] svc.CoinGecko.GetCoin('axie-infinity') failed")
	// }

	// roninData, err, _ := e.svc.CoinGecko.GetCoin("ronin")
	// if err != nil {
	// 	l.Error(err, "[entity.listRoninWalletStakings] svc.CoinGecko.GetCoin(ronin) failed")
	// }

	prices, err := e.svc.CoinGecko.GetCoinPrice([]string{"axie-infinity", "ronin"}, "usd")
	if err != nil {
		l.Error(err, "[entity.listRoninWalletStakings] svc.CoinGecko.GetCoinPrice failed")
	}

	res = []response.WalletStakingData{
		{
			TokenName: "Axie Infinity",
			Symbol:    strings.ToUpper("axs"),
			Amount:    axsStakingAmount,
			Reward:    axsRewards,
			Price:     prices["axie-infinity"],
		},
		{
			TokenName: "Ronin",
			Symbol:    strings.ToUpper("ron"),
			Amount:    ronStakingAmount,
			Reward:    ronRewrds,
			Price:     prices["ronin"],
		},
	}

	return res, nil
}

func (e *Entity) ListWalletNfts(req request.ListWalletAssetsRequest) ([]response.WalletNftData, error) {
	req.Standardize()

	// Ronin
	if req.Type == "ron" {
		return e.listAxieNfts(req)
	}

	// TODO: support Solana, Sui & other EVM-based chains
	return nil, nil
}

func (e *Entity) listAxieNfts(req request.ListWalletAssetsRequest) ([]response.WalletNftData, error) {
	var result []response.WalletNftData

	// no cache -> re-fetch
	res, err := e.svc.Skymavis.GetOwnedNfts(strings.ToLower(req.Address))
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.listAxieNfts] svc.Skymavis.GetOwnedAxies() failed")
		return nil, err
	}

	axies := make([]response.WalletNftMetadata, len(res.Data.Axies.Results))
	accessories := make([]response.WalletNftMetadata, len(res.Data.Equipments.Results))
	lands := make([]response.WalletNftMetadata, len(res.Data.Lands.Results))
	items := make([]response.WalletNftMetadata, len(res.Data.Items.Results))
	result = []response.WalletNftData{
		{Total: res.Data.Axies.Total, Tokens: axies, CollectionName: "Axie"},
		{Total: res.Data.Equipments.Total, Tokens: accessories, CollectionName: "Axie Accessory"},
		{Total: res.Data.Lands.Total, Tokens: lands, CollectionName: "Axie Land"},
		{Total: res.Data.Items.Total, Tokens: items, CollectionName: "Axie Land Item"},
	}
	// Axie collection
	for i, axie := range res.Data.Axies.Results {
		axies[i] = response.WalletNftMetadata{
			MarketplaceURL: fmt.Sprintf("https://app.axieinfinity.com/marketplace/axies/%s/", axie.ID),
			TokenName:      axie.Name,
			Image:          axie.Image,
		}
	}

	// Accessory collection
	for i, acc := range res.Data.Equipments.Results {
		accessories[i] = response.WalletNftMetadata{
			Image:          fmt.Sprintf("https://cdn.axieinfinity.com/marketplace-website/accessories/%s.png", acc.Alias),
			MarketplaceURL: fmt.Sprintf("https://app.axieinfinity.com/marketplace/accessories/%s/my-inventory/", acc.Alias),
			TokenName:      acc.Name,
		}
	}

	// Land collection
	for i, land := range res.Data.Lands.Results {
		lands[i] = response.WalletNftMetadata{
			Image:          fmt.Sprintf("https://cdn.axieinfinity.com/avatars/land/square/square_%d_%d.png", land.Col, land.Row),
			MarketplaceURL: fmt.Sprintf("https://app.axieinfinity.com/marketplace/lands/%d/%d/", land.Col, land.Row),
			TokenName:      fmt.Sprintf("%s plot (%d, %d)", land.LandType, land.Col, land.Row),
		}
	}

	// Land Item collection
	for i, item := range res.Data.Items.Results {
		items[i] = response.WalletNftMetadata{
			Image:          item.FigureURL,
			MarketplaceURL: fmt.Sprintf("https://app.axieinfinity.com/marketplace/items/%s/%d/", item.Alias, item.ItemID),
			TokenName:      fmt.Sprintf("%s #%d", item.Name, item.ItemID),
		}
	}

	return result, nil
}

func (e *Entity) calculateEvmTokenBalance(item krystal.Balance, chainID int) (bal, quote float64) {
	balance, ok := new(big.Float).SetString(item.Balance)
	if !ok {
		return
	}
	parsedBal, _ := balance.Float64()
	bal = parsedBal / math.Pow10(item.Token.Decimals)
	if strings.EqualFold(item.Token.Symbol, "icy") && chainID == 137 {
		quote = 1.5 * bal
	} else {
		quote = item.Quotes.Usd.Value
	}
	return
}

func (e *Entity) UpdateTrackingInfo(req request.UpdateTrackingInfoRequest) (*model.UserWalletWatchlistItem, error) {
	tx, ff := e.repo.Store.NewTransaction()

	var (
		wallet *model.UserWalletWatchlistItem
		err    error
	)
	if req.Alias != "" {
		wallet, err = tx.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{
			UserID:    req.UserID,
			Query:     req.Alias,
			ForUpdate: false,
		})
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.UpdateTrackingInfo] tx.UserWalletWatchlistItem.GetOne failed")
			return nil, ff.Rollback(err)
		}
	}

	if wallet != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.UpdateTrackingInfo] alias already exists")
		return nil, ff.Rollback(baseerr.ErrAliasAlreadyExisted)
	}

	wallet, err = tx.UserWalletWatchlistItem.GetOne(userwalletwatchlistitem.GetOneQuery{
		UserID:    req.UserID,
		Query:     req.Address,
		ForUpdate: true,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.UpdateTrackingInfo] tx.UserWalletWatchlistItem.GetOne failed")

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ff.Rollback(baseerr.ErrRecordNotFound)
		}
		return nil, ff.Rollback(err)
	}

	wallet.Alias = req.Alias

	if err := tx.UserWalletWatchlistItem.Update(wallet); err != nil {
		e.log.Fields(logger.Fields{"wallet": wallet}).Error(err, "[entity.UpdateTrackingInfo] tx.UserWalletWatchlistItem.Update failed")
		return nil, ff.Rollback(err)
	}

	return wallet, ff.Commit()
}
