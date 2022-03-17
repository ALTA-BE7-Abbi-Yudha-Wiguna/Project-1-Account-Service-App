package main

import (
	"Project1-Account-Service-App/config"
	"Project1-Account-Service-App/entities"
	"fmt"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func Init() {
	dbConn = config.InitDB()
	InitialMigration()
}

func InitialMigration() {

	dbConn.AutoMigrate(entities.Transfer{})
	dbConn.AutoMigrate(entities.TopUp{})
	dbConn.AutoMigrate(entities.User{})

}

func main() {
	var hasil string
	Init()

	fmt.Println("Account Service App")
	fmt.Println("=====================")
	fmt.Println("Menu")
	fmt.Println("No 1. Registrasi")
	fmt.Println("No 2. Tampilkan Profile User")
	fmt.Println("No 3. Update User")
	fmt.Println("No 4. Delete User")
	fmt.Println("No 5. Top Up")
	fmt.Println("No 6. Transfer")
	fmt.Println("No 7. History Top Up")
	fmt.Println("No 8. History Transfer")
	fmt.Println("No 0. Keluar")
	fmt.Println("=====================")
	fmt.Println("Pillih Menu:")
	var menu string
	fmt.Scanln(&menu)

	switch menu {
	case "1":
	case "2":
	case "3":
	case "4":
	case "5":
	case "6":
	case "7":
	case "8":
	case "0":
	}

}
