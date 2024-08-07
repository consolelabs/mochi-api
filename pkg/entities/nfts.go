package entities

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/consolelabs/mochi-typeset/common/contract/abi/typeset"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/consts"
	"github.com/defipod/mochi/pkg/contract/nekostaking"
	"github.com/defipod/mochi/pkg/contracts/erc721"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	baseerrs "github.com/defipod/mochi/pkg/model/errors"
	"github.com/defipod/mochi/pkg/repo/user_nft_balance"
	usernftwatchlistitem "github.com/defipod/mochi/pkg/repo/user_nft_watchlist_items"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/service/mochiprofile"
	"github.com/defipod/mochi/pkg/util"
	maputils "github.com/defipod/mochi/pkg/util/map"
	sliceutils "github.com/defipod/mochi/pkg/util/slice"
)

var (
	mapChainChainId = map[string]string{
		"eth":    "1",
		"heco":   "128",
		"bsc":    "56",
		"matic":  "137",
		"op":     "10",
		"btt":    "199",
		"okt":    "66",
		"movr":   "1285",
		"celo":   "42220",
		"metis":  "1088",
		"cro":    "25",
		"xdai":   "0x64",
		"boba":   "288",
		"ftm":    "250",
		"avax":   "0xa86a",
		"arb":    "42161",
		"aurora": "1313161554",
		"base":   "8453",
	}
)

func (e *Entity) GetNFTDetail(symbol, tokenID, guildID string, queryByAddress bool) (*response.IndexerGetNFTTokenDetailResponseWithSuggestions, error) {
	suggest := []response.CollectionSuggestions{}
	// handle query by address
	// TODO(trkhoi): find way to detect non-evem address
	if queryByAddress {
		data, err := e.GetNFTDetailByAddress(symbol, tokenID)
		if err != nil {
			e.log.Errorf(err, "[e.GetNFTDetailByAddress] failed to get nft collection by address")
			return nil, err
		}
		return data, nil
	}

	// get collection
	collections, err := e.repo.NFTCollection.GetBySymbolorName(symbol)
	// cannot find collection => return suggested collections
	if err != nil || len(collections) == 0 {
		suggest, err = e.GetNFTSuggestion(symbol, tokenID)
		if err != nil {
			e.log.Errorf(err, "[repo.NFTCollection.GetBySymbolorName] failed to get nft collection by symbol %s", symbol)
			return nil, err
		}
		return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
			Data:        &response.IndexerNFTTokenDetailData{},
			Suggestions: suggest,
		}, nil
	}

	// found multiple symbols => only suggest those
	if len(collections) != 1 {
		var defaultSymbol *response.CollectionSuggestions
		// check default symbol
		symbols, _ := e.GetDefaultCollectionSymbol(guildID)
		for _, col := range collections {
			// if found default symbol
			if len(symbols) != 0 && checkIsDefaultSymbol(symbols, &col) {
				def := response.CollectionSuggestions{
					Address: col.Address,
					Chain:   util.ConvertChainIDToChain(col.ChainID),
					Name:    col.Name,
					Symbol:  col.Symbol,
				}
				defaultSymbol = &def
			}
			suggest = append(suggest, response.CollectionSuggestions{
				Name:    col.Name,
				Symbol:  col.Symbol,
				Address: col.Address,
				Chain:   util.ConvertChainIDToChain(col.ChainID),
			})
		}
		return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
			Data:          &response.IndexerNFTTokenDetailData{},
			Suggestions:   suggest,
			DefaultSymbol: defaultSymbol,
		}, nil
	}

	// db returned 1 collection
	collection := collections[0]
	data, err := e.getTokenDetailFromIndexer(collection.Address, tokenID)
	if err != nil {
		e.log.Errorf(err, "[e.getTokenDetailFromIndexer] failed to get nft indexer detail")
		return nil, err
	}

	finalData := make([]response.NftListingMarketplace, 0)
	if len(data.Data.Marketplace) > 0 {
		for _, marketplace := range data.Data.Marketplace {
			marketplace.ItemUrl = util.GetTokenMarketplaceUrl(collection.Address, collection.Symbol, tokenID, marketplace.PlatformName)
			decimals, _ := strconv.Atoi(marketplace.PaymentTokenDecimals)
			marketplace.ListingPrice = fmt.Sprintf("%.2f", util.StringWeiToEther(marketplace.ListingPrice, decimals))
			marketplace.FloorPrice = fmt.Sprintf("%.2f", util.StringWeiToEther(marketplace.FloorPrice, decimals))
			finalData = append(finalData, marketplace)
		}
		data.Data.Marketplace = finalData
	}

	// empty response
	if data == nil {
		e.log.Infof("[indexer.GetNFTDetail] no nft data from indexer")
		err := fmt.Errorf("no nft data from indexer")
		return nil, err
	}

	data.Data.Image = util.StandardizeUri(data.Data.Image)
	return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
		Data:        &data.Data,
		Suggestions: suggest,
	}, nil
}

func (e *Entity) GetNFTActivity(collectionAddress, tokenID, query string) (*response.GetNFTActivityResponse, error) {
	res, err := e.getNFTActivityFromIndexer(collectionAddress, tokenID, query)
	if err != nil {
		e.log.Errorf(err, "[e.getNFTActivityFromIndexer] failed to get nft indexer activity")
		return nil, err
	}

	// empty response
	if res == nil {
		e.log.Infof("[indexer.GetNFTActivity] no nft data from indexer")
		err := fmt.Errorf("no nft activity data from indexer")
		return nil, err
	}

	return &response.GetNFTActivityResponse{
		Data: response.GetNFTActivityData{
			Data: res.Data,
			Metadata: util.Pagination{
				Page:  res.Page,
				Size:  res.Size,
				Total: res.Total,
			},
		},
	}, nil
}

func (e *Entity) GetNFTTokenTransactionHistory(collectionAddress, tokenID string) (*response.IndexerGetNFTTokenTxHistoryResponse, error) {
	res, err := e.getNftTokenTransactionHistory(collectionAddress, tokenID)
	if err != nil {
		e.log.Errorf(err, "[e.GetNFTTokenTransactionHistory] failed to get nft indexer activity")
		return nil, err
	}

	return res, nil
}

func checkIsDefaultSymbol(defaults []model.GuildConfigDefaultCollection, symbol *model.NFTCollection) bool {
	for _, def := range defaults {
		if def.Address == symbol.Address && def.ChainID == symbol.ChainID {
			return true
		}
	}
	return false
}

type NFTCollectionData struct {
	TokenAddress string    `json:"token_address"`
	Name         string    `json:"name"`
	Symbol       string    `json:"symbol"`
	ContractType string    `json:"contract_type"`
	SyncedAt     time.Time `json:"synced_at"`
}

type MoralisMessageFail struct {
	Message string `json:"message"`
}

func (e *Entity) GetNFTSuggestion(symbol string, tokenID string) ([]response.CollectionSuggestions, error) {
	// get collections that are correct 50%
	matches, err := e.repo.NFTCollection.GetSuggestionsBySymbolorName(symbol, len(symbol)/2)
	if err != nil {
		if err.Error() == "record not found" {
			e.log.Info("[repo.NFTCollection.GetSuggestionsBySymbolorName] found no suggestions")
			return nil, fmt.Errorf("found no suggestions")
		} else {
			e.log.Errorf(err, "[repo.NFTCollection.GetSuggestionsBySymbolorName] failed to get nft suggestions for symbol %s", symbol)
			return nil, fmt.Errorf("[repo.NFTCollection.GetSuggestionsBySymbolorName] failed to get nft suggestions: %s", err)
		}
	}

	res := []response.CollectionSuggestions{}
	for _, col := range matches {
		chainId, _ := strconv.Atoi(col.ChainID)
		res = append(res, response.CollectionSuggestions{
			Name:    col.Name,
			Symbol:  col.Symbol,
			Address: col.Address,
			Chain:   util.ConvertChainIDToChain(col.ChainID),
			ChainId: int64(chainId),
		})
	}

	return res, nil
}

