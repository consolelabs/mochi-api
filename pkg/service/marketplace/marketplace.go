package marketplace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/defipod/mochi/pkg/config"
	res "github.com/defipod/mochi/pkg/response"
)

type marketplace struct {
	config *config.Config
}

func NewMarketplace(cfg *config.Config) Service {
	return &marketplace{
		config: cfg,
	}
}

func (e *marketplace) ConvertPaintswapToFtmAddress(paintswapMarketplace string) string {
	splittedPaintswap := strings.Split(paintswapMarketplace, "/")
	return splittedPaintswap[len(splittedPaintswap)-1]
}

func (e *marketplace) ConvertOpenseaToEthAddress(openseaMarketplace string) string {
	splittedOpensea := strings.Split(openseaMarketplace, "/")
	collectionSymbol := splittedOpensea[len(splittedOpensea)-1]
	openseaCollection, _ := e.GetCollectionFromOpensea(collectionSymbol)
	return openseaCollection.Collection.PrimaryAssetContracts[0].Address
}

func (e *marketplace) ConvertQuixoticToOpAddress(quixoticMarketplace string) string {
	splittedQuixotic := strings.Split(quixoticMarketplace, "/")
	collectionSymbol := splittedQuixotic[len(splittedQuixotic)-1]
	quixoticCollection, _ := e.GetCollectionFromQuixotic(collectionSymbol)
	return quixoticCollection.Address
}

func (e *marketplace) HandleMarketplaceLink(contractAddress, chain string) string {
	switch strings.Contains(contractAddress, "/") {
	case false:
		return contractAddress
	case true:
		switch chain {
		case "paintswap":
			return e.ConvertPaintswapToFtmAddress(contractAddress)
		case "opensea":
			return e.ConvertOpenseaToEthAddress(contractAddress)
		case "quixotic":
			return e.ConvertQuixoticToOpAddress(contractAddress)
		default:
			return e.ConvertPaintswapToFtmAddress(contractAddress)
		}
	default:
		return contractAddress
	}
}

func (e *marketplace) GetCollectionFromOpensea(collectionSymbol string) (*res.OpenseaGetCollectionResponse, error) {
	url := fmt.Sprintf("%s/api/v1/collection/%s", e.config.MarketplaceBaseUrl.Opensea, collectionSymbol)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-API-KEY", e.config.MarketplaceApiKey.Opensea)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err = errBody.ReadFrom(response.Body)
		if err != nil {
			return nil, fmt.Errorf("openseaGetCollection - failed to read response: %v", err)
		}

		err = fmt.Errorf("GetNFTCollections - failed to get opensea collections with symbol=%s: %v", collectionSymbol, errBody.String())
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := &res.OpenseaGetCollectionResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// TODO: test with real api key
func (e *marketplace) GetCollectionFromQuixotic(collectionSymbol string) (*res.QuixoticCollectionResponse, error) {
	url := fmt.Sprintf("%s/api/v1/collection/%s", e.config.MarketplaceBaseUrl.Quixotic, collectionSymbol)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-API-KEY", e.config.MarketplaceApiKey.Quixotic)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err = errBody.ReadFrom(response.Body)
		if err != nil {
			return nil, fmt.Errorf("quixoticGetCollection - failed to read response: %v", err)
		}

		err = fmt.Errorf("GetNFTCollections - failed to get quixotic collections with symbol=%s: %v", collectionSymbol, errBody.String())
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := &res.QuixoticCollectionResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *marketplace) GetCollectionFromPaintswap(address string) (*res.PaintswapCollectionResponse, error) {
	url := fmt.Sprintf("%s/v2/collections/%s", e.config.MarketplaceBaseUrl.Painswap, address)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err = errBody.ReadFrom(response.Body)
		if err != nil {
			return nil, fmt.Errorf("paintswapGetAssetContract - failed to read response: %v", err)
		}

		err = fmt.Errorf("GetNFTCollections - failed to get Paintswap asset contract with address=%s: %v", address, errBody.String())
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := &res.PaintswapCollectionResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (e *marketplace) GetOpenseaAssetContract(address string) (*res.OpenseaAssetContractResponse, error) {
	url := fmt.Sprintf("%s/api/v1/asset_contract/%s", e.config.MarketplaceBaseUrl.Opensea, address)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-API-KEY", e.config.MarketplaceApiKey.Opensea)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err = errBody.ReadFrom(response.Body)
		if err != nil {
			return nil, fmt.Errorf("openseaGetAssetContract - failed to read response: %v", err)
		}

		err = fmt.Errorf("GetNFTCollections - failed to get opensea asset contract with address=%s: %v", address, errBody.String())
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	data := &res.OpenseaAssetContractResponse{}
	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
