package processor

import (
	"github.com/defipod/mochi/pkg/model"
)

type Service interface {
	CreateUserTransaction(createUserTransactionRequest model.CreateUserTransaction) (*model.CreateUserTxResponse, error)
	GetUserFactionXp(userDiscordId string) (*model.GetUserFactionXpsResponse, error)
}
