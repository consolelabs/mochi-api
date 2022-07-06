package entities

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/contracts/erc1155"
	"github.com/defipod/mochi/pkg/contracts/erc721"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/service/indexer"
	"github.com/defipod/mochi/pkg/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
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
	}
)

func (e *Entity) GetNFTDetail(symbol, tokenID string) (*response.IndexerNFTToken, error) {
	// get collection
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	// cannot find collection in db
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("database: record nft collection not found")
		} else {
			err = fmt.Errorf("failed to get nft collection : %v", err)
		}
		return nil, err
	}

	data, err := e.indexer.GetNFTDetail(collection.Address, tokenID)
	// cannot find collection in indexer
	if err != nil {
		if err.Error() == "record not found" {
			err = fmt.Errorf("indexer: record nft not found")
		} else {
			err = fmt.Errorf("failed to get nft from indexer: %v", err)
		}
		return nil, err
	}

	// empty response
	if data == nil {
		err := fmt.Errorf("no nft data from indexer")
		return nil, err
	}

	return data, nil
}

func (e *Entity) GetNFTDetailWithName(name, tokenID string) (*response.IndexerNFTToken, error) {
	// get collection
	collection, err := e.repo.NFTCollection.GetByName(name)
	// cannot find collection in db
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("database: record nft collection not found")
		} else {
			err = fmt.Errorf("failed to get nft collection : %v", err)
		}
		return nil, err
	}

	data, err := e.indexer.GetNFTDetail(collection.Address, tokenID)
	// cannot find collection in indexer
	if err != nil {
		if err.Error() == "record not found" {
			err = fmt.Errorf("indexer: record nft not found")
		} else {
			err = fmt.Errorf("failed to get nft from indexer: %v", err)
		}
		return nil, err
	}

	// empty response
	if data == nil {
		err := fmt.Errorf("no nft data from indexer")
		return nil, err
	}

	return data, nil
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

func (e *Entity) CheckExistNftCollection(address string) (bool, error) {
	_, err := e.repo.NFTCollection.GetByAddress(address)
	// cannot find collection in db
	if err != nil {
		if err.Error() == "record not found" {
			return false, nil
		} else {
			err = errors.New("failed to get nft collection")
			return false, err
		}
	}
	return true, nil
}

func (e *Entity) CheckIsSync(address string) (bool, error) {
	indexerContract, err := e.indexer.GetNFTContract(address)
	if err != nil {
		return false, err
	}

	return indexerContract.IsSynced, nil
}

func GetNFTCollectionFromMoralis(address, chain string, cfg config.Config) (*NFTCollectionData, error) {
	colData := &NFTCollectionData{}
	moralisApi := "https://deep-index.moralis.io/api/v2/nft/%s/metadata?chain=%s"
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(moralisApi, address, chain), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-API-Key", cfg.MoralisXApiKey)
	q := req.URL.Query()

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		mesErr := &MoralisMessageFail{}
		mes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(mes, &mesErr)
		if err != nil {
			return nil, err
		}
		err = fmt.Errorf("%v", mesErr.Message)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &colData)
	if err != nil {
		return nil, err
	}

	return colData, nil
}

func (e *Entity) CreateNFTCollection(req request.CreateNFTCollectionRequest) (nftCollection *model.NFTCollection, err error) {
	address := e.HandleMarketplaceLink(req.Address, req.ChainID)
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

	if checkExistNFT {
		is_sync, err := e.CheckIsSync(checksumAddress)
		if err != nil {
			e.log.Errorf(err, "[e.CheckIsSync] failed to check if nft is synced: %v", err)
			return nil, err
		}

		if !is_sync {
			return nil, fmt.Errorf("Already added. Nft is in sync progress")
		} else {
			return nil, fmt.Errorf("Already added. Nft is done with sync")
		}
	}

	req.Address = checksumAddress
	convertedChainId := util.ConvertChainToChainId(req.ChainID)
	chainID, err := strconv.Atoi(convertedChainId)
	if err != nil {
		e.log.Errorf(err, "[util.ConvertChainToChainId] failed to convert chain to chainId: %v", err)
		return nil, fmt.Errorf("Failed to convert chain to chainId: %v", err)
	}
	// query name and symbol from contract
	name, symbol, err := e.abi.GetNameAndSymbol(req.Address, int64(chainID))
	if err != nil {
		e.log.Errorf(err, "[GetNameAndSymbol] cannot get name and symbol of contract: %s | chainId %d", req.Address, chainID)
		return nil, fmt.Errorf("Cannot get name and symbol of contract: %v", err)
	}

	err = e.indexer.CreateERC721Contract(indexer.CreateERC721ContractRequest{
		Address: req.Address,
		ChainID: chainID,
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
	})
	if err != nil {
		e.log.Errorf(err, "[repo.NFTCollection.Create] cannot add collection: %v", err)
		return nil, fmt.Errorf("Cannot add collection: %v", err)
	}

	return
}

