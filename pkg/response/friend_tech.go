package response

import (
	"time"

	"github.com/defipod/mochi/pkg/model"
)

type FriendTechKeysResponse struct {
	Data []FriendTechKey `json:"data"`
}

type FriendTechKey struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Address         string    `json:"address"`
	TwitterUsername string    `json:"twitterUsername"`
	TwitterPfpUrl   string    `json:"twitterPfpUrl"`
	ProfileChecked  bool      `json:"profileChecked"`
	Price           float64   `json:"price"`
	Supply          int       `json:"supply"`
	Holders         int       `json:"holders"`
}

type GetUserFriendTechKeyWatchlistResponse struct {
	Data []FriendTechKeyWatchlistItemRespose `json:"data"`
}

type TrackFriendTechKeyResponse struct {
	Data FriendTechKeyWatchlistItemRespose `json:"data"`
}

type FriendTechKeyWatchlistItemRespose struct {
	Id              int            `json:"id"`
	KeyAddress      string         `json:"key_address"`
	ProfileId       string         `json:"profile_id"`
	IncreaseAlertAt int            `json:"increase_alert_at"`
	DecreaseAlertAt int            `json:"decrease_alert_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	Metadata        *FriendTechKey `json:"metadata"`
}

func TrackingFriendTechKeyModelToResponse(m model.FriendTechKeyWatchlistItem, metadata *FriendTechKey) *FriendTechKeyWatchlistItemRespose {
	return &FriendTechKeyWatchlistItemRespose{
		Id:              m.Id,
		KeyAddress:      m.KeyAddress,
		ProfileId:       m.ProfileId,
		IncreaseAlertAt: m.IncreaseAlertAt,
		DecreaseAlertAt: m.DecreaseAlertAt,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
		Metadata:        metadata,
	}
}
