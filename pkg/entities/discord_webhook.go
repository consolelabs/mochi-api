package entities

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/defipod/mochi/pkg/model"
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

func (e *Entity) createNftTokenModel(nftSale request.HandleNftWebhookRequest, collection *model.NFTCollection, indexerTokenRes *response.IndexerGetNFTTokenDetailResponse) (*nftTokenModel, error) {
	// calculate last price, price, pnl, sub pnl
	price := util.StringWeiToEther(nftSale.Price.Amount, nftSale.Price.Token.Decimal)
	lastPrice := util.StringWeiToEther(nftSale.LastPrice.Amount, nftSale.LastPrice.Token.Decimal)
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
		marketplaceLink = "[" + marketplace + "](" + util.GetURLMarketPlace(nftSale.Marketplace) + strings.ToLower(nftSale.CollectionAddress) + ")"
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
			Image:           "",
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

func (e *Entity) SendNftSalesToChannel(nftSale request.HandleNftWebhookRequest) error {
	collection, err := e.repo.NFTCollection.GetByAddress(nftSale.CollectionAddress)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] cannot get collection by address %s", nftSale.CollectionAddress)
		return err
	}

	indexerTokenRes, err := e.indexer.GetNFTDetail(nftSale.CollectionAddress, nftSale.TokenId)
	if err != nil {
		e.log.Infof("[indexer.GetNFTDetail] cannot get token from indexer by address %s and token %s", nftSale.CollectionAddress, nftSale.TokenId)
	}

	// create nft model for both case indexer has data or not
	nftToken, err := e.createNftTokenModel(nftSale, collection, indexerTokenRes)
	if err != nil {
		e.log.Errorf(err, "[createNftTokenModel] cannot create nft token model")
		return err
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
			Value:  util.FormatCryptoPrice(*nftToken.Price) + " " + strings.ToUpper(nftSale.Price.Token.Symbol),
			Inline: true,
		},
		{
			Name:   "Last Price",
			Value:  util.FormatCryptoPrice(*nftToken.LastPrice) + " " + strings.ToUpper(nftSale.LastPrice.Token.Symbol),
			Inline: true,
		},
		{
			Name: "PnL",
			// + " " + strings.ToUpper(nftSale.Price.Token.Symbol)
			Value:  util.GetGainEmoji(nftToken.Pnl) + util.FormatCryptoPrice(*nftToken.Pnl) + nftToken.SubPnlDisplay,
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
		Description: nftToken.Name + " sold!",
		Color:       int(util.RarityColors(nftToken.RarityRate)),
		Image: &discordgo.MessageEmbedImage{
			URL: util.StandardizeUri(nftToken.Image),
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}}
	resp, _ := e.GetAllNFTSalesTracker()

	for _, saleChannel := range resp {
		if saleChannel.ContractAddress == "*" || nftSale.CollectionAddress == saleChannel.ContractAddress {
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
