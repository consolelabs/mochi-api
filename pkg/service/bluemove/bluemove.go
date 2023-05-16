package bluemove

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type bluemoveService struct {
	config *config.Config
	logger logger.Logger
}

func New(cfg *config.Config, l logger.Logger) Service {
	return &bluemoveService{
		config: cfg,
		logger: l,
	}
}

func (b *bluemoveService) GetCollections(chainId, page, pageSize string) (*response.BluemoveCollectionsResponse, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/collections?pagination[page]=%s&pagination[pageSize]=%s", b.ChooseBluemoveChain(chainId), page, pageSize), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := &response.BluemoveCollectionsResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *bluemoveService) SelectBluemoveCollection(collectionAddress, chainId string) (*model.NFTCollection, error) {
	page := 0
	pageSize := 100
	mapCollections := make([]response.BluemoveCollectionDetail, 0)
	for {
		collections, err := b.GetCollections(chainId, fmt.Sprintf("%d", page), fmt.Sprintf("%d", pageSize))
		if err != nil {
			return nil, err
		}

		mapCollections = append(mapCollections, collections.Data...)

		if int64(page) == collections.Meta.Pagination.PageCount {
			break
		}
		page++
	}

	// select collection
	for _, collection := range mapCollections {
		if collection.Attributes.Creator == collectionAddress || util.GetSuiAddressCollection(collection.Attributes.Type) == collectionAddress {
			return &model.NFTCollection{
				ChainID: chainId,
				Name:    collection.Attributes.Name,
				Symbol:  util.GetSymbolSuiCollection(collection.Attributes.Slug),
				Address: collectionAddress,
				Image:   collection.Attributes.Uri,
				Author:  strconv.Itoa(int(collection.Id)),
			}, nil
		}
	}
	return nil, errors.New("not found collection")
}

func (b *bluemoveService) ChooseBluemoveChain(chainId string) string {
	switch chainId {
	case "9999":
		return b.config.MarketplaceBaseUrl.BluemoveAptos
	case "9996":
		return b.config.MarketplaceBaseUrl.BluemoveSuiMainnet
	default:
		return b.config.MarketplaceBaseUrl.BluemoveAptos
	}
}
