package controllers

import (
	"Project1-Account-Service-App/entities"
	"fmt"

	"gorm.io/gorm"
)

func Registrasi(db *gorm.DB) (entities.User, error) {
	var Nama, Email, Password, NoHP string

	fmt.Println("Masukkan Nama:")
	fmt.Scanln(&Nama)
	fmt.Println("Masukkan Email:")
	fmt.Scanln(&Email)
	fmt.Println("Masukkan Password:")
	fmt.Scanln(&Password)
	fmt.Println("Masukkan NoHP:")
	fmt.Scanln(&NoHP)

	var users entities.User
	users = entities.User{
		Nama:     Nama,
		Email:    Email,
		Password: Password,
		NoHP:     NoHP,
	}

	db.Create(&users)

	return users, nil
}
