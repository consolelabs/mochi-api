package entities

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	nftaddrequesthistory "github.com/defipod/mochi/pkg/repo/nft_add_request_history"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type nftTokenModel struct {
	RarityDisplay   string
	MarketplaceLink string
	Price           *big.Float
	LastPrice       *big.Float
	Pnl             *big.Float
	SubPnlDisplay   string
	SubPnlPer       *big.Float
	Name            string
	RarityRate      string
	RankDisplay     string
	Image           string
	Marketplace     string
	TokenID         string
}

func (e *Entity) NotifyNftCollectionIntegration(req request.NotifyCompleteNftIntegrationRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		e.log.Error(err, "[entity.SendCollectionIntegrationToMochiLogs] repo.NFTCollection.GetByAddress() failed")
		return err
	}
	history, err := e.repo.NftAddRequestHistory.GetOne(nftaddrequesthistory.GetOneQuery{Address: req.CollectionAddress, ChainID: req.ChainID})
	if err != nil {
		e.log.Fields(logger.Fields{"address": req.CollectionAddress}).Error(err, "[entity.SendCollectionIntegrationToMochiLogs] repo.NftAddRequestHistory.GetOne() failed")
		return err
	}

	chain := strings.ToUpper(util.ConvertChainIDToChain(collection.ChainID))
	// send logs to mochi
	err = e.svc.Discord.NotifyCompleteCollectionIntegration(history.GuildID, collection.Name, collection.Symbol, chain, collection.Image)
	if err != nil {
		e.log.Error(err, "[entity.SendCollectionIntegrationToMochiLogs] svc.Discord.NotifyCompleteCollectionIntegration() failed")
		return err
	}

	// reply to orignal command
	description := fmt.Sprintf("👉 Your collection is being processed. We will let you know when it's ready to use.\n👉 To track other collection in `$nft list`, run `$nft %s <token_id>`", collection.Address)
	_, err = e.discord.ChannelMessageSendEmbedReply(history.ChannelID, &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://cdn.discordapp.com/emojis/977508805011181638.png?size=240&quality=lossless",
			Name:    fmt.Sprintf("%s has been added", collection.Name),
		},
		Description: description,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: collection.Image,
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}, &discordgo.MessageReference{
		GuildID:   history.GuildID,
		ChannelID: history.ChannelID,
		MessageID: history.MessageID,
	})

	return err
}

func (e *Entity) NotifyNftCollectionAdd(req request.NotifyNftCollectionAddRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		e.log.Error(err, "[entity.NotifyNftCollectionAdd] repo.NFTCollection.GetByAddress() failed")
		return err
	}
	history, err := e.repo.NftAddRequestHistory.GetOne(nftaddrequesthistory.GetOneQuery{Address: req.CollectionAddress, ChainID: req.ChainID})
	if err != nil {
		e.log.Fields(logger.Fields{"address": req.CollectionAddress}).Error(err, "[entity.NotifyNftCollectionAdd] repo.NftAddRequestHistory.GetOne() failed")
		return err
	}

	chain := strings.ToUpper(util.ConvertChainIDToChain(collection.ChainID))
	// send logs to mochi
	err = e.svc.Discord.NotifyAddNewCollection(history.GuildID, collection.Name, collection.Symbol, chain, collection.Image)
	if err != nil {
		e.log.Error(err, "[entity.NotifyNftCollectionAdd] svc.Discord.NotifyAddNewCollection() failed")
		return err
	}

	// reply to orignal command
	description := fmt.Sprintf("👉 Your collection is being processed. We will let you know when it's ready to use.\n👉 You can track other collections in `$nft list`.\n👉 You can track sales movement by running `$sales track <channel> %s %s`.", collection.Address, collection.ChainID)
	_, err = e.discord.ChannelMessageSendEmbedReply(history.ChannelID, &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://cdn.discordapp.com/emojis/977508805011181638.png?size=240&quality=lossless",
			Name:    fmt.Sprintf("%s has been added", collection.Name),
		},
		Description: description,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: collection.Image,
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}, &discordgo.MessageReference{
		GuildID:   history.GuildID,
		ChannelID: history.ChannelID,
		MessageID: history.MessageID,
	})

	return err
}

