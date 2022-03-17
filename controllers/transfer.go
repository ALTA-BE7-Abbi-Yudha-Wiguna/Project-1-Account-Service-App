package controllers

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"withgorm/entities"
)

func Transfer(db *gorm.DB) (entities.Transfer, error) {
	var noHptujuan, noHppengirim string
	var nominal int64
	var transfer entities.Transfer

	fmt.Println("Masukan No.Hp Pengirim:")
	fmt.Scanln(&noHppengirim)
	fmt.Println("Masukan No.Hp Tujuan:")
	fmt.Scanln(&noHptujuan)
	fmt.Println("Masukan Nominal:")
	fmt.Scanln(&nominal)

	var pengirim entities.User
	err := db.Where("no_hp = ?", noHppengirim).First(&pengirim).Error
	if err != nil {
		// panic(tx.Error)
		return transfer, err
	}

	if pengirim.Saldo < nominal {

		return transfer, errors.New("Saldo tidak cukup")
	}
	var users entities.User
	err = db.Where("no_hp = ?", noHptujuan).First(&users).Error
	if err != nil {
		// panic(tx.Error)
		return transfer, err
	}

	users.Saldo += nominal
	db.Where("no_hp = ?", noHptujuan).Updates(&users)

	pengirim.Saldo -= nominal
	db.Where("no_hp = ?", noHppengirim).Updates(&pengirim)

	transfer = entities.Transfer{
		NominalTransfer: nominal,
		NoHpTujuan:      noHptujuan,
		NoHpPengirim:    noHppengirim,
		UserID:          users.ID,
	}

	db.Save(&transfer)

	return transfer, nil
}
