package offchain_tip_bot_transfer_histories

import (
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/response"
)

type Store interface {
	CreateTransferHistories(transferHistories []model.OffchainTipBotTransferHistory) ([]model.OffchainTipBotTransferHistory, error)
	GetByUserDiscordId(userDiscordId string) (transferHistories []model.OffchainTipBotTransferHistory, err error)
	TotalFeeFromWithdraw() ([]response.TotalFeeWithdraw, error)
	GetTransactionsByQuery(receiverId, senderId, token string) ([]model.OffchainTipBotTransferHistory, error)
	GetTotalTransactionByGuildAndToken(guildId, token string) ([]model.OffchainTipBotTransferHistory, int64, error)
	GetTotalTransactionByGuild(guildId string) (count int64, err error)
}
