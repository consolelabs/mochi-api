package entities

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/util"
	"gorm.io/gorm"
)

func (e *Entity) SendNftSalesToChannel(nftSale request.HandleNftWebhookRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(nftSale.CollectionAddress)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.handleNotAddedCollection(nftSale)
		} else {
			e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] cannot get collection by address %s", nftSale.CollectionAddress)
			return err
		}

	}

	indexerTokenRes, err := e.indexer.GetNFTDetail(nftSale.CollectionAddress, nftSale.TokenId)
	if err != nil {
		e.log.Errorf(err, "[indexer.GetNFTDetail] cannot get token from indexer by address %s and token %s", nftSale.CollectionAddress, nftSale.TokenId)
		return err
	}
	indexerToken := indexerTokenRes.Data

	// calculate last price, price, pnl, sub pnl
	price := util.StringWeiToEther(nftSale.Price.Amount, nftSale.Price.Token.Decimal)
	lastPrice := util.StringWeiToEther(nftSale.LastPrice.Amount, nftSale.LastPrice.Token.Decimal)
	pnl := new(big.Float)
	pnl = pnl.Sub(price, lastPrice)
	subPnl := new(big.Float).Quo(pnl, lastPrice)
	subPnlPer := subPnl.Mul(subPnl, big.NewFloat(100))

	subPnlDisplay := ""
	if util.FormatCryptoPrice(*lastPrice) != "0" {
		subPnlDisplay = " `" + util.GetChangePnl(pnl) + " " + fmt.Sprintf("%.2f", subPnlPer.Abs(subPnlPer)) + "%`"
	}

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

	// handle marketplace
	marketplace := strings.ToUpper(string(nftSale.Marketplace[0])) + nftSale.Marketplace[1:]
	marketplaceLink := ""
	if strings.ToLower(nftSale.Marketplace) == "opensea" {
		res, err := e.marketplace.GetOpenseaAssetContract(nftSale.CollectionAddress)
		if err != nil {
			e.log.Errorf(err, "[marketplace.GetOpenseaAssetContrace] cannot get opensea data")
			return fmt.Errorf("cannot get opensea data. Error: %v", err)
		}
		marketplaceLink = "[" + marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + res.Collection.UrlName + ")"
	} else {
		marketplaceLink = "[" + marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + strings.ToLower(nftSale.CollectionAddress) + ")"
	}

	// handle image
	image := indexerToken.ImageCDN
	if image == "" {
		image = indexerToken.Image
	}

	data := []*discordgo.MessageEmbedField{
		{
			Name:   "Rarity",
			Value:  rarityDisplay,
			Inline: true,
		},
		{

			Name:   "Rank",
			Value:  rankDisplay,
			Inline: true,
		},
		{
			Name:   "Marketplace",
			Value:  marketplaceLink,
			Inline: true,
		},
		{
			Name:   "Transaction",
			Value:  "[" + util.ShortenAddress(nftSale.Transaction) + "]" + "(" + util.GetTransactionUrl(nftSale.Marketplace) + strings.ToLower(nftSale.Transaction) + ")",
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
			Name:   "Hodl",
			Value:  strconv.Itoa(util.SecondsToDays(nftSale.Hodl)) + " days",
			Inline: true,
		},
		{
			Name:   "Price",
			Value:  util.FormatCryptoPrice(*price) + " " + strings.ToUpper(nftSale.Price.Token.Symbol),
			Inline: true,
		},
		{
			Name:   "Last Price",
			Value:  util.FormatCryptoPrice(*lastPrice) + " " + strings.ToUpper(nftSale.LastPrice.Token.Symbol),
			Inline: true,
		},
		{
			Name: "PnL",
			// + " " + strings.ToUpper(nftSale.Price.Token.Symbol)
			Value:  util.GetGainEmoji(pnl) + util.FormatCryptoPrice(*pnl) + subPnlDisplay,
			Inline: true,
		},
	}
	// finalize message nft sales
	messageSale := []*discordgo.MessageEmbed{{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    collection.Name,
			IconURL: collection.Image,
		},
		Fields:      data,
		Description: indexerToken.Name + " sold!",
		Color:       int(util.RarityColors(rarityRate)),
		Image: &discordgo.MessageEmbedImage{
			URL: util.StandardizeUri(image),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}}
	resp, _ := e.GetAllNFTSalesTracker()

	for _, saleChannel := range resp {
		if saleChannel.ContractAddress == "*" || nftSale.CollectionAddress == saleChannel.ContractAddress {
			_, err := e.discord.ChannelMessageSendEmbeds(saleChannel.ChannelID, messageSale)
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to sale channel. CollectionName: %s, TokenName: %s", collection.Name, indexerToken.Name)
				return fmt.Errorf("cannot send message to sale channel. Error: %v", err)
			}

			sub := ""
			if util.FormatCryptoPrice(*lastPrice) != "0" {
				sub = util.GetChangePnl(pnl) + fmt.Sprintf("%.2f", subPnlPer.Abs(subPnlPer))
			}
			// add sales message to database
			err = e.HandleMochiSalesMessage(&request.TwitterSalesMessage{
				TokenName:         indexerToken.Name,
				CollectionName:    collection.Name,
				Price:             util.FormatCryptoPrice(*price) + " " + strings.ToUpper(nftSale.Price.Token.Symbol),
				SellerAddress:     util.ShortenAddress(nftSale.From),
				BuyerAddress:      util.ShortenAddress(nftSale.To),
				Marketplace:       marketplace,
				MarketplaceURL:    util.GetStringBetweenParentheses(marketplaceLink),
				Image:             image,
				TxURL:             util.GetTransactionUrl(nftSale.Marketplace) + strings.ToLower(nftSale.Transaction),
				CollectionAddress: collection.Address,
				TokenID:           indexerToken.TokenID,
				SubPnl:            sub,
				Pnl:               util.FormatCryptoPrice(*pnl),
				Hodl:              strconv.Itoa(util.SecondsToDays(nftSale.Hodl)),
			})
			if err != nil {
				e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot handle mochi sales msg. CollectionName: %s, TokenName: %s", collection.Name, indexerToken.Name)
				return fmt.Errorf("cannot handle mochi sales msg. Error: %v", err)
			}
		}
	}
	return nil
}

