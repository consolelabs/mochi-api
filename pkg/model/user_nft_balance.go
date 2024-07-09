package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type metadata map[string]interface{}

type UserNFTBalance struct {
	UserAddress     string         `json:"user_address"`
	ChainType       JSONNullString `json:"chain_type"`
	NFTCollectionID uuid.NullUUID  `json:"nft_collection_id" swaggertype:"string"`
	Balance         int            `json:"balance"`
	StakingNekos    int            `json:"staking_nekos"`
	ProfileID       string         `json:"profile_id"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Metadata        metadata       `json:"metadata" gorm:"type:jsonb"`
}

func (m *metadata) Scan(value interface{}) error {
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("user_nft_balance.metadata has unsupported type: %v", v)
	}

	return json.Unmarshal(bytes, m)
}

func (m metadata) Value() (driver.Value, error) {
	bytes, err := json.Marshal(m)
	return string(bytes), err
}

type UserNFTBalancesByGuild struct {
	UserDiscordId string `json:"user_discord_id"`
	TotalBalance  int64  `json:"total_balance"`
}

type UserAddressNFTBalancesByGuild struct {
	UserAddress     string `json:"user_address"`
	TotalBalance    int64  `json:"total_balance"`
	StakingNeko     int64  `json:"staking_neko"`
	NftCollectionID string `json:"nft_collection_id"`
}

type UserNFTBalanceIdentify struct {
	UserDiscordId   string `json:"user_discord_id"`
	NftCollectionId string `json:"nft_collection_id"`
}

type PodTownUserNFTBalance struct {
	ProfileId   string `json:"profile_id"`
	DiscordId   string `json:"discord_id"`
	UserAddress string `json:"user_address"`
	Neko        int    `json:"neko"`
	Rabby       int    `json:"rabby"`
	Fukuro      int    `json:"fukuro"`
	Gm          int    `json:"gm"`
}