func (e *Entity) getTokenDetailFromIndexer(address string, tokenID string) (*response.IndexerGetNFTTokenDetailResponse, error) {
	data, err := e.indexer.GetNFTDetail(address, tokenID)
	// cannot find collection in indexer
	if err != nil {
		if err.Error() == "record not found" {
			e.log.Errorf(err, "[indexer.GetNFTDetail] indexer: record nft not found")
			err = baseerrs.ErrTokenNotFound
		} else {
			e.log.Errorf(err, "[indexer.GetNFTDetail] failed to get nft from indexer")
			err = fmt.Errorf("[e.GetNFTDetail] failed to get nft from indexer: %v", err)
		}
		return nil, err
	}
	return data, nil
}

func (e *Entity) getNFTActivityFromIndexer(collectionAddress, tokenID, query string) (*response.IndexerGetNFTActivityResponse, error) {
	data, err := e.indexer.GetNFTActivity(collectionAddress, tokenID, query)
	if err != nil {
		if err.Error() == "record not found" {
			e.log.Errorf(err, "[indexer.GetNFTActivity] indexer: record nft activity not found")
			err = fmt.Errorf("token not found")
		} else {
			e.log.Errorf(err, "[indexer.GetNFTActivity] failed to get nft activity from indexer")
			err = fmt.Errorf("[e.GetNFTActivity] failed to get nft activity from indexer: %v", err)
		}
		return nil, err
	}
	return data, nil
}

func (e *Entity) getNftTokenTransactionHistory(collectionAddress, tokenID string) (*response.IndexerGetNFTTokenTxHistoryResponse, error) {
	data, err := e.indexer.GetNFTTokenTxHistory(collectionAddress, tokenID)
	if err != nil {
		e.log.Fields(logger.Fields{"collectionAddress": collectionAddress, "tokenID": tokenID}).Errorf(err, "[indexer.GetNFTTokenTxHistory] failed to get nft token tx history from indexer")
		return nil, err
	}

	return data, nil
}

func (e *Entity) GetNFTDetailByAddress(address string, tokenID string) (*response.IndexerGetNFTTokenDetailResponseWithSuggestions, error) {
	exist, err := e.CheckExistNftCollection(address)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] failed to get nft collection by address %s", address)
		return nil, err
	}

	if !exist {
		return nil, fmt.Errorf("record not found")
	}

	collection, err := e.repo.NFTCollection.GetByAddress(address)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] failed to get nft collection")
		return nil, err
	}

	data, err := e.getTokenDetailFromIndexer(address, tokenID)
	if err != nil {
		e.log.Errorf(err, "[e.getTokenDetailFromIndexer] failed to get nft indexer detail")
		return nil, err
	}

	if data.Data.TokenID == "" {
		return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
			Data: nil,
		}, nil
	}

	finalData := make([]response.NftListingMarketplace, 0)
	if len(data.Data.Marketplace) > 0 {
		for _, marketplace := range data.Data.Marketplace {
			marketplace.ItemUrl = util.GetTokenMarketplaceUrl(collection.Address, collection.Symbol, tokenID, marketplace.PlatformName)
			decimals, _ := strconv.Atoi(marketplace.PaymentTokenDecimals)
			marketplace.ListingPrice = fmt.Sprintf("%.2f", util.StringWeiToEther(marketplace.ListingPrice, decimals))
			marketplace.FloorPrice = fmt.Sprintf("%.2f", util.StringWeiToEther(marketplace.FloorPrice, decimals))
			finalData = append(finalData, marketplace)
		}
		data.Data.Marketplace = finalData
	}

	data.Data.Image = util.StandardizeUri(data.Data.Image)

	return &response.IndexerGetNFTTokenDetailResponseWithSuggestions{
		Data: &data.Data,
	}, nil

}

func (e *Entity) CheckExistNftCollection(address string) (bool, error) {
	_, err := e.repo.NFTCollection.GetByAddress(address)
	// cannot find collection in db
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		} else {
			e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] failed to get nft collection by address %s", address)
			err = errors.New("failed to get nft collection")
			return false, err
		}
	}
	return true, nil
}

func (e *Entity) CheckIsSync(address string) (bool, error) {
	indexerContract, err := e.indexer.GetNFTContract(address)
	if err != nil {
		e.log.Errorf(err, "[indexer.GetNFTContract] failed to get nft contract by address %s", address)
		return false, err
	}

	return indexerContract.IsSynced, nil
}

func (e *Entity) CreateSolanaNFTCollection(req request.CreateNFTCollectionRequest) (nftCollection *model.NFTCollection, err error) {
	collectionAddress := "solscan-" + req.Address
	checkExistNFT, err := e.CheckExistNftCollection(collectionAddress)
	if err != nil {
		e.log.Errorf(err, "[e.CheckExistNftCollection] failed to check if nft exist: %v", err)
		return nil, err
	}

	if checkExistNFT {
		return e.handleExistCollection(req)
	}

	// get solana metadata collection from blockchain api
	solanaCollection, err := e.svc.Solscan.GetCollectionBySolscanId(req.Address)
	if err != nil {
		e.log.Errorf(err, "[e.svc.Solscan.GetSolanaCollection] failed to get solana collection: %v", err)
		return nil, err
	}

	solToken, err := e.svc.Solscan.GetNftTokenFromCollection(req.Address, "1")
	if err != nil {
		e.log.Errorf(err, "[e.svc.Solscan.GetSolanaCollection] failed to get solana collection: %v", err)
		return nil, err
	}
	collectionSymbol := ""
	if len(solToken.Data.ListNfts) > 0 {
		collectionSymbol = solToken.Data.ListNfts[0].NftSymbol
	}
	collectionName := ""
	if len(solanaCollection.Data.Collections) > 0 {
		collectionName = solanaCollection.Data.Collections[0].NftCollectionName
	}

	convertedChainId := util.ConvertChainToChainId(req.ChainID)
	chainID, _ := strconv.Atoi(convertedChainId)

	err = e.indexer.CreateERC721Contract(indexer.CreateERC721ContractRequest{
		Address: collectionAddress,
		ChainID: chainID,
		Name:    collectionName,
		// Symbol:       solanaCollection.Data.Data.Symbol,
		MessageID:    req.MessageID,
		PriorityFlag: req.PriorityFlag,
	})
	if err != nil {
		e.log.Errorf(err, "[CreateERC721Contract] failed to create erc721 contract: %v", err)
		return nil, fmt.Errorf("Failed to create erc721 contract: %v", err)
	}

	history := model.NftAddRequestHistory{
		Address:   collectionAddress,
		ChainID:   int64(chainID),
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		MessageID: req.MessageID,
	}
	err = e.repo.NftAddRequestHistory.UpsertOne(history)
	if err != nil {
		e.log.Errorf(err, "[CreateERC721Contract] repo.NftAddRequestHistory.UpsertOne() failed")
		return nil, fmt.Errorf("Failed to create erc721 contract: %v", err)
	}

	nftCollection, err = e.repo.NFTCollection.Create(model.NFTCollection{
		Address:    collectionAddress,
		Symbol:     collectionSymbol,
		Name:       collectionName,
		ChainID:    convertedChainId,
		ERCFormat:  "ERC721",
		IsVerified: true,
		Author:     req.Author,
		// Image:      solanaCollection.Data.Data.Avatar,
	})
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.Create] cannot add collection: %v", err)
		return nil, fmt.Errorf("Cannot add collection: %v", err)
	}

	return
}