func (e *Entity) handleNotAddedCollection(nftSale request.HandleNftWebhookRequest) error {
	// convert marketplace to chain id
	chainID := util.ConvertMarkplaceToChainId(nftSale.Marketplace)

	// query name and symbol from contract
	name, symbol, err := e.abi.GetNameAndSymbol(nftSale.CollectionAddress, int64(chainID))
	if err != nil {
		e.log.Errorf(err, "[e.abi.GetNameAndSymbol] cannot get name and symbol of contract: %s | chainId %d", nftSale.CollectionAddress, chainID)
		return err
	}

	// get image from marketplace
	image, err := e.getImageFromMarketPlace(int(chainID), nftSale.CollectionAddress)
	if err != nil {
		e.log.Errorf(err, "[e.getImageFromMarketPlace] failed to get image from market place: %v", err)
		return err
	}

	// add indexer
	err = e.indexer.CreateERC721Contract(indexer.CreateERC721ContractRequest{
		Address: nftSale.CollectionAddress,
		ChainID: int(chainID),
	})
	if err != nil && err.Error() != "block number not synced yet, TODO: add to queue and try later" {
		e.log.Errorf(err, "[CreateERC721Contract] failed to create erc721 contract: %v", err)
		return nil
	}
	// add collection
	_, err = e.repo.NFTCollection.Create(model.NFTCollection{
		Address:    nftSale.CollectionAddress,
		Symbol:     symbol,
		Name:       name,
		ChainID:    strconv.Itoa(int(chainID)),
		ERCFormat:  "ERC721",
		IsVerified: true,
		Image:      image,
	})
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.Create] cannot add collection: %v", err)
		return err
	}

	// notify added collection
	err = e.svc.Discord.NotifyAddNewCollection("962589711841525780", name, symbol, util.ConvertChainIDToChain(strconv.Itoa(int(chainID))), image)
	if err != nil {
		e.log.Errorf(err, "[e.svc.Discord.NotifyAddNewCollection] cannot send embed message: %v", err)
		return err
	}
	return nil
}

