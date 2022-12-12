package response

import (
	"github.com/google/uuid"

	"github.com/defipod/mochi/pkg/model"
)

type NFTSalesTrackerResponse struct {
	ContractAddress string `json:"contract_address"`
	Platform        string `json:"platform"`
	GuildID         string `json:"guild_id"`
	ChannelID       string `json:"channel_id"`
}

type NFTSalesTrackerData struct {
	ID                      uuid.NullUUID                 `json:"id" swaggertype:"string"`
	ContractAddress         string                        `json:"contract_address"`
	Platform                string                        `json:"platform"`
	SalesConfigID           uuid.NullUUID                 `json:"sales_config_id" swaggertype:"string"`
	GuildConfigSalesTracker model.GuildConfigSalesTracker `json:"guild_config_sales_tracker"`
	Name                    string                        `json:"name"`
	Chain                   model.Chain                   `json:"chain"`
}

type NFTSalesTrackerGuildResponse struct {
	ID         uuid.NullUUID         `json:"id" swaggertype:"string"`
	GuildID    string                `json:"guild_id"`
	ChannelID  string                `json:"channel_id"`
	Collection []NFTSalesTrackerData `json:"collection"`
}