func (e *Entity) CreateBluemoveNFTCollection(req request.CreateNFTCollectionRequest) (nftCollection *model.NFTCollection, err error) {
	checkExistNFT, err := e.CheckExistNftCollection(req.Address)
	if err != nil {
		e.log.Errorf(err, "[e.CheckExistNftCollection] failed to check if nft exist: %v", err)
		return nil, err
	}

	if checkExistNFT {
		return e.handleExistCollection(req)
	}

	convertedChainId := util.ConvertChainToChainId(req.ChainID)
	chainID, _ := strconv.Atoi(convertedChainId)
	collection, err := e.svc.Bluemove.SelectBluemoveCollection(req.Address, convertedChainId)
	if err != nil {
		e.log.Errorf(err, "[e.svc.Bluemove.SelectBluemoveCollection] failed to get bluemove collection: %v", err)
		return nil, err
	}

	history := model.NftAddRequestHistory{
		Address:   req.Address,
		ChainID:   int64(chainID),
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		MessageID: req.MessageID,
	}
	err = e.repo.NftAddRequestHistory.UpsertOne(history)
	if err != nil {
		e.log.Errorf(err, "[CreateERC721Contract] repo.NftAddRequestHistory.UpsertOne() failed")
		return nil, fmt.Errorf("Failed to create erc721 contract: %v", err)
	}

	err = e.indexer.CreateERC721Contract(indexer.CreateERC721ContractRequest{
		Address:      collection.Author,
		ChainID:      chainID,
		Name:         collection.Name,
		Symbol:       collection.Symbol,
		PriorityFlag: req.PriorityFlag,
	})
	if err != nil {
		e.log.Errorf(err, "[CreateERC721Contract] failed to create erc721 contract: %v", err)
		return nil, fmt.Errorf("Failed to create erc721 contract: %v", err)
	}

	nftCollection, err = e.repo.NFTCollection.Create(model.NFTCollection{
		Address:    collection.Address,
		Symbol:     collection.Symbol,
		Name:       collection.Name,
		ChainID:    convertedChainId,
		ERCFormat:  "ERC721",
		IsVerified: true,
		Author:     req.Author,
		Image:      collection.Image,
	})
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.Create] cannot add collection: %v", err)
		return nil, fmt.Errorf("Cannot add collection: %v", err)
	}
	return
}

func (e *Entity) CreateEVMNFTCollection(req request.CreateNFTCollectionRequest) (nftCollection *model.NFTCollection, err error) {
	address := e.HandleMarketplaceLink(req.Address, req.ChainID)
	if address == "collection does not have an address" {
		e.log.Infof("[e.HandleMarketplaceLink] collection %s does not have address", req.Address)
		return nil, fmt.Errorf("Collection does not have an address")
	}

	checksumAddress, err := util.ConvertToChecksumAddr(address)
	if err != nil {
		e.log.Errorf(err, "[util.ConvertToChecksumAddr] failed to convert checksum address: %v", err)
		return nil, fmt.Errorf("Failed to validate address: %v", err)
	}

	checkExistNFT, err := e.CheckExistNftCollection(checksumAddress)
	if err != nil {
		e.log.Errorf(err, "[e.CheckExistNftCollection] failed to check if nft exist: %v", err)
		return nil, err
	}

	req.Address = checksumAddress
	if checkExistNFT {
		return e.handleExistCollection(req)
	}

	convertedChainId := util.ConvertChainToChainId(req.ChainID)
	chainID, err := strconv.Atoi(convertedChainId)
	if err != nil {
		e.log.Errorf(err, "[util.ConvertChainToChainId] failed to convert chain to chainId: %v", err)
		return nil, fmt.Errorf("Failed to convert chain to chainId: %v", err)
	}
	colNolFound := false

	image, err := e.getImageFromMarketPlace(chainID, req.Address)
	if err != nil {
		e.log.Errorf(err, "[e.getImageFromMarketPlace] failed to get image from market place: %v", err)
		colNolFound = strings.Contains(err.Error(), "Collection not found") && strings.Contains(err.Error(), "paintswap")
		// if the error not related to collection not found should return
		if !colNolFound {
			return nil, err
		}
	}

	// query name and symbol from contract
	name, symbol, err := e.abi.GetNameAndSymbol(req.Address, int64(chainID))
	if err != nil {
		e.log.Errorf(err, "[GetNameAndSymbol] cannot get name and symbol of contract: %s | chainId %d", req.Address, chainID)
		if strings.Contains(err.Error(), "no contract code at given address") {
			return nil, baseerrs.ErrRecordNotFound
		}
		return nil, fmt.Errorf("Cannot get name and symbol of contract: %v", err)
	}
	// check if collection not found in paintswap then skip get from paintswap
	if !colNolFound {
		// host image to cloud if necessary
		image, err = e.svc.Cloud.HostImageToGCS(image, strings.ReplaceAll(name, " ", ""))
		if err != nil {
			e.log.Errorf(err, "[cloud.HostImageToGCS] failed to host image to GCS: %v", err)
			return nil, err
		}
	}

	// store nft add request
	history := model.NftAddRequestHistory{
		Address:   req.Address,
		ChainID:   int64(chainID),
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
		MessageID: req.MessageID,
	}
	err = e.repo.NftAddRequestHistory.UpsertOne(history)
	if err != nil {
		e.log.Errorf(err, "[CreateERC721Contract] repo.NftAddRequestHistory.UpsertOne() failed")
		return nil, fmt.Errorf("Failed to create erc721 contract: %v", err)
	}

	err = e.indexer.CreateERC721Contract(indexer.CreateERC721ContractRequest{
		Address:      req.Address,
		ChainID:      chainID,
		Name:         name,
		Symbol:       symbol,
		PriorityFlag: req.PriorityFlag,
	})
	if err != nil {
		e.log.Errorf(err, "[CreateERC721Contract] failed to create erc721 contract: %v", err)
		return nil, fmt.Errorf("Failed to create erc721 contract: %v", err)
	}

	nftCollection, err = e.repo.NFTCollection.Create(model.NFTCollection{
		Address:    req.Address,
		Symbol:     symbol,
		Name:       name,
		ChainID:    convertedChainId,
		ERCFormat:  "ERC721",
		IsVerified: true,
		Author:     req.Author,
		Image:      image,
	})
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.Create] cannot add collection: %v", err)
		return nil, fmt.Errorf("Cannot add collection: %v", err)
	}

	//Add collection to podtown
	go CreatePodtownNFTCollection(model.NFTCollection{
		Address:   req.Address,
		Symbol:    symbol,
		Name:      name,
		ChainID:   convertedChainId,
		ERCFormat: "ERC721",
	}, e.cfg)

	return
}

func (e *Entity) handleExistCollection(req request.CreateNFTCollectionRequest) (*model.NFTCollection, error) {
	isSync, err := e.CheckIsSync(req.Address)
	if err != nil {
		e.log.Errorf(err, "[e.CheckIsSync] failed to check if nft is synced: %v", err)
		return nil, err
	}

	if !isSync {
		e.log.Infof("[e.CheckIsSync] Already added. Nft is in sync progress")
		return nil, fmt.Errorf("Already added. Nft is in sync progress")
	} else {
		e.log.Infof("[e.CheckIsSync] Already added. Nft is done with sync")
		return nil, fmt.Errorf("Already added. Nft is done with sync")
	}
}

