package controllers

import (
	"Project1-Account-Service-App/entities"
	"fmt"

	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB) (entities.User, error) {
	var noHp string
	fmt.Println("Masukan No Hp:")
	fmt.Scanln(&noHp)

	fmt.Println("Pilih menu Update:")
	fmt.Println("1. Update Nama:")
	fmt.Println("2. Update Email:")
	fmt.Println("1. Update Password:")

	var menuUpdate, namaBaru, emailBaru, passwordBaru string
	fmt.Scanln(&menuUpdate)
	var users entities.User

	switch menuUpdate {
	case "1":

		fmt.Println("Masukan Nama Baru: ")
		fmt.Scanln(&namaBaru)

		users.Nama = namaBaru

	case "2":

		fmt.Println("Masukan Email Baru: ")
		fmt.Scanln(&emailBaru)

		users.Email = emailBaru

	case "3":

		fmt.Println("Masukan Password Baru: ")
		fmt.Scanln(&passwordBaru)

		users.Password = passwordBaru

	}

	db.Where("no_hp = ?", noHp).Updates(&users)
	return users, nil
}
