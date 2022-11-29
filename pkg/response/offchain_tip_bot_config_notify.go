package response

import "github.com/defipod/mochi/pkg/model"

type ListConfigNotifyResponse struct {
	Data []model.OffchainTipBotConfigNotify `json:"data"`
}
