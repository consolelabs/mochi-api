package model

import "time"

type ProductHashtagAlias struct {
	Id               int64           `json:"id"`
	ProductHashtagId int64           `json:"product_hashtag_id"`
	Alias            string          `json:"alias"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	ProductHashtag   *ProductHashtag `json:"product_hashtag" gorm:"foreignkey:ProductHashtagId;references:Id;<-:false"`
}
