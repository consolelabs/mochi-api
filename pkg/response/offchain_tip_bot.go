package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
	"github.com/google/uuid"
)

type OffchainTipBotCreateAssignContractResponse struct {
	Id          uuid.UUID                    `json:"id"`
	TokenId     string                       `json:"token_id"`
	ChainId     string                       `json:"chain_id"`
	UserId      string                       `json:"user_id"`
	ContractId  string                       `json:"contract_id"`
	Status      int                          `json:"status"`
	ExpiredTime time.Time                    `json:"expired_time"`
	Contract    model.OffchainTipBotContract `json:"contract"`
}