func (e *Entity) SendNftAddedCollection(nftAddedCollection request.HandleNftWebhookRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(nftAddedCollection.CollectionAddress)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] cannot get collection. CollectionAddress: %s", nftAddedCollection.CollectionAddress)
		return err
	}

	channelNewCollection := "964780299307343912"
	messageAddedNewCollection := []*discordgo.MessageEmbed{{
		Title:       "New collection: " + collection.Name,
		Description: "We're happy to announce that " + collection.Name + " ranking is available.\n\n" + "You can check your rank using:\n" + "`$nft " + strings.ToLower(collection.Symbol) + " <token_id>`\n\n" + ":warning: Remeber that ranks are calculated using metadata, wrong and bad metadata can impact ranks as well.\n:warning:Ranks are not a financial indicator.\n",
		Color:       0xFCD3C1,
		Timestamp:   time.Now().Format(time.RFC3339),
		Image: &discordgo.MessageEmbedImage{
			URL: collection.Image,
		},
	}}

	_, err = e.discord.ChannelMessageSendEmbeds(channelNewCollection, messageAddedNewCollection)
	if err != nil {
		e.log.Errorf(err, "[discord.ChannelMessageSendEmbeds] cannot send message to new added collection channel. CollectionAddress: %s, Chain: %s", nftAddedCollection.CollectionAddress, nftAddedCollection.ChainId)
		return fmt.Errorf("cannot send message to new added collection channel. Error: %v", err)
	}
	return nil
}

func (e *Entity) SendStealAlert(price float64, address string, marketplace string, token string, image string, name string) error {
	var floor float64 = 0
	var average float64 = 0
	url := ""
	switch marketplace {
	case "opensea":
		//ETH: opensea -> asseet_contract -> slug -> collection/slug -> floor
		res, err := e.marketplace.GetOpenseaAssetContract(address)
		if err != nil {
			return err
		}

		collection, err := e.marketplace.GetCollectionFromOpensea(res.Collection.UrlName)
		if err != nil {
			return err
		}

		floor = collection.Collection.Stats.FloorPrice
		average = collection.Collection.Stats.AveragePrice
		url = fmt.Sprintf("https://opensea.io/assets/ethereum/%s/%s", address, token)

	case "paintswap":
		// FTM: paintswap
		res, err := e.marketplace.GetCollectionFromPaintswap(address)
		if err != nil {
			return err
		}
		floorPrice, _ := util.StringWeiToEther(res.Collection.Stats.FloorPrice, 18).Float64()
		floor = floorPrice
		avgPrice, _ := util.StringWeiToEther(res.Collection.Stats.AveragePrice, 18).Float64()
		average = avgPrice
		url = fmt.Sprintf("https://paintswap.finance/marketplace/assets/%s/%s", address, token)

	case "optimism":
		// OP: quixotic -> collection/slug == collection/address
		res, err := e.marketplace.GetCollectionFromQuixotic(address)
		if err != nil {
			return err
		}
		length := len(strconv.Itoa(int(res.FloorPrice)))
		floorPrice, _ := util.StringWeiToEther(strconv.Itoa(int(res.FloorPrice)), length).Float64()
		floor = floorPrice
		//api does not have average price
		url = fmt.Sprintf("https://quixotic.io/asset/%s/%s", address, token)
	}
	if price < floor {
		err := e.svc.Discord.NotifyStealFloorPrice(price, floor, url, name, image)
		if err != nil {
			return err
		}
	} else if price < average {
		err := e.svc.Discord.NotifyStealAveragePrice(price, average, url, name, image)
		if err != nil {
			return err
		}
	}
	return nil
}
