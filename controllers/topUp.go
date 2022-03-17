package controllers

import (
	"Project1-Account-Service-App/entities"
	"fmt"
	"gorm.io/gorm"
)

func TopUp(db *gorm.DB) (entities.TopUp, error) {
	var noHp string
	var topup entities.TopUp
	var nominal int64
	fmt.Println("Masukan No.Hp:")
	fmt.Scanln(&noHp)
	fmt.Println("Masukan Nominal:")
	fmt.Scanln(&nominal)

	var users entities.User
	err := db.Where("no_hp = ?", noHp).First(&users).Error
	if err != nil {
		// panic(tx.Error)
		return topup, err
	}

	users.Saldo += nominal
	db.Where("no_hp = ?", noHp).Updates(&users)

	topup = entities.TopUp{
		NominalTransfer: nominal,
		NoHpTujuan:      noHp,
		UserID:          users.ID,
	}
	db.Save(&topup)

	return topup, nil
}
