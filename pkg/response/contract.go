package response

import "github.com/defipod/mochi/pkg/model"

type GetAssignedContract struct {
	Data model.OffchainTipBotAssignContract `json:"data"`
}
