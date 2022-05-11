package entities

import (
	"fmt"
	"net/http"

	"encoding/json"
	"github.com/defipod/mochi/pkg/config"
	"io/ioutil"
	"strings"
	"time"
)

const (
	fantomChainID     = 250
	fantomChainSymbol = "ftm"
)

func (e *Entity) GetNFTDetail(symbol, tokenId string) (nfts *NFTDetailData, err error) {
	// get collection
	collection, err := e.repo.NFTCollection.GetBySymbol(symbol)
	if err != nil {
		err = fmt.Errorf("failed to get nft collection : %v", err)
		return
	}
	chain := ""
	switch collection.ChainID {
	case fantomChainID:
		chain = fantomChainSymbol
	}

	nfts, err = GetNFTDetailFromMoralis(strings.ToLower(collection.Address), tokenId, chain, e.cfg)
	if err != nil {
		err = fmt.Errorf("failed to get user NFTS: %v", err)
		return
	}
	if nfts == nil {
		err = fmt.Errorf("response from debank nil")
		return
	}
	return
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