func (e *Entity) getImageFromMarketPlace(chainID int, address string) (string, error) {
	if chainID == 1 {
		// TODO(trkhoi): opensea api key is expired, get new one
		// collection, err := e.marketplace.GetOpenseaAssetContract(address)
		// if err != nil {
		// 	e.log.Errorf(err, "[GetOpenseaAssetContract] cannot get contract: %s | chainId %d", address, chainID)
		// 	return "", fmt.Errorf("Cannot get contract: %v", err)
		// }
		return "", nil
	}
	if chainID == 250 {
		collection, err := e.marketplace.GetCollectionFromPaintswap(address)
		if err != nil {
			e.log.Errorf(err, "[GetCollectionFromPaintswap] cannot get collection: %s | chainId %d", address, chainID)
			return "", fmt.Errorf("Cannot get collection from paintswap: %v", err)
		}
		return collection.Collection.Image, nil
	}
	if chainID == 10 {
		// TODO: get image, alchemy response does not include image
		// collection, err := e.marketplace.GetCollectionFromQuixotic(address)
		// if err != nil {
		// 	e.log.Errorf(err, "[GetCollectionFromQuixotic] cannot get collection: %s | chainId %d", address, chainID)
		// 	return "", fmt.Errorf("Cannot get collection: %v", err)
		// }
		// return collection.ImageUrl, nil
		return "", nil
	}

	return "", nil
}

func CreatePodtownNFTCollection(collection model.NFTCollection, cfg config.Config) (err error) {
	body, err := json.Marshal(collection)
	if err != nil {
		return err
	}
	jsonBody := bytes.NewBuffer(body)
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/nft/collection", cfg.PodtownServerHost), jsonBody)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return nil
}

func (e *Entity) ListAllNFTCollections() ([]model.NFTCollection, error) {
	return e.repo.NFTCollection.ListAll()
}

func (e *Entity) ListAllNFTCollectionConfigs() ([]model.NFTCollectionConfig, error) {
	return e.repo.NFTCollection.ListAllNFTCollectionConfigs()
}

func (e *Entity) GetNFTBalanceFunc(config model.NFTCollectionConfig) (func(address string) (*big.Int, error), error) {
	chainID, err := strconv.Atoi(config.ChainID)
	if err != nil {
		e.log.Errorf(err, "[strconv.Atoi] failed to convert chain id %s to int", config.ChainID)
		return nil, fmt.Errorf("failed to convert chain id %s to int: %v", config.ChainID, err)
	}

	chain, err := e.repo.Chain.GetByID(chainID)
	if err != nil {
		e.log.Errorf(err, "[repo.Chain.GetByID] failed to get chain by id %s", config.ChainID)
		return nil, fmt.Errorf("failed to get chain by id %s: %v", config.ChainID, err)
	}

	client, err := ethclient.Dial(chain.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to chain client: %v", err.Error())
	}

	var balanceOf func(string) (*big.Int, error)
	switch strings.ToLower(config.ERCFormat) {
	case "721", "erc721":
		contract721, err := erc721.NewErc721(common.HexToAddress(config.Address), client)
		if err != nil {
			e.log.Errorf(err, "[erc721.NewErc721] failed to init erc721 contract")
			return nil, fmt.Errorf("failed to init erc721 contract: %v", err.Error())
		}

		balanceOf = func(address string) (*big.Int, error) {
			b, err := contract721.BalanceOf(nil, common.HexToAddress(address))
			if err != nil {
				e.log.Errorf(err, "[contract721.BalanceOf] failed to get balance of %s in chain %s", address, config.ChainID)
				return nil, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, config.ChainID, err.Error())
			}
			return b, nil
		}

	case "1155", "erc1155":
		contract1155, err := typeset.NewErc1155(common.HexToAddress(config.Address), client)
		if err != nil {
			e.log.Errorf(err, "[erc1155.NewErc1155] failed to init erc1155 contract")
			return nil, fmt.Errorf("failed to init erc1155 contract: %v", err.Error())
		}

		balanceOf = func(address string) (*big.Int, error) {
			total := big.NewInt(int64(0))
			for index := 0; index <= consts.MaxTokenId; index++ {
				tokenID, err := strconv.ParseInt(strconv.Itoa(index), 10, 64)
				if err != nil {
					e.log.Errorf(err, "[strconv.ParseInt] token id is not valid")
					return nil, fmt.Errorf("token id is not valid")
				}
				b, err := contract1155.BalanceOf(nil, common.HexToAddress(address), big.NewInt(tokenID))
				if err != nil {
					e.log.Errorf(err, "[contract1155.BalanceOf] failed to get balance of %s in chain %s", address, config.ChainID)
					return nil, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, config.ChainID, err.Error())
				}
				total = total.Add(total, b)

			}
			return total, nil
		}

	default:
		e.log.Errorf(err, "[GetNFTBalanceFunc] erc format %s not supported", config.ERCFormat)
		return nil, fmt.Errorf("erc format %s not supported", config.ERCFormat)
	}

	return balanceOf, nil
}

func (e *Entity) NewUserNFTBalance(balance model.UserNFTBalance) error {
	err := e.repo.UserNFTBalance.Upsert(balance)
	if err != nil {
		e.log.Errorf(err, "[repo.UserNFTBalance.Upsert] failed to upsert user nft balance")
		return fmt.Errorf("failed to upsert user nft balance: %v", err.Error())
	}
	return nil
}

func (e *Entity) GetNFTCollectionTickers(req request.GetNFTCollectionTickersRequest, rawQuery string) (*response.IndexerNFTCollectionTickersResponse, error) {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Infof("Nft colletion ticker not found. CollectionAddress: %s", req.CollectionAddress)
			return &response.IndexerNFTCollectionTickersResponse{Data: nil}, nil
		}
		e.log.Errorf(err, "[repo.NFTCollection.GetBySymbol] failed to get nft collection by.  CollectionAddress: %s", req.CollectionAddress)
		return nil, err
	}

	res, err := e.indexer.GetNFTCollectionTickers(collection.Address, rawQuery)
	if err != nil {
		if err.Error() == "record not found" {
			e.log.Infof("[indexer.GetNFTCollectionTickers] Indexer does not have ticker for this collection. CollectionAddress: %s", req.CollectionAddress)
			return nil, err
		}
		e.log.Errorf(err, "[indexer.GetNFTCollectionTickers] failed to get nft collection tickers by %s and %s", collection.Address, rawQuery)
		return nil, err
	}

	for _, ts := range res.Data.Tickers.Timestamps {
		time := time.UnixMilli(ts)
		res.Data.Tickers.Times = append(res.Data.Tickers.Times, time.Format("01-02"))
	}
	return res, nil
}

func (e *Entity) GetNFTCollections(p string, s string) (*response.NFTCollectionsResponse, error) {
	page, _ := strconv.Atoi(p)
	size, _ := strconv.Atoi(s)
	data, total, err := e.repo.NFTCollection.ListAllWithPaging(page, size)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.ListAllWithPaging] failed to list all nft collection with paging")
		return nil, err
	}

	for i := range data {
		data[i].Image = util.StandardizeUri(data[i].Image)
	}

	return &response.NFTCollectionsResponse{
		Data: response.NFTCollectionsData{
			Metadata: util.Pagination{
				Page:  int64(page),
				Size:  int64(size),
				Total: total,
			},
			Data: data,
		},
	}, err
}

func (e *Entity) GetNFTTokens(symbol, query string) (*response.IndexerGetNFTTokensResponse, error) {
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.GetBySymbol] failed to get nft collection by symbol %s", symbol)
		return nil, err
	}
	if collection.Address == "" {
		e.log.Errorf(err, "[GetNFTTokens] invalid address - collection %s", collection.ID.UUID)
		return nil, fmt.Errorf("invalid address - collection %s", collection.ID.UUID)
	}
	data, err := e.svc.Indexer.GetNFTTokens(collection.Address, query)
	if err != nil {
		e.log.Errorf(err, "[svc.Indexer.GetNFTTokens] failed to get nft tokens by %s and  %s", collection.Address, query)
		return nil, err
	}
	return data, nil
}

