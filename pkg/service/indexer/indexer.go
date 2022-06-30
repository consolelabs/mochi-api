package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	res "github.com/defipod/mochi/pkg/response"
)

type CreateERC721ContractRequest struct {
	Address string `json:"address"`
	ChainID int    `json:"chain_id"`
}

type indexer struct {
	cfg config.Config
	log logger.Logger
}

func NewIndexer(cfg config.Config, log logger.Logger) Service {
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

func (i *indexer) GetNFTCollectionTickers(address string) (*res.IndexerNFTCollectionTickersResponse, error) {
	url := fmt.Sprintf("%s/api/v1/nft/ticker/%s", i.cfg.IndexerServerHost, address)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err = errBody.ReadFrom(response.Body)
		if err != nil {
			return nil, fmt.Errorf("GetNFTCollection - failed to read response: %v", err)
		}

		err = fmt.Errorf("GetNFTCollection - failed to get nft collection info %s: %v", address, errBody.String())
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := &res.IndexerNFTCollectionTickersResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return data, nil
}

func (i *indexer) GetNFTCollections(query string) (*res.IndexerGetNFTCollectionsResponse, error) {

	url := fmt.Sprintf("%s/api/v1/nft?%s", i.cfg.IndexerServerHost, query)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err = errBody.ReadFrom(response.Body)
		if err != nil {
			return nil, fmt.Errorf("GetNFTCollections - failed to read response: %v", err)
		}

		err = fmt.Errorf("GetNFTCollections - failed to filter nft collections with query=%s: %v", query, errBody.String())
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := &res.IndexerGetNFTCollectionsResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return data, nil
}

func (i *indexer) GetNFTTokens(address, query string) (*res.IndexerGetNFTTokensResponse, error) {

	url := fmt.Sprintf("%s/api/v1/nft/%s?%s", i.cfg.IndexerServerHost, address, query)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err = errBody.ReadFrom(response.Body)
		if err != nil {
			return nil, fmt.Errorf("GetNFTTokens - failed to read response: %v", err)
		}

		err = fmt.Errorf("GetNFTTokens - failed to filter nft tokens with symbol=%s, query=%s: %v", address, query, errBody.String())
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := &res.IndexerGetNFTTokensResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return data, nil
}

func (i *indexer) GetNFTTradingVolume() ([]res.NFTTradingVolume, error) {
	url := fmt.Sprintf("%s/api/v1/nft/ticker", i.cfg.IndexerServerHost)

	nftList := res.NFTTradingVolumeResponse{}
	client := &http.Client{Timeout: time.Second * 30}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(b), &nftList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch nft indexer")
	}

	return nftList.Data, nil
}

func (i *indexer) GetNFTDetail(collectionAddress, tokenID string) (*res.IndexerNFTToken, error) {
	//--check data in sync
	contract, err := i.GetNFTContract(collectionAddress)
	if err != nil {
		return nil, err
	}
	if !contract.IsSynced {
		err = fmt.Errorf("data not in sync")
		return nil, err
	}
	//--
	data := &res.IndexerNFTToken{}
	errorMsg := &res.ErrorMessage{}
	url := "%s/api/v1/nft/%s/%s"
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(url, i.cfg.IndexerServerHost, collectionAddress, tokenID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	// err still == nil even if indexer return error
	if err != nil {
		err = fmt.Errorf("GetNFTDetail - failed to get record")
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("GetNFTDetail - failed to read response body")
		return nil, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		err = fmt.Errorf("GetNFTDetail - failed to unmarshal response data")
		return nil, err
	}
	// parse error msg, empty if not
	if err := json.Unmarshal(body, &errorMsg); err != nil {
		err = fmt.Errorf("GetNFTDetail - failed to unmarshal response data")
		return nil, err
	}
	// if id too large api will return out of range and unsigned number instead
	if strings.Contains(errorMsg.Error, "record not found") || strings.Contains(errorMsg.Error, "out of range") || strings.Contains(errorMsg.Error, "unsigned number") {
		err = fmt.Errorf("record not found")
		return nil, err
	}
	return data, nil
}

func (i *indexer) GetNftSales(addr string, platform string) (*res.NftSalesResponse, error) {
	data := &res.NftSalesResponse{}
	url := "%s/api/v1/nft/sales?collection_address=%s&platform=%s"
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf(url, i.cfg.IndexerServerHost, addr, platform), nil)
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
		err = fmt.Errorf("GetNFTSales - failed to read response body")
		return nil, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		err = fmt.Errorf("GetNFTSales - failed to unmarshal response data")
		return nil, err
	}

	return data, nil
}

func (i *indexer) GetNFTContract(address string) (*res.IndexerContract, error) {
	url := fmt.Sprintf("%s/api/v1/contract/%s", i.cfg.IndexerServerHost, address)

	contract := &res.IndexerContract{}
	client := &http.Client{Timeout: time.Second * 30}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(b), &contract)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("GETNFTContract - failed to unmarshal data")
	}

	return contract, nil
}
