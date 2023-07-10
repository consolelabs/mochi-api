package request

import (
	"strings"

	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/model/errors"
)

type WatchlistBaseRequest struct {
	ProfileID string `uri:"id" binding:"required" json:"-"`
}

type GetTrackingWalletsRequest struct {
	ProfileID string
	IsOwner   bool   `json:"-"`
	GuildID   string `json:"-"`
}

type GetOneWalletRequest struct {
	WatchlistBaseRequest
	AliasOrAddress string `uri:"address" binding:"required"`
}

type TrackWalletRequest struct {
	ProfileID string `json:"-"`
	Address   string `json:"address" binding:"required"`
	Alias     string `json:"alias"`
	ChainType string `json:"chain_type" binding:"required"`
	Type      string `json:"type" binding:"required"`
	IsOwner   bool   `json:"is_owner"`
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
}

func (req TrackWalletRequest) RequestToUserWalletWatchlistItemModel() (model.UserWalletWatchlistItem, error) {
	chainType := model.ChainType(strings.ToLower(req.ChainType))
	if !chainType.IsValid() {
		return model.UserWalletWatchlistItem{}, errors.ErrInvalidChainType
	}

	trackingType := model.TrackingType(strings.ToLower(req.Type))
	if !trackingType.IsValid() {
		return model.UserWalletWatchlistItem{}, errors.ErrInvalidTrackingType
	}

	return model.UserWalletWatchlistItem{
		ProfileID: req.ProfileID,
		Address:   req.Address,
		Alias:     req.Alias,
		ChainType: chainType,
		Type:      trackingType,
		IsOwner:   req.IsOwner,
	}, nil
}

type UntrackWalletRequest struct {
	WatchlistBaseRequest
	Address string `json:"address"`
	Alias   string `json:"alias"`
}

type ListWalletAssetsRequest struct {
	WatchlistBaseRequest
	Address string `uri:"address" binding:"required"`
	Type    string `uri:"type" binding:"required"`
}

type ListWalletTransactionsRequest struct {
	WatchlistBaseRequest
	Address string `uri:"address" binding:"required"`
	Type    string `uri:"type" binding:"required"`
}

type GenerateWalletVerificationRequest struct {
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
	UserID    string `json:"-"`
}

func (r *ListWalletAssetsRequest) Standardize() {
	addr := strings.ToLower(r.Address)
	if strings.HasPrefix(addr, "ronin:") {
		r.Address = "0x" + r.Address[6:]
		r.Type = "ron"
	}
}

func (r *ListWalletTransactionsRequest) Standardize() {
	addr := strings.ToLower(r.Address)
	if strings.HasPrefix(addr, "ronin:") {
		r.Address = "0x" + r.Address[6:]
		r.Type = "ron"
	}
}

func (r *GetOneWalletRequest) Standardize() {
	addr := strings.ToLower(r.AliasOrAddress)
	if strings.HasPrefix(addr, "ronin:") && len(addr) == 46 {
		r.AliasOrAddress = "0x" + r.AliasOrAddress[6:]
	}
}

func (r *TrackWalletRequest) Standardize() {
	addr := strings.ToLower(r.Address)
	if strings.HasPrefix(addr, "ronin:") && len(addr) == 46 {
		r.Address = "0x" + r.Address[6:]
		r.ChainType = "ron"
	}
}

type UpdateTrackingInfoRequest struct {
	WatchlistBaseRequest
	Address string `uri:"address" binding:"required" json:"-"`
	// Request body, only update the following fields
	Alias string `json:"alias"`
}