func (e *Entity) GetDetailNftCollection(symbol string, queryByAddress bool) (*model.NFTCollectionDetail, error) {
	if queryByAddress {
		data, err := e.repo.NFTCollection.GetByAddress(symbol)
		if err != nil {
			e.log.Errorf(err, "[repo.NFTCollection.GetByAddress] failed to get nft collection by address")
			return nil, err
		}
		return e.GetCollectionWithChainDetail(data), nil
	}

	collection, err := e.repo.NFTCollection.GetBySymbolorName(symbol)
	if err != nil || len(collection) == 0 {
		e.log.Errorf(err, "[repo.NFTCollection.GetBySymbolorName] failed to get nft collection by %s", symbol)
		return nil, err
	}

	return e.GetCollectionWithChainDetail(&collection[0]), nil
}

func (e *Entity) GetCollectionWithChainDetail(collection *model.NFTCollection) *model.NFTCollectionDetail {
	chainId, _ := strconv.Atoi(collection.ChainID)
	var chainPtr *model.Chain

	// if chain not exist return chain=nil
	chain, err := e.repo.Chain.GetByID(chainId)
	chainPtr = &chain
	if err != nil {
		e.log.Infof("[e.repo.Chain.GetByID] failed to get chain: %s", collection.ChainID)
		chainPtr = nil
	}
	return &model.NFTCollectionDetail{
		ID:         collection.ID,
		Address:    collection.Address,
		Name:       collection.Name,
		Symbol:     collection.Symbol,
		ChainID:    collection.ChainID,
		Chain:      chainPtr,
		ERCFormat:  collection.ERCFormat,
		IsVerified: collection.IsVerified,
		CreatedAt:  collection.CreatedAt,
		Image:      util.StandardizeUri(collection.Image),
		Author:     collection.Author,
	}
}

func (e *Entity) GetNewListedNFTCollection(interval string, page string, size string) (*response.NFTNewListed, error) {
	itv, _ := strconv.Atoi(interval)
	pg, _ := strconv.Atoi(page)
	lim, _ := strconv.Atoi(size)
	data, total, err := e.repo.NFTCollection.GetNewListed(itv, pg, lim)
	for i, ele := range data {
		chainId, _ := strconv.Atoi(ele.ChainID)
		chain, err := e.repo.Chain.GetByID(chainId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				e.log.Infof("Not have chain with id: ", chainId)
				data[i].Chain = ""
				continue
			}
			e.log.Errorf(err, "[repo.Chain.GetByID] failed to get chain %d", chainId)
			return nil, err
		}
		data[i].Chain = chain.Name
	}

	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			e.log.Info("[handler.GetNewListedNFTCollection] - no new collection")
			return &response.NFTNewListed{
				Data: nil,
			}, nil
		}
		e.log.Errorf(err, "[repo.NFTCollection.GetNewListed] failed to get recent collections")
		return nil, err
	}

	return &response.NFTNewListed{
		Metadata: util.Pagination{
			Page:  int64(pg),
			Size:  int64(lim),
			Total: total,
		},
		Data: data,
	}, nil
}

func (e *Entity) GetNftMetadataAttrIcon() (*response.NftMetadataAttrIconResponse, error) {
	data, err := e.indexer.GetNftMetadataAttrIcon()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *Entity) GetCollectionCount() (*response.NFTCollectionCount, error) {
	chains, err := e.repo.Chain.GetAll()
	if err != nil {
		e.log.Errorf(err, "[e.GetCollectionCount] - cannot get all chain")
		return nil, err
	}
	total := 0
	data := []response.NFTChainCollectionCount{}
	for _, chain := range chains {
		_, nr, err := e.repo.NFTCollection.GetByChain(chain.ID)
		if err != nil {
			e.log.Errorf(err, fmt.Sprintf("[e.GetCollectionCount] - cannot count number of %s collections", chain.Currency))
			return nil, err
		}
		total += nr
		data = append(data, response.NFTChainCollectionCount{Chain: chain, Count: nr})
	}
	return &response.NFTCollectionCount{
		Total: total,
		Data:  data,
	}, nil
}

func (e *Entity) GetNFTCollectionByAddressChain(address, chainId string) (*response.GetNFTCollectionByAddressChain, error) {
	collection, err := e.repo.NFTCollection.GetByAddressChainId(address, chainId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Infof("Collection is not exist, address: %s", address)
			return nil, nil
		}
		e.log.Errorf(err, "[repo.NFTCollection.GetNFTCollectionByAddress] failed to get nft collection by address %s", address)
		return nil, err
	}

	collectionMetadata, err := e.indexer.GetNftCollectionMetadata(address, chainId)
	if err != nil {
		e.log.Fields(logger.Fields{"address": address, "chainId": chainId}).Error(err, "[repo.NFTCollection.GetNFTCollectionByAddress] failed to get nft collection metadata")
		return nil, err
	}

	marketplaces := []string{}
	if len(collectionMetadata.Data.Marketplace) > 0 {
		for _, itm := range collectionMetadata.Data.Marketplace {
			marketplaces = append(marketplaces, itm.PlatformName)
		}
	}

	return &response.GetNFTCollectionByAddressChain{
		ID:           collection.ID,
		Address:      collection.Address,
		Name:         collection.Name,
		Symbol:       collection.Symbol,
		ChainID:      collection.ChainID,
		ERCFormat:    collection.ERCFormat,
		IsVerified:   collection.IsVerified,
		CreatedAt:    collection.CreatedAt,
		Image:        util.StandardizeUri(collection.Image),
		Author:       collection.Author,
		Description:  collectionMetadata.Data.Description,
		Discord:      collectionMetadata.Data.Discord,
		Twitter:      collectionMetadata.Data.Twitter,
		Website:      collectionMetadata.Data.Website,
		Marketplaces: marketplaces,
	}, nil
}

func (e *Entity) UpdateNFTCollection(address string) error {
	collection, err := e.repo.NFTCollection.GetByAddress(address)
	if err != nil {
		e.log.Errorf(err, "[e.UpdateNFTCollection] cannot get address")
		return err
	}
	// if image already valid, function return same string
	image, err := e.svc.Cloud.HostImageToGCS(collection.Image, strings.ReplaceAll(collection.Name, " ", ""))
	if err != nil {
		e.log.Errorf(err, "[e.UpdateNFTCollection] cannot host image")
		return err
	}
	if image != collection.Image {
		err := e.repo.NFTCollection.UpdateImage(address, image)
		if err != nil {
			e.log.Errorf(err, "[e.UpdateNFTCollection] cannot update image")
			return err
		}
	}
	return nil
}

func (e *Entity) AddNftWatchlist(req request.AddNftWatchlistRequest) (*response.NftWatchlistSuggestResponse, error) {
	if req.CollectionAddress == "" && req.Chain == "" {
		suggestNftCollection, collection, err := e.SuggestCollection(req)
		if err != nil {
			e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[e.SuggestCollection] failed to get nft collection suggestion")
			return nil, err
		}

		// case symbol has suggestion
		if suggestNftCollection != nil && collection == nil {
			return &response.NftWatchlistSuggestResponse{Data: suggestNftCollection}, nil
		}

		if collection != nil && suggestNftCollection == nil {
			req.CollectionAddress = collection.Address
			req.Chain = util.ConvertChainIDToChain(collection.ChainID)
		}
	}

	// case has collection address + chain id / not have but suggest return 1 result
	collection, err := e.repo.NFTCollection.GetByAddressChainId(req.CollectionAddress, util.ConvertChainToChainId(req.Chain))
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[e.AddNftWatchlist] failed to get nft collection by address and chain id")
		return nil, err
	}
	chainID, _ := strconv.Atoi(collection.ChainID)
	listQ := usernftwatchlistitem.UserNftWatchlistQuery{CollectionAddress: req.CollectionAddress, ProfileID: req.ProfileID, ChainID: collection.ChainID, Symbol: collection.Symbol}
	_, total, err := e.repo.UserNftWatchlistItem.List(listQ)
	if err != nil {
		e.log.Fields(logger.Fields{"listQ": listQ}).Error(err, "[entity.AddNftToWatchlist] repo.UserWatchlistItem.List() failed")
		return nil, err
	}
	if total == 1 {
		return nil, baseerrs.ErrConflict
	}

	err = e.repo.UserNftWatchlistItem.Create(&model.UserNftWatchlistItem{
		ProfileID:         req.ProfileID,
		Symbol:            collection.Symbol,
		CollectionAddress: collection.Address,
		ChainId:           int64(chainID),
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[UserNftWatchlistItem.Create] failed to create user nft wl")
		return nil, err
	}

	return &response.NftWatchlistSuggestResponse{Data: nil}, nil
}

