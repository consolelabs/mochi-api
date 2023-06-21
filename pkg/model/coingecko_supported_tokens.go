package model

import "gorm.io/datatypes"

type CoingeckoSupportedTokens struct {
	ID              string  `json:"id"`
	Symbol          string  `json:"symbol"`
	Name            string  `json:"name"`
	CurrentPrice    float64 `json:"current_price" gorm:"-"`
	MostPopular     bool    `json:"most_popular" gorm:"-"`
	DetailPlatforms []byte  `json:"detail_platforms"`

	CoingeckoInfo datatypes.JSON `json:"coingecko_info" gorm:"type:jsonb;default:null"`
	// Contracts   map[string]string `json:"contracts" gorm:"-"`
	// Websites    map[string]string `json:"websites" gorm:"-"`
	// Explorers   map[string]string `json:"explorers" gorm:"-"`
	// Communities map[string]string `json:"communities" gorm:"-"`
	// Tags        []string          `json:"tags" gorm:"-"`
}

type CoingeckoDetailPlatform struct {
	ChainId int64  `json:"chain_id"`
	Address string `json:"address"`
}