func PutSyncMoralisNFTCollection(address, chain string, cfg config.Config) (err error) {
	time.Sleep(1000)
	moralisApi := "https://deep-index.moralis.io/api/v2/nft/%s/sync?chain=%s"
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf(moralisApi, address, chain), nil)
	if err != nil {
		return
	}
	req.Header.Add("X-API-Key", cfg.MoralisXApiKey)
	q := req.URL.Query()

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		mesErr := &MoralisMessageFail{}
		mes, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(mes, &mesErr)
		if err != nil {
			return
		}
		err = fmt.Errorf("%v", mesErr.Message)
		return
	}

	return nil
}

func (e *Entity) ListAllNFTCollections() ([]model.NFTCollection, error) {
	return e.repo.NFTCollection.ListAll()
}

func (e *Entity) ListAllNFTCollectionConfigs() ([]model.NFTCollectionConfig, error) {
	return e.repo.NFTCollection.ListAllNFTCollectionConfigs()
}

func (e *Entity) GetNFTBalanceFunc(config model.NFTCollectionConfig) (func(address string) (int, error), error) {

	chainID, err := strconv.Atoi(config.ChainID)
	if err != nil {
		return nil, fmt.Errorf("failed to convert chain id %s to int: %v", config.ChainID, err)
	}

	chain, err := e.repo.Chain.GetByID(chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain by id %s: %v", config.ChainID, err)
	}

	client, err := ethclient.Dial(chain.RPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to chain client: %v", err.Error())
	}

	var balanceOf func(string) (int, error)
	switch strings.ToLower(config.ERCFormat) {
	case "721", "erc721":
		contract721, err := erc721.NewErc721(common.HexToAddress(config.Address), client)
		if err != nil {
			return nil, fmt.Errorf("failed to init erc721 contract: %v", err.Error())
		}

		balanceOf = func(address string) (int, error) {
			b, err := contract721.BalanceOf(nil, common.HexToAddress(address))
			if err != nil {
				return 0, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, config.ChainID, err.Error())
			}
			return int(b.Int64()), nil
		}

	case "1155", "erc1155":
		contract1155, err := erc1155.NewErc1155(common.HexToAddress(config.Address), client)
		if err != nil {
			return nil, fmt.Errorf("failed to init erc1155 contract: %v", err.Error())
		}

		tokenID, err := strconv.ParseInt(config.TokenID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("token id is not valid")
		}

		balanceOf = func(address string) (int, error) {
			b, err := contract1155.BalanceOf(nil, common.HexToAddress(address), big.NewInt(tokenID))
			if err != nil {
				return 0, fmt.Errorf("failed to get balance of %s in chain %s: %v", address, config.ChainID, err.Error())
			}
			return int(b.Int64()), nil
		}

	default:
		return nil, fmt.Errorf("erc format %s not supported", config.ERCFormat)
	}

	return balanceOf, nil
}

func (e *Entity) NewUserNFTBalance(balance model.UserNFTBalance) error {
	err := e.repo.UserNFTBalance.Upsert(balance)
	if err != nil {
		return fmt.Errorf("failed to upsert user nft balance: %v", err.Error())
	}
	return nil
}

func (e *Entity) GetNFTCollection(symbol string) (*response.IndexerNFTCollectionTickersResponse, error) {
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	data, err := e.indexer.GetNFTCollectionTickers(collection.Address)
	if err != nil {
		return nil, err
	}

	for _, ts := range data.Tickers.Timestamps {
		time := time.UnixMilli(ts)
		data.Tickers.Times = append(data.Tickers.Times, time.Format("01-02"))
	}
	return data, nil
}

