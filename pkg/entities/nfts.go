package entities

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
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

func (e *Entity) GetNFTDetail(symbol, tokenId string) (nftsResponse *NFTTokenResponse, err error) {
	// get collection
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		err = fmt.Errorf("failed to get nft collection : %v", err)
		return nil, err
	}
	nftsResponse, err = GetNFTDetailFromIndexer(collection.Address, tokenId)
	if err != nil {
		err = fmt.Errorf("failed to get NFT from indexer: %v", err)
		return nil, err
	}

	if nftsResponse == nil {
		err = fmt.Errorf("response nfts from indexer nil")
		return nil, err
	}

	return nftsResponse, nil
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
	TokenId           uint64          `json:"token_id"`
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

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	Count     int    `json:"count"`
	Rarity    string `json:"rarity"`
	Frequency string `json:"frequency"`
}

func GetNFTDetailFromIndexer(address, tokenId string) (*NFTTokenResponse, error) {
	nftsData := &NFTTokenResponse{}
	//TODO change to prod
	indexerAPI := "https://develop-api.indexer.console.so/api/v1/nft/%s/%s"
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(indexerAPI, address, tokenId), nil)
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
	err = json.Unmarshal(body, &nftsData)
	if err != nil {
		return nil, err
	}
	return nftsData, nil
}
