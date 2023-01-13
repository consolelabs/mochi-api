package entities

import (
	"fmt"
	"strings"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetNftSales(addr string, platform string) (*response.NftSalesResponse, error) {
	nft, err := e.indexer.GetNftSales(addr, platform)
	if err != nil {
		err = fmt.Errorf("failed to get sales from indexer: %v", err)
		return nil, err
	}

	return nft, nil
}

func (e *Entity) CreateSalesTrackerConfig(req request.CreateSalesTrackerConfigRequest) error {
	if !strings.EqualFold(req.ContractAddress, "all") {
		checkExistNFT, err := e.CheckExistNftCollection(req.ContractAddress)
		if err != nil {
			e.log.Errorf(err, "[e.CheckExistNftCollection] failed to check if nft exist: %v", err)
			return err
		}

		if !checkExistNFT {
			e.log.Fields(logger.Fields{"guildID": req.GuildID, "contractAddress": req.ContractAddress}).Info("[e.CreateSalesTracker] Collection has not been added.")
			return fmt.Errorf("Collection has not been added.")
		}
	}

	// create config sales tracker
	if strings.EqualFold(req.ContractAddress, "all") {
		req.ContractAddress = "*"
	}

	if strings.EqualFold(req.Chain, "all") {
		req.Chain = "*"
	}

	err := e.repo.GuildConfigSalesTracker.Create(&model.GuildConfigSalesTracker{
		GuildID:         req.GuildID,
		ChannelID:       req.ChannelID,
		Chain:           req.Chain,
		ContractAddress: req.ContractAddress,
	})
	if err != nil {
		e.log.Errorf(err, "[e.repo.GuildConfigSalesTracker.Create] failed to create sales tracker: %v", err)
		return err
	}
	return nil
}

func (e *Entity) GetSalesTrackerConfig(guildID string) ([]model.GuildConfigSalesTracker, error) {
	config, err := e.repo.GuildConfigSalesTracker.GetByGuildID(guildID)
	if err != nil {
		e.log.Errorf(err, "[e.repo.GuildConfigSalesTracker.GetByGuildID] failed to get sales tracker: %v", err)
		return nil, err
	}

	return config, nil
}

func (e *Entity) GetAllSalesTrackerConfig() ([]model.GuildConfigSalesTracker, error) {
	return e.repo.GuildConfigSalesTracker.GetAllSalesTrackerConfig()
}