func (e *Entity) GetNFTCollections(query string) (*response.IndexerGetNFTCollectionsResponse, error) {
	data, err := e.svc.Indexer.GetNFTCollections(query)
	if err != nil {
		return nil, err
	}

	chainMap := make(map[int]model.Chain)
	if chains, err := e.repo.Chain.GetAll(); err == nil {
		for _, chain := range chains {
			chainMap[chain.ID] = chain
		}
	}

	// TODO: remove after indexer data's better
	collectionMap := make(map[string]model.NFTCollection)
	if collections, err := e.repo.NFTCollection.ListAll(); err == nil {
		for _, col := range collections {
			collectionMap[col.Address] = col
		}
	}

	for i, collection := range data.Data {
		data.Data[i].Image = util.StandardizeUri(collection.Image)
		if chain, ok := chainMap[collection.ChainId]; ok {
			data.Data[i].Chain = &chain
		}

		// TODO: remove after indexer data's better
		if idxCol, ok := collectionMap[collection.Address]; ok {
			if collection.Name == "" && idxCol.Name != "" {
				data.Data[i].Name = idxCol.Name
			}
			if collection.Symbol == "" && idxCol.Symbol != "" {
				data.Data[i].Symbol = idxCol.Symbol
			}
		}
	}

	return data, nil
}

func (e *Entity) GetNFTTokens(symbol, query string) (*response.IndexerGetNFTTokensResponse, error) {
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}
	if collection.Address == "" {
		return nil, fmt.Errorf("invalid address - collection %s", collection.ID.UUID)
	}
	data, err := e.svc.Indexer.GetNFTTokens(collection.Address, query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (e *Entity) CreateNFTSalesTracker(addr string, platform string, guildID string) error {
	checksum, err := util.ConvertToChecksumAddr(addr)
	if err != nil {
		e.log.Errorf(err, "[util.ConvertToChecksumAddr] cannot convert to checksum")
		return fmt.Errorf("invalid contract address")
	}
	config, err := e.GetSalesTrackerConfig(guildID)
	if err != nil {
		return err
	}

	return e.repo.NFTSalesTracker.FirstOrCreate(&model.InsertNFTSalesTracker{
		ContractAddress: checksum,
		Platform:        platform,
		SalesConfigID:   config.ID.UUID.String(),
	})
}

func (e *Entity) GetDetailNftCollection(symbol string) (*model.NFTCollection, error) {
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		return nil, err
	}

	res, err := e.svc.Indexer.GetNFTCollections(fmt.Sprintf("address=%s", collection.Address))
	if err == nil && len(res.Data) > 0 {
		collection.Image = util.StandardizeUri(res.Data[0].Image)
	}
	return collection, nil
}

func (e *Entity) GetDetailNftCollectionWithName(name string) (*model.NFTCollection, error) {
	collection, err := e.repo.NFTCollection.GetByName(name)
	if err != nil {
		return nil, err
	}

	res, err := e.svc.Indexer.GetNFTCollections(fmt.Sprintf("address=%s", collection.Address))
	if err == nil && len(res.Data) > 0 {
		collection.Image = util.StandardizeUri(res.Data[0].Image)
	}
	return collection, nil
}

func (e *Entity) GetAllNFTSalesTracker() ([]response.NFTSalesTrackerResponse, error) {
	resp := []response.NFTSalesTrackerResponse{}
	data, err := e.repo.NFTSalesTracker.GetAll()
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		resp = append(resp, response.NFTSalesTrackerResponse{
			ContractAddress: item.ContractAddress,
			Platform:        item.Platform,
			GuildID:         item.GuildConfigSalesTracker.GuildID,
			ChannelID:       item.GuildConfigSalesTracker.ChannelID,
		})
	}
	return resp, nil
}

func (e *Entity) GetNewListedNFTCollection(interval string, page string, size string) (*response.NFTNewListedResponse, error) {
	itv, _ := strconv.Atoi(interval)
	pg, _ := strconv.Atoi(page)
	lim, _ := strconv.Atoi(size)
	data, total, err := e.repo.NFTCollection.GetNewListed(itv, pg, lim)
	for i, ele := range data {
		chainId, _ := strconv.Atoi(ele.ChainID)
		chain, err := e.repo.Chain.GetByID(chainId)
		if err != nil {
			return nil, err
		}
		data[i].Chain = chain.Name
	}
	return &response.NFTNewListedResponse{
		Pagination: util.Pagination{
			Page:  int64(pg),
			Size:  int64(lim),
			Total: total,
		},
		Data: data,
	}, err
}
