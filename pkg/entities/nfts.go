package entities

import (
	"fmt"
	"net/http"

	"encoding/json"
	"io/ioutil"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/config"
)

const (
	fantomChainID     = 250
	fantomChainSymbol = "ftm"
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
	switch collection.ChainID {
	case fantomChainID:
		chain = fantomChainSymbol
	}

	nfts, err := GetNFTDetailFromMoralis(strings.ToLower(collection.Address), tokenId, chain, e.cfg)
	if err != nil {
		err = fmt.Errorf("failed to get user NFTS: %v", err)
		return nil, err
	}

	if nfts == nil {
		err = fmt.Errorf("response from debank nil")
		return nil, err
	}

	meta := nfts.Metadata
	var metaParse Metadata
	_ = json.Unmarshal([]byte(meta), &metaParse)
	fmt.Println(metaParse)
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
