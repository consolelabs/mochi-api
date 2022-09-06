package response

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
)

type NFTSalesTrackerResponse struct {
	ContractAddress string `json:"contract_address"`
	Platform        string `json:"platform"`
	GuildID         string `json:"guild_id"`
	ChannelID       string `json:"channel_id"`
}

type NFTSalesTrackerData struct {
	ID                      uuid.NullUUID                 `json:"id"`
	ContractAddress         string                        `json:"contract_address"`
	Platform                string                        `json:"platform"`
	SalesConfigID           uuid.NullUUID                 `json:"sales_config_id"`
	GuildConfigSalesTracker model.GuildConfigSalesTracker `json:"guild_config_sales_tracker"`
	Name                    string                        `json:"name"`
	Chain                   model.Chain                   `json:"chain"`
}

type NFTSalesTrackerGuildResponse struct {
	ID         uuid.NullUUID         `json:"id"`
	GuildID    string                `json:"guild_id"`
	ChannelID  string                `json:"channel_id"`
	Collection []NFTSalesTrackerData `json:"collection"`
}
