package bluemove

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	res "github.com/defipod/mochi/pkg/response"
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

func (s *bluemoveService) GetCollection(collectionAddress, chainId string) (*model.NFTCollection, error) {
	var client = &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("%s/collections?collection_address=%s&chain_id=%s", s.config.MarketplaceBaseUrl.Bluemove, collectionAddress, chainId), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	res := &res.BluemoveCollectionResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return &model.NFTCollection{
		Address:    collectionAddress,
		Name:       res.Data.Name,
		Symbol:     res.Data.Symbol,
		ChainID:    chainId,
		ERCFormat:  res.Data.ERCFormat,
		IsVerified: res.Data.IsVerified,
		Image:      res.Data.Image,
		Author:     res.Data.Author,
	}, nil
}
