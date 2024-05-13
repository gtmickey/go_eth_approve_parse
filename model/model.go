package model

import (
	"gorm.io/gorm"
)

type Approval struct {
	gorm.Model
	Contract    string `json:"contract" gorm:"primaryKey"`
	Owner       string `json:"owner" gorm:"primaryKey"`
	Spender     string `json:"spender" gorm:"primaryKey"`
	Value       string `json:"value"`
	TxHash      string `json:"tx_hash"`
	BlockNumber uint64 `json:"block_number"`
	LogIndex    uint   `json:"log_index"`
}
