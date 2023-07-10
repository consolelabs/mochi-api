package model

import (
	"time"
)

type UserWalletWatchlist struct {
	Following []UserWalletWatchlistItem `json:"following"`
	Tracking  []UserWalletWatchlistItem `json:"tracking"`
	Copying   []UserWalletWatchlistItem `json:"copying"`
}

type UserWalletWatchlistItem struct {
	ProfileID   string       `json:"profile_id"`
	Address     string       `json:"address"`
	Alias       string       `json:"alias"`
	ChainType   ChainType    `json:"chain_type"`
	Type        TrackingType `json:"type"`
	IsOwner     bool         `json:"is_owner"`
	CreatedAt   time.Time    `json:"created_at"`
	NetWorth    float64      `json:"net_worth" gorm:"-"`
	FetchedData bool         `json:"fetched_data" gorm:"-"`
}

// TrackingType is the type of tracking
type TrackingType string

const (
	// TrackingTypeFollow represents that user want add the wallet to the watchlist
	TrackingTypeFollow TrackingType = "follow"
	// TrackingTypeTrack represents that user want to get notification when there is a new transaction
	TrackingTypeTrack TrackingType = "track"
	// TrackingTypeCopy represents that user want to copy trades from the wallet
	TrackingTypeCopy TrackingType = "copy"
)

// String returns the string representation of the tracking type
func (t TrackingType) String() string {
	return string(t)
}

// IsValid returns true if the tracking type is valid
func (t TrackingType) IsValid() bool {
	switch t {
	case TrackingTypeFollow, TrackingTypeTrack, TrackingTypeCopy:
		return true
	default:
		return false
	}
}

// ChainType is the type of chain
type ChainType string

const (
	// ChainTypeEvm represents the evm chain
	ChainTypeEvm ChainType = "evm"
	// ChainTypeRonin represents the ronin chain
	ChainTypeRonin ChainType = "ron"
	// ChainTypeSolana represents the solana chain
	ChainTypeSolana ChainType = "sol"
	// ChainTypeSui represents the sui chain
	ChainTypeSui ChainType = "sui"
	// ChainTypeAptos represents the aptos chain
	ChainTypeAptos ChainType = "apt"
)

// String returns the string representation of the chain type
func (t ChainType) String() string {
	return string(t)
}

// IsValid returns true if the chain type is valid
func (t ChainType) IsValid() bool {
	switch t {
	case ChainTypeEvm, ChainTypeRonin, ChainTypeSolana, ChainTypeSui, ChainTypeAptos:
		return true
	default:
		return false
	}
}