func (e *Entity) NotifyNftCollectionSync(req request.NotifyCompleteNftSyncRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		e.log.Error(err, "[entity.NotifyNftCollectionSync] repo.NFTCollection.GetByAddress() failed")
		return err
	}
	history, err := e.repo.NftAddRequestHistory.GetOne(nftaddrequesthistory.GetOneQuery{Address: req.CollectionAddress, ChainID: req.ChainID})
	if err != nil {
		e.log.Fields(logger.Fields{"address": req.CollectionAddress}).Error(err, "[entity.NotifyNftCollectionSync] repo.NftAddRequestHistory.GetOne() failed")
		return err
	}

	chain := strings.ToUpper(util.ConvertChainIDToChain(collection.ChainID))
	// send logs to mochi
	err = e.svc.Discord.NotifyCompleteCollectionSync(history.GuildID, collection.Name, collection.Symbol, chain, collection.Image)
	if err != nil {
		e.log.Error(err, "[entity.NotifyNftCollectionSync] svc.Discord.NotifyCompleteCollectionSync() failed")
		return err
	}

	// reply to orignal command
	description := fmt.Sprintf("👉 To check rarity, run `$nft %s <token_id>`.\n👉 To track sales, run `$sales track <channel> %s %s`.", collection.Symbol, collection.Address, collection.ChainID)
	_, err = e.discord.ChannelMessageSendEmbedReply(history.ChannelID, &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			IconURL: "https://cdn.discordapp.com/emojis/977508805011181638.png?size=240&quality=lossless",
			Name:    fmt.Sprintf("%s is ready to use", collection.Name),
		},
		Description: description,
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: collection.Image,
		},
		Timestamp: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}, &discordgo.MessageReference{
		GuildID:   history.GuildID,
		ChannelID: history.ChannelID,
		MessageID: history.MessageID,
	})

	return err
}

