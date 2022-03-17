package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitDB() *gorm.DB {
	connectionString := os.Getenv("CONNECTION_STRING")

	//var err error
	var DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return DB
}
