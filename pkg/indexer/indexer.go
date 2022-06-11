package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
)

type CreateERC721ContractRequest struct {
	Address string `json:"address"`
	ChainID int    `json:"chain_id"`
}

type indexer struct {
	cfg config.Config
	log logger.Logger
}

func NewIndexer(cfg config.Config, log logger.Logger) Indexer {
	return &indexer{
		cfg: cfg,
		log: log,
	}
}

func (i *indexer) CreateERC721Contract(req CreateERC721ContractRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	jsonBody := bytes.NewBuffer(body)

	url := fmt.Sprintf("%s/api/v1/contract/erc721", i.cfg.IndexerServerHost)
	request, err := http.NewRequest("POST", url, jsonBody)
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err := errBody.ReadFrom(response.Body)
		if err != nil {
			return fmt.Errorf("CreateERC721Contract - failed to read response: %v", err)
		}
		i.log.Errorf(err, "CreateERC721Contract error: %s | chain_id %d", req.Address, req.ChainID)
		return fmt.Errorf("CreateERC721Contract - failed to create erc721 contract: %v", errBody.String())
	}

	defer response.Body.Close()
	return nil
}
