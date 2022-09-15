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

func (pg *pg) GetByAddressChainId(address, chainId string) (*model.NFTCollection, error) {
	var collection model.NFTCollection
	err := pg.db.Table("nft_collections").Where("lower(address) = ?", strings.ToLower(address)).Where("chain_id = ?", chainId).First(&collection).Error
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

func (pg *pg) GetSuggestionsBySymbolorName(name string, len int) ([]model.NFTCollection, error) {
	var collection []model.NFTCollection
	err := pg.db.Table("nft_collections").Where(fmt.Sprintf("LEVENSHTEIN(symbol, '%s') <= %v OR LEVENSHTEIN(name, '%s') <= %v", name, len, name, len)).Order(fmt.Sprintf("(LENGTH(symbol) = LENGTH('%s')) DESC, levenshtein(symbol, '%s')", name, name)).Limit(5).Find(&collection).Error
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (pg *pg) GetBySymbolorName(name string) ([]model.NFTCollection, error) {
	var collection []model.NFTCollection
	err := pg.db.Table("nft_collections").Where("lower(name) = lower(?) OR lower(symbol) = lower(?)", name, name).
		Where("is_verified = ?", true).Find(&collection).Error
	if err != nil {
		return nil, err
	}
	return collection, nil
}

func (pg *pg) GetByID(id string) (*model.NFTCollection, error) {
	var collection model.NFTCollection
	return &collection, pg.db.Table("nft_collections").Where("id = ?", id).First(&collection).Error
}

func (pg *pg) GetByChain(chain int) ([]model.NFTCollection, int, error) {
	var collections []model.NFTCollection
	var count int64
	err := pg.db.Table("nft_collections").Where(fmt.Sprintf("chain_id='%v'", chain)).Find(&collections).Count(&count)
	if err.Error != nil {
		return nil, 0, err.Error
	}
	return collections, int(count), nil
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
	rows, err := pg.db.Raw(`
	SELECT DISTINCT ON (guild_config_nft_roles.nft_collection_id)
		guild_config_nft_roles.nft_collection_id,
		nft_collections.erc_format,
		nft_collections.address,
		nft_collections. "name",
		nft_collections.chain_id
	FROM
		guild_config_nft_roles
		INNER JOIN nft_collections ON guild_config_nft_roles.nft_collection_id = nft_collections.id
	WHERE
		NOT erc_format = '1155'
	GROUP BY
		nft_collections.id,
		guild_config_nft_roles.id;
	`).Rows()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		tmp := model.NFTCollectionConfig{}
		if err := rows.Scan(&tmp.ID, &tmp.ERCFormat, &tmp.Address, &tmp.Name, &tmp.ChainID); err != nil {
			return nil, err
		}
		configs = append(configs, tmp)
	}

	return configs, nil
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

func (pg *pg) UpdateImage(address string, image string) error {
	return pg.db.Table("nft_collections").Where("address=?", address).Update("image", image).Error
}
