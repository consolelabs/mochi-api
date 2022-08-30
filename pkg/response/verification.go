package response

import "github.com/defipod/mochi/pkg/model"

type NewGuildConfigWalletVerificationMessageResponse struct {
	Status string                                      `json:"status"`
	Data   *model.GuildConfigWalletVerificationMessage `json:"Data"`
}

type GenerateVerificationResponse struct {
	Status string `json:"status"`
	Code   string `json:"code"`
}
