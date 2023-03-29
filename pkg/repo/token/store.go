package token

import "github.com/defipod/mochi/pkg/model"

type ListQuery struct {
	Offset int
	Limit  int
}

type Store interface {
	GetBySymbol(symbol string, botSupported bool) (model.Token, error)
	GetAllSupported(q ListQuery) ([]model.Token, int64, error)
	GetByAddress(address string, chainID int) (*model.Token, error)
	GetDefaultTokens() ([]model.Token, error)
	CreateOne(token *model.Token) error
	UpsertOne(token model.Token) error
	GetAll() ([]model.Token, error)
	Get(id int) (model *model.Token, err error)
	GetSupportedTokenByGuildId(guildID string) ([]model.Token, error)
	GetOneBySymbol(symbol string) (*model.Token, error)
	GetDefaultTokenByGuildID(guildID string) (model.Token, error)
	GetByChainID(chainID int) ([]model.Token, error)
}
