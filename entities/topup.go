package entities

import "gorm.io/gorm"

type TopUp struct {
	gorm.Model

	NominalTransfer int64  `json:"nominaltransfer" from:"nama"`
	NoHpTujuan      string `json:"nohptujuan" from:"nohptujuan"`
	UserID          uint
}
