package response

import "github.com/defipod/mochi/pkg/model"

type ListEmojisResponse struct {
	Data []*model.Emojis `json:"data"`
}