func (e *Entity) SuggestCollection(req request.AddNftWatchlistRequest) (*response.NftWatchlistSuggest, *model.NFTCollection, error) {
	// get collection
	e.log.Infof("DEBUG PROD: [handler.GetSuggestionNFTCollections] Checking req: ", req)

	suggest := []response.CollectionSuggestions{}
	collections, err := e.repo.NFTCollection.GetBySymbolorName(req.CollectionSymbol)

	e.log.Info("DEBUG PROD: [entities.SuggestCollection] GetBySymbolorName()")
	for i, collection := range collections {
		e.log.Info("DEBUG PROD: [entities.SuggestCollection] Result of query database")
		e.log.Infof("DEBUG PROD: [entities.SuggestCollection] ", i, " - ", collection.Address, " - ", collection.Symbol, " - ", collection.Name)
	}

	// cannot find collection => return suggested collections
	if err != nil || len(collections) == 0 {
		e.log.Info("DEBUG PROD: [entities.SuggestCollection] In case cant find collection")
		suggest, err = e.GetNFTSuggestion(req.CollectionSymbol, "")
		if err != nil {
			e.log.Errorf(err, "[repo.NFTCollection.GetBySymbolorName] failed to get nft collection by symbol %s", req.CollectionSymbol)
			return nil, nil, err
		}
		return &response.NftWatchlistSuggest{
			Suggestions: suggest,
		}, nil, nil
	}

	// found multiple symbols => only suggest those
	if len(collections) > 1 {
		e.log.Info("DEBUG PROD: [entities.SuggestCollection] In case find multiple collections")
		var defaultSymbol *response.CollectionSuggestions
		// check default symbol
		symbols, _ := e.GetDefaultCollectionSymbol(req.GuildID)
		for _, col := range collections {
			// if found default symbol
			chainId, _ := strconv.Atoi(col.ChainID)
			if len(symbols) != 0 && checkIsDefaultSymbol(symbols, &col) {
				def := response.CollectionSuggestions{
					Address: col.Address,
					Chain:   util.ConvertChainIDToChain(col.ChainID),
					Name:    col.Name,
					Symbol:  col.Symbol,
					ChainId: int64(chainId),
				}
				defaultSymbol = &def
			}
			suggest = append(suggest, response.CollectionSuggestions{
				Name:    col.Name,
				Symbol:  col.Symbol,
				Address: col.Address,
				Chain:   util.ConvertChainIDToChain(col.ChainID),
				ChainId: int64(chainId),
			})
		}
		return &response.NftWatchlistSuggest{
			Suggestions:   suggest,
			DefaultSymbol: defaultSymbol,
		}, nil, nil
	}

	e.log.Info("DEBUG PROD: [entities.SuggestCollection] In case find 1 collection")
	return nil, &collections[0], nil
}

func (e *Entity) DeleteNftWatchlist(req request.DeleteNftWatchlistRequest) error {
	rows, err := e.repo.UserNftWatchlistItem.Delete(req.ProfileID, req.Symbol)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.DeleteNftWatchlist] repo.UserNftWatchlistItem.Delete() failed")
	}
	if rows == 0 {
		e.log.Fields(logger.Fields{"req": req}).Error(err, "[entity.DeleteNftWatchlist] item not found")
		return baseerrs.ErrRecordNotFound
	}
	return err
}

func (e *Entity) GetNftWatchlist(req *request.ListTrackingNftsRequest) (*response.GetNftWatchlistResponse, error) {
	q := usernftwatchlistitem.UserNftWatchlistQuery{
		ProfileID: req.ProfileID,
		Offset:    req.Page * req.Size,
		Limit:     req.Size,
	}
	list, _, err := e.repo.UserNftWatchlistItem.List(q)
	if err != nil {
		e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.GetNftWatchlist] repo.UserWatchlistItem.List() failed")
		return nil, err
	}

	if len(list) == 0 {
		list = e.getDefaultNFTWatchlistItems()
	}

	res := make([]response.GetNftWatchlist, 0)
	for _, itm := range list {
		data, err := e.indexer.GetNFTCollectionTickersForWl(itm.CollectionAddress)
		if err != nil {
			e.log.Fields(logger.Fields{"query": q}).Error(err, "[entity.GetNftWatchlist] indexer.GetNFTCollectionTickersForWl failed")
			return nil, err
		}

		collection, err := e.repo.NFTCollection.GetByAddressChainId(itm.CollectionAddress, strconv.Itoa(int(itm.ChainId)))
		if err != nil {
			e.log.Fields(logger.Fields{"CollectionAddress": itm.CollectionAddress, "ChainId": itm.ChainId}).Error(err, "[entity.GetNftWatchlist] repo.NFTCollection.GetByAddressChainId failed")
			return nil, err
		}

		if data == nil {
			res = append(res, response.GetNftWatchlist{
				Symbol:                   itm.Symbol,
				Image:                    collection.Image,
				IsPair:                   false,
				Name:                     collection.Name,
				PriceChangePercentage24h: 0,
				SparkLineIn7d: response.SparkLineIn7d{
					Price: []float64{},
				},
				FloorPrice:                        0,
				PriceChangePercentage7dInCurrency: 0,
			})
			continue
		}

		price := make([]float64, 0)
		for _, ticker := range data.Data.Tickers.Prices {
			bigFloatFloorPrice7d := util.StringWeiToEther(ticker.Amount, int(ticker.Token.Decimals))
			floatFloorPrice7d, _ := bigFloatFloorPrice7d.Float64()
			price = append(price, floatFloorPrice7d)
		}
		// NFT collection does not have floor price
		amount := "0"
		decimals := 0
		token := response.IndexerToken{}
		if data.Data.FloorPrice != nil {
			amount = data.Data.FloorPrice.Amount
			decimals = int(data.Data.FloorPrice.Token.Decimals)
			token = data.Data.FloorPrice.Token
		}
		floatFloorPrice, _ := util.StringWeiToEther(amount, decimals).Float64()

		priceChangePercentage7dInCurrency := 0.0
		if len(price) > 0 {
			priceChangePercentage7dInCurrency = (price[len(price)-1] - price[0]) / price[0]
		}

		itmRes := response.GetNftWatchlist{
			Symbol:                   itm.Symbol,
			Image:                    collection.Image,
			IsPair:                   false,
			Name:                     collection.Name,
			PriceChangePercentage24h: 0,
			SparkLineIn7d: response.SparkLineIn7d{
				Price: price,
			},
			FloorPrice:                        floatFloorPrice,
			PriceChangePercentage7dInCurrency: priceChangePercentage7dInCurrency,
			Token:                             token,
		}
		res = append(res, itmRes)
	}
	return &response.GetNftWatchlistResponse{Data: res}, nil
}