func (e *Entity) NotifySaleMarketplace(nftSale request.NotifySaleMarketplaceRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(nftSale.Address)
	if err != nil {
		e.log.Infof("Collection not exist yet, adding it to database")
		// TODO(trkhoi): handle for evm chain
		if nftSale.ChainId == 9999 || nftSale.ChainId == 9997 {
			e.CreateBluemoveNFTCollection(request.CreateNFTCollectionRequest{
				Address: nftSale.Address,
				ChainID: strconv.Itoa(int(nftSale.ChainId)),

				Author:       "393034938028392449",
				PriorityFlag: false,
			})
		}

		if nftSale.ChainId == 66 {
			e.CreateEVMNFTCollection(request.CreateNFTCollectionRequest{
				Address:      nftSale.Address,
				ChainID:      strconv.Itoa(int(nftSale.ChainId)),
				Author:       "393034938028392449",
				GuildID:      "891310117658705931",
				MessageID:    nftSale.Address,
				PriorityFlag: false,
			})
		}
		return nil
		// e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] cannot get collection by address %s", nftSale.Address)
		// return err
	}

	indexerTokenRes, err := e.indexer.GetNFTDetail(nftSale.Address, nftSale.TokenId)
	if err != nil {
		e.log.Infof("[indexer.GetNFTDetail] cannot get token from indexer by address %s and token %s", nftSale.Address, nftSale.TokenId)
	}

	// create nft model for both case indexer has data or not
	nftToken, err := e.createNftTokenModel(nftSale, collection, indexerTokenRes)
	if err != nil {
		e.log.Errorf(err, "[createNftTokenModel] cannot create nft token model")
		return err
	}

	txUrl := ""
	if nftSale.Transaction != "" {
		txUrl = "[" + util.ShortenAddress(nftSale.Transaction) + "]" + "(" + util.GetTransactionUrl(nftSale.Marketplace) + strings.ToLower(nftSale.Transaction) + ")"
	}

	data := []*discordgo.MessageEmbedField{
		{
			Name:   "Rarity",
			Value:  nftToken.RarityDisplay,
			Inline: true,
		},
		{

			Name:   "Rank",
			Value:  nftToken.RankDisplay,
			Inline: true,
		},
		{
			Name:   "Marketplace",
			Value:  nftToken.MarketplaceLink,
			Inline: true,
		},
		{
			Name:   "From",
			Value:  "[" + util.ShortenAddress(nftSale.From) + "]" + "(" + util.GetWalletUrl(nftSale.Marketplace) + strings.ToLower(nftSale.From) + ")",
			Inline: true,
		},
		{
			Name:   "To",
			Value:  "[" + util.ShortenAddress(nftSale.To) + "]" + "(" + util.GetWalletUrl(nftSale.Marketplace) + strings.ToLower(nftSale.To) + ")",
			Inline: true,
		},
		{
			Name:   "Price",
			Value:  util.FormatCryptoPrice(*nftToken.Price) + " " + strings.ToUpper(nftSale.Price.Token.Symbol),
			Inline: true,
		},
		{
			Name:   "PnL",
			Value:  util.GetGainEmoji(nftToken.Pnl) + util.FormatCryptoPrice(*nftToken.Pnl) + nftToken.SubPnlDisplay,
			Inline: true,
		},
	}

	if util.SecondsToDays(nftSale.Hodl) > 0 {
		data = append(data, &discordgo.MessageEmbedField{
			Name:   "Hodl",
			Value:  strconv.Itoa(util.SecondsToDays(nftSale.Hodl)) + " days",
			Inline: true,
		})
	}

	if txUrl != "" {
		data = append(data, &discordgo.MessageEmbedField{
			Name:   "Transaction",
			Value:  txUrl,
			Inline: true,
		})
	}

	if util.FormatCryptoPrice(*nftToken.LastPrice) != "0" {
		data = append(data, &discordgo.MessageEmbedField{
			Name:   "Last Price",
			Value:  util.FormatCryptoPrice(*nftToken.LastPrice) + " " + strings.ToUpper(nftSale.LastPrice.Token.Symbol),
			Inline: true,
		})
	}
	// finalize message nft sales
	messageSale := []*discordgo.MessageEmbed{{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    collection.Name,
			IconURL: collection.Image,
		},
		Fields:      data,
		Description: nftToken.Name + " sold!",
		Color:       int(util.RarityColors(nftToken.RarityRate)),
		Image: &discordgo.MessageEmbedImage{
			URL: util.StandardizeUri(nftToken.Image),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}}
	resp, _ := e.GetAllSalesTrackerConfig()

	for _, saleChannel := range resp {
		if (saleChannel.ContractAddress == "*" && saleChannel.Chain == "*") || (saleChannel.ContractAddress == "*" && saleChannel.Chain == util.ConvertChainIDToChain(collection.ChainID)) || nftSale.Address == saleChannel.ContractAddress {
			_, err := e.discord.ChannelMessageSendEmbeds(saleChannel.ChannelID, messageSale)
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to sale channel. CollectionName: %s, TokenName: %s", collection.Name, nftToken.Name)
				return fmt.Errorf("cannot send message to sale channel. Error: %v", err)
			}

			sub := ""
			if util.FormatCryptoPrice(*nftToken.LastPrice) != "0" {
				sub = util.GetChangePnl(nftToken.Pnl) + fmt.Sprintf("%.2f", nftToken.SubPnlPer.Abs(nftToken.SubPnlPer))
			}
			// add sales message to database
			err = e.HandleMochiSalesMessage(&request.TwitterSalesMessage{
				TokenName:         nftToken.Name,
				CollectionName:    collection.Name,
				Price:             util.FormatCryptoPrice(*nftToken.Price) + " " + strings.ToUpper(nftSale.Price.Token.Symbol),
				SellerAddress:     util.ShortenAddress(nftSale.From),
				BuyerAddress:      util.ShortenAddress(nftSale.To),
				Marketplace:       nftToken.Marketplace,
				MarketplaceURL:    util.GetStringBetweenParentheses(nftToken.MarketplaceLink),
				Image:             nftToken.Image,
				TxURL:             util.GetTransactionUrl(nftSale.Marketplace) + strings.ToLower(nftSale.Transaction),
				CollectionAddress: collection.Address,
				TokenID:           nftToken.TokenID,
				SubPnl:            sub,
				Pnl:               util.FormatCryptoPrice(*nftToken.Pnl),
				Hodl:              strconv.Itoa(util.SecondsToDays(nftSale.Hodl)),
			})
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot handle mochi sales msg. CollectionName: %s, TokenName: %s", collection.Name, nftToken.Name)
				return fmt.Errorf("cannot handle mochi sales msg. Error: %v", err)
			}
		}
	}
	return nil
}

