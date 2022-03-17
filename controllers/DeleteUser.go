package controllers

import (
	"Project1-Account-Service-App/entities"
	"fmt"

	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB) (entities.User, error) {
	var noHp string
	fmt.Println("Masukan No Hp:")
	fmt.Scanln(&noHp)

	var users entities.User
	err := db.Where("no_hp = ?", noHp).Delete(&users).Error
	if err != nil {
		// panic(tx.Error)
		return users, err
	}
	return users, nil
}
