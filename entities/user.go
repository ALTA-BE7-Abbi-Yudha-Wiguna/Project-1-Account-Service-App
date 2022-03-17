package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	NoHP     string `json:"nohp" form:"nohP" gorm:"unique"`
	Saldo    int64  `json:"saldo" form:"saldo"`
	Transfer []Transfer
	TopUp    []TopUp
}
