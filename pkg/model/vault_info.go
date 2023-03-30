package model

import "time"

type VaultInfo struct {
	Id              int64     `json:"id"`
	Description     string    `json:"description"`
	ModStep         string    `json:"mod_step"`
	NormalStep      string    `json:"normal_step"`
	InstructionLink string    `json:"instruction_link"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
