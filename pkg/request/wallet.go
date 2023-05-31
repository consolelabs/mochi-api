package request

import (
	"strings"
)

type WalletBaseRequest struct {
	UserID string `uri:"id" binding:"required"`
}

type GetTrackingWalletsRequest struct {
	WalletBaseRequest
	IsOwner bool   `json:"-"`
	GuildID string `json:"-"`
}

type GetOneWalletRequest struct {
	WalletBaseRequest
	AliasOrAddress string `uri:"address" binding:"required"`
}

type TrackWalletRequest struct {
	UserID    string `json:"user_id" binding:"required"`
	Address   string `json:"address" binding:"required"`
	Alias     string `json:"alias"`
	Type      string `json:"type" binding:"required"`
	IsOwner   bool   `json:"is_owner"`
	ChannelID string `json:"channel_id"`
	MessageID string `json:"message_id"`
}

type UntrackWalletRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Address string `json:"address"`
	Alias   string `json:"alias"`
}

type ListWalletAssetsRequest struct {
	WalletBaseRequest
	Address string `uri:"address" binding:"required"`
	Type    string `uri:"type" binding:"required"`
}

type ListWalletTransactionsRequest struct {
	WalletBaseRequest
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
	}
}

func (r *ListWalletTransactionsRequest) Standardize() {
	addr := strings.ToLower(r.Address)
	if strings.HasPrefix(addr, "ronin:") {
		r.Address = "0x" + r.Address[6:]
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
	}
}
