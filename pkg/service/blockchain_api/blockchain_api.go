package blockchainapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/model"
)

type blockchainapi struct {
	config *config.Config
}

func NewService(cfg *config.Config) Service {
	return &blockchainapi{
		config: cfg,
	}
}

var blockChainAPIBaseURL = "https://api.blockchainapi.com/v1/solana/nft/mainnet-beta/"

func (s *blockchainapi) GetSolanaCollection(address string) (*model.SolanaCollectionMetadata, error) {
	var client = &http.Client{}
	req, err := http.NewRequest("GET", blockChainAPIBaseURL+address, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("APIKeyID", s.config.BlockChainAPIKeyID)
	req.Header.Add("APISecretKey", s.config.BlockChainAPISecretKey)

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	res := &model.SolanaCollectionMetadata{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