func (e *Entity) createNftTokenModel(nftSale request.NotifySaleMarketplaceRequest, collection *model.NFTCollection, indexerTokenRes *response.IndexerGetNFTTokenDetailResponse) (*nftTokenModel, error) {
	// calculate last price, price, pnl, sub pnl
	price := util.StringWeiToEther(nftSale.Price.Amount, nftSale.Price.Token.Decimals)
	lastPrice := util.StringWeiToEther(nftSale.LastPrice.Amount, nftSale.LastPrice.Token.Decimals)
	pnl := new(big.Float)
	pnl = pnl.Sub(price, lastPrice)

	subPnl := new(big.Float)
	if lastPrice.Cmp(big.NewFloat(0)) == 0 {
		subPnl = big.NewFloat(0)
	} else {
		subPnl = new(big.Float).Quo(pnl, lastPrice)
	}

	subPnlPer := subPnl.Mul(subPnl, big.NewFloat(100))

	subPnlDisplay := ""
	if util.FormatCryptoPrice(*lastPrice) != "0" {
		subPnlDisplay = " `" + util.GetChangePnl(pnl) + " " + fmt.Sprintf("%.2f", subPnlPer.Abs(subPnlPer)) + "%`"
	}

	// handle marketplace
	marketplace := strings.ToUpper(string(nftSale.Marketplace[0])) + nftSale.Marketplace[1:]
	marketplaceLink := ""
	if strings.ToLower(nftSale.Marketplace) == "opensea" {
		// TODO(trkhoi): renew expired opensea api key
		// res, err := e.marketplace.GetOpenseaAssetContract(nftSale.CollectionAddress)
		// if err != nil {
		// 	e.log.Errorf(err, "[marketplace.GetOpenseaAssetContrace] cannot get opensea data")
		// 	return nil, fmt.Errorf("cannot get opensea data. Error: %v", err)
		// }
		// res.Collection.UrlName

		marketplaceLink = "[" + marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + "test_opensea" + ")"
	} else {
		marketplaceLink = "[" + marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + strings.ToLower(nftSale.Address) + ")"
	}

	// case indexer not have data in nft_token -> return
	if indexerTokenRes == nil {
		return &nftTokenModel{
			RarityDisplay:   "N/A",
			MarketplaceLink: marketplaceLink,
			RankDisplay:     "N/A",
			Price:           price,
			LastPrice:       lastPrice,
			Pnl:             pnl,
			SubPnlDisplay:   subPnlDisplay,
			SubPnlPer:       subPnlPer,
			Name:            collection.Name + " #" + nftSale.TokenId,
			RarityRate:      "N/A",
			Image:           nftSale.Image,
			Marketplace:     marketplace,
			TokenID:         nftSale.TokenId,
		}, nil
	}

	indexerToken := indexerTokenRes.Data

	// handle rarity, rank
	rankDisplay := ""
	rarityDisplay := ""
	rarityRate := ""

	if indexerToken.Rarity == nil {
		rankDisplay = "N/A"
		rarityDisplay = "N/A"
		rarityRate = ""
	} else {
		if indexerToken.Rarity.Rarity == "" {
			rarityDisplay = "N/A"
			rarityRate = ""
		} else {
			rarityDisplay = indexerToken.Rarity.Rarity
			rarityDisplay = util.RarityEmoji(rarityDisplay) + " " + rarityDisplay
			rarityRate = indexerToken.Rarity.Rarity
		}

		if indexerToken.Rarity.Rank == 0 {
			rankDisplay = "N/A"
		} else {
			rankDisplay = strconv.Itoa(int(indexerToken.Rarity.Rank))
			rankDisplay = "<:cup:985137841027821589> " + rankDisplay
		}
	}

	// handle image
	image := indexerToken.ImageCDN
	if image == "" {
		image = indexerToken.Image
	}

	return &nftTokenModel{
		RarityDisplay:   rarityDisplay,
		MarketplaceLink: marketplaceLink,
		Price:           price,
		RankDisplay:     rankDisplay,
		LastPrice:       lastPrice,
		Pnl:             pnl,
		SubPnlDisplay:   subPnlDisplay,
		SubPnlPer:       subPnlPer,
		Name:            indexerToken.Name,
		RarityRate:      rarityRate,
		Image:           image,
		Marketplace:     marketplace,
		TokenID:         nftSale.TokenId,
	}, nil
}
