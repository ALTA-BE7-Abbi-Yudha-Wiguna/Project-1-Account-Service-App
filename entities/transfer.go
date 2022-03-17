package entities

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model

	NominalTransfer int64  `json:"nominaltransfer" from:"nama"`
	NoHpTujuan      string `json:"nohptujuan" from:"nohptujuan"`
	NoHpPengirim    string `json:"nohppengirim" from:"nohppengirim"`
	UserID          uint
}
