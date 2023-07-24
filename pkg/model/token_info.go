package model

import "gorm.io/datatypes"

type TokenInfo struct {
	Id   string         `json:"id"`
	Data datatypes.JSON `json:"data"`
}