func (e *Entity) getDefaultNFTWatchlistItems() []model.UserNftWatchlistItem {
	return []model.UserNftWatchlistItem{
		{CollectionAddress: "0x7acee5d0acc520fab33b3ea25d4feef1ffebde73", Symbol: "NEKO", ChainId: 250},
		{CollectionAddress: "0x2ab5c606a5aa2352f8072b9e2e8a213033e2c4c9", Symbol: "MGC", ChainId: 250},
	}
}

func (e *Entity) GetGuildDefaultNftTicker(req request.GetGuildDefaultNftTickerRequest) (*model.GuildConfigDefaultCollection, error) {
	defaultNftTicker, err := e.repo.GuildConfigDefaultCollection.GetOneByGuildIDAndQuery(req.GuildID, req.Query)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"guild_id": req.GuildID, "query": req.Query}).Error(err, "[entity.GetGuildDefaultNftTicker] repo.GuildConfigDefaultNftTicker.GetOneByGuildIDAndQuery() failed")
		return nil, err
	}
	return defaultNftTicker, nil
}

func (e *Entity) SetGuildDefaultNftTicker(req request.GuildConfigDefaultNftTickerRequest) error {
	return e.repo.GuildConfigDefaultCollection.Upsert(&model.GuildConfigDefaultCollection{
		GuildID: req.GuildID,
		Address: req.CollectionAddress,
		Symbol:  req.Symbol,
		ChainID: fmt.Sprint(req.ChainId),
	})
}

func (e *Entity) GetSuggestionNftCollections(query string) ([]response.CollectionSuggestions, error) {
	// case query (symbol) has 1 collection -> return in collection
	// has more than 1 collection -> return collectionSuggestions
	collectionSuggestions, collection, error := e.SuggestCollection(request.AddNftWatchlistRequest{
		CollectionSymbol: strings.ToUpper(query),
	})

	if error != nil {
		e.log.Fields(logger.Fields{"query": query}).Error(error, "[entity.GetSuggestionNftCollections] GetNFTSuggestion() failed")
		return nil, error
	}

	if collectionSuggestions == nil {
		chainId, _ := strconv.Atoi(collection.ChainID)
		return []response.CollectionSuggestions{{
			Name:    collection.Name,
			Symbol:  collection.Symbol,
			Address: collection.Address,
			Chain:   util.ConvertChainToChainId(collection.ChainID),
			ChainId: int64(chainId),
		}}, nil
	}

	return collectionSuggestions.Suggestions, nil
}

func (e *Entity) GetNFTTokenTickers(req request.GetNFTTokenTickersRequest, rawQuery string) (*response.IndexerNFTTokenTickersData, error) {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			e.log.Infof("Nft colletion not found. CollectionAddress: %s", req.CollectionAddress)
			return nil, nil
		}
		e.log.Errorf(err, "[repo.NFTCollection.GetBySymbol] failed to get nft collection by.  CollectionAddress: %s", req.CollectionAddress)
		return nil, err
	}

	res, err := e.indexer.GetNFTTokenTickers(req.CollectionAddress, req.TokenID, rawQuery)
	if err != nil {
		if err.Error() == "record not found" {
			e.log.Infof("[indexer.GetNFTTokenTickers] Indexer does not have ticker for this token. CollectionAddress: %s, tokenID: %s", req.CollectionAddress, req.TokenID)
			return nil, err
		}
		e.log.Errorf(err, "[indexer.GetNFTTokenTickers] failed to get nft token tickers. Address: %s, TokenID: %s, query: %s", collection.Address, req.TokenID, rawQuery)
		return nil, err
	}

	for _, ts := range res.Tickers.Timestamps {
		time := time.UnixMilli(ts)
		res.Tickers.Times = append(res.Tickers.Times, time.Format("01-02"))
	}
	return res, nil
}

func (e *Entity) TotalNftCollection() (int64, error) {
	return e.repo.NFTCollection.TotalNftCollection()
}

func (e *Entity) GetProfileNftBalance(req request.GetProfileNFTsRequest) (*response.ProfileNftBalancesData, error) {
	collection, err := e.repo.NFTCollection.GetByAddress(req.CollectionAddress)
	if err != nil {
		e.log.Error(err, "[entity.GetProfileNftBalance] repo.NFTCollection.GetByAddress() failed")
		return nil, err
	}

	collectionID := collection.ID.UUID.String()
	total, err := e.repo.UserNFTBalance.TotalBalance(collectionID)
	if err != nil {
		e.log.Error(err, "[entity.GetProfileNftBalance] repo.UserNFTBalance.TotalBalance() failed")
		return nil, err
	}

	bals, err := e.repo.UserNFTBalance.List(user_nft_balance.ListQuery{ProfileID: req.ProfileID, CollectionID: collectionID})
	if err != nil {
		e.log.Error(err, "[entity.GetProfileNftBalance] repo.UserNFTBalance.List() failed")
		return nil, err
	}

	data := response.ProfileNftBalancesData{
		ProfileID: req.ProfileID,
		Total:     total,
		Items:     bals,
	}
	if !sliceutils.IsEmpty(bals) {
		data.Balance = sliceutils.Reduce(bals, func(acc int, curr model.UserNFTBalance) int {
			return acc + curr.Balance + curr.StakingNekos
		}, 0)
		if total != 0 {
			data.Share = float64(data.Balance) / float64(total)
		}

	}

	// go e.fetchNftBalance(req.ProfileID, collection)

	return &data, nil
}

func (e *Entity) fetchNftBalance(profileID string, collection *model.NFTCollection) {
	if profileID == "" {
		return
	}
	if !strings.EqualFold(collection.ERCFormat, model.ErcFormat721) || !collection.IsVerified {
		return
	}

	// only support EVM
	// TODO: why collection.chainID is string?
	chainID, err := strconv.Atoi(collection.ChainID)
	if err != nil {
		return
	}
	chain, err := e.repo.Chain.GetByID(chainID)
	if err != nil {
		e.log.Errorf(err, "[entity.fetchNftBalance] GetNFTBalanceFunc() failed")
		return
	}
	if !strings.EqualFold(chain.Type, model.ChainTypeEvm.String()) {
		return
	}

	balanceOf, err := e.GetNFTBalanceFunc(model.NFTCollectionConfig{
		ID:        collection.ID,
		ERCFormat: collection.ERCFormat,
		Address:   collection.Address,
		Name:      collection.Name,
		ChainID:   collection.ChainID,
	})
	if err != nil {
		e.log.Errorf(err, "[entity.fetchNftBalance] GetNFTBalanceFunc() failed")
		return
	}

	profile, err := e.svc.MochiProfile.GetByID(profileID)
	if err != nil {
		e.log.Errorf(err, "[entity.fetchNftBalance] svc.MochiProfile.GetByID() failed")
		return
	}

	if profile == nil || profile.AssociatedAccounts == nil {
		return
	}

	evmAccs := sliceutils.Filter(profile.AssociatedAccounts, func(aa mochiprofile.AssociatedAccount) bool {
		return aa.Platform == mochiprofile.PlatformEVM
	})

	now := time.Now().UTC()
	for _, acc := range evmAccs {
		// fetch nft balance
		bal, err := balanceOf(acc.PlatformIdentifier)
		if err != nil {
			e.log.Errorf(err, "[entity.fetchNftBalance] balanceOf() failed")
			continue
		}

		// update db record
		err = e.repo.UserNFTBalance.Upsert(model.UserNFTBalance{
			NFTCollectionID: collection.ID,
			UserAddress:     strings.ToLower(acc.PlatformIdentifier),
			ChainType:       model.JSONNullString{NullString: sql.NullString{String: chain.Type, Valid: true}},
			Balance:         int(bal.Int64()),
			ProfileID:       profileID,
			UpdatedAt:       now,
		})
		if err != nil {
			e.log.Errorf(err, "[entity.fetchNftBalance] repo.UserNFTBalance.Upsert() failed")
			continue
		}
	}
}

