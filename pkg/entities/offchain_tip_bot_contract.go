package entities

import (
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
	"gorm.io/gorm"
)

func (e *Entity) OffchainTipBotCreateAssignContract(ac *model.OffchainTipBotAssignContract) (*response.OffchainTipBotCreateAssignContractResponse, error) {
	err := e.repo.OffchainTipBotContract.CreateAssignContract(ac)
	if err != nil {
		e.log.Fields(logger.Fields{"model": ac}).Error(err, "[entities.OffchainTipBotCreateAssignContract] - failed to create assign contract")
		return nil, err
	}

	userAssignedContract, err := e.repo.OffchainTipBotContract.GetAssignedContractByUserId(ac.UserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		e.log.Fields(logger.Fields{"model": ac}).Error(err, "[entities.OffchainTipBotCreateAssignContract] - failed to get assigned contract")
		return nil, err
	}

	response := &response.OffchainTipBotCreateAssignContractResponse{
		Id:          ac.ID,
		TokenId:     ac.TokenID,
		ChainId:     ac.ChainID,
		UserId:      ac.UserID,
		ContractId:  ac.ContractID,
		Status:      ac.Status,
		ExpiredTime: ac.ExpiredTime,
		Contract:    userAssignedContract.OffchainTipBotContract,
	}

	return response, nil
}
