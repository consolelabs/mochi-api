package token

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	GetBySymbol(symbol string, botSupported bool) (model.Token, error)
	GetAllSupported() ([]model.Token, error)
	GetByAddress(address string, chainID int) (*model.Token, error)
	GetDefaultTokens() ([]model.Token, error)
	CreateOne(token *model.Token) error
	UpsertOne(token model.Token) error
	GetAll() ([]model.Token, error)
	GetAllSupportedToken(guildID string) ([]model.Token, error)
	GetOneBySymbol(symbol string) (*model.Token, error)
}