func (e *Entity) GetNekoBalanceFunc(config model.NFTCollectionConfig) (func(address string) (*big.Int, *big.Int, error), error) {
	chainID, err := strconv.Atoi(config.ChainID)
	if err != nil {
		e.log.Errorf(err, "[strconv.Atoi] failed to convert chain id %s to int", config.ChainID)
		return nil, fmt.Errorf("failed to convert chain id %s to int: %v", config.ChainID, err)
	}

	chain, err := e.repo.Chain.GetByID(chainID)
	if err != nil {
		e.log.Errorf(err, "[repo.Chain.GetByID] failed to get chain by id %s", config.ChainID)
		return nil, fmt.Errorf("failed to get chain by id %s: %v", config.ChainID, err)
	}

	client, err := ethclient.Dial(chain.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to chain client: %v", err.Error())
	}

	var balanceOf func(string) (*big.Int, *big.Int, error)
	contract721, err := erc721.NewErc721(common.HexToAddress(config.Address), client)
	if err != nil {
		e.log.Errorf(err, "[erc721.NewErc721] failed to init erc721 contract")
		return nil, err
	}

	nekoStakingAddr := "0xD28Cf82b9B8ee25E3C82923aDF6aA6CC2f220932"
	stakingInstance, err := nekostaking.NewNekostaking(common.HexToAddress(nekoStakingAddr), client)
	if err != nil {
		e.log.Errorf(err, "[erc721.NewErc721] nekostaking.NewNekoStaking() failed")
		return nil, err
	}

	balanceOf = func(address string) (holding *big.Int, staking *big.Int, err error) {
		exist, err := e.repo.UserNFTBalance.IsExists(config.ID.UUID.String(), address)
		if exist {
			return nil, nil, errors.New("existed")
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func(addr string) {
			defer wg.Done()
			holding, err = contract721.BalanceOf(nil, common.HexToAddress(address))
			if err != nil {
				e.log.Errorf(err, "[contract721.BalanceOf] failed to get balance of %s in chain %s", address, config.ChainID)
				// return nil, nil, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, config.ChainID, err.Error())
			}
		}(address)

		go func(addr string) {
			defer wg.Done()
			staking, err = stakingInstance.Balances(nil, common.HexToAddress(address))
			if err != nil {
				e.log.Errorf(err, "[contract721.BalanceOf] failed to get balance of %s in chain %s", address, config.ChainID)
				// return nil, nil, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, config.ChainID, err.Error())
			}
		}(address)

		wg.Wait()

		return holding, staking, nil
	}

	return balanceOf, nil
}

// TODO: remove after airdrop
func (e *Entity) FetchHolders(col model.NFTCollection) ([]model.UserNFTBalance, error) {
	chainID, err := strconv.Atoi(col.ChainID)
	if err != nil {
		e.log.Errorf(err, "[FetchHolders] strconv.Atoi() failed")
		return nil, err
	}

	chain, err := e.repo.Chain.GetByID(chainID)
	if err != nil {
		e.log.Errorf(err, "[FetchHolders] repo.Chain.GetByID() failed")
		return nil, err
	}

	client, err := ethclient.Dial(chain.RPC)
	if err != nil {
		e.log.Errorf(err, "[FetchHolders] ethclient.Dial() failed")
		return nil, err
	}

	tokenTracker, err := erc721.NewErc721(common.HexToAddress(col.Address), client)
	if err != nil {
		e.log.Errorf(err, "[FetchHolders] erc721.NewErc721() failed")
		return nil, err
	}

	// tokenTracker, err := neko.NewNeko(common.HexToAddress(col.Address), client)
	// if err != nil {
	// 	e.log.Errorf(err, "[FetchHolders] neko.NewNeko() failed")
	// 	return nil, err
	// }

	// stakingAddr := consts.NekoStakingContractAddress
	// stakingInstance, err := nekostaking.NewNekostaking(common.HexToAddress(stakingAddr), client)
	// if err != nil {
	// 	e.log.Errorf(err, "[FetchHolders] nekostaking.NewNekoStaking() failed")
	// 	return nil, err
	// }

	// stakings, err := nekoInstance.TokensOfOwner(nil, common.HexToAddress(stakingAddr))
	// if err != nil {
	// 	e.log.Errorf(err, "[FetchHolders] nekoInstance.TokensOfOwner() failed")
	// 	return nil, err
	// }
	// stakingIds := sliceutils.Map(stakings, func(t *big.Int) int64 {
	// 	return t.Int64()
	// })

	var unstakings []*big.Int
	for tokenId := 0; tokenId < col.TotalSupply; tokenId++ {
		// if sliceutils.Contains(stakingIds, int64(tokenId)) {
		// 	continue
		// }
		unstakings = append(unstakings, big.NewInt(int64(tokenId)))
	}

	// var wg sync.WaitGroup
	// wg.Add(2)

	holders := make(map[string]int)
	tokensByHolder := make(map[string][]int64)
	// go func(tokenIds []*big.Int) {
	tokenIds := unstakings
	for _, tokenId := range tokenIds {
		owner, err := tokenTracker.OwnerOf(&bind.CallOpts{}, tokenId)
		if err != nil {
			e.log.Errorf(err, "[FetchHolders] failed to get balance of %d in chain %s", tokenId.Int64(), col.ChainID)
			continue
		}
		e.log.Infof("Token %d owner: %s", tokenId.Int64(), owner.Hex())

		k := strings.ToLower(owner.Hex())
		holders[k]++
		tokensByHolder[k] = append(tokensByHolder[k], tokenId.Int64())
	}
	// 	wg.Done()
	// }(unstakings)

	stakers := make(map[string]int)
	tokensByStaker := make(map[string][]int64)
	// go func(tokenIds []*big.Int) {
	// 	for _, tokenId := range tokenIds {
	// 		staker, err := stakingInstance.Owners(&bind.CallOpts{}, tokenId)
	// 		if err != nil {
	// 			e.log.Errorf(err, "[FetchHolders] stakingInstance.Owners() failed")
	// 			continue
	// 		}

	// 		k := strings.ToLower(staker.Hex())
	// 		stakers[k]++
	// 		tokensByStaker[k] = append(tokensByStaker[k], tokenId.Int64())
	// 	}
	// 	wg.Done()
	// }(stakings)

	// wg.Wait()

	evmAccs, err := e.ListAllWalletAddresses()
	if err != nil {
		e.log.Error(err, "[FetchHolders] ListAllWalletAddresses() failed")
		return nil, err
	}

	evms := map[string]string{}
	for _, acc := range evmAccs {
		k := strings.ToLower(acc.Address)
		evms[k] = acc.ProfileId
	}

	wallets := append(maputils.Keys(holders), maputils.Keys(stakers)...)
	data := sliceutils.Map(wallets, func(w string) model.UserNFTBalance {
		return model.UserNFTBalance{
			UserAddress:     w,
			ChainType:       model.JSONNullString{NullString: sql.NullString{String: "evm", Valid: true}},
			NFTCollectionID: col.ID,
			Balance:         holders[w],
			ProfileID:       evms[w],
			StakingNekos:    stakers[w],
			Metadata:        map[string]interface{}{"holding": tokensByHolder[w], "staking": tokensByStaker[w]},
		}
	})

	return data, nil
}

func (e *Entity) ExistedNekoHolder(colId, addr string) bool {
	existed, _ := e.repo.UserNFTBalance.IsExists(colId, addr)
	return existed
}

func (e *Entity) GetPodTownNFTBalances(collectionAddresses []string) ([]model.PodTownUserNFTBalance, error) {
	return e.repo.UserNFTBalance.GetPodTownUserNFTBalances(collectionAddresses)
}
