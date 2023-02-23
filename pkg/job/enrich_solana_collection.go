package job

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/service"
)

type enrichSolanaCollection struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewEnrichSolanaCollectionJob(e *entities.Entity, svc *service.Service, l logger.Logger) Job {
	return &enrichSolanaCollection{
		entity:  e,
		service: svc,
		log:     l,
	}
}

type Temp struct {
	CollectionId string `json:"collection_id"`
}

type TempData struct {
	Data []Temp `json:"data"`
}

func (job *enrichSolanaCollection) Run() error {
	content, err := ioutil.ReadFile("/Users/trkhoi/Downloads/solana_collection_top_100.json")
	if err != nil {
		log.Fatal(err)
	}
	collections := TempData{}
	err = json.Unmarshal(content, &collections)
	if err != nil {
		log.Fatal(err)
	}

	for _, collection := range collections.Data {
		job.log.Infof("Adding collection, id: ", collection.CollectionId)
		_, err := job.entity.CreateSolanaNFTCollection(request.CreateNFTCollectionRequest{
			Address:      collection.CollectionId,
			Chain:        "",
			ChainID:      "sol",
			Author:       "393034938028392449",
			GuildID:      "981852028970082344",
			MessageID:    collection.CollectionId,
			ChannelID:    "1052079279619457095",
			PriorityFlag: false,
		})
		if err != nil {
			job.log.Errorf(err, "Error adding collection, id: ", collection.CollectionId)
			continue
		}
	}
	// job.entity.GetTop10kCollectionSolana()
	job.log.Infof("Finish adding 10k collection")
	return nil
}
