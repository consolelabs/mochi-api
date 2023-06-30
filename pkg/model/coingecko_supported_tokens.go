package model

type CoingeckoSupportedTokens struct {
	ID              string  `json:"id"`
	Symbol          string  `json:"symbol"`
	Name            string  `json:"name"`
	CurrentPrice    float64 `json:"current_price" gorm:"-"`
	MostPopular     bool    `json:"most_popular" gorm:"-"`
	DetailPlatforms []byte  `json:"detail_platforms"`
	IsNative        bool    `json:"is_native"`
}

type CoingeckoDetailPlatform struct {
	ChainId int64  `json:"chain_id"`
	Address string `json:"address"`
}
