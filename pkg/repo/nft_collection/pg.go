package nftcollection

import (
	"fmt"
	"strings"

	"github.com/defipod/mochi/pkg/model"
	"gorm.io/gorm"
)

type pg struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) Store {
	return &pg{
		db: db,
	}
}

func (pg *pg) GetByAddress(address string) (*model.NFTCollection, error) {
	var collection model.NFTCollection
	err := pg.db.Table("nft_collections").Where("lower(address) = ?", strings.ToLower(address)).First(&collection).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (pg *pg) GetBySymbol(symbol string) (*model.NFTCollection, error) {
	var collection model.NFTCollection
	err := pg.db.Table("nft_collections").Where("lower(symbol) = lower(?)", symbol).
		Where("is_verified = ?", true).First(&collection).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (pg *pg) GetByName(name string) (*model.NFTCollection, error) {
	var collection model.NFTCollection
	err := pg.db.Table("nft_collections").Where("lower(name) = lower(?)", name).
		Where("is_verified = ?", true).First(&collection).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}

func (pg *pg) GetByID(id string) (*model.NFTCollection, error) {
	var collection model.NFTCollection
	return &collection, pg.db.Table("nft_collections").Where("id = ?", id).First(&collection).Error
}

func (pg *pg) Create(collection model.NFTCollection) (*model.NFTCollection, error) {
	return &collection, pg.db.Table("nft_collections").Create(&collection).Error
}

func (pg *pg) ListAll() ([]model.NFTCollection, error) {
	var collections []model.NFTCollection
	return collections, pg.db.Table("nft_collections").Find(&collections).Error
}
func (pg *pg) ListAllWithPaging(page int, size int) ([]model.NFTCollection, int64, error) {
	var collection []model.NFTCollection
	var count int64
	return collection, count, pg.db.Table("nft_collections").
		Count(&count).
		Limit(size).
		Offset(size * page).
		Find(&collection).Error
}

func (pg *pg) ListAllNFTCollectionConfigs() ([]model.NFTCollectionConfig, error) {
	var configs []model.NFTCollectionConfig
	return configs, pg.db.
		Table("nft_collections").
		Select("nft_collections.*, token_id").
		Joins("left join guild_config_nft_roles on guild_config_nft_roles.nft_collection_id = nft_collections.id").
		Where("not (erc_format = '1155' and token_id is null)").
		Group("nft_collections.id, token_id").
		Find(&configs).Error
}

func (pg *pg) ListByGuildID(guildID string) ([]model.NFTCollection, error) {
	var collections []model.NFTCollection
	return collections, pg.db.Table("nft_collections").
		Joins("left join guild_config_nft_roles on guild_config_nft_roles.nft_collection_id = nft_collections.id").
		Where("guild_id = ?", guildID).
		Find(&collections).Error
}

func (pg *pg) GetNewListed(interval int, page int, size int) ([]model.NewListedNFTCollection, int64, error) {
	var collection []model.NewListedNFTCollection
	var count int64
	return collection, count, pg.db.Table("nft_collections").
		Where(fmt.Sprintf("created_at > now() - interval '%v days'", interval)). //error if uses placeholder
		Order("created_at DESC").
		Count(&count).
		Limit(size).
		Offset(size * page).
		Find(&collection).Error
}
