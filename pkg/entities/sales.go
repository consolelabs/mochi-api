package entities

import (
	"fmt"

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

func (e *Entity) CreateSalesTracker(req request.NFTSalesTrackerRequest) error {
	err := e.UpsertSalesTrackerConfig(request.UpsertSalesTrackerConfigRequest{
		GuildID:   req.GuildID,
		ChannelID: req.ChannelID,
	})
	if err != nil {
		e.log.Errorf(err, "[e.UpsertSalesTrackerConfig] cannot upsert sales tracker config. GuildId: %s, ChannelId %s", req.GuildID, req.ChannelID)
		return fmt.Errorf("cannot upsert sales tracker config. Error: %v", err)
	}

	err = e.CreateNFTSalesTracker(req.ContractAddress, req.Platform, req.GuildID)
	if err != nil {
		e.log.Errorf(err, "[e.CreateNFTSalesTracker] cannot create nft sales tracker. Contract: %s, Platform: %s, GuildId: %s", req.ContractAddress, req.Platform, req.GuildID)
		return fmt.Errorf("cannot create nft sales tracker. Error: %v", err)
	}
	return nil
}
