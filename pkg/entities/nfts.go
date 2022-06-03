package entities

import (
	"fmt"
	"math/big"
	"net/http"

	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"strconv"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/contracts/erc1155"
	"github.com/defipod/mochi/pkg/contracts/erc721"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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

type NFTDetailDataResponse struct {
	TokenAddress string   `json:"token_address"`
	TokenId      string   `json:"token_id"`
	ContractType string   `json:"contract_type"`
	TokenUri     string   `json:"token_uri"`
	Metadata     Metadata `json:"metadata"`
	SyncedAt     string   `json:"synced_at"`
	Amount       string   `json:"amount"`
	Name         string   `json:"name"`
	Symbol       string   `json:"symbol"`
}
type Metadata struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	TokenId     int         `json:"token_id"`
	Attributes  []Attribute `json:"attributes"`
	Image       string      `json:"image"`
	Rarity      interface{} `json:"rarity"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	Count     int    `json:"count"`
	Rarity    string `json:"rarity"`
	Frequency string `json:"frequency"`
}

func (e *Entity) GetNFTDetail(symbol, tokenId string) (nftsResponse *NFTDetailDataResponse, err error) {
	// get collection
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		err = fmt.Errorf("failed to get nft collection : %v", err)
		return nil, err
	}
	chain := ""
	for k, v := range mapChainChainId {
		if strings.ToLower(v) == strings.ToLower(collection.ChainID) {
			chain = k
		}
	}

	nfts := &NFTDetailData{}
	//support for nft rabby // fukuro- get from backendapi
	switch symbol {
	case "rabby", "fukuro":
		nftsResponse, err = GetNFTDetailFromPodtown(*collection, collection.Address, tokenId, symbol)
		if err != nil {
			err = fmt.Errorf("failed to get user NFTS: %v", err)
			return nil, err
		}
		return
	default:
		nfts, err = GetNFTDetailFromMoralis(strings.ToLower(collection.Address), tokenId, chain, e.cfg)
		if err != nil {
			err = fmt.Errorf("failed to get user NFTS: %v", err)
			return nil, err
		}
	}

	if nfts == nil {
		err = fmt.Errorf("response from debank nil")
		return nil, err
	}

	meta := nfts.Metadata
	if nfts.Metadata == "" && nfts.TokenUri != "" {
		client := &http.Client{
			Timeout: time.Second * 60,
		}

		req, err := http.NewRequest("GET", nfts.TokenUri, nil)
		if err != nil {
			return nil, err
		}

		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		meta = string(body)
	}

	var metaParse Metadata
	_ = json.Unmarshal([]byte(meta), &metaParse)
	nftResponse := NFTDetailDataResponse{
		TokenAddress: nfts.TokenAddress,
		TokenId:      nfts.TokenId,
		ContractType: nfts.ContractType,
		TokenUri:     nfts.TokenUri,
		SyncedAt:     nfts.SyncedAt,
		Amount:       nfts.Amount,
		Name:         nfts.Name,
		Symbol:       nfts.Symbol,
		Metadata:     metaParse,
	}
	return &nftResponse, nil
}

type NFTDetailData struct {
	TokenAddress string `json:"token_address"`
	TokenId      string `json:"token_id"`
	ContractType string `json:"contract_type"`
	TokenUri     string `json:"token_uri"`
	Metadata     string `json:"metadata"`
	SyncedAt     string `json:"synced_at"`
	Amount       string `json:"amount"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
}

func GetNFTDetailFromMoralis(address, tokenId, chain string, cfg config.Config) (*NFTDetailData, error) {
	nftsData := &NFTDetailData{}
	moralisApi := "https://deep-index.moralis.io/api/v2/nft/%s/%s?chain=%s&format=decimal"
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(moralisApi, address, tokenId, chain), nil)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &nftsData)
	if err != nil {
		return nil, err
	}

	return nftsData, nil
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
	collection, err := GetNFTCollectionFromMoralis(strings.ToLower(req.Address), req.Chain, e.cfg)
	if err != nil {
		err = fmt.Errorf("failed to get collection NFT from moralis: %v", err)
		return
	}
	if collection == nil {
		err = fmt.Errorf("response collection from moralis nil")
		return
	}

	nftCollection, err = e.repo.NFTCollection.Create(model.NFTCollection{
		Address:   collection.TokenAddress,
		Symbol:    collection.Symbol,
		Name:      collection.Name,
		ChainID:   req.ChainID,
		ERCFormat: collection.ContractType,
	})
	if err != nil {
		err = fmt.Errorf("failed to create collection NFTS: %v", err)
		return
	}
	go PutSyncMoralisNFTCollection(strings.ToLower(req.Address), req.Chain, e.cfg)

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

type NFTTokenResponse struct {
	TokenID           uint64          `json:"tokenId"`
	CollectionAddress string          `json:"collection_address"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Amount            uint64          `json:"amount"`
	Image             string          `json:"image"`
	ImageCDN          string          `json:"image_cdn"`
	ThumbnailCDN      string          `json:"thumbnail_cdn"`
	ImageContentType  string          `json:"image_content_type"`
	Rarity            *NFTTokenRarity `json:"rarity"`
	Attributes        []Attribute     `json:"attributes"`
	MetadataId        string          `json:"metadata_id"`
}

type NFTTokenRarity struct {
	Rank   uint64 `json:"rank"`
	Score  string `json:"score"`
	Total  uint64 `json:"total"`
	Rarity string `json:"rarity,omitempty"`
}

func GetNFTDetailFromPodtown(collection model.NFTCollection, address, tokenId, symbol string) (*NFTDetailDataResponse, error) {
	nftsData := &NFTDetailDataResponse{}
	var r NFTTokenResponse
	podtown := "https://backend.pod.so/api/v1/nft/%s/items/%s"
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(podtown, address, tokenId), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}
	nftsData = &NFTDetailDataResponse{
		TokenAddress: r.CollectionAddress,
		TokenId:      strconv.FormatUint(r.TokenID, 10),
		ContractType: "ERC721",
		Amount:       strconv.FormatUint(r.Amount, 10),
		Name:         r.Name,
		Symbol:       symbol,
		Metadata: Metadata{
			Name:        r.Name,
			Description: r.Description,
			TokenId:     int(r.TokenID),
			Attributes:  r.Attributes,
			Image:       r.Image,
			Rarity:      r.Rarity,
		},
		TokenUri: fmt.Sprintf("https://backend.pod.so/api/v1/nft/%s/metadata/%s", strings.ToLower(collection.Symbol), r.MetadataId),
	}
	return nftsData, nil
}

func (e *Entity) ListAllNFTCollections() ([]model.NFTCollection, error) {
	return e.repo.NFTCollection.ListAll()
}

func (e *Entity) ListAllNFTCollectionConfigs() ([]model.NFTCollectionConfig, error) {
	return e.repo.NFTCollection.ListAllNFTCollectionConfigs()
}

func (e *Entity) GetNFTBalanceFunc(config model.NFTCollectionConfig) (func(address string) (int, error), error) {

	var rpcUrl string
	switch config.ChainID {
	case "1":
		rpcUrl = e.cfg.EthereumRPC
	case "56":
		rpcUrl = e.cfg.BscRPC
	case "250":
		rpcUrl = e.cfg.FantomRPC
	default:
		return nil, fmt.Errorf("chain id %s not supported", config.ChainID)
	}

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to chain client: %v", err.Error())
	}

	var balanceOf func(string) (int, error)
	switch config.ERCFormat {
	case "721":
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

	case "1155":
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
