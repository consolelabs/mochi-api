package model

type NotificationGroup string

const (
	NotificationGroupWallet    NotificationGroup = "wallet"
	NotificationGroupCommunity NotificationGroup = "community"
	NotificationGroupApp       NotificationGroup = "app"
)

// "key" must follow this format [role]_[action]_[status]
// roles: send/receive/*
// action: tx action
// status: tx status
// e.g. receive_tip_success
type NotificationFlag struct {
	Key         string            `json:"key"`
	Group       NotificationGroup `json:"group"`
	Description string            `json:"description"`
}
